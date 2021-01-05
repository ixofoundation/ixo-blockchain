package keeper
//
//import (
//	"github.com/cosmos/cosmos-sdk/codec"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	abci "github.com/tendermint/tendermint/abci/types"
//)
//
//const (
//	QueryOracles = "queryOracles"
//)
//
//func NewQuerier(k Keeper) sdk.Querier {
//	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
//		switch path[0] {
//		case QueryOracles:
//			return queryOracles(ctx, k)
//		default:
//			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown oracles query endpoint")
//		}
//	}
//}
//
//func queryOracles(ctx sdk.Context, k Keeper) ([]byte, error) {
//	oracles := k.GetOracles(ctx)
//
//	res, err := codec.MarshalJSONIndent(k.cdc, oracles)
//	if err != nil {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "failed to marshal JSON")
//	}
//	return res, nil
//}
