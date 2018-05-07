package did

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k DidKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case AddDidMsg:
			return handleAddDidDocMsg(ctx, k, msg)
		case DidMsg:
			return handleGetDidDocdMsg(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleAddDidDocMsg(ctx sdk.Context, k DidKeeper, msg AddDidMsg) sdk.Result {
	fmt.Println("Handler")
	fmt.Println(msg)
	fmt.Println(msg.DidDoc)
	//	newDidDoc := k.dm.NewDidDoc(ctx, msg)
	newDidDoc := msg.DidDoc
	didDoc, err := k.AddDidDoc(ctx, newDidDoc)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
		Data: k.dm.encodeDid(didDoc),
	}
}

func handleGetDidDocdMsg(ctx sdk.Context, k DidKeeper, msg DidMsg) sdk.Result {
	didDoc := k.GetDidDoc(ctx, msg.Did)

	return sdk.Result{
		Code: sdk.CodeOK,
		Data: k.dm.encodeDid(didDoc),
	}
}
