package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TokenDoc(c context.Context, req *types.QueryTokenDocRequest) (*types.QueryTokenDocResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	token, err := k.GetToken(ctx, req.Minter, req.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &types.QueryTokenDocResponse{TokenDoc: token}, nil
}

func (k Keeper) TokenList(c context.Context, req *types.QueryTokenListRequest) (*types.QueryTokenListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokens []types.Token
	ctx := sdk.UnwrapSDKContext(c)
	tokensStore := k.GetMinterTokensStore(ctx, req.Minter)

	pageRes, err := query.Paginate(tokensStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.Token
		if err := k.cdc.Unmarshal(value, &token); err != nil {
			return err
		}

		tokens = append(tokens, token)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTokenListResponse{TokenDocs: tokens, Pagination: pageRes}, nil
}

func (k Keeper) TokenMetadata(c context.Context, req *types.QueryTokenMetadataRequest) (*types.QueryTokenMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	tokenProperties, token, err := k.GetTokenById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	properties := types.TokenMetadataProperties{
		Class:           token.Class,
		Collection:      tokenProperties.Collection,
		Cap:             token.Cap.String(),
		LinkedResources: tokenProperties.TokenData,
	}

	return &types.QueryTokenMetadataResponse{
		Name:        token.Name,
		Description: token.Description,
		Decimals:    "0",
		Image:       token.Image,
		Index:       tokenProperties.Index,
		Properties:  &properties,
	}, nil
}
