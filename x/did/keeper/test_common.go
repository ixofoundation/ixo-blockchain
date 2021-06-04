package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/ixofoundation/ixo-blockchain/cmd"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)


func CreateTestInput() (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
	appl := cmd.Setup(false)
	ctx := appl.BaseApp.NewContext(false, tmproto.Header{})

	appl.DidKeeper = NewKeeper(appl.AppCodec(), appl.GetKey(types.StoreKey))

	return appl.LegacyAmino(), appl, ctx

	//storeKey := sdk.NewKVStoreKey(types.StoreKey)
	//
	//db := dbm.NewMemDB()
	//ms := store.NewCommitMultiStore(db)
	//ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	//_ = ms.LoadLatestVersion()
	//ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
	//cdc := codec.New()
	//keeper := NewKeeper(cdc, storeKey)
	//
	//return ctx, keeper, cdc
}