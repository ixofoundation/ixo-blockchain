package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis init for the node config
func InitGenesis(ctx sdk.Context, keeper Keeper, data []*GenesisState) error {
	for _, w := range data {
		keeper.SetNode(ctx, KeyNodeID, w.Did)
		keeper.SetNode(ctx, KeyEthWallet, w.EthWallet)
	}
	return nil
}

// WriteGenesis writes state to genesis
func WriteGenesis(ctx sdk.Context, keeper Keeper) []*GenesisState {

	return []*GenesisState{
		&GenesisState{

			Did:       keeper.GetNode(ctx, KeyNodeID),
			EthWallet: keeper.GetNode(ctx, KeyEthWallet),
		},
	}
}

// DefaultGenesis default config for genesis
func DefaultGenesis(did string, ethWallet string) *GenesisState {
	return &GenesisState{
		Did:       did,
		EthWallet: ethWallet,
	}
}
