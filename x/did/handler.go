package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgAddDid:
			res, err := msgServer.AddDid(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			//return handleMsgAddDidDoc(ctx, k, msg)
		case *types.MsgAddCredential:
			res, err := msgServer.AddCredential(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			//return handleMsgAddCredential(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized did Msg type: %v", msg.Type())
		}
	}
}

// functions previously here are now in msg_server.go

//func handleMsgAddDidDoc(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddDid) (*sdk.Result, error) {
//	didDoc := types.NewBaseDidDoc(msg.Did, msg.PubKey)
//
//	err := k.SetDidDoc(ctx, didDoc)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitEvents(sdk.Events{
//		sdk.NewEvent(
//			types.EventTypeAddDidDoc,
//			sdk.NewAttribute(types.AttributeKeyDid, msg.Did),
//			sdk.NewAttribute(types.AttributeKeyPubKey, msg.PubKey),
//		),
//		sdk.NewEvent(
//			sdk.EventTypeMessage,
//			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
//		),
//	})
//	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
//	//return &sdk.Result{Events: ctx.EventManager().Events()}, nil
//}
//
//func handleMsgAddCredential(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddCredential) (*sdk.Result, error) {
//	err := k.AddCredentials(ctx, msg.DidCredential.Claim.Id, msg.DidCredential)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitEvents(sdk.Events{
//		sdk.NewEvent(
//			types.EventTypeAddCredential,
//			sdk.NewAttribute(types.AttributeKeyCredType, fmt.Sprint(msg.DidCredential.CredType)),
//			sdk.NewAttribute(types.AttributeKeyIssuer, msg.DidCredential.Issuer),
//			sdk.NewAttribute(types.AttributeKeyIssued, msg.DidCredential.Issued),
//			sdk.NewAttribute(types.AttributeKeyClaimID, msg.DidCredential.Claim.Id),
//			sdk.NewAttribute(types.AttributeKeyKYCValidated, strconv.FormatBool(msg.DidCredential.Claim.KYCValidated)),
//		),
//		sdk.NewEvent(
//			sdk.EventTypeMessage,
//			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
//		),
//	})
//	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
//	//return &sdk.Result{Events: ctx.EventManager().Events()}, nil
//}
