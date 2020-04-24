package keeper

import (
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

// SetParams sets the list of registered oracles
func (k Keeper) SetOracle(ctx sdk.Context, oracle types.Oracle) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOraclePrefixKey(ixo.Did(oracle))
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(oracle))
}
