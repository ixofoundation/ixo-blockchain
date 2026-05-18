package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

func TestDefaultGenesisState(t *testing.T) {
	gs := types.DefaultGenesisState()
	require.NotNil(t, gs)
	require.Empty(t, gs.IidDocs)
}
