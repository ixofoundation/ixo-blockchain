package entity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateEntity:
			res, err := msgServer.CreateEntity(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateEntity:
			res, err := msgServer.UpdateEntity(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateEntityVerified:
			res, err := msgServer.UpdateEntityVerified(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgTransferEntity:
			res, err := msgServer.TransferEntity(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateEntityAccount:
			res, err := msgServer.CreateEntityAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgGrantEntityAccountAuthz:
			res, err := msgServer.GrantEntityAccountAuthz(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			// err := sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized bonds Msg type: %v", msg.Type())
			err := sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unrecognized entity Msg type")
			return nil, err
		}
	}
}
