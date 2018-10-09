package pool

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
)

// Keeper manages validator pool
type Keeper struct {
	key  sdk.StoreKey
	cdc  *wire.Codec
	pool ValidatorPool
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *wire.Codec, key sdk.StoreKey) Keeper {
	pool := NewValidatorPool()
	return Keeper{
		key,
		cdc,
		pool,
	}
}
