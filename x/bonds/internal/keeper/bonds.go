package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
)

func (k Keeper) GetBondIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.BondsKeyPrefix)
}

func (k Keeper) GetBond(ctx sdk.Context, token string) (bond types.Bond, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !k.BondExists(ctx, token) {
		return
	}
	bz := store.Get(types.GetBondKey(token))
	k.cdc.MustUnmarshalBinaryBare(bz, &bond)
	return bond, true
}

func (k Keeper) MustGetBond(ctx sdk.Context, token string) types.Bond {
	bond, found := k.GetBond(ctx, token)
	if !found {
		panic(fmt.Sprintf("bond '%s' not found\n", token))
	}
	return bond
}

func (k Keeper) MustGetBondByKey(ctx sdk.Context, key []byte) types.Bond {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("bond not found")
	}
	bz := store.Get(key)
	var bond types.Bond
	k.cdc.MustUnmarshalBinaryBare(bz, &bond)
	return bond
}

func (k Keeper) BondExists(ctx sdk.Context, token string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetBondKey(token))
}

func (k Keeper) SetBond(ctx sdk.Context, token string, bond types.Bond) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBondKey(token), k.cdc.MustMarshalBinaryBare(bond))
}

func (k Keeper) GetReserveBalances(ctx sdk.Context, token string) sdk.Coins {
	// TODO: investigate ways to prevent reserve address from being reused since this affects calculations
	bond := k.MustGetBond(ctx, token)
	return k.CoinKeeper.GetCoins(ctx, bond.ReserveAddress)
}

func (k Keeper) GetSupplyAdjustedForBuy(ctx sdk.Context, token string) sdk.Coin {
	bond := k.MustGetBond(ctx, token)
	batch := k.MustGetBatch(ctx, token)
	supply := bond.CurrentSupply
	return supply.Add(batch.TotalBuyAmount)
}

func (k Keeper) GetSupplyAdjustedForSell(ctx sdk.Context, token string) sdk.Coin {
	bond := k.MustGetBond(ctx, token)
	batch := k.MustGetBatch(ctx, token)
	supply := bond.CurrentSupply
	return supply.Sub(batch.TotalSellAmount)
}

func (k Keeper) SetCurrentSupply(ctx sdk.Context, token string, currentSupply sdk.Coin) {
	if currentSupply.IsNegative() {
		panic("current supply cannot be negative")
	}
	bond := k.MustGetBond(ctx, token)
	bond.CurrentSupply = currentSupply
	k.SetBond(ctx, token, bond)
}
