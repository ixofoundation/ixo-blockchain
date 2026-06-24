package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/x/epochs/types"
)

// KeeperTestSuite covers x/epochs keeper, BeginBlocker, hooks, queries, and
// genesis behaviour. The chain seeds three default epochs (`day`, `hour`,
// `week`) at genesis — most tests deliberately use a fresh identifier to
// avoid coupling with that fixture.
type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryServer types.QueryServer
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) { suite.Run(t, new(KeeperTestSuite)) }

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.queryServer = *s.App.EpochsKeeper
	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func (s *KeeperTestSuite) goCtx() context.Context { return s.Ctx }

// newEpoch builds an EpochInfo ready to add to the keeper. Using
// time.Now().UTC() ensures BlockTime semantics match.
func (s *KeeperTestSuite) newEpoch(id string, dur time.Duration) types.EpochInfo {
	return types.EpochInfo{
		Identifier:              id,
		StartTime:               time.Time{}, // keeper fills with BlockTime
		Duration:                dur,
		CurrentEpoch:            0,
		CurrentEpochStartHeight: 0,
		CurrentEpochStartTime:   time.Time{},
		EpochCountingStarted:    false,
	}
}
