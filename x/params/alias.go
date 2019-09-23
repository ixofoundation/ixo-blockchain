package params

import (
	"github.com/ixofoundation/ixo-cosmos/x/params/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/params/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	
	DefaultCodeSpace = types.DefaultCodeSpace
)

type (
	Keeper = keeper.Keeper
)

var (
	NewKeeper = keeper.NewKeeper
	ModuleCdc = types.ModuleCdc
)
