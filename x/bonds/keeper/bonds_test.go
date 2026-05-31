package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/lib/ixo"
	bondskeeper "github.com/ixofoundation/ixo-blockchain/v7/x/bonds/keeper"
	"github.com/ixofoundation/ixo-blockchain/v7/x/bonds/types"
)

func (s *KeeperTestSuite) TestSetGetBond_RoundTrip() {
	s.SetupTest()
	bondDid := s.seedBond("alpha", types.HatchState)
	got, found := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().True(found)
	s.Require().Equal("alpha", got.Token)
	s.Require().Equal(types.HatchState.String(), got.State)
}

func (s *KeeperTestSuite) TestBondExists() {
	s.SetupTest()
	s.seedBond("beta", types.HatchState)
	s.Require().True(s.App.BondsKeeper.BondExists(s.Ctx, "did:ixo:bond-beta"))
	s.Require().False(s.App.BondsKeeper.BondExists(s.Ctx, "did:ixo:nonexistent"))
}

func (s *KeeperTestSuite) TestBondDid_Lookup() {
	s.SetupTest()
	bondDid := s.seedBond("gamma", types.HatchState)
	got, found := s.App.BondsKeeper.GetBondDid(s.Ctx, "gamma")
	s.Require().True(found)
	s.Require().Equal(bondDid, got)
}

func (s *KeeperTestSuite) TestSetCurrentSupply() {
	s.SetupTest()
	bondDid := s.seedBond("supply", types.HatchState)
	s.App.BondsKeeper.SetCurrentSupply(s.Ctx, bondDid, sdk.NewCoin("supply", math.NewInt(42)))
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal(int64(42), got.CurrentSupply.Amount.Int64())
}

// TestSetBondState confirms the keeper persists state changes and emits
// BondUpdatedEvent. The state machine itself (HATCH→OPEN→SETTLE) is
// validated independently in types_test.go.
func (s *KeeperTestSuite) TestSetBondState_EmitsEvent() {
	s.SetupTest()
	bondDid := s.seedBond("statetest", types.HatchState)
	s.App.BondsKeeper.SetBondState(s.Ctx, bondDid, types.OpenState.String())
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal(types.OpenState.String(), got.State)
	s.AssertEventEmitted(s.Ctx, "ixo.bonds.v1beta1.BondUpdatedEvent", 1)
}

// TestDepositIntoReserve and TestWithdrawFromReserve exercise the bank
// interactions through the bonds reserve module account.
func (s *KeeperTestSuite) TestDepositIntoReserve_AndWithdraw() {
	s.SetupTest()
	bondDid := s.seedBond("res", types.OpenState)
	denom := ixo.IxoNativeToken
	depositor := apptesting.RandomAccountAddress()
	s.FundAcc(depositor, sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(1_000))))

	err := s.App.BondsKeeper.DepositIntoReserve(s.Ctx, bondDid, depositor, sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(500))))
	s.Require().NoError(err)
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal(math.NewInt(500), got.CurrentReserve.AmountOf(denom))
	s.Require().Equal(math.NewInt(500), got.AvailableReserve.AmountOf(denom))

	withdrawTo := apptesting.RandomAccountAddress()
	err = s.App.BondsKeeper.WithdrawFromReserve(s.Ctx, bondDid, withdrawTo, sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(200))))
	s.Require().NoError(err)
	got, _ = s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal(math.NewInt(300), got.CurrentReserve.AmountOf(denom))
	s.Require().Equal(math.NewInt(200), s.App.BankKeeper.GetBalance(s.Ctx, withdrawTo, denom).Amount)
}

// TestMsgBuy_FirstSwapperBuy exercises the full liquidity-init special case:
// when a SwapperFunction bond has CurrentSupply=0, the first MsgBuy treats
// MaxPrices as the seed liquidity, mints CurrentSupply bond tokens and sends
// them to the buyer, and updates CurrentReserve.
func (s *KeeperTestSuite) TestMsgBuy_FirstSwapperBuy() {
	s.SetupTest()
	bondDid, buyer, buyerDID := s.seedSwapperBondWithBuyer("swap")

	maxPrices := sdk.NewCoins(
		sdk.NewCoin("uatom", math.NewInt(1_000)),
		sdk.NewCoin(ixo.IxoNativeToken, math.NewInt(1_000)),
	)
	s.FundAcc(buyer, maxPrices)

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.Buy(s.Ctx, &types.MsgBuy{
		BuyerDid:     buyerDID,
		BondDid:      bondDid,
		Amount:       sdk.NewCoin("swap", math.NewInt(100)),
		MaxPrices:    maxPrices,
		BuyerAddress: buyer.String(),
	})
	s.Require().NoError(err)

	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal(int64(100), got.CurrentSupply.Amount.Int64())
	s.Require().Equal(maxPrices.Sort(), got.CurrentReserve.Sort())

	bal := s.App.BankKeeper.GetBalance(s.Ctx, buyer, "swap")
	s.Require().Equal(int64(100), bal.Amount.Int64())
	s.Require().True(s.App.BankKeeper.GetBalance(s.Ctx, buyer, "uatom").IsZero())
	s.Require().True(s.App.BankKeeper.GetBalance(s.Ctx, buyer, ixo.IxoNativeToken).IsZero())
}

