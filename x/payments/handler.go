package payments

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	iterator := keeper.GetSubscriptionIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())

		// Skip if should not effect
		if !subscription.ShouldEffect(ctx) {
			continue
		}

		// Effect subscription payment
		err := keeper.EffectSubscriptionPayment(ctx, subscription.Id)
		if err != nil {
			panic(err) // TODO: maybe shouldn't panic?
		}

		// Note: if payment can be re-effected immediately, this should be done
		// in the next block to prevent spending too much time effecting payments

		// Get updated subscription
		subscription, err = keeper.GetSubscription(ctx, subscription.Id)
		if err != nil {
			panic(err)
		}

		// Delete subscription if it has completed
		if subscription.IsComplete() {
			// TODO: delete subscription
		}

		// Note: no need to save the subscription, as it is being saved by the
		// functions operating on it, such as EffectSubscriptionPayment()
	}
	return []abci.ValidatorUpdate{}
}

func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgSetPaymentContractAuthorisation:
			res, err := msgServer.SetPaymentContractAuthorisation(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreatePaymentTemplate:
			res, err := msgServer.CreatePaymentTemplate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreatePaymentContract:
			res, err := msgServer.CreatePaymentContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateSubscription:
			res, err := msgServer.CreateSubscription(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgGrantDiscount:
			res, err := msgServer.GrantDiscount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRevokeDiscount:
			res, err := msgServer.RevokeDiscount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgEffectPayment:
			res, err := msgServer.EffectPayment(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized payments Msg type: %v", msg.Type())
		}
	}
}
