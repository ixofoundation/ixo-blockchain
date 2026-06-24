package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/x/claims/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v7/x/iid/ante"
)

// TestClaimsIidTxMsgMembership locks in which claims messages are subject to the
// IID ante's "the proto signer must control the GetIidController DID" check.
//
// Regression guard for the 2026-06 delegated-claims break: MsgSubmitClaim,
// MsgEvaluateClaim, MsgCreateClaimAuthorization, MsgClaimIntent and
// MsgDisputeClaim have a proto signer (admin_address or agent_address) while
// their *_did field points at a DIFFERENT party (the agent / creator) and is
// attribution only. The signer may be a delegated agent or an entity module
// account (e.g. the SUPA onboarding "fee" module account) whose address differs
// from that DID. They MUST NOT be IidTxMsg — authorization is enforced in the
// keeper (collection.Admin == admin_address, SubmitClaimAuthorization grant, or
// dispute deposit), which holds on every route. Re-adding GetIidController to
// any of them would make the IID ante wrongly require the signer to control the
// agent's DID and break authz-delegated / on-behalf claims.
//
// Only MsgAdjudicateDispute remains IidTxMsg: its proto signer IS the party
// identified by the DID (adjudicator_address↔AdjudicatorDid). The keeper's
// AuthorizeAdjudicator itself requires the signer to control AdjudicatorDid
// (iidDoc.HasRelationship), with no bypass — the entity branch is payout-routing
// only — so the ante check is provably consistent with the keeper and can never
// reject a flow the keeper accepts.
func TestClaimsIidTxMsgMembership(t *testing.T) {
	notIidTxMsg := []any{
		&types.MsgSubmitClaim{},
		&types.MsgEvaluateClaim{},
		&types.MsgCreateClaimAuthorization{},
		&types.MsgClaimIntent{},
		&types.MsgDisputeClaim{},
	}
	for _, m := range notIidTxMsg {
		_, ok := m.(iidante.IidTxMsg)
		require.Falsef(t, ok, "%T must NOT implement IidTxMsg (signer may be a delegated/module-account agent; *_did is attribution only)", m)
	}

	mustBeIidTxMsg := []any{
		&types.MsgAdjudicateDispute{},
	}
	for _, m := range mustBeIidTxMsg {
		_, ok := m.(iidante.IidTxMsg)
		require.Truef(t, ok, "%T must implement IidTxMsg (signer is the DID's own party; keeper enforces the same binding)", m)
	}
}
