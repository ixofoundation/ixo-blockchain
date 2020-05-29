package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmDB "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-blockchain/x/fees"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func CreateTestInput() (sdk.Context, Keeper, *codec.Codec, fees.Keeper, bank.Keeper) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	actStoreKey := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")
	keyFees := sdk.NewKVStoreKey(fees.StoreKey)

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyFees, sdk.StoreTypeIAVL, nil)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
	cdc := MakeTestCodec()

	pk1 := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	accountKeeper := auth.NewAccountKeeper(
		cdc, actStoreKey, pk1.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount,
	)

	feesSubspace := pk1.Subspace(fees.DefaultParamspace)
	projectSubspace := pk1.Subspace(types.DefaultParamspace)

	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk1.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, nil)
	feeKeeper := fees.NewKeeper(cdc, keyFees, feesSubspace, bankKeeper, nil)
	keeper := NewKeeper(cdc, storeKey, projectSubspace, accountKeeper, feeKeeper)

	feeKeeper.SetParams(ctx, fees.DefaultParams())

	return ctx, keeper, cdc, feeKeeper, bankKeeper
}

func MakeTestCodec() *codec.Codec {
	return codec.New()
}
