package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// --------------------------
// LIQUID STAKE
// --------------------------
const MsgTypeLiquidStake = "liquid_stake"

var _ sdk.Msg = &MsgLiquidStake{}

// NewMsgLiquidStake creates a new MsgLiquidStake.
func NewMsgLiquidStake(
	liquidStaker sdk.AccAddress,
	amount sdk.Coin,
) *MsgLiquidStake {
	return &MsgLiquidStake{
		DelegatorAddress: liquidStaker.String(),
		Amount:           amount,
	}
}

func (m *MsgLiquidStake) Route() string { return RouterKey }

func (m *MsgLiquidStake) Type() string { return MsgTypeLiquidStake }

func (m *MsgLiquidStake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.DelegatorAddress); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address %q: %v", m.DelegatorAddress, err)
	}
	if ok := m.Amount.IsZero(); ok {
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
const MsgTypeLiquidUnstake = "liquid_unstake"

var _ sdk.Msg = &MsgLiquidUnstake{}

// NewMsgLiquidUnstake creates a new MsgLiquidUnstake.
func NewMsgLiquidUnstake(
	liquidStaker sdk.AccAddress,
	amount sdk.Coin,
) *MsgLiquidUnstake {
	return &MsgLiquidUnstake{
		DelegatorAddress: liquidStaker.String(),
		Amount:           amount,
	}
}

func (m *MsgLiquidUnstake) Route() string { return RouterKey }

func (m *MsgLiquidUnstake) Type() string { return MsgTypeLiquidUnstake }

func (m *MsgLiquidUnstake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.DelegatorAddress); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address %q: %v", m.DelegatorAddress, err)
	}
	if ok := m.Amount.IsZero(); ok {
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
// UPDATE PARAMS
// --------------------------
const MsgTypeUpdateParams = "update_params"

var _ sdk.Msg = &MsgUpdateParams{}

// NewMsgUpdateParams creates a new MsgUpdateParams.
func NewMsgUpdateParams(authority sdk.AccAddress, amount Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority.String(),
		Params:    amount,
	}
}

func (m *MsgUpdateParams) Route() string {
	return RouterKey
}

// Type should return the action
func (m *MsgUpdateParams) Type() string {
	return MsgTypeUpdateParams
}

func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}

	err := m.Params.Validate()
	if err != nil {
		return err
	}
	return nil
}

// --------------------------
// UPDATE WHITELISTED VALIDATORS
// --------------------------
const MsgTypeUpdateWhitelistedValidators = "update_whitelisted_validators"

var _ sdk.Msg = &MsgUpdateWhitelistedValidators{}

// NewMsgUpdateWhitelistedValidators creates a new MsgUpdateWhitelistedValidators.
func NewMsgUpdateWhitelistedValidators(authority sdk.AccAddress, list []WhitelistedValidator) *MsgUpdateWhitelistedValidators {
	return &MsgUpdateWhitelistedValidators{
		Authority:             authority.String(),
		WhitelistedValidators: list,
	}
}

func (m *MsgUpdateWhitelistedValidators) Route() string {
	return RouterKey
}

// Type should return the action
func (m *MsgUpdateWhitelistedValidators) Type() string {
	return MsgTypeUpdateWhitelistedValidators
}

func (m *MsgUpdateWhitelistedValidators) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}

	err := validateWhitelistedValidators(m.WhitelistedValidators)
	if err != nil {
		return err
	}

	return nil
}

// --------------------------
// UPDATE WEIGHTED REWARDS RECEIVERS
// --------------------------
const MsgTypeUpdateWeightedRewardsReceivers = "update_weighted_rewards_receivers"

var _ sdk.Msg = &MsgUpdateWeightedRewardsReceivers{}

// NewMsgUpdateWeightedRewardsReceivers creates a new MsgUpdateWeightedRewardsReceivers.
func NewMsgUpdateWeightedRewardsReceivers(authority sdk.AccAddress, list []WeightedAddress) *MsgUpdateWeightedRewardsReceivers {
	return &MsgUpdateWeightedRewardsReceivers{
		Authority:                authority.String(),
		WeightedRewardsReceivers: list,
	}
}

func (m *MsgUpdateWeightedRewardsReceivers) Route() string {
	return RouterKey
}

// Type should return the action
func (m *MsgUpdateWeightedRewardsReceivers) Type() string {
	return MsgTypeUpdateWeightedRewardsReceivers
}

func (m *MsgUpdateWeightedRewardsReceivers) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}

	err := validateWeightedRewardsReceivers(m.WeightedRewardsReceivers)
	if err != nil {
		return err
	}

	return nil
}

// --------------------------
// SET MODULE PAUSED
// --------------------------
const MsgTypeSetModulePaused = "set_module_paused"

var _ sdk.Msg = &MsgSetModulePaused{}

// NewMsgSetModulePaused creates a new MsgSetModulePaused.
func NewMsgSetModulePaused(authority sdk.AccAddress, isPaused bool) *MsgSetModulePaused {
	return &MsgSetModulePaused{
		Authority: authority.String(),
		IsPaused:  isPaused,
	}
}

func (m *MsgSetModulePaused) Route() string {
	return RouterKey
}

// Type should return the action
func (m *MsgSetModulePaused) Type() string {
	return MsgTypeSetModulePaused
}

func (m *MsgSetModulePaused) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}

	return nil
}

// --------------------------
// BURN
// --------------------------
const MsgTypeBurn = "burn"

var _ sdk.Msg = &MsgBurn{}

// NewMsgBurn creates a new MsgBurn.
func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coin) *MsgBurn {
	return &MsgBurn{
		Burner: burner.String(),
		Amount: amount,
	}
}

func (m *MsgBurn) Route() string { return RouterKey }

func (m *MsgBurn) Type() string { return MsgTypeBurn }

func (m *MsgBurn) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Burner); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid burner address %q: %v", m.Burner, err)
	}
	if ok := m.Amount.IsZero(); ok {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "burning amount must not be zero")
	}
	if err := m.Amount.Validate(); err != nil {
		return err
	}
	// do static validation for the Coin denom, only uixo is allowed
	if m.Amount.Denom != "uixo" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "burning amount must be in uixo")
	}
	return nil
}
