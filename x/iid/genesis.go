package iid

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	//save did docs to the store

	go func() {
		for _, iid := range gs.IidDocs {
			k.SetDidDocument(ctx, []byte(iid.Id), iid)
		}
	}()

	go func() {
		for _, meta := range gs.IidMeta {
			k.SetDidMetadata(ctx, []byte(meta.Id), meta)
		}
	}()

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	dds := k.GetAllDidDocuments(ctx)
	var metas []types.IidMetadata
	for _, iid := range dds {
		meta, _ := k.GetDidMetadata(ctx, []byte(iid.Id))
		metas = append(metas, meta)
	}

	return &types.GenesisState{
		IidDocs: dds,
		IidMeta: metas,
	}
}
