package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"
)

const (
	QueryBondDoc = "queryBondDoc"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req types.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryBondDoc:
			return queryBondDoc(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown bond query endpoint")
		}
	}
}

func queryBondDoc(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	storedDoc, err := k.GetBondDoc(ctx, path[0])
	if err != nil {
		return nil, err
	}

	res, errRes := codec.MarshalJSONIndent(k.cdc, storedDoc)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err))
	}

	return res, nil
}
