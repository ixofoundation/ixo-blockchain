package claims_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	claimstypes "github.com/ixofoundation/ixo-blockchain/v7/x/claims/types"
)

// ClaimsNoAnteTestSuite dispatches claims messages through the app's REAL
// MsgServiceRouter via RunMsg, which skips the ante handler — reproducing the
// CosmWasm / ICA-host / authz routes. It asserts that admin-gated claims
// operations authorise against the (signer) AdminAddress vs the state-pinned
// collection admin in the KEEPER, so authorization holds without the ante.
type ClaimsNoAnteTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestClaimsNoAnteTestSuite(t *testing.T) { suite.Run(t, new(ClaimsNoAnteTestSuite)) }

func (s *ClaimsNoAnteTestSuite) SetupTest() { s.Setup() }

func (s *ClaimsNoAnteTestSuite) seedAdminCollection(id string) string {
	admin := apptesting.RandomAccountAddress().String()
	s.App.ClaimsKeeper.SetCollection(s.Ctx, claimstypes.Collection{
		Id:       id,
		Entity:   "did:ixo:entity-claim",
		Admin:    admin,
		Protocol: "did:ixo:protocol-claim",
		State:    claimstypes.CollectionState_open,
		Intents:  claimstypes.CollectionIntentOptions_allow,
	})
	return admin
}

// TestUpdateCollectionState_NoAnte_RejectsNonAdmin: a non-admin signer is
// rejected by the keeper even with the ante skipped.
func (s *ClaimsNoAnteTestSuite) TestUpdateCollectionState_NoAnte_RejectsNonAdmin() {
	_ = s.seedAdminCollection("noante-claims")
	attacker := apptesting.RandomAccountAddress().String()

	_, err := s.RunMsg(&claimstypes.MsgUpdateCollectionState{
		CollectionId: "noante-claims",
		State:        claimstypes.CollectionState_paused,
		AdminAddress: attacker,
	})
	s.Require().Error(err)

	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "noante-claims")
	s.Require().Equal(claimstypes.CollectionState_open, got.State, "state must be unchanged")
}

// TestUpdateCollectionState_NoAnte_AllowsAdmin: the pinned admin succeeds on the
// same no-ante route.
func (s *ClaimsNoAnteTestSuite) TestUpdateCollectionState_NoAnte_AllowsAdmin() {
	admin := s.seedAdminCollection("noante-claims-ok")

	_, err := s.RunMsg(&claimstypes.MsgUpdateCollectionState{
		CollectionId: "noante-claims-ok",
		State:        claimstypes.CollectionState_paused,
		AdminAddress: admin,
	})
	s.Require().NoError(err)

	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "noante-claims-ok")
	s.Require().Equal(claimstypes.CollectionState_paused, got.State)
}
