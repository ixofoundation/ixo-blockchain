package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ixofoundation/ixo-blockchain/v2/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/token keeper providing gRPC method
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

func (q Querier) TokenDoc(c context.Context, req *types.QueryTokenDocRequest) (*types.QueryTokenDocResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	token, err := q.Keeper.GetToken(ctx, req.Minter, req.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &types.QueryTokenDocResponse{TokenDoc: token}, nil
}

func (q Querier) TokenList(c context.Context, req *types.QueryTokenListRequest) (*types.QueryTokenListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokens []types.Token
	ctx := sdk.UnwrapSDKContext(c)
	tokensStore := q.Keeper.GetMinterTokensStore(ctx, req.Minter)

	pageRes, err := query.Paginate(tokensStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.Token
		if err := q.Keeper.cdc.Unmarshal(value, &token); err != nil {
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

func (q Querier) TokenMetadata(c context.Context, req *types.QueryTokenMetadataRequest) (*types.QueryTokenMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	tokenProperties, token, err := q.Keeper.GetTokenById(ctx, req.Id)
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
