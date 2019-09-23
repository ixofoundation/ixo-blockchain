package keeper

import (
	"encoding/json"
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"
)

const (
	QueryProjectDoc     = "queryProjectDoc"
	QueryProjectAccount = "queryProjectAccount"
	QueryProjectTx      = "queryProjectTx"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req types.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryProjectDoc:
			return queryProjectDoc(ctx, path[1:], k)
		case QueryProjectAccount:
			return queryProjectAccount(ctx, path[1:], k)
		case QueryProjectTx:
			return queryProjectTx(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown project query endpoint")
		}
	}
}

func queryProjectDoc(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	storedDoc, err := k.GetProjectDoc(ctx, path[0])
	if err != nil {
		return nil, err
	}
	
	res, errRes := codec.MarshalJSONIndent(k.cdc, storedDoc)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err))
	}
	
	return res, nil
}

func queryProjectAccount(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	
	resp := k.GetAccountMap(ctx, path[0])
	res, err := json.Marshal(resp)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err.Error()))
	}
	
	return res, nil
}

func queryProjectTx(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	info, err := k.GetProjectWithdrawalTransactions(ctx, path[0])
	if err != nil {
		return nil, err
	}
	
	res, errRes := codec.MarshalJSONIndent(k.cdc, info)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err.Error()))
	}
	
	return res, nil
}
