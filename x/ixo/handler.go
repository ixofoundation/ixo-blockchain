package ixo

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// func NewHandler(k Keeper) sdk.Handler {
// 	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
// 		switch msg := msg.(type) {
// 		case AddEthWalletMsg:
// 			return handleAddEthWalletMsg(ctx, k, msg)
// 		default:
// 			return sdk.ErrUnknownRequest("No match for message type.").Result()
// 		}
// 	}
// }

func handleAddEthWalletMsg(ctx sdk.Context, k Keeper, msg AddEthWalletMsg) sdk.Result {
	k.SetEthAddress(ctx, msg.Data.Id, msg.Data.WalletAddress)

	return sdk.Result{
		Code: sdk.ABCICodeOK,
	}
}

// InitGenesis sets the fees onto the chain
func InitGenesis(ctx sdk.Context, ixoKeeper Keeper, genesisState GenesisState) error {

	address := genesisState.FoundationWallet
	if isEthAddress(address) {
		panic("ixo Foundation wallet is not set in genesis file")
	}
	ixoKeeper.SetEthAddress(ctx, KeyFoundationWalletID, address)

	address = genesisState.AuthContractAddress
	if isEthAddress(address) {
		panic("Auth Contract Address is not set in genesis file")
	}
	ixoKeeper.SetEthAddress(ctx, KeyAuthContractAddress, address)

	address = genesisState.IxoTokenContractAddress
	if isEthAddress(address) {
		panic("Ixo ERC20 Token Contract Address is not set in genesis file")
	}
	ixoKeeper.SetEthAddress(ctx, KeyIxoTokenContractAddress, address)

	address = genesisState.ProjectRegistryContractAddress
	if isEthAddress(address) {
		panic("Project Registry Contract Address is not set in genesis file")
	}
	ixoKeeper.SetEthAddress(ctx, KeyProjectRegistryContractAddress, address)

	address = genesisState.ProjectWalletAuthoriserAddress
	if isEthAddress(address) {
		panic("Project Wallet Authoriser Contract Address is not set in genesis file")
	}
	ixoKeeper.SetEthAddress(ctx, KeyProjectWalletAuthoriserContractAddress, address)
	
	return nil
}

func isEthAddress(address string) bool {
	return (strings.HasPrefix(address, "0x") && len(address) != 42)
}
