package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	abci "github.com/tendermint/tendermint/abci/types"
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

	res, errRes := codec.MarshalJSONIndent(k.cdc, didDoc)
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
	var didDocs []exported.DidDoc
	didDocs = k.GetAllDidDocs(ctx)

	res, errRes := json.Marshal(didDocs)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", errRes.Error()))
	}

	return res, nil
}
