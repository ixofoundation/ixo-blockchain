package project

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paymentskeeper "github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	types "github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func NewHandler(k keeper.Keeper, pk paymentskeeper.Keeper, bk bankkeeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k, bk, pk)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateProject:
			res, err := msgServer.CreateProject(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateProjectStatus:
			res, err := msgServer.UpdateProjectStatus(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateAgent:
			res, err := msgServer.CreateAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateAgent:
			res, err := msgServer.UpdateAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateClaim:
			res, err := msgServer.CreateClaim(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateEvaluation:
			res, err := msgServer.CreateEvaluation(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawFunds:
			res, err := msgServer.WithdrawFunds(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"unrecognized project Msg type: %v", msg.Type())
		}
	}
}
