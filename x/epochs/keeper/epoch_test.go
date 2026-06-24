package keeper_test

import (
	"time"

	"github.com/ixofoundation/ixo-blockchain/v8/x/epochs/types"
)

func (s *KeeperTestSuite) TestAddEpochInfo_HappyPath() {
	s.SetupTest()
	ep := s.newEpoch("test", time.Minute)
	err := s.App.EpochsKeeper.AddEpochInfo(s.Ctx, ep)
	s.Require().NoError(err)

	got := s.App.EpochsKeeper.GetEpochInfo(s.Ctx, "test")
	s.Require().Equal("test", got.Identifier)
	s.Require().Equal(time.Minute, got.Duration)
	s.Require().Equal(s.Ctx.BlockTime(), got.StartTime,
		"AddEpochInfo should default StartTime to current block time")
	s.Require().Equal(s.Ctx.BlockHeight(), got.CurrentEpochStartHeight)
}

func (s *KeeperTestSuite) TestAddEpochInfo_DuplicateRejected() {
	s.SetupTest()
	ep := s.newEpoch("dup", time.Hour)
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, ep))
	err := s.App.EpochsKeeper.AddEpochInfo(s.Ctx, ep)
	s.Require().ErrorContains(err, "already exists")
}

func (s *KeeperTestSuite) TestAddEpochInfo_InvalidIsRejected() {
	s.SetupTest()
	cases := []struct {
		name string
		ep   types.EpochInfo
		err  string
	}{
		{"empty identifier", types.EpochInfo{Duration: time.Minute}, "identifier should NOT be empty"},
		{"zero duration", types.EpochInfo{Identifier: "x"}, "duration should NOT be 0"},
		{"negative current epoch", types.EpochInfo{Identifier: "x", Duration: time.Minute, CurrentEpoch: -1}, "non-negative"},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			err := s.App.EpochsKeeper.AddEpochInfo(s.Ctx, tc.ep)
			s.Require().ErrorContains(err, tc.err)
		})
	}
}

func (s *KeeperTestSuite) TestDeleteEpochInfo() {
	s.SetupTest()
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, s.newEpoch("doomed", time.Minute)))
	s.Require().Equal("doomed", s.App.EpochsKeeper.GetEpochInfo(s.Ctx, "doomed").Identifier)

	s.App.EpochsKeeper.DeleteEpochInfo(s.Ctx, "doomed")
	got := s.App.EpochsKeeper.GetEpochInfo(s.Ctx, "doomed")
	s.Require().Equal("", got.Identifier, "GetEpochInfo on missing identifier returns zero value")
}

func (s *KeeperTestSuite) TestAllEpochInfos_IncludesGenesisEpochs() {
	s.SetupTest()
	all := s.App.EpochsKeeper.AllEpochInfos(s.Ctx)
	// Genesis seeds day/hour/week. Order is by store key (lexicographic).
	idents := make(map[string]bool, len(all))
	for _, e := range all {
		idents[e.Identifier] = true
	}
	s.Require().True(idents["day"], "default genesis must seed `day`")
	s.Require().True(idents["hour"], "default genesis must seed `hour`")
	s.Require().True(idents["week"], "default genesis must seed `week`")
}

func (s *KeeperTestSuite) TestNumBlocksSinceEpochStart() {
	s.SetupTest()
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, s.newEpoch("blk", time.Minute)))

	// At creation, BlockHeight - CurrentEpochStartHeight = 0
	n, err := s.App.EpochsKeeper.NumBlocksSinceEpochStart(s.Ctx, "blk")
	s.Require().NoError(err)
	s.Require().Equal(int64(0), n)

	// Advance 5 blocks
	ctx5 := s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + 5)
	n, err = s.App.EpochsKeeper.NumBlocksSinceEpochStart(ctx5, "blk")
	s.Require().NoError(err)
	s.Require().Equal(int64(5), n)

	// Unknown identifier returns an error
	_, err = s.App.EpochsKeeper.NumBlocksSinceEpochStart(s.Ctx, "unknown")
	s.Require().ErrorContains(err, "not found")
}
