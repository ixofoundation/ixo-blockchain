package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/ixofoundation/ixo-cosmos/types"
	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// SetFiatPeg:
func (k Keeper) SetFiatPeg(ctx sdk.Context, fiatPeg types.FiatPeg) {
	store := ctx.KVStore(k.storeKey)

	fiatPegHash := fiatTypes.FiatPegHashStoreKey(fiatPeg.GetPegHash())
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(fiatPeg)
	store.Set(fiatPegHash, bz)
}

// returns fiatPeg by pegHash
func (k Keeper) GetFiatPeg(ctx sdk.Context, pegHash types.PegHash) (fiatPeg types.FiatPeg, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	fiatPegKey := fiatTypes.FiatPegHashStoreKey(pegHash)
	bz := store.Get(fiatPegKey)
	if bz == nil {
		return nil, fiatTypes.ErrInvalidPegHash(fiatTypes.DefaultCodeSpace)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fiatPeg)
	return fiatPeg, nil
}

// get all fiatPegs => []FiatPeg from store
func (k Keeper) GetFiatPegs(ctx sdk.Context) (fiatPegs []types.FiatPeg) {
	k.IterateFiatPegs(ctx, func(fiatPeg types.FiatPeg) (stop bool) {
		fiatPegs = append(fiatPegs, fiatPeg)
		return false
	},
	)
	return
}

func (k Keeper) IterateFiatPegs(ctx sdk.Context, handler func(fiatPeg types.FiatPeg) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, fiatTypes.PegHashKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var fiatPeg types.FiatPeg
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &fiatPeg)
		if handler(fiatPeg) {
			break
		}
	}
}

func (k Keeper) GetFiatPegDetails(ctx sdk.Context, buyerAddress sdk.AccAddress, sellerAddress sdk.AccAddress,
	hash types.PegHash) (types.FiatPeg, sdk.Error) {

	pegHash := types.PegHash(append(append(buyerAddress.Bytes(), sellerAddress.Bytes()...), hash.Bytes()...))
	_fiatPeg, err := k.GetFiatPeg(ctx, pegHash)
	if err != nil {
		return nil, err
	}
	return _fiatPeg, nil
}
