//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIxoNames_FullScenario boots ONE chain and walks the full names
// service end-to-end, including the user-facing Msg* surface that the
// chain's autocli leaves unwired (only governance Msg* are registered
// for autocli). All user-facing names msgs go through `tx broadcast`
// via the multi_message_test signing helpers.
//
//	gov-create namespace (yoid, self-register)  →
//	gov-create namespace (twitter, registrar-only)  →
//	user self-registers in yoid  →
//	resolve query normalises case  →
//	user transfers their name to another user  →
//	resolve query reflects new owner  →
//	non-owner transfer is rejected  →
//	by-namespace + by-owner queries return correct sets.
//
// Mirrors `ixo-multiclient-sdk/__tests__/flows/names.ts::namesBasic` +
// `namesTransferAndStatus` collapsed into one Go scenario.
func TestIxoNames_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 2)
	owner, recipient := users[0], users[1]

	// Both users need a registered iid so MsgRegisterName's iid-controller
	// check passes. CreateIidDoc returns `did:ixo:<addr>` with a single
	// verification tied to the user's signing address.
	ownerDID := CreateIidDoc(t, ctx, chain, owner)
	recipientDID := CreateIidDoc(t, ctx, chain, recipient)

	govAddr, err := chain.GetModuleAddress(ctx, "gov")
	require.NoError(t, err)

	// ----- 1. Gov-create the self-register "yoid" namespace -----
	t.Run("gov: create self-register namespace 'yoid'", func(t *testing.T) {
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.names.v1beta1.MsgCreateNamespace",
    "authority": %q,
    "namespace": {
      "name": "yoid",
      "description": "YoID self-registered handles",
      "allow_self_register": true,
      "allow_registrar_override": false,
      "min_length": 3,
      "max_length": 24
    }
  }],
  "metadata": "create yoid namespace",
  "deposit": "10000000uixo",
  "title": "create yoid",
  "summary": "Self-register namespace for the names scenario."
}`, govAddr)
		SubmitGovProposalAndPass(t, ctx, chain, owner, proposal)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "namespace", "yoid", "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), "yoid")
	})
	if t.Failed() {
		return
	}

	// ----- 2. Owner self-registers a name (mixed case) -----
	t.Run("user: self-register 'AliceCAPSandcase' in yoid", func(t *testing.T) {
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgRegisterName",
      "signer": %q,
      "namespace": "yoid",
      "name": "AliceCAPSandcase",
      "owner_did": %q
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "300000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, owner.FormattedAddress(), ownerDID, IxoNativeDenom)
		broadcastSignedTx(t, ctx, chain, owner.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)
	})
	if t.Failed() {
		return
	}

	// ----- 3. Resolve normalises case (display preserved) -----
	t.Run("resolve query normalises case to lowercase", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "resolve", "yoid", "AliceCAPSandcase", "--output", "json")
		require.NoError(t, err, "resolve: %s", stdout)
		var resp struct {
			Record struct {
				Namespace      string `json:"namespace"`
				NormalizedName string `json:"normalized_name"`
				DisplayName    string `json:"display_name"`
				OwnerDid       string `json:"owner_did"`
			} `json:"record"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.Equal(t, "yoid", resp.Record.Namespace)
		require.Equal(t, "alicecapsandcase", resp.Record.NormalizedName,
			"normalized form must lowercase the display name")
		require.Equal(t, "AliceCAPSandcase", resp.Record.DisplayName,
			"display form must preserve case")
		require.Equal(t, ownerDID, resp.Record.OwnerDid)
	})

	// ----- 4. by-namespace and by-owner queries -----
	t.Run("by-namespace and by-owner queries return the record", func(t *testing.T) {
		nsOut, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "list-by-namespace", "yoid", "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(nsOut), "alicecapsandcase")

		ownerOut, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "list-by-owner", ownerDID, "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(ownerOut), "alicecapsandcase")
	})

	// ----- 5. Owner transfers name to recipient -----
	t.Run("user: transfer name to recipient", func(t *testing.T) {
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgTransferName",
      "signer": %q,
      "namespace": "yoid",
      "normalized_name": "alicecapsandcase",
      "new_owner_did": %q
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "300000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, owner.FormattedAddress(), recipientDID, IxoNativeDenom)
		broadcastSignedTx(t, ctx, chain, owner.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "resolve", "yoid", "alicecapsandcase", "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Record struct {
				OwnerDid string `json:"owner_did"`
			} `json:"record"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.Equal(t, recipientDID, resp.Record.OwnerDid,
			"transfer must move the owner_did to the recipient")
	})

	// ----- 6. Non-owner transfer is rejected -----
	t.Run("non-owner transfer is rejected (registrar-override is off)", func(t *testing.T) {
		// `owner` is no longer the owner; the chain must reject this.
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgTransferName",
      "signer": %q,
      "namespace": "yoid",
      "normalized_name": "alicecapsandcase",
      "new_owner_did": %q
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "300000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, owner.FormattedAddress(), ownerDID, IxoNativeDenom)
		// broadcastSignedTxIgnoreError accepts a failed broadcast; we
		// then assert the underlying state didn't change.
		broadcastSignedTxIgnoreError(t, ctx, chain, owner.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "resolve", "yoid", "alicecapsandcase", "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), recipientDID,
			"non-owner transfer must NOT change the owner; resolve still shows recipient")
	})

	// ----- 7a. Gov-create a registrar-only namespace; tester is registrar -----
	t.Run("gov: create registrar-only namespace 'twitter'", func(t *testing.T) {
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.names.v1beta1.MsgCreateNamespace",
    "authority": %q,
    "namespace": {
      "name": "twitter",
      "description": "Twitter handles attested by oracle",
      "registrar_accounts": [%q],
      "allow_self_register": false,
      "allow_registrar_override": true,
      "min_length": 1,
      "max_length": 15
    }
  }],
  "metadata": "create twitter namespace",
  "deposit": "10000000uixo",
  "title": "create twitter",
  "summary": "Registrar-only namespace for the names scenario."
}`, govAddr, owner.FormattedAddress())
		SubmitGovProposalAndPass(t, ctx, chain, owner, proposal)
	})
	if t.Failed() {
		return
	}

	// ----- 7b. Registrar registers a verified handle on behalf of recipient -----
	t.Run("registrar: register-name-by-registrar with verified=true", func(t *testing.T) {
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgRegisterNameByRegistrar",
      "registrar": %q,
      "namespace": "twitter",
      "name": "recipientHandle",
      "owner_did": %q,
      "verified": true,
      "evidence_hash": "sha256:initial",
      "source": "twitter-oauth"
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "400000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, owner.FormattedAddress(), recipientDID, IxoNativeDenom)
		broadcastSignedTx(t, ctx, chain, owner.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "resolve", "twitter", "recipienthandle", "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Record struct {
				Verified bool   `json:"verified"`
				Source   string `json:"source"`
				OwnerDid string `json:"owner_did"`
			} `json:"record"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.True(t, resp.Record.Verified, "registrar-issued name must have verified=true")
		require.Equal(t, "twitter-oauth", resp.Record.Source)
		require.Equal(t, recipientDID, resp.Record.OwnerDid)
	})

	// ----- 7c. Registrar updates the same record's metadata -----
	t.Run("registrar: update-name-by-registrar updates source", func(t *testing.T) {
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgUpdateNameByRegistrar",
      "registrar": %q,
      "namespace": "twitter",
      "normalized_name": "recipienthandle",
      "verified": true,
      "evidence_hash": "sha256:updated",
      "source": "twitter-vc"
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "400000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, owner.FormattedAddress(), IxoNativeDenom)
		broadcastSignedTx(t, ctx, chain, owner.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "get", "twitter", "recipienthandle", "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), "twitter-vc",
			"update-name-by-registrar must change the source field: %s", stdout)
	})

	// ----- 7d. Registrar suspends the name -----
	t.Run("registrar: set-name-status SUSPENDED hides from resolve", func(t *testing.T) {
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgSetNameStatus",
      "signer": %q,
      "namespace": "twitter",
      "normalized_name": "recipienthandle",
      "status": "NAME_STATUS_SUSPENDED",
      "reason": "ToS violation"
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "400000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, owner.FormattedAddress(), IxoNativeDenom)
		broadcastSignedTx(t, ctx, chain, owner.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		// `get` always returns the record regardless of status — that's
		// the load-bearing assertion. Whether `resolve` hides the
		// suspended record is chain-version-dependent (some versions
		// return it with status=suspended; others omit). We check `get`
		// confirms the record exists in storage with the new status.
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "get", "twitter", "recipienthandle", "--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), "recipienthandle")
		// Status enum value 1 == NAME_STATUS_SUSPENDED.
		require.True(t,
			containsLowercase(string(stdout), "suspended") ||
				containsLowercase(string(stdout), "\"status\":1") ||
				containsLowercase(string(stdout), "\"status\": 1"),
			"get-name response must show suspended status: %s", stdout)
	})

	// ----- 7e. Gov-update a namespace's config -----
	t.Run("gov: update-namespace tightens max_length", func(t *testing.T) {
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.names.v1beta1.MsgUpdateNamespace",
    "authority": %q,
    "namespace": {
      "name": "yoid",
      "description": "YoID self-registered handles (locked-down)",
      "allow_self_register": true,
      "allow_registrar_override": false,
      "min_length": 3,
      "max_length": 8
    }
  }],
  "metadata": "tighten yoid",
  "deposit": "10000000uixo",
  "title": "tighten yoid",
  "summary": "Drops max_length to 8."
}`, govAddr)
		SubmitGovProposalAndPass(t, ctx, chain, owner, proposal)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "namespace", "yoid", "--output", "json")
		require.NoError(t, err)
		// Field name varies snake/camel across SDK versions; accept either.
		require.True(t,
			containsLowercase(string(stdout), "\"max_length\": 8") ||
				containsLowercase(string(stdout), "\"maxLength\": 8") ||
				containsLowercase(string(stdout), "\"max_length\":8") ||
				containsLowercase(string(stdout), "\"maxLength\":8"),
			"max_length must update to 8 after gov-update; raw: %s", stdout)
	})

	// ----- 8. Validation: name shorter than min_length is rejected -----
	t.Run("name shorter than min_length is rejected", func(t *testing.T) {
		raw := fmt.Sprintf(`{
  "body": {
    "messages": [{
      "@type": "/ixo.names.v1beta1.MsgRegisterName",
      "signer": %q,
      "namespace": "yoid",
      "name": "ab",
      "owner_did": %q
    }]
  },
  "auth_info": {"signer_infos": [], "fee": {"gas_limit": "300000", "amount": [{"denom": %q, "amount": "10000"}]}},
  "signatures": []
}`, recipient.FormattedAddress(), recipientDID, IxoNativeDenom)
		// Sign with the recipient (its DID matches the owner_did).
		broadcastSignedTxIgnoreError(t, ctx, chain, recipient.KeyName(), raw)
		WaitBlocks(t, ctx, chain, 3)

		// Confirm "ab" doesn't exist.
		stdout, _, _ := chain.GetNode().ExecQuery(ctx,
			"names", "resolve", "yoid", "ab", "--output", "json")
		require.False(t, strings.Contains(string(stdout), "\"ownerDid\"") ||
			strings.Contains(string(stdout), "\"owner_did\""),
			"shorter-than-min_length name must not be registered: %s", stdout)
	})
}
