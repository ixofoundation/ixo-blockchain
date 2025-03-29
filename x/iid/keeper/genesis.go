package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/iid/types"
)

// InitGenesis initializes the x/iid module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	for _, iid := range gs.IidDocs {
		k.SetDidDocument(ctx, []byte(iid.Id), iid)
	}
}

// ExportGenesis returns the x/iid module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	iidDocs := k.GetAllDidDocuments(ctx)

	return &types.GenesisState{
		IidDocs: iidDocs,
	}
}
