package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v2/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v2/x/iid/types"
)

var (
	_ iidante.IidTxMsg = &MsgSubmitClaim{}
	_ iidante.IidTxMsg = &MsgEvaluateClaim{}
	_ iidante.IidTxMsg = &MsgDisputeClaim{}
)

// --------------------------
// CREATE COLLECTION
// --------------------------
const TypeMsgCreateCollection = "create_collection"

var _ sdk.Msg = &MsgCreateCollection{}

func (msg MsgCreateCollection) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateCollection) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateCollection) Type() string { return TypeMsgCreateCollection }

func (msg MsgCreateCollection) Route() string { return RouterKey }

// --------------------------
// SUBMIT CLAIM
// --------------------------
const TypeMsgSubmitClaim = "submit_claim"

var _ sdk.Msg = &MsgSubmitClaim{}

func (msg MsgSubmitClaim) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgSubmitClaim) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgSubmitClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSubmitClaim) Type() string { return TypeMsgSubmitClaim }

func (msg MsgSubmitClaim) Route() string { return RouterKey }

// --------------------------
// EVALUATE CLAIM
// --------------------------
const TypeMsgEvaluateClaim = "evaluate_claim"

var _ sdk.Msg = &MsgEvaluateClaim{}

func (msg MsgEvaluateClaim) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgEvaluateClaim) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgEvaluateClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgEvaluateClaim) Type() string { return TypeMsgEvaluateClaim }

func (msg MsgEvaluateClaim) Route() string { return RouterKey }

// --------------------------
// DISPUTE CLAIM
// --------------------------
const TypeMsgDisputeClaim = "dispute_claim"

var _ sdk.Msg = &MsgDisputeClaim{}

func (msg MsgDisputeClaim) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgDisputeClaim) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgDisputeClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDisputeClaim) Type() string { return TypeMsgDisputeClaim }

func (msg MsgDisputeClaim) Route() string { return RouterKey }

// --------------------------
// WITHDRAW PAYMENT
// --------------------------
const TypeMsgWithdrawPayment = "withdraw_payment"

var _ sdk.Msg = &MsgWithdrawPayment{}

func (msg MsgWithdrawPayment) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgWithdrawPayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawPayment) Type() string { return TypeMsgWithdrawPayment }

func (msg MsgWithdrawPayment) Route() string { return RouterKey }
