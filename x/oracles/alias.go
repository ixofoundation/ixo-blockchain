package oracles

import (
	"github.com/ixofoundation/ixo-cosmos/x/oracles/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/oracles/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	MintCap = types.MintCap
	BurnCap = types.BurnCap
	SendCap = types.SendCap

	DefaultCodeSpace = types.DefaultCodeSpace
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	Oracle                  = types.Oracle
	Oracles                 = types.Oracles
	OracleTokenCapability   = types.OracleTokenCap
	OracleTokenCapabilities = types.OracleTokenCaps
	TokenCapability         = types.TokenCap
	TokenCapabilities       = types.TokenCapabilities
)

var (
	// function aliases
	NewKeeper  = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
