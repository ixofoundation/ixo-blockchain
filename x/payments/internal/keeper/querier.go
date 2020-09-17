package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryParams          = "queryParams"
	QueryPaymentTemplate = "queryPaymentTemplate"
	QueryPaymentContract = "queryPaymentContract"
	QuerySubscription    = "querySubscription"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryPaymentTemplate:
			return queryPaymentTemplate(ctx, path[1:], k)
		case QueryPaymentContract:
			return queryPaymentContract(ctx, path[1:], k)
		case QuerySubscription:
			return querySubscription(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown payments query endpoint")
		}
	}
}

func queryPaymentTemplate(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	templateId := path[0]

	template, err := k.GetPaymentTemplate(ctx, templateId)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf(
			"payment template '%s' does not exist", templateId))
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, template)
	if err2 != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err2.Error()))
	}

	return res, nil
}

func queryPaymentContract(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	contractId := path[0]

	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf(
			"payment contract '%s' does not exist", contractId))
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, contract)
	if err2 != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err2.Error()))
	}

	return res, nil
}

func querySubscription(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	subscriptionId := path[0]

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf(
			"subscription '%s' does not exist", subscriptionId))
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, subscription)
	if err2 != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err2.Error()))
	}

	return res, nil
}
