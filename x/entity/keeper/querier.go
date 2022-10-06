package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryEntityDoc = "queryEntityDoc"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Unknown project query endpoint")
		}
	}
}

// func queryProjectDoc(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
// 	storedDoc, err := k.GetProjectDoc(ctx, path[0])
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, errRes := codec.MarshalJSONIndent(legacyQuerierCdc, storedDoc)
// 	if errRes != nil {
// 		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal iid: %s", err)
// 	}

// 	return res, nil
// }

// func queryProjectAccounts(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
// 	resp := k.GetAccountMap(ctx, path[0])
// 	res, err := json.Marshal(resp)
// 	if err != nil {
// 		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal iid %s", err)
// 	}

// 	return res, nil
// }

// func queryProjectTx(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
// 	info, err := k.GetProjectWithdrawalTransactions(ctx, path[0])
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, info)
// 	if err != nil {
// 		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal iid %s", err)
// 	}

// 	return res, nil
// }

// func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
// 	params := k.GetParams(ctx)

// 	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
// 	if err != nil {
// 		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal iid %s", err)
// 	}

// 	return res, nil
// }
