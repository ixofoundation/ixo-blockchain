package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/x/bonds/types"
)

func TestBondState_IsValidProgressionFrom(t *testing.T) {
	cases := []struct {
		from, to types.BondState
		ok       bool
	}{
		{types.HatchState, types.OpenState, true},
		{types.HatchState, types.FailedState, true},
		{types.HatchState, types.SettleState, false},
		{types.OpenState, types.SettleState, true},
		{types.OpenState, types.FailedState, true},
		{types.OpenState, types.HatchState, false},
		{types.SettleState, types.OpenState, false},
		{types.FailedState, types.OpenState, false},
	}
	for _, tc := range cases {
		got := tc.to.IsValidProgressionFrom(tc.from)
		require.Equalf(t, tc.ok, got, "IsValidProgressionFrom(%s -> %s)", tc.from, tc.to)
	}
}

func TestBondStateFromString_RoundTrip(t *testing.T) {
	for _, s := range []types.BondState{types.HatchState, types.OpenState, types.SettleState, types.FailedState} {
		require.Equal(t, s, types.BondStateFromString(s.String()))
	}
}
