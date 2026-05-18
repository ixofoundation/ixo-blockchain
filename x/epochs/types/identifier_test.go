package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v6/x/epochs/types"
)

func TestValidateEpochIdentifierString(t *testing.T) {
	require.NoError(t, types.ValidateEpochIdentifierString("day"))
	require.ErrorContains(t, types.ValidateEpochIdentifierString(""), "empty distribution epoch identifier")
}

func TestValidateEpochIdentifierInterface(t *testing.T) {
	require.NoError(t, types.ValidateEpochIdentifierInterface("day"))
	require.ErrorContains(t, types.ValidateEpochIdentifierInterface(123), "invalid parameter type")
	require.ErrorContains(t, types.ValidateEpochIdentifierInterface(""), "empty distribution epoch identifier")
}
