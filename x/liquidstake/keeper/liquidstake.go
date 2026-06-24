package keeper

import (
	"sort"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

// ----------------------------------------------------------------------------
// Pause helpers
// ----------------------------------------------------------------------------

// PoolEffectivelyPaused reports whether the given pool's stake-related
// operations should be blocked. A pool is paused if EITHER the global
// ModuleParams.module_paused flag is true OR the pool's own .Paused flag
// is true.
func (k Keeper) PoolEffectivelyPaused(ctx sdk.Context, p types.Pool) bool {
	return k.IsModulePaused(ctx) || p.Paused
}

// ----------------------------------------------------------------------------
// Per-pool NetAmountState
// ----------------------------------------------------------------------------

// GetNetAmountStateForPool computes a pool's NetAmountState from current chain
// state (delegations held by its proxy account, unbonding entries, pending
// rewards, and spendable balance). Never persisted — recomputed every call.
func (k Keeper) GetNetAmountStateForPool(ctx sdk.Context, p types.Pool) types.NetAmountState {
	proxyAcc := p.GetProxyAccount()
	totalRemainingRewards, totalDelShares, totalLiquidTokens := k.CheckDelegationStates(ctx, proxyAcc)

	totalUnbondingBalance := math.ZeroInt()
	ubds, err := k.stakingKeeper.GetAllUnbondingDelegations(ctx, proxyAcc)
	if err != nil {
		panic(err) // should never happen
	}
	for _, ubd := range ubds {
		for _, entry := range ubd.Entries {
			// Use slashing-applied balance, not initial balance.
			totalUnbondingBalance = totalUnbondingBalance.Add(entry.Balance)
		}
	}

	nas := types.NetAmountState{
		StkixoTotalSupply:     k.bankKeeper.GetSupply(ctx, p.LiquidBondDenom).Amount,
		TotalDelShares:        totalDelShares,
		TotalLiquidTokens:     totalLiquidTokens,
		TotalRemainingRewards: totalRemainingRewards,
		TotalUnbondingBalance: totalUnbondingBalance,
		ProxyAccBalance:       k.GetProxyAccBalance(ctx, proxyAcc).Amount,
	}
	nas.NetAmount = nas.CalcNetAmount()
	// StakeRate is fixed at 1.0 by the pool's mint design (see
	// NativeTokenToStkIXO). UnstakeRate captures all value accrual.
	nas.StakeRate = math.LegacyNewDecFromInt(types.NativeTokenToStkIXO(math.NewInt(1)))
	nas.UnstakeRate = types.StkIXOToNativeToken(math.NewInt(1), nas.StkixoTotalSupply, nas.NetAmount)
	return nas
}

// GetProxyAccBalance returns the proxy account's spendable balance in the
// chain's bond denom. Used as the source of newly-staked or
// not-yet-redelegated funds.
func (k Keeper) GetProxyAccBalance(ctx sdk.Context, proxyAcc sdk.AccAddress) sdk.Coin {
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	return sdk.NewCoin(bondDenom, k.bankKeeper.SpendableCoins(ctx, proxyAcc).AmountOf(bondDenom))
}

// ----------------------------------------------------------------------------
// LiquidStake
// ----------------------------------------------------------------------------

// LiquidStake delegates stakingCoin (denominated in bondDenom) through the
// pool's proxy account to its active whitelisted validators in proportion
// to their target weights, then mints the pool's LST denom to liquidStaker.
//
// Permission model preserved from pre-v7: only the pool's
// whitelist_admin_address may invoke LiquidStake.
func (k Keeper) LiquidStake(
	ctx sdk.Context, p types.Pool, liquidStaker sdk.AccAddress, stakingCoin sdk.Coin,
) (stkIXOMintAmount math.Int, err error) {
	moduleParams := k.GetModuleParams(ctx)

	if k.PoolEffectivelyPaused(ctx, p) {
		return math.ZeroInt(), types.ErrPoolPaused
	}

	// only allow the whitelistedAdminAddress to do liquid staking, this might be removed later
	if p.WhitelistAdminAddress != liquidStaker.String() {
		return math.ZeroInt(), types.ErrRestrictedToWhitelistedAdminAddress
	}

	// check minimum liquid stake amount
	if stakingCoin.Amount.LT(moduleParams.MinLiquidStakeAmount) {
		return math.ZeroInt(), types.ErrLessThanMinLiquidStakeAmount
	}

	// check bond denomination
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	if stakingCoin.Denom != bondDenom {
		return math.ZeroInt(), errorsmod.Wrapf(
			types.ErrInvalidBondDenom, "invalid coin denomination: got %s, expected %s", stakingCoin.Denom, bondDenom,
		)
	}

	whitelistedValsMap := p.WhitelistedValsMap()
	activeVals := k.GetActiveLiquidValidatorsForPool(ctx, p.PoolId, whitelistedValsMap)
	if activeVals.Len() == 0 {
		return math.ZeroInt(), types.ErrActiveLiquidValidatorsNotExists
	}

	totalActiveWeight := activeVals.TotalWeight(whitelistedValsMap)
	activeWeightQuorum := math.LegacyNewDecFromInt(totalActiveWeight).
		Quo(math.LegacyNewDecFromInt(types.TotalValidatorWeight))
	if activeWeightQuorum.LT(types.ActiveLiquidValidatorsWeightQuorum) {
		k.Logger(ctx).Error(
			"active liquid validators weight quorum not reached",
			"pool_id", p.PoolId,
			"active_weight_quorum", activeWeightQuorum.String(),
			"min_active_weight_quorum", types.ActiveLiquidValidatorsWeightQuorum.String(),
		)
		return math.ZeroInt(), errorsmod.Wrapf(
			types.ErrActiveLiquidValidatorsWeightQuorumNotReached, "%s < %s",
			activeWeightQuorum.String(), types.ActiveLiquidValidatorsWeightQuorum.String(),
		)
	}

	// NetAmountState must be sampled BEFORE moving the staker's coin into the
	// proxy account, otherwise the freshly-arrived balance would inflate
	// nas.NetAmount and skew the mint ratio.
	nas := k.GetNetAmountStateForPool(ctx, p)

	// send staking coin to liquid stake proxy account to proxy delegation, need sufficient spendable balances
	proxyAcc := p.GetProxyAccount()
	if err := k.bankKeeper.SendCoins(ctx, liquidStaker, proxyAcc, sdk.NewCoins(stakingCoin)); err != nil {
		return math.ZeroInt(), err
	}

	// MintAmount is fixed at 1:1 (NativeTokenToStkIXO returns its input
	// unchanged); value accrual is captured on the unstake side. The
	// infinite-c-value guard still catches the pathological case where
	// supply is positive but every backing token has been slashed or
	// withdrawn — leaving stk holders unable to redeem.
	if nas.StkixoTotalSupply.IsPositive() && nas.NetAmount.IsZero() {
		// this case must not be reachable, consider stopping module for investigation
		// c_value -> inf
		k.Logger(ctx).Error(
			"infinite c value", "pool_id", p.PoolId, "net_amount_state", nas.String(),
		)
		return math.ZeroInt(), types.ErrInsufficientProxyAccBalance
	}
	stkIXOMintAmount = types.NativeTokenToStkIXO(stakingCoin.Amount)
	if !stkIXOMintAmount.IsPositive() {
		return math.ZeroInt(), types.ErrTooSmallLiquidStakeAmount
	}

	// mint on module acc and send.
	//
	// Wrap the SDK context with AuthorizeLSTMintContext so the
	// bank MintCoinsRestriction installed in app/keepers/keepers.go
	// allows the mint despite p.LiquidBondDenom being a pool-claimed
	// denom. Any other module attempting to mint into this denom would
	// lack the sentinel and be rejected.
	mintCoin := sdk.NewCoins(sdk.NewCoin(p.LiquidBondDenom, stkIXOMintAmount))
	if err := k.bankKeeper.MintCoins(types.AuthorizeLSTMintContext(ctx), types.ModuleName, mintCoin); err != nil {
		return stkIXOMintAmount, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidStaker, mintCoin); err != nil {
		k.Logger(ctx).Error("failed to send minted coins to liquid staker", "pool_id", p.PoolId, "error", err)
		return stkIXOMintAmount, err
	}

	if err := k.LiquidDelegate(ctx, proxyAcc, activeVals, stakingCoin.Amount, whitelistedValsMap); err != nil {
		return stkIXOMintAmount, err
	}
	return stkIXOMintAmount, nil
}

// ----------------------------------------------------------------------------
// LiquidDelegate / DelegateWithCap (per-validator delegation primitives)
// ----------------------------------------------------------------------------

// LiquidDelegate splits stakingAmt across activeVals proportionally to their
// target weights and issues each delegation through the staking module's
// MsgDelegate handler routed via DelegateWithCap (which respects the
// chain-level liquid-staking cap).
func (k Keeper) LiquidDelegate(
	ctx sdk.Context, proxyAcc sdk.AccAddress, activeVals types.ActiveLiquidValidators, stakingAmt math.Int, whitelistedValsMap types.WhitelistedValsMap,
) error {
	// crumb may occur due to a decimal point error in dividing the staking amount into the weight of liquid validators, It added on first active liquid validator
	weightedAmt, crumb := types.DivideByWeight(activeVals, stakingAmt, whitelistedValsMap)
	if len(weightedAmt) == 0 {
		k.Logger(ctx).Error(
			"invalid active liquid validators",
			"active_validators", activeVals,
			"amount", stakingAmt.String(),
		)
		return types.ErrInvalidActiveLiquidValidators
	}
	// Allocate the rounding crumb to the first active validator.
	weightedAmt[0] = weightedAmt[0].Add(crumb)

	for i, val := range activeVals {
		if !weightedAmt[i].IsPositive() {
			continue
		}
		validator, _ := k.stakingKeeper.GetValidator(ctx, val.GetOperator())
		if err := k.DelegateWithCap(ctx, proxyAcc, validator, weightedAmt[i]); err != nil {
			return errorsmod.Wrapf(err, "failed to delegate to validator %s", val.GetOperator())
		}
	}
	return nil
}

// DelegateWithCap routes a MsgDelegate through the staking MsgServer so the
// chain-level liquid-staking cap accounting (LiquidShares / GlobalCap) is
// applied. Direct stakingKeeper.Delegate would bypass that bookkeeping.
func (k Keeper) DelegateWithCap(
	ctx sdk.Context, delegatorAddress sdk.AccAddress, validator stakingtypes.Validator, bondAmt math.Int,
) error {
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	msgDelegate := &stakingtypes.MsgDelegate{
		DelegatorAddress: delegatorAddress.String(),
		ValidatorAddress: validator.OperatorAddress,
		Amount:           sdk.NewCoin(bondDenom, bondAmt),
	}
	handler := k.router.Handler(msgDelegate)
	if handler == nil {
		k.Logger(ctx).Error("failed to find delegate handler")
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message route: %s", sdk.MsgTypeURL(msgDelegate))
	}
	res, err := handler(ctx, msgDelegate)
	if err != nil {
		k.Logger(ctx).Error("failed to execute delegate msg", "error", err, "msg", msgDelegate.String())
		return errorsmod.Wrapf(types.ErrDelegationFailed, "failed to send delegate msg with err: %v", err)
	}
	ctx.EventManager().EmitEvents(res.GetEvents())

	if len(res.MsgResponses) != 1 {
		return errorsmod.Wrapf(
			types.ErrInvalidResponse,
			"expected msg response should be exactly 1, got: %v, responses: %v",
			len(res.MsgResponses), res.MsgResponses,
		)
	}
	var msgDelegateResponse stakingtypes.MsgDelegateResponse
	if err := k.cdc.Unmarshal(res.MsgResponses[0].Value, &msgDelegateResponse); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal delegate tx response message: %v", err)
	}
	return nil
}

