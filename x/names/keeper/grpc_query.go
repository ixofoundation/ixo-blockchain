package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ixofoundation/ixo-blockchain/v6/x/names/types"
)

var _ types.QueryServer = Keeper{}

// Namespace returns a single Namespace by name.
func (k Keeper) Namespace(c context.Context, req *types.QueryNamespaceRequest) (*types.QueryNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	ctx := sdk.UnwrapSDKContext(c)
	ns, found := k.GetNamespace(ctx, req.Name)
	if !found {
		return nil, status.Error(codes.NotFound, errorsmod.Wrapf(types.ErrNamespaceNotFound, "namespace %q", req.Name).Error())
	}
	return &types.QueryNamespaceResponse{Namespace: ns}, nil
}

// Namespaces returns all Namespaces, paginated.
func (k Keeper) Namespaces(c context.Context, req *types.QueryNamespacesRequest) (*types.QueryNamespacesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	out, pageRes, err := k.PaginateNamespaces(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if out == nil {
		out = []types.Namespace{}
	}
	return &types.QueryNamespacesResponse{Namespaces: out, Pagination: pageRes}, nil
}

// ResolveName looks up an ACTIVE NameRecord. Callers may pass the display
// form; the server normalizes before lookup.
func (k Keeper) ResolveName(c context.Context, req *types.QueryResolveNameRequest) (*types.QueryResolveNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Namespace == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "namespace and name are required")
	}
	ctx := sdk.UnwrapSDKContext(c)
	normalized := types.NormalizeName(req.Name)
	record, found := k.GetNameRecord(ctx, req.Namespace, normalized)
	if !found || record.Status != types.NAME_STATUS_ACTIVE {
		return nil, status.Error(codes.NotFound, errorsmod.Wrapf(types.ErrNameNotFound, "%s/%s", req.Namespace, normalized).Error())
	}
	return &types.QueryResolveNameResponse{Record: record}, nil
}

// GetName returns a NameRecord regardless of status.
func (k Keeper) GetName(c context.Context, req *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Namespace == "" || req.NormalizedName == "" {
		return nil, status.Error(codes.InvalidArgument, "namespace and normalized_name are required")
	}
	ctx := sdk.UnwrapSDKContext(c)
	record, found := k.GetNameRecord(ctx, req.Namespace, req.NormalizedName)
	if !found {
		return nil, status.Error(codes.NotFound, errorsmod.Wrapf(types.ErrNameNotFound, "%s/%s", req.Namespace, req.NormalizedName).Error())
	}
	return &types.QueryGetNameResponse{Record: record}, nil
}

// NamesByNamespace returns paginated NameRecords under a namespace.
func (k Keeper) NamesByNamespace(c context.Context, req *types.QueryNamesByNamespaceRequest) (*types.QueryNamesByNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Namespace == "" {
		return nil, status.Error(codes.InvalidArgument, "namespace is required")
	}
	ctx := sdk.UnwrapSDKContext(c)
	if !k.HasNamespace(ctx, req.Namespace) {
		return nil, status.Error(codes.NotFound, errorsmod.Wrapf(types.ErrNamespaceNotFound, "namespace %q", req.Namespace).Error())
	}
	out, pageRes, err := k.PaginateNamesByNamespace(ctx, req.Namespace, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if out == nil {
		out = []types.NameRecord{}
	}
	return &types.QueryNamesByNamespaceResponse{Records: out, Pagination: pageRes}, nil
}

// NamesByOwner returns paginated NameRecords owned by a DID across every
// namespace.
func (k Keeper) NamesByOwner(c context.Context, req *types.QueryNamesByOwnerRequest) (*types.QueryNamesByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.OwnerDid == "" {
		return nil, status.Error(codes.InvalidArgument, "owner_did is required")
	}
	ctx := sdk.UnwrapSDKContext(c)
	out, pageRes, err := k.PaginateNamesByOwner(ctx, req.OwnerDid, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if out == nil {
		out = []types.NameRecord{}
	}
	return &types.QueryNamesByOwnerResponse{Records: out, Pagination: pageRes}, nil
}
