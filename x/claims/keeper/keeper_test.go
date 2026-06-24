package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/x/claims/types"
)

// KeeperTestSuite covers x/claims keeper-level CRUD across the four
// persistent entities (Collection / Claim / Dispute / Intent), plus
// genesis. Wasm-dependent flows (CreateCollection / SubmitClaim /
// EvaluateClaim / WithdrawPayment all call the cw20 / cw1155 contracts to
// settle payments) live in tests/interchaintest/.
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

func (s *KeeperTestSuite) seedCollection(id, entity, admin string) types.Collection {
	c := types.Collection{
		Id:       id,
		Entity:   entity,
		Admin:    admin,
		Protocol: "did:ixo:protocol",
		State:    types.CollectionState_open,
	}
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)
	return c
}

func (s *KeeperTestSuite) seedClaim(collectionID, claimID, agentDid, agentAddr string) types.Claim {
	c := types.Claim{
		CollectionId: collectionID,
		AgentDid:     agentDid,
		AgentAddress: agentAddr,
		ClaimId:      claimID,
	}
	s.App.ClaimsKeeper.SetClaim(s.Ctx, c)
	return c
}
