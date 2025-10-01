package types

import (
	fmt "fmt"
	"slices"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/token/types/contracts/ixo1155"
)

func (p *CW1155Payment) Validate(allowZero bool) error {
	if p != nil {
		_, err := sdk.AccAddressFromBech32(p.Address)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
		}
		for _, tokenId := range p.TokenId {
			if iidtypes.IsEmpty(tokenId) {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "token_id cannot be empty")
			}
		}
		if !allowZero && p.Amount == 0 {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "payment amount cannot be 0")
		}
	}
	return nil
}

// ValidateCW1155Payments validates CW1155Payments by checking if:
// - addresses are valid
// - token ids are valid (not empty)
// - amounts are positive (default as type is Uint)
// - no duplicates in addresses
func ValidateCW1155Payments(p []*CW1155Payment, allowZero bool) error {
	if len(p) == 0 {
		return nil
	}

	// validate each payment and ensure no duplicates
	addresses := make(map[string]struct{}, len(p))
	for _, cw1155Payment := range p {
		if _, ok := addresses[cw1155Payment.Address]; ok {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cw1155 payments cannot have duplicate addresses (%s)", cw1155Payment.Address)
		}
		addresses[cw1155Payment.Address] = struct{}{}
		if err := cw1155Payment.Validate(allowZero); err != nil {
			return err
		}
	}
	return nil
}

// IsZero returns true if there are no payments or all payments are zero.
func IsZeroCW1155Payments(cw1155Payments []*CW1155Payment) bool {
	if len(cw1155Payments) == 0 {
		return true
	}
	for _, cw1155Payment := range cw1155Payments {
		if cw1155Payment.Amount != 0 {
			return false
		}
	}
	return true
}

// Checks if the CW1155 payments are within the max constraints provided
func IsCW1155PaymentsInMaxConstraints(cw1155Payments []*CW1155Payment, maxCw1155Payments []*CW1155Payment) bool {
	maxPaymentsMap := make(map[string]*CW1155Payment)
	for _, maxPayment := range maxCw1155Payments {
		maxPaymentsMap[maxPayment.Address] = maxPayment
	}

	for _, payment := range cw1155Payments {
		maxPayment, ok := maxPaymentsMap[payment.Address]
		// first check if address is in max payments and amount is less than max amount
		if !ok || payment.Amount > maxPayment.Amount {
			return false
		}
		// if max payments has token ids, then check if token ids are in max payments
		// otherwise it means any token id is allowed as no restrictions on max payments
		// if max payments has at least 1 token id, then payment token ids can't be empty as it must conform to max payments
		if len(maxPayment.TokenId) > 0 {
			if len(payment.TokenId) == 0 {
				return false
			}
			for _, tokenId := range payment.TokenId {
				if !slices.Contains(maxPayment.TokenId, tokenId) {
					return false
				}
			}
		}
	}
	return true
}

func (batch *CW1155IntentPaymentToken) GetWasmTransferBatch() ixo1155.Batch {
	return []string{batch.TokenId, fmt.Sprint(batch.Amount), ""}
}
