package keeper_test

import (
	"time"

	"github.com/ixofoundation/ixo-blockchain/v7/x/epochs/types"
)

// TestGenesis_RoundTrip: import a known set of epochs, export, ensure both
// the count and the per-epoch fields survive.
func (s *KeeperTestSuite) TestGenesis_RoundTrip() {
	s.SetupTest()

	// Wipe default genesis epochs to keep this test focused on round-trip.
	for _, e := range s.App.EpochsKeeper.AllEpochInfos(s.Ctx) {
		s.App.EpochsKeeper.DeleteEpochInfo(s.Ctx, e.Identifier)
	}

	gs := types.GenesisState{Epochs: []types.EpochInfo{
		{Identifier: "minute", Duration: time.Minute},
		{Identifier: "hour", Duration: time.Hour},
	}}

	s.App.EpochsKeeper.InitGenesis(s.Ctx, gs)
	out := s.App.EpochsKeeper.ExportGenesis(s.Ctx)
	s.Require().Len(out.Epochs, 2)

	got := map[string]types.EpochInfo{}
	for _, e := range out.Epochs {
		got[e.Identifier] = e
	}
	s.Require().Equal(time.Minute, got["minute"].Duration)
	s.Require().Equal(time.Hour, got["hour"].Duration)
}
