package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/ixofoundation/ixo-cosmos/types"
)

const (
	QueryFiat = "queryFiat"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abciTypes.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryFiat:
			return queryFiat(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown fiatFactory query endpoint")
		}
	}
}

// query Fiat handler
func queryFiat(ctx sdk.Context, path []string, keeper Keeper) ([]byte, sdk.Error) {
	pegHashHex, err := types.GetPegHashHex(path[1])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse the pegHash %s", err))
	}

	fiatPeg, err := keeper.GetFiatPeg(ctx, pegHashHex)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("%s", err))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, fiatPeg)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err.Error()))
	}
	return res, nil
}

