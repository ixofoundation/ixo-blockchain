package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgSend:
			return handleMsgSend(ctx, k, msg)
		case MsgSendOnBehalfOf:
			return handleMsgSendOnBehalfOf(ctx, k, msg)
		case MsgMint:
			return handleMsgMint(ctx, k, msg)
		case MsgBurn:
			return handleMsgBurn(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}

	// TODO: be able to disable sends/mints/burns globally
	// TODO: be able to blacklist addresses/DIDs
}

func handleMsgSend(ctx sdk.Context, k keeper.Keeper, msg types.MsgSend) sdk.Result {

	if err := k.Send(ctx, msg.FromDid, msg.ToDid, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSendOnBehalfOf(ctx sdk.Context, k keeper.Keeper, msg types.MsgSendOnBehalfOf) sdk.Result {

	if err := k.SendOnBehalfOf(ctx, msg.FromDid, msg.ToDid, msg.OracleDid, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgMint(ctx sdk.Context, k keeper.Keeper, msg types.MsgMint) sdk.Result {

	if err := k.Mint(ctx, msg.OracleDid, msg.ToDid, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgBurn(ctx sdk.Context, k keeper.Keeper, msg types.MsgBurn) sdk.Result {

	if err := k.Burn(ctx, msg.OracleDid, msg.FromDid, msg.Amount); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}
