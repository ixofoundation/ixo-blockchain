package types

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// WhitelistedValsMap indexes WhitelistedValidator entries by operator address
// for O(1) lookups inside hot paths.
type WhitelistedValsMap map[string]WhitelistedValidator

func (whitelistedValsMap WhitelistedValsMap) IsListed(operatorAddr string) bool {
	_, ok := whitelistedValsMap[operatorAddr]
	return ok
}

func GetWhitelistedValsMap(whitelistedValidators []WhitelistedValidator) WhitelistedValsMap {
	out := make(WhitelistedValsMap, len(whitelistedValidators))
	for _, wv := range whitelistedValidators {
		out[wv.ValidatorAddress] = wv
	}
	return out
}

// Validate checks the LiquidValidator's operator address parses as bech32.
func (v LiquidValidator) Validate() error {
	_, err := sdk.ValAddressFromBech32(v.OperatorAddress)
	return err
}

func (v LiquidValidator) GetOperator() sdk.ValAddress {
	if v.OperatorAddress == "" {
		return nil
	}
	addr, err := sdk.ValAddressFromBech32(v.OperatorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// GetDelShares returns the delegation shares this liquid validator holds for
// the given proxy account. The proxy account is now per-pool, so callers
// must pass it explicitly rather than reading a module-global address.
func (v LiquidValidator) GetDelShares(ctx context.Context, sk StakingKeeper, proxyAcc sdk.AccAddress) math.LegacyDec {
	del, err := sk.GetDelegation(ctx, proxyAcc, v.GetOperator())
	if err != nil {
		return math.LegacyZeroDec()
	}
	return del.GetShares()
}

// GetLiquidTokens returns the slashing-applied token-equivalent of the proxy
// account's delegation shares to this validator. If onlyBonded is true,
// returns zero for non-bonded validators.
func (v LiquidValidator) GetLiquidTokens(ctx context.Context, sk StakingKeeper, proxyAcc sdk.AccAddress, onlyBonded bool) math.Int {
	delShares := v.GetDelShares(ctx, sk, proxyAcc)
	if !delShares.IsPositive() {
		return math.ZeroInt()
	}
	val, err := sk.Validator(ctx, v.GetOperator())
	if err != nil {
		return math.ZeroInt()
	}
	if onlyBonded && !val.IsBonded() {
		return math.ZeroInt()
	}
	return val.TokensFromSharesTruncated(delShares).TruncateInt()
}

// GetWeight returns this validator's whitelisted target weight if it's
// active, otherwise zero. Inactive validators receive zero new delegations
// during rebalancing.
func (v LiquidValidator) GetWeight(whitelistedValsMap WhitelistedValsMap, active bool) math.Int {
	if wv, ok := whitelistedValsMap[v.OperatorAddress]; ok && active {
		return wv.TargetWeight
	}
	return math.ZeroInt()
}

func (v LiquidValidator) GetStatus(activeCondition bool) ValidatorStatus {
	if activeCondition {
		return ValidatorStatusActive
	}
	return ValidatorStatusInactive
}

// ActiveCondition is the canonical predicate for "may this liquid validator
// receive new delegations": whitelisted by its pool, present in the staking
// module with sane delegation/exchange-rate state, and not tombstoned.
func ActiveCondition(validator stakingtypes.Validator, whitelisted, tombstoned bool) bool {
	return whitelisted &&
		!tombstoned &&
		validator.GetStatus() != stakingtypes.Unspecified &&
		!validator.GetTokens().IsNil() &&
		!validator.GetDelegatorShares().IsNil() &&
		!validator.InvalidExRate()
}

// LiquidValidators is a collection of LiquidValidator.
type (
	LiquidValidators       []LiquidValidator
	ActiveLiquidValidators LiquidValidators
)

// MinMaxGap returns the validators with the largest positive and largest
// negative gap from their target weights, plus the redelegation amount that
// would close them. lastRedelegation signals that the source validator has
// no remaining target weight, so its full balance can be redelegated away.
func (vs LiquidValidators) MinMaxGap(targetMap, liquidTokenMap map[string]math.Int) (minGapVal, maxGapVal LiquidValidator, amountNeeded math.Int, lastRedelegation bool) {
	maxGap := math.ZeroInt()
	minGap := math.ZeroInt()

	for _, val := range vs {
		gap := liquidTokenMap[val.OperatorAddress].Sub(targetMap[val.OperatorAddress])
		if gap.GT(maxGap) {
			maxGap = gap
			maxGapVal = val
		}
		if gap.LT(minGap) {
			minGap = gap
			minGapVal = val
		}
	}
	amountNeeded = math.MinInt(maxGap, minGap.Abs())
	lastRedelegation = amountNeeded.IsPositive() &&
		!targetMap[maxGapVal.OperatorAddress].IsPositive() &&
		liquidTokenMap[maxGapVal.OperatorAddress].Equal(amountNeeded)

	return minGapVal, maxGapVal, amountNeeded, lastRedelegation
}

func (vs LiquidValidators) Len() int { return len(vs) }

// TotalLiquidTokens sums the (slashing-applied) token-equivalent of the proxy
// account's delegations across the given validators, returning both the
// aggregate total and the per-validator breakdown for downstream use.
func (vs LiquidValidators) TotalLiquidTokens(ctx context.Context, sk StakingKeeper, proxyAcc sdk.AccAddress, onlyBonded bool) (math.Int, map[string]math.Int) {
	total := math.ZeroInt()
	perValidator := map[string]math.Int{}
	for _, lv := range vs {
		liquid := lv.GetLiquidTokens(ctx, sk, proxyAcc, onlyBonded)
		perValidator[lv.OperatorAddress] = liquid
		total = total.Add(liquid)
	}
	return total, perValidator
}

// Map returns a set-valued map keyed by OperatorAddress, used for O(1)
// presence checks.
func (vs LiquidValidators) Map() map[string]struct{} {
	out := map[string]struct{}{}
	for _, val := range vs {
		out[val.OperatorAddress] = struct{}{}
	}
	return out
}

func (avs ActiveLiquidValidators) Len() int { return LiquidValidators(avs).Len() }

func (avs ActiveLiquidValidators) TotalActiveLiquidTokens(ctx context.Context, sk StakingKeeper, proxyAcc sdk.AccAddress, onlyBonded bool) (math.Int, map[string]math.Int) {
	return LiquidValidators(avs).TotalLiquidTokens(ctx, sk, proxyAcc, onlyBonded)
}

// TotalWeight sums the target weights of every active validator under the
// given pool whitelist. Used to verify the active-weight quorum before
// accepting a LiquidStake.
func (avs ActiveLiquidValidators) TotalWeight(whitelistedValsMap WhitelistedValsMap) math.Int {
	totalWeight := math.ZeroInt()
	for _, val := range avs {
		totalWeight = totalWeight.Add(val.GetWeight(whitelistedValsMap, true))
	}
	return totalWeight
}

// NativeTokenToStkIXO is fixed at 1:1 by design: 1 native token always mints
// exactly 1 LST regardless of accrued rewards. Value accrual surfaces only
// on the unstake side via NetAmount/Supply (see StkIXOToNativeToken).
//
// LiquidStake is currently restricted to the pool admin, so the asymmetric
// mint/burn design has no dilution effect in practice. If LiquidStake is
// ever opened to multiple stakers per pool, this formula must be revisited
// (a new staker would otherwise claim a pro-rata share of already-accrued
// rewards).
func NativeTokenToStkIXO(nativeTokenAmount math.Int) (stkIXOAmount math.Int) {
	return nativeTokenAmount
}

// StkIXOToNativeToken returns stkIXOAmount * netAmount / stkIXOTotalSupply
// with truncating integer division. Returns zero if total supply is zero
// (caller's responsibility to handle).
func StkIXOToNativeToken(stkIXOAmount, stkIXOTotalSupplyAmount math.Int, netAmount math.LegacyDec) (nativeTokenAmount math.LegacyDec) {
	if stkIXOTotalSupplyAmount.IsZero() {
		return math.LegacyZeroDec()
	}
	return math.LegacyNewDecFromInt(stkIXOAmount).MulTruncate(netAmount).Quo(math.LegacyNewDecFromInt(stkIXOTotalSupplyAmount)).TruncateDec()
}

// DeductFeeRate returns input * (1 - feeRate) with truncation.
func DeductFeeRate(input, feeRate math.LegacyDec) (feeDeductedOutput math.LegacyDec) {
	return input.MulTruncate(math.LegacyOneDec().Sub(feeRate)).TruncateDec()
}

// CalcNetAmount returns the net amount of native tokens backing a pool's LST
// supply: liquid (delegated, slashing-applied) tokens plus tokens currently
// in the unbonding queue from the proxy account.
func (nas NetAmountState) CalcNetAmount() math.LegacyDec {
	return math.LegacyNewDecFromInt(nas.TotalLiquidTokens.Add(nas.TotalUnbondingBalance))
}

type LiquidValidatorStates []LiquidValidatorState

func MustMarshalLiquidValidator(cdc codec.BinaryCodec, val *LiquidValidator) []byte {
	return cdc.MustMarshal(val)
}

func MustUnmarshalLiquidValidator(cdc codec.BinaryCodec, value []byte) LiquidValidator {
	validator, err := UnmarshalLiquidValidator(cdc, value)
	if err != nil {
		panic(err)
	}
	return validator
}

func UnmarshalLiquidValidator(cdc codec.BinaryCodec, value []byte) (val LiquidValidator, err error) {
	err = cdc.Unmarshal(value, &val)
	return val, err
}

func (w *WhitelistedValidator) GetValidatorAddress() sdk.ValAddress {
	valAddr, err := sdk.ValAddressFromBech32(w.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return valAddr
}
