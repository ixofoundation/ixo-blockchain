package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

func TestIsValidDID(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"did:ixo:abc123", true},
		{"did:cosmos:foo", true},
		{"not-a-did", false},
		{"", false},
		{"did::abc", false},
		{"did:ixo:", false},
	}
	for _, tc := range cases {
		got := types.IsValidDID(tc.in)
		require.Equalf(t, tc.want, got, "IsValidDID(%q)", tc.in)
	}
}

func TestIsValidDIDURL(t *testing.T) {
	require.True(t, types.IsValidDIDURL("did:ixo:abc#key-1"))
	require.True(t, types.IsValidDIDURL("did:ixo:abc"))
	require.False(t, types.IsValidDIDURL("not-a-did-url"))
}

func TestBlockchainAccountID(t *testing.T) {
	bid := types.NewBlockchainAccountID("ixo1abc")
	require.Equal(t, "ixo1abc", string(bid))
	require.Equal(t, "ixo1abc", bid.GetAddress(), "no `:` in source means GetAddress returns the whole string")

	caip := types.NewBlockchainAccountID("cosmos:cosmoshub-4:cosmos1abc")
	require.Equal(t, "cosmos1abc", caip.GetAddress(), "GetAddress returns the suffix after the last `:`")
	require.True(t, caip.MatchAddress("cosmos1abc"))
	require.False(t, caip.MatchAddress("cosmos1xyz"))
}

// IsReservedDid recognises every prefix in ReservedDidPrefixes and lets
// ordinary DIDs through. Regression for ca41cb4d feat(iid): block
// module-reserved DID namespaces on MsgCreateIidDocument.
func TestIsReservedDid(t *testing.T) {
	// Every registered prefix is recognised.
	require.NotEmpty(t, types.ReservedDidPrefixes,
		"ReservedDidPrefixes must list at least one module prefix")
	for _, p := range types.ReservedDidPrefixes {
		gotPrefix, ok := types.IsReservedDid(p + "anything")
		require.True(t, ok, "IsReservedDid must accept %q", p+"anything")
		require.Equal(t, p, gotPrefix, "returned prefix must match the registered one")

		// The bare prefix itself counts as reserved — len(id) == len(p) is the boundary.
		gotPrefix, ok = types.IsReservedDid(p)
		require.True(t, ok, "IsReservedDid must accept bare prefix %q", p)
		require.Equal(t, p, gotPrefix)
	}

	// Non-reserved DIDs pass through.
	for _, free := range []string{
		"did:ixo:abc123",
		"did:ixo:not-entity-prefixed",
		"did:cosmos:foo",
		"",
	} {
		_, ok := types.IsReservedDid(free)
		require.False(t, ok, "IsReservedDid must NOT flag %q", free)
	}
}

// ValidateBasic rejects reserved-namespace DIDs at the antehandler. The
// handler is also defended in depth (see keeper test), but ValidateBasic
// is what runs on normal user-broadcast txs.
func TestMsgCreateIidDocument_ValidateBasic_ReservedNamespace(t *testing.T) {
	// Generate a valid bech32-encoded signer so ValidateBasic reaches the
	// reserved-namespace check (signer is validated first).
	signer := sdk.AccAddress("front-running-attacker").String()
	for _, prefix := range types.ReservedDidPrefixes {
		msg := types.MsgCreateIidDocument{
			Id:     prefix + "front-running-id",
			Signer: signer,
			Verifications: []*types.Verification{{
				Relationships: []string{"authentication"},
				Method: &types.VerificationMethod{
					Id:         prefix + "front-running-id#key-1",
					Controller: prefix + "front-running-id",
					Type:       "CosmosAccountAddress",
					VerificationMaterial: &types.VerificationMethod_BlockchainAccountID{
						BlockchainAccountID: signer,
					},
				},
			}},
		}
		err := msg.ValidateBasic()
		require.ErrorIs(t, err, types.ErrReservedDidNamespace,
			"ValidateBasic must reject prefix %q (got %v)", prefix, err)
	}
}

// IXO-2045: ValidateMsgCreateDIDForm is the policy the iid module enforces
// on every MsgCreateIidDocument — both in ValidateBasic and (defense in
// depth) inside the msg-server handler so Cosmwasm contracts cannot bypass
// it. This table-driven test pins each allow/deny path with the matching
// sentinel error so a future refactor that drops a branch fails loudly.
func TestValidateMsgCreateDIDForm(t *testing.T) {
	signer := sdk.AccAddress("signer-bytes-20-chars").String()
	other := sdk.AccAddress("another-acct-20-chars").String()
	contract := sdk.AccAddress("contract-addr-20-byte").String()

	cases := []struct {
		name    string
		id      string
		signer  string
		errLike error // nil => expected to pass
	}{
		{"did:ixo:<signer> ok", "did:ixo:" + signer, signer, nil},
		{"did:ixo:wasm:<contract> ok (signer != contract is fine)",
			"did:ixo:wasm:" + contract, signer, nil},

		{"did:ixo:<other> rejected (account != signer)",
			"did:ixo:" + other, signer, types.ErrDIDAccountSignerMismatch},
		{"did:ixo:wasm:<not-bech32> rejected",
			"did:ixo:wasm:not-a-bech32", signer, types.ErrDIDFormNotAllowed},
		{"did:ixo:wasm:<empty> rejected",
			"did:ixo:wasm:", signer, types.ErrDIDFormNotAllowed},
		{"did:ixo:wasm:<addr>:more rejected (nested sub-method)",
			"did:ixo:wasm:" + contract + ":more", signer, types.ErrDIDFormNotAllowed},
		{"did:ixo:haha:haha rejected (unknown sub-method)",
			"did:ixo:haha:haha", signer, types.ErrDIDFormNotAllowed},
		{"did:ixo:<not-bech32> rejected (account not bech32)",
			"did:ixo:not-a-bech32", signer, types.ErrDIDFormNotAllowed},
		{"did:ixo: (empty suffix) rejected",
			"did:ixo:", signer, types.ErrDIDFormNotAllowed},
		{"did:cosmos:foo rejected (wrong method)",
			"did:cosmos:foo", signer, types.ErrDIDFormNotAllowed},
		{"did:x:ixo:abc rejected (DidChainPrefix not a user namespace)",
			"did:x:ixo:abc", signer, types.ErrDIDFormNotAllowed},

		// Reserved namespace continues to be flagged distinctly.
		{"did:ixo:entity:* rejected (reserved)",
			"did:ixo:entity:abc", signer, types.ErrReservedDidNamespace},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateMsgCreateDIDForm(tc.id, tc.signer)
			if tc.errLike == nil {
				require.NoError(t, err, "id=%q signer=%q", tc.id, tc.signer)
				return
			}
			require.ErrorIs(t, err, tc.errLike,
				"id=%q signer=%q (got %v, expected %v)", tc.id, tc.signer, err, tc.errLike)
		})
	}
}
