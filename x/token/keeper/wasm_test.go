package keeper_test

import (
	"crypto/md5"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/mock/gomock"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
	tokenkeeper "github.com/ixofoundation/ixo-blockchain/v8/x/token/keeper"
	"github.com/ixofoundation/ixo-blockchain/v8/x/token/testutil"
	"github.com/ixofoundation/ixo-blockchain/v8/x/token/types"
)

// withMockedWasm builds a fresh token keeper that delegates Instantiate /
// Execute to a gomock instead of the live wasmd keeper. The mock controller
// is registered with t.Cleanup so all EXPECTed calls are verified after the
// test runs.
//
// Returns (mockWasm, msgServer, keeper) — keeper is used by the caller to
// seed state directly via SetToken / SetTokenProperties when needed.
func (s *KeeperTestSuite) withMockedWasm() (*testutil.MockWasmKeeper, types.MsgServer, *tokenkeeper.Keeper) {
	ctrl := gomock.NewController(s.T())
	s.T().Cleanup(ctrl.Finish)

	mockWasm := testutil.NewMockWasmKeeper(ctrl)

	// Build a fresh Keeper that points at the live store + iid keeper but
	// uses our mock for wasm. The store key + paramSpace come from the
	// running app so persisted state is shared with the rest of the test.
	k := tokenkeeper.NewKeeper(
		s.App.AppCodec(),
		s.App.GetKey(types.StoreKey),
		s.App.IidKeeper,
		mockWasm,
		s.App.GetSubspace(types.ModuleName),
	)
	return mockWasm, tokenkeeper.NewMsgServerImpl(&k), &k
}

// seedClassDID registers a DID document for `class` so CreateToken's
// "class did document not found" guard passes.
func (s *KeeperTestSuite) seedClassDID(class string) {
	meta := iidtypes.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := iidtypes.IidDocument{Id: class, Metadata: &meta}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(class), doc)
}

func (s *KeeperTestSuite) TestMsgCreateToken_HappyPath() {
	s.SetupTest()
	mockWasm, ms, _ := s.withMockedWasm()

	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()
	classDID := "did:ixo:protocol-1"
	s.seedClassDID(classDID)

	// Wasm Instantiate gets called once with the encoded init msg; we don't
	// inspect bytes here (gomock would need the exact ixo1155.Marshal output)
	// — gomock.Any() is enough for shape verification, and the contract addr
	// we return is what gets persisted on the Token.
	mockWasm.EXPECT().
		Instantiate(gomock.Any(), gomock.Any(), minter, minter, gomock.Any(), gomock.Any(), gomock.Any()).
		Return(contractAddr, []byte(nil), nil).
		Times(1)

	_, err := ms.CreateToken(s.Ctx, &types.MsgCreateToken{
		Minter:      minter.String(),
		Class:       iidtypes.DIDFragment(classDID),
		Name:        "carbon-credit-001",
		Description: "test carbon credit",
		Image:       "ipfs://x",
		TokenType:   "ixo1155",
		Cap:         math.NewUint(1_000_000),
	})
	s.Require().NoError(err)

	// Persisted Token has the contract address returned by the wasm mock.
	got, err := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), contractAddr.String())
	s.Require().NoError(err)
	s.Require().Equal("carbon-credit-001", got.Name)
	s.Require().Equal(classDID, got.Class)
	s.Require().True(got.Supply.IsZero())
	s.AssertEventEmitted(s.Ctx, "ixo.token.v1beta1.TokenCreatedEvent", 1)
}

func (s *KeeperTestSuite) TestMsgCreateToken_MissingClassDID() {
	s.SetupTest()
	_, ms, _ := s.withMockedWasm()

	minter := apptesting.RandomAccountAddress()
	_, err := ms.CreateToken(s.Ctx, &types.MsgCreateToken{
		Minter: minter.String(),
		Class:  "did:ixo:does-not-exist",
		Name:   "no-class",
	})
	s.Require().ErrorContains(err, "class did document not found")
}

func (s *KeeperTestSuite) TestMsgCreateToken_DuplicateName() {
	s.SetupTest()
	mockWasm, ms, _ := s.withMockedWasm()

	classDID := "did:ixo:protocol-2"
	s.seedClassDID(classDID)
	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	mockWasm.EXPECT().Instantiate(
		gomock.Any(), gomock.Any(), minter, minter, gomock.Any(), gomock.Any(), gomock.Any(),
	).Return(contractAddr, []byte(nil), nil).Times(1)

	first := &types.MsgCreateToken{
		Minter: minter.String(), Class: iidtypes.DIDFragment(classDID), Name: "dup", TokenType: "ixo1155", Cap: math.NewUint(1),
	}
	_, err := ms.CreateToken(s.Ctx, first)
	s.Require().NoError(err)

	// Second create with same Name must be rejected before wasm is invoked.
	_, err = ms.CreateToken(s.Ctx, first)
	s.Require().ErrorContains(err, "token name is already taken")
}

