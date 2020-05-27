package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"strings"
)

func NewMsgSetFeeContractAuthorisation(feeContractId uint64, authorised bool,
	payerDid sovrin.SovrinDid) MsgSetFeeContractAuthorisation {
	return MsgSetFeeContractAuthorisation{
		SignBytes:     "",
		PubKey:        payerDid.VerifyKey,
		PayerDid:      payerDid.Did,
		FeeContractId: feeContractId,
		Authorised:    authorised,
	}
}

func NewMsgCreateFee(feeContent FeeContent,
	creatorDid sovrin.SovrinDid) MsgCreateFee {
	return MsgCreateFee{
		SignBytes:  "",
		PubKey:     creatorDid.VerifyKey,
		CreatorDid: creatorDid.Did,
		FeeContent: feeContent,
	}
}

func NewMsgCreateFeeContract(feeId uint64, payer sdk.AccAddress,
	canDeauthorise bool, creatorDid sovrin.SovrinDid) MsgCreateFeeContract {
	return MsgCreateFeeContract{
		SignBytes:      "",
		PubKey:         creatorDid.VerifyKey,
		CreatorDid:     creatorDid.Did,
		FeeId:          feeId,
		Payer:          payer,
		CanDeauthorise: canDeauthorise,
	}
}
func NewMsgGrantFeeDiscount(feeContractId, discountId uint64,
	recipient sdk.AccAddress, creatorDid sovrin.SovrinDid) MsgGrantFeeDiscount {
	return MsgGrantFeeDiscount{
		SignBytes:             "",
		PubKey:                creatorDid.VerifyKey,
		FeeContractCreatorDid: creatorDid.Did,
		FeeContractId:         feeContractId,
		DiscountId:            discountId,
		Recipient:             recipient,
	}
}

func NewMsgRevokeFeeDiscount(feeContractId, discountId uint64,
	holder sdk.AccAddress, creatorDid sovrin.SovrinDid) MsgRevokeFeeDiscount {
	return MsgRevokeFeeDiscount{
		SignBytes:             "",
		PubKey:                creatorDid.VerifyKey,
		FeeContractCreatorDid: creatorDid.Did,
		FeeContractId:         feeContractId,
		DiscountId:            discountId,
		Holder:                holder,
	}
}

func NewMsgChargeFee(feeContractId uint64, creatorDid sovrin.SovrinDid) MsgChargeFee {
	return MsgChargeFee{
		SignBytes:             "",
		PubKey:                creatorDid.VerifyKey,
		FeeContractCreatorDid: creatorDid.Did,
		FeeContractId:         feeContractId,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
