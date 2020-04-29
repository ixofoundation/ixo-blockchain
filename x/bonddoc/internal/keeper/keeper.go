package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type Keeper struct {
	cdc      *codec.Codec
	storeKey sdk.StoreKey
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: key,
	}
}

func (k Keeper) GetBondDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.BondKey)
}

func (k Keeper) MustGetBondDocByKey(ctx sdk.Context, key []byte) types.StoredBondDoc {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("bond doc not found")
	}

	bz := store.Get(key)
	var bondDoc types.MsgCreateBond
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &bondDoc)

	return &bondDoc
}

func (k Keeper) BondDocExists(ctx sdk.Context, bondDid ixo.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetBondPrefixKey(bondDid))
}

func (k Keeper) GetBondDoc(ctx sdk.Context, bondDid ixo.Did) (types.StoredBondDoc, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetBondPrefixKey(bondDid)

	bz := store.Get(key)
	if bz == nil {
		return nil, did.ErrorInvalidDid(types.DefaultCodeSpace, "Invalid BondDid Address")
	}

	var bondDoc types.MsgCreateBond
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &bondDoc)

	return &bondDoc, nil
}

func (k Keeper) SetBondDoc(ctx sdk.Context, bondDoc types.StoredBondDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetBondPrefixKey(bondDoc.GetBondDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(bondDoc))
}

func (k Keeper) UpdateBondDoc(ctx sdk.Context, newBondDoc types.StoredBondDoc) (types.StoredBondDoc, sdk.Error) {
	existedDoc, _ := k.GetBondDoc(ctx, newBondDoc.GetBondDid())
	if existedDoc == nil {

		return nil, did.ErrorInvalidDid(types.DefaultCodeSpace, "ProjectDoc details are not exist")
	} else {

		existedDoc.SetStatus(newBondDoc.GetStatus())
		k.SetBondDoc(ctx, newBondDoc)

		return newBondDoc, nil
	}
}
