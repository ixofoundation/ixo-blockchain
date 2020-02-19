package bonddoc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type InternalAccountID = string

func NewHandler(k Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case CreateBondMsg:
			return handleCreateBondMsg(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleCreateBondMsg(ctx sdk.Context, k Keeper, msg CreateBondMsg) sdk.Result {

	err := k.SetBondDoc(ctx, &msg)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}
