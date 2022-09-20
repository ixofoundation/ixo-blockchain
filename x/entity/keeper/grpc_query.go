package keeper

import (
	"context"

	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) EntityDoc(c context.Context, req *types.QueryEntityDocRequest) (*types.QueryEntityDocResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)

	// storedDoc, err := k.GetProjectDoc(ctx, req.ProjectDid)
	// if err != nil {
	// 	return nil, err
	// }

	return &types.QueryEntityDocResponse{}, nil
}

func (k Keeper) EntityList(c context.Context, req *types.QueryEntityListRequest) (*types.QueryEntityListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)

	// storedDoc, err := k.GetProjectDoc(ctx, req.ProjectDid)
	// if err != nil {
	// 	return nil, err
	// }

	return &types.QueryEntityListResponse{}, nil
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
