package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (p *CW20Payment) Validate(allowZero bool) error {
	if p == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "cw20 payment cannot be nil")
	}
	_, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
	}
	if !allowZero && p.Amount == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "cw20 payment amount cannot be 0")
	}
	return nil
}

// ValidateCW20Payments validates CW20Payments by checking if:
// - addresses are valid
// - amounts are positive (default as type is Uint)
// - no duplicates in addresses
func ValidateCW20Payments(p []*CW20Payment, allowZero bool) error {
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
		if err := cw20Payment.Validate(allowZero); err != nil {
			return err
		}
	}

	return nil
}

// IsZero returns true if there are no payments or all payments are zero.
func IsZeroCW20Payments(cw20Payments []*CW20Payment) bool {
	if len(cw20Payments) == 0 {
		return true
	}
	for _, cw20Payment := range cw20Payments {
		if cw20Payment.Amount != 0 {
			return false
		}
	}
	return true
}

// Checks if the CW20 payments are within the max constraints provided
func IsCW20PaymentsInMaxConstraints(cw20Payments []*CW20Payment, maxCw20Payments []*CW20Payment) bool {
	maxPaymentsMap := make(map[string]*CW20Payment)
	for _, maxPayment := range maxCw20Payments {
		maxPaymentsMap[maxPayment.Address] = maxPayment
	}

	for _, payment := range cw20Payments {
		maxPayment, ok := maxPaymentsMap[payment.Address]
		if !ok || payment.Amount > maxPayment.Amount {
			return false
		}
	}
	return true
}
