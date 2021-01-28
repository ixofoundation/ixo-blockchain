package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"strconv"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddDid(goCtx context.Context, msg *types.MsgAddDid) (*types.MsgAddDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	didDoc := types.NewBaseDidDoc(msg.Did, msg.PubKey)

	err := k.SetDidDoc(ctx, didDoc)
	if err != nil {
		return nil, err
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
	return &types.MsgAddDidResponse{}, nil //&sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func (k msgServer) AddCredential(goCtx context.Context, msg *types.MsgAddCredential) (*types.MsgAddCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.AddCredentials(ctx, msg.DidCredential.Claim.Id, *msg.DidCredential)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddCredential,
			sdk.NewAttribute(types.AttributeKeyCredType, fmt.Sprint(msg.DidCredential.Credtype)),
			sdk.NewAttribute(types.AttributeKeyIssuer, msg.DidCredential.Issuer),
			sdk.NewAttribute(types.AttributeKeyIssued, msg.DidCredential.Issued),
			sdk.NewAttribute(types.AttributeKeyClaimID, msg.DidCredential.Claim.Id),
			sdk.NewAttribute(types.AttributeKeyKYCValidated, strconv.FormatBool(msg.DidCredential.Claim.KYCvalidated)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &types.MsgAddCredentialResponse{}, nil //&sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}