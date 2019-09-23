package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

type Keeper struct {
	cdc          *codec.Codec
	paramsKeeper params.Keeper
}

func NewKeeper(cdc *codec.Codec, paramsKeeper params.Keeper) Keeper {
	return Keeper{
		cdc:          cdc,
		paramsKeeper: paramsKeeper,
	}
}

func (k Keeper) SetContract(ctx sdk.Context, key string, value string) {
	k.paramsKeeper.Setter().SetString(ctx, MakeContractKey(key), value)
}

func (k Keeper) GetContract(ctx sdk.Context, key string) string {
	r, err := k.paramsKeeper.Getter().GetString(ctx, MakeContractKey(key))
	if err != nil {
		panic(err)
	}
	return r
}

func MakeContractKey(key string) string {
	return "contract/" + key
}
