package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v5/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v5/x/iid/types"
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

func (msg MsgCreateCollection) Type() string { return TypeMsgCreateCollection }

func (msg MsgCreateCollection) Route() string { return RouterKey }

// --------------------------
// SUBMIT CLAIM
// --------------------------
const TypeMsgSubmitClaim = "submit_claim"

var _ sdk.Msg = &MsgSubmitClaim{}

func (msg MsgSubmitClaim) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgSubmitClaim) Type() string { return TypeMsgSubmitClaim }

func (msg MsgSubmitClaim) Route() string { return RouterKey }

// --------------------------
// EVALUATE CLAIM
// --------------------------
const TypeMsgEvaluateClaim = "evaluate_claim"

var _ sdk.Msg = &MsgEvaluateClaim{}

func (msg MsgEvaluateClaim) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgEvaluateClaim) Type() string { return TypeMsgEvaluateClaim }

func (msg MsgEvaluateClaim) Route() string { return RouterKey }

// --------------------------
// DISPUTE CLAIM
// --------------------------
const TypeMsgDisputeClaim = "dispute_claim"

var _ sdk.Msg = &MsgDisputeClaim{}

func (msg MsgDisputeClaim) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgDisputeClaim) Type() string { return TypeMsgDisputeClaim }

func (msg MsgDisputeClaim) Route() string { return RouterKey }

// --------------------------
// WITHDRAW PAYMENT
// --------------------------
const TypeMsgWithdrawPayment = "withdraw_payment"

var _ sdk.Msg = &MsgWithdrawPayment{}

func (msg MsgWithdrawPayment) Type() string { return TypeMsgWithdrawPayment }

func (msg MsgWithdrawPayment) Route() string { return RouterKey }

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
const TypeMsgUpdateCollectionState = "update_collection_state"

var _ sdk.Msg = &MsgUpdateCollectionState{}

func (msg MsgUpdateCollectionState) Type() string { return TypeMsgUpdateCollectionState }

func (msg MsgUpdateCollectionState) Route() string { return RouterKey }

// --------------------------
// UPDATE COLLECTION DATES
// --------------------------
const TypeMsgUpdateCollectionDates = "update_collection_dates"

var _ sdk.Msg = &MsgUpdateCollectionDates{}

func (msg MsgUpdateCollectionDates) Type() string { return TypeMsgUpdateCollectionDates }

func (msg MsgUpdateCollectionDates) Route() string { return RouterKey }

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
const TypeMsgUpdateCollectionPayments = "update_collection_payments"

var _ sdk.Msg = &MsgUpdateCollectionPayments{}

func (msg MsgUpdateCollectionPayments) Type() string { return TypeMsgUpdateCollectionPayments }

func (msg MsgUpdateCollectionPayments) Route() string { return RouterKey }

// --------------------------
// UPDATE COLLECTION INTENTS
// --------------------------
const TypeMsgUpdateCollectionIntents = "update_collection_intents"

var _ sdk.Msg = &MsgUpdateCollectionIntents{}

func (msg MsgUpdateCollectionIntents) Type() string { return TypeMsgUpdateCollectionIntents }

func (msg MsgUpdateCollectionIntents) Route() string { return RouterKey }

// --------------------------
// CLAIM INTENT
// --------------------------
const TypeMsgClaimIntent = "claim_intent"

var _ sdk.Msg = &MsgClaimIntent{}

func (msg MsgClaimIntent) Type() string { return TypeMsgClaimIntent }

func (msg MsgClaimIntent) GetIidController() iidtypes.DIDFragment { return msg.AgentDid }

func (msg MsgClaimIntent) Route() string { return RouterKey }

// --------------------------
// CREATE CLAIM AUTHORIZATION
// --------------------------
const TypeMsgCreateClaimAuthorization = "create_claim_authorization"

var _ sdk.Msg = &MsgCreateClaimAuthorization{}

func (msg MsgCreateClaimAuthorization) Type() string { return TypeMsgCreateClaimAuthorization }

func (msg MsgCreateClaimAuthorization) GetIidController() iidtypes.DIDFragment { return msg.CreatorDid }

func (msg MsgCreateClaimAuthorization) Route() string { return RouterKey }
