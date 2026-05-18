package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	liquidstakekeeper "github.com/ixofoundation/ixo-blockchain/v6/x/liquidstake/keeper"
	"github.com/ixofoundation/ixo-blockchain/v6/x/liquidstake/types"
)

// govAuthority returns the gov module account bech32 address — required as
// the Authority for governance-only flows (UpdateModuleParams, SetModulePaused).
func (s *KeeperTestSuite) govAuthority() string {
	return authtypes.NewModuleAddress(govtypes.ModuleName).String()
}

// TestMsgSetPoolPaused: authorised pool admin flips Pool.Paused.
func (s *KeeperTestSuite) TestMsgSetPoolPaused() {
	s.SetupTest()
	want := s.seedPool("paused-test", "upaused")

	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	_, err := msgServer.SetPoolPaused(s.Ctx, &types.MsgSetPoolPaused{
		Authority: want.WhitelistAdminAddress,
		PoolId:    want.PoolId,
		IsPaused:  true,
	})
	s.Require().NoError(err)
	got, _ := s.App.LiquidStakeKeeper.GetPool(s.Ctx, want.PoolId)
	s.Require().True(got.Paused)
}

// TestMsgSetPoolPaused_Unauthorised rejects a stranger.
func (s *KeeperTestSuite) TestMsgSetPoolPaused_Unauthorised() {
	s.SetupTest()
	want := s.seedPool("auth-test", "uauth")

	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	stranger := apptesting.RandomAccountAddress()
	_, err := msgServer.SetPoolPaused(s.Ctx, &types.MsgSetPoolPaused{
		Authority: stranger.String(),
		PoolId:    want.PoolId,
		IsPaused:  true,
	})
	s.Require().Error(err)
}

// TestMsgSetModulePaused: governance flips the module-wide kill switch.
func (s *KeeperTestSuite) TestMsgSetModulePaused() {
	s.SetupTest()
	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	_, err := msgServer.SetModulePaused(s.Ctx, &types.MsgSetModulePaused{
		Authority: s.govAuthority(),
		IsPaused:  true,
	})
	s.Require().NoError(err)
	params := s.App.LiquidStakeKeeper.GetModuleParams(s.Ctx)
	s.Require().True(params.ModulePaused)

	// Restore.
	_, _ = msgServer.SetModulePaused(s.Ctx, &types.MsgSetModulePaused{
		Authority: s.govAuthority(),
		IsPaused:  false,
	})
}

// TestMsgSetModulePaused_NonGovRejected: non-gov authority is rejected.
func (s *KeeperTestSuite) TestMsgSetModulePaused_NonGovRejected() {
	s.SetupTest()
	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	stranger := apptesting.RandomAccountAddress()
	_, err := msgServer.SetModulePaused(s.Ctx, &types.MsgSetModulePaused{
		Authority: stranger.String(),
		IsPaused:  true,
	})
	s.Require().Error(err)
}

// TestMsgBurn: burner sends uixo to the module then it's burned.
func (s *KeeperTestSuite) TestMsgBurn() {
	s.SetupTest()
	burner := apptesting.RandomAccountAddress()
	s.FundAcc(burner, sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000))))

	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	_, err := msgServer.Burn(s.Ctx, &types.MsgBurn{
		Burner: burner.String(),
		Amount: sdk.NewCoin("uixo", math.NewInt(400)),
	})
	s.Require().NoError(err)

	bal := s.App.BankKeeper.GetBalance(s.Ctx, burner, "uixo").Amount
	s.Require().Equal(int64(600), bal.Int64())
}

// TestMsgBurn_NonUixoRejected: only uixo is burnable through this module.
func (s *KeeperTestSuite) TestMsgBurn_NonUixoRejected() {
	s.SetupTest()
	burner := apptesting.RandomAccountAddress()
	s.FundAcc(burner, sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(1_000))))

	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	_, err := msgServer.Burn(s.Ctx, &types.MsgBurn{
		Burner: burner.String(),
		Amount: sdk.NewCoin("uatom", math.NewInt(100)),
	})
	s.Require().ErrorContains(err, "burning amount must be in uixo")
}