func (s *KeeperTestSuite) TestMsgCreateToken_InstantiateError() {
	s.SetupTest()
	mockWasm, ms, _ := s.withMockedWasm()
	classDID := "did:ixo:protocol-3"
	s.seedClassDID(classDID)
	minter := apptesting.RandomAccountAddress()

	mockWasm.EXPECT().Instantiate(
		gomock.Any(), gomock.Any(), minter, minter, gomock.Any(), gomock.Any(), gomock.Any(),
	).Return(sdk.AccAddress(nil), []byte(nil), fmt.Errorf("wasm boom")).Times(1)

	_, err := ms.CreateToken(s.Ctx, &types.MsgCreateToken{
		Minter: minter.String(), Class: iidtypes.DIDFragment(classDID), Name: "fails", TokenType: "ixo1155", Cap: math.NewUint(1),
	})
	s.Require().ErrorContains(err, "wasm boom")

	// And no Token was persisted.
	_, err = s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), "any")
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestMsgMintToken_HappyPath() {
	s.SetupTest()
	mockWasm, ms, k := s.withMockedWasm()

	minter := apptesting.RandomAccountAddress()
	owner := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	// Pre-seed the Token directly via the keeper.
	tok := types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-mint", Name: "mintme",
		Cap: math.NewUint(1_000), Supply: math.ZeroUint(),
	}
	k.SetToken(s.Ctx, tok)

	mockWasm.EXPECT().Execute(
		gomock.Any(), contractAddr, minter, gomock.Any(), gomock.Any(),
	).Return([]byte(nil), nil).Times(1)

	_, err := ms.MintToken(s.Ctx, &types.MsgMintToken{
		Minter: minter.String(), ContractAddress: contractAddr.String(), Owner: owner.String(),
		MintBatch: []*types.MintBatch{
			{Name: "mintme", Index: "1", Amount: math.NewUint(50), Collection: "col"},
		},
	})
	s.Require().NoError(err)

	got, _ := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), contractAddr.String())
	s.Require().Equal(uint64(50), got.Supply.Uint64())

	// TokenProperties persisted under the deterministic md5(name+index) id.
	tokenID := fmt.Sprintf("%x", md5.Sum([]byte("mintme1")))
	tp, err := s.App.TokenKeeper.GetTokenProperties(s.Ctx, tokenID)
	s.Require().NoError(err)
	s.Require().Equal("mintme", tp.Name)
}

func (s *KeeperTestSuite) TestMsgMintToken_PausedRejected() {
	s.SetupTest()
	_, ms, k := s.withMockedWasm()
	minter := apptesting.RandomAccountAddress()
	owner := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	k.SetToken(s.Ctx, types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-pause", Name: "paused",
		Cap: math.NewUint(1_000), Supply: math.ZeroUint(),
		Paused: true,
	})

	_, err := ms.MintToken(s.Ctx, &types.MsgMintToken{
		Minter: minter.String(), ContractAddress: contractAddr.String(), Owner: owner.String(),
		MintBatch: []*types.MintBatch{{Name: "paused", Index: "1", Amount: math.NewUint(1)}},
	})
	s.Require().ErrorContains(err, "token is paused")
}

func (s *KeeperTestSuite) TestMsgMintToken_OverCap() {
	s.SetupTest()
	_, ms, k := s.withMockedWasm()
	minter := apptesting.RandomAccountAddress()
	owner := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	k.SetToken(s.Ctx, types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-cap", Name: "capped",
		Cap: math.NewUint(10), Supply: math.NewUint(8),
	})

	_, err := ms.MintToken(s.Ctx, &types.MsgMintToken{
		Minter: minter.String(), ContractAddress: contractAddr.String(), Owner: owner.String(),
		MintBatch: []*types.MintBatch{{Name: "capped", Index: "1", Amount: math.NewUint(5)}},
	})
	s.Require().ErrorContains(err, "is greater than token cap")
}

func (s *KeeperTestSuite) TestMsgTransferToken_HappyPath() {
	s.SetupTest()
	mockWasm, ms, k := s.withMockedWasm()

	owner := apptesting.RandomAccountAddress()
	recipient := apptesting.RandomAccountAddress()
	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	// Pre-seed Token + TokenProperties so GetTokenById resolves.
	k.SetToken(s.Ctx, types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-tr", Name: "transferme",
		Cap: math.NewUint(1_000), Supply: math.NewUint(100),
	})
	tokenID := fmt.Sprintf("%x", md5.Sum([]byte("transferme1")))
	k.SetTokenProperties(s.Ctx, types.TokenProperties{Id: tokenID, Index: "1", Name: "transferme"})

	mockWasm.EXPECT().Execute(
		gomock.Any(), contractAddr, owner, gomock.Any(), gomock.Any(),
	).Return([]byte(nil), nil).Times(1)

	_, err := ms.TransferToken(s.Ctx, &types.MsgTransferToken{
		Owner: owner.String(), Recipient: recipient.String(),
		Tokens: []*types.TokenBatch{{Id: tokenID, Amount: math.NewUint(10)}},
	})
	s.Require().NoError(err)
	s.AssertEventEmitted(s.Ctx, "ixo.token.v1beta1.TokenTransferredEvent", 1)
}

