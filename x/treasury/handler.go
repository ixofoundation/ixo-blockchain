package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
)

func NewHandler(bk bank.Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgSend:
			return handleMsgSend(ctx, bk, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgSend(ctx sdk.Context, bk bank.Keeper, msg types.MsgSend) sdk.Result {
	// TODO?
	//if !k.GetSendEnabled(ctx) {
	//	return types.ErrSendDisabled(k.Codespace()).Result()
	//}

	// TODO?
	//if k.BlacklistedAddr(msg.ToAddress) {
	//	return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed to receive transactions", msg.ToAddress)).Result()
	//}

	fromAddress := types.DidToAddr(msg.FromDid)
	toAddress := types.DidToAddr(msg.ToDid)

	err := bk.SendCoins(ctx, fromAddress, toAddress, msg.Amount)
	if err != nil {
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
