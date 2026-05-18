package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v6/x/entity/types"
)

// KeeperTestSuite covers x/entity keeper-level CRUD, account creation, and
// genesis. The wasm-dependent surface (CreateEntity instantiates an NFT
// cw721 contract; GetCurrentOwner queries the NFT contract; UpdateEntity /
// TransferEntity / UpdateEntityVerified all proxy through the contract) is
// covered in tests/interchaintest/ where a real wasmvm is available.
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

// seedEntity inserts a valid Entity directly via SetEntity, bypassing the
// wasm-dependent CreateEntity flow.
func (s *KeeperTestSuite) seedEntity(id string) types.Entity {
	now := time.Now().UTC()
	meta := types.NewEntityMetadata([]byte("v1"), now)
	e := types.Entity{
		Id:       id,
		Type:     "asset",
		Status:   1,
		Metadata: &meta,
	}
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(id), e)
	return e
}
