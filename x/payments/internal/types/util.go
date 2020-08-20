package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"strings"
)

func NewMsgSetPaymentContractAuthorisation(contractId string, authorised bool,
	payerDid did.Did) MsgSetPaymentContractAuthorisation {
	return MsgSetPaymentContractAuthorisation{
		PayerDid:          payerDid,
		PaymentContractId: contractId,
		Authorised:        authorised,
	}
}

func NewMsgCreatePaymentTemplate(template PaymentTemplate,
	creatorDid did.Did) MsgCreatePaymentTemplate {
	return MsgCreatePaymentTemplate{
		CreatorDid:      creatorDid,
		PaymentTemplate: template,
	}
}

func NewMsgCreatePaymentContract(templateId, contractId string,
	payer sdk.AccAddress, canDeauthorise bool, discountId sdk.Uint,
	creatorDid did.Did) MsgCreatePaymentContract {
	return MsgCreatePaymentContract{
		CreatorDid:        creatorDid,
		PaymentTemplateId: templateId,
		PaymentContractId: contractId,
		Payer:             payer,
		CanDeauthorise:    canDeauthorise,
		DiscountId:        discountId,
	}
}

func NewMsgCreateSubscription(subscriptionId, contractId string, maxPeriods sdk.Uint,
	period Period, creatorDid did.Did) MsgCreateSubscription {
	return MsgCreateSubscription{
		CreatorDid:        creatorDid,
		SubscriptionId:    subscriptionId,
		PaymentContractId: contractId,
		MaxPeriods:        maxPeriods,
		Period:            period,
	}
}

func NewMsgGrantDiscount(contractId string, discountId sdk.Uint,
	recipient sdk.AccAddress, creatorDid did.Did) MsgGrantDiscount {
	return MsgGrantDiscount{
		SenderDid:         creatorDid,
		PaymentContractId: contractId,
		DiscountId:        discountId,
		Recipient:         recipient,
	}
}

func NewMsgRevokeDiscount(contractId string, holder sdk.AccAddress,
	creatorDid did.Did) MsgRevokeDiscount {
	return MsgRevokeDiscount{
		SenderDid:         creatorDid,
		PaymentContractId: contractId,
		Holder:            holder,
	}
}

func NewMsgEffectPayment(contractId string, creatorDid did.Did) MsgEffectPayment {
	return MsgEffectPayment{
		SenderDid:         creatorDid,
		PaymentContractId: contractId,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err error) {
	if strings.TrimSpace(value) == "" {
		return false, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "%s is empty", name)
	} else {
		return true, nil
	}
}
