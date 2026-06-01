//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestIxoClaimsDisputesMultiAdjudicator_FullScenario mirrors the SDK's
// claimsDisputesMultiAdjudicator: a collection with TWO whitelisted
// adjudicators (charlie 20%, oracle 50%), and two separate disputes that
// each get resolved by a different adjudicator. The chain reads the
// percentage from the AdjudicationDid entry matching the
// adjudicator_did supplied on MsgAdjudicateDispute — so charlie's
// dispute payouts use 20% and oracle's use 50%, even within the same
// collection.
//
// Load-bearing assertion: the actual_penalty_paid is the SAME for both
// (both slash the same fixed `penalty_amount_per_dispute`), but the
// 80/20-vs-50/50 split differs based on the adjudicator's
// reward_percentage. We measure the difference indirectly via the
// adjudicator's wallet delta (charlie gets +20% of penalty, oracle gets
// +50% of penalty).
func TestIxoClaimsDisputesMultiAdjudicator_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 6)
	creator, alice, tester, bob, charlie, oracle := users[0], users[1], users[2], users[3], users[4], users[5]

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
  "title": "register cw721",
  "summary": "Multi-adjudicator setup."
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
	})
	if t.Failed() {
		return
	}

	t.Run("fund admin + agents", func(t *testing.T) {
		for _, target := range []string{
			adminAddress,
			alice.FormattedAddress(), tester.FormattedAddress(),
			bob.FormattedAddress(), charlie.FormattedAddress(),
			oracle.FormattedAddress(),
		} {
			out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
				"bank", "send", creator.FormattedAddress(), target,
				"100000000uixo", "--gas", "auto", "--gas-adjustment", "1.5",
			)
			require.NoError(t, err, "fund: %s", out)
		}
	})
	if t.Failed() {
		return
	}

	t.Run("setup: register IIDs", func(t *testing.T) {
		for _, u := range []ibcUserLike{alice, tester, bob, charlie, oracle} {
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

	// Collection: TWO adjudicators with DIFFERENT reward percentages.
	var collectionID string
	t.Run("create-collection with two adjudicators (charlie 20%, oracle 50%)", func(t *testing.T) {
		startDate := time.Now().UTC().Format(time.RFC3339)
		endDate := time.Now().Add(72 * time.Hour).UTC().Format(time.RFC3339)
		charlieDid := "did:ixo:" + charlie.FormattedAddress()
		oracleDid := "did:ixo:" + oracle.FormattedAddress()
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
  "service_agent_deposit_required": [{"denom":"uixo","amount":"10000000"}],
  "evaluator_deposit_required":     [{"denom":"uixo","amount":"10000000"}],
  "dispute_deposit_amount":         [{"denom":"uixo","amount":"3000000"}],
  "penalty_amount_per_dispute":     [{"denom":"uixo","amount":"10000000"}],
  "min_deposit_period": 0,
  "adjudicators": [
    {"did": %q, "reward_percentage": "20.000000000000000000"},
    {"did": %q, "reward_percentage": "50.000000000000000000"}
  ]
}`,
			entityID, entityID, creator.FormattedAddress(),
			startDate, endDate,
			adminAddress, adminAddress, adminAddress, adminAddress,
			charlieDid, oracleDid)
		out, err := chain.GetNode().ExecTx(ctx, creator.KeyName(),
			"claims", "create-collection", collectionDoc,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "create-collection multi-adj: %s", out)

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
	})
	if t.Failed() || collectionID == "" {
		return
	}

	// Authz grants and deposit setup.
	grant := func(t *testing.T, grantee ibcUserLike, authTypeURL string) {
		t.Helper()
		var innerStripped string
		if authTypeURL == "/ixo.claims.v1beta1.SubmitClaimAuthorization" {
			innerStripped = fmt.Sprintf(`"admin":%q,"constraints":[{"collection_id":%q,"agent_quota":"100","max_amount":[],"max_cw20_payment":[],"max_cw1155_payment":[],"intent_duration_ns":"0s","member_address":""}]`,
				adminAddress, collectionID)
		} else {
			innerStripped = fmt.Sprintf(`"admin":%q,"constraints":[{"collection_id":%q,"agent_quota":"100","max_custom_amount":[],"max_custom_cw20_payment":[],"max_custom_cw1155_payment":[]}]`,
				adminAddress, collectionID)
		}
		grantBody := fmt.Sprintf(`"id":%q,"name":"admin","grantee_address":%q,"owner_address":%q,"grant":{"authorization":{"@type":%q,%s},"expiration":"2030-01-01T00:00:00Z"}`,
			entityID, grantee.FormattedAddress(), creator.FormattedAddress(),
			authTypeURL, innerStripped)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.entity.v1beta1.MsgGrantEntityAccountAuthz",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, grantBody)
		broadcastSignedTx(t, ctx, chain, creator.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)
	}
	t.Run("grant alice + tester authz", func(t *testing.T) {
		grant(t, alice, "/ixo.claims.v1beta1.SubmitClaimAuthorization")
		grant(t, tester, "/ixo.claims.v1beta1.EvaluateClaimAuthorization")
	})

	deposit := func(t *testing.T, agent ibcUserLike, amount string) {
		t.Helper()
		body := fmt.Sprintf(`"collection_id":%q,"agent_address":%q,"amount":[{"denom":"uixo","amount":%q}]`,
			collectionID, agent.FormattedAddress(), amount)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAddPerformanceDeposit",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTx(t, ctx, chain, agent.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)
	}
	t.Run("alice + tester top up performance deposits", func(t *testing.T) {
		deposit(t, alice, "10000000")
		deposit(t, tester, "20000000") // 20M so two slashes of 10M each can fire
	})

	// Helpers
	submitClaim := func(t *testing.T, agent ibcUserLike, claimID string) {
		t.Helper()
		inner := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgSubmitClaim","claim_id":%q,"collection_id":%q,"agent_did":"did:ixo:%s#key-1","agent_address":%q,"admin_address":%q}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"0","payer":"","granter":""}},"signatures":[]}`,
			claimID, collectionID, agent.FormattedAddress(), agent.FormattedAddress(), adminAddress)
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(inner), "submit.json"))
		_, err := chain.GetNode().ExecTx(ctx, agent.KeyName(),
			"authz", "exec", chain.GetNode().HomeDir()+"/submit.json",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err)
		WaitBlocks(t, ctx, chain, 2)
	}
	evalClaim := func(t *testing.T, agent ibcUserLike, claimID string, status int) {
		t.Helper()
		inner := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgEvaluateClaim","claim_id":%q,"collection_id":%q,"oracle":"did:ixo:%s","agent_did":"did:ixo:%s#key-1","agent_address":%q,"admin_address":%q,"status":%d,"reason":1,"verification_proof":"sha256:proof"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"0","payer":"","granter":""}},"signatures":[]}`,
			claimID, collectionID, agent.FormattedAddress(), agent.FormattedAddress(),
			agent.FormattedAddress(), adminAddress, status)
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(inner), "eval.json"))
		_, err := chain.GetNode().ExecTx(ctx, agent.KeyName(),
			"authz", "exec", chain.GetNode().HomeDir()+"/eval.json",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err)
		WaitBlocks(t, ctx, chain, 2)
	}
	dispute := func(t *testing.T, agent ibcUserLike, claimID, proof string, targetRole int) {
		t.Helper()
		body := fmt.Sprintf(`"agent_did":"did:ixo:%s#key-1","agent_address":%q,"subject_id":%q,"dispute_type":1,"target_role":%d,"data":{"type":"application/vnd.ixo+json","proof":%q,"uri":%q,"encrypted":false}`,
			agent.FormattedAddress(), agent.FormattedAddress(), claimID, targetRole,
			proof, "ipfs://"+proof)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgDisputeClaim",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTx(t, ctx, chain, agent.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)
	}
	adjudicate := func(t *testing.T, adj ibcUserLike, claimID string, targetRole int, dataProof string) {
		t.Helper()
		adjDid := "did:ixo:" + adj.FormattedAddress()
		body := fmt.Sprintf(`"subject_id":%q,"target_role":%d,"adjudicator_did":%q,"adjudicator_address":%q,"outcome":1,"data":{"type":"application/vnd.ixo+json","proof":%q,"uri":%q,"encrypted":false},"penalty_amount":[]`,
			claimID, targetRole, adjDid, adj.FormattedAddress(), dataProof, "ipfs://"+dataProof)
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAdjudicateDispute",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTx(t, ctx, chain, adj.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)
	}
	walletBalance := func(t *testing.T, addr string) int64 {
		t.Helper()
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"bank", "balance", addr, "uixo", "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Balance struct {
				Amount string `json:"amount"`
			} `json:"balance"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		v, _ := strconv.ParseInt(resp.Balance.Amount, 10, 64)
		return v
	}

	const claim1 = "MULTI-1"
	const claim2 = "MULTI-2"

	// =======================================================
	// First dispute: bob disputes claim 1 EVALUATOR, charlie (20%) resolves
	// =======================================================
	t.Run("submit+evaluate+dispute claim 1, charlie AWARDS @ 20%", func(t *testing.T) {
		submitClaim(t, alice, claim1)
		evalClaim(t, tester, claim1, 1 /* APPROVED */)
		dispute(t, bob, claim1, "p-multi-1", 2 /* EVALUATOR */)

		// charlie's wallet pre-adjudication
		charlieBefore := walletBalance(t, charlie.FormattedAddress())
		adjudicate(t, charlie, claim1, 2 /* EVALUATOR */, "adj-multi-1")
		charlieAfter := walletBalance(t, charlie.FormattedAddress())

		// charlie should have gained ~20% of 10_000_000 = 2_000_000 uixo
		// (minus tx fees for the adjudicate broadcast).
		delta := charlieAfter - charlieBefore
		require.Greater(t, delta, int64(1_500_000),
			"charlie's 20%% adjudicator share must land in wallet (delta ≥ ~1.5M after fees): %d", delta)
		require.Less(t, delta, int64(2_500_000),
			"charlie's 20%% share must NOT exceed ~2.5M: %d", delta)
	})

	// =======================================================
	// Second dispute: bob disputes claim 2 EVALUATOR, oracle (50%) resolves
	// Compares 50% payout vs 20% payout.
	// =======================================================
	t.Run("submit+evaluate+dispute claim 2, oracle AWARDS @ 50%", func(t *testing.T) {
		submitClaim(t, alice, claim2)
		evalClaim(t, tester, claim2, 1 /* APPROVED */)
		dispute(t, bob, claim2, "p-multi-2", 2 /* EVALUATOR */)

		oracleBefore := walletBalance(t, oracle.FormattedAddress())
		adjudicate(t, oracle, claim2, 2 /* EVALUATOR */, "adj-multi-2")
		oracleAfter := walletBalance(t, oracle.FormattedAddress())

		// oracle should have gained ~50% of 10_000_000 = 5_000_000 uixo.
		delta := oracleAfter - oracleBefore
		require.Greater(t, delta, int64(4_500_000),
			"oracle's 50%% adjudicator share must land in wallet (delta ≥ ~4.5M after fees): %d", delta)
		require.Less(t, delta, int64(5_500_000),
			"oracle's 50%% share must NOT exceed ~5.5M: %d", delta)
	})

	// =======================================================
	// Sanity: the chain accepted BOTH per-adjudicator percentages on the
	// SAME collection. If multi-adjudicator weren't wired, only the first
	// entry's percentage would apply.
	// =======================================================
	t.Run("final: both AWARDED disputes counted", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Collection map[string]json.RawMessage `json:"collection"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.Equal(t, "2", trimQuotes(string(resp.Collection["disputes_awarded"])),
			"both disputes resolved AWARDED → counter = 2")
	})
}
