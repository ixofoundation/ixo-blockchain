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

// SetFiatAccount:
func (k Keeper) SetFiatAccount(ctx sdk.Context, fiatAccount types.FiatAccount) {
	store := ctx.KVStore(k.storeKey)

	fiatAccountKey := fiatTypes.FiatAccountStoreKey(fiatAccount.GetAddress())
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(fiatAccount)
	store.Set(fiatAccountKey, bz)
}

// returns fiat account by address
func (k Keeper) GetFiatAccount(ctx sdk.Context, address sdk.AccAddress) (fiatAccount types.FiatAccount, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	fiatAccountKey := fiatTypes.FiatAccountStoreKey(address)
	bz := store.Get(fiatAccountKey)
	if bz == nil {
		return nil, fiatTypes.ErrInvalidPegHash(fiatTypes.DefaultCodeSpace)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fiatAccount)
	return fiatAccount, nil
}

func (k Keeper) IterateFiatAccounts(ctx sdk.Context, handler func(fiatAccount types.FiatAccount) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, fiatTypes.FiatAccountKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var fiatAccount types.FiatAccount
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &fiatAccount)
		if handler(fiatAccount) {
			break
		}
	}
}

// get all fiat accounts => []FiatAccounts from store
func (k Keeper) GetFiatAccounts(ctx sdk.Context) (fiatAccounts []types.FiatAccount) {
	k.IterateFiatAccounts(ctx, func(fiatAccount types.FiatAccount) (stop bool) {
		fiatAccounts = append(fiatAccounts, fiatAccount)
		return false
	},
	)
	return
}

func (k Keeper) IssueFiats(ctx sdk.Context, issueFiat fiatTypes.IssueFiat) sdk.Error {
	return nil
}

func (k Keeper) RedeemFiats(ctx sdk.Context, redeemFiat fiatTypes.RedeemFiat) sdk.Error {
	return nil
}

func (k Keeper) SendFiats(ctx sdk.Context, sendFiat fiatTypes.SendFiat) sdk.Error {
	return nil
}
