package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func (k Keeper) GetBondIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.BondsKeyPrefix)
}

func (k Keeper) GetNumberOfBonds(ctx sdk.Context) sdk.Int {
	var count sdk.Int
	iterator := k.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &bond)
		count = count.AddRaw(1)
	}
	return count
}

func (k Keeper) GetBond(ctx sdk.Context, bondDid ixo.Did) (bond types.Bond, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !k.BondExists(ctx, bondDid) {
		return
	}
	bz := store.Get(types.GetBondKey(bondDid))
	k.cdc.MustUnmarshalBinaryBare(bz, &bond)
	return bond, true
}

func (k Keeper) MustGetBond(ctx sdk.Context, bondDid ixo.Did) types.Bond {
	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		panic(fmt.Sprintf("bond '%s' not found\n", bondDid))
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

func (k Keeper) BondExists(ctx sdk.Context, bondDid ixo.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetBondKey(bondDid))
}

func (k Keeper) SetBond(ctx sdk.Context, bondDid ixo.Did, bond types.Bond) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBondKey(bondDid), k.cdc.MustMarshalBinaryBare(bond))
}

func (k Keeper) GetReserveBalances(ctx sdk.Context, bondDid ixo.Did) sdk.Coins {
	// TODO: investigate ways to prevent reserve address from being reused since this affects calculations
	bond := k.MustGetBond(ctx, bondDid)
	return k.CoinKeeper.GetCoins(ctx, bond.ReserveAddress)
}

func (k Keeper) GetSupplyAdjustedForBuy(ctx sdk.Context, bondDid ixo.Did) sdk.Coin {
	bond := k.MustGetBond(ctx, bondDid)
	batch := k.MustGetBatch(ctx, bondDid)
	supply := bond.CurrentSupply
	return supply.Add(batch.TotalBuyAmount)
}

func (k Keeper) GetSupplyAdjustedForSell(ctx sdk.Context, bondDid ixo.Did) sdk.Coin {
	bond := k.MustGetBond(ctx, bondDid)
	batch := k.MustGetBatch(ctx, bondDid)
	supply := bond.CurrentSupply
	return supply.Sub(batch.TotalSellAmount)
}

func (k Keeper) SetCurrentSupply(ctx sdk.Context, bondDid ixo.Did, currentSupply sdk.Coin) {
	if currentSupply.IsNegative() {
		panic("current supply cannot be negative")
	}
	bond := k.MustGetBond(ctx, bondDid)
	bond.CurrentSupply = currentSupply
	k.SetBond(ctx, bondDid, bond)
}
