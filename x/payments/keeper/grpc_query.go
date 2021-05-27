package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) PaymentTemplate(c context.Context, req *types.QueryPaymentTemplateRequest) (*types.QueryPaymentTemplateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	templateId := req.PaymentTemplateId

	template, err := k.GetPaymentTemplate(ctx, templateId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"payment template '%s' does not exist", templateId)
	}

	return &types.QueryPaymentTemplateResponse{PaymentTemplate: template}, nil
}

func (k Keeper) PaymentContract(c context.Context, req *types.QueryPaymentContractRequest) (*types.QueryPaymentContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	contractId := req.PaymentContractId

	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "payment contract '%s' does not exist", contractId)
	}

	return &types.QueryPaymentContractResponse{PaymentContract: contract}, nil
}

func (k Keeper) PaymentContractsByIdPrefix(c context.Context, req *types.QueryPaymentContractsByIdPrefixRequest) (*types.QueryPaymentContractsByIdPrefixResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	contractIdPrefix := req.PaymentContractsIdPrefix

	contracts := k.GetPaymentContractsByPrefix(ctx, contractIdPrefix)

	return &types.QueryPaymentContractsByIdPrefixResponse{PaymentContracts: contracts}, nil
}

func (k Keeper) Subscription(c context.Context, req *types.QuerySubscriptionRequest) (*types.QuerySubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscriptionId := req.SubscriptionId

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "subscription '%s' does not exist", subscriptionId)
	}

	return &types.QuerySubscriptionResponse{Subscription: subscription}, nil
}