// TestMsgBuy_BondNotFound rejects a buy targeting an unknown bond.
func (s *KeeperTestSuite) TestMsgBuy_BondNotFound() {
	s.SetupTest()
	_, buyer, buyerDID := s.seedSwapperBondWithBuyer("buynope")
	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.Buy(s.Ctx, &types.MsgBuy{
		BuyerDid: buyerDID,
		BondDid:  "did:ixo:does-not-exist",
		Amount:   sdk.NewCoin("buynope", math.NewInt(1)),
		MaxPrices: sdk.NewCoins(
			sdk.NewCoin(ixo.IxoNativeToken, math.NewInt(1)),
		),
		BuyerAddress: buyer.String(),
	})
	s.Require().ErrorContains(err, "bond does not exist")
}

// TestMsgBuy_TokenMismatch rejects a buy whose Amount.Denom is not the
// bond's token.
func (s *KeeperTestSuite) TestMsgBuy_TokenMismatch() {
	s.SetupTest()
	bondDid, buyer, buyerDID := s.seedSwapperBondWithBuyer("mismatch")
	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.Buy(s.Ctx, &types.MsgBuy{
		BuyerDid:     buyerDID,
		BondDid:      bondDid,
		Amount:       sdk.NewCoin("not-the-token", math.NewInt(1)),
		MaxPrices:    sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, math.NewInt(1))),
		BuyerAddress: buyer.String(),
	})
	s.Require().ErrorContains(err, "bond token does not match bond")
}

// TestMsgEditBond updates the metadata fields of a bond. The msg uses
// DoNotModifyField sentinels for fields that should not change.
func (s *KeeperTestSuite) TestMsgEditBond() {
	s.SetupTest()
	bondDid := s.seedBond("editme", types.HatchState)
	// EditBond requires bond.CreatorDid == msg.EditorDid; seedBond doesn't
	// set CreatorDid, so we patch it before the test.
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	got.CreatorDid = "did:ixo:creator#key-1"
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, got)

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.EditBond(s.Ctx, &types.MsgEditBond{
		BondDid:                bondDid,
		Name:                   "renamed",
		Description:            types.DoNotModifyField,
		OrderQuantityLimits:    types.DoNotModifyField,
		SanityRate:             types.DoNotModifyField,
		SanityMarginPercentage: types.DoNotModifyField,
		EditorDid:              "did:ixo:creator#key-1",
		EditorAddress:          apptesting.RandomAccountAddress().String(),
	})
	s.Require().NoError(err)
	got, _ = s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal("renamed", got.Name)
	s.Require().Equal("test bond", got.Description, "unchanged because DoNotModifyField")
}

// TestMsgEditBond_WrongEditor rejects a non-creator's edit attempt.
func (s *KeeperTestSuite) TestMsgEditBond_WrongEditor() {
	s.SetupTest()
	bondDid := s.seedBond("permissions", types.HatchState)
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	got.CreatorDid = "did:ixo:owner#key-1"
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, got)

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.EditBond(s.Ctx, &types.MsgEditBond{
		BondDid:                bondDid,
		Name:                   types.DoNotModifyField,
		Description:            types.DoNotModifyField,
		OrderQuantityLimits:    types.DoNotModifyField,
		SanityRate:             types.DoNotModifyField,
		SanityMarginPercentage: types.DoNotModifyField,
		EditorDid:              "did:ixo:stranger#key-1",
		EditorAddress:          apptesting.RandomAccountAddress().String(),
	})
	s.Require().ErrorContains(err, "editor must be the creator of the bond")
}

func (s *KeeperTestSuite) TestReservedBondToken() {
	s.SetupTest()
	params := s.App.BondsKeeper.GetParams(s.Ctx)
	if len(params.ReservedBondTokens) > 0 {
		s.Require().True(s.App.BondsKeeper.ReservedBondToken(s.Ctx, params.ReservedBondTokens[0]))
	}
	s.Require().False(s.App.BondsKeeper.ReservedBondToken(s.Ctx, "definitely-not-reserved-12345"))
}

