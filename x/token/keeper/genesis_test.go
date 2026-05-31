package keeper_test

import (
	"cosmossdk.io/math"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/x/token/types"
)

func (s *KeeperTestSuite) TestGenesis_RoundTrip() {
	s.SetupTest()

	minter := apptesting.RandomAccountAddress()
	contract := apptesting.RandomAccountAddress()

	gs := types.GenesisState{
		Params: types.DefaultParams(),
		Tokens: []types.Token{
			{
				Minter:          minter.String(),
				ContractAddress: contract.String(),
				Class:           "did:ixo:abc",
				Name:            "alpha",
				Description:     "imported via genesis",
				Image:           "ipfs://x",
				Type:            "credit",
				Cap:             math.NewUint(123),
				Supply:          math.ZeroUint(),
			},
		},
		TokenProperties: []types.TokenProperties{
			{Id: "tp1", Name: "alpha", Index: "1"},
		},
	}

	s.App.TokenKeeper.InitGenesis(s.Ctx, gs)
	out := s.App.TokenKeeper.ExportGenesis(s.Ctx)
	s.Require().Len(out.Tokens, 1)
	s.Require().Equal("alpha", out.Tokens[0].Name)
	s.Require().Len(out.TokenProperties, 1)
}

func (s *KeeperTestSuite) TestDefaultGenesis_IsValid() {
	gs := types.DefaultGenesisState()
	s.Require().NoError(gs.Validate())
}
