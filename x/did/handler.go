package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k DidKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case AddDidMsg:
			return handleAddDidDocMsg(ctx, k, msg)
		case AddCredentialMsg:
			return handleAddCredentialMsg(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleAddDidDocMsg(ctx sdk.Context, k DidKeeper, msg AddDidMsg) sdk.Result {
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

func handleAddCredentialMsg(ctx sdk.Context, k DidKeeper, msg AddCredentialMsg) sdk.Result {
	didDoc, err := k.AddCredential(ctx, msg.DidCredential.Claim.Id, msg.DidCredential)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
		Data: k.dm.encodeDid(didDoc),
	}
}
