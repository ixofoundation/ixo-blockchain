package bonds_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v8/x/bonds/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// BondsDisabledE2ETestSuite drives the disabled bonds module through the real,
// fully-wired IxoApp: messages go through the actual MsgServiceRouter (hitting
// the registered disabled msg server) and the no-op batch EndBlocker runs as
// part of the app's EndBlocker.
type BondsDisabledE2ETestSuite struct {
	apptesting.KeeperTestHelper
}

func TestBondsDisabledE2ETestSuite(t *testing.T) {
	suite.Run(t, new(BondsDisabledE2ETestSuite))
}

func (s *BondsDisabledE2ETestSuite) SetupTest() { s.Setup() }

// TestBondsMsgsRejectedByRouter confirms that bonds messages, dispatched
// through the app's real MsgServiceRouter, reach the registered disabled msg
// server and are rejected with ErrBondsModuleDisabled. This is the
// authoritative disable layer and covers every routed path (authz dispatch,
// CosmWasm stargate, ICA-host) since they all funnel through this same router.
//
// We use MakeOutcomePayment (the drain message used in the incident) and
// WithdrawShare (the cash-out message). Both are fully valid here so they pass
// the router's ValidateBasic step and actually reach the disabled handler —
// proving the rejection comes from the disable, not from input validation.
func (s *BondsDisabledE2ETestSuite) TestBondsMsgsRejectedByRouter() {
	addr := apptesting.RandomAccountAddress().String()
	did := "did:ixo:" + addr
	frag := iidtypes.DIDFragment(did + "#v1")

	msgs := []sdk.Msg{
		&bondstypes.MsgMakeOutcomePayment{BondDid: "did:ixo:bond-x", SenderDid: frag, SenderAddress: addr, Amount: math.NewInt(1)},
		&bondstypes.MsgWithdrawShare{BondDid: "did:ixo:bond-x", RecipientDid: frag, RecipientAddress: addr},
	}

	for _, msg := range msgs {
		s.Run(sdk.MsgTypeURL(msg), func() {
			// Sanity: the message itself is valid, so any rejection is from the disable.
			s.Require().NoError(msg.(sdk.HasValidateBasic).ValidateBasic())
			_, err := s.RunMsg(msg)
			s.Require().Error(err)
			s.Require().ErrorIs(err, bondstypes.ErrBondsModuleDisabled)
		})
	}
}

// TestEndBlockerIsNoOp seeds a bond + batch with blocks remaining and confirms
// the app EndBlocker does NOT decrement the batch / process orders — proving
// the bonds batch EndBlocker is neutralised. With the original EndBlocker this
// batch's BlocksRemaining would drop from 5 to 4.
func (s *BondsDisabledE2ETestSuite) TestEndBlockerIsNoOp() {
	addr := apptesting.RandomAccountAddress()
	const bondDid = "did:ixo:bond-endblock"
	bond := bondstypes.NewBond(
		"abc", "name", "desc",
		iidtypes.DIDFragment("did:ixo:creator"),
		iidtypes.DIDFragment("did:ixo:controller"),
		iidtypes.DIDFragment("did:ixo:oracle"),
		bondstypes.SwapperFunction, bondstypes.FunctionParams{}, []string{"uixo"},
		math.LegacyZeroDec(), math.LegacyZeroDec(),
		addr, addr,
		sdk.NewCoin("abc", math.NewInt(1_000_000)), sdk.Coins{},
		math.LegacyZeroDec(), math.LegacyZeroDec(),
		false, false, false, math.NewUint(10), math.ZeroInt(),
		bondstypes.OpenState, bondDid,
	)
	s.App.BondsKeeper.SetBond(s.Ctx, bondDid, bond)
	s.App.BondsKeeper.SetBatch(s.Ctx, bondDid, bondstypes.NewBatch(bondDid, bond.Token, math.NewUint(5)))

	// Run the app's EndBlocker (which invokes the bonds module EndBlock).
	s.EndBlock()

	got := s.App.BondsKeeper.MustGetBatch(s.Ctx, bondDid)
	s.Require().Equal(math.NewUint(5).String(), got.BlocksRemaining.String(),
		"bonds EndBlocker must be a no-op while the module is disabled")
}
