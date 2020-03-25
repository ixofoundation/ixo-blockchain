package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/fees/internal/types"
)

const (
	QueryFees = "queryFees"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abciTypes.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryFees:
			return queryFees(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown fees query endpoint")
		}
	}
}

func queryFees(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	fees := make(map[string]int64)
	for _, feeKey := range types.AllFees {
		resDec := k.GetDec(ctx, feeKey)
		fees[feeKey] = resDec.RoundInt64()
	}

	res, err := codec.MarshalJSONIndent(k.cdc, fees)
	if err != nil {
		return nil, types.ErrorUnmarshalFees()
	}

	return res, nil
}
