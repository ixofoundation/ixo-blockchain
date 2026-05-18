package apptesting

import (
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var testDescription = stakingtypes.NewDescription(
	"test_moniker",
	"test_identity",
	"test_website",
	"test_security_contact",
	"test_details",
)

// SetupValidator creates a fresh validator with self-bond = DefaultPowerReduction
// and forces it into bondStatus regardless of the staking module's normal flow.
// Returns the validator's operator address.
//
// The bond denom is read from the staking params (initialised to `uixo` by
// app.Setup), so this works on a vanilla ixo test app.
func (s *KeeperTestHelper) SetupValidator(bondStatus stakingtypes.BondStatus) sdk.ValAddress {
	valPub := secp256k1.GenPrivKey().PubKey()
	valAddr := sdk.ValAddress(valPub.Address())

	stakingParams, err := s.App.StakingKeeper.GetParams(s.Ctx)
	s.Require().NoError(err)
	bondDenom := stakingParams.BondDenom
	bondAmt := sdk.DefaultPowerReduction
	selfBond := sdk.NewCoins(sdk.NewCoin(bondDenom, bondAmt))

	s.FundAcc(sdk.AccAddress(valAddr), selfBond)

	zeroDec := sdkmath.LegacyZeroDec()
	zeroCommission := stakingtypes.NewCommissionRates(zeroDec, zeroDec, zeroDec)

	createMsg, err := stakingtypes.NewMsgCreateValidator(
		valAddr.String(),
		valPub,
		sdk.NewCoin(bondDenom, bondAmt),
		testDescription,
		zeroCommission,
	)
	s.Require().NoError(err)

	stakingMsgSvr := stakingkeeper.NewMsgServerImpl(s.App.StakingKeeper)
	_, err = stakingMsgSvr.CreateValidator(s.Ctx, createMsg)
	s.Require().NoError(err)

	val, err := s.App.StakingKeeper.GetValidator(s.Ctx, valAddr)
	s.Require().NoError(err)

	val = val.UpdateStatus(bondStatus)
	s.Require().NoError(s.App.StakingKeeper.SetValidator(s.Ctx, val))

	consAddr, err := val.GetConsAddr()
	s.Require().NoError(err)

	signingInfo := slashingtypes.NewValidatorSigningInfo(
		sdk.ConsAddress(consAddr),
		s.Ctx.BlockHeight(),
		0,
		time.Unix(0, 0),
		false,
		0,
	)
	s.Require().NoError(s.App.SlashingKeeper.SetValidatorSigningInfo(s.Ctx, consAddr, signingInfo))

	return valAddr
}

// SetupMultipleValidators creates n bonded validators and returns their
// operator addresses as bech32 strings.
func (s *KeeperTestHelper) SetupMultipleValidators(n int) []string {
	out := make([]string, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, s.SetupValidator(stakingtypes.Bonded).String())
	}
	return out
}
