package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new fees genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.SetParams(ctx, data.Params)
	keeper.SetFeeID(ctx, data.StartingFeeId)
	keeper.SetFeeContractID(ctx, data.StartingFeeContractId)
	keeper.SetSubscriptionID(ctx, data.StartingFeeContractId)
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

	// Export fee ID and fee contract ID
	startingFeeID, _ := keeper.GetFeeID(ctx)
	startingFeeContractID, _ := keeper.GetFeeContractID(ctx)
	startingSubscriptionID, _ := keeper.GetSubscriptionID(ctx)

	return NewGenesisState(params, fees, feeContracts, subscriptions,
		startingFeeID, startingFeeContractID, startingSubscriptionID)
}
