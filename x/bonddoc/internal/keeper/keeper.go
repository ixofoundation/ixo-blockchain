package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
	didTypes "github.com/ixofoundation/ixo-cosmos/x/did"
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

func (k Keeper) GetBondDoc(ctx sdk.Context, bondDid ixo.Did) (types.StoredBondDoc, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetBondPrefixKey(bondDid)

	bz := store.Get(key)
	if bz == nil {
		return nil, didTypes.ErrorInvalidDid(types.DefaultCodeSpace, "Invalid BondDid Address")
	}

	var bondDoc types.CreateBondMsg
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &bondDoc)

	return &bondDoc, nil
}

func (k Keeper) SetBondDoc(ctx sdk.Context, bondDoc types.StoredBondDoc) sdk.Error {
	existedDoc, err := k.GetBondDoc(ctx, bondDoc.GetBondDid())
	if existedDoc != nil {
		return didTypes.ErrorInvalidDid(types.DefaultCodeSpace, fmt.Sprintf("Bond already exists %s", err))
	}

	k.AddBondDoc(ctx, bondDoc)

	return nil
}

func (k Keeper) AddBondDoc(ctx sdk.Context, bondDoc types.StoredBondDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetBondPrefixKey(bondDoc.GetBondDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(bondDoc))
}
