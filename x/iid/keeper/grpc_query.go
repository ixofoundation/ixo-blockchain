package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

var _ types.QueryServer = Keeper{}

// IidDocuments implements the DidDocuments gRPC method
func (k Keeper) IidDocuments(
	c context.Context,
	req *types.QueryIidDocumentsRequest,
) (*types.QueryIidDocumentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var iidDocs []types.IidDocument
	iidDocStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.DidDocumentKey)

	pageRes, err := query.Paginate(iidDocStore, req.Pagination, func(key []byte, value []byte) error {
		var iidDoc types.IidDocument
		if err := k.cdc.Unmarshal(value, &iidDoc); err != nil {
			return err
		}

		iidDocs = append(iidDocs, iidDoc)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryIidDocumentsResponse{
		IidDocuments: iidDocs,
		Pagination: pageRes,
	}, nil
}

// IidDocument implements the IidDocument gRPC method
func (k Keeper) IidDocument(
	c context.Context,
	req *types.QueryIidDocumentRequest,
) (*types.QueryIidDocumentResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "iid Doc id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	doc, err := k.ResolveDid(ctx, types.DID(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &types.QueryIidDocumentResponse{
		IidDocument: doc,
	}, nil
}
