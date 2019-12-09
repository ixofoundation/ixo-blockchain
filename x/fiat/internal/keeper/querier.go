package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/codec"
)

const (
	QueryFiatAccount     = "queryFiatAccount"
	QueryAllFiatAccounts = "queryAllFiatAccounts"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abciTypes.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryFiatAccount:
			return queryFiatAccount(ctx, path[1:], k)
		case QueryAllFiatAccounts:
			return queryAllFiatAccounts(ctx, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown fiat query endpoint")
		}
	}
}

// query Fiat handler
func queryFiatAccount(ctx sdk.Context, path []string, keeper Keeper) ([]byte, sdk.Error) {

	addr, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse the address %s", err))
	}

	fiatAccount, err := keeper.GetFiatAccount(ctx, addr)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("%s", err))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, fiatAccount)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err.Error()))
	}
	return res, nil
}

func queryAllFiatAccounts(ctx sdk.Context, keeper Keeper) ([]byte, sdk.Error) {

	fiatAccounts := keeper.GetFiatAccounts(ctx)
	// if len(fiatAccounts) == 0 {
	// 	return []byte("No Fiat Accounts created."), nil
	// }

	res, err := codec.MarshalJSONIndent(keeper.cdc, fiatAccounts)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", err.Error()))
	}
	return res, nil
}
