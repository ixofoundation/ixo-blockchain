package keeper

import (
	"fmt"
	"strconv"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ixofoundation/ixo-blockchain/v7/ixomath"
	"github.com/ixofoundation/ixo-blockchain/v7/x/liquidstake/types"
)

// TryRedelegation attempts a single redelegation, atomically through a cached
// context so a transitive-redelegation rejection does not leak partial state
// into the parent transaction.
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
	if _, err = k.stakingKeeper.ValidateUnbondAmount(ctx, re.Delegator, srcVal, re.Amount); err != nil {
		return time.Time{}, fmt.Errorf("failed to validate unbond amount: %w", err)
	}

	// On the final transfer-out from a source validator, redelegate the
	// validator's full liquid-token balance (not the requested amount,
	// which may have been rounded down) so no dust is left behind.
	amt := re.Amount
	if re.Last {
		amt = re.SrcValidator.GetLiquidTokens(ctx, k.stakingKeeper, re.Delegator, false)
	}
	cachedCtx, writeCache := ctx.CacheContext()
	completionTime, err = k.RedelegateWithCap(cachedCtx, re.Delegator, srcVal, dstVal, amt)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to begin redelegation: %w", err)
	}
	writeCache()
	return completionTime, nil
}

// Rebalance reshuffles delegation amounts among the pool's liquid validators
// so each one's stake matches its target weight. Acts only when the per-pool
// max gap exceeds rebalancingTrigger * totalLiquidTokens.
//
// Returns every redelegation it tried (including failed ones with .Error
// populated), useful for tests and emitted-event accounting.
func (k Keeper) Rebalance(
	ctx sdk.Context,
	poolID string,
	proxyAcc sdk.AccAddress,
	liquidVals types.LiquidValidators,
	whitelistedValsMap types.WhitelistedValsMap,
	rebalancingTrigger math.LegacyDec,
) (redelegations []types.Redelegation) {
	totalLiquidTokens, liquidTokenMap := liquidVals.TotalLiquidTokens(ctx, k.stakingKeeper, proxyAcc, false)
	if !totalLiquidTokens.IsPositive() {
		return redelegations
	}

	weightMap, totalWeight := k.GetWeightMap(ctx, liquidVals, whitelistedValsMap)
	if !totalWeight.IsPositive() {
		return redelegations
	}

	// Per-validator target = totalLiquidTokens * weight / totalWeight.
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
	// Add the rounding crumb to the first non-zero-target validator.
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
		minVal, maxVal, amountNeeded, last := liquidVals.MinMaxGap(targetMap, liquidTokenMap)
		// Skip if the gap is below the trigger on the very first iteration —
		// avoids churn for routine reward fluctuations.
		if amountNeeded.IsZero() || (i == 0 && !amountNeeded.GT(rebalancingThresholdAmt)) {
			break
		}

		// Locally update liquidTokenMap so the next MinMaxGap reflects the
		// transfer we're about to attempt.
		liquidTokenMap[maxVal.OperatorAddress] = liquidTokenMap[maxVal.OperatorAddress].Sub(amountNeeded)
		liquidTokenMap[minVal.OperatorAddress] = liquidTokenMap[minVal.OperatorAddress].Add(amountNeeded)

		redelegation := types.Redelegation{
			Delegator:    proxyAcc,
			SrcValidator: maxVal,
			DstValidator: minVal,
			Amount:       amountNeeded,
			Last:         last,
		}
		if _, err := k.TryRedelegation(ctx, redelegation); err != nil {
			redelegation.Error = err
			failCount++
			k.Logger(ctx).Error(
				"redelegation failed",
				"pool_id", poolID,
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
		if err := ctx.EventManager().EmitTypedEvent(
			&types.RebalancedLiquidStakeEvent{
				PoolId:                poolID,
				Delegator:             proxyAcc.String(),
				RedelegationCount:     uint32(len(redelegations)),
				RedelegationFailCount: uint32(failCount),
			},
		); err != nil {
			k.Logger(ctx).Error("failed to emit rebalanced liquid stake event", "pool_id", poolID, "error", err)
		}
		k.Logger(ctx).Info(
			"Rebalance",
			"pool_id", poolID,
			"module", types.ModuleName,
			"delegator", proxyAcc.String(),
			"redelegation_count", strconv.Itoa(len(redelegations)),
			"redelegation_fail_count", strconv.Itoa(failCount),
		)
	}
	return redelegations
}

// UpdateLiquidValidatorSet refreshes a single pool's LiquidValidator set so
// any newly whitelisted active validators are persisted, and (when
// redelegate=true) attempts to rebalance towards the configured weights.
//
// Returns redelegations for testing visibility; production callers (BeginBlock,
// epoch hook) discard the return value.
func (k Keeper) UpdateLiquidValidatorSet(ctx sdk.Context, p types.Pool, redelegate bool) (redelegations []types.Redelegation) {
	liquidValidators := k.GetAllLiquidValidatorsForPool(ctx, p.PoolId)
	liquidValsMap := liquidValidators.Map()
	whitelistedValsMap := p.WhitelistedValsMap()

	for _, wv := range p.WhitelistedValidators {
		if _, ok := liquidValsMap[wv.ValidatorAddress]; !ok {
			lv := types.LiquidValidator{OperatorAddress: wv.ValidatorAddress}
			if k.IsActiveLiquidValidator(ctx, lv, whitelistedValsMap) {
				k.SetLiquidValidator(ctx, p.PoolId, lv)
				liquidValidators = append(liquidValidators, lv)
				if err := ctx.EventManager().EmitTypedEvent(
					&types.AddLiquidValidatorEvent{PoolId: p.PoolId, Validator: lv.OperatorAddress},
				); err != nil {
					k.Logger(ctx).Error("failed to emit add liquid validator event", "pool_id", p.PoolId, "error", err)
				}
			}
		}
	}

	// Tombstone-state changes are also handled inside Rebalance via
	// IsActiveLiquidValidator / GetWeightMap.
	if redelegate {
		redelegations = k.Rebalance(
			ctx, p.PoolId, p.GetProxyAccount(),
			liquidValidators, whitelistedValsMap,
			types.RebalancingTrigger,
		)
		// Inactive liquid validators stay in the set; LiquidUnstake drains
		// them ahead of healthy ones (see PrioritiseInactiveLiquidValidators).
		return redelegations
	}
	return nil
}

// AutocompoundStakingRewards withdraws all staking rewards for a pool, takes
// the autocompound fee, distributes the configured weighted-rewards portion,
// and re-stakes whatever remains. Each step uses a cached context so a
// failure in one phase does not corrupt prior state.
func (k Keeper) AutocompoundStakingRewards(ctx sdk.Context, p types.Pool, whitelistedValsMap types.WhitelistedValsMap) {
	proxyAcc := p.GetProxyAccount()
	k.WithdrawLiquidRewards(ctx, proxyAcc)

	// skip when no active liquid validator
	activeVals := k.GetActiveLiquidValidatorsForPool(ctx, p.PoolId, whitelistedValsMap)
	if len(activeVals) == 0 {
		return
	}

	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}

	// Source of compoundable funds: every spendable bondDenom token in the
	// proxy account (rewards just withdrawn plus any leftover from prior
	// epochs).
	proxyAccBalance := k.GetProxyAccBalance(ctx, proxyAcc)
	autoCompoundableAmount := proxyAccBalance.Amount

	autocompoundFee := sdk.NewCoin(bondDenom, math.ZeroInt())
	if !p.AutocompoundFeeRate.IsZero() && autoCompoundableAmount.IsPositive() {
		autocompoundFee = sdk.NewCoin(bondDenom, p.AutocompoundFeeRate.MulInt(autoCompoundableAmount).TruncateInt())
	}
	delegatableAmount := autoCompoundableAmount.Sub(autocompoundFee.Amount)

	// Distribute weighted rewards first so receivers cannot be diluted by
	// re-staked amounts that haven't yet earned them.
	rewardsCoin := sdk.NewCoin(bondDenom, delegatableAmount)
	totalDistributedAmount, err := k.DistributeWeightedRewards(ctx, proxyAcc, rewardsCoin, p.WeightedRewardsReceivers)
	if err != nil {
		// Don't panic in BeginBlock; next epoch will retry.
		k.Logger(ctx).Error("failed to distribute weighted rewards", "pool_id", p.PoolId, "error", err)
		return
	}
	delegatableAmount = delegatableAmount.Sub(totalDistributedAmount)

	// if there are still delegatable amount, re-stake the accumulated rewards
	if delegatableAmount.IsPositive() {
		cachedCtx, writeCache := ctx.CacheContext()
		if err = k.LiquidDelegate(cachedCtx, proxyAcc, activeVals, delegatableAmount, whitelistedValsMap); err != nil {
			k.Logger(ctx).Error("failed to re-stake the accumulated rewards", "pool_id", p.PoolId, "error", err)
			return
			// hitting global liquid cap surfaces here too — let next epoch retry
		}
		writeCache()
	}

	// Pay out the autocompound fee last so it does not consume from the
	// re-staking budget if delegation fails.
	feeAccountAddr, err := sdk.AccAddressFromBech32(p.FeeAccountAddress)
	if err != nil {
		k.Logger(ctx).Error("invalid fee_account_address", "pool_id", p.PoolId, "error", err)
		return
	}
	if autocompoundFee.Amount.IsPositive() {
		if err := k.bankKeeper.SendCoins(ctx, proxyAcc, feeAccountAddr, sdk.NewCoins(autocompoundFee)); err != nil {
			k.Logger(ctx).Error("failed to send autocompound fee to fee account", "pool_id", p.PoolId, "error", err)
			return
		}
	}

	if err := ctx.EventManager().EmitTypedEvent(
		&types.AutocompoundStakingRewardsEvent{
			PoolId:                p.PoolId,
			Delegator:             proxyAcc.String(),
			TotalAmount:           sdk.NewCoin(bondDenom, autoCompoundableAmount),
			FeeAmount:             autocompoundFee,
			RedelegateAmount:      sdk.NewCoin(bondDenom, delegatableAmount),
			WeightedRewardsAmount: sdk.NewCoin(bondDenom, totalDistributedAmount),
		},
	); err != nil {
		k.Logger(ctx).Error("failed to emit autocompound staking rewards event", "pool_id", p.PoolId, "error", err)
		return
	}
	k.Logger(ctx).Info(
		"AutocompoundStakingRewards",
		"pool_id", p.PoolId,
		"module", types.ModuleName,
		"delegator", proxyAcc.String(),
		"autocompound_amount", delegatableAmount.String(),
		"autocompound_fee", autocompoundFee.String(),
	)
}