func (s *KeeperTestSuite) TestMsgRetireToken_HappyPath() {
	s.SetupTest()
	mockWasm, ms, k := s.withMockedWasm()

	owner := apptesting.RandomAccountAddress()
	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	k.SetToken(s.Ctx, types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-ret", Name: "retireme",
		Cap: math.NewUint(1_000), Supply: math.NewUint(100),
	})
	tokenID := fmt.Sprintf("%x", md5.Sum([]byte("retireme1")))
	k.SetTokenProperties(s.Ctx, types.TokenProperties{Id: tokenID, Index: "1", Name: "retireme"})

	mockWasm.EXPECT().Execute(
		gomock.Any(), contractAddr, owner, gomock.Any(), gomock.Any(),
	).Return([]byte(nil), nil).Times(1)

	_, err := ms.RetireToken(s.Ctx, &types.MsgRetireToken{
		Owner:        owner.String(),
		Tokens:       []*types.TokenBatch{{Id: tokenID, Amount: math.NewUint(5)}},
		Reason:       "burnt to retire credit",
		Jurisdiction: "ZA",
	})
	s.Require().NoError(err)

	got, _ := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), contractAddr.String())
	s.Require().Len(got.Retired, 1)
	s.Require().Equal(uint64(5), got.Retired[0].Amount.Uint64())
}

func (s *KeeperTestSuite) TestMsgCancelToken_HappyPath() {
	s.SetupTest()
	mockWasm, ms, k := s.withMockedWasm()

	owner := apptesting.RandomAccountAddress()
	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	k.SetToken(s.Ctx, types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-cancel", Name: "cancelme",
		Cap: math.NewUint(1_000), Supply: math.NewUint(100),
	})
	tokenID := fmt.Sprintf("%x", md5.Sum([]byte("cancelme1")))
	k.SetTokenProperties(s.Ctx, types.TokenProperties{Id: tokenID, Index: "1", Name: "cancelme"})

	mockWasm.EXPECT().Execute(
		gomock.Any(), contractAddr, owner, gomock.Any(), gomock.Any(),
	).Return([]byte(nil), nil).Times(1)

	_, err := ms.CancelToken(s.Ctx, &types.MsgCancelToken{
		Owner:  owner.String(),
		Tokens: []*types.TokenBatch{{Id: tokenID, Amount: math.NewUint(7)}},
		Reason: "wrong issuance",
	})
	s.Require().NoError(err)

	got, _ := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), contractAddr.String())
	s.Require().Equal(uint64(93), got.Supply.Uint64(), "Supply must drop by Cancelled amount")
	s.Require().Len(got.Cancelled, 1)
}

func (s *KeeperTestSuite) TestMsgTransferCredit_HappyPath() {
	s.SetupTest()
	mockWasm, ms, k := s.withMockedWasm()

	owner := apptesting.RandomAccountAddress()
	minter := apptesting.RandomAccountAddress()
	contractAddr := apptesting.RandomAccountAddress()

	k.SetToken(s.Ctx, types.Token{
		Minter: minter.String(), ContractAddress: contractAddr.String(),
		Class: "did:ixo:protocol-credit", Name: "creditme",
		Cap: math.NewUint(1_000), Supply: math.NewUint(100),
	})
	tokenID := fmt.Sprintf("%x", md5.Sum([]byte("creditme1")))
	k.SetTokenProperties(s.Ctx, types.TokenProperties{Id: tokenID, Index: "1", Name: "creditme"})

	mockWasm.EXPECT().Execute(
		gomock.Any(), contractAddr, owner, gomock.Any(), gomock.Any(),
	).Return([]byte(nil), nil).Times(1)

	_, err := ms.TransferCredit(s.Ctx, &types.MsgTransferCredit{
		Owner:           owner.String(),
		Tokens:          []*types.TokenBatch{{Id: tokenID, Amount: math.NewUint(3)}},
		Reason:          "credit transfer",
		Jurisdiction:    "ZA",
		AuthorizationId: "auth-1",
	})
	s.Require().NoError(err)

	got, _ := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), contractAddr.String())
	s.Require().Len(got.Transferred, 1)
}

