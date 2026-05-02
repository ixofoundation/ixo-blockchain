package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Compile-time assertions that every concrete request type implements sdk.Msg.
var (
	_ sdk.Msg = &MsgLiquidStake{}
	_ sdk.Msg = &MsgLiquidUnstake{}
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgUpdateModuleParams{}
	_ sdk.Msg = &MsgUpdatePool{}
	_ sdk.Msg = &MsgUpdateWhitelistedValidators{}
	_ sdk.Msg = &MsgUpdateWeightedRewardsReceivers{}
	_ sdk.Msg = &MsgSetPoolPaused{}
	_ sdk.Msg = &MsgSetModulePaused{}
	_ sdk.Msg = &MsgBurn{}
)

// Stable, human-readable type identifiers used for telemetry/event filtering.
const (
	MsgTypeLiquidStake                     = "liquid_stake"
	MsgTypeLiquidUnstake                   = "liquid_unstake"
	MsgTypeCreatePool                      = "create_pool"
	MsgTypeUpdateModuleParams              = "update_module_params"
	MsgTypeUpdatePool                      = "update_pool"
	MsgTypeUpdateWhitelistedValidators     = "update_whitelisted_validators"
	MsgTypeUpdateWeightedRewardsReceivers  = "update_weighted_rewards_receivers"
	MsgTypeSetPoolPaused                   = "set_pool_paused"
	MsgTypeSetModulePaused                 = "set_module_paused"
	MsgTypeBurn                            = "burn"
)

// --------------------------
// LIQUID STAKE
// --------------------------

func NewMsgLiquidStake(liquidStaker sdk.AccAddress, poolID string, amount sdk.Coin) *MsgLiquidStake {
	return &MsgLiquidStake{
		DelegatorAddress: liquidStaker.String(),
		PoolId:           poolID,
		Amount:           amount,
	}
}

func (m *MsgLiquidStake) Route() string { return RouterKey }
func (m *MsgLiquidStake) Type() string  { return MsgTypeLiquidStake }

func (m *MsgLiquidStake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.DelegatorAddress); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address %q: %v", m.DelegatorAddress, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	if m.Amount.IsZero() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "staking amount must not be zero")
	}
	if err := m.Amount.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgLiquidStake) GetDelegator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// --------------------------
// LIQUID UNSTAKE
// --------------------------

func NewMsgLiquidUnstake(liquidStaker sdk.AccAddress, poolID string, amount sdk.Coin) *MsgLiquidUnstake {
	return &MsgLiquidUnstake{
		DelegatorAddress: liquidStaker.String(),
		PoolId:           poolID,
		Amount:           amount,
	}
}

func (m *MsgLiquidUnstake) Route() string { return RouterKey }
func (m *MsgLiquidUnstake) Type() string  { return MsgTypeLiquidUnstake }

func (m *MsgLiquidUnstake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.DelegatorAddress); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address %q: %v", m.DelegatorAddress, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	if m.Amount.IsZero() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "unstaking amount must not be zero")
	}
	if err := m.Amount.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgLiquidUnstake) GetDelegator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// --------------------------
// CREATE POOL
// --------------------------

func NewMsgCreatePool(authority sdk.AccAddress, poolID, liquidBondDenom string, initialAdmin, initialFeeAccount sdk.AccAddress) *MsgCreatePool {
	return &MsgCreatePool{
		Authority:                authority.String(),
		PoolId:                   poolID,
		LiquidBondDenom:          liquidBondDenom,
		InitialAdminAddress:      initialAdmin.String(),
		InitialFeeAccountAddress: initialFeeAccount.String(),
	}
}

func (m *MsgCreatePool) Route() string { return RouterKey }
func (m *MsgCreatePool) Type() string  { return MsgTypeCreatePool }

func (m *MsgCreatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	if err := validateLiquidBondDenom(m.LiquidBondDenom); err != nil {
		return errors.Wrap(ErrInvalidLiquidBondDenom, err.Error())
	}
	// Pool creation requires a real admin: the admin is the only address
	// allowed to call LiquidStake against the pool, so a pool with empty
	// admin would be permanently inert until governance issues a follow-up
	// UpdatePool. Force the issue here.
	if err := validateFeeAccountAddress(m.InitialAdminAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "initial_admin_address: "+err.Error())
	}
	if err := validateFeeAccountAddress(m.InitialFeeAccountAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "initial_fee_account_address: "+err.Error())
	}
	return nil
}

// --------------------------
// UPDATE MODULE PARAMS
// --------------------------

func NewMsgUpdateModuleParams(authority sdk.AccAddress, params ModuleParams) *MsgUpdateModuleParams {
	return &MsgUpdateModuleParams{Authority: authority.String(), ModuleParams: params}
}

func (m *MsgUpdateModuleParams) Route() string { return RouterKey }
func (m *MsgUpdateModuleParams) Type() string  { return MsgTypeUpdateModuleParams }