// DistributeWeightedRewards splits rewardsCoin among receivers proportionally
// to their weights, sending from sender. Returns the aggregate amount sent.
// Receivers' weights are pre-validated to sum to <= 1; the unused remainder
// is left for the autocompounder to re-stake.
func (k Keeper) DistributeWeightedRewards(ctx sdk.Context, sender sdk.AccAddress, rewardsCoin sdk.Coin, weightedRewardsReceivers []types.WeightedAddress) (ixomath.Int, error) {
	// counter for total distributed amount, used instead of rewardsCoin to avoid rounding discrepancies.
	totalDistributedAmount := ixomath.ZeroInt()
	if rewardsCoin.IsZero() || len(weightedRewardsReceivers) == 0 {
		return totalDistributedAmount, nil
	}

	for _, w := range weightedRewardsReceivers {
		portion, err := getProportions(rewardsCoin, w.Weight)
		if err != nil {
			return ixomath.Int{}, err
		}
		if !portion.Amount.IsPositive() {
			continue
		}
		receiverAddr, err := sdk.AccAddressFromBech32(w.Address)
		if err != nil {
			return ixomath.Int{}, err
		}
		if err := k.bankKeeper.SendCoins(ctx, sender, receiverAddr, sdk.NewCoins(portion)); err != nil {
			return ixomath.Int{}, err
		}
		totalDistributedAmount = totalDistributedAmount.Add(portion.Amount)
	}
	return totalDistributedAmount, nil
}

// getProportions gets the balance of the input coin returns a coin according to the
// allocation ratio. Returns error if ratio is greater than 1.
func getProportions(coin sdk.Coin, ratio ixomath.Dec) (sdk.Coin, error) {
	if ratio.GT(ixomath.OneDec()) {
		return sdk.Coin{}, types.ErrRatioMoreThanOne
	}
	return sdk.NewCoin(coin.Denom, coin.Amount.ToLegacyDec().Mul(ratio).TruncateInt()), nil
}
