package fiat

import (
	"fmt"

	cTypes "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) cTypes.Handler {
	return func(ctx cTypes.Context, msg cTypes.Msg) cTypes.Result {
		ctx = ctx.WithEventManager(cTypes.NewEventManager())

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
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return cTypes.ErrUnknownRequest(errMsg).Result()

		}

	}
}

// func handleMsgChangeBids(ctx cTypes.Context, negotiationKeeper Keeper, changeBids []ChangeBid) cTypes.Result {

// 	for _, changeBid := range changeBids {

// 		err := negotiationKeeper.ChangeNegotiationBidWithACL(ctx, changeBid)
// 		if err != nil {
// 			return err.Result()
// 		}
// 	}

// 	return cTypes.Result{
// 		Events: ctx.EventManager().Events(),
// 	}
// }

// func handleMsgConfirmBids(ctx cTypes.Context, negotitationKeeper Keeper, confirmBids []ConfirmBid) cTypes.Result {
// 	for _, confirmBid := range confirmBids {
// 		err := negotitationKeeper.ConfirmNegotiationBidWithACL(ctx, confirmBid)
// 		if err != nil {
// 			return err.Result()
// 		}
// 	}

// 	return cTypes.Result{
// 		Events: ctx.EventManager().Events(),
// 	}

// }
