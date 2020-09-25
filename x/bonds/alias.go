package bonds

import (
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
)

//noinspection GoNameStartsWithPackageName
const (
	BondsMintBurnAccount       = types.BondsMintBurnAccount
	BatchesIntermediaryAccount = types.BatchesIntermediaryAccount
	BondsReserveAccount        = types.BondsReserveAccount

	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	RouterKey         = types.RouterKey
)

//noinspection GoNameStartsWithPackageName
var (
	// function aliases
	RegisterInvariants = keeper.RegisterInvariants
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	RegisterCodec      = types.RegisterCodec

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc            = types.ModuleCdc
	BondsKeyPrefix       = types.BondsKeyPrefix
	BatchesKeyPrefix     = types.BatchesKeyPrefix
	LastBatchesKeyPrefix = types.LastBatchesKeyPrefix
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
