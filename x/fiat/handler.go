package fiat

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(fiatKeeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgIssueFiats:
			return handleMsgIssueFiats(ctx, fiatKeeper, msg.IssueFiats)
		case MsgRedeemFiats:
			return handleMsgRedeemFiats(ctx, fiatKeeper, msg.RedeemFiats)
		case MsgSendFiats:
			return handleMsgSendFiats(ctx, fiatKeeper, msg.SendFiats)
		default:
			return sdk.ErrInternal(fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)).Result()
		}

	}
}

func handleMsgIssueFiats(ctx sdk.Context, fiatKeeper Keeper, issueFiats []IssueFiat) sdk.Result {

	for _, issueFiat := range issueFiats {

		err := fiatKeeper.IssueFiats(ctx, issueFiat)
		if err != nil {
			return err.Result()
		}
	}

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgRedeemFiats(ctx sdk.Context, fiatKeeper Keeper, redeemFiats []RedeemFiat) sdk.Result {
	for _, redeemFiat := range redeemFiats {
		err := fiatKeeper.RedeemFiats(ctx, redeemFiat)
		if err != nil {
			return err.Result()
		}
	}

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}

}

func handleMsgSendFiats(ctx sdk.Context, fiatKeeper Keeper, sendFiats []SendFiat) sdk.Result {
	for _, sendFiat := range sendFiats {
		err := fiatKeeper.SendFiats(ctx, sendFiat)
		if err != nil {
			return err.Result()
		}
	}

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}

}
