package keeper

import (
	"fmt"
	"strconv"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ixofoundation/ixo-blockchain/v5/ixomath"
	"github.com/ixofoundation/ixo-blockchain/v5/x/liquidstake/types"
)

// GetProxyAccBalance returns the available spendable balance of the proxy account for the native token.
func (k Keeper) GetProxyAccBalance(ctx sdk.Context, proxyAcc sdk.AccAddress) (balance sdk.Coin) {
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	return sdk.NewCoin(bondDenom, k.bankKeeper.SpendableCoins(ctx, proxyAcc).AmountOf(bondDenom))
}

// TryRedelegation attempts redelegation, which is applied only when successful through cached context because there is a constraint that fails if already receiving redelegation.
func (k Keeper) TryRedelegation(ctx sdk.Context, re types.Redelegation) (completionTime time.Time, err error) {
	dstVal := re.DstValidator.GetOperator()
	srcVal := re.SrcValidator.GetOperator()

	// check the source validator already has receiving transitive redelegation
	hasReceiving, err := k.stakingKeeper.HasReceivingRedelegation(ctx, re.Delegator, srcVal)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to check has receiving redelegation: %w", err)
	}
	if hasReceiving {
		return time.Time{}, stakingtypes.ErrTransitiveRedelegation
	}

	// calculate delShares from tokens with validation
	_, err = k.stakingKeeper.ValidateUnbondAmount(
		ctx, re.Delegator, srcVal, re.Amount,
	)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to validate unbond amount: %w", err)
	}

	// when last, full redelegation of shares from delegation
	amt := re.Amount
	if re.Last {
		amt = re.SrcValidator.GetLiquidTokens(ctx, k.stakingKeeper, false)
	}
	cachedCtx, writeCache := ctx.CacheContext()
	completionTime, err = k.RedelegateWithCap(cachedCtx, re.Delegator, srcVal, dstVal, amt)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to begin redelegation: %w", err)
	}
	writeCache()
	return completionTime, nil
}

// Rebalance argument liquidVals containing ValidatorStatusActive which is containing just added on whitelist(liquidToken 0) and ValidatorStatusInactive to delist
func (k Keeper) Rebalance(
	ctx sdk.Context,
	proxyAcc sdk.AccAddress,
	liquidVals types.LiquidValidators,
	whitelistedValsMap types.WhitelistedValsMap,
	rebalancingTrigger math.LegacyDec,
) (redelegations []types.Redelegation) {
	totalLiquidTokens, liquidTokenMap := liquidVals.TotalLiquidTokens(ctx, k.stakingKeeper, false)
	if !totalLiquidTokens.IsPositive() {
		return redelegations
	}

	weightMap, totalWeight := k.GetWeightMap(ctx, liquidVals, whitelistedValsMap)

	// no active liquid validators
	if !totalWeight.IsPositive() {
		return redelegations
	}

	// calculate rebalancing target map
	targetMap := map[string]math.Int{}
	totalTargetMap := math.ZeroInt()
	for _, val := range liquidVals {
		targetMap[val.OperatorAddress] = totalLiquidTokens.Mul(weightMap[val.OperatorAddress]).Quo(totalWeight)
		totalTargetMap = totalTargetMap.Add(targetMap[val.OperatorAddress])
	}
	crumb := totalLiquidTokens.Sub(totalTargetMap)
	if !totalTargetMap.IsPositive() {
		return redelegations
	}
	// crumb to first non zero liquid validator
	for _, val := range liquidVals {
		if targetMap[val.OperatorAddress].IsPositive() {
			targetMap[val.OperatorAddress] = targetMap[val.OperatorAddress].Add(crumb)
			break
		}
	}

	failCount := 0
	rebalancingThresholdAmt := rebalancingTrigger.Mul(math.LegacyNewDecFromInt(totalLiquidTokens)).TruncateInt()
	redelegations = make([]types.Redelegation, 0, liquidVals.Len())

	for i := 0; i < liquidVals.Len(); i++ {
		// get min, max of liquid token gap
		minVal, maxVal, amountNeeded, last := liquidVals.MinMaxGap(targetMap, liquidTokenMap)
		if amountNeeded.IsZero() || (i == 0 && !amountNeeded.GT(rebalancingThresholdAmt)) {
			break
		}

		// sync liquidTokenMap applied rebalancing
		liquidTokenMap[maxVal.OperatorAddress] = liquidTokenMap[maxVal.OperatorAddress].Sub(amountNeeded)
		liquidTokenMap[minVal.OperatorAddress] = liquidTokenMap[minVal.OperatorAddress].Add(amountNeeded)

		// try redelegation from max validator to min validator
		redelegation := types.Redelegation{
			Delegator:    proxyAcc,
			SrcValidator: maxVal,
			DstValidator: minVal,
			Amount:       amountNeeded,
			Last:         last,
		}

		_, err := k.TryRedelegation(ctx, redelegation)
		if err != nil {
			redelegation.Error = err
			failCount++

			k.Logger(ctx).Error(
				"redelegation failed",
				"delegator", proxyAcc.String(),
				"src_validator", maxVal.OperatorAddress,
				"dst_validator", minVal.OperatorAddress,
				"amount", amountNeeded.String(),
				"error", err.Error(),
			)
		}

		redelegations = append(redelegations, redelegation)
	}

	if len(redelegations) != 0 {
		// Emit the events
		if err := ctx.EventManager().EmitTypedEvent(
			&types.RebalancedLiquidStakeEvent{
				Delegator:             proxyAcc.String(),
				RedelegationCount:     strconv.Itoa(len(redelegations)),
				RedelegationFailCount: strconv.Itoa(failCount),
			},
		); err != nil {
			k.Logger(ctx).Error(
				"failed to emit rebalanced liquid stake event",
				"error", err,
			)
		}
		k.Logger(ctx).Info(
			"Rebalance",
			"module", types.ModuleName,
			"delegator", proxyAcc.String(),
			"redelegation_count", strconv.Itoa(len(redelegations)),
			"redelegation_fail_count", strconv.Itoa(failCount),
		)
	}

	return redelegations
}

