package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryOracles = "queryOracles"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryOracles:
			return queryOracles(ctx, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown oracles query endpoint")
		}
	}
}

func queryOracles(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	oracles := k.GetOracles(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, oracles)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}
