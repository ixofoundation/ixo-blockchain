package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-cosmos/x/did/internal/types"
)

func CreateTestInput() (sdk.Context, Keeper, *codec.Codec) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	_ = ms.LoadLatestVersion()
	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
	cdc := codec.New()
	keeper := NewKeeper(cdc, storeKey)

	return ctx, keeper, cdc
}
