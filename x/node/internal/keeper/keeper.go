package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/node/internal/types"
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

func (k Keeper) SetNode(ctx sdk.Context, key string, value string) {
	k.paramsKeeper.Setter().SetString(ctx, MakeNodeKey(key), value)
}

func (k Keeper) GetNode(ctx sdk.Context, key string) (string, sdk.Error) {
	r, err := k.paramsKeeper.Getter().GetString(ctx, MakeNodeKey(key))
	if err != nil {
		return "", types.ErrorInvalidQueryNode()
	}

	return r, nil
}

func MakeNodeKey(key string) string {
	return "node/" + key
}