func (k Keeper) UpdateLiquidValidatorSet(ctx sdk.Context, redelegate bool) (redelegations []types.Redelegation) {
	params := k.GetParams(ctx)
	liquidValidators := k.GetAllLiquidValidators(ctx)
	liquidValsMap := liquidValidators.Map()
	whitelistedValsMap := types.GetWhitelistedValsMap(params.WhitelistedValidators)

	// Set Liquid validators for added whitelist validators
	for _, wv := range params.WhitelistedValidators {
		if _, ok := liquidValsMap[wv.ValidatorAddress]; !ok {
			lv := types.LiquidValidator{
				OperatorAddress: wv.ValidatorAddress,
			}
			if k.IsActiveLiquidValidator(ctx, lv, whitelistedValsMap) {
				k.SetLiquidValidator(ctx, lv)
				liquidValidators = append(liquidValidators, lv)
				// Emit the events
				if err := ctx.EventManager().EmitTypedEvent(
					&types.AddLiquidValidatorEvent{
						Validator: lv.OperatorAddress,
					},
				); err != nil {
					k.Logger(ctx).Error(
						"failed to emit add liquid validator event",
						"error", err,
					)
				}
			}
		}
	}

	// rebalancing based updated liquid validators status with threshold, try by cachedCtx
	// tombstone status also handled on Rebalance
	if redelegate {
		redelegations = k.Rebalance(
			ctx,
			types.LiquidStakeProxyAcc,
			liquidValidators,
			whitelistedValsMap,
			types.RebalancingTrigger,
		)

		// if there are inactive liquid validators, do not unbond,
		// instead let validator selection and rebalancing take care of it.

		return redelegations
	}
	return nil
}

