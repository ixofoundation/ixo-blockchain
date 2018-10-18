package ixo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

const KeyFoundationWalletID = "FOUNDATION"

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

// GetETHWallet returns the ETH addr for the ID.
func (k Keeper) GetETHWallet(ctx sdk.Context, id string) string {
	wallet, err := k.paramsKeeper.Getter().GetString(ctx, k.getKeyForEthWallet(id))
	if err != nil {
		panic(err)
	}
	return wallet
}

// SetETHWallet sets the ETH addr of for the ID
func (k Keeper) SetETHWallet(ctx sdk.Context, id string, walletAddr string) {
	setter := k.paramsKeeper.Setter()
	setter.SetString(ctx, k.getKeyForEthWallet(id), walletAddr)
}

func (k Keeper) getKeyForEthWallet(id string) string {
	if len(id) == 0 {
		panic("EthWalletID is empty")
	}
	return "ETH-" + id
}