// UnbondWithCap performs the LSM tokenize → bank-send → redeem sequence so
// the unbonding queue is owned by the user (not the pool's proxy account).
// The result is that the user can retrieve their unbonded coins from the
// chain's standard unbonding flow even after our keeper has no further
// involvement.
func (k Keeper) UnbondWithCap(
	ctx sdk.Context,
	delegatorAddress sdk.AccAddress,
	validatorAddress sdk.ValAddress,
	amount sdk.Coin,
	userAddress sdk.AccAddress,
) (math.Int, error) {
	// [1] Tokenize the proxy's delegation into LSM shares owned by the user.
	lsmTokenizeMsg := &stakingtypes.MsgTokenizeShares{
		DelegatorAddress:    delegatorAddress.String(),
		ValidatorAddress:    validatorAddress.String(),
		Amount:              amount,
		TokenizedShareOwner: userAddress.String(),
	}
	handler := k.router.Handler(lsmTokenizeMsg)
	if handler == nil {
		k.Logger(ctx).Error("failed to find tokenize handler")
		return math.ZeroInt(), sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(lsmTokenizeMsg))
	}
	msgResp, err := handler(ctx, lsmTokenizeMsg)
	if err != nil {
		k.Logger(ctx).Error("failed to execute tokenize msg", "error", err, "msg", lsmTokenizeMsg.String())
		return math.ZeroInt(), types.ErrLSMTokenizeFailed.Wrapf("error: %s; message: %v", err.Error(), lsmTokenizeMsg)
	}
	ctx.EventManager().EmitEvents(msgResp.GetEvents())
	if len(msgResp.MsgResponses) != 1 {
		return math.ZeroInt(), errorsmod.Wrapf(
			types.ErrInvalidResponse,
			"expected msg response should be exactly 1, got: %v, responses: %v",
			len(msgResp.MsgResponses), msgResp.MsgResponses,
		)
	}
	var lsmTokenizeResp stakingtypes.MsgTokenizeSharesResponse
	if err = k.cdc.Unmarshal(msgResp.MsgResponses[0].Value, &lsmTokenizeResp); err != nil {
		return math.ZeroInt(), errorsmod.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal tokenize tx response message: %v", err)
	}

	// [2] Send the freshly-minted LSM shares from the proxy to the user.
	if err = k.bankKeeper.SendCoins(ctx, delegatorAddress, userAddress, sdk.NewCoins(lsmTokenizeResp.Amount)); err != nil {
		return math.ZeroInt(), err
	}

	// [3] Redeem the LSM shares back into a delegation owned by the user;
	// from this point the user owns the delegation and can unbond directly.
	lsmRedeemMsg := &stakingtypes.MsgRedeemTokensForShares{
		DelegatorAddress: userAddress.String(),
		Amount:           lsmTokenizeResp.Amount,
	}
	handler = k.router.Handler(lsmRedeemMsg)
	if handler == nil {
		k.Logger(ctx).Error("failed to find redeem handler")
		return math.ZeroInt(), sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(lsmRedeemMsg))
	}
	msgResp, err = handler(ctx, lsmRedeemMsg)
	if err != nil {
		k.Logger(ctx).Error("failed to execute redeem msg", "error", err, "msg", lsmRedeemMsg.String())
		return math.ZeroInt(), types.ErrLSMRedeemFailed.Wrapf("error: %s; message: %v", err.Error(), lsmRedeemMsg)
	}
	ctx.EventManager().EmitEvents(msgResp.GetEvents())
	if len(msgResp.MsgResponses) != 1 {
		return math.ZeroInt(), errorsmod.Wrapf(
			types.ErrInvalidResponse,
			"expected msg response should be exactly 1, got: %v, responses: %v",
			len(msgResp.MsgResponses), msgResp.MsgResponses,
		)
	}
	var lsmRedeemResp stakingtypes.MsgRedeemTokensForSharesResponse
	if err = k.cdc.Unmarshal(msgResp.MsgResponses[0].Value, &lsmRedeemResp); err != nil {
		return math.ZeroInt(), errorsmod.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal redeem tx response message: %v", err)
	}

	// [4] Now actually unbond — this triggers the standard Cosmos unbonding queue
	// from the user's account, completing in unbonding_time.
	msgUndelegate := &stakingtypes.MsgUndelegate{
		DelegatorAddress: userAddress.String(),
		ValidatorAddress: validatorAddress.String(),
		Amount:           lsmRedeemResp.Amount,
	}
	handler = k.router.Handler(msgUndelegate)
	if handler == nil {
		k.Logger(ctx).Error("failed to find undelegate handler")
		return math.ZeroInt(), sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(msgUndelegate))
	}
	msgResp, err = handler(ctx, msgUndelegate)
	if err != nil {
		k.Logger(ctx).Error("failed to execute undelegate msg", "error", err, "msg", msgUndelegate.String())
		return math.ZeroInt(), errorsmod.Wrapf(types.ErrUnstakeFailed, "failed to send undelegate msg with err: %v", err)
	}
	ctx.EventManager().EmitEvents(msgResp.GetEvents())


	if len(msgResp.MsgResponses) != 1 {
		return math.ZeroInt(), errorsmod.Wrapf(
			types.ErrInvalidResponse,
			"expected msg response should be exactly 1, got: %v, responses: %v",
			len(msgResp.MsgResponses), msgResp.MsgResponses,
		)
	}

	var msgUndelegateResp stakingtypes.MsgUndelegateResponse
	if err = k.cdc.Unmarshal(msgResp.MsgResponses[0].Value, &msgUndelegateResp); err != nil {
		return math.ZeroInt(), errorsmod.Wrapf(
			sdkerrors.ErrJSONUnmarshal,
			"cannot unmarshal msg undelegate tx response message: %v",
			err,
		)
	}

	return lsmRedeemResp.Amount.Amount, nil
}

