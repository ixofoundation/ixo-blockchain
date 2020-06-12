package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"strings"
)

func NewMsgSetPaymentContractAuthorisation(contractId string, authorised bool,
	payerDid sovrin.SovrinDid) MsgSetPaymentContractAuthorisation {
	return MsgSetPaymentContractAuthorisation{
		PubKey:            payerDid.VerifyKey,
		PayerDid:          payerDid.Did,
		PaymentContractId: contractId,
		Authorised:        authorised,
	}
}

func NewMsgCreatePaymentTemplate(template PaymentTemplate,
	creatorDid sovrin.SovrinDid) MsgCreatePaymentTemplate {
	return MsgCreatePaymentTemplate{
		PubKey:          creatorDid.VerifyKey,
		CreatorDid:      creatorDid.Did,
		PaymentTemplate: template,
	}
}

func NewMsgCreatePaymentContract(templateId, contractId string,
	payer sdk.AccAddress, canDeauthorise bool, discountId sdk.Uint,
	creatorDid sovrin.SovrinDid) MsgCreatePaymentContract {
	return MsgCreatePaymentContract{
		PubKey:            creatorDid.VerifyKey,
		CreatorDid:        creatorDid.Did,
		PaymentTemplateId: templateId,
		PaymentContractId: contractId,
		Payer:             payer,
		CanDeauthorise:    canDeauthorise,
		DiscountId:        discountId,
	}
}

func NewMsgCreateSubscription(subscriptionId, contractId string, maxPeriods sdk.Uint,
	period Period, creatorDid sovrin.SovrinDid) MsgCreateSubscription {
	return MsgCreateSubscription{
		PubKey:            creatorDid.VerifyKey,
		CreatorDid:        creatorDid.Did,
		SubscriptionId:    subscriptionId,
		PaymentContractId: contractId,
		MaxPeriods:        maxPeriods,
		Period:            period,
	}
}

func NewMsgGrantDiscount(contractId string, discountId sdk.Uint,
	recipient sdk.AccAddress, creatorDid sovrin.SovrinDid) MsgGrantDiscount {
	return MsgGrantDiscount{
		PubKey:            creatorDid.VerifyKey,
		SenderDid:         creatorDid.Did,
		PaymentContractId: contractId,
		DiscountId:        discountId,
		Recipient:         recipient,
	}
}

func NewMsgRevokeDiscount(contractId string, holder sdk.AccAddress,
	creatorDid sovrin.SovrinDid) MsgRevokeDiscount {
	return MsgRevokeDiscount{
		PubKey:            creatorDid.VerifyKey,
		SenderDid:         creatorDid.Did,
		PaymentContractId: contractId,
		Holder:            holder,
	}
}

func NewMsgEffectPayment(contractId string, creatorDid sovrin.SovrinDid) MsgEffectPayment {
	return MsgEffectPayment{
		PubKey:            creatorDid.VerifyKey,
		SenderDid:         creatorDid.Did,
		PaymentContractId: contractId,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
