package keeper

import (
	"context"

	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TokenDoc(c context.Context, req *types.QueryTokenDocRequest) (*types.QueryTokenDocResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)

	// storedDoc, err := k.GetProjectDoc(ctx, req.ProjectDid)
	// if err != nil {
	// 	return nil, err
	// }

	return &types.QueryTokenDocResponse{}, nil
}

func (k Keeper) TokenList(c context.Context, req *types.QueryTokenListRequest) (*types.QueryTokenListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)

	// wasmkeeper.Querier(&k.WasmKeeper).SmartContractState()
	// k.WasmKeeper.Execute()

	// storedDoc, err := k.GetProjectDoc(ctx, req.ProjectDid)
	// if err != nil {
	// 	return nil, err
	// }

	return &types.QueryTokenListResponse{}, nil
}

func (k Keeper) TokenConfig(c context.Context, req *types.QueryTokenConfigRequest) (*types.QueryTokenConfigResponse, error) {
	return &types.QueryTokenConfigResponse{}, nil
}

// func (k Keeper) ProjectAccounts(c context.Context, req *types.QueryProjectAccountsRequest) (*types.QueryProjectAccountsResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)

// 	resp := k.GetAccountMap(ctx, req.ProjectDid)

// 	return &types.QueryProjectAccountsResponse{AccountMap: &resp}, nil
// }

// func (k Keeper) ProjectTx(c context.Context, req *types.QueryProjectTxRequest) (*types.QueryProjectTxResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)

// 	info, err := k.GetProjectWithdrawalTransactions(ctx, req.ProjectDid)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &types.QueryProjectTxResponse{Txs: &info}, nil
// }

// func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	params := k.GetParams(ctx)

// 	return &types.QueryParamsResponse{Params: params}, nil
// }