// RedelegateWithCap routes a MsgBeginRedelegate through the staking MsgServer
// so the cap accounting is applied, but account for the amount of liquid staked shares and check
// against liquid staking cap.
func (k Keeper) RedelegateWithCap(
	ctx sdk.Context,
	delegatorAddress sdk.AccAddress,
	validatorSrc sdk.ValAddress,
	validatorDst sdk.ValAddress,
	bondAmt math.Int,
) (time.Time, error) {
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	msgRedelegate := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    delegatorAddress.String(),
		ValidatorSrcAddress: validatorSrc.String(),
		ValidatorDstAddress: validatorDst.String(),
		Amount:              sdk.NewCoin(bondDenom, bondAmt),
	}
	handler := k.router.Handler(msgRedelegate)
	if handler == nil {
		k.Logger(ctx).Error("failed to find redelegate handler")
		return time.Time{}, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message route: %s", sdk.MsgTypeURL(msgRedelegate))
	}
	res, err := handler(ctx, msgRedelegate)
	if err != nil {
		k.Logger(ctx).Error("failed to execute redelegate msg", "error", err, "msg", msgRedelegate.String())
		return time.Time{}, errorsmod.Wrapf(types.ErrRedelegateFailed, "failed to send redelegate msg with err: %v", err)
	}
	ctx.EventManager().EmitEvents(res.GetEvents())

	if len(res.MsgResponses) != 1 {
		return time.Time{}, errorsmod.Wrapf(
			types.ErrInvalidResponse,
			"expected msg response should be exactly 1, got: %v, responses: %v",
			len(res.MsgResponses), res.MsgResponses,
		)
	}
	var msgRedelegateResponse stakingtypes.MsgBeginRedelegateResponse
	if err = k.cdc.Unmarshal(res.MsgResponses[0].Value, &msgRedelegateResponse); err != nil {
		return time.Time{}, errorsmod.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal redelegate tx response message: %v", err)
	}
	return msgRedelegateResponse.CompletionTime, nil
}

