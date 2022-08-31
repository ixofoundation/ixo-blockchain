package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

var _ types.QueryServer = Keeper{}

// DidDocuments implements the DidDocuments gRPC method
func (k Keeper) IidDocuments(
	c context.Context,
	req *types.QueryIidDocumentsRequest,
) (*types.QueryIidDocumentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	dids := k.GetAllDidDocuments(ctx)

	return &types.QueryIidDocumentsResponse{
		IidDocuments: dids,
	}, nil
}

// DidDocument implements the DidDocument gRPC method
func (k Keeper) IidDocument(
	c context.Context,
	req *types.QueryIidDocumentRequest,
) (*types.QueryIidDocumentResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "verifiable credential id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	doc, _, err := k.ResolveDid(ctx, types.DID(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &types.QueryIidDocumentResponse{
		IidDocument: doc,
	}, nil
}
