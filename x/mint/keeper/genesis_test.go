package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v7/x/mint/types"
)

func (s *KeeperTestSuite) TestGenesis_RoundTrip() {
	s.SetupTest()

	gs := s.App.MintKeeper.ExportGenesis(s.Ctx)
	s.Require().NotNil(gs)
	s.Require().NotNil(gs.Minter)
	s.Require().NotNil(gs.Params)

	// Re-init from the export and confirm params survive.
	s.App.MintKeeper.InitGenesis(s.Ctx, gs)
	got := s.App.MintKeeper.GetParams(s.Ctx)
	s.Require().Equal(gs.Params.MintDenom, got.MintDenom)
}

func (s *KeeperTestSuite) TestDefaultGenesis_IsValid() {
	gs := types.DefaultGenesisState()
	s.Require().NoError(types.ValidateGenesis(*gs))
}
