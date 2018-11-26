package node

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

// SetNode The node param values
func (k Keeper) SetNode(ctx sdk.Context, key string, value string) {
	k.paramsKeeper.Setter().SetString(ctx, MakeNodeKey(key), value)
}

// GetNode The node param values
func (k Keeper) GetNode(ctx sdk.Context, key string) string {
	r, err := k.paramsKeeper.Getter().GetString(ctx, MakeNodeKey(key))
	if err != nil {
		panic(err)
	}
	return r
}

// MakeNodeKey Makes node key
func MakeNodeKey(key string) string {
	return "node/" + key
}
