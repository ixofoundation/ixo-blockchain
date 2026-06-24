package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/x/claims/types"
)

func (s *KeeperTestSuite) TestCollection_RoundTrip() {
	s.SetupTest()
	admin := apptesting.RandomAccountAddress().String()
	want := s.seedCollection("1", "did:ixo:entity", admin)

	got, err := s.App.ClaimsKeeper.GetCollection(s.Ctx, "1")
	s.Require().NoError(err)
	s.Require().Equal(want.Entity, got.Entity)
	s.Require().Equal(want.Admin, got.Admin)
	s.Require().Equal(types.CollectionState_open, got.State)
}

func (s *KeeperTestSuite) TestCollection_NotFound() {
	s.SetupTest()
	_, err := s.App.ClaimsKeeper.GetCollection(s.Ctx, "missing")
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestGetCollections() {
	s.SetupTest()
	admin := apptesting.RandomAccountAddress().String()
	s.seedCollection("1", "did:ixo:e1", admin)
	s.seedCollection("2", "did:ixo:e2", admin)
	all := s.App.ClaimsKeeper.GetCollections(s.Ctx)
	s.Require().Len(all, 2)
}

func (s *KeeperTestSuite) TestClaim_RoundTrip() {
	s.SetupTest()
	agent := apptesting.RandomAccountAddress()
	want := s.seedClaim("1", "claim-1", "did:ixo:agent", agent.String())

	got, err := s.App.ClaimsKeeper.GetClaim(s.Ctx, "claim-1")
	s.Require().NoError(err)
	s.Require().Equal(want.AgentDid, got.AgentDid)
	s.Require().Equal(want.CollectionId, got.CollectionId)
}

func (s *KeeperTestSuite) TestGetClaims() {
	s.SetupTest()
	agent := apptesting.RandomAccountAddress()
	s.seedClaim("1", "claim-a", "did:ixo:a", agent.String())
	s.seedClaim("1", "claim-b", "did:ixo:b", agent.String())
	all := s.App.ClaimsKeeper.GetClaims(s.Ctx)
	s.Require().Len(all, 2)
}

func (s *KeeperTestSuite) TestDispute_RoundTrip() {
	s.SetupTest()
	d := types.Dispute{
		SubjectId: "claim-x",
		Type:      1,
		Data: &types.DisputeData{
			Uri:    "ipfs://disp",
			Type:   "wrongdoing",
			Proof:  "proof-1",
			Encrypted: false,
		},
	}
	s.App.ClaimsKeeper.SetDispute(s.Ctx, d)

	got, err := s.App.ClaimsKeeper.GetDispute(s.Ctx, "proof-1")
	s.Require().NoError(err)
	s.Require().Equal("claim-x", got.SubjectId)
}

func (s *KeeperTestSuite) TestIntent_RoundTrip() {
	s.SetupTest()
	agent := apptesting.RandomAccountAddress()
	from := apptesting.RandomAccountAddress()
	escrow := apptesting.RandomAccountAddress()
	intent := types.Intent{
		Id:            "intent-1",
		AgentDid:      "did:ixo:agent",
		AgentAddress:  agent.String(),
		CollectionId:  "1",
		FromAddress:   from.String(),
		EscrowAddress: escrow.String(),
	}
	s.App.ClaimsKeeper.SetIntent(s.Ctx, intent)

	got, err := s.App.ClaimsKeeper.GetIntent(s.Ctx, agent.String(), "1", "intent-1")
	s.Require().NoError(err)
	s.Require().Equal("intent-1", got.Id)

	all := s.App.ClaimsKeeper.GetIntents(s.Ctx)
	s.Require().Len(all, 1)
}

func (s *KeeperTestSuite) TestParams_RoundTrip() {
	s.SetupTest()
	got := s.App.ClaimsKeeper.GetParams(s.Ctx)
	s.Require().NotNil(got)
}
