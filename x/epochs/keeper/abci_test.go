package keeper_test

import (
	"time"

	"github.com/ixofoundation/ixo-blockchain/v6/x/epochs/types"
)

// TestBeginBlocker_StartsCountingOnFirstTriggerBlock verifies the
// shouldInitialEpochStart branch: the very first block whose time >= StartTime
// flips EpochCountingStarted to true and sets CurrentEpoch=1, even if the
// duration hasn't elapsed yet.
func (s *KeeperTestSuite) TestBeginBlocker_StartsCountingOnFirstTriggerBlock() {
	s.SetupTest()
	id := "abci-init"
	// Use a future StartTime so the first BeginBlocker call (with current
	// block time < StartTime) does NOT start the epoch.
	startTime := s.Ctx.BlockTime().Add(time.Hour)
	ep := types.EpochInfo{
		Identifier:              id,
		StartTime:               startTime,
		Duration:                time.Hour * 24,
		CurrentEpoch:            0,
		CurrentEpochStartHeight: s.Ctx.BlockHeight(),
		CurrentEpochStartTime:   time.Time{},
		EpochCountingStarted:    false,
	}
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, ep))

	// BeginBlocker before StartTime: must NOT start.
	s.App.EpochsKeeper.BeginBlocker(s.Ctx)
	got := s.App.EpochsKeeper.GetEpochInfo(s.Ctx, id)
	s.Require().False(got.EpochCountingStarted, "epoch must not start while ctx.BlockTime() < StartTime")

	// Advance block time past StartTime — counting must start.
	s.Ctx = s.Ctx.WithBlockTime(startTime.Add(time.Second))
	s.App.EpochsKeeper.BeginBlocker(s.Ctx)
	got = s.App.EpochsKeeper.GetEpochInfo(s.Ctx, id)
	s.Require().True(got.EpochCountingStarted, "first trigger block must start counting")
	s.Require().Equal(int64(1), got.CurrentEpoch)
	s.Require().Equal(startTime, got.CurrentEpochStartTime)
}

// TestBeginBlocker_RollsOverAfterDuration confirms that once Duration has
// elapsed past CurrentEpochStartTime, BeginBlocker increments CurrentEpoch.
func (s *KeeperTestSuite) TestBeginBlocker_RollsOverAfterDuration() {
	s.SetupTest()
	id := "abci-roll"
	dur := time.Minute
	start := s.Ctx.BlockTime()
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, types.EpochInfo{
		Identifier: id, StartTime: start, Duration: dur,
	}))

	// First BeginBlocker: starts counting at epoch 1.
	s.Ctx = s.Ctx.WithBlockTime(start)
	s.App.EpochsKeeper.BeginBlocker(s.Ctx)
	s.Require().Equal(int64(1), s.App.EpochsKeeper.GetEpochInfo(s.Ctx, id).CurrentEpoch)

	// Advance past Duration → roll to epoch 2.
	s.Ctx = s.Ctx.WithBlockTime(start.Add(dur).Add(time.Second))
	s.App.EpochsKeeper.BeginBlocker(s.Ctx)
	got := s.App.EpochsKeeper.GetEpochInfo(s.Ctx, id)
	s.Require().Equal(int64(2), got.CurrentEpoch)
	s.Require().Equal(start.Add(dur), got.CurrentEpochStartTime, "epoch start should advance by exactly one duration")
}

// TestBeginBlocker_NoOpWhenNotElapsed: between epochs, BeginBlocker should
// leave state unchanged.
func (s *KeeperTestSuite) TestBeginBlocker_NoOpWhenNotElapsed() {
	s.SetupTest()
	id := "abci-noop"
	dur := time.Hour
	start := s.Ctx.BlockTime()
	s.Require().NoError(s.App.EpochsKeeper.AddEpochInfo(s.Ctx, types.EpochInfo{
		Identifier: id, StartTime: start, Duration: dur,
	}))

	s.App.EpochsKeeper.BeginBlocker(s.Ctx) // start counting (epoch 1)
	before := s.App.EpochsKeeper.GetEpochInfo(s.Ctx, id)

	// Advance some time but NOT past Duration.
	s.Ctx = s.Ctx.WithBlockTime(start.Add(dur / 2))
	s.App.EpochsKeeper.BeginBlocker(s.Ctx)
	after := s.App.EpochsKeeper.GetEpochInfo(s.Ctx, id)
	s.Require().Equal(before.CurrentEpoch, after.CurrentEpoch)
	s.Require().Equal(before.CurrentEpochStartTime, after.CurrentEpochStartTime)
}
