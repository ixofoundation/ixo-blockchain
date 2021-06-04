package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/ixofoundation/ixo-blockchain/cmd"
	paymentskeeper "github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	paymentstypes "github.com/ixofoundation/ixo-blockchain/x/payments/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"

	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func CreateTestInput(t *testing.T, isCheckTx bool) (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
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

	reservedIdPrefixes := make([]string, 0)

	appl.BankKeeper = bankkeeper.NewBaseKeeper(
		appl.AppCodec(),
		appl.GetKey(banktypes.StoreKey),
		appl.AccountKeeper,
		appl.ParamsKeeper.Subspace("bank"), //TODO (Stef) no banktypes.DefaultParamspace
		blockedAddrs,
	)
	appl.PaymentsKeeper = paymentskeeper.NewKeeper(
		appl.AppCodec(),
		appl.GetKey(paymentstypes.StoreKey),
		appl.BankKeeper,
		appl.DidKeeper,
		reservedIdPrefixes,
	)
	appl.ProjectKeeper = NewKeeper(
		appl.AppCodec(),
		appl.GetKey(types.StoreKey),
		appl.ParamsKeeper.Subspace(types.DefaultParamspace),
		appl.AccountKeeper,
		appl.DidKeeper,
		appl.PaymentsKeeper,
	)

	return appl.LegacyAmino(), appl, ctx

	//keyProject := sdk.NewKVStoreKey(types.StoreKey)
	//keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	//keyParams := sdk.NewKVStoreKey("subspace")
	//tkeyParams := sdk.NewTransientStoreKey("transient_params")
	//keyPayments := sdk.NewKVStoreKey(payments.StoreKey)
	//keyDid := sdk.NewKVStoreKey(did.StoreKey)

	//db := tmDB.NewMemDB()
	//ms := store.NewCommitMultiStore(db)
	//ms.MountStoreWithDB(keyProject, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(keyPayments, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(keyDid, sdk.StoreTypeIAVL, nil)
	//err := ms.LoadLatestVersion()
	//require.Nil(t, err)

	//ctx := sdk.NewContext(ms, abci.Header{ChainID: "foochainid"}, isCheckTx, log.NewNopLogger())
	//ctx = ctx.WithConsensusParams(
	//	&abci.ConsensusParams{
	//		Validator: &abci.ValidatorParams{
	//			PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519},
	//		},
	//	},
	//)
	////cdc := MakeTestCodec()
	//
	//feeCollectorAcc := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
	//notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
	//bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)
	//
	//blacklistedAddrs := make(map[string]bool)
	//blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
	//blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	//blacklistedAddrs[bondPool.GetAddress().String()] = true
	//
	//reservedIdPrefixes := make([]string, 0)
	//
	//pk := params.NewKeeper(cdc, keyParams, tkeyParams)
	//accountKeeper := auth.NewAccountKeeper(cdc, keyAcc, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	//bankKeeper := bank.NewBaseKeeper(accountKeeper, pk.Subspace(bank.DefaultParamspace), blacklistedAddrs)
	//didKeeper := did.NewKeeper(cdc, keyDid)
	//paymentsKeeper := payments.NewKeeper(cdc, keyPayments, bankKeeper, didKeeper, reservedIdPrefixes)
	//
	//keeper := NewKeeper(
	//	cdc, keyProject, pk.Subspace(types.DefaultParamspace), accountKeeper, didKeeper, paymentsKeeper,
	//)
	//
	//return ctx, keeper, paymentsKeeper, bankKeeper
}

//func MakeTestCodec() *codec.Codec {
//	var cdc = codec.New()
//	types.RegisterCodec(cdc)
//	sdk.RegisterCodec(cdc)
//	codec.RegisterCrypto(cdc)
//	auth.RegisterCodec(cdc)
//
//	return cdc
//}
