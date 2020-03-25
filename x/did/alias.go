package did

import (
	"github.com/ixofoundation/ixo-cosmos/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/did/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	DefaultCodeSpace = types.DefaultCodeSpace
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	ErrorInvalidDid = types.ErrorInvalidDid
)
