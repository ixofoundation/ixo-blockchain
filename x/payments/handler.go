package payments

import (
	"fmt"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
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

func NewHandler(k Keeper, bk bankkeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgSetPaymentContractAuthorisation:
			return handleMsgSetPaymentContractAuthorisation(ctx, k, msg)
		case MsgCreatePaymentTemplate:
			return handleMsgCreatePaymentTemplate(ctx, k, bk, msg)
		case MsgCreatePaymentContract:
			return handleMsgCreatePaymentContract(ctx, k, bk, msg)
		case MsgCreateSubscription:
			return handleMsgCreateSubscription(ctx, k, msg)
		case MsgGrantDiscount:
			return handleMsgGrantDiscount(ctx, k, msg)
		case MsgRevokeDiscount:
			return handleMsgRevokeDiscount(ctx, k, msg)
		case MsgEffectPayment:
			return handleMsgEffectPayment(ctx, k, bk, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized payments Msg type: %v", msg.Type())
		}
	}
}

func handleMsgSetPaymentContractAuthorisation(ctx sdk.Context, k Keeper, msg MsgSetPaymentContractAuthorisation) (*sdk.Result, error) {

	// Get payment contract
	contract, err := k.GetPaymentContract(ctx, msg.PaymentContractId)
	if err != nil {
		return nil, err
	}

	// Get payer address
	payerDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.PayerDid)
	if err != nil {
		return nil, err
	}
	payerAddr := payerDidDoc.Address()

	// Confirm that signer is actually the payer in the payment contract
	if !payerAddr.Equals(contract.Payer) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")

	}

	// Set authorised status
	err = k.SetPaymentContractAuthorised(ctx, msg.PaymentContractId, msg.Authorised)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePaymentContractAuthorisation,
			sdk.NewAttribute(types.AttributeKeyPayerDid, msg.PayerDid),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyAuthorised, strconv.FormatBool(msg.Authorised)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgCreatePaymentTemplate(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, msg MsgCreatePaymentTemplate) (*sdk.Result, error) {

	// Ensure that payment template doesn't already exist
	if k.PaymentTemplateExists(ctx, msg.PaymentTemplate.Id) {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyExists,
			"payment template '%s' already exists", msg.PaymentTemplate.Id)
	}

	// Ensure that payment template ID is not reserved
	if k.PaymentTemplateIdReserved(msg.PaymentTemplate.Id) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed as it is "+
			"using a reserved prefix", msg.PaymentTemplate.Id)
	}

	// Create and validate payment template
	if err := msg.PaymentTemplate.Validate(); err != nil {
		return nil, err
	}

	// Submit payment template
	k.SetPaymentTemplate(ctx, msg.PaymentTemplate)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePaymentTemplate,
			sdk.NewAttribute(types.AttributeKeyCreatorDid, msg.CreatorDid),
			sdk.NewAttribute(types.AttributeKeyAttributeKeyId, msg.PaymentTemplate.Id),
			sdk.NewAttribute(types.AttributeKeyPaymentAmount, msg.PaymentTemplate.PaymentAmount.String()),
			sdk.NewAttribute(types.AttributeKeyPaymentMinimum, msg.PaymentTemplate.PaymentMinimum.String()),
			sdk.NewAttribute(types.AttributeKeyPaymentMaximum, msg.PaymentTemplate.PaymentMaximum.String()),
			sdk.NewAttribute(types.AttributeKeyDiscounts, fmt.Sprint(msg.PaymentTemplate.Discounts)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgCreatePaymentContract(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper,
	msg MsgCreatePaymentContract) (*sdk.Result, error) {

	// Ensure that payment contract doesn't already exist
	if k.PaymentContractExists(ctx, msg.PaymentContractId) {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyExists,
			"payment contract '%s' already exists", msg.PaymentContractId)
	}

	// Ensure that payment contract ID is not reserved
	if k.PaymentContractIdReserved(msg.PaymentContractId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed as it is "+
			"using a reserved prefix", msg.PaymentContractId)
	}

	// Ensure payer is not a blocked address
	if bk.BlockedAddr(msg.Payer) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed "+
			"to receive transactions", msg.Payer)
	}

	// Confirm that payment template exists
	if !k.PaymentTemplateExists(ctx, msg.PaymentTemplateId) {
		return nil, fmt.Errorf("invalid payment template")
	}

	// Get creator address
	cretorDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.CreatorDid)
	if err != nil {
		return nil, err
	}
	creatorAddr := cretorDidDoc.Address()

	// Create payment contract and validate
	authorised := false
	contract := NewPaymentContract(
		msg.PaymentContractId, msg.PaymentTemplateId, creatorAddr, msg.Payer,
		msg.Recipients, msg.CanDeauthorise, authorised, msg.DiscountId)
	if err := contract.Validate(); err != nil {
		return nil, err
	}

	// Ensure no blocked address in wallet distribution
	for _, share := range msg.Recipients {
		if bk.BlockedAddr(share.Address) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed "+
				"to receive transactions", share.Address)
		}
	}

	// Submit payment contract
	k.SetPaymentContract(ctx, contract)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePaymentContract,
			sdk.NewAttribute(types.AttributeKeyCreatorDid, msg.CreatorDid),
			sdk.NewAttribute(types.AttributeKeyPaymentTemplateId, msg.PaymentTemplateId),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyPayer, msg.Payer.String()),
			sdk.NewAttribute(types.AttributeKeyRecipients, fmt.Sprint(msg.Recipients)),
			sdk.NewAttribute(types.AttributeKeyDiscountId, msg.DiscountId.String()),
			sdk.NewAttribute(types.AttributeKeyCanDeauthorise, strconv.FormatBool(msg.CanDeauthorise)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgCreateSubscription(ctx sdk.Context, k Keeper,
	msg MsgCreateSubscription) (*sdk.Result, error) {

	// Ensure that subscription doesn't already exist
	if k.SubscriptionExists(ctx, msg.SubscriptionId) {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyExists,
			"subscription '%s' already exists", msg.SubscriptionId)
	}

	// Ensure that subscription ID is not reserved
	if k.SubscriptionIdReserved(msg.SubscriptionId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed as it is "+
			"using a reserved prefix", msg.SubscriptionId)
	}

	// Get payment contract
	contract, err := k.GetPaymentContract(ctx, msg.PaymentContractId)
	if err != nil {
		return nil, err
	}

	// Get creator address
	cretorDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.CreatorDid)
	if err != nil {
		return nil, err
	}
	creatorAddr := cretorDidDoc.Address()

	// Confirm that signer is actually the creator of the payment contract
	if !creatorAddr.Equals(contract.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract creator")
	}

	// Create subscription and validate
	subscription := NewSubscription(msg.SubscriptionId,
		msg.PaymentContractId, msg.MaxPeriods, msg.Period)
	if err := subscription.Validate(); err != nil {
		return nil, err
	}

	// Submit subscription
	k.SetSubscription(ctx, subscription)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateSubscription,
			sdk.NewAttribute(types.AttributeKeySubscriptionId, msg.SubscriptionId),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyMaxPeriods, msg.MaxPeriods.String()),
			sdk.NewAttribute(types.AttributeKeyPeriod, msg.Period.GetPeriodUnit()),
			sdk.NewAttribute(types.AttributeKeyCreatorDid, msg.CreatorDid),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgGrantDiscount(ctx sdk.Context, k Keeper, msg MsgGrantDiscount) (*sdk.Result, error) {

	// Get PaymentContract
	contract, err := k.GetPaymentContract(ctx, msg.PaymentContractId)
	if err != nil {
		return nil, err
	}

	// Get creator address
	creatorDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.SenderDid)
	if err != nil {
		return nil, err
	}
	creatorAddr := creatorDidDoc.Address()

	// Confirm that signer is actually the creator of the payment contract
	if !creatorAddr.Equals(contract.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
			"signer must be payment contract creator")

	}

	// Confirm that discount ID is in the template (to avoid invalid discount IDs)
	found, err := k.DiscountIdExists(ctx, contract.PaymentTemplateId, msg.DiscountId)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidId,
			"discount ID not in payment template's discount list")
	}

	// Grant the discount
	err = k.GrantDiscount(ctx, contract.Id, msg.DiscountId)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeGrantDiscount,
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyDiscountId, msg.DiscountId.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgRevokeDiscount(ctx sdk.Context, k Keeper, msg MsgRevokeDiscount) (*sdk.Result, error) {

	// Get PaymentContract
	contract, err := k.GetPaymentContract(ctx, msg.PaymentContractId)
	if err != nil {
		return nil, err
	}

	// Get creator address
	cretorDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.SenderDid)
	if err != nil {
		return nil, err
	}
	creatorAddr := cretorDidDoc.Address()

	// Confirm that signer is actually the creator of the payment contract
	if !creatorAddr.Equals(contract.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract creator")

	}

	// Revoke the discount
	err = k.RevokeDiscount(ctx, contract.Id)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRevokeDiscount,
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyHolder, msg.Holder.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgEffectPayment(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, msg MsgEffectPayment) (*sdk.Result, error) {

	// Get payment contract
	contract, err := k.GetPaymentContract(ctx, msg.PaymentContractId)
	if err != nil {
		return nil, err
	}

	// Get creator address
	cretorDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.SenderDid)
	if err != nil {
		return nil, err
	}
	creatorAddr := cretorDidDoc.Address()

	// Confirm that signer is actually the creator of the payment contract
	if !creatorAddr.Equals(contract.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract creator")
	}

	// Effect payment
	effected, err := k.EffectPayment(ctx, bk, msg.PaymentContractId)
	if err != nil {
		return nil, err
	}

	// Payment not effected but no error, meaning that payment should have been effected
	if !effected {
		return nil, fmt.Errorf("payment not effected due to unknown reason")
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEffectPayment,
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
