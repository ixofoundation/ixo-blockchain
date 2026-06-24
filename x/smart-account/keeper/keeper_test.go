package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/x/smart-account/types"
)

// KeeperTestSuite covers x/smart-account keeper-level operations: adding
// and removing authenticators on an account, listing them, the active-state
// switch, and the global authenticator-id counter. The full ante-handler
// chain (authentication during tx-flow, circuit-breaker behaviour, classic
// vs authenticator dispatch) is exercised in Phase 2 ante tests.
type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) { suite.Run(t, new(KeeperTestSuite)) }

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func (s *KeeperTestSuite) goCtx() context.Context { return s.Ctx }
