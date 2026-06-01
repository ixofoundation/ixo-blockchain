package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Per-pool defaults applied when MsgCreatePool does not specify a value.
// Pools must explicitly opt-in to non-zero fees through MsgUpdatePool.
var (
	DefaultUnstakeFeeRate      = math.LegacyZeroDec()
	DefaultAutocompoundFeeRate = math.LegacyZeroDec()
)

// Module-wide invariants over the per-pool validator set:
//
//   - TotalValidatorWeight: WhitelistedValidator.TargetWeight values must
//     sum to exactly this value within a single pool. Same value applies to
//     every pool independently.
//   - ActiveLiquidValidatorsWeightQuorum: a pool's active (non-tombstoned,
//     bonded) validators must collectively hold at least this fraction of
//     the pool's total weight before any LiquidStake will succeed.
//   - RebalancingTrigger: per-validator weight deviations smaller than this
//     fraction of the pool's totalLiquidTokens are ignored by Rebalance.
var (
	TotalValidatorWeight               = math.NewInt(10_000)
	ActiveLiquidValidatorsWeightQuorum = math.LegacyNewDecWithPrec(3333, 4) // 0.3333
	RebalancingTrigger                 = math.LegacyNewDecWithPrec(1, 3)    // 0.001
)

// String renders Pool as indented JSON. Required because the proto disables
// gogoproto's auto-generated stringer.
func (p Pool) String() string {
	out, _ := json.MarshalIndent(p, "", "  ")
	return string(out)
}

// GetProxyAccount returns the bech32-decoded proxy account address. Panics if
// the stored address is malformed; CreatePool guarantees it is well-formed,
// so a panic here indicates corrupted state.
func (p Pool) GetProxyAccount() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(p.ProxyAccountAddress)
	if err != nil {
		panic(fmt.Errorf("pool %q: malformed proxy_account_address %q: %w", p.PoolId, p.ProxyAccountAddress, err))
	}
	return addr
}

// WhitelistedValsMap returns this pool's whitelisted validators indexed by
// validator operator address.
func (p Pool) WhitelistedValsMap() WhitelistedValsMap {
	return GetWhitelistedValsMap(p.WhitelistedValidators)
}

// Validate fully validates a Pool record, including pool_id format, denom
// shape, all addresses, fee rates, validator-weight invariants, and weighted
// rewards receivers. Used by genesis import and as a defence-in-depth check
// before persisting a Pool.
func (p Pool) Validate() error {
	if err := ValidatePoolID(p.PoolId); err != nil {
		return err
	}
	if err := validateLiquidBondDenom(p.LiquidBondDenom); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(p.ProxyAccountAddress); err != nil {
		return fmt.Errorf("invalid proxy_account_address %q: %w", p.ProxyAccountAddress, err)
	}
	if err := validateUnstakeFeeRate(p.UnstakeFeeRate); err != nil {
		return err
	}
	if err := validateAutocompoundFeeRate(p.AutocompoundFeeRate); err != nil {
		return err
	}
	if err := validateFeeAccountAddress(p.FeeAccountAddress); err != nil {
		return err
	}
	if err := validateOptionalAdminAddress(p.WhitelistAdminAddress); err != nil {
		return err
	}
	if err := ValidateWhitelistedValidators(p.WhitelistedValidators); err != nil {
		return err
	}
	if err := ValidateWeightedRewardsReceivers(p.WeightedRewardsReceivers); err != nil {
		return err
	}
	return nil
}

func validateLiquidBondDenom(v string) error {
	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("liquid bond denom cannot be blank")
	}
	return sdk.ValidateDenom(v)
}

func validateUnstakeFeeRate(v math.LegacyDec) error {
	if v.IsNil() {
		return fmt.Errorf("unstake fee rate must not be nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("unstake fee rate must not be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("unstake fee rate too large: %s", v)
	}
	return nil
}

func validateAutocompoundFeeRate(v math.LegacyDec) error {
	if v.IsNil() {
		return fmt.Errorf("autocompound fee rate must not be nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("autocompound fee rate must not be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("autocompound fee rate too large: %s", v)
	}
	return nil
}

func validateFeeAccountAddress(v string) error {
	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid fee_account_address %q: %w", v, err)
	}
	return nil
}

// validateOptionalAdminAddress allows an empty admin address (matches the
// pre-v7 default for newly created pools that have not yet appointed an
// admin), but otherwise requires a valid bech32.
func validateOptionalAdminAddress(v string) error {
	if v == "" {
		return nil
	}
	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid whitelist_admin_address %q: %w", v, err)
	}
	return nil
}

// ValidateWhitelistedValidators checks that every entry has a valid validator
// address, a positive non-nil target weight, and that addresses are unique.
// Note: it does NOT enforce sum == TotalValidatorWeight here, because Pool
// genesis validation also has to accept a freshly created (empty) pool.
// The sum check lives in msg_server when a non-empty list is being applied.
func ValidateWhitelistedValidators(wvs []WhitelistedValidator) error {
	seen := map[string]struct{}{}
	for _, wv := range wvs {
		if _, err := sdk.ValAddressFromBech32(wv.ValidatorAddress); err != nil {
			return fmt.Errorf("invalid validator_address %q: %w", wv.ValidatorAddress, err)
		}
		if wv.TargetWeight.IsNil() {
			return fmt.Errorf("liquidstake validator target weight must not be nil")
		}
		if !wv.TargetWeight.IsPositive() {
			return fmt.Errorf("liquidstake validator target weight must be positive: %s", wv.TargetWeight)
		}
		if _, dup := seen[wv.ValidatorAddress]; dup {
			return fmt.Errorf("liquidstake validator cannot be duplicated: %s", wv.ValidatorAddress)
		}
		seen[wv.ValidatorAddress] = struct{}{}
	}
	return nil
}

// ValidateWeightedRewardsReceivers checks every receiver's address and
// weight, and that the total weight does not exceed 1.
func ValidateWeightedRewardsReceivers(receivers []WeightedAddress) error {
	if len(receivers) == 0 {
		return nil
	}
	weightSum := math.LegacyZeroDec()
	for i, w := range receivers {
		if _, err := sdk.AccAddressFromBech32(w.Address); err != nil {
			return fmt.Errorf("weighted_rewards_receivers[%d]: invalid address: %w", i, err)
		}
		if w.Weight.IsNil() || !w.Weight.IsPositive() {
			return fmt.Errorf("weighted_rewards_receivers[%d]: weight must be positive", i)
		}
		if w.Weight.GT(math.LegacyOneDec()) {
			return fmt.Errorf("weighted_rewards_receivers[%d]: weight cannot exceed 1", i)
		}
		weightSum = weightSum.Add(w.Weight)
	}
	if weightSum.GT(math.LegacyOneDec()) {
		return fmt.Errorf("sum of weighted_rewards_receivers weights cannot exceed 1: got %s", weightSum)
	}
	return nil
}
