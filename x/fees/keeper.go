package fees

import "github.com/ixofoundation/ixo-cosmos/x/params"

// Keeper manages global parameter store
type Keeper struct {
	paramsKeeper params.Keeper
}

// NewKeeper constructs a new Keeper
func NewKeeper(paramsKeeper params.Keeper) Keeper {
	return Keeper{
		paramsKeeper: paramsKeeper,
	}
}

// InitKeeper constructs a new Keeper with initial parameters
func InitKeeper(paramsKeeper params.Keeper) Keeper {
	k := NewKeeper(paramsKeeper)

	return k
}