// ----------------------------------------------------------------------------
// LiquidUnstake
// ----------------------------------------------------------------------------

// LiquidUnstake burns the user's stkIXO denominated in pool.LiquidBondDenom
// and unbonds the corresponding NetAmount-derived stake from the pool's
// validators back to the user. Returns: completion time of the unbonding,
// total amount initiating unbonding (bondDenom), the unbonding entries
// created, total amount transferred immediately (e.g. when proxy held
// spendable balance), and any error.
func (k Keeper) LiquidUnstake(
	ctx sdk.Context, p types.Pool, liquidStaker sdk.AccAddress, unstakingStkIXO sdk.Coin,
) (time.Time, math.Int, []stakingtypes.UnbondingDelegation, math.Int, error) {
	if k.PoolEffectivelyPaused(ctx, p) {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), types.ErrPoolPaused
	}

	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}

	if unstakingStkIXO.Denom != p.LiquidBondDenom {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), errorsmod.Wrapf(
			types.ErrPoolDenomMismatch, "got %s, expected %s", unstakingStkIXO.Denom, p.LiquidBondDenom,
		)
	}

	nas := k.GetNetAmountStateForPool(ctx, p)
	if unstakingStkIXO.Amount.GT(nas.StkixoTotalSupply) || nas.StkixoTotalSupply.IsZero() {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), types.ErrInvalidStkIXOSupply
	}

	// UnstakeAmount = NetAmount * stkIXO/totalSupply * (1 - unstakeFeeRate)
	unbondingAmount := types.StkIXOToNativeToken(unstakingStkIXO.Amount, nas.StkixoTotalSupply, nas.NetAmount)
	unbondingAmount = types.DeductFeeRate(unbondingAmount, p.UnstakeFeeRate)
	unbondingAmountInt := unbondingAmount.TruncateInt()
	if !unbondingAmountInt.IsPositive() {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), types.ErrTooSmallLiquidUnstakingAmount
	}

	// Burn the staker's stkIXO before unbonding to avoid two parties claiming
	// the same backing tokens if anything mid-flight fails.
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, liquidStaker, types.ModuleName, sdk.NewCoins(unstakingStkIXO)); err != nil {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(unstakingStkIXO)); err != nil {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), err
	}

	proxyAcc := p.GetProxyAccount()
	liquidVals := k.GetAllLiquidValidatorsForPool(ctx, p.PoolId)
	totalLiquidTokens, liquidTokenMap := liquidVals.TotalLiquidTokens(ctx, k.stakingKeeper, proxyAcc, false)

	// If the pool has no live delegations, satisfy the request from the
	// proxy account's spendable balance (e.g. rewards waiting for the next
	// autocompound). This handles the post-rebalance edge case where stake
	// is briefly between validators.
	if !totalLiquidTokens.IsPositive() {
		if nas.ProxyAccBalance.GTE(unbondingAmountInt) {
			if err := k.bankKeeper.SendCoins(
				ctx, proxyAcc, liquidStaker,
				sdk.NewCoins(sdk.NewCoin(bondDenom, unbondingAmountInt)),
			); err != nil {
				return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), err
			}
			return time.Time{}, math.ZeroInt(), nil, unbondingAmountInt, nil
		}
		k.Logger(ctx).Error(
			"non-positive total liquid tokens",
			"pool_id", p.PoolId,
			"validators", liquidVals,
			"total_liquid_tokens", totalLiquidTokens.String(),
		)
		// error case where there is a quantity that are unbonding balance or remaining rewards that is not re-stake or withdrawn in netAmount.
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), types.ErrInsufficientProxyAccBalance
	}

	// fail when no liquid validators to unbond
	if liquidVals.Len() == 0 {
		k.Logger(ctx).Error("no liquid validators to unbond", "pool_id", p.PoolId, "validators", liquidVals)
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), types.ErrLiquidValidatorsNotExists
	}

	// Drain inactive validators first so we move stake away from
	// no-longer-whitelisted operators before touching active ones.
	liquidVals = k.PrioritiseInactiveLiquidValidators(ctx, liquidVals)

	// crumb may occur due to a decimal error in dividing the unstaking stkIXO into the weight of liquid validators, it will remain in the NetAmount
	unbondingAmounts, crumb := types.DivideByCurrentWeight(liquidVals, unbondingAmount, totalLiquidTokens, liquidTokenMap)
	if !unbondingAmount.Sub(crumb).IsPositive() {
		return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), types.ErrTooSmallLiquidUnstakingAmount
	}

	totalReturnAmount := math.ZeroInt()
	var ubdTime time.Time
	ubds := make([]stakingtypes.UnbondingDelegation, 0, len(liquidVals))
	for i, val := range liquidVals {
		if !unbondingAmounts[i].IsPositive() {
			continue
		}
		weightedShare, err := k.stakingKeeper.ValidateUnbondAmount(ctx, proxyAcc, val.GetOperator(), unbondingAmounts[i].TruncateInt())
		if err != nil {
			k.Logger(ctx).Error(
				"failed to validate unbond amount",
				"error", err,
				"pool_id", p.PoolId,
				"validator", val.GetOperator().String(),
				"amount", unbondingAmounts[i].TruncateInt().String(),
			)
			return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), err
		}
		if !weightedShare.IsPositive() {
			continue
		}

		ubdT, returnAmount, ubd, err := k.LiquidUnbond(
			ctx, proxyAcc, liquidStaker, val.GetOperator(), weightedShare, true,
			sdk.NewCoin(bondDenom, unbondingAmounts[i].TruncateInt()),
		)
		if err != nil {
			return time.Time{}, math.ZeroInt(), nil, math.ZeroInt(), err
		}
		ubdTime = ubdT
		ubds = append(ubds, ubd)
		totalReturnAmount = totalReturnAmount.Add(returnAmount)
	}

	return ubdTime, totalReturnAmount, ubds, math.ZeroInt(), nil
}

