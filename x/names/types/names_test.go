package types_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

func TestNormalizeName(t *testing.T) {
	for _, tc := range []struct {
		in, want string
	}{
		{"Alice", "alice"},
		{"  bob ", "bob"},
		{"BOB", "bob"},
		{"NoChange", "nochange"},
		{"", ""},
	} {
		require.Equalf(t, tc.want, types.NormalizeName(tc.in), "NormalizeName(%q)", tc.in)
	}
}

func TestValidateNamespace(t *testing.T) {
	base := types.Namespace{Name: "handles", AllowSelfRegister: true, MinLength: 3, MaxLength: 32}
	cases := []struct {
		name   string
		mut    func(ns *types.Namespace)
		errSub string
	}{
		{"happy", func(*types.Namespace) {}, ""},
		{"empty name", func(ns *types.Namespace) { ns.Name = "" }, "name is required"},
		{"name too long", func(ns *types.Namespace) { ns.Name = strings.Repeat("a", types.MaxNamespaceNameLength+1) }, "longer than"},
		{"invalid charset in name", func(ns *types.Namespace) { ns.Name = "Capital" }, "lowercase ASCII"},
		{"description too long", func(ns *types.Namespace) { ns.Description = strings.Repeat("d", types.MaxNamespaceDescriptionLength+1) }, "description longer than"},
		{"max_length=0", func(ns *types.Namespace) { ns.MaxLength = 0 }, "max_length must be > 0"},
		{"max_length above cap", func(ns *types.Namespace) { ns.MaxLength = types.MaxNameLengthCap + 1 }, "exceeds chain cap"},
		{"min > max", func(ns *types.Namespace) { ns.MinLength = 100; ns.MaxLength = 10 }, "must be <= max_length"},
		{"regex too long", func(ns *types.Namespace) { ns.Regex = strings.Repeat("a", types.MaxNamespaceRegexLength+1) }, "regex longer than"},
		{"bad regex", func(ns *types.Namespace) { ns.Regex = "([invalid" }, "regex did not compile"},
		{"no auth path", func(ns *types.Namespace) { ns.AllowSelfRegister = false; ns.RegistrarAccounts = nil }, "must allow self-registration"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ns := base
			tc.mut(&ns)
			err := types.ValidateNamespace(ns)
			if tc.errSub == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.errSub)
		})
	}
}

func TestValidateNameAgainstNamespace(t *testing.T) {
	ns := types.Namespace{Name: "handles", AllowSelfRegister: true, MinLength: 3, MaxLength: 8}
	require.NoError(t, types.ValidateNameAgainstNamespace(ns, "alice"))
	require.ErrorContains(t, types.ValidateNameAgainstNamespace(ns, "ab"), "shorter than namespace min_length")
	require.ErrorContains(t, types.ValidateNameAgainstNamespace(ns, "thisistoolong"), "longer than namespace max_length")
	require.ErrorContains(t, types.ValidateNameAgainstNamespace(ns, "Alice"), "lowercase ASCII")
	require.ErrorContains(t, types.ValidateNameAgainstNamespace(ns, "al!ce"), "lowercase ASCII")
	// regex narrows further
	ns.Regex = `^a.*`
	require.NoError(t, types.ValidateNameAgainstNamespace(ns, "alice"))
	require.ErrorContains(t, types.ValidateNameAgainstNamespace(ns, "bob"), "does not match")
}

func TestValidateRecordMetadata(t *testing.T) {
	require.NoError(t, types.ValidateRecordMetadata("", ""))
	require.NoError(t, types.ValidateRecordMetadata("sha256:abc", "linkedin"))
	require.ErrorContains(t,
		types.ValidateRecordMetadata(strings.Repeat("a", types.MaxNameRecordEvidenceHashLength+1), ""),
		"evidence_hash longer than",
	)
	require.ErrorContains(t,
		types.ValidateRecordMetadata("", strings.Repeat("s", types.MaxNameRecordSourceLength+1)),
		"source longer than",
	)
}

func TestHasRegistrar(t *testing.T) {
	ns := types.Namespace{RegistrarAccounts: []string{"a", "b"}}
	require.True(t, types.HasRegistrar(ns, "a"))
	require.True(t, types.HasRegistrar(ns, "b"))
	require.False(t, types.HasRegistrar(ns, "c"))
	require.False(t, types.HasRegistrar(types.Namespace{}, "a"))
}
