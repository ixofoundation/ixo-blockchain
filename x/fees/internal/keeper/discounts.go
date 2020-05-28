package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

// -------------------------------------------------------- DiscountHolders

func (k Keeper) GetFeeDiscountHoldersIterator(ctx sdk.Context, feeId string, discountId uint64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.GetDiscountHoldersKey(feeId, discountId))
}

func (k Keeper) GetFeeDiscountsHoldersIterator(ctx sdk.Context, feeId string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.GetDiscountsHoldersKey(feeId))
}

func (k Keeper) GetFeesDiscountsHoldersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.DiscountHolderKeyPrefix)
}

func (k Keeper) DiscountHolderExists(ctx sdk.Context, feeId string, discountId uint64, holder sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetDiscountHolderKey(feeId, discountId, holder))
}

func (k Keeper) GetFirstDiscountHeld(ctx sdk.Context, feeId string,
	holder sdk.AccAddress) (discountId uint64, holdsDiscount bool, err sdk.Error) {
	// Get specified fee
	fee, err := k.GetFee(ctx, feeId)
	if err != nil {
		return 0, false, err
	}

	// Find first discount
	for _, discount := range fee.Content.Discounts {
		if k.DiscountHolderExists(ctx, feeId, discount.Id, holder) {
			return discount.Id, true, nil
		}
	}

	// Not holding a discount is not considered and error
	return 0, false, nil
}

func (k Keeper) MustGetDiscountHolderByKey(ctx sdk.Context, key []byte) types.DiscountHolder {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("discount holder not found")
	}

	bz := store.Get(key)
	var discountHolder types.DiscountHolder
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &discountHolder)

	return discountHolder
}

func (k Keeper) SetDiscountHolder(ctx sdk.Context, discountHolder types.DiscountHolder) {
	store := ctx.KVStore(k.storeKey)
	dc := discountHolder
	key := types.GetDiscountHolderKey(dc.FeeId, dc.DiscountId, dc.Holder)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(dc))
}
