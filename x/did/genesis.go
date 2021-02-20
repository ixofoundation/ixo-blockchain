package did

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	// Initialise did docs
	for _, d := range gs.Diddocs {
		dd, ok := d.GetCachedValue().(exported.DidDoc)
		if !ok {
			panic("expected DidDoc")
		}

		k.AddDidDoc(ctx, dd)
	}

	//for _, d := range gs.Diddocs {
	//	keeper.AddDidDoc(ctx, d)
	//}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	dds := k.GetAllDidDocs(ctx)
	diddocs := make([]*codectypes.Any, len(dds))
	for i, dd := range dds {
		msg, ok := dd.(proto.Message)
		if !ok {
			panic(fmt.Errorf("cannot proto marshal %T", dd))
		}
		any, err := codectypes.NewAnyWithValue(msg)
		if err != nil {
			panic(err)
		}
		diddocs[i] = any
	}

	return types.GenesisState{
		Diddocs: diddocs, //keeper.GetAllDidDocs(ctx)
	}
}
