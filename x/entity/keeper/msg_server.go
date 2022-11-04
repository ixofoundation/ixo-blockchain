package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
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

func (s msgServer) CreateEntity(goCtx context.Context, msg *types.MsgCreateEntity) (*types.MsgCreateEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resp, err := s.Keeper.CreateEntity(ctx, msg)
	return &resp, err
}

func (s msgServer) UpdateEntity(goCtx context.Context, msg *types.MsgUpdateEntity) (*types.MsgUpdateEntityResponse, error) {
	return &types.MsgUpdateEntityResponse{}, nil
}

func (s msgServer) TransferEntity(goCtx context.Context, msg *types.MsgTransferEntity) (*types.MsgTransferEntityResponse, error) {
	return s.Keeper.TransferEntity(sdk.UnwrapSDKContext(goCtx), msg)

}
