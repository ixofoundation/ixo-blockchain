package types_test

import (
	"strings"
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/x/names/types"
)

func mkAddr() string {
	return sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address()).String()
}

func validNamespace() *types.Namespace {
	return &types.Namespace{
		Name:              "handles",
		Description:       "test",
		AllowSelfRegister: true,
		MinLength:         3,
		MaxLength:         32,
	}
}

// Test all ValidateBasic implementations against a matrix of common failure
// modes plus a happy path. Catches regressions where a new field is added to
// a Msg* struct without a corresponding ValidateBasic check.
func TestMsgCreateNamespace_ValidateBasic(t *testing.T) {
	auth := mkAddr()
	for _, tc := range []struct {
		name string
		msg  types.MsgCreateNamespace
		err  string
	}{
		{"happy", types.MsgCreateNamespace{Authority: auth, Namespace: validNamespace()}, ""},
		{"bad authority", types.MsgCreateNamespace{Authority: "not-an-address", Namespace: validNamespace()}, "invalid authority"},
		{"nil namespace", types.MsgCreateNamespace{Authority: auth, Namespace: nil}, "namespace is required"},
		{"bad registrar addr", types.MsgCreateNamespace{Authority: auth, Namespace: &types.Namespace{
			Name: "verified", Description: "x", MinLength: 3, MaxLength: 32,
			AllowSelfRegister: false, RegistrarAccounts: []string{"not-bech32"},
		}}, "invalid address"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.err == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.err)
		})
	}
}

func TestMsgUpdateNamespace_ValidateBasic(t *testing.T) {
	auth := mkAddr()
	require.NoError(t, (&types.MsgUpdateNamespace{Authority: auth, Namespace: validNamespace()}).ValidateBasic())
	require.ErrorContains(t,
		(&types.MsgUpdateNamespace{Authority: "bad", Namespace: validNamespace()}).ValidateBasic(),
		"invalid authority",
	)
	require.ErrorContains(t,
		(&types.MsgUpdateNamespace{Authority: auth, Namespace: nil}).ValidateBasic(),
		"namespace is required",
	)
}

func TestMsgRegisterName_ValidateBasic(t *testing.T) {
	signer := mkAddr()
	for _, tc := range []struct {
		name string
		msg  types.MsgRegisterName
		err  string
	}{
		{"happy", types.MsgRegisterName{Signer: signer, Namespace: "n", Name: "alice", OwnerDid: "did:ixo:abc"}, ""},
		{"bad signer", types.MsgRegisterName{Signer: "bad", Namespace: "n", Name: "alice", OwnerDid: "did:ixo:abc"}, "invalid address"},
		{"empty namespace", types.MsgRegisterName{Signer: signer, Namespace: "", Name: "alice", OwnerDid: "did:ixo:abc"}, "namespace is required"},
		{"empty name", types.MsgRegisterName{Signer: signer, Namespace: "n", Name: "", OwnerDid: "did:ixo:abc"}, "name is required"},
		{"bad DID", types.MsgRegisterName{Signer: signer, Namespace: "n", Name: "alice", OwnerDid: "not-a-did"}, "invalid DID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.err == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.err)
		})
	}
}

func TestMsgRegisterNameByRegistrar_ValidateBasic(t *testing.T) {
	r := mkAddr()
	t.Run("happy", func(t *testing.T) {
		require.NoError(t, (&types.MsgRegisterNameByRegistrar{
			Registrar: r, Namespace: "n", Name: "alice", OwnerDid: "did:ixo:abc",
		}).ValidateBasic())
	})
	t.Run("oversized evidence hash", func(t *testing.T) {
		require.ErrorContains(t,
			(&types.MsgRegisterNameByRegistrar{
				Registrar: r, Namespace: "n", Name: "alice", OwnerDid: "did:ixo:abc",
				EvidenceHash: strings.Repeat("a", types.MaxNameRecordEvidenceHashLength+1),
			}).ValidateBasic(),
			"evidence_hash longer than",
		)
	})
}

func TestMsgUpdateNameByRegistrar_ValidateBasic(t *testing.T) {
	r := mkAddr()
	require.NoError(t, (&types.MsgUpdateNameByRegistrar{
		Registrar: r, Namespace: "n", NormalizedName: "alice",
	}).ValidateBasic())
	require.ErrorContains(t,
		(&types.MsgUpdateNameByRegistrar{Registrar: "bad", Namespace: "n", NormalizedName: "alice"}).ValidateBasic(),
		"invalid address",
	)
	require.ErrorContains(t,
		(&types.MsgUpdateNameByRegistrar{Registrar: r, Namespace: "", NormalizedName: "alice"}).ValidateBasic(),
		"namespace is required",
	)
}

func TestMsgTransferName_ValidateBasic(t *testing.T) {
	s := mkAddr()
	require.NoError(t, (&types.MsgTransferName{
		Signer: s, Namespace: "n", NormalizedName: "alice", NewOwnerDid: "did:ixo:b",
	}).ValidateBasic())
	require.ErrorContains(t,
		(&types.MsgTransferName{
			Signer: s, Namespace: "n", NormalizedName: "alice", NewOwnerDid: "not-a-did",
		}).ValidateBasic(),
		"invalid DID",
	)
}

func TestMsgSetNameStatus_ValidateBasic(t *testing.T) {
	s := mkAddr()
	require.NoError(t, (&types.MsgSetNameStatus{
		Signer: s, Namespace: "n", NormalizedName: "alice", Status: types.NAME_STATUS_ACTIVE,
	}).ValidateBasic())
	require.ErrorContains(t,
		(&types.MsgSetNameStatus{Signer: s, Namespace: "n", NormalizedName: "alice", Status: types.NameStatus(99)}).ValidateBasic(),
		"unsupported target status",
	)
}
