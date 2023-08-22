package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ixofoundation/ixo-blockchain/x/claims/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/claims keeper providing gRPC method
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

func (q Querier) Collection(c context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	collection, err := q.Keeper.GetCollection(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectionResponse{Collection: collection}, nil
}

func (q Querier) CollectionList(c context.Context, req *types.QueryCollectionListRequest) (*types.QueryCollectionListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var collections []types.Collection
	ctx := sdk.UnwrapSDKContext(c)
	collectionsStore := prefix.NewStore(ctx.KVStore(q.Keeper.storeKey), types.CollectionKey)

	pageRes, err := query.Paginate(collectionsStore, req.Pagination, func(key []byte, value []byte) error {
		var collection types.Collection
		if err := q.Keeper.cdc.Unmarshal(value, &collection); err != nil {
			return err
		}

		collections = append(collections, collection)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCollectionListResponse{Collections: collections, Pagination: pageRes}, nil
}

func (q Querier) Claim(c context.Context, req *types.QueryClaimRequest) (*types.QueryClaimResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	claim, err := q.Keeper.GetClaim(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryClaimResponse{Claim: claim}, nil
}

func (q Querier) ClaimList(c context.Context, req *types.QueryClaimListRequest) (*types.QueryClaimListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claims []types.Claim
	ctx := sdk.UnwrapSDKContext(c)
	claimsStore := prefix.NewStore(ctx.KVStore(q.Keeper.storeKey), types.ClaimKey)

	pageRes, err := query.Paginate(claimsStore, req.Pagination, func(key []byte, value []byte) error {
		var claim types.Claim
		if err := q.Keeper.cdc.Unmarshal(value, &claim); err != nil {
			return err
		}

		claims = append(claims, claim)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryClaimListResponse{Claims: claims, Pagination: pageRes}, nil
}

func (q Querier) Dispute(c context.Context, req *types.QueryDisputeRequest) (*types.QueryDisputeResponse, error) {
	if req == nil || req.Proof == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	Dispute, err := q.Keeper.GetDispute(ctx, req.Proof)
	if err != nil {
		return nil, err
	}

	return &types.QueryDisputeResponse{Dispute: Dispute}, nil
}

func (q Querier) DisputeList(c context.Context, req *types.QueryDisputeListRequest) (*types.QueryDisputeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var disputes []types.Dispute
	ctx := sdk.UnwrapSDKContext(c)
	disputesStore := prefix.NewStore(ctx.KVStore(q.Keeper.storeKey), types.DisputeKey)

	pageRes, err := query.Paginate(disputesStore, req.Pagination, func(key []byte, value []byte) error {
		var dispute types.Dispute
		if err := q.Keeper.cdc.Unmarshal(value, &dispute); err != nil {
			return err
		}

		disputes = append(disputes, dispute)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDisputeListResponse{Disputes: disputes, Pagination: pageRes}, nil
}
