package iid

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	for _, iid := range gs.IidDocs {
		k.SetDidDocument(ctx, []byte(iid.Id), iid)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	iidDocs := k.GetAllDidDocuments(ctx)

	return &types.GenesisState{
		IidDocs: iidDocs,
	}
}