func (m *MsgUpdateModuleParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	return m.ModuleParams.Validate()
}

// --------------------------
// UPDATE POOL
// --------------------------

func NewMsgUpdatePool(authority sdk.AccAddress, msg MsgUpdatePool) *MsgUpdatePool {
	return &MsgUpdatePool{
		Authority:             authority.String(),
		PoolId:                msg.PoolId,
		UnstakeFeeRate:        msg.UnstakeFeeRate,
		FeeAccountAddress:     msg.FeeAccountAddress,
		AutocompoundFeeRate:   msg.AutocompoundFeeRate,
		WhitelistAdminAddress: msg.WhitelistAdminAddress,
	}
}

func (m *MsgUpdatePool) Route() string { return RouterKey }
func (m *MsgUpdatePool) Type() string  { return MsgTypeUpdatePool }

func (m *MsgUpdatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	if err := validateUnstakeFeeRate(m.UnstakeFeeRate); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if err := validateAutocompoundFeeRate(m.AutocompoundFeeRate); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if err := validateFeeAccountAddress(m.FeeAccountAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if err := validateOptionalAdminAddress(m.WhitelistAdminAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	return nil
}

// --------------------------
// UPDATE WHITELISTED VALIDATORS
// --------------------------

func NewMsgUpdateWhitelistedValidators(authority sdk.AccAddress, poolID string, list []WhitelistedValidator) *MsgUpdateWhitelistedValidators {
	return &MsgUpdateWhitelistedValidators{
		Authority:             authority.String(),
		PoolId:                poolID,
		WhitelistedValidators: list,
	}
}

func (m *MsgUpdateWhitelistedValidators) Route() string { return RouterKey }
func (m *MsgUpdateWhitelistedValidators) Type() string  { return MsgTypeUpdateWhitelistedValidators }

func (m *MsgUpdateWhitelistedValidators) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	if err := ValidateWhitelistedValidators(m.WhitelistedValidators); err != nil {
		return errors.Wrap(ErrWhitelistedValidatorsList, err.Error())
	}
	return nil
}

// --------------------------
// UPDATE WEIGHTED REWARDS RECEIVERS
// --------------------------

func NewMsgUpdateWeightedRewardsReceivers(authority sdk.AccAddress, poolID string, list []WeightedAddress) *MsgUpdateWeightedRewardsReceivers {
	return &MsgUpdateWeightedRewardsReceivers{
		Authority:                authority.String(),
		PoolId:                   poolID,
		WeightedRewardsReceivers: list,
	}
}

func (m *MsgUpdateWeightedRewardsReceivers) Route() string { return RouterKey }
func (m *MsgUpdateWeightedRewardsReceivers) Type() string  { return MsgTypeUpdateWeightedRewardsReceivers }

func (m *MsgUpdateWeightedRewardsReceivers) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	if err := ValidateWeightedRewardsReceivers(m.WeightedRewardsReceivers); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return nil
}

// --------------------------
// SET POOL PAUSED
// --------------------------

func NewMsgSetPoolPaused(authority sdk.AccAddress, poolID string, isPaused bool) *MsgSetPoolPaused {
	return &MsgSetPoolPaused{
		Authority: authority.String(),
		PoolId:    poolID,
		IsPaused:  isPaused,
	}
}

func (m *MsgSetPoolPaused) Route() string { return RouterKey }
func (m *MsgSetPoolPaused) Type() string  { return MsgTypeSetPoolPaused }

func (m *MsgSetPoolPaused) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	if err := ValidatePoolID(m.PoolId); err != nil {
		return errors.Wrap(ErrInvalidPoolID, err.Error())
	}
	return nil
}

// --------------------------
// SET MODULE PAUSED (global)
// --------------------------

func NewMsgSetModulePaused(authority sdk.AccAddress, isPaused bool) *MsgSetModulePaused {
	return &MsgSetModulePaused{Authority: authority.String(), IsPaused: isPaused}
}

func (m *MsgSetModulePaused) Route() string { return RouterKey }
func (m *MsgSetModulePaused) Type() string  { return MsgTypeSetModulePaused }

func (m *MsgSetModulePaused) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}
	return nil
}

// --------------------------
// BURN
// --------------------------

func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coin) *MsgBurn {
	return &MsgBurn{Burner: burner.String(), Amount: amount}
}

func (m *MsgBurn) Route() string { return RouterKey }
func (m *MsgBurn) Type() string  { return MsgTypeBurn }

func (m *MsgBurn) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Burner); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid burner address %q: %v", m.Burner, err)
	}
	if m.Amount.IsZero() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "burning amount must not be zero")
	}
	if err := m.Amount.Validate(); err != nil {
		return err
	}
	if m.Amount.Denom != "uixo" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "burning amount must be in uixo")
	}
	return nil
}
