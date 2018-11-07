package contracts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

// Keeper manages contract params
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

// SetContract The contract param values
func (k Keeper) SetContract(ctx sdk.Context, key string, value string) {
	k.paramsKeeper.Setter().SetString(ctx, MakeContractKey(key), value)
}

// GetContract The contract param values
func (k Keeper) GetContract(ctx sdk.Context, key string) string {
	r, err := k.paramsKeeper.Getter().GetString(ctx, MakeContractKey(key))
	if err != nil {
		panic(err)
	}
	return r
}

// MakeContractKey Makes contract key
func MakeContractKey(key string) string {
	return "contract/" + key
}
