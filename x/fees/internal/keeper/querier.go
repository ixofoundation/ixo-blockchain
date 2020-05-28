package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryParams      = "queryParams"
	QueryFee         = "queryFee"
	QueryFeeContract = "queryFeeContract"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryParams:
			return queryParams(ctx, k)
		case QueryFee:
			return queryFee(ctx, path[1:], k)
		case QueryFeeContract:
			return queryFeeContract(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown fees query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}

func queryFee(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	feeId := path[0]

	fee, err := k.GetFee(ctx, feeId)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf(
			"fee '%s' does not exist", feeId))
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, fee)
	if err2 != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err2.Error()))
	}

	return res, nil
}

func queryFeeContract(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	feeContractId := path[0]

	feeContract, err := k.GetFeeContract(ctx, feeContractId)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf(
			"fee contract '%s' does not exist", feeContractId))
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, feeContract)
	if err2 != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err2.Error()))
	}

	return res, nil
}
