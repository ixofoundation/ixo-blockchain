package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v8/x/entity/types"
)

func (s *KeeperTestSuite) TestSetGetEntity_RoundTrip() {
	s.SetupTest()
	id := "did:ixo:entity-1"
	want := s.seedEntity(id)

	got, found := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(id))
	s.Require().True(found)
	s.Require().Equal(want.Id, got.Id)
	s.Require().Equal(want.Type, got.Type)
}

func (s *KeeperTestSuite) TestGetEntity_NotFound() {
	s.SetupTest()
	_, found := s.App.EntityKeeper.GetEntity(s.Ctx, []byte("nope"))
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestGetAllEntity() {
	s.SetupTest()
	s.seedEntity("did:ixo:entity-1")
	s.seedEntity("did:ixo:entity-2")
	s.seedEntity("did:ixo:entity-3")
	all := s.App.EntityKeeper.GetAllEntity(s.Ctx)
	s.Require().Len(all, 3)
}

func (s *KeeperTestSuite) TestGetAllEntityWithCondition() {
	s.SetupTest()
	a := s.seedEntity("did:ixo:entity-asset")
	a.Type = "asset"
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(a.Id), a)
	b := s.seedEntity("did:ixo:entity-asset-2")
	b.Type = "asset"
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(b.Id), b)
	c := s.seedEntity("did:ixo:entity-protocol")
	c.Type = "protocol"
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(c.Id), c)

	assets := s.App.EntityKeeper.GetAllEntityWithCondition(
		s.Ctx, types.EntityKey,
		func(e types.Entity) bool { return e.Type == "asset" },
	)
	s.Require().Len(assets, 2)
}

func (s *KeeperTestSuite) TestCreateNewAccount() {
	s.SetupTest()
	id := "did:ixo:account-test"
	addr, err := s.App.EntityKeeper.CreateNewAccount(s.Ctx, id, "primary")
	s.Require().NoError(err)
	s.Require().NotNil(addr)

	// Re-creating the same module account is rejected.
	_, err = s.App.EntityKeeper.CreateNewAccount(s.Ctx, id, "primary")
	s.Require().ErrorContains(err, "account already exists")
}

func (s *KeeperTestSuite) TestParams_RoundTrip() {
	s.SetupTest()
	got := s.App.EntityKeeper.GetParams(s.Ctx)
	s.Require().NotNil(got)
}
