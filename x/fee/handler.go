package fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// switch msg := msg.(type) {
		// case oracle.Msg:
		// 	return keeper.Oracle
		// 	return oracle.Handle(ctx sdk.Context, p Payload) sdk.Error {
		// 		switch p := p.(type) {
		// 		case Payload:
		// 			return handleMyPayload(ctx, keeper, p)
		// 		}
		// 	}
		// }
		return sdk.ErrUnknownRequest("No match for message type.").Result()
	}
}
