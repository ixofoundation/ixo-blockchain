package module

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	// Initialise did docs
	for _, d := range gs.DidDocs {
		dd, ok := d.GetCachedValue().(exported.DidDoc)
		if !ok {
			panic("expected DidDoc")
		}

		k.AddDidDoc(ctx, dd)
	}

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
		DidDocs: diddocs,
	}
}
