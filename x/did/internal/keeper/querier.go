package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const (
	QueryDidDoc     = "queryDidDoc"
	QueryAllDids    = "queryAllDids"
	QueryAllDidDocs = "queryAllDidDocs"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryDidDoc:
			return queryDidDoc(ctx, path[1:], k)
		case QueryAllDids:
			return queryAllDids(ctx, k)
		case QueryAllDidDocs:
			return queryAllDidDocs(ctx, k)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown did query endpoint")
		}
	}
}

func queryDidDoc(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	didDoc, err := k.GetDidDoc(ctx, path[0])
	if err != nil {
		return nil, err
	}

	res, errRes := json.Marshal(didDoc)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", errRes))
	}

	return res, nil
}

func queryAllDids(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	allDids := k.GetAddDids(ctx)

	res, errRes := json.Marshal(allDids)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", errRes.Error()))
	}

	return res, nil
}

func queryAllDidDocs(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	var didDocs []ixo.DidDoc
	didDocs = k.GetAllDidDocs(ctx)

	res, errRes := json.Marshal(didDocs)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", errRes.Error()))
	}

	return res, nil
}
