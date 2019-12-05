package fiat

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewAnteHandler(fiatKeeper Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {

		ctx = ctx.WithEventManager(sdk.NewEventManager())

		msg := tx.GetMsgs()[0]

		switch msg := msg.(type) {
		case MsgIssueFiats:
			return handleMsgIssueFiats(ctx, fiatKeeper, msg.IssueFiats)
		case MsgRedeemFiats:
			return handleMsgRedeemFiats(ctx, fiatKeeper, msg.RedeemFiats)
		case MsgSendFiats:
			return handleMsgSendFiats(ctx, fiatKeeper, msg.SendFiats)
		default:
			return ctx, sdk.ErrInternal(fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)).Result(), true
		}

	}
}

func handleMsgIssueFiats(ctx sdk.Context, fiatKeeper Keeper, issueFiats []IssueFiat) (_ sdk.Context, _ sdk.Result, abort bool) {

	for _, issueFiat := range issueFiats {

		err := fiatKeeper.IssueFiats(ctx, issueFiat)
		if err != nil {
			return ctx, err.Result(), true
		}
	}

	return ctx, sdk.Result{
		Events: ctx.EventManager().Events(),
	}, false
}

func handleMsgRedeemFiats(ctx sdk.Context, fiatKeeper Keeper, redeemFiats []RedeemFiat) (_ sdk.Context, _ sdk.Result, abort bool) {
	for _, redeemFiat := range redeemFiats {
		err := fiatKeeper.RedeemFiats(ctx, redeemFiat)
		if err != nil {
			return ctx, err.Result(), true
		}
	}

	return ctx, sdk.Result{
		Events: ctx.EventManager().Events(),
	}, false

}

func handleMsgSendFiats(ctx sdk.Context, fiatKeeper Keeper, sendFiats []SendFiat) (_ sdk.Context, _ sdk.Result, abort bool) {
	for _, sendFiat := range sendFiats {
		err := fiatKeeper.SendFiats(ctx, sendFiat)
		if err != nil {
			return ctx, err.Result(), true
		}
	}

	return ctx, sdk.Result{
		Events: ctx.EventManager().Events(),
	}, false

}
