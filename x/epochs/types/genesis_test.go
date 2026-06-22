package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/x/epochs/types"
)

func TestDefaultGenesis_HasDailyHourlyWeekly(t *testing.T) {
	gs := types.DefaultGenesis()
	idents := map[string]bool{}
	for _, e := range gs.Epochs {
		idents[e.Identifier] = true
	}
	// Assert the required production epochs are present rather than pinning an
	// exact count: local/devnet test builds may also enable a "2min" epoch
	// (see DefaultGenesis), so the count varies but these must always exist.
	require.True(t, idents["day"])
	require.True(t, idents["hour"])
	require.True(t, idents["week"])
}

func TestGenesisState_Validate(t *testing.T) {
	cases := []struct {
		name string
		gs   types.GenesisState
		err  string
	}{
		{
			name: "happy",
			gs:   *types.DefaultGenesis(),
		},
		{
			name: "duplicate identifier",
			gs: types.GenesisState{Epochs: []types.EpochInfo{
				{Identifier: "x", Duration: time.Minute},
				{Identifier: "x", Duration: time.Hour},
			}},
			err: "should be unique",
		},
		{
			name: "internally invalid epoch",
			gs:   types.GenesisState{Epochs: []types.EpochInfo{{Identifier: "x"}}},
			err:  "duration should NOT be 0",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.gs.Validate()
			if tc.err == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.err)
		})
	}
}

func TestEpochInfo_Validate(t *testing.T) {
	cases := []struct {
		name string
		ep   types.EpochInfo
		err  string
	}{
		{"happy", types.EpochInfo{Identifier: "x", Duration: time.Minute}, ""},
		{"empty identifier", types.EpochInfo{Duration: time.Minute}, "identifier should NOT be empty"},
		{"zero duration", types.EpochInfo{Identifier: "x"}, "duration should NOT be 0"},
		{"negative current epoch", types.EpochInfo{Identifier: "x", Duration: time.Minute, CurrentEpoch: -1}, "non-negative"},
		{"negative start height", types.EpochInfo{Identifier: "x", Duration: time.Minute, CurrentEpochStartHeight: -1}, "non-negative"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.ep.Validate()
			if tc.err == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.err)
		})
	}
}
