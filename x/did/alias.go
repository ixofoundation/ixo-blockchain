package did

import (
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	DefaultCodespace = types.DefaultCodespace
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	Did    = exported.Did
	DidDoc = exported.DidDoc
	IxoDid = exported.IxoDid
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

	IsValidDid      = types.IsValidDid
	DidToAddr       = types.DidToAddr
	StringToAddr    = types.StringToAddr
	UnmarshalIxoDid = types.UnmarshalIxoDid
)