// LiquidUnbond invokes the LSM tokenize/redeem/undelegate sequence on a
// single validator delegation. checkMaxEntries bounds queue depth.
func (k Keeper) LiquidUnbond(
	ctx sdk.Context, proxyAcc, liquidStaker sdk.AccAddress, valAddr sdk.ValAddress, shares math.LegacyDec, checkMaxEntries bool, unbondAmount sdk.Coin,
) (time.Time, math.Int, stakingtypes.UnbondingDelegation, error) {
	if _, err := k.stakingKeeper.GetValidator(ctx, valAddr); err != nil {
		return time.Time{}, math.ZeroInt(), stakingtypes.UnbondingDelegation{}, stakingtypes.ErrNoDelegatorForAddress
	}

	// If checkMaxEntries is true, perform a maximum limit unbonding entries check.
	hasMax, err := k.stakingKeeper.HasMaxUnbondingDelegationEntries(ctx, liquidStaker, valAddr)
	if err != nil {
		return time.Time{}, math.ZeroInt(), stakingtypes.UnbondingDelegation{}, err
	}
	if checkMaxEntries && hasMax {
		return time.Time{}, math.ZeroInt(), stakingtypes.UnbondingDelegation{}, stakingtypes.ErrMaxUnbondingDelegationEntries
	}

	// unbond from proxy account
	returnAmount, err := k.UnbondWithCap(ctx, proxyAcc, valAddr, unbondAmount, liquidStaker)
	if err != nil {
		return time.Time{}, math.ZeroInt(), stakingtypes.UnbondingDelegation{}, err
	}

	// Unbonding from proxy account, but queues to liquid staker.
	unbondingTime, err := k.stakingKeeper.UnbondingTime(ctx)
	if err != nil {
		return time.Time{}, math.ZeroInt(), stakingtypes.UnbondingDelegation{}, err
	}
	completionTime := ctx.BlockHeader().Time.Add(unbondingTime)
	ubd, err := k.stakingKeeper.GetUnbondingDelegation(ctx, liquidStaker, valAddr)
	if err != nil {
		k.Logger(ctx).Error("failed to find unbonding delegation",
			"delegator", liquidStaker.String(),
			"validator", valAddr.String(),
		)
		return time.Time{}, math.ZeroInt(), stakingtypes.UnbondingDelegation{}, types.ErrInvalidResponse.Wrap("expected undelegation entry, found none")
	}
	return completionTime, returnAmount, ubd, nil
}

