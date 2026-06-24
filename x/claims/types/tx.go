package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v7/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

// IidTxMsg subjects a message to the IID ante check "the proto signer must
// control the GetIidController DID". Among claims messages only
// MsgAdjudicateDispute is safe to subject to it: its signer is
// adjudicator_address and the keeper itself already requires the signer to
// control AdjudicatorDid (AuthorizeAdjudicator), so the two always match.
//
// MsgSubmitClaim, MsgEvaluateClaim, MsgCreateClaimAuthorization, MsgClaimIntent
// and MsgDisputeClaim are intentionally NOT IidTxMsg. Their *_did field
// (agent_did / creator_did) is attribution only and is legitimately DECOUPLED
// from the proto signer: submit/evaluate/createauth are signed by admin_address
// (the collection admin), and intent/dispute are signed by agent_address, which
// may be a delegated agent or an entity module account (e.g. a "fee" module
// account that receives the payout) whose address differs from the agent_did on
// the signed VC. Authorization is enforced in the keeper (collection.Admin /
// SubmitClaimAuthorization grant / dispute deposit), which holds on every route
// (top-level, authz.MsgExec, ICA, wasm). Subjecting them to the IID ante would
// wrongly require the signer to control the agent's DID and break delegated /
// on-behalf claims (e.g. the SUPA onboarding fee flow's module-account agent).
var (
	_ iidante.IidTxMsg = &MsgAdjudicateDispute{}
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

// NOTE: not IidTxMsg — authorized via collection.Admin in the keeper. AgentDid
// is attribution only (see the IidTxMsg comment block above).

func (msg MsgSubmitClaim) Type() string { return TypeMsgSubmitClaim }

func (msg MsgSubmitClaim) Route() string { return RouterKey }

// --------------------------
// EVALUATE CLAIM
// --------------------------
const TypeMsgEvaluateClaim = "evaluate_claim"

var _ sdk.Msg = &MsgEvaluateClaim{}

// NOTE: not IidTxMsg — authorized via collection.Admin in the keeper. AgentDid
// is attribution only (see the IidTxMsg comment block above).

func (msg MsgEvaluateClaim) Type() string { return TypeMsgEvaluateClaim }

func (msg MsgEvaluateClaim) Route() string { return RouterKey }

// --------------------------
// DISPUTE CLAIM
// --------------------------
const TypeMsgDisputeClaim = "dispute_claim"

var _ sdk.Msg = &MsgDisputeClaim{}

// NOTE: not IidTxMsg — signer agent_address may be a delegated/module-account
// agent whose address differs from AgentDid (attribution). Authorized in the
// keeper via the dispute deposit + collection membership.

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

// NOTE: not IidTxMsg — signer agent_address may be a delegated/module-account
// agent (e.g. the SUPA onboarding "fee" module account) whose address differs
// from AgentDid (attribution). Authorized in the keeper via the agent's
// SubmitClaimAuthorization grant from collection.Admin.

func (msg MsgClaimIntent) Type() string { return TypeMsgClaimIntent }

func (msg MsgClaimIntent) Route() string { return RouterKey }

// --------------------------
// CREATE CLAIM AUTHORIZATION
// --------------------------
const TypeMsgCreateClaimAuthorization = "create_claim_authorization"

var _ sdk.Msg = &MsgCreateClaimAuthorization{}

// NOTE: not IidTxMsg — authorized via collection.Admin in the keeper. CreatorDid
// is attribution only (see the IidTxMsg comment block above).

func (msg MsgCreateClaimAuthorization) Type() string { return TypeMsgCreateClaimAuthorization }

func (msg MsgCreateClaimAuthorization) Route() string { return RouterKey }

// --------------------------
// UPDATE COLLECTION DISPUTE CONFIG
// --------------------------
const TypeMsgUpdateCollectionDisputeConfig = "update_collection_dispute_config"

var _ sdk.Msg = &MsgUpdateCollectionDisputeConfig{}

func (msg MsgUpdateCollectionDisputeConfig) Type() string {
	return TypeMsgUpdateCollectionDisputeConfig
}

func (msg MsgUpdateCollectionDisputeConfig) Route() string { return RouterKey }

// --------------------------
// ADD PERFORMANCE DEPOSIT
// --------------------------
const TypeMsgAddPerformanceDeposit = "add_performance_deposit"

var _ sdk.Msg = &MsgAddPerformanceDeposit{}

func (msg MsgAddPerformanceDeposit) Type() string { return TypeMsgAddPerformanceDeposit }

func (msg MsgAddPerformanceDeposit) Route() string { return RouterKey }

// --------------------------
// WITHDRAW PERFORMANCE DEPOSIT
// --------------------------
const TypeMsgWithdrawPerformanceDeposit = "withdraw_performance_deposit"

var _ sdk.Msg = &MsgWithdrawPerformanceDeposit{}

func (msg MsgWithdrawPerformanceDeposit) Type() string { return TypeMsgWithdrawPerformanceDeposit }

func (msg MsgWithdrawPerformanceDeposit) Route() string { return RouterKey }

// --------------------------
// ADJUDICATE DISPUTE
// --------------------------
const TypeMsgAdjudicateDispute = "adjudicate_dispute"

var _ sdk.Msg = &MsgAdjudicateDispute{}

func (msg MsgAdjudicateDispute) Type() string { return TypeMsgAdjudicateDispute }

func (msg MsgAdjudicateDispute) GetIidController() iidtypes.DIDFragment {
	return iidtypes.DIDFragment(msg.AdjudicatorDid)
}

func (msg MsgAdjudicateDispute) Route() string { return RouterKey }
