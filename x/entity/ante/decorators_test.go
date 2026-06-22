package ante_test

import (
	"testing"

	"cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	entityante "github.com/ixofoundation/ixo-blockchain/v7/x/entity/ante"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v7/x/entity/types"
)

// EntityAnteTestSuite covers BlockNftContractTransferForEntityDecorator: it
// must reject any MsgExecuteContract whose target contract address matches
// the params.NftContractAddress, and pass everything else through.
type EntityAnteTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestEntityAnteTestSuite(t *testing.T) { suite.Run(t, new(EntityAnteTestSuite)) }

func (s *EntityAnteTestSuite) SetupTest() { s.Setup() }

// minimalTx implements sdk.Tx with a single message — enough to exercise the
// decorator's per-message inspection.
type minimalTx struct{ msgs []sdk.Msg }

func (t *minimalTx) GetMsgs() []sdk.Msg                                { return t.msgs }
func (t *minimalTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error)   { return nil, nil }
func (t *minimalTx) ValidateBasic() error                              { return nil }

func (s *EntityAnteTestSuite) TestBlockNftContractTransfer_BlocksTargetContract() {
	s.SetupTest()

	// Configure the entity params with a known NFT contract address.
	nftAddr := apptesting.RandomAccountAddress()
	params := entitytypes.DefaultParams()
	params.NftContractAddress = nftAddr.String()
	s.App.EntityKeeper.SetParams(s.Ctx, &params)

	dec := entityante.NewBlockNftContractTransferForEntityDecorator(s.App.EntityKeeper)

	// Tx targeting the NFT contract directly — must be rejected.
	tx := &minimalTx{msgs: []sdk.Msg{&wasmtypes.MsgExecuteContract{
		Sender:   apptesting.RandomAccountAddress().String(),
		Contract: nftAddr.String(),
		Msg:      []byte(`{"transfer":{}}`),
		Funds:    sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1))),
	}}}

	_, err := dec.AnteHandle(s.Ctx, tx, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) { return ctx, nil })
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "cannot execute contract set as the entity nft contract address")
}

// TestBlockNftContractTransfer_BlocksInsideMsgExec is the F3 regression: a
// MsgExecuteContract against the NFT contract hidden inside an authz.MsgExec
// must also be blocked (previously the decorator only inspected top-level msgs).
func (s *EntityAnteTestSuite) TestBlockNftContractTransfer_BlocksInsideMsgExec() {
	s.SetupTest()

	nftAddr := apptesting.RandomAccountAddress()
	params := entitytypes.DefaultParams()
	params.NftContractAddress = nftAddr.String()
	s.App.EntityKeeper.SetParams(s.Ctx, &params)

	dec := entityante.NewBlockNftContractTransferForEntityDecorator(s.App.EntityKeeper)

	inner := &wasmtypes.MsgExecuteContract{
		Sender:   apptesting.RandomAccountAddress().String(),
		Contract: nftAddr.String(),
		Msg:      []byte(`{"transfer_nft":{}}`),
	}
	exec := authz.NewMsgExec(apptesting.RandomAccountAddress(), []sdk.Msg{inner})
	tx := &minimalTx{msgs: []sdk.Msg{&exec}}

	called := false
	_, err := dec.AnteHandle(s.Ctx, tx, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		called = true
		return ctx, nil
	})
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "cannot execute contract set as the entity nft contract address")
	s.Require().False(called)
}

func (s *EntityAnteTestSuite) TestBlockNftContractTransfer_AllowsOtherContracts() {
	s.SetupTest()

	nftAddr := apptesting.RandomAccountAddress()
	params := entitytypes.DefaultParams()
	params.NftContractAddress = nftAddr.String()
	s.App.EntityKeeper.SetParams(s.Ctx, &params)

	dec := entityante.NewBlockNftContractTransferForEntityDecorator(s.App.EntityKeeper)

	other := apptesting.RandomAccountAddress()
	tx := &minimalTx{msgs: []sdk.Msg{&wasmtypes.MsgExecuteContract{
		Sender:   apptesting.RandomAccountAddress().String(),
		Contract: other.String(),
		Msg:      []byte(`{"action":{}}`),
	}}}

	called := false
	_, err := dec.AnteHandle(s.Ctx, tx, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		called = true
		return ctx, nil
	})
	s.Require().NoError(err)
	s.Require().True(called, "decorator must pass through to the next handler when contract is not the NFT address")
}

func (s *EntityAnteTestSuite) TestBlockNftContractTransfer_AllowsNonWasmMessages() {
	s.SetupTest()
	dec := entityante.NewBlockNftContractTransferForEntityDecorator(s.App.EntityKeeper)

	// A non-wasm message — should pass through regardless of params state.
	// We re-use bank's MsgSend as a stand-in non-wasm message.
	tx := &minimalTx{msgs: []sdk.Msg{&entitytypes.MsgCreateEntity{}}}

	_, err := dec.AnteHandle(s.Ctx, tx, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) { return ctx, nil })
	s.Require().NoError(err)
}
