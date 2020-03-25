package contracts

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/types"
)

func InitGenesis(ctx sdk.Context, contractKeeper keeper.Keeper, genesisState types.GenesisState) []abciTypes.ValidatorUpdate {

	address := genesisState.FoundationWallet
	if isEthAddress(address) {
		panic("ixo Foundation wallet is not set in genesis file")
	}

	contractKeeper.SetContract(ctx, types.KeyFoundationWallet, address)
	address = genesisState.AuthContractAddress
	if isEthAddress(address) {
		panic("Auth Contract Address is not set in genesis file")
	}
	contractKeeper.SetContract(ctx, types.KeyAuthContractAddress, address)

	address = genesisState.IxoTokenContractAddress
	if isEthAddress(address) {
		panic("Ixo ERC20 Token Contract Address is not set in genesis file")
	}

	contractKeeper.SetContract(ctx, types.KeyIxoTokenContractAddress, address)
	address = genesisState.ProjectRegistryContractAddress
	if isEthAddress(address) {
		panic("Project Registry Contract Address is not set in genesis file")
	}

	contractKeeper.SetContract(ctx, types.KeyProjectRegistryContractAddress, address)
	address = genesisState.ProjectWalletAuthoriserAddress
	if isEthAddress(address) {
		panic("Project Wallet Authoriser Contract Address is not set in genesis file")
	}

	contractKeeper.SetContract(ctx, types.KeyProjectWalletAuthoriserContractAddress, address)

	return nil
}

func isEthAddress(address string) bool {
	return strings.HasPrefix(address, "0x") && len(address) != 42
}

func WriteGenesis(ctx sdk.Context, keeper keeper.Keeper) []*types.GenesisState {

	return []*types.GenesisState{
		{
			AuthContractAddress:            keeper.GetContract(ctx, types.KeyAuthContractAddress),
			FoundationWallet:               keeper.GetContract(ctx, types.KeyFoundationWallet),
			IxoTokenContractAddress:        keeper.GetContract(ctx, types.KeyIxoTokenContractAddress),
			ProjectRegistryContractAddress: keeper.GetContract(ctx, types.KeyProjectRegistryContractAddress),
			ProjectWalletAuthoriserAddress: keeper.GetContract(ctx, types.KeyProjectWalletAuthoriserContractAddress),
		},
	}
}

func DefaultGenesis() types.GenesisState {

	return types.GenesisState{
		IxoTokenContractAddress:        "Enter ETH Ixo Token contract address",
		AuthContractAddress:            "Enter ETH auth contract address",
		ProjectWalletAuthoriserAddress: "Enter ETH project wallet authoriser contract address",
		FoundationWallet:               "Enter ETH wallet address to accumulate foundations tokens",
		ProjectRegistryContractAddress: "Enter ETH project registry contract address",
	}
}
