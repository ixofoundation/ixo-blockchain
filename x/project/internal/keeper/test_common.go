package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	cParams "github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmDB "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/params"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func CreateTestInput() (sdk.Context, Keeper, *codec.Codec, fees.Keeper, bank.Keeper, params.Keeper) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	actStoreKey := sdk.NewKVStoreKey(auth.StoreKey)
	keyParam := sdk.NewKVStoreKey(params.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey("transient_params")
	keyFee := sdk.NewKVStoreKey(fees.StoreKey)

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyParam, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyFee, sdk.StoreTypeIAVL, nil)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
	cdc := MakeTestCodec()

	paramsKeeper := params.NewKeeper(cdc, keyParam)

	pk1 := cParams.NewKeeper(cdc, keyParam, tkeyParams, cParams.DefaultCodespace)
	accountKeeper := auth.NewAccountKeeper(
		cdc, actStoreKey, pk1.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount,
	)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk1.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, nil)
	feeKeeper := fees.NewKeeper(cdc, paramsKeeper)
	keeper := NewKeeper(cdc, storeKey, accountKeeper, feeKeeper)

	return ctx, keeper, cdc, feeKeeper, bankKeeper, paramsKeeper
}

func MakeTestCodec() *codec.Codec {
	return codec.New()
}
