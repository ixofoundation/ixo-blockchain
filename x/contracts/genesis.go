package contracts

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets the fees onto the chain
func InitGenesis(ctx sdk.Context, contractKeeper Keeper, genesisState GenesisState) error {

	address := genesisState.FoundationWallet
	if isEthAddress(address) {
		panic("ixo Foundation wallet is not set in genesis file")
	}
	contractKeeper.SetContract(ctx, KeyFoundationWallet, address)

	address = genesisState.AuthContractAddress
	if isEthAddress(address) {
		panic("Auth Contract Address is not set in genesis file")
	}
	contractKeeper.SetContract(ctx, KeyAuthContractAddress, address)

	address = genesisState.IxoTokenContractAddress
	if isEthAddress(address) {
		panic("Ixo ERC20 Token Contract Address is not set in genesis file")
	}
	contractKeeper.SetContract(ctx, KeyIxoTokenContractAddress, address)

	address = genesisState.ProjectRegistryContractAddress
	if isEthAddress(address) {
		panic("Project Registry Contract Address is not set in genesis file")
	}
	contractKeeper.SetContract(ctx, KeyProjectRegistryContractAddress, address)

	address = genesisState.ProjectWalletAuthoriserAddress
	if isEthAddress(address) {
		panic("Project Wallet Authoriser Contract Address is not set in genesis file")
	}
	contractKeeper.SetContract(ctx, KeyProjectWalletAuthoriserContractAddress, address)

	return nil
}

func isEthAddress(address string) bool {
	return (strings.HasPrefix(address, "0x") && len(address) != 42)
}

// WriteGenesis returns a GenesisState for a given context and keeper. The
// GenesisState will contain the pool, params, validators, and bonds found in
// the keeper.
func WriteGenesis(ctx sdk.Context, keeper Keeper) []*GenesisState {

	return []*GenesisState{
		&GenesisState{
			AuthContractAddress:            keeper.GetContract(ctx, KeyAuthContractAddress),
			FoundationWallet:               keeper.GetContract(ctx, KeyFoundationWallet),
			IxoTokenContractAddress:        keeper.GetContract(ctx, KeyIxoTokenContractAddress),
			ProjectRegistryContractAddress: keeper.GetContract(ctx, KeyProjectRegistryContractAddress),
			ProjectWalletAuthoriserAddress: keeper.GetContract(ctx, KeyProjectWalletAuthoriserContractAddress),
		},
	}
}

// DefaultGenesis returns a the defaultGenesisState for a given context and keeper.
func DefaultGenesis() GenesisState {

	return GenesisState{
		FoundationWallet:               "Enter ETH wallet address to accumulate foundations tokens",
		AuthContractAddress:            "Enter ETH auth contract address",
		IxoTokenContractAddress:        "Enter ETH Ixo Token contract address",
		ProjectRegistryContractAddress: "Enter ETH project registry contract address",
		ProjectWalletAuthoriserAddress: "Enter ETH project wallet authoriser contract address",
	}
}
