package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ixofoundation/ixo-blockchain/v3/x/claims/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) Collection(c context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	collection, err := k.GetCollection(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectionResponse{Collection: collection}, nil
}

func (k Keeper) CollectionList(c context.Context, req *types.QueryCollectionListRequest) (*types.QueryCollectionListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var collections []types.Collection
	ctx := sdk.UnwrapSDKContext(c)
	collectionsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.CollectionKey)

	pageRes, err := query.Paginate(collectionsStore, req.Pagination, func(key []byte, value []byte) error {
		var collection types.Collection
		if err := k.cdc.Unmarshal(value, &collection); err != nil {
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

func (k Keeper) Claim(c context.Context, req *types.QueryClaimRequest) (*types.QueryClaimResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	claim, err := k.GetClaim(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryClaimResponse{Claim: claim}, nil
}

func (k Keeper) ClaimList(c context.Context, req *types.QueryClaimListRequest) (*types.QueryClaimListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claims []types.Claim
	ctx := sdk.UnwrapSDKContext(c)
	claimsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ClaimKey)

	pageRes, err := query.Paginate(claimsStore, req.Pagination, func(key []byte, value []byte) error {
		var claim types.Claim
		if err := k.cdc.Unmarshal(value, &claim); err != nil {
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

func (k Keeper) Dispute(c context.Context, req *types.QueryDisputeRequest) (*types.QueryDisputeResponse, error) {
	if req == nil || req.Proof == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	Dispute, err := k.GetDispute(ctx, req.Proof)
	if err != nil {
		return nil, err
	}

	return &types.QueryDisputeResponse{Dispute: Dispute}, nil
}

func (k Keeper) DisputeList(c context.Context, req *types.QueryDisputeListRequest) (*types.QueryDisputeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var disputes []types.Dispute
	ctx := sdk.UnwrapSDKContext(c)
	disputesStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.DisputeKey)

	pageRes, err := query.Paginate(disputesStore, req.Pagination, func(key []byte, value []byte) error {
		var dispute types.Dispute
		if err := k.cdc.Unmarshal(value, &dispute); err != nil {
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

func (k Keeper) Intent(c context.Context, req *types.QueryIntentRequest) (*types.QueryIntentResponse, error) {
	if req == nil || req.Id == "" || req.AgentAddress == "" || req.CollectionId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	Dispute, err := k.GetIntent(ctx, req.AgentAddress, req.CollectionId, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryIntentResponse{Intent: Dispute}, nil
}

func (k Keeper) IntentList(c context.Context, req *types.QueryIntentListRequest) (*types.QueryIntentListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var intents []types.Intent
	ctx := sdk.UnwrapSDKContext(c)
	disputesStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.IntentKey)

	pageRes, err := query.Paginate(disputesStore, req.Pagination, func(key []byte, value []byte) error {
		var intent types.Intent
		if err := k.cdc.Unmarshal(value, &intent); err != nil {
			return err
		}

		intents = append(intents, intent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryIntentListResponse{Intents: intents, Pagination: pageRes}, nil
}
