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
	// function aliases
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	ErrorInvalidDid = types.ErrorInvalidDid
)
