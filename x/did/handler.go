package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddDid:
			return handleMsgAddDidDoc(ctx, k, msg)
		case types.MsgAddCredential:
			return handleMsgAddCredential(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgAddDidDoc(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddDid) sdk.Result {
	newDidDoc := msg.DidDoc

	if len(newDidDoc.Credentials) > 0 {
		return sdk.ErrUnknownRequest("Cannot add a new DID with existing Credentials").Result()
	}

	err := k.SetDidDoc(ctx, newDidDoc)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgAddCredential(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddCredential) sdk.Result {
	err := k.AddCredentials(ctx, msg.DidCredential.Claim.Id, msg.DidCredential)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}