// TestMsgUpdateBondState_AugmentedHatchToOpen exercises the AugmentedFunction
// state-machine progression. UpdateBondState requires bond.FunctionType ==
// AugmentedFunction; for SwapperFunction bonds it returns
// ErrFunctionNotAvailableForFunctionType.
func (s *KeeperTestSuite) TestMsgUpdateBondState_AugmentedHatchToOpen() {
	s.SetupTest()
	bondDid := s.seedBond("aug", types.HatchState)
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	got.FunctionType = types.AugmentedFunction
	got.ControllerDid = "did:ixo:controller#key-1"
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, got)
	s.App.BondsKeeper.SetBatch(s.Ctx, bondDid, types.NewBatch(bondDid, got.Token, math.NewUint(1)))

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.UpdateBondState(s.Ctx, &types.MsgUpdateBondState{
		BondDid:       bondDid,
		State:         types.OpenState.String(),
		EditorDid:     "did:ixo:controller#key-1",
		EditorAddress: apptesting.RandomAccountAddress().String(),
	})
	s.Require().NoError(err)

	got, _ = s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().Equal(types.OpenState.String(), got.State)
}

// TestMsgUpdateBondState_NotAugmentedRejected
func (s *KeeperTestSuite) TestMsgUpdateBondState_NotAugmentedRejected() {
	s.SetupTest()
	bondDid := s.seedBond("not-aug", types.HatchState)
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.App.BondsKeeper.SetBatch(s.Ctx, bondDid, types.NewBatch(bondDid, got.Token, math.NewUint(1)))

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.UpdateBondState(s.Ctx, &types.MsgUpdateBondState{
		BondDid:   bondDid,
		State:     types.OpenState.String(),
		EditorDid: "did:ixo:editor#key-1",
	})
	s.Require().Error(err)
}

// TestMsgUpdateBondState_InvalidProgression: HATCH -> SETTLE is invalid.
func (s *KeeperTestSuite) TestMsgUpdateBondState_InvalidProgression() {
	s.SetupTest()
	bondDid := s.seedBond("invalid-prog", types.HatchState)
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	got.FunctionType = types.AugmentedFunction
	got.ControllerDid = "did:ixo:c#key-1"
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, got)
	s.App.BondsKeeper.SetBatch(s.Ctx, bondDid, types.NewBatch(bondDid, got.Token, math.NewUint(1)))

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.UpdateBondState(s.Ctx, &types.MsgUpdateBondState{
		BondDid:   bondDid,
		State:     types.SettleState.String(),
		EditorDid: "did:ixo:c#key-1",
	})
	s.Require().Error(err)
}

// TestMsgMakeOutcomePayment routes a payment through the iid + DID
// resolution path. The bond must be in OPEN state.
func (s *KeeperTestSuite) TestMsgMakeOutcomePayment() {
	s.SetupTest()
	bondDid, sender, senderDID := s.seedSwapperBondWithBuyer("oc-pay")
	// seedSwapperBondWithBuyer leaves the bond in HATCH; bump to OPEN.
	got, _ := s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	got.State = types.OpenState.String()
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, got)

	// Fund the sender with both reserve denoms — bond.GetNewReserveCoins(N)
	// returns N of every reserve token.
	s.FundAcc(sender, sdk.NewCoins(
		sdk.NewCoin(ixo.IxoNativeToken, math.NewInt(1_000)),
		sdk.NewCoin("uatom", math.NewInt(1_000)),
	))

	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.MakeOutcomePayment(s.Ctx, &types.MsgMakeOutcomePayment{
		SenderDid:     senderDID,
		Amount:        math.NewInt(100),
		BondDid:       bondDid,
		SenderAddress: sender.String(),
	})
	s.Require().NoError(err)

	got, _ = s.App.BondsKeeper.GetBond(s.Ctx, bondDid)
	s.Require().False(got.CurrentOutcomePaymentReserve.IsZero())
}

// TestMsgMakeOutcomePayment_WrongState
func (s *KeeperTestSuite) TestMsgMakeOutcomePayment_WrongState() {
	s.SetupTest()
	bondDid, sender, senderDID := s.seedSwapperBondWithBuyer("oc-pay-wrong")
	msgServer := bondskeeper.NewMsgServerImpl(s.App.BondsKeeper)
	_, err := msgServer.MakeOutcomePayment(s.Ctx, &types.MsgMakeOutcomePayment{
		SenderDid:     senderDID,
		Amount:        math.NewInt(1),
		BondDid:       bondDid,
		SenderAddress: sender.String(),
	})
	s.Require().Error(err)
}
