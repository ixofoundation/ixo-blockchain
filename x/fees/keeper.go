package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

// Keeper manages global parameter store
type Keeper struct {
	cdc          *wire.Codec
	paramsKeeper params.Keeper
}

// NewKeeper constructs a new Keeper
func NewKeeper(cdc *wire.Codec, paramsKeeper params.Keeper) Keeper {
	return Keeper{
		cdc:          cdc,
		paramsKeeper: paramsKeeper,
	}
}

// InitKeeper constructs a new Keeper with initial parameters
func InitKeeper(cdc *wire.Codec, paramsKeeper params.Keeper) Keeper {
	k := NewKeeper(cdc, paramsKeeper)

	return k
}

// set the percentage value
func (k Keeper) SetRat(ctx sdk.Context, key string, value sdk.Rat) {
	k.paramsKeeper.Setter().SetRat(ctx, MakeFeeKey(key), value)
}

// set the amount value
func (k Keeper) SetInt64(ctx sdk.Context, key string, value int64) {
	k.paramsKeeper.Setter().SetInt64(ctx, MakeFeeKey(key), value)
}

// set the percentage value
func (k Keeper) GetRat(ctx sdk.Context, key string) sdk.Rat {
	r, err := k.paramsKeeper.Getter().GetRat(ctx, MakeFeeKey(key))
	if err != nil {
		panic(err)
	}
	return r
}

// set the amount value
func (k Keeper) GetInt64(ctx sdk.Context, key string) int64 {
	i, err := k.paramsKeeper.Getter().GetInt64(ctx, MakeFeeKey(key))
	if err != nil {
		panic(err)
	}
	return i
}

func MakeFeeKey(key string) string {
	return "fee/" + key
}
