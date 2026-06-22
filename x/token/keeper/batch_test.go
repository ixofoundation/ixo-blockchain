package keeper_test

import (
	"cosmossdk.io/math"
	"go.uber.org/mock/gomock"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/x/token/types"
)

// seedTokenClass seeds a Token (keyed by name) with a given contract + supply,
// plus a TokenProperties row mapping id -> name, so GetTokenById(id) resolves it.
func (s *KeeperTestSuite) seedTokenClass(name, id, contract string, supply math.Uint) {
	s.App.TokenKeeper.SetToken(s.Ctx, types.Token{
		Minter:          apptesting.RandomAccountAddress().String(),
		ContractAddress: contract,
		Class:           "did:ixo:test-class",
		Name:            name,
		Cap:             math.NewUint(1_000_000),
		Supply:          supply,
	})
	s.App.TokenKeeper.SetTokenProperties(s.Ctx, types.TokenProperties{Id: id, Name: name, Index: id})
}

func (s *KeeperTestSuite) TestValidateTokenBatch_EmptyBatch() {
	s.SetupTest()
	_, err := s.App.TokenKeeper.ValidateTokenBatch(s.Ctx, nil)
	s.Require().ErrorIs(err, types.ErrTokenBatchInvalid)
}

func (s *KeeperTestSuite) TestValidateTokenBatch_EmptyId() {
	s.SetupTest()
	_, err := s.App.TokenKeeper.ValidateTokenBatch(s.Ctx, []*types.TokenBatch{{Id: "", Amount: math.NewUint(1)}})
	s.Require().ErrorIs(err, types.ErrTokenBatchInvalid)
}

func (s *KeeperTestSuite) TestValidateTokenBatch_ZeroAmount() {
	s.SetupTest()
	s.seedTokenClass("alpha", "a1", apptesting.RandomAccountAddress().String(), math.NewUint(10))
	_, err := s.App.TokenKeeper.ValidateTokenBatch(s.Ctx, []*types.TokenBatch{{Id: "a1", Amount: math.ZeroUint()}})
	s.Require().ErrorIs(err, types.ErrTokenAmountIncorrect)
}

func (s *KeeperTestSuite) TestValidateTokenBatch_MixedContract() {
	s.SetupTest()
	s.seedTokenClass("alpha", "a1", apptesting.RandomAccountAddress().String(), math.NewUint(10))
	s.seedTokenClass("beta", "b1", apptesting.RandomAccountAddress().String(), math.NewUint(10))
	_, err := s.App.TokenKeeper.ValidateTokenBatch(s.Ctx, []*types.TokenBatch{
		{Id: "a1", Amount: math.NewUint(1)},
		{Id: "b1", Amount: math.NewUint(1)},
	})
	s.Require().ErrorIs(err, types.ErrTokenContractMismatch)
}

func (s *KeeperTestSuite) TestValidateTokenBatch_HappyPath() {
	s.SetupTest()
	contract := apptesting.RandomAccountAddress().String()
	s.seedTokenClass("alpha", "a1", contract, math.NewUint(10))
	// a second id of the SAME class/contract.
	s.App.TokenKeeper.SetTokenProperties(s.Ctx, types.TokenProperties{Id: "a2", Name: "alpha", Index: "a2"})

	got, err := s.App.TokenKeeper.ValidateTokenBatch(s.Ctx, []*types.TokenBatch{
		{Id: "a1", Amount: math.NewUint(1)},
		{Id: "a2", Amount: math.NewUint(2)},
	})
	s.Require().NoError(err)
	s.Require().Equal(contract, got)
}

// TestCancelToken_ExceedsSupply_NoPanic is the F4 regression: cancelling more
// than the current supply must return an error, NOT panic on math.Uint
// underflow (the previous unchecked token.Supply.Sub(amount)).
func (s *KeeperTestSuite) TestCancelToken_ExceedsSupply_NoPanic() {
	s.SetupTest()
	mockWasm, ms, _ := s.withMockedWasm()

	contract := apptesting.RandomAccountAddress().String()
	s.seedTokenClass("gamma", "g1", contract, math.NewUint(3))
	owner := apptesting.RandomAccountAddress()

	// The cw1155 burn executes before the supply check; mock it as successful.
	mockWasm.EXPECT().
		Execute(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(nil), nil).
		Times(1)

	s.Require().NotPanics(func() {
		_, err := ms.CancelToken(s.Ctx, &types.MsgCancelToken{
			Owner:  owner.String(),
			Tokens: []*types.TokenBatch{{Id: "g1", Amount: math.NewUint(10)}},
		})
		s.Require().ErrorIs(err, types.ErrTokenAmountIncorrect)
	})

	// Supply untouched.
	_, token, err := s.App.TokenKeeper.GetTokenById(s.Ctx, "g1")
	s.Require().NoError(err)
	s.Require().Equal(math.NewUint(3).String(), token.Supply.String())
}
