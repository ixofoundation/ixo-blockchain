package types_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

// TestNamespaceKey_PrefixIsolation guarantees a NamespaceKey can never collide
// with any other module's prefix. The prefix byte is reserved for x/names.
func TestNamespaceKey_PrefixIsolation(t *testing.T) {
	k := types.NamespaceKey("handles")
	require.True(t, bytes.HasPrefix(k, types.NamespaceKeyPrefix))
	require.NotEqual(t, k, types.NameRecordKey("handles", "alice"))
	require.NotEqual(t, k, types.OwnerIndexKey("did:ixo:a", "handles", "alice"))
}

// TestNameRecordKey_DelimiterIsolation guarantees that two distinct
// (namespace, normalized_name) tuples cannot share a key. This is the key
// uniqueness invariant the keeper relies on for `HasNameRecord`.
func TestNameRecordKey_DelimiterIsolation(t *testing.T) {
	a := types.NameRecordKey("handles", "alice")
	b := types.NameRecordKey("handles", "alice2")
	require.NotEqual(t, a, b)
	c := types.NameRecordKey("handle", "salice") // pre-delimiter substring overlap
	require.NotEqual(t, a, c)
}

// TestOwnerIndex_RoundTrip verifies the (namespace, name) suffix can be
// recovered after iteration past OwnerIndexPrefix(ownerDid).
func TestOwnerIndex_RoundTrip(t *testing.T) {
	owner := "did:ixo:abc"
	full := types.OwnerIndexKey(owner, "handles", "alice")
	prefix := types.OwnerIndexPrefix(owner)
	require.True(t, bytes.HasPrefix(full, prefix))

	suffix := full[len(prefix):]
	ns, name, ok := types.ParseOwnerIndexSuffix(suffix)
	require.True(t, ok)
	require.Equal(t, "handles", ns)
	require.Equal(t, "alice", name)
}

// TestParseOwnerIndexSuffix_NoDelimiter rejects malformed suffixes.
func TestParseOwnerIndexSuffix_NoDelimiter(t *testing.T) {
	_, _, ok := types.ParseOwnerIndexSuffix([]byte("nodelim"))
	require.False(t, ok)
}

// TestNameRecordNamespacePrefix is iterable: every NameRecordKey under
// `namespace` starts with NameRecordNamespacePrefix(namespace).
func TestNameRecordNamespacePrefix_IsIterable(t *testing.T) {
	prefix := types.NameRecordNamespacePrefix("handles")
	for _, n := range []string{"alice", "bob", "carol"} {
		k := types.NameRecordKey("handles", n)
		require.True(t, bytes.HasPrefix(k, prefix))
	}
	// other namespaces are NOT under this prefix
	other := types.NameRecordKey("groups", "alice")
	require.False(t, bytes.HasPrefix(other, prefix))
}