// PrioritiseInactiveLiquidValidators sorts the slice in-place so inactive
// liquid validators come first. Used by LiquidUnstake so funds drain away
// from de-whitelisted/tombstoned validators ahead of healthy ones.
func (k Keeper) PrioritiseInactiveLiquidValidators(
	ctx sdk.Context, vs types.LiquidValidators,
) types.LiquidValidators {
	sort.SliceStable(vs, func(i, j int) bool {
		vs1, vs1err := k.stakingKeeper.GetValidator(ctx, vs[i].GetOperator())
		vs2, vs2err := k.stakingKeeper.GetValidator(ctx, vs[j].GetOperator())

		switch {
		case vs1err != nil && vs2err != nil:
			// only one case when less
			return true
		case vs1err == nil && vs2err == nil:
			// both exist, compare status
			vs1Active := vs[i].GetStatus(types.ActiveCondition(vs1, true, k.IsTombstoned(ctx, vs1)))
			vs2Active := vs[j].GetStatus(types.ActiveCondition(vs2, true, k.IsTombstoned(ctx, vs2)))
			return vs1Active != types.ValidatorStatusActive && vs2Active == types.ValidatorStatusActive
		default:
			return false
		}
	})
	return vs
}

// CheckDelegationStates aggregates the proxy account's delegation positions:
// total unwithdrawn rewards, total delegation shares, and total
// (slashing-applied) liquid tokens. Uses a CacheContext to keep
// reward-period side effects out of the parent transaction.
func (k Keeper) CheckDelegationStates(ctx sdk.Context, proxyAcc sdk.AccAddress) (math.LegacyDec, math.LegacyDec, math.Int) {
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	totalRewards := math.LegacyZeroDec()
	totalDelShares := math.LegacyZeroDec()
	totalLiquidTokens := math.ZeroInt()

	// Cache ctx for calculate rewards
	cachedCtx, _ := ctx.CacheContext()
	k.stakingKeeper.IterateDelegations(
		cachedCtx, proxyAcc,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr, err := k.stakingKeeper.ValidatorAddressCodec().StringToBytes(del.GetValidatorAddr())
			if err != nil {
				panic(err)
			}
			val, err := k.stakingKeeper.Validator(cachedCtx, valAddr)
			if err != nil {
				panic(err) // should never happen
			}
			endingPeriod, err := k.distrKeeper.IncrementValidatorPeriod(cachedCtx, val)
			if err != nil {
				panic(err) // should never happen
			}
			delReward, err := k.distrKeeper.CalculateDelegationRewards(cachedCtx, val, del, endingPeriod)
			if err != nil {
				panic(err) // should never happen
			}
			delShares := del.GetShares()
			if delShares.IsPositive() {
				totalDelShares = totalDelShares.Add(delShares)
				liquidTokens := val.TokensFromSharesTruncated(delShares).TruncateInt()
				totalLiquidTokens = totalLiquidTokens.Add(liquidTokens)
				totalRewards = totalRewards.Add(delReward.AmountOf(bondDenom).TruncateDec())
			}
			return false
		},
	)
	return totalRewards, totalDelShares, totalLiquidTokens
}

