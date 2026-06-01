package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/x/token/keeper"
	"github.com/ixofoundation/ixo-blockchain/v7/x/token/types"
)

// KeeperTestSuite covers the non-wasm surface of x/token: keeper CRUD,
// PauseToken/StopToken (which do not call into the ixo1155 contract), and
// genesis. Wasm-dependent flows (CreateToken, MintToken, TransferToken,
// RetireToken, CancelToken, TransferCredit) require a live wasmvm and are
// covered in tests/interchaintest/ instead.
type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	msgServer   types.MsgServer
	queryServer types.QueryServer
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) { suite.Run(t, new(KeeperTestSuite)) }

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.msgServer = keeper.NewMsgServerImpl(&s.App.TokenKeeper)
	s.queryServer = s.App.TokenKeeper
	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func (s *KeeperTestSuite) goCtx() context.Context { return s.Ctx }

// seedToken inserts a Token directly via SetToken (skipping CreateToken's
// wasm contract instantiation) so we can exercise downstream flows.
func (s *KeeperTestSuite) seedToken(name string) (sdk.AccAddress, types.Token) {
	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()
	t := types.Token{
		Minter:          minter.String(),
		ContractAddress: contractAddr.String(),
		Class:           "did:ixo:test-class",
		Name:            name,
		Description:     "test",
		Image:           "ipfs://x",
		Type:            "carbon-credit",
		Cap:             math.NewUint(1_000_000),
		Supply:          math.ZeroUint(),
		Paused:          false,
		Stopped:         false,
	}
	s.App.TokenKeeper.SetToken(s.Ctx, t)
	return minter, t
}
