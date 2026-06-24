package apptesting_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/lib/ixo"
)

// SmokeTestSuite verifies the apptesting harness boots a fresh IxoApp,
// produces a usable context, and lets us mutate state via the bank keeper.
// If this suite ever fails, every downstream module test will too — fix here
// first.
type SmokeTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestSmokeTestSuite(t *testing.T) {
	suite.Run(t, new(SmokeTestSuite))
}

func (s *SmokeTestSuite) SetupTest() {
	s.Setup()
}

// TestSetup_FundsAndAccounts confirms test accounts exist and FundAcc / balance
// queries round-trip through the bank keeper.
func (s *SmokeTestSuite) TestSetup_FundsAndAccounts() {
	s.Require().Len(s.TestAccs, 3, "expected 3 baseline test accounts")

	addr := s.TestAccs[0]
	denom := ixo.IxoNativeToken
	want := sdkmath.NewInt(123_456_789)

	gotBefore := s.App.BankKeeper.GetBalance(s.Ctx, addr, denom)
	s.Require().True(gotBefore.IsZero(), "fresh account should start at zero, got %s", gotBefore)

	s.FundAcc(addr, sdk.NewCoins(sdk.NewCoin(denom, want)))

	gotAfter := s.App.BankKeeper.GetBalance(s.Ctx, addr, denom)
	s.Require().Equal(want, gotAfter.Amount, "FundAcc should top up the account")
}

// TestSetup_ChainID confirms the ctx is initialised on the test chain id.
func (s *SmokeTestSuite) TestSetup_ChainID() {
	s.Require().Equal("ixo-test-1", s.Ctx.ChainID())
}

// TestSetup_ValidatorBonded confirms a single bonded validator was registered
// in the staking module by the genesis builder.
func (s *SmokeTestSuite) TestSetup_ValidatorBonded() {
	vals, err := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	s.Require().NoError(err)
	s.Require().Len(vals, 1, "expected one bonded validator from GenesisStateWithValSet")
	s.Require().True(vals[0].IsBonded(), "validator should be bonded")
}

// TestSetupValidator_AddsBondedValidator confirms the SetupValidator helper
// adds a fresh bonded validator on top of the genesis one.
func (s *SmokeTestSuite) TestSetupValidator_AddsBondedValidator() {
	before, err := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	s.Require().NoError(err)

	_ = s.SetupValidator(stakingtypes.Bonded)

	after, err := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	s.Require().NoError(err)
	s.Require().Equal(len(before)+1, len(after))
}
