package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

func TestDefaultGenesisState_IsValid(t *testing.T) {
	gs := types.DefaultGenesisState()
	require.NotNil(t, gs)
	require.Empty(t, gs.Namespaces)
	require.Empty(t, gs.Names)
	require.NoError(t, gs.Validate())
}

// TestGenesisState_Validate_Cases covers the validation rules as described in
// (gs GenesisState) Validate(). Each case targets one specific branch.
func TestGenesisState_Validate_Cases(t *testing.T) {
	mkNS := func(name string) types.Namespace {
		return types.Namespace{Name: name, AllowSelfRegister: true, MinLength: 3, MaxLength: 32}
	}
	mkRec := func(ns, n, did string) types.NameRecord {
		return types.NameRecord{Namespace: ns, NormalizedName: n, OwnerDid: did, Status: types.NAME_STATUS_ACTIVE}
	}

	cases := []struct {
		name   string
		gs     types.GenesisState
		errSub string
	}{
		{
			name: "happy: one namespace, one record",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles")},
				Names:      []types.NameRecord{mkRec("handles", "alice", "did:ixo:a")},
			},
		},
		{
			name: "duplicate namespace",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles"), mkNS("handles")},
			},
			errSub: "duplicate namespace",
		},
		{
			name: "internally invalid namespace",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{{Name: "", AllowSelfRegister: true, MaxLength: 32}},
			},
			errSub: "name is required",
		},
		{
			name: "record references unknown namespace",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles")},
				Names:      []types.NameRecord{mkRec("missing", "alice", "did:ixo:a")},
			},
			errSub: "references unknown namespace",
		},
		{
			name: "record empty normalized_name",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles")},
				Names:      []types.NameRecord{mkRec("handles", "", "did:ixo:a")},
			},
			errSub: "empty normalized_name",
		},
		{
			name: "record not in normalized form",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles")},
				Names:      []types.NameRecord{mkRec("handles", "Alice", "did:ixo:a")},
			},
			errSub: "is not in normalized form",
		},
		{
			name: "record empty owner_did",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles")},
				Names:      []types.NameRecord{mkRec("handles", "alice", "")},
			},
			errSub: "empty owner_did",
		},
		{
			name: "duplicate record in namespace",
			gs: types.GenesisState{
				Namespaces: []types.Namespace{mkNS("handles")},
				Names: []types.NameRecord{
					mkRec("handles", "alice", "did:ixo:a"),
					mkRec("handles", "alice", "did:ixo:b"),
				},
			},
			errSub: "duplicate name",
		},
		{
			name: "valid_until without AllowExpiry is rejected",
			gs: func() types.GenesisState {
				ns := mkNS("handles") // AllowExpiry = false
				rec := mkRec("handles", "alice", "did:ixo:a")
				rec.ValidUntil = 1234
				return types.GenesisState{
					Namespaces: []types.Namespace{ns},
					Names:      []types.NameRecord{rec},
				}
			}(),
			errSub: "namespace \"handles\" forbids expiry",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.gs.Validate()
			if tc.errSub == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.errSub)
		})
	}
}
