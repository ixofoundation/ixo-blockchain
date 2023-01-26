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

	// wasmkeeper.Querier(&k.WasmKeeper).SmartContractState()
	// k.WasmKeeper.Execute()

	// storedDoc, err := k.GetProjectDoc(ctx, req.ProjectDid)
	// if err != nil {
	// 	return nil, err
	// }

	return &types.QueryEntityListResponse{}, nil
}

func (k Keeper) EntityConfig(c context.Context, req *types.QueryEntityConfigRequest) (*types.QueryEntityConfigResponse, error) {
	return &types.QueryEntityConfigResponse{}, nil
}
