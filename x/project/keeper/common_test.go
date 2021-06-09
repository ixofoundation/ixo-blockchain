package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/ixofoundation/ixo-blockchain/cmd"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

func CreateTestInput() (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
	appl := cmd.Setup(false)
	ctx := appl.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithConsensusParams(
		&abci.ConsensusParams{
			Validator: &tmproto.ValidatorParams{
				PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519},
			},
		},
	)

	feeCollectorAcc := authtypes.NewEmptyModuleAccount(authtypes.FeeCollectorName)
	notBondedPool := authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName, authtypes.Burner, authtypes.Staking)
	bondPool := authtypes.NewEmptyModuleAccount(stakingtypes.BondedPoolName, authtypes.Burner, authtypes.Staking)

	blockedAddrs := make(map[string]bool)
	blockedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blockedAddrs[notBondedPool.GetAddress().String()] = true
	blockedAddrs[bondPool.GetAddress().String()] = true

	appl.BankKeeper = bankkeeper.NewBaseKeeper(
		appl.AppCodec(),
		appl.GetKey(banktypes.StoreKey),
		appl.AccountKeeper,
		appl.GetSubspace(banktypes.ModuleName),
		blockedAddrs,
	)

	return appl.LegacyAmino(), appl, ctx
}
