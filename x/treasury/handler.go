package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgSend:
			return handleMsgSend(ctx, k, msg)
		case MsgOracleTransfer:
			return handleMsgOracleTransfer(ctx, k, msg)
		case MsgOracleMint:
			return handleMsgOracleMint(ctx, k, msg)
		case MsgOracleBurn:
			return handleMsgOracleBurn(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized treasury Msg type: %v", msg.Type())
		}
	}
	// TODO: be able to disable sends/mints/burns globally
	// TODO: be able to blacklist addresses/DIDs
}

func handleMsgSend(ctx sdk.Context, k keeper.Keeper, msg types.MsgSend) (*sdk.Result, error) {

	if err := k.Send(ctx, msg.FromDid, msg.ToDidOrAddr, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSend,
			sdk.NewAttribute(types.AttributeKeyTxHashFromDid, msg.FromDid),
			sdk.NewAttribute(types.AttributeKeyToDidOrAddr, msg.ToDidOrAddr),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgOracleTransfer(ctx sdk.Context, k keeper.Keeper, msg types.MsgOracleTransfer) (*sdk.Result, error) {

	if err := k.OracleTransfer(ctx, msg.FromDid, msg.ToDidOrAddr, msg.OracleDid, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeOracleTransfer,
			sdk.NewAttribute(types.AttributeKeyOracleDid, msg.OracleDid),
			sdk.NewAttribute(types.AttributeKeyTxHashFromDid, msg.FromDid),
			sdk.NewAttribute(types.AttributeKeyToDidOrAddr, msg.ToDidOrAddr),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyProof, msg.Proof),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgOracleMint(ctx sdk.Context, k keeper.Keeper, msg types.MsgOracleMint) (*sdk.Result, error) {

	if err := k.OracleMint(ctx, msg.OracleDid, msg.ToDidOrAddr, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeOracleMint,
			sdk.NewAttribute(types.AttributeKeyOracleDid, msg.OracleDid),
			sdk.NewAttribute(types.AttributeKeyToDidOrAddr, msg.ToDidOrAddr),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyProof, msg.Proof),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgOracleBurn(ctx sdk.Context, k keeper.Keeper, msg types.MsgOracleBurn) (*sdk.Result, error) {

	if err := k.OracleBurn(ctx, msg.OracleDid, msg.FromDid, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeOracleBurn,
			sdk.NewAttribute(types.AttributeKeyTxHashFromDid, msg.FromDid),
			sdk.NewAttribute(types.AttributeKeyOracleDid, msg.OracleDid),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyProof, msg.Proof),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