// AutocompoundStakingRewards withdraws staking rewards and re-stakes when over threshold.
func (k Keeper) AutocompoundStakingRewards(ctx sdk.Context, whitelistedValsMap types.WhitelistedValsMap) {
	// withdraw rewards of LiquidStakeProxyAcc
	k.WithdrawLiquidRewards(ctx, types.LiquidStakeProxyAcc)

	// skip when no active liquid validator
	activeVals := k.GetActiveLiquidValidators(ctx, whitelistedValsMap)
	if len(activeVals) == 0 {
		return
	}

	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}

	// use all available funds in the proxy account as autocompoundable amount
	proxyAccBalance := k.GetProxyAccBalance(ctx, types.LiquidStakeProxyAcc)
	autoCompoundableAmount := proxyAccBalance.Amount

	// calculate autocompounding fee
	params := k.GetParams(ctx)
	autocompoundFee := sdk.NewCoin(bondDenom, math.ZeroInt())
	if !params.AutocompoundFeeRate.IsZero() && autoCompoundableAmount.IsPositive() {
		autocompoundFee = sdk.NewCoin(
			bondDenom,
			params.AutocompoundFeeRate.MulInt(autoCompoundableAmount).TruncateInt(),
		)
	}

	// delegatableAmount is the autoCompoundableAmount minus the autocompoundFee
	delegatableAmount := autoCompoundableAmount.Sub(autocompoundFee.Amount)

	// distribute rewards to weighted rewards receivers first
	rewardsCoin := sdk.NewCoin(bondDenom, delegatableAmount)
	totalDistributedAmount, err := k.DistributeWeightedRewards(ctx, rewardsCoin, params.WeightedRewardsReceivers)
	if err != nil {
		// skip errors as don't want to panic in beginblock. Can see logs for fails
		// and next epoch will attempt distribution of rewards again
		k.Logger(ctx).Error(
			"failed to distribute weighted rewards",
			"error", err,
		)
		return
	}

	delegatableAmount = delegatableAmount.Sub(totalDistributedAmount)

	// if there are still delegatable amount, re-stake the accumulated rewards
	if delegatableAmount.IsPositive() {
		// re-staking of the accumulated rewards
		cachedCtx, writeCache := ctx.CacheContext()
		err = k.LiquidDelegate(cachedCtx, types.LiquidStakeProxyAcc, activeVals, delegatableAmount, whitelistedValsMap)
		if err != nil {
			k.Logger(ctx).Error(
				"failed to re-stake the accumulated rewards",
				"error", err,
			)
			return
			// skip errors as they might occur due to reaching global liquid cap
		}
		writeCache()
	}

	// move autocompounding fee from the balance to fee account
	feeAccountAddr := sdk.MustAccAddressFromBech32(params.FeeAccountAddress)
	err = k.bankKeeper.SendCoins(ctx, types.LiquidStakeProxyAcc, feeAccountAddr, sdk.NewCoins(autocompoundFee))
	if err != nil {
		k.Logger(ctx).Error(
			"failed to send autocompound fee to fee account",
			"error", err,
		)
		return
	}

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.AutocompoundStakingRewardsEvent{
			Delegator:             types.LiquidStakeProxyAcc.String(),
			TotalAmount:           autoCompoundableAmount.String(),
			FeeAmount:             autocompoundFee.String(),
			RedelegateAmount:      delegatableAmount.String(),
			WeightedRewardsAmount: totalDistributedAmount.String(),
		},
	); err != nil {
		k.Logger(ctx).Error(
			"failed to emit autocompound staking rewards event",
			"error", err,
		)
		return
	}
	k.Logger(ctx).Info(
		"AutocompoundStakingRewards",
		"module", types.ModuleName,
		"delegator", types.LiquidStakeProxyAcc.String(),
		"autocompound_amount", delegatableAmount.String(),
		"autocompound_fee", autocompoundFee.String(),
	)
}

// DistributeWeightedRewards distributes the staking rewards to the given list of weightedRewardsReceivers based on their
// weights, if there are any. Returns the total amount distributed.
func (k Keeper) DistributeWeightedRewards(ctx sdk.Context, rewardsCoin sdk.Coin, weightedRewardsReceivers []types.WeightedAddress) (ixomath.Int, error) {
	// counter for total distributed amount, used instead of rewardsCoin to avoid rounding discrepancies.
	totalDistributedAmount := ixomath.ZeroInt()

	// if rewardsCoin is empty, or no weighted rewards receivers provided, return 0
	if rewardsCoin.IsZero() || len(weightedRewardsReceivers) == 0 {
		return totalDistributedAmount, nil
	}

	// allocate weighted rewards to addresses by weight
	for _, w := range weightedRewardsReceivers {
		weightedRewardPortionCoin, err := getProportions(rewardsCoin, w.Weight)
		if err != nil {
			return ixomath.Int{}, err
		}

		// distribute impact rewards to the address
		weightedRewardsAddr, err := sdk.AccAddressFromBech32(w.Address)
		if err != nil {
			return ixomath.Int{}, err
		}
		err = k.bankKeeper.SendCoins(ctx, types.LiquidStakeProxyAcc, weightedRewardsAddr, sdk.NewCoins(weightedRewardPortionCoin))
		if err != nil {
			return ixomath.Int{}, err
		}

		// update total distributed amount
		totalDistributedAmount = totalDistributedAmount.Add(weightedRewardPortionCoin.Amount)
	}

	return totalDistributedAmount, nil
}

// getProportions gets the balance of the inout coin returns a coin according to the
// allocation ratio. Returns error if ratio is greater than 1.
func getProportions(coin sdk.Coin, ratio ixomath.Dec) (sdk.Coin, error) {
	if ratio.GT(ixomath.OneDec()) {
		return sdk.Coin{}, types.ErrRatioMoreThanOne
	}
	return sdk.NewCoin(coin.Denom, coin.Amount.ToLegacyDec().Mul(ratio).TruncateInt()), nil
}
