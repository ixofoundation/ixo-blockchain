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
// MsgEvaluateClaim and MsgCreateClaimAuthorization have proto signer
// admin_address (the collection admin / authorizer) while their *_did field
// points at a DIFFERENT party (the agent / creator) and is attribution only.
// They MUST NOT be IidTxMsg — their authorization is enforced in the keeper via
// collection.Admin == admin_address (+ authz grants), which holds on every
// route. Re-adding GetIidController to any of them would make the IID ante
// wrongly require the admin to control the agent's DID and break authz-delegated
// claim submission/evaluation.
//
// MsgDisputeClaim and MsgAdjudicateDispute MUST remain IidTxMsg: their proto
// signer IS the party identified by the DID (agent_address↔AgentDid,
// adjudicator_address↔AdjudicatorDid), so the ante check is correct for them.
func TestClaimsIidTxMsgMembership(t *testing.T) {
	notIidTxMsg := []any{
		&types.MsgSubmitClaim{},
		&types.MsgEvaluateClaim{},
		&types.MsgCreateClaimAuthorization{},
	}
	for _, m := range notIidTxMsg {
		_, ok := m.(iidante.IidTxMsg)
		require.Falsef(t, ok, "%T must NOT implement IidTxMsg (signer is admin_address; *_did is attribution only)", m)
	}

	mustBeIidTxMsg := []any{
		&types.MsgDisputeClaim{},
		&types.MsgAdjudicateDispute{},
	}
	for _, m := range mustBeIidTxMsg {
		_, ok := m.(iidante.IidTxMsg)
		require.Truef(t, ok, "%T must implement IidTxMsg (signer is the DID's own party)", m)
	}
}
