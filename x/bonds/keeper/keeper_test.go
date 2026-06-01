package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/lib/ixo"
	"github.com/ixofoundation/ixo-blockchain/v7/x/bonds/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

// KeeperTestSuite covers x/bonds keeper-level operations: CRUD, reserve
// movement, state machine, and genesis. The bonding-curve math itself
// (Buy/Sell/Swap pricing, batch settlement) lives in the existing
// x/bonds/types/bondingfunctions_test.go pure-function tests and is exercised
// end-to-end via tests/interchaintest/ where a real validator-set + epochs
// drive the BeginBlock/EndBlock batch lifecycle.
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

// seedBond inserts a minimal Bond directly via SetBond (skipping CreateBond's
// validation chain, which requires a DID document, function-param checks, and
// reserve-token validation). Returns the bond's DID.
func (s *KeeperTestSuite) seedBond(token string, state types.BondState) string {
	bondDid := "did:ixo:bond-" + token
	feeAddr := apptesting.RandomAccountAddress()
	withdrawalAddr := apptesting.RandomAccountAddress()
	reserveDenom := ixo.IxoNativeToken

	b := types.Bond{
		Token:                        token,
		Name:                         "Test Bond " + token,
		Description:                  "test bond",
		FeeAddress:                   feeAddr.String(),
		ReserveWithdrawalAddress:     withdrawalAddr.String(),
		ReserveTokens:                []string{reserveDenom},
		FunctionType:                 types.SwapperFunction,
		MaxSupply:                    sdk.NewCoin(token, math.NewInt(1_000_000)),
		CurrentSupply:                sdk.NewCoin(token, math.ZeroInt()),
		CurrentReserve:               sdk.NewCoins(),
		AvailableReserve:             sdk.NewCoins(),
		CurrentOutcomePaymentReserve: sdk.NewCoins(),
		AllowSells:                   true,
		AllowReserveWithdrawals:      false,
		AlphaBond:                    false,
		State:                        state.String(),
		BondDid:                      bondDid,
		// Required-but-unused decimal/uint fields. These are mandatory because
		// math.LegacyDec / math.Uint zero-values panic on marshal.
		TxFeePercentage:        math.LegacyZeroDec(),
		ExitFeePercentage:      math.LegacyZeroDec(),
		SanityRate:             math.LegacyZeroDec(),
		SanityMarginPercentage: math.LegacyZeroDec(),
		BatchBlocks:            math.NewUint(1),
		OutcomePayment:         math.ZeroInt(),
	}
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, b)
	s.App.BondsKeeper.SetBondDid(s.Ctx, token, bondDid)
	return bondDid
}

// seedSwapperBondWithBuyer registers a bond that:
//   - uses SwapperFunction with two reserve tokens (uixo, uatom)
//   - is in HATCH state with zero supply (so the next Buy is the
//     "first swapper buy" liquidity-init special case)
//
// Plus a buyer addr + DID with a verification method that resolves to the
// buyer's bech32 address. Returns (bondDid, buyerAddr, buyerDIDFragment).
func (s *KeeperTestSuite) seedSwapperBondWithBuyer(token string) (string, sdk.AccAddress, iidtypes.DIDFragment) {
	bondDid := "did:ixo:bond-" + token
	feeAddr := apptesting.RandomAccountAddress()
	withdrawalAddr := apptesting.RandomAccountAddress()

	b := types.Bond{
		Token:                        token,
		Name:                         "Swapper Bond " + token,
		Description:                  "swapper test",
		FeeAddress:                   feeAddr.String(),
		ReserveWithdrawalAddress:     withdrawalAddr.String(),
		ReserveTokens:                []string{"uatom", ixo.IxoNativeToken},
		FunctionType:                 types.SwapperFunction,
		MaxSupply:                    sdk.NewCoin(token, math.NewInt(1_000_000)),
		CurrentSupply:                sdk.NewCoin(token, math.ZeroInt()),
		CurrentReserve:               sdk.NewCoins(),
		AvailableReserve:             sdk.NewCoins(),
		CurrentOutcomePaymentReserve: sdk.NewCoins(),
		AllowSells:                   true,
		State:                        types.HatchState.String(),
		BondDid:                      bondDid,
		TxFeePercentage:              math.LegacyZeroDec(),
		ExitFeePercentage:            math.LegacyZeroDec(),
		SanityRate:                   math.LegacyZeroDec(),
		SanityMarginPercentage:       math.LegacyZeroDec(),
		BatchBlocks:                  math.NewUint(1),
		OutcomePayment:               math.ZeroInt(),
	}
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, b)
	s.App.BondsKeeper.SetBondDid(s.Ctx, token, bondDid)

	// Seed an empty batch — Buy reads MustGetBatch.
	s.App.BondsKeeper.SetBatch(s.Ctx, bondDid, types.NewBatch(bondDid, b.Token, math.NewUint(1)))

	// Buyer + DID document.
	buyer := apptesting.RandomAccountAddress()
	buyerDid := "did:ixo:buyer-" + token
	methodID := buyerDid + "#key-1"
	vm := iidtypes.NewVerificationMethod(methodID, iidtypes.DID(buyerDid),
		iidtypes.NewBlockchainAccountID(buyer.String()))
	meta := iidtypes.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := iidtypes.IidDocument{
		Id:                 buyerDid,
		VerificationMethod: []*iidtypes.VerificationMethod{&vm},
		Authentication:     []string{methodID},
		Metadata:           &meta,
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(buyerDid), doc)

	return bondDid, buyer, iidtypes.DIDFragment(methodID)
}
