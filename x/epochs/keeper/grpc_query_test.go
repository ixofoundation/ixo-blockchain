package keeper_test

import (
	"time"

	"github.com/ixofoundation/ixo-blockchain/v6/x/epochs/types"
)

func (s *KeeperTestSuite) TestQueryEpochInfos() {
	s.SetupTest()
	resp, err := s.queryClient.EpochInfos(s.goCtx(), &types.QueryEpochsInfoRequest{})
	s.Require().NoError(err)
	// Default genesis seeds three epochs.
	s.Require().GreaterOrEqual(len(resp.Epochs), 3)
}

func (s *KeeperTestSuite) TestQueryCurrentEpoch() {
	s.SetupTest()
	id := "tickled"
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, types.EpochInfo{
		Identifier: id, StartTime: s.Ctx.BlockTime(), Duration: time.Minute,
	}))

	// Tick BeginBlocker so CurrentEpoch jumps from 0 to 1.
	s.App.EpochsKeeper.BeginBlocker(s.Ctx)

	resp, err := s.queryClient.CurrentEpoch(s.goCtx(), &types.QueryCurrentEpochRequest{Identifier: id})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.CurrentEpoch)

	// Unknown identifier returns an error.
	_, err = s.queryClient.CurrentEpoch(s.goCtx(), &types.QueryCurrentEpochRequest{Identifier: "missing"})
	s.Require().ErrorContains(err, "not available identifier")

	// Empty identifier rejected with InvalidArgument.
	_, err = s.queryClient.CurrentEpoch(s.goCtx(), &types.QueryCurrentEpochRequest{Identifier: ""})
	s.Require().ErrorContains(err, "identifier is empty")

	// Direct keeper call — nil request guard.
	_, err = s.queryServer.CurrentEpoch(s.goCtx(), nil)
	s.Require().ErrorContains(err, "empty request")
}
