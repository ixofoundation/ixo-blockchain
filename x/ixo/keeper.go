package ixo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

const KeyFoundationWalletID = "Foundation"
const KeyAuthContractAddress = "AuthContract"
const KeyIxoTokenContractAddress = "IxoTokenContract"
const KeyProjectRegistryContractAddress = "ProjectRegistryContract"

// IXO Keeper manages ixo params
type Keeper struct {
	cdc          *wire.Codec
	paramsKeeper params.Keeper
}

// NewKeeper constructs a new Keeper
func NewKeeper(cdc *wire.Codec, paramsKeeper params.Keeper) Keeper {
	return Keeper{
		cdc:          cdc,
		paramsKeeper: paramsKeeper,
	}
}

// GetConfigString returns the String for the ID.
func (k Keeper) GetConfigString(ctx sdk.Context, id string) string {
	value, err := k.paramsKeeper.Getter().GetString(ctx, k.getKeyForConfig(id))
	if err != nil {
		panic(err)
	}
	return value
}

// SetConfigString sets the value of for the ID
func (k Keeper) SetConfigString(ctx sdk.Context, id string, value string) {
	setter := k.paramsKeeper.Setter()
	setter.SetString(ctx, k.getKeyForConfig(id), value)
}

// GetEthAddress returns the ETH addr for the ID.
func (k Keeper) GetEthAddress(ctx sdk.Context, id string) string {
	wallet, err := k.paramsKeeper.Getter().GetString(ctx, k.getKeyForEthAddress(id))
	if err != nil {
		panic(err)
	}
	return wallet
}

// SetEthAddress sets the ETH addr of for the ID
func (k Keeper) SetEthAddress(ctx sdk.Context, id string, walletAddr string) {
	setter := k.paramsKeeper.Setter()
	setter.SetString(ctx, k.getKeyForEthAddress(id), walletAddr)
}

func (k Keeper) getKeyForEthAddress(id string) string {
	if len(id) == 0 {
		panic("EthWalletID is empty")
	}
	return "ETH-" + id
}

func (k Keeper) getKeyForConfig(id string) string {
	if len(id) == 0 {
		panic("Config id is empty")
	}
	return "CFG-" + id
}
