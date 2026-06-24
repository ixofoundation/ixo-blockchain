package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

func (s *KeeperTestSuite) TestPool_RoundTrip() {
	s.SetupTest()
	want := s.seedPool("zero", "uzero")
	got, found := s.App.LiquidStakeKeeper.GetPool(s.Ctx, "zero")
	s.Require().True(found)
	s.Require().Equal(want.LiquidBondDenom, got.LiquidBondDenom)
	s.Require().Equal(want.ProxyAccountAddress, got.ProxyAccountAddress)
}

func (s *KeeperTestSuite) TestHasPool() {
	s.SetupTest()
	s.seedPool("zero", "uzero")
	s.Require().True(s.App.LiquidStakeKeeper.HasPool(s.Ctx, "zero"))
	s.Require().False(s.App.LiquidStakeKeeper.HasPool(s.Ctx, "ghost"))
}

func (s *KeeperTestSuite) TestHasPoolWithDenom() {
	s.SetupTest()
	s.seedPool("zero", "uzero")
	s.seedPool("qi", "uqi")
	s.Require().True(s.App.LiquidStakeKeeper.HasPoolWithDenom(s.Ctx, "uzero"))
	s.Require().True(s.App.LiquidStakeKeeper.HasPoolWithDenom(s.Ctx, "uqi"))
	s.Require().False(s.App.LiquidStakeKeeper.HasPoolWithDenom(s.Ctx, "uother"))
}

func (s *KeeperTestSuite) TestGetAllPools() {
	s.SetupTest()
	s.seedPool("a", "ua")
	s.seedPool("b", "ub")
	s.seedPool("c", "uc")
	all := s.App.LiquidStakeKeeper.GetAllPools(s.Ctx)
	s.Require().Len(all, 3)
}

func (s *KeeperTestSuite) TestLiquidValidator_RoundTrip() {
	s.SetupTest()
	s.seedPool("zero", "uzero")
	valAddr := sdk.ValAddress(secp256k1.GenPrivKey().PubKey().Address())

	lv := types.LiquidValidator{
		OperatorAddress: valAddr.String(),
	}
	s.App.LiquidStakeKeeper.SetLiquidValidator(s.Ctx, "zero", lv)

	got, found := s.App.LiquidStakeKeeper.GetLiquidValidator(s.Ctx, "zero", valAddr)
	s.Require().True(found)
	s.Require().Equal(valAddr.String(), got.OperatorAddress)

	all := s.App.LiquidStakeKeeper.GetAllLiquidValidatorsForPool(s.Ctx, "zero")
	s.Require().Len(all, 1)

	s.App.LiquidStakeKeeper.RemoveLiquidValidator(s.Ctx, "zero", lv)
	_, found = s.App.LiquidStakeKeeper.GetLiquidValidator(s.Ctx, "zero", valAddr)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestModuleParams_RoundTrip() {
	s.SetupTest()
	got := s.App.LiquidStakeKeeper.GetModuleParams(s.Ctx)
	s.Require().NotNil(got)
}
