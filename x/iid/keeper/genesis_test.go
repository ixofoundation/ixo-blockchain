package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// TestGenesis_RoundTrip seeds two DID documents through SetDidDocument,
// exports genesis, then re-imports into a fresh keeper and asserts each DID
// is still queryable.
func (s *KeeperTestSuite) TestGenesis_RoundTrip() {
	s.SetupTest()
	_, didA := s.freshSigner()
	_, didB := s.freshSigner()

	gs := s.App.IidKeeper.ExportGenesis(s.Ctx)
	ids := map[string]bool{}
	for _, d := range gs.IidDocs {
		ids[d.Id] = true
	}
	s.Require().True(ids[didA])
	s.Require().True(ids[didB])

	// Reset the suite so we have a clean keeper, then InitGenesis from the export
	// and assert documents survive.
	s.SetupTest()
	s.App.IidKeeper.InitGenesis(s.Ctx, *gs)
	gotA, foundA := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(didA))
	s.Require().True(foundA)
	s.Require().Equal(didA, gotA.Id)
	gotB, foundB := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(didB))
	s.Require().True(foundB)
	s.Require().Equal(didB, gotB.Id)
}

func (s *KeeperTestSuite) TestDefaultGenesis_IsEmpty() {
	gs := types.DefaultGenesisState()
	s.Require().NotNil(gs)
	s.Require().Empty(gs.IidDocs)
}
