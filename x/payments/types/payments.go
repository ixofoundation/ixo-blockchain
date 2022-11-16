package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (pt PaymentTemplate) GetDiscountPercent(discountId sdk.Uint) (sdk.Dec, error) {
	for _, discount := range pt.Discounts {
		if discount.Id.Equal(discountId) {
			return discount.Percent, nil
		}
	}
	return sdk.Dec{}, ErrDiscountIdIsNotInTemplate
}

func (pt PaymentTemplate) Validate() error {
	// Validate ID
	if !IsValidPaymentTemplateId(pt.Id) {
		return sdkerrors.Wrap(ErrInvalidId, "payment template ID invalid")
	}

	// Validate payment amount, minimum, maximum
	amt := &pt.PaymentAmount
	min := &pt.PaymentMinimum
	max := &pt.PaymentMaximum
	if !amt.IsValid() {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "PaymentAmount coins invalid")
	} else if !min.IsValid() {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "PaymentMinimum coins invalid")
	} else if !max.IsValid() {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "PaymentMaximum coins invalid")
	} else if min.IsAnyGT(*max) {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "min pay includes value greater than max pay")
	} else if !min.DenomsSubsetOf(*amt) {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "min pay includes denom not in pay amount")
	} else if !max.DenomsSubsetOf(*amt) {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "max pay includes denom not in pay amount")
	}

	// Validate discounts
	if err := Discounts(pt.Discounts).Validate(); err != nil {
		return err
	}

	return nil
}

func NewPaymentContract(id, templateId string, creator, payer sdk.AccAddress,
	recipients Distribution, canDeauthorise, authorised bool,
	discountId sdk.Uint) PaymentContract {
	return PaymentContract{
		Id:                id,
		PaymentTemplateId: templateId,
		Creator:           creator.String(),
		Payer:             payer.String(),
		Recipients:        recipients,
		CumulativePay:     sdk.NewCoins(),
		CurrentRemainder:  sdk.NewCoins(),
		CanDeauthorise:    canDeauthorise,
		Authorised:        authorised,
		DiscountId:        discountId,
	}
}

func NewPaymentContractNoDiscount(id, templateId string, creator,
	payer sdk.AccAddress, recipients Distribution, canDeauthorise,
	authorised bool) PaymentContract {
	return NewPaymentContract(
		id, templateId, creator, payer, recipients,
		canDeauthorise, authorised, sdk.ZeroUint(),
	)
}

func (pc PaymentContract) Validate() error {
	// Validate ID
	if !IsValidPaymentContractId(pc.Id) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	// Validate coins
	if !pc.CumulativePay.IsValid() {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "CumulativePay coins invalid")
	} else if !pc.CurrentRemainder.IsValid() {
		return sdkerrors.Wrap(ErrInvalidPaymentTemplate, "CurrentRemainder coins invalid")
	}

	// Validate addresses
	if strings.TrimSpace(pc.Creator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty creator address")
	} else if strings.TrimSpace(pc.Payer) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty payer address")
	}

	// Validate IDs
	if !IsValidPaymentTemplateId(pc.PaymentTemplateId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment template ID invalid")
	}

	// Validate recipient distribution
	if err := Distribution(pc.Recipients).Validate(); err != nil {
		return err
	}

	return nil
}

func (pc PaymentContract) IsFirstPayment() bool {
	return pc.CumulativePay.IsZero()
}

// CanEffectPayment False if not authorised or the (non-zero!) max has been reached
func (pc PaymentContract) CanEffectPayment(template PaymentTemplate) bool {
	if template.Id != pc.PaymentTemplateId {
		panic("payment template ID mismatch in CanEffectPayment")
	}
	max := template.PaymentMaximum
	return pc.Authorised && (max.IsZero() || max.IsAllGT(pc.CumulativePay))
}
