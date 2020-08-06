package did

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
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
	didDoc := types.NewBaseDidDoc(msg.Did, msg.PubKey)

	err := k.SetDidDoc(ctx, didDoc)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddDidDoc,
			sdk.NewAttribute(types.AttributeKeyDid, msg.Did),
			sdk.NewAttribute(types.AttributeKeyPubKey, msg.PubKey),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgAddCredential(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddCredential) sdk.Result {
	err := k.AddCredentials(ctx, msg.DidCredential.Claim.Id, msg.DidCredential)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddCredential,
			sdk.NewAttribute(types.AttributeKeyCredType, fmt.Sprint(msg.DidCredential.CredType)),
			sdk.NewAttribute(types.AttributeKeyIssuer, msg.DidCredential.Issuer),
			sdk.NewAttribute(types.AttributeKeyIssued, msg.DidCredential.Issued),
			sdk.NewAttribute(types.AttributeKeyClaimID, msg.DidCredential.Claim.Id),
			sdk.NewAttribute(types.AttributeKeyKYCValidated, strconv.FormatBool(msg.DidCredential.Claim.KYCValidated)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return sdk.Result{Events: ctx.EventManager().Events()}
}
