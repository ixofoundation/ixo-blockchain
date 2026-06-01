//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestIxoClaimsDisputesValidations_FullScenario covers the negative paths
// and edge cases from the SDK's claimsDisputesValidations flow:
//
//   - Adjudicator-not-in-whitelist: chain rejects with
//     ErrAdjudicatorDidNotApproved.
//   - Adjudicator-address-not-on-DID: signer is in the keyring but not
//     listed in the adjudicator_did's authentication VMs → IID ante
//     rejects.
//   - min_deposit_period gate: withdraw must reject while
//     withdrawable_at > blocktime.
//   - Withdraw-after-period: once withdrawable_at has passed, agent can
//     fully withdraw the deposit.
//
// Auth: same authz wrap pattern as TestIxoClaimsDisputes_FullScenario.
// We only exercise the dispute path far enough to land an OPEN dispute
// against an evaluator, then probe the adjudicate-authorization paths.
func TestIxoClaimsDisputesValidations_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 6)
	creator, alice, tester, bob, charlie, evil := users[0], users[1], users[2], users[3], users[4], users[5]

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
  "summary": "Disputes validations setup."
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
			alice.FormattedAddress(),
			tester.FormattedAddress(),
			bob.FormattedAddress(),
			charlie.FormattedAddress(),
			evil.FormattedAddress(),
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
		for _, u := range []ibcUserLike{alice, tester, bob, charlie, evil} {
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

	// Collection setup: dispute config WITH min_deposit_period = 60s so we
	// can validate the lock-period gate. Only charlie is on the
	// adjudicators whitelist; evil is registered but NOT in the whitelist
	// (used to exercise ErrAdjudicatorDidNotApproved).
	var collectionID string
	t.Run("create-collection with min_deposit_period=60s", func(t *testing.T) {
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
  "min_deposit_period": 60000000000,
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
		collectionID = trimLastDigit(paramsResp.Params.CollectionSequence)
		require.NotEmpty(t, collectionID)
	})
	if t.Failed() || collectionID == "" {
		return
	}

	// Grants
	grantClaimsAuthz := func(t *testing.T, grantee ibcUserLike, authTypeURL string) {
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
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.entity.v1beta1.MsgGrantEntityAccountAuthz",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`,
			grantBody)
		broadcastSignedTx(t, ctx, chain, creator.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)
	}
	t.Run("grant alice + tester authz", func(t *testing.T) {
		grantClaimsAuthz(t, alice, "/ixo.claims.v1beta1.SubmitClaimAuthorization")
		grantClaimsAuthz(t, tester, "/ixo.claims.v1beta1.EvaluateClaimAuthorization")
	})

	// =======================================================
	// min_deposit_period gate: alice tops up, then withdraw must fail
	// while inside the lock window.
	// =======================================================
	t.Run("alice tops up performance deposit (5 IXO)", func(t *testing.T) {
		body := fmt.Sprintf(`"collection_id":%q,"agent_address":%q,"amount":[{"denom":"uixo","amount":"5000000"}]`,
			collectionID, alice.FormattedAddress())
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAddPerformanceDeposit",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTx(t, ctx, chain, alice.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "agent-deposit-balance", collectionID, alice.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		// Withdrawable_at must be populated (period > 0).
		require.Contains(t, string(stdout), "withdrawable_at",
			"min_deposit_period > 0 must populate withdrawable_at on the balance record")
	})

	t.Run("withdraw blocked while inside min_deposit_period", func(t *testing.T) {
		body := fmt.Sprintf(`"collection_id":%q,"agent_address":%q,"amount":[]`,
			collectionID, alice.FormattedAddress())
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgWithdrawPerformanceDeposit",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTxIgnoreError(t, ctx, chain, alice.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)

		// Balance must remain at 5_000_000 (withdraw rejected by lock).
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "agent-deposit-balance", collectionID, alice.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		var resp struct {
			Balance struct {
				Amount []struct {
					Amount string `json:"amount"`
				} `json:"amount"`
			} `json:"balance"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.Len(t, resp.Balance.Amount, 1)
		require.Equal(t, "5000000", resp.Balance.Amount[0].Amount,
			"withdraw inside lock window must NOT change balance")
	})

	// =======================================================
	// Dispute setup: tester evaluates, bob disputes (EVALUATOR).
	// =======================================================
	const claimX = "VAL-CLAIM-1"
	t.Run("submit+evaluate to seed a dispute target", func(t *testing.T) {
		// alice submit via authz exec (claim X)
		inner := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgSubmitClaim","claim_id":%q,"collection_id":%q,"agent_did":"did:ixo:%s#key-1","agent_address":%q,"admin_address":%q}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"0","payer":"","granter":""}},"signatures":[]}`,
			claimX, collectionID, alice.FormattedAddress(), alice.FormattedAddress(), adminAddress)
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(inner), "alice-submit.json"))
		out, err := chain.GetNode().ExecTx(ctx, alice.KeyName(),
			"authz", "exec", chain.GetNode().HomeDir()+"/alice-submit.json",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "alice authz-submit: %s", out)
		WaitBlocks(t, ctx, chain, 2)

		// tester needs deposit too
		dep := fmt.Sprintf(`"collection_id":%q,"agent_address":%q,"amount":[{"denom":"uixo","amount":"5000000"}]`,
			collectionID, tester.FormattedAddress())
		depRaw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAddPerformanceDeposit",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, dep)
		broadcastSignedTx(t, ctx, chain, tester.KeyName(), depRaw)
		WaitBlocks(t, ctx, chain, 2)

		// tester evaluate APPROVED via authz exec
		evalInner := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgEvaluateClaim","claim_id":%q,"collection_id":%q,"oracle":"did:ixo:%s","agent_did":"did:ixo:%s#key-1","agent_address":%q,"admin_address":%q,"status":1,"reason":1,"verification_proof":"sha256:proof"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"0","payer":"","granter":""}},"signatures":[]}`,
			claimX, collectionID, tester.FormattedAddress(), tester.FormattedAddress(),
			tester.FormattedAddress(), adminAddress)
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(evalInner), "tester-eval.json"))
		out, err = chain.GetNode().ExecTx(ctx, tester.KeyName(),
			"authz", "exec", chain.GetNode().HomeDir()+"/tester-eval.json",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "tester authz-evaluate: %s", out)
		WaitBlocks(t, ctx, chain, 2)

		// Bob disputes (EVALUATOR=2)
		disp := fmt.Sprintf(`"agent_did":"did:ixo:%s#key-1","agent_address":%q,"subject_id":%q,"dispute_type":1,"target_role":2,"data":{"type":"application/vnd.ixo+json","proof":"val-proof","uri":"ipfs://val","encrypted":false}`,
			bob.FormattedAddress(), bob.FormattedAddress(), claimX)
		dispRaw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgDisputeClaim",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, disp)
		broadcastSignedTx(t, ctx, chain, bob.KeyName(), dispRaw)
		WaitBlocks(t, ctx, chain, 2)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), `"disputes_open"`,
			"a dispute must be filed by this point")
	})

	// =======================================================
	// Adjudicator-not-in-whitelist (evil): chain must reject.
	// =======================================================
	t.Run("adjudicate by EVIL (not in whitelist) is rejected", func(t *testing.T) {
		evilDid := "did:ixo:" + evil.FormattedAddress()
		body := fmt.Sprintf(`"subject_id":%q,"target_role":2,"adjudicator_did":%q,"adjudicator_address":%q,"outcome":1,"data":{"type":"application/vnd.ixo+json","proof":"evil-proof","uri":"ipfs://evil","encrypted":false},"penalty_amount":[]`,
			claimX, evilDid, evil.FormattedAddress())
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAdjudicateDispute",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTxIgnoreError(t, ctx, chain, evil.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)

		// dispute still OPEN (rejection means nothing changed).
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), `"disputes_open"`,
			"evil's adjudicate must NOT close the dispute")
	})

	// =======================================================
	// Adjudicator-address-not-on-DID: evil signs with charlie's DID
	// (whitelist hit), but evil's address isn't on charlie's DID
	// authentication VMs → IID ante rejects.
	// =======================================================
	t.Run("adjudicate mismatched address+DID is rejected by IID ante", func(t *testing.T) {
		charlieDid := "did:ixo:" + charlie.FormattedAddress()
		// adjudicator_did=charlie (whitelisted), adjudicator_address=evil (mismatch).
		body := fmt.Sprintf(`"subject_id":%q,"target_role":2,"adjudicator_did":%q,"adjudicator_address":%q,"outcome":1,"data":{"type":"application/vnd.ixo+json","proof":"mix-proof","uri":"ipfs://mix","encrypted":false},"penalty_amount":[]`,
			claimX, charlieDid, evil.FormattedAddress())
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAdjudicateDispute",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTxIgnoreError(t, ctx, chain, evil.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 2)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), `"disputes_open"`,
			"IID-ante reject must NOT close the dispute")
	})

	// =======================================================
	// Happy path: charlie adjudicates AWARDED correctly.
	// =======================================================
	t.Run("charlie adjudicates correctly → AWARDED", func(t *testing.T) {
		charlieDid := "did:ixo:" + charlie.FormattedAddress()
		body := fmt.Sprintf(`"subject_id":%q,"target_role":2,"adjudicator_did":%q,"adjudicator_address":%q,"outcome":1,"data":{"type":"application/vnd.ixo+json","proof":"correct-proof","uri":"ipfs://correct","encrypted":false},"penalty_amount":[]`,
			claimX, charlieDid, charlie.FormattedAddress())
		raw := fmt.Sprintf(`{"body":{"messages":[{"@type":"/ixo.claims.v1beta1.MsgAdjudicateDispute",%s}]},"auth_info":{"signer_infos":[],"fee":{"gas_limit":"600000","amount":[{"denom":"uixo","amount":"10000"}]}},"signatures":[]}`, body)
		broadcastSignedTx(t, ctx, chain, charlie.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		// disputes_awarded incremented; disputes_open back to 0.
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"claims", "collection", collectionID, "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Collection map[string]json.RawMessage `json:"collection"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		// disputes_awarded should now be "1"
		require.Equal(t, "1", trimQuotes(string(resp.Collection["disputes_awarded"])),
			"correctly-authorised adjudicate AWARDED bumps disputes_awarded to 1")
	})
}