// TestMsgCreatePool: governance creates a new pool with a valid id+denom.
func (s *KeeperTestSuite) TestMsgCreatePool() {
	s.SetupTest()
	ms := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)

	admin := apptesting.RandomAccountAddress()
	feeAcc := apptesting.RandomAccountAddress()
	resp, err := ms.CreatePool(s.Ctx, &types.MsgCreatePool{
		Authority:                s.govAuthority(),
		PoolId:                   "newpool",
		LiquidBondDenom:          "unewpool",
		InitialAdminAddress:      admin.String(),
		InitialFeeAccountAddress: feeAcc.String(),
	})
	s.Require().NoError(err)
	s.Require().NotEmpty(resp.ProxyAccountAddress)

	got, found := s.App.LiquidStakeKeeper.GetPool(s.Ctx, "newpool")
	s.Require().True(found)
	s.Require().Equal("unewpool", got.LiquidBondDenom)
	s.Require().Equal(admin.String(), got.WhitelistAdminAddress)
	s.Require().Equal(feeAcc.String(), got.FeeAccountAddress)
}

// TestMsgCreatePool_NonGovRejected
func (s *KeeperTestSuite) TestMsgCreatePool_NonGovRejected() {
	s.SetupTest()
	ms := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	stranger := apptesting.RandomAccountAddress()
	_, err := ms.CreatePool(s.Ctx, &types.MsgCreatePool{
		Authority:                stranger.String(),
		PoolId:                   "stranger",
		LiquidBondDenom:          "ustr",
		InitialAdminAddress:      stranger.String(),
		InitialFeeAccountAddress: stranger.String(),
	})
	s.Require().Error(err)
}

// TestMsgUpdatePool: pool admin updates the pool's mutable fields.
func (s *KeeperTestSuite) TestMsgUpdatePool() {
	s.SetupTest()
	pool := s.seedPool("upd-pool", "uupdpool")

	ms := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	newFeeAcc := apptesting.RandomAccountAddress()
	newAdmin := apptesting.RandomAccountAddress()
	_, err := ms.UpdatePool(s.Ctx, &types.MsgUpdatePool{
		Authority:             pool.WhitelistAdminAddress,
		PoolId:                pool.PoolId,
		UnstakeFeeRate:        pool.UnstakeFeeRate,
		FeeAccountAddress:     newFeeAcc.String(),
		AutocompoundFeeRate:   pool.AutocompoundFeeRate,
		WhitelistAdminAddress: newAdmin.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.LiquidStakeKeeper.GetPool(s.Ctx, pool.PoolId)
	s.Require().Equal(newFeeAcc.String(), got.FeeAccountAddress)
	s.Require().Equal(newAdmin.String(), got.WhitelistAdminAddress)
}

// TestMsgUpdatePool_Unauthorised
func (s *KeeperTestSuite) TestMsgUpdatePool_Unauthorised() {
	s.SetupTest()
	pool := s.seedPool("unauth-pool", "uunauth")
	ms := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)
	stranger := apptesting.RandomAccountAddress()

	_, err := ms.UpdatePool(s.Ctx, &types.MsgUpdatePool{
		Authority:             stranger.String(),
		PoolId:                pool.PoolId,
		UnstakeFeeRate:        pool.UnstakeFeeRate,
		FeeAccountAddress:     stranger.String(),
		AutocompoundFeeRate:   pool.AutocompoundFeeRate,
		WhitelistAdminAddress: stranger.String(),
	})
	s.Require().Error(err)
}

// TestMsgUpdateModuleParams: governance updates module params in full.
func (s *KeeperTestSuite) TestMsgUpdateModuleParams() {
	s.SetupTest()
	msgServer := liquidstakekeeper.NewMsgServerImpl(s.App.LiquidStakeKeeper)

	current := s.App.LiquidStakeKeeper.GetModuleParams(s.Ctx)
	current.ModulePaused = true

	_, err := msgServer.UpdateModuleParams(s.Ctx, &types.MsgUpdateModuleParams{
		Authority:    s.govAuthority(),
		ModuleParams: current,
	})
	s.Require().NoError(err)
	s.Require().True(s.App.LiquidStakeKeeper.GetModuleParams(s.Ctx).ModulePaused)
}
