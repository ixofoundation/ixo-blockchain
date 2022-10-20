package entity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	//save did docs to the store
	// for i, iid := range gs.IidDocs {
	// 	k.SetDidDocument(ctx, []byte(iid.Id), iid)
	// 	metadata := gs.IidMeta[i]
	// 	k.SetDidMetadata(ctx, []byte(iid.Id), metadata)
	// }

	// Initialise params
	// k.SetParams(ctx, gs.Params)

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// dds := k.(ctx)
	// var metas []types.IidMetadata
	// for _, iid := range dds {
	// 	meta, _ := k.GetDidMetadata(ctx, []byte(iid.Id))
	// 	metas = append(metas, meta)
	// }

	params := k.GetParams(ctx)

	return &types.GenesisState{
		Params: params,
	}
}
