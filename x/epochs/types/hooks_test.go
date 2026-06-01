package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/x/epochs/types"
)

// MultiEpochHooks fan-out (BeforeEpochStart / AfterEpochEnd) requires a
// fully-initialised sdk.Context with a logger and a cache-able multistore,
// so the runtime dispatch is exercised in the keeper integration tests via
// BeginBlocker rather than at the types layer. Here we only check the parts
// that are pure value semantics.

func TestMultiEpochHooks_GetModuleName(t *testing.T) {
	multi := types.NewMultiEpochHooks()
	require.Equal(t, types.ModuleName, multi.GetModuleName(),
		"MultiEpochHooks.GetModuleName must always return the epochs module name")
}

func TestMultiEpochHooks_NewVariadic(t *testing.T) {
	require.Len(t, types.NewMultiEpochHooks(), 0)
	require.Len(t, types.NewMultiEpochHooks(types.MultiEpochHooks{}...), 0)
}
