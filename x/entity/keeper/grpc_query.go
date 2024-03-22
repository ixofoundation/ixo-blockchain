package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ixofoundation/ixo-blockchain/v3/x/entity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/entity keeper providing gRPC method
// handlers.
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: q.Keeper.GetParams(ctx)}, nil
}

func (q Querier) Entity(c context.Context, req *types.QueryEntityRequest) (*types.QueryEntityResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	doc, entity, err := q.Keeper.ResolveEntity(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryEntityResponse{IidDocument: doc, Entity: entity}, nil
}

func (q Querier) EntityMetaData(c context.Context, req *types.QueryEntityMetadataRequest) (*types.QueryEntityMetadataResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_, entity, err := q.Keeper.ResolveEntity(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryEntityMetadataResponse{Entity: entity}, nil
}

func (q Querier) EntityIidDocument(c context.Context, req *types.QueryEntityIidDocumentRequest) (*types.QueryEntityIidDocumentResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	doc, _, err := q.Keeper.ResolveEntity(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryEntityIidDocumentResponse{IidDocument: doc}, nil
}

func (q Querier) EntityVerified(c context.Context, req *types.QueryEntityVerifiedRequest) (*types.QueryEntityVerifiedResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_, entity, err := q.Keeper.ResolveEntity(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryEntityVerifiedResponse{EntityVerified: entity.EntityVerified}, nil
}

func (q Querier) EntityList(c context.Context, req *types.QueryEntityListRequest) (*types.QueryEntityListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var entities []types.Entity
	ctx := sdk.UnwrapSDKContext(c)
	entityStore := prefix.NewStore(ctx.KVStore(q.Keeper.storeKey), types.EntityKey)

	pageRes, err := query.Paginate(entityStore, req.Pagination, func(key []byte, value []byte) error {
		var entity types.Entity
		if err := q.Keeper.cdc.Unmarshal(value, &entity); err != nil {
			return err
		}

		entities = append(entities, entity)
		return nil
	})

	// pageRes, err := query.FilteredPaginate(entityStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
	// 	if accumulate {
	// 		var e types.EntityDoc
	// 		if err := q.Keeper.cdc.Unmarshal(value, &e); err != nil {
	// 			return false, err
	// 		}
	// 		entities = append(entities, e)
	// 	}
	// 	return true, nil
	// })

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEntityListResponse{Entities: entities, Pagination: pageRes}, nil
}
