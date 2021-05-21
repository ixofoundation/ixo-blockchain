package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"strconv"
)

type msgServer struct {
	Keeper
	BankKeeper     bankkeeper.Keeper
}

// NewMsgServerImpl returns an implementation of the gov MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, bankKeeper bankkeeper.Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper, BankKeeper: bankKeeper}
}

func (k msgServer) SetPaymentContractAuthorisation(goCtx context.Context, msg *types.MsgSetPaymentContractAuthorisation) (*types.MsgSetPaymentContractAuthorisationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	contractPayerAddr, err := sdk.AccAddressFromBech32(contract.Payer)
	if err != nil {
		return nil, err
	}
	if !payerAddr.Equals(contractPayerAddr) {
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

	return &types.MsgSetPaymentContractAuthorisationResponse{}, nil
}

func (k msgServer) CreatePaymentTemplate(goCtx context.Context, msg *types.MsgCreatePaymentTemplate) (*types.MsgCreatePaymentTemplateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	return &types.MsgCreatePaymentTemplateResponse{}, nil
}

func (k msgServer) CreatePaymentContract(goCtx context.Context, msg *types.MsgCreatePaymentContract) (*types.MsgCreatePaymentContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	bk := k.BankKeeper

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
	payerAddr, err := sdk.AccAddressFromBech32(msg.Payer)
	if err != nil {
		return nil, err
	}
	if bk.BlockedAddr(payerAddr) {
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
	contract := types.NewPaymentContract(
		msg.PaymentContractId, msg.PaymentTemplateId, creatorAddr, payerAddr,
		msg.Recipients, msg.CanDeauthorise, authorised, msg.DiscountId)
	if err := contract.Validate(); err != nil {
		return nil, err
	}

	// Ensure no blocked address in wallet distribution
	for _, share := range msg.Recipients {
		shareAddr, err := sdk.AccAddressFromBech32(share.Address)
		if err != nil {
			return nil, err
		}
		if bk.BlockedAddr(shareAddr) {
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
			sdk.NewAttribute(types.AttributeKeyPayer, msg.Payer),
			sdk.NewAttribute(types.AttributeKeyRecipients, fmt.Sprint(msg.Recipients)),
			sdk.NewAttribute(types.AttributeKeyDiscountId, msg.DiscountId.String()),
			sdk.NewAttribute(types.AttributeKeyCanDeauthorise, strconv.FormatBool(msg.CanDeauthorise)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgCreatePaymentContractResponse{}, nil
}

func (k msgServer) CreateSubscription(goCtx context.Context, msg *types.MsgCreateSubscription) (*types.MsgCreateSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	contractCreatorAddr, err := sdk.AccAddressFromBech32(contract.Creator)
	if err != nil {
		return nil, err
	}
	if !creatorAddr.Equals(contractCreatorAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract creator")
	}

	// Create subscription and validate
	subscription := types.NewSubscription(msg.SubscriptionId,
		msg.PaymentContractId, msg.MaxPeriods, msg.GetPeriod())
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
			sdk.NewAttribute(types.AttributeKeyPeriod, msg.GetPeriod().GetPeriodUnit()),
			sdk.NewAttribute(types.AttributeKeyCreatorDid, msg.CreatorDid),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgCreateSubscriptionResponse{}, nil
}

func (k msgServer) GrantDiscount(goCtx context.Context, msg *types.MsgGrantDiscount) (*types.MsgGrantDiscountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	contractCreatorAddr, err := sdk.AccAddressFromBech32(contract.Payer)
	if err != nil {
		return nil, err
	}
	if !creatorAddr.Equals(contractCreatorAddr) {
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
	err = k.Keeper.GrantDiscount(ctx, contract.Id, msg.DiscountId)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeGrantDiscount,
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyDiscountId, msg.DiscountId.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgGrantDiscountResponse{}, nil
}

func (k msgServer) RevokeDiscount(goCtx context.Context, msg *types.MsgRevokeDiscount) (*types.MsgRevokeDiscountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	contractCreatorAddr, err := sdk.AccAddressFromBech32(contract.Creator)
	if err != nil {
		return nil, err
	}
	if !creatorAddr.Equals(contractCreatorAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract creator")

	}

	// Revoke the discount
	err = k.Keeper.RevokeDiscount(ctx, contract.Id)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRevokeDiscount,
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyPaymentContractId, msg.PaymentContractId),
			sdk.NewAttribute(types.AttributeKeyHolder, msg.Holder),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgRevokeDiscountResponse{}, nil
}

func (k msgServer) EffectPayment(goCtx context.Context, msg *types.MsgEffectPayment) (*types.MsgEffectPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	bk := k.BankKeeper

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
	contractCreatorAddr, err := sdk.AccAddressFromBech32(contract.Creator)
	if err != nil {
		return nil, err
	}
	if !creatorAddr.Equals(contractCreatorAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract creator")
	}

	// Effect payment
	effected, err := k.Keeper.EffectPayment(ctx, bk, msg.PaymentContractId)
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

	return &types.MsgEffectPaymentResponse{}, nil
}