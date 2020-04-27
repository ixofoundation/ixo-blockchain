package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/oracles/internal/types"
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

// GetOracles returns the list of registered oracles
func (k Keeper) GetOracles(ctx sdk.Context) (oracles types.Oracles) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.OracleKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var oracle types.Oracle
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &oracle)
		oracles = append(oracles, oracle)
	}

	return oracles
}

// MustGetOracle returns the oracle if it exists
func (k Keeper) MustGetOracle(ctx sdk.Context, oracleDid ixo.Did) types.Oracle {

	store := ctx.KVStore(k.storeKey)
	if !k.OracleExists(ctx, oracleDid) {
		panic(fmt.Sprintf("oracle not found for %s\n", oracleDid))
	}

	bz := store.Get(types.GetOraclePrefixKey(oracleDid))
	var oracle types.Oracle
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &oracle)

	return oracle
}

// OracleExists checks if an oracle exists
func (k Keeper) OracleExists(ctx sdk.Context, oracleDid ixo.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetOraclePrefixKey(oracleDid))
}

// SetOracle registers an oracle
func (k Keeper) SetOracle(ctx sdk.Context, oracle types.Oracle) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOraclePrefixKey(oracle.OracleDid)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(oracle))
}
