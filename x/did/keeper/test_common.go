package keeper

///*
//func CreateTestInput() (sdk.Context, Keeper, *codec.Codec) {
//	storeKey := sdk.NewKVStoreKey(types.StoreKey)
//
//	db := dbm.NewMemDB()
//	ms := store.NewCommitMultiStore(db)
//	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
//	_ = ms.LoadLatestVersion()
//	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
//	cdc := codec.New()
//	keeper := NewKeeper(cdc, storeKey)
//
//	return ctx, keeper, cdc
//}
//*/
//
//// TODO tests now generate app (simapp.Setup) instead of using CreateTestInput()
//
//func CreateTestInput() (sdk.Context, Keeper, codec.LegacyAmino) {
//	storeKey := sdk.NewKVStoreKey(types.StoreKey)
//
//	db := dbm.NewMemDB()
//	ms := store.NewCommitMultiStore(db)
//	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
//	_ = ms.LoadLatestVersion()
//	ctx := sdk.NewContext(ms, tmproto.Header{}, true, log.NewNopLogger())
//	marshaler, legacyAmino:= app.MakeCodecs() //using app here gives import cycle not allowed error
//	keeper := NewKeeper(marshaler, storeKey)
//
//	return ctx, keeper, *legacyAmino
//}