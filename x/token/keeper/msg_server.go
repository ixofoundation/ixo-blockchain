package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
)

type msgServer struct {
	Keeper Keeper
}

// NewMsgServerImpl returns an implementation of the project MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

func (s msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resp, err := s.Keeper.CreateToken(ctx, msg)
	return &resp, err
}

func (s msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	return &types.MsgUpdateTokenResponse{}, nil
}

func (s msgServer) TransferToken(goCtx context.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	return s.Keeper.TransferToken(sdk.UnwrapSDKContext(goCtx), msg)

}
