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

func InitKeeper(cdc *codec.Codec, paramsKeeper params.Keeper) Keeper {
	k := NewKeeper(cdc, paramsKeeper)

	return k
}

func (k Keeper) SetDec(ctx sdk.Context, key string, value sdk.Dec) {
	k.paramsKeeper.Setter().SetDec(ctx, MakeFeeKey(key), value)
}

func (k Keeper) GetDec(ctx sdk.Context, key string) sdk.Dec {
	dec, err := k.paramsKeeper.Getter().GetDec(ctx, MakeFeeKey(key))
	if err != nil {
		panic(err)
	}

	return dec
}

func MakeFeeKey(key string) string {
	return "fee/" + key
}
