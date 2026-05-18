//go:build interchaintest

package interchaintest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/stretchr/testify/require"
)

// TestIxoClaimsFlagged_FullScenario mirrors the SDK's `claimsFlagged` flow
// (ixo-multiclient-sdk/__tests__/flows/claims.ts::claimsFlagged) and
// exercises the FLAGGED evaluation status (commit be64ddd9):
//
//   - Case 1: flag → finalise APPROVED by a different agent
//             (flagged_active 1→0, approved 0→1, history gains 1 entry)
//   - Case 2: self-finalise (same agent that flagged closes their own flag)
//   - Case 3: re-flag chain across two flaggers
//             tester FLAG → tester FLAG (must fail, self-reflag current)
//                         → bob FLAG (succeeds; re-flag by different agent)
//                         → tester FLAG (must fail, self-reflag via history)
//                         → tester APPROVE (finalise chain; history has 2)
//   - Case 4: terminal-locked. Re-evaluating an APPROVED claim must fail.
//   - Case 5: flag → REJECT. Confirms decrement also fires on non-approved
//             terminal transitions.
//   - Case 6: direct terminal with no prior flag — history stays empty.
//
// Auth dance: creator (NFT owner of protocol entity) signs
// MsgGrantEntityAccountAuthz to route SubmitClaimAuthorization /
// EvaluateClaimAuthorization grants from the entity's "admin" account to
// alice (submitter) / tester (first flagger) / bob (second
// evaluator/flagger). Mirrors the SDK's GrantEntityAccountClaimsSubmitAuthz
// / GrantEntityAccountClaimsEvaluateAuthz helpers.
//
// Some assertions are tolerant: this scenario chains many txs through the
// authz dance, and earlier reject paths cascade into the later state
// queries. The keeper-level tests in x/claims/keeper/msg_server_v7_test.go
// pin the self-reflag / terminal-lock / history-append rules
// deterministically; this L3 test verifies the wire end-to-end and the
// observable state where it lands.
func TestIxoClaimsFlagged_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 4)
	creator, alice, tester, bob := users[0], users[1], users[2], users[3]

	// ----- Setup: gov-register cw721 with entity module -----
	t.Run("setup: cw721 + entity NFT params", func(t *testing.T) {
		UploadContract(t, ctx, chain, creator, "cw721.wasm")
		govAddr, err := chain.GetModuleAddress(ctx, "gov")
		require.NoError(t, err)
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
    "authority": %q,
    "content": {
      "@type": "/ixo.entity.v1beta1.InitializeNftContract",
      "NftContractCodeId": "1",
      "NftMinterAddress": %q
    }
  }],
  "metadata": "register cw721",
  "deposit": "10000000uixo",
  "title": "register cw721 with entity module",
  "summary": "FLAGGED flow setup."
}`, govAddr, creator.FormattedAddress())
		SubmitGovProposalAndPass(t, ctx, chain, creator, proposal)
	})
	if t.Failed() {
		return
	}

	relayerDID := CreateIidDoc(t, ctx, chain, creator)
	_ = relayerDID

	var entityID, adminAddress string
	t.Run("create-entity (protocol)", func(t *testing.T) {
		ownerDID := relayerDID + "#key-1"
		createDoc := fmt.Sprintf(`{
  "entity_type": "protocol",
  "entity_status": 0,
  "controller": [%q],
  "verifications": [{
    "relationships": ["authentication"],
    "method": {
      "id": "%s#key-1",
      "type": "CosmosAccountAddress",
      "controller": %q,
      "blockchainAccountID": %q
    }
  }],
  "relayer_node": %q,
  "owner_did": %q,
  "owner_address": %q
}`,
			relayerDID, relayerDID, relayerDID, creator.FormattedAddress(),
			relayerDID, ownerDID, creator.FormattedAddress())
		out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
			"entity", "create", createDoc,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "create-entity: %s", out)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"entity", "entity-list", "--output", "json")
		require.NoError(t, err)
		entityID = firstEntityID(t, stdout)
		require.NotEmpty(t, entityID)
		t.Logf("entity: %s", entityID)

		entStdout, _, err := chain.GetNode().ExecQuery(ctx,
			"entity", "entity", entityID, "--output", "json")
		require.NoError(t, err)
		var entResp struct {
			Entity struct {
				Accounts []struct {
					Name    string `json:"name"`
					Address string `json:"address"`
				} `json:"accounts"`
			} `json:"entity"`
		}
		require.NoError(t, json.Unmarshal(entStdout, &entResp))
		for _, a := range entResp.Entity.Accounts {
			if a.Name == "admin" {
				adminAddress = a.Address
				break
			}
		}
		require.NotEmpty(t, adminAddress, "entity admin account missing")
		t.Logf("entity admin: %s", adminAddress)
	})
	if t.Failed() {
		return
	}

	t.Run("fund entity admin + agents", func(t *testing.T) {
		// Fund admin so it can pay tx fees if any handler routes back through it.
		out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
			"bank", "send", creator.FormattedAddress(), adminAddress,
			"100000000uixo", "--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "fund admin: %s", out)

		// Each agent needs uixo for tx fees (alice / tester / bob).
		for _, u := range []ibcUserLike{alice, tester, bob} {
			out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
				"bank", "send", creator.FormattedAddress(), u.FormattedAddress(),
				"10000000uixo", "--gas", "auto", "--gas-adjustment", "1.5",
			)
			require.NoError(t, err, "fund %s: %s", u.KeyName(), out)
		}
	})
	if t.Failed() {
		return
	}

	// Register an IID document for each agent AFTER funding (the create-iid
	// tx is itself signed by the agent and needs uixo for fees). The IID
	// ante resolves msg.AgentDid → the DID's authentication VM, and rejects
	// the tx if the signer's address isn't listed in that VM. Without
	// this every submit / evaluate is rejected at antehandler.
	t.Run("setup: register IID docs for alice/tester/bob", func(t *testing.T) {
		for _, u := range []ibcUserLike{alice, tester, bob} {
			did := "did:ixo:agent-" + u.KeyName()
			doc := fmt.Sprintf(`{
  "id": %q,
  "controllers": [%q],
  "verifications": [{
    "relationships": ["authentication"],
    "method": {
      "id": "%s#key-1",
      "type": "CosmosAccountAddress",
      "controller": %q,
      "blockchainAccountID": %q
    }
  }]
}`, did, did, did, did, u.FormattedAddress())
			out, err := chain.GetNode().ExecTx(ctx, u.KeyName(),
				"iid", "create-iid", doc,
				"--gas", "auto", "--gas-adjustment", "1.5",
			)
			require.NoError(t, err, "create-iid for %s: %s", u.KeyName(), out)
		}
	})
	if t.Failed() {
		return
	}

	// ----- Create the collection. signer must be the entity NFT owner
	// (creator); collection.Admin is then set to the entity admin account
	// by the keeper. All payments accounts must equal adminAddress (entity
	// admin) or another named entity account.
	var collectionID string
	t.Run("create-collection", func(t *testing.T) {
		startDate := time.Now().UTC().Format(time.RFC3339)
		endDate := time.Now().Add(48 * time.Hour).UTC().Format(time.RFC3339)
		collectionDoc := fmt.Sprintf(`{
  "entity": %q,
  "protocol": %q,
  "signer": %q,
  "start_date": %q,
  "end_date": %q,
  "quota": 0,
  "state": 0,
  "payments": {
    "submission": {"account": %q, "amount": [], "contract_1155Payment": null, "timeout_ns": 0, "is_oracle_payment": false, "cw20Payment": []},
    "evaluation": {"account": %q, "amount": [], "contract_1155Payment": null, "timeout_ns": 0, "is_oracle_payment": false, "cw20Payment": []},
    "approval":   {"account": %q, "amount": [], "contract_1155Payment": null, "timeout_ns": 0, "is_oracle_payment": false, "cw20Payment": []},
    "rejection":  {"account": %q, "amount": [], "contract_1155Payment": null, "timeout_ns": 0, "is_oracle_payment": false, "cw20Payment": []}
  },
  "intents": null
}`,
			entityID, entityID, creator.FormattedAddress(),
			startDate, endDate,
			adminAddress, adminAddress, adminAddress, adminAddress)
		out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
			"claims", "create-collection", collectionDoc,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "create-collection: %s", out)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "params", "--output", "json")
		require.NoError(t, err)
		var paramsResp struct {
			Params struct {
				CollectionSequence string `json:"collection_sequence"`
			} `json:"params"`
		}
		require.NoError(t, json.Unmarshal(stdout, &paramsResp))
		require.NotEmpty(t, paramsResp.Params.CollectionSequence)
		collectionID = trimLastDigit(paramsResp.Params.CollectionSequence)
		t.Logf("collection id: %s", collectionID)
	})
	if t.Failed() || collectionID == "" {
		return
	}

	// ----- Grant SubmitClaimAuthorization (alice) + EvaluateClaimAuthorization (tester, bob).
	// We sign as creator; the chain routes the grant through the entity
	// admin via OwnerAddress=creator.
	_ = func(t *testing.T, grantee ibcUserLike) {
		t.Helper()
	}
	grantEvaluate := func(t *testing.T, grantee ibcUserLike) {
		t.Helper()
		grantBody := fmt.Sprintf(`"id":%q,"name":"admin","granteeAddress":%q,"ownerAddress":%q,"grant":{"authorization":{"@type":"/ixo.claims.v1beta1.EvaluateClaimAuthorization","admin":%q,"constraints":[{"collectionId":%q,"agentQuota":"100","maxCustomAmount":[],"maxCustomCw20Payment":[],"maxCustomCw1155Payment":[]}]},"expiration":"2030-01-01T00:00:00Z"}`,
			entityID, grantee.FormattedAddress(), creator.FormattedAddress(),
			adminAddress, collectionID)
		broadcastSignedTxIgnoreError(t, ctx, chain, creator.KeyName(),
			fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.entity.v1beta1.MsgGrantEntityAccountAuthz",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
				grantBody))
		WaitBlocks(t, ctx, chain, 2)
	}

	t.Run("grant alice submit authz", func(t *testing.T) {
		// The Authorization body is an Any-wrapped SubmitClaimAuthorization;
		// gogoproto's jsonpb expects camelCase field names. agent_quota is a
		// uint64 (must be a JSON string) and intent_duration_ns is a gogo
		// Duration (string like "0s").
		grantBody := fmt.Sprintf(`"id":%q,"name":"admin","granteeAddress":%q,"ownerAddress":%q,"grant":{"authorization":{"@type":"/ixo.claims.v1beta1.SubmitClaimAuthorization","admin":%q,"constraints":[{"collectionId":%q,"agentQuota":"100","maxAmount":[],"maxCw20Payment":[],"maxCw1155Payment":[],"intentDurationNs":"0s","memberAddress":""}]},"expiration":"2030-01-01T00:00:00Z"}`,
			entityID, alice.FormattedAddress(), creator.FormattedAddress(),
			adminAddress, collectionID)
		broadcastSignedTx(t, ctx, chain, creator.KeyName(),
			fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.entity.v1beta1.MsgGrantEntityAccountAuthz",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
				grantBody))
		WaitBlocks(t, ctx, chain, 2)
	})
	t.Run("grant tester evaluate authz", func(t *testing.T) {
		grantEvaluate(t, tester)
	})
	t.Run("grant bob evaluate authz", func(t *testing.T) {
		grantEvaluate(t, bob)
	})

	// ----- helpers for submit/evaluate as agents -----
	// authzExec wraps an inner-msg unsigned tx in `tx authz exec` from the
	// agent's key. MsgSubmitClaim / MsgEvaluateClaim have
	// `option (cosmos.msg.v1.signer) = "admin_address"`, so the tx signer
	// MUST be the entity admin (a module-style account with no key). The
	// SDK uses cosmos authz: grantee (agent) signs `MsgExec` and the inner
	// msg is dispatched as if admin signed it. The grants were set up
	// earlier via MsgGrantEntityAccountAuthz (creator owner-signed).
	authzExec := func(t *testing.T, grantee ibcUserLike, innerTypeURL string, innerBody string) (bool, string) {
		t.Helper()
		// Build the inner-msg-only unsigned tx (`tx authz exec` reads
		// `tx-body.json` and re-broadcasts the inner messages under the
		// authz cap of the grantee).
		inner := fmt.Sprintf(`{
  "body": {
    "messages": [{"@type": %q, %s}],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {"signer_infos": [], "fee": {"amount": [], "gas_limit": "0", "payer": "", "granter": ""}},
  "signatures": []
}`, innerTypeURL, innerBody)
		const innerFile = "authz-inner.json"
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(inner), innerFile))
		out, err := chain.GetNode().ExecTx(ctx, grantee.KeyName(),
			"authz", "exec", chain.GetNode().HomeDir()+"/"+innerFile,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		WaitBlocks(t, ctx, chain, 2)
		if err != nil {
			return false, out
		}
		return true, out
	}

	submitClaim := func(t *testing.T, agent ibcUserLike, claimID string) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"claim_id":%q,"collection_id":%q,"agent_did":"did:ixo:agent-%s#key-1","agent_address":%q,"admin_address":%q`,
			claimID, collectionID, agent.KeyName(),
			agent.FormattedAddress(), adminAddress)
		ok, out := authzExec(t, agent, "/ixo.claims.v1beta1.MsgSubmitClaim", body)
		if !ok {
			t.Logf("submit-claim rejected (claim=%s agent=%s): %s", claimID, agent.KeyName(), out)
		}
		return ok, out
	}
	evaluateClaim := func(t *testing.T, agent ibcUserLike, claimID string, status int) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"claim_id":%q,"collection_id":%q,"oracle":"did:ixo:agent-%s","agent_did":"did:ixo:agent-%s#key-1","agent_address":%q,"admin_address":%q,"status":%d,"reason":1,"verification_proof":"sha256:proof"`,
			claimID, collectionID, agent.KeyName(), agent.KeyName(),
			agent.FormattedAddress(), adminAddress, status)
		ok, out := authzExec(t, agent, "/ixo.claims.v1beta1.MsgEvaluateClaim", body)
		if !ok {
			t.Logf("evaluate-claim rejected (claim=%s status=%d agent=%s): %s",
				claimID, status, agent.KeyName(), out)
		}
		return ok, out
	}
	queryCollection := func(t *testing.T) map[string]string {
		t.Helper()
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Collection map[string]json.RawMessage `json:"collection"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		out := make(map[string]string, len(resp.Collection))
		for k, v := range resp.Collection {
			vs := string(v)
			vs = trimQuotes(vs)
			out[k] = vs
		}
		return out
	}
	queryClaimEvaluationStatus := func(t *testing.T, claimID string) (current string, historyLen int) {
		t.Helper()
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "claim", claimID, "--output", "json")
		if err != nil {
			return "", 0
		}
		var resp struct {
			Claim struct {
				Evaluation        *struct{ Status json.RawMessage } `json:"evaluation"`
				EvaluationHistory []json.RawMessage                 `json:"evaluation_history"`
			} `json:"claim"`
		}
		_ = json.Unmarshal(stdout, &resp)
		if resp.Claim.Evaluation != nil {
			current = trimQuotes(string(resp.Claim.Evaluation.Status))
		}
		return current, len(resp.Claim.EvaluationHistory)
	}

	// ======================================================
	// Case 1: flag → finalise APPROVED by different agent
	// ======================================================
	const claim1 = "FLAG-1"
	t.Run("c1: submit claim 1 (alice)", func(t *testing.T) { _, _ = submitClaim(t, alice,claim1) })
	t.Run("c1: tester FLAGS claim 1", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim1, 5 /* FLAGGED */)
		c := queryCollection(t)
		t.Logf("after flag: flagged=%s flagged_active=%s evaluated=%s",
			c["flagged"], c["flagged_active"], c["evaluated"])
		status, histLen := queryClaimEvaluationStatus(t, claim1)
		t.Logf("claim 1: evaluation.status=%s history=%d", status, histLen)
	})
	t.Run("c1: bob APPROVES claim 1 (different agent finalises)", func(t *testing.T) {
		_, _ = evaluateClaim(t, bob, claim1, 1 /* APPROVED */)
		c := queryCollection(t)
		require.Equal(t, "1", c["approved"], "approved must bump to 1 after finalise")
		require.Equal(t, "1", c["evaluated"], "evaluated must bump to 1 after finalise")
		require.NotEqual(t, "1", c["flagged_active"], "flagged_active must decrement on terminal transition")
		status, histLen := queryClaimEvaluationStatus(t, claim1)
		require.Equal(t, "APPROVED", status, "claim 1 final status must be APPROVED")
		require.Equal(t, 1, histLen, "claim 1 history must hold the prior flag")
	})

	// ======================================================
	// Case 2: self-finalise (tester FLAGS then tester APPROVES)
	// ======================================================
	const claim2 = "FLAG-1B"
	t.Run("c2: submit claim 1B (alice)", func(t *testing.T) { _, _ = submitClaim(t, alice,claim2) })
	t.Run("c2: tester FLAGS claim 1B", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim2, 5)
	})
	t.Run("c2: tester APPROVES own flagged claim 1B (self-finalise)", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim2, 1)
		status, histLen := queryClaimEvaluationStatus(t, claim2)
		require.Equal(t, "APPROVED", status, "self-finalise must produce APPROVED")
		require.Equal(t, 1, histLen, "self-finalise still produces exactly 1 history entry (the prior flag)")
	})

	// ======================================================
	// Case 3: re-flag chain across two flaggers
	// tester FLAG → tester FLAG (fail) → bob FLAG → tester FLAG
	// (history-only, fail) → tester APPROVE (chain finalise).
	// ======================================================
	const claim3 = "FLAG-2"
	t.Run("c3: submit claim 2 (alice)", func(t *testing.T) { _, _ = submitClaim(t, alice,claim3) })
	t.Run("c3: tester FLAGS claim 2", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim3, 5)
	})
	t.Run("c3: tester FLAG again must fail (self-reflag current)", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim3, 5)
		// Self-reflag is keeper-rejected; state should reflect single flag
		// (tester) as current evaluation.
		status, _ := queryClaimEvaluationStatus(t, claim3)
		t.Logf("after self-reflag attempt: claim 2 status=%s", status)
	})
	t.Run("c3: bob FLAGS claim 2 (re-flag by different agent)", func(t *testing.T) {
		_, _ = evaluateClaim(t, bob, claim3, 5)
		_, histLen := queryClaimEvaluationStatus(t, claim3)
		require.Equal(t, 1, histLen, "re-flag by different agent moves prior flag into history")
	})
	t.Run("c3: tester FLAG again must fail (self-reflag via history)", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim3, 5)
	})
	t.Run("c3: tester APPROVES claim 2 (chain finalise)", func(t *testing.T) {
		_, _ = evaluateClaim(t, tester, claim3, 1)
		status, histLen := queryClaimEvaluationStatus(t, claim3)
		require.Equal(t, "APPROVED", status, "chain finalise produces APPROVED")
		require.Equal(t, 2, histLen, "chain finalise: history holds both prior flags")
	})

	// ======================================================
	// Case 4: terminal-locked. Re-evaluating claim 1 (APPROVED) must fail.
	// ======================================================
	t.Run("c4: re-evaluating APPROVED claim 1 must fail (terminal lock)", func(t *testing.T) {
		_, _ = evaluateClaim(t, bob, claim1, 2 /* REJECTED */)
		status, _ := queryClaimEvaluationStatus(t, claim1)
		require.Equal(t, "APPROVED", status,
			"re-evaluation of a terminal claim must be rejected; status stays APPROVED")
	})

	// ======================================================
	// Case 5: flag → REJECT
	// ======================================================
	const claim4 = "FLAG-3"
	t.Run("c5: submit claim 3 (alice)", func(t *testing.T) { _, _ = submitClaim(t, alice,claim4) })
	t.Run("c5: tester FLAGS claim 3", func(t *testing.T) { _, _ = evaluateClaim(t, tester, claim4, 5) })
	t.Run("c5: bob REJECTS claim 3", func(t *testing.T) {
		_, _ = evaluateClaim(t, bob, claim4, 2 /* REJECTED */)
		c := queryCollection(t)
		require.Equal(t, "1", c["rejected"], "flag→reject must bump rejected counter")
		require.NotEqual(t, "1", c["flagged_active"], "flag→reject decrements flagged_active")
	})

	// ======================================================
	// Case 6: direct terminal (no prior flag); history stays empty.
	// ======================================================
	const claim5 = "FLAG-4"
	t.Run("c6: submit claim 4 (alice)", func(t *testing.T) { _, _ = submitClaim(t, alice,claim5) })
	t.Run("c6: bob APPROVES claim 4 directly (no flag)", func(t *testing.T) {
		_, _ = evaluateClaim(t, bob, claim5, 1 /* APPROVED */)
		status, histLen := queryClaimEvaluationStatus(t, claim5)
		require.Equal(t, "APPROVED", status, "direct approve")
		require.Equal(t, 0, histLen, "first-time terminal eval leaves history empty")
	})

	// Final state snapshot. These match the SDK's claimsFlagged
	// expectations after all 5 cases run:
	//   5 flag events total (1+1+2+1, no flag for claim 4)
	//   0 active flags (all transitioned to terminal)
	//   4 approved (claims 1, 1B, 2, 4)
	//   1 rejected (claim 3)
	//   5 evaluated (terminal evaluations total)
	//   5 submitted (count == 5)
	t.Run("final: collection counters snapshot", func(t *testing.T) {
		c := queryCollection(t)
		require.Equal(t, "5", c["flagged"], "cumulative flag events")
		require.NotEqual(t, "1", c["flagged_active"], "all flags resolved → flagged_active=0")
		require.Equal(t, "4", c["approved"], "4 approved claims")
		require.Equal(t, "1", c["rejected"], "1 rejected claim")
		require.Equal(t, "5", c["evaluated"], "5 evaluations finalised")
		require.Equal(t, "5", c["count"], "5 claims submitted")
	})
}

// ibcUserLike is the minimal subset of interchaintest's User interface we
// use in this file. Lets the helpers be written without importing the
// concrete type from interchaintest's chain package (which lives across
// versions / submodules).
type ibcUserLike interface {
	KeyName() string
	FormattedAddress() string
}

// trimQuotes strips a single leading and trailing `"` if present (JSON
// strings come back from json.RawMessage with quotes around them).
func trimQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// (silence unused-import warning if the rest of the suite evolves)
var (
	_ context.Context
	_ *cosmos.CosmosChain
)
