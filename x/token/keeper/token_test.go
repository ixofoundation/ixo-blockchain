package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v7/x/token/types"
)

func (s *KeeperTestSuite) TestSetGetToken_RoundTrip() {
	s.SetupTest()
	minter, t := s.seedToken("alpha")

	got, err := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), t.ContractAddress)
	s.Require().NoError(err)
	s.Require().Equal(t.Name, got.Name)
	s.Require().Equal(t.Class, got.Class)
}

func (s *KeeperTestSuite) TestGetToken_NotFound() {
	s.SetupTest()
	_, err := s.App.TokenKeeper.GetToken(s.Ctx, "ixo1deadbeef", "ixo1deadbeef")
	s.Require().ErrorContains(err, "token not found")
}

func (s *KeeperTestSuite) TestGetTokens_AllAndByMinter() {
	s.SetupTest()
	a, ta := s.seedToken("alpha")
	_ = ta
	_, _ = s.seedToken("beta")

	all := s.App.TokenKeeper.GetTokens(s.Ctx)
	s.Require().Len(all, 2)

	mt := s.App.TokenKeeper.GetMinterTokens(s.Ctx, a.String())
	s.Require().GreaterOrEqual(len(mt), 1)
}

func (s *KeeperTestSuite) TestCheckTokensDuplicateName() {
	s.SetupTest()
	_, _ = s.seedToken("unique-name")
	s.Require().True(s.App.TokenKeeper.CheckTokensDuplicateName(s.Ctx, "unique-name"))
	s.Require().False(s.App.TokenKeeper.CheckTokensDuplicateName(s.Ctx, "fresh-name"))
}

func (s *KeeperTestSuite) TestGetTokenByName() {
	s.SetupTest()
	_, t := s.seedToken("findme")
	got, found := s.App.TokenKeeper.GetTokenByName(s.Ctx, "findme")
	s.Require().True(found)
	s.Require().Equal(t.ContractAddress, got.ContractAddress)

	_, found = s.App.TokenKeeper.GetTokenByName(s.Ctx, "missing")
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestTokenProperties_RoundTrip() {
	s.SetupTest()
	tp := types.TokenProperties{
		Id:    "tok-id-1",
		Name:  "alpha",
		Index: "1",
	}
	s.App.TokenKeeper.SetTokenProperties(s.Ctx, tp)

	got, err := s.App.TokenKeeper.GetTokenProperties(s.Ctx, "tok-id-1")
	s.Require().NoError(err)
	s.Require().Equal(tp.Name, got.Name)

	_, err = s.App.TokenKeeper.GetTokenProperties(s.Ctx, "missing")
	s.Require().ErrorContains(err, "token properties not found")
}