// WithdrawLiquidRewards iterates every delegation held by proxyAcc and
// withdraws accrued rewards through distribution module's MsgServer.
// Errors are logged and swallowed: a single failed withdrawal must not
// abort BeginBlock.
func (k Keeper) WithdrawLiquidRewards(ctx sdk.Context, proxyAcc sdk.AccAddress) {
	k.stakingKeeper.IterateDelegations(
		ctx, proxyAcc,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			msgWithdraw := &distributiontypes.MsgWithdrawDelegatorReward{
				DelegatorAddress: proxyAcc.String(),
				ValidatorAddress: del.GetValidatorAddr(),
			}
			handler := k.router.Handler(msgWithdraw)
			if handler == nil {
				k.Logger(ctx).Error("could not find distribution handler for withdraw rewards msg")
				return true
			}
			res, err := handler(ctx, msgWithdraw)
			if err != nil {
				k.Logger(ctx).Error(
					"failed to execute withdraw rewards msg",
					"msg", msgWithdraw.String(),
					"error", err,
				)
				// will be retried next epoch
			} else {
				ctx.EventManager().EmitEvents(res.GetEvents())
			}
			return false
		},
	)
}

// ----------------------------------------------------------------------------
// LiquidValidator queries (per-pool aggregations)
// ----------------------------------------------------------------------------

