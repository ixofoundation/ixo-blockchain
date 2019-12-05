package fiat

import (
	"fmt"

	cTypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewAnteHandler(fiatKeeper Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {

		ctx = ctx.WithEventManager(cTypes.NewEventManager())

		msg := tx.GetMsgs()[0]

		switch msg := msg.(type) {
		// case MsgChangeBuyerBids:
		// 	return handleMsgChangeBids(ctx, k, msg.ChangeBids)
		// case MsgChangeSellerBids:
		// 	return handleMsgChangeBids(ctx, k, msg.ChangeBids)
		// case MsgConfirmSellerBids:
		// 	return handleMsgConfirmBids(ctx, k, msg.ConfirmBids)
		// case MsgConfirmBuyerBids:
		// 	return handleMsgConfirmBids(ctx, k, msg.ConfirmBids)

		default:
			return ctx, sdk.ErrInternal(fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)).Result(), true
		}

	}
}
