package bonddoc

import (
	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
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

	MsgCreateBond       = types.MsgCreateBond
	MsgUpdateBondStatus = types.MsgUpdateBondStatus

	StoredBondDoc = types.StoredBondDoc
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
)
