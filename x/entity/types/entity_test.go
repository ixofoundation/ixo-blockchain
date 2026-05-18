package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v6/x/entity/types"
)

func TestIsEmpty(t *testing.T) {
	require.True(t, types.IsEmpty(""))
	require.True(t, types.IsEmpty("   "))
	require.False(t, types.IsEmpty("x"))
}

func TestIsValidEntity(t *testing.T) {
	now := time.Now().UTC()
	good := &types.Entity{
		Id: "did:ixo:abc",
		Metadata: &types.EntityMetadata{
			VersionId: "v1",
			Created:   &now,
		},
	}
	require.True(t, types.IsValidEntity(good))
	require.False(t, types.IsValidEntity(nil))

	bad := *good
	bad.Metadata = &types.EntityMetadata{Created: &now}
	require.False(t, types.IsValidEntity(&bad), "empty version id is invalid")

	bad2 := *good
	bad2.Metadata = &types.EntityMetadata{VersionId: "v1"}
	require.False(t, types.IsValidEntity(&bad2), "nil Created is invalid")
}

func TestNewEntityMetadata(t *testing.T) {
	now := time.Now().UTC()
	m := types.NewEntityMetadata([]byte("seed"), now)
	require.NotEmpty(t, m.VersionId)
	require.NotNil(t, m.Created)
	require.Equal(t, now, *m.Created)
}

func TestGetModuleAccountAddress(t *testing.T) {
	a := types.GetModuleAccountAddress("did:ixo:abc", "primary")
	b := types.GetModuleAccountAddress("did:ixo:abc", "secondary")
	c := types.GetModuleAccountAddress("did:ixo:abd", "primary")
	require.NotEqual(t, a, b, "different fragments produce different addresses")
	require.NotEqual(t, a, c, "different ids produce different addresses")
}
