package keeper

import (
	"context"
	"errors"
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidDoc(c context.Context, req *types.QueryDidDocRequest) (*types.QueryDidDocResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: no DID specified")
	}

	ctx := sdk.UnwrapSDKContext(c)

	didDoc, err := k.GetDidDoc(ctx, req.Did)
	if err != nil {
		return nil, err
	}

	any, err := codectypes.NewAnyWithValue(didDoc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types.QueryDidDocResponse{Diddoc: any}, nil
}

func (k Keeper) AllDids(c context.Context, req *types.QueryAllDidsRequest) (*types.QueryAllDidsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	allDids := k.GetAllDids(ctx)

	return &types.QueryAllDidsResponse{Dids: allDids}, nil
}

func (k Keeper) AllDidDocs(c context.Context, req *types.QueryAllDidDocsRequest) (*types.QueryAllDidDocsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var didDocs []exported.DidDoc
	didDocs = k.GetAllDidDocs(ctx)

	anys := make([]*codectypes.Any, len(didDocs))

	for i, dd := range didDocs {
		msg, ok := dd.(proto.Message)
		if !ok {
			panic(fmt.Errorf("cannot proto marshal %T", dd))
		}
		any, err := codectypes.NewAnyWithValue(msg)
		if err != nil {
			anys[i] = any
		}
	}

	return &types.QueryAllDidDocsResponse{Diddocs: anys}, nil
}

func (k Keeper) AddressFromDid(c context.Context, req *types.QueryAddressFromDidRequest) (*types.QueryAddressFromDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: no DID specified")
	}

	ctx := sdk.UnwrapSDKContext(c)

	didDoc, err := k.GetDidDoc(ctx, req.Did)
	if err != nil {
		return nil, err
	}

	return &types.QueryAddressFromDidResponse{Address: didDoc.Address().String()}, nil
}

func (k Keeper) AddressFromBase58EncodedPubkey(c context.Context, req *types.QueryAddressFromBase58EncodedPubkeyRequest) (*types.QueryAddressFromBase58EncodedPubkeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: no pubkey specified")
	}

	if !types.IsValidPubKey(req.PubKey) {
		return &types.QueryAddressFromBase58EncodedPubkeyResponse{}, errors.New("input is not a valid base-58 encoded pubKey")
	}

	accAddress := exported.VerifyKeyToAddr(req.PubKey)

	return &types.QueryAddressFromBase58EncodedPubkeyResponse{Address: accAddress.String()}, nil
}