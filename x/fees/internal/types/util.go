package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"strings"
)

func NewMsgSetFeeContractAuthorisation(feeContractId string, authorised bool,
	payerDid sovrin.SovrinDid) MsgSetFeeContractAuthorisation {
	return MsgSetFeeContractAuthorisation{
		SignBytes:     "",
		PubKey:        payerDid.VerifyKey,
		PayerDid:      payerDid.Did,
		FeeContractId: feeContractId,
		Authorised:    authorised,
	}
}

func NewMsgCreateFee(feeId string, feeContent FeeContent,
	creatorDid sovrin.SovrinDid) MsgCreateFee {
	return MsgCreateFee{
		SignBytes:  "",
		PubKey:     creatorDid.VerifyKey,
		CreatorDid: creatorDid.Did,
		FeeId:      feeId,
		FeeContent: feeContent,
	}
}

func NewMsgCreateFeeContract(feeId, feeContractId string, payer sdk.AccAddress,
	canDeauthorise bool, discountId sdk.Uint, creatorDid sovrin.SovrinDid) MsgCreateFeeContract {
	return MsgCreateFeeContract{
		SignBytes:      "",
		PubKey:         creatorDid.VerifyKey,
		CreatorDid:     creatorDid.Did,
		FeeId:          feeId,
		FeeContractId:  feeContractId,
		Payer:          payer,
		CanDeauthorise: canDeauthorise,
		DiscountId:     discountId,
	}
}

func NewMsgCreateSubscription(subscriptionId string, subscriptionContent SubscriptionContent,
	creatorDid sovrin.SovrinDid) MsgCreateSubscription {
	return MsgCreateSubscription{
		SignBytes:           "",
		PubKey:              creatorDid.VerifyKey,
		CreatorDid:          creatorDid.Did,
		SubscriptionId:      subscriptionId,
		SubscriptionContent: subscriptionContent,
	}
}

func NewMsgGrantFeeDiscount(feeContractId string, discountId sdk.Uint,
	recipient sdk.AccAddress, creatorDid sovrin.SovrinDid) MsgGrantFeeDiscount {
	return MsgGrantFeeDiscount{
		SignBytes:     "",
		PubKey:        creatorDid.VerifyKey,
		SenderDid:     creatorDid.Did,
		FeeContractId: feeContractId,
		DiscountId:    discountId,
		Recipient:     recipient,
	}
}

func NewMsgRevokeFeeDiscount(feeContractId string, holder sdk.AccAddress,
	creatorDid sovrin.SovrinDid) MsgRevokeFeeDiscount {
	return MsgRevokeFeeDiscount{
		SignBytes:     "",
		PubKey:        creatorDid.VerifyKey,
		SenderDid:     creatorDid.Did,
		FeeContractId: feeContractId,
		Holder:        holder,
	}
}

func NewMsgChargeFee(feeContractId string, creatorDid sovrin.SovrinDid) MsgChargeFee {
	return MsgChargeFee{
		SignBytes:     "",
		PubKey:        creatorDid.VerifyKey,
		SenderDid:     creatorDid.Did,
		FeeContractId: feeContractId,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
