package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (p *CW20Payment) Validate() error {
	if p == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "cw20 payment cannot be nil")
	}
	_, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
	}
	return nil
}

// ValidateCW20Payments validates CW20Payments by checking if:
// - addresses are valid
// - amounts are positive (default as type is Uint)
// - no duplicates in addresses
func ValidateCW20Payments(p []*CW20Payment) error {
	if len(p) == 0 {
		return nil
	}

	// check if addresses are valid and no duplicates
	addresses := make(map[string]struct{}, len(p))
	for _, cw20Payment := range p {
		if _, ok := addresses[cw20Payment.Address]; ok {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cw20 payments cannot have duplicate addresses (%s)", cw20Payment.Address)
		}
		addresses[cw20Payment.Address] = struct{}{}
		if err := cw20Payment.Validate(); err != nil {
			return err
		}
	}

	return nil
}
