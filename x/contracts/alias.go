package contracts

import (
	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	
	DefaultCodeSpace = types.DefaultCodeSpace
	
	KeyIxoTokenContractAddress                = types.KeyIxoTokenContractAddress
	KeyProjectRegistryContractAddress         = types.KeyProjectRegistryContractAddress
	KeyAuthContractAddress                    = types.KeyAuthContractAddress
	KeyProjectWalletAuthoriserContractAddress = types.KeyProjectWalletAuthoriserContractAddress
	KeyFoundationWallet                       = types.KeyFoundationWallet
)

type (
	GenesisState = types.GenesisState
	Keeper       = keeper.Keeper
)

var (
	NewKeeper  = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier
	ModuleCdc  = types.ModuleCdc
)
