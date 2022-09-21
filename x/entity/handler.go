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
		case *types.MsgUpdateEntityStatus:
			res, err := msgServer.UpdateEntityStatus(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateEntityConfig:
			res, err := msgServer.UpdateEntityConfig(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgUpdateProjectStatus:
		// 	res, err := msgServer.UpdateProjectStatus(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgCreateAgent:
		// 	res, err := msgServer.CreateAgent(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgUpdateAgent:
		// 	res, err := msgServer.UpdateAgent(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgCreateClaim:
		// 	res, err := msgServer.CreateClaim(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgCreateEvaluation:
		// 	res, err := msgServer.CreateEvaluation(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgWithdrawFunds:
		// 	res, err := msgServer.WithdrawFunds(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgUpdateProjectDoc:
		// 	res, err := msgServer.UpdateProjectDoc(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		default:
			// err := sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized bonds Msg type: %v", msg.Type())
			err := sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unrecognized bonds Msg")
			return nil, err
		}
	}
}
