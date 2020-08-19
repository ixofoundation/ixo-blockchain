package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmDB "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-blockchain/x/payments"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func CreateTestInput() (sdk.Context, Keeper, *codec.Codec,
	payments.Keeper, bank.Keeper) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	actStoreKey := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")
	keyPayments := sdk.NewKVStoreKey(payments.StoreKey)
	keyDid := sdk.NewKVStoreKey(did.StoreKey)

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyPayments, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyDid, sdk.StoreTypeIAVL, nil)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
	cdc := MakeTestCodec()

	pk1 := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(
		cdc, actStoreKey, pk1.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount,
	)

	paymentsSubspace := pk1.Subspace(payments.DefaultParamspace)
	projectSubspace := pk1.Subspace(types.DefaultParamspace)

	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk1.Subspace(bank.DefaultParamspace), nil)
	didKeeper := did.NewKeeper(cdc, keyDid)
	paymentsKeeper := payments.NewKeeper(cdc, keyPayments, paymentsSubspace, bankKeeper, didKeeper, nil)
	keeper := NewKeeper(cdc, storeKey, projectSubspace, accountKeeper, didKeeper, paymentsKeeper)

	paymentsKeeper.SetParams(ctx, payments.DefaultParams())

	return ctx, keeper, cdc, paymentsKeeper, bankKeeper
}

func MakeTestCodec() *codec.Codec {
	return codec.New()
}
