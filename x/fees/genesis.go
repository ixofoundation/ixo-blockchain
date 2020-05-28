package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new fees genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Init params
	keeper.SetParams(ctx, data.Params)

	// Init fees
	for _, f := range data.Fees {
		keeper.SetFee(ctx, f)
	}

	// Init fee contracts
	for _, fc := range data.FeeContracts {
		keeper.SetFeeContract(ctx, fc)
	}

	// Init subscriptions
	for _, s := range data.Subscriptions {
		keeper.SetSubscription(ctx, s)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	// Export params
	params := keeper.GetParams(ctx)

	// Export fees
	var fees []Fee
	iterator := keeper.GetFeeIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		fee := keeper.MustGetFeeByKey(ctx, iterator.Key())
		fees = append(fees, fee)
	}

	// Export fee contracts
	var feeContracts []FeeContract
	iterator = keeper.GetFeeContractIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		feeContract := keeper.MustGetFeeContractByKey(ctx, iterator.Key())
		feeContracts = append(feeContracts, feeContract)
	}

	// Export subscriptions
	var subscriptions []Subscription
	iterator = keeper.GetSubscriptionIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())
		subscriptions = append(subscriptions, subscription)
	}

	return NewGenesisState(params, fees, feeContracts, subscriptions)
}
