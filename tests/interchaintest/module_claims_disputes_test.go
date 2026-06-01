//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestIxoClaimsDisputes_FullScenario mirrors the SDK's
// `claimsDisputesBasic` flow (ixo-multiclient-sdk/__tests__/flows/claims.ts).
// Walks the v7 dispute lifecycle end-to-end:
//
//   - Collection created with inline dispute config:
//       SA / EA deposit:  5_000_000 uixo
//       Dispute deposit:  3_000_000 uixo
//       Penalty:          5_000_000 uixo
//       Adjudicator:      charlie's DID, reward = 20%
//   - Deposit gate: submit / evaluate blocked until the agent tops up
//   - Bob disputes (target=EVALUATOR), charlie AWARDS → tester slashed,
//     80% to bob (winner), 20% to charlie (adjudicator)
//   - Re-dispute against the same (subject, role) is permanently blocked
//   - Alice's separate claim disputed (target=SUBMITTER) and DISMISSED →
//     bob loses dispute deposit, 80% to alice (vindicated), 20% to charlie
//
// Auth: MsgSubmitClaim / MsgEvaluateClaim require entity-admin signing
// (their proto signer = admin_address). We wrap them in `tx authz exec`
// after granting alice/tester via MsgGrantEntityAccountAuthz from the
// entity NFT owner (creator). MsgDisputeClaim / MsgAdjudicateDispute use
// their own agent_did / adjudicator_did and sign directly.
func TestIxoClaimsDisputes_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 5)
	creator, alice, tester, bob, charlie := users[0], users[1], users[2], users[3], users[4]

	// ----- Gov-register cw721 with entity module -----
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
  "summary": "Disputes flow setup."
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
		require.NotEmpty(t, adminAddress)
		t.Logf("entity %s admin %s", entityID, adminAddress)
	})
	if t.Failed() {
		return
	}

	t.Run("fund admin + agents", func(t *testing.T) {
		for _, target := range []string{
			adminAddress,
			alice.FormattedAddress(),
			tester.FormattedAddress(),
			bob.FormattedAddress(),
			charlie.FormattedAddress(),
		} {
			out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
				"bank", "send", creator.FormattedAddress(), target,
				"100000000uixo", "--gas", "auto", "--gas-adjustment", "1.5",
			)
			require.NoError(t, err, "fund %s: %s", target, out)
		}
	})
	if t.Failed() {
		return
	}

	// Register IID docs for alice/tester/bob/charlie so the IID ante
	// can resolve msg.AgentDid / msg.AdjudicatorDid to a registered
	// document with the agent's address as a VM. Use did:ixo:<bech32>
	// (matches SDK convention).
	t.Run("setup: register IIDs", func(t *testing.T) {
		for _, u := range []ibcUserLike{alice, tester, bob, charlie} {
			did := "did:ixo:" + u.FormattedAddress()
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

	// ----- Create collection with v7 dispute config inline -----
	var collectionID string
	t.Run("create-collection with dispute config", func(t *testing.T) {
		startDate := time.Now().UTC().Format(time.RFC3339)
		endDate := time.Now().Add(72 * time.Hour).UTC().Format(time.RFC3339)
		charlieDid := "did:ixo:" + charlie.FormattedAddress()
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
  "intents": null,
  "service_agent_deposit_required": [{"denom":"uixo","amount":"5000000"}],
  "evaluator_deposit_required":     [{"denom":"uixo","amount":"5000000"}],
  "dispute_deposit_amount":         [{"denom":"uixo","amount":"3000000"}],
  "penalty_amount_per_dispute":     [{"denom":"uixo","amount":"5000000"}],
  "min_deposit_period": 0,
  "adjudicators": [{"did": %q, "reward_percentage": "20.000000000000000000"}]
}`,
			entityID, entityID, creator.FormattedAddress(),
			startDate, endDate,
			adminAddress, adminAddress, adminAddress, adminAddress,
			charlieDid)
		out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
			"claims", "create-collection", collectionDoc,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "create-collection with dispute config: %s", out)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "params", "--output", "json")
		require.NoError(t, err)
		var paramsResp struct {
			Params struct {
				CollectionSequence string `json:"collection_sequence"`
			} `json:"params"`
		}
		require.NoError(t, json.Unmarshal(stdout, &paramsResp))
		collectionID = trimLastDigit(paramsResp.Params.CollectionSequence)
		require.NotEmpty(t, collectionID)
		t.Logf("collection id: %s", collectionID)
	})
	if t.Failed() || collectionID == "" {
		return
	}

	// ----- Authz: alice submit, tester evaluate.
	// Wrapping via MsgGrantEntityAccountAuthz so the grant is routed
	// through the entity admin account by the chain keeper.
	grantClaimsAuthz := func(t *testing.T, grantee ibcUserLike, authTypeURL string) {
		t.Helper()
		var inner string
		if authTypeURL == "/ixo.claims.v1beta1.SubmitClaimAuthorization" {
			inner = fmt.Sprintf(`{"admin":%q,"constraints":[{"collection_id":%q,"agent_quota":"100","max_amount":[],"max_cw20_payment":[],"max_cw1155_payment":[],"intent_duration_ns":"0s","member_address":""}]}`,
				adminAddress, collectionID)
		} else {
			inner = fmt.Sprintf(`{"admin":%q,"constraints":[{"collection_id":%q,"agent_quota":"100","max_custom_amount":[],"max_custom_cw20_payment":[],"max_custom_cw1155_payment":[]}]}`,
				adminAddress, collectionID)
		}
		grantBody := fmt.Sprintf(`"id":%q,"name":"admin","grantee_address":%q,"owner_address":%q,"grant":{"authorization":{"@type":%q,%s[1:len(%s)-1]},"expiration":"2030-01-01T00:00:00Z"}`,
			entityID, grantee.FormattedAddress(), creator.FormattedAddress(),
			authTypeURL, inner, inner)
		// The above Sprintf inserts JSON object literal as a string — the
		// Sprintf trick of stripping outer braces with [1:len-1] doesn't
		// work in Sprintf. Build it manually instead.
		innerStripped := inner[1 : len(inner)-1] // strip outer { }
		grantBody = fmt.Sprintf(`"id":%q,"name":"admin","grantee_address":%q,"owner_address":%q,"grant":{"authorization":{"@type":%q,%s},"expiration":"2030-01-01T00:00:00Z"}`,
			entityID, grantee.FormattedAddress(), creator.FormattedAddress(),
			authTypeURL, innerStripped)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.entity.v1beta1.MsgGrantEntityAccountAuthz",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
			grantBody)
		broadcastSignedTx(t, ctx, chain, creator.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)
	}
	t.Run("grant alice SubmitClaimAuthorization", func(t *testing.T) {
		grantClaimsAuthz(t, alice, "/ixo.claims.v1beta1.SubmitClaimAuthorization")
	})
	t.Run("grant tester EvaluateClaimAuthorization", func(t *testing.T) {
		grantClaimsAuthz(t, tester, "/ixo.claims.v1beta1.EvaluateClaimAuthorization")
	})

	// ----- helpers for authz-exec wrapped submit / evaluate -----
	authzExec := func(t *testing.T, grantee ibcUserLike, innerTypeURL, innerBody string) (bool, string) {
		t.Helper()
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
		return err == nil, out
	}
	submitClaim := func(t *testing.T, agent ibcUserLike, claimID string) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"claim_id":%q,"collection_id":%q,"agent_did":"did:ixo:%s#key-1","agent_address":%q,"admin_address":%q`,
			claimID, collectionID, agent.FormattedAddress(),
			agent.FormattedAddress(), adminAddress)
		return authzExec(t, agent, "/ixo.claims.v1beta1.MsgSubmitClaim", body)
	}
	evaluateClaim := func(t *testing.T, agent ibcUserLike, claimID string, status int) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"claim_id":%q,"collection_id":%q,"oracle":"did:ixo:%s","agent_did":"did:ixo:%s#key-1","agent_address":%q,"admin_address":%q,"status":%d,"reason":1,"verification_proof":"sha256:proof"`,
			claimID, collectionID, agent.FormattedAddress(), agent.FormattedAddress(),
			agent.FormattedAddress(), adminAddress, status)
		return authzExec(t, agent, "/ixo.claims.v1beta1.MsgEvaluateClaim", body)
	}
	addPerformanceDeposit := func(t *testing.T, agent ibcUserLike, amount string) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"collection_id":%q,"agent_address":%q,"amount":[{"denom":"uixo","amount":%q}]`,
			collectionID, agent.FormattedAddress(), amount)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAddPerformanceDeposit",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
			body)
		signed, ok := trySignMultiMsgTx(t, ctx, chain, agent.KeyName(), raw)
		if !ok {
			return false, "sign failed"
		}
		out, _, err := chain.GetNode().Exec(ctx, []string{
			chain.Config().Bin,
			"tx", "broadcast", chain.GetNode().HomeDir() + "/" + signed,
			"--chain-id", chain.Config().ChainID,
			"--home", chain.GetNode().HomeDir(),
			"--node", "tcp://" + chain.GetNode().HostName() + ":26657",
			"-y", "--output", "json",
		}, nil)
		WaitBlocks(t, ctx, chain, 2)
		return err == nil, string(out)
	}
	withdrawPerformanceDeposit := func(t *testing.T, agent ibcUserLike) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"collection_id":%q,"agent_address":%q,"amount":[]`,
			collectionID, agent.FormattedAddress())
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgWithdrawPerformanceDeposit",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
			body)
		signed, ok := trySignMultiMsgTx(t, ctx, chain, agent.KeyName(), raw)
		if !ok {
			return false, "sign failed"
		}
		out, _, err := chain.GetNode().Exec(ctx, []string{
			chain.Config().Bin,
			"tx", "broadcast", chain.GetNode().HomeDir() + "/" + signed,
			"--chain-id", chain.Config().ChainID,
			"--home", chain.GetNode().HomeDir(),
			"--node", "tcp://" + chain.GetNode().HostName() + ":26657",
			"-y", "--output", "json",
		}, nil)
		WaitBlocks(t, ctx, chain, 2)
		return err == nil, string(out)
	}
	disputeClaim := func(t *testing.T, agent ibcUserLike, claimID, proof string, targetRole int) (bool, string) {
		t.Helper()
		body := fmt.Sprintf(`"agent_did":"did:ixo:%s#key-1","agent_address":%q,"subject_id":%q,"dispute_type":1,"target_role":%d,"data":{"type":"application/vnd.ixo+json","proof":%q,"uri":%q,"encrypted":false}`,
			agent.FormattedAddress(), agent.FormattedAddress(),
			claimID, targetRole, proof, "ipfs://"+proof)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgDisputeClaim",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
			body)
		signed, ok := trySignMultiMsgTx(t, ctx, chain, agent.KeyName(), raw)
		if !ok {
			return false, "sign failed"
		}
		out, _, err := chain.GetNode().Exec(ctx, []string{
			chain.Config().Bin,
			"tx", "broadcast", chain.GetNode().HomeDir() + "/" + signed,
			"--chain-id", chain.Config().ChainID,
			"--home", chain.GetNode().HomeDir(),
			"--node", "tcp://" + chain.GetNode().HostName() + ":26657",
			"-y", "--output", "json",
		}, nil)
		WaitBlocks(t, ctx, chain, 2)
		return err == nil, string(out)
	}
	adjudicateDispute := func(t *testing.T, claimID string, targetRole, outcome int, dataProof string) (bool, string) {
		t.Helper()
		charlieDid := "did:ixo:" + charlie.FormattedAddress()
		body := fmt.Sprintf(`"subject_id":%q,"target_role":%d,"adjudicator_did":%q,"adjudicator_address":%q,"outcome":%d,"data":{"type":"application/vnd.ixo+json","proof":%q,"uri":%q,"encrypted":false},"penalty_amount":[]`,
			claimID, targetRole, charlieDid, charlie.FormattedAddress(),
			outcome, dataProof, "ipfs://"+dataProof)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAdjudicateDispute",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
			body)
		signed, ok := trySignMultiMsgTx(t, ctx, chain, charlie.KeyName(), raw)
		if !ok {
			return false, "sign failed"
		}
		out, _, err := chain.GetNode().Exec(ctx, []string{
			chain.Config().Bin,
			"tx", "broadcast", chain.GetNode().HomeDir() + "/" + signed,
			"--chain-id", chain.Config().ChainID,
			"--home", chain.GetNode().HomeDir(),
			"--node", "tcp://" + chain.GetNode().HostName() + ":26657",
			"-y", "--output", "json",
		}, nil)
		WaitBlocks(t, ctx, chain, 3)
		t.Logf("adjudicate broadcast: err=%v out=%s", err, string(out))
		return err == nil, string(out)
	}
	queryAgentDeposit := func(t *testing.T, agent ibcUserLike) (string, bool) {
		t.Helper()
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "agent-deposit-balance", collectionID, agent.FormattedAddress(),
			"--output", "json")
		if err != nil {
			return "", false
		}
		var resp struct {
			Balance struct {
				Amount []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"amount"`
			} `json:"balance"`
		}
		if jerr := json.Unmarshal(stdout, &resp); jerr != nil {
			return "", false
		}
		if len(resp.Balance.Amount) == 0 {
			return "0", true
		}
		return resp.Balance.Amount[0].Amount, true
	}
	// queryCollectionField returns a Collection scalar field. Cosmos-SDK
	// proto-JSON encodes uint64 zeroes as omitted fields, so "" is
	// equivalent to "0" — the helper normalises that for assertion clarity.
	queryCollectionField := func(t *testing.T, field string) string {
		t.Helper()
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Collection map[string]json.RawMessage `json:"collection"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		v, ok := resp.Collection[field]
		if !ok {
			return "0"
		}
		s := trimQuotes(string(v))
		if s == "" {
			return "0"
		}
		return s
	}

	// =======================================================
	// 1. Deposit-gate negative: alice can't submit without deposit
	// =======================================================
	const claim1 = "DISP-1"
	t.Run("alice submit blocked: no performance deposit", func(t *testing.T) {
		ok, out := submitClaim(t, alice, claim1)
		require.False(t, ok, "submit must fail when deposit < required; got out=%s", out)
	})

	// =======================================================
	// 2. Alice tops up — submit gate passes
	// =======================================================
	t.Run("alice tops up performance deposit (5 IXO)", func(t *testing.T) {
		ok, out := addPerformanceDeposit(t, alice, "5000000")
		require.True(t, ok, "add-performance-deposit must succeed: %s", out)
		bal, found := queryAgentDeposit(t, alice)
		require.True(t, found, "alice deposit balance must exist after top-up")
		require.Equal(t, "5000000", bal, "alice deposit = 5_000_000 uixo")
	})

	t.Run("alice submits claim 1 (deposit gate passes)", func(t *testing.T) {
		ok, out := submitClaim(t, alice, claim1)
		require.True(t, ok, "submit must succeed after deposit gate satisfied: %s", out)
	})

	// =======================================================
	// 3. Tester deposit gate
	// =======================================================
	t.Run("tester evaluate blocked: no deposit", func(t *testing.T) {
		ok, out := evaluateClaim(t, tester, claim1, 1 /* APPROVED */)
		require.False(t, ok, "evaluate must fail before deposit; got out=%s", out)
	})
	t.Run("tester tops up evaluator deposit (5 IXO)", func(t *testing.T) {
		ok, out := addPerformanceDeposit(t, tester, "5000000")
		require.True(t, ok, "tester add-performance-deposit: %s", out)
	})
	t.Run("tester evaluates claim 1 APPROVED", func(t *testing.T) {
		ok, out := evaluateClaim(t, tester, claim1, 1)
		require.True(t, ok, "evaluate must succeed after deposit: %s", out)
	})

	// =======================================================
	// 4. Bob disputes claim 1 EVALUATOR (= tester)
	// =======================================================
	t.Run("bob disputes claim 1 against EVALUATOR", func(t *testing.T) {
		ok, out := disputeClaim(t, bob, claim1, "dispute-1-proof", 2 /* EVALUATOR */)
		require.True(t, ok, "dispute must land: %s", out)
		require.Equal(t, "1", queryCollectionField(t, "disputes_open"),
			"disputes_open bumped to 1 after dispute filing")
	})

	t.Run("tester withdraw blocked while disputed", func(t *testing.T) {
		// `tx broadcast` returns code=0 (mempool accepted) even when the
		// on-chain execution rejects. Verify by reading state: tester's
		// balance must remain at 5_000_000 (untouched).
		_, _ = withdrawPerformanceDeposit(t, tester)
		bal, found := queryAgentDeposit(t, tester)
		require.True(t, found, "tester balance must still exist after blocked withdraw")
		require.Equal(t, "5000000", bal,
			"withdraw must be a no-op while tester has an open dispute")
	})

	// =======================================================
	// 5. Charlie AWARDS dispute 1 (slash tester)
	// =======================================================
	t.Run("charlie adjudicates dispute 1 AWARDED", func(t *testing.T) {
		ok, out := adjudicateDispute(t, claim1, 2 /* EVALUATOR */, 1 /* AWARDED */, "adj-1-proof")
		require.True(t, ok, "adjudicate AWARDED must land: %s", out)

		// Post-AWARDED: tester's balance drained (5M penalty = 5M required = 0 remaining)
		bal, found := queryAgentDeposit(t, tester)
		if found {
			require.Equal(t, "0", bal, "tester balance must be drained to 0 after AWARDED slash")
		}

		require.Equal(t, "0", queryCollectionField(t, "disputes_open"),
			"disputes_open back to 0 after resolution")
		require.Equal(t, "1", queryCollectionField(t, "disputes_awarded"),
			"disputes_awarded incremented to 1")
	})

	// =======================================================
	// 6. Cannot re-dispute (subject, EVALUATOR) after AWARDED
	// =======================================================
	t.Run("re-dispute claim 1 EVALUATOR is blocked after AWARDED", func(t *testing.T) {
		// Broadcast may return code=0 (mempool acceptance) even when the
		// chain ultimately rejects with ErrDisputeAlreadyAwardedForSubjectRole.
		// Verify post-state: disputes_open must remain at 0.
		_, _ = disputeClaim(t, bob, claim1, "dispute-1-retry", 2)
		require.Equal(t, "0", queryCollectionField(t, "disputes_open"),
			"AWARDED permanently blocks new disputes on the same (subject, role)")
	})

	// =======================================================
	// 7. DISMISSED branch: dispute against alice as SUBMITTER
	// =======================================================
	const claim2 = "DISP-2"
	t.Run("tester re-deposits + alice submits claim 2 + APPROVED", func(t *testing.T) {
		// tester needs to top up again (drained by AWARDED).
		ok, _ := addPerformanceDeposit(t, tester, "5000000")
		require.True(t, ok)

		ok, out := submitClaim(t, alice, claim2)
		require.True(t, ok, "alice submit claim 2: %s", out)
		ok, out = evaluateClaim(t, tester, claim2, 1 /* APPROVED */)
		require.True(t, ok, "tester evaluate claim 2 APPROVED: %s", out)
	})

	t.Run("bob disputes claim 2 against SUBMITTER", func(t *testing.T) {
		ok, out := disputeClaim(t, bob, claim2, "dispute-2-proof", 1 /* SUBMITTER */)
		require.True(t, ok, "dispute SUBMITTER must land: %s", out)
		require.Equal(t, "1", queryCollectionField(t, "disputes_open"))
	})

	t.Run("charlie DISMISSES dispute 2", func(t *testing.T) {
		ok, out := adjudicateDispute(t, claim2, 1 /* SUBMITTER */, 2 /* DISMISSED */, "adj-2-proof")
		require.True(t, ok, "adjudicate DISMISSED must land: %s", out)
		require.Equal(t, "1", queryCollectionField(t, "disputes_dismissed"),
			"disputes_dismissed bumped to 1")
		// Alice's deposit balance should be untouched on DISMISSED.
		bal, found := queryAgentDeposit(t, alice)
		require.True(t, found, "alice balance must persist through DISMISSED")
		require.Equal(t, "5000000", bal, "alice deposit unchanged on DISMISSED (her stake survives)")
	})

	// Final counters check.
	t.Run("final: collection dispute counters", func(t *testing.T) {
		require.Equal(t, "0", queryCollectionField(t, "disputes_open"),
			"all disputes resolved")
		require.Equal(t, "1", queryCollectionField(t, "disputes_awarded"))
		require.Equal(t, "1", queryCollectionField(t, "disputes_dismissed"))
		require.Equal(t, "2", queryCollectionField(t, "count"), "2 claims submitted")
		require.Equal(t, "2", queryCollectionField(t, "evaluated"), "2 claims evaluated APPROVED")
		require.Equal(t, "2", queryCollectionField(t, "approved"))
	})
}
