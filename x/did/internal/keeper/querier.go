package keeper

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryDidDoc     = "queryDidDoc"
	QueryAllDids    = "queryAllDids"
	QueryAllDidDocs = "queryAllDidDocs"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryDidDoc:
			return queryDidDoc(ctx, path[1:], k, legacyQuerierCdc)
		case QueryAllDids:
			return queryAllDids(ctx, k)
		case QueryAllDidDocs:
			return queryAllDidDocs(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Unknown did query endpoint")
		}
	}
}

func queryDidDoc(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	didDoc, err := k.GetDidDoc(ctx, path[0])
	if err != nil {
		return nil, err
	}

	//res, errRes := codec.MarshalJSONIndent(k.cdc, didDoc)
	res, errRes := codec.MarshalJSONIndent(legacyQuerierCdc, didDoc)
	if errRes != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal data %s", errRes.Error())
	}

	return res, nil
}

func queryAllDids(ctx sdk.Context, k Keeper) ([]byte, error) {
	allDids := k.GetAddDids(ctx)

	res, errRes := json.Marshal(allDids)
	if errRes != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal data %s", errRes.Error())
	}

	return res, nil
}

func queryAllDidDocs(ctx sdk.Context, k Keeper) ([]byte, error) {
	var didDocs []exported.DidDoc
	didDocs = k.GetAllDidDocs(ctx)

	res, errRes := json.Marshal(didDocs)
	if errRes != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal data %s", errRes.Error())
	}

	return res, nil
}
