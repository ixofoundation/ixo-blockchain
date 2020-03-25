package node

import (
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/node/internal/keeper"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) []abci.ValidatorUpdate {

	for _, w := range data.ETHGenesisStates {
		keeper.SetNode(ctx, KeyNodeID, w.Did)
		keeper.SetNode(ctx, KeyEthWallet, w.EthWallet)
	}

	return []abci.ValidatorUpdate{}
}

func WriteGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	did, _ := keeper.GetNode(ctx, KeyNodeID)
	wallet, _ := keeper.GetNode(ctx, KeyEthWallet)

	return GenesisState{
		ETHGenesisStates: []ETHGenesisState{{Did: did, EthWallet: wallet}},
	}
}

func DefaultGenesis() *GenesisState {
	var secret string
	_, secret, err := server.GenerateCoinKey()
	if err != nil {
		panic(err)
	}

	ethereumAddr, err := ixo.IxoAppGenEthWallet()
	ethWallet := ethereumAddr

	if err != nil {
		panic(err)
	}

	sovrinDid := sovrin.FromMnemonic(secret)
	did := "did:ixo:" + sovrinDid.Did

	return &GenesisState{[]ETHGenesisState{{Did: did, EthWallet: ethWallet}}}
}
