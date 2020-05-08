package oracles

import (
	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	MintCap     = types.MintCap
	BurnCap     = types.BurnCap
	TransferCap = types.TransferCap

	DefaultCodeSpace = types.DefaultCodeSpace
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	Oracle          = types.Oracle
	Oracles         = types.Oracles
	OracleTokenCap  = types.OracleTokenCap
	OracleTokenCaps = types.OracleTokenCaps
	TokenCap        = types.TokenCap
	TokenCaps       = types.TokenCaps
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
