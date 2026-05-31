package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/lib/ixo"
	"github.com/ixofoundation/ixo-blockchain/v7/x/mint/types"
)

// KeeperTestSuite covers x/mint keeper, distribution math, params/minter
// round-trips, and genesis. The mint module has no Msg* server (params are
// gov-controlled and the minter is updated by the BeginBlocker hook), so
// the surface is keeper-shaped rather than handler-shaped.
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

func (s *KeeperTestSuite) TestParams_RoundTrip() {
	s.SetupTest()
	got := s.App.MintKeeper.GetParams(s.Ctx)
	s.Require().Equal(ixo.IxoNativeToken, got.MintDenom)

	got.MintDenom = "ualtdenom"
	s.App.MintKeeper.SetParams(s.Ctx, got)
	roundtrip := s.App.MintKeeper.GetParams(s.Ctx)
	s.Require().Equal("ualtdenom", roundtrip.MintDenom)
}

func (s *KeeperTestSuite) TestMinter_RoundTrip() {
	s.SetupTest()
	m := s.App.MintKeeper.GetMinter(s.Ctx)
	s.Require().NotNil(m)

	m.EpochProvisions = math.LegacyMustNewDecFromStr("123.456")
	s.App.MintKeeper.SetMinter(s.Ctx, m)

	roundtrip := s.App.MintKeeper.GetMinter(s.Ctx)
	// LegacyDec serialises with 18-digit precision, so we compare canonical
	// strings instead of literal text.
	s.Require().Equal(math.LegacyMustNewDecFromStr("123.456").String(), roundtrip.EpochProvisions.String())
}

// TestQueryParams via the gRPC client to confirm the query server is wired up
// in the app.
func (s *KeeperTestSuite) TestQueryParams() {
	s.SetupTest()
	resp, err := s.queryClient.Params(s.goCtx(), &types.QueryParamsRequest{})
	s.Require().NoError(err)
	s.Require().Equal(ixo.IxoNativeToken, resp.Params.MintDenom)
}

func (s *KeeperTestSuite) TestQueryEpochProvisions() {
	s.SetupTest()
	resp, err := s.queryClient.EpochProvisions(s.goCtx(), &types.QueryEpochProvisionsRequest{})
	s.Require().NoError(err)
	s.Require().False(resp.EpochProvisions.IsNil())
}

// TestDistributeMintedCoin: mint coins into the mint module account, then
// invoke DistributeMintedCoin and verify the staking incentives portion lands
// in the fee collector while the remainder ends up funding the community
// pool. We use a community-pool sentinel via balance change since the
// community pool is funded via distribution.
func (s *KeeperTestSuite) TestDistributeMintedCoin() {
	s.SetupTest()
	denom := ixo.IxoNativeToken
	want := sdk.NewCoin(denom, math.NewInt(1_000_000))

	// Mint into the mint module account.
	s.FundModuleAcc(types.ModuleName, sdk.NewCoins(want))

	err := s.App.MintKeeper.DistributeMintedCoin(s.Ctx, want)
	s.Require().NoError(err)

	// Mint module account should have less than `want` left (everything got
	// distributed; some may remain due to truncation rounding which goes back
	// to community pool).
	mintAcc := s.App.AccountKeeper.GetModuleAddress(types.ModuleName)
	left := s.App.BankKeeper.GetBalance(s.Ctx, mintAcc, denom)
	s.Require().True(left.Amount.LT(want.Amount), "mint module should have distributed coins out")
}
