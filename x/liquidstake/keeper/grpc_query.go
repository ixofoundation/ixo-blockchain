package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

// Querier is a thin facade over Keeper for the gRPC server. Wrapping the
// keeper avoids method-name collisions with the embedded Keeper's helpers.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// ModuleParams returns the current global module parameters.
func (k Querier) ModuleParams(c context.Context, _ *types.QueryModuleParamsRequest) (*types.QueryModuleParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryModuleParamsResponse{ModuleParams: k.GetModuleParams(ctx)}, nil
}

// Pool returns a single pool by id.
func (k Querier) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if err := types.ValidatePoolID(req.PoolId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.GetPool(ctx, req.PoolId)
	if !found {
		return nil, status.Error(codes.NotFound, types.ErrPoolNotFound.Wrap(req.PoolId).Error())
	}
	return &types.QueryPoolResponse{Pool: pool}, nil
}

// Pools returns every registered pool, paginated.
func (k Querier) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	poolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)

	var pools []types.Pool
	pageRes, err := query.Paginate(poolStore, req.Pagination, func(_ []byte, value []byte) error {
		var p types.Pool
		if err := k.cdc.Unmarshal(value, &p); err != nil {
			return err
		}
		pools = append(pools, p)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryPoolsResponse{Pools: pools, Pagination: pageRes}, nil
}

// LiquidValidators returns the per-pool liquid validators with state.
func (k Querier) LiquidValidators(c context.Context, req *types.QueryLiquidValidatorsRequest) (*types.QueryLiquidValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if err := types.ValidatePoolID(req.PoolId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.GetPool(ctx, req.PoolId)
	if !found {
		return nil, status.Error(codes.NotFound, types.ErrPoolNotFound.Wrap(req.PoolId).Error())
	}
	return &types.QueryLiquidValidatorsResponse{LiquidValidators: k.GetAllLiquidValidatorStatesForPool(ctx, pool)}, nil
}

// States returns the per-pool NetAmountState (rates, supplies, balances).
func (k Querier) States(c context.Context, req *types.QueryStatesRequest) (*types.QueryStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if err := types.ValidatePoolID(req.PoolId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.GetPool(ctx, req.PoolId)
	if !found {
		return nil, status.Error(codes.NotFound, types.ErrPoolNotFound.Wrap(req.PoolId).Error())
	}
	return &types.QueryStatesResponse{NetAmountState: k.GetNetAmountStateForPool(ctx, pool)}, nil
}