// GetActiveLiquidValidatorsForPool returns the per-pool LiquidValidators
// that meet ActiveCondition under the given pool whitelist.
func (k Keeper) GetActiveLiquidValidatorsForPool(ctx sdk.Context, poolID string, whitelistedValsMap types.WhitelistedValsMap) types.ActiveLiquidValidators {
	out := types.ActiveLiquidValidators{}
	k.IterateLiquidValidatorsForPool(ctx, poolID, func(val types.LiquidValidator) bool {
		if k.IsActiveLiquidValidator(ctx, val, whitelistedValsMap) {
			out = append(out, val)
		}
		return false
	})
	return out
}

// GetAllLiquidValidatorStatesForPool returns every LiquidValidator known to a
// pool annotated with current weight/status/del-shares/liquid-tokens.
func (k Keeper) GetAllLiquidValidatorStatesForPool(ctx sdk.Context, p types.Pool) []types.LiquidValidatorState {
	whitelistedValsMap := p.WhitelistedValsMap()
	proxyAcc := p.GetProxyAccount()
	out := []types.LiquidValidatorState{}
	k.IterateLiquidValidatorsForPool(ctx, p.PoolId, func(lv types.LiquidValidator) bool {
		active := k.IsActiveLiquidValidator(ctx, lv, whitelistedValsMap)
		out = append(out, types.LiquidValidatorState{
			OperatorAddress: lv.OperatorAddress,
			Weight:          lv.GetWeight(whitelistedValsMap, active),
			Status:          lv.GetStatus(active),
			DelShares:       lv.GetDelShares(ctx, k.stakingKeeper, proxyAcc),
			LiquidTokens:    lv.GetLiquidTokens(ctx, k.stakingKeeper, proxyAcc, false),
		})
		return false
	})
	return out
}

// IsActiveLiquidValidator checks ActiveCondition for a single validator
// against a pool's whitelist map.
func (k Keeper) IsActiveLiquidValidator(ctx sdk.Context, lv types.LiquidValidator, whitelistedValsMap types.WhitelistedValsMap) bool {
	val, err := k.stakingKeeper.GetValidator(ctx, lv.GetOperator())
	if err != nil {
		return false
	}
	return types.ActiveCondition(val, whitelistedValsMap.IsListed(lv.OperatorAddress), k.IsTombstoned(ctx, val))
}

func (k Keeper) IsTombstoned(ctx sdk.Context, val stakingtypes.Validator) bool {
	consPk, err := val.ConsPubKey()
	if err != nil {
		return false
	}
	return k.slashingKeeper.IsTombstoned(ctx, sdk.ConsAddress(consPk.Address()))
}

// GetWeightMap returns a per-validator active-weight map plus the sum.
// Inactive validators contribute zero so the map shape is preserved for
// downstream consumers (Rebalance) that index by operator address.
func (k Keeper) GetWeightMap(ctx sdk.Context, liquidVals types.LiquidValidators, whitelistedValsMap types.WhitelistedValsMap) (map[string]math.Int, math.Int) {
	weightMap := map[string]math.Int{}
	totalWeight := math.ZeroInt()
	for _, val := range liquidVals {
		weight := val.GetWeight(whitelistedValsMap, k.IsActiveLiquidValidator(ctx, val, whitelistedValsMap))
		totalWeight = totalWeight.Add(weight)
		weightMap[val.OperatorAddress] = weight
	}
	return weightMap, totalWeight
}
