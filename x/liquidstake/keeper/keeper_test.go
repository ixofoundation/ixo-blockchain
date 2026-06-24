package keeper_test

import (
	"context"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

// KeeperTestSuite covers x/liquidstake at the keeper layer: pool CRUD,
// validator-set tracking, params round-trip, and genesis. The full liquid
// staking flow (LiquidStake / LiquidUnstake / rebalance / Burn) needs a real
// staking keeper with bonded validators driving epoch-aligned delegations
// and is exercised in tests/interchaintest/.
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

// seedPool inserts a Pool struct directly via SetPool, bypassing the
// CreatePool message-handler validation chain (which requires a registered
// proxy account, pool denom uniqueness checks, etc.).
func (s *KeeperTestSuite) seedPool(poolID, liquidDenom string) types.Pool {
	proxy := apptesting.RandomAccountAddress()
	p := types.Pool{
		PoolId:                poolID,
		LiquidBondDenom:       liquidDenom,
		ProxyAccountAddress:   proxy.String(),
		WhitelistAdminAddress: apptesting.RandomAccountAddress().String(),
		FeeAccountAddress:     apptesting.RandomAccountAddress().String(),
		UnstakeFeeRate:        sdkmath.LegacyZeroDec(),
		AutocompoundFeeRate:   sdkmath.LegacyZeroDec(),
	}
	s.App.LiquidStakeKeeper.SetPool(s.Ctx, p)
	return p
}
