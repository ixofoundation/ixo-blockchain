package payments

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
)

// InitGenesis new payments genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) {
	// Init payment templates
	for _, pt := range data.PaymentTemplates {
		keeper.SetPaymentTemplate(ctx, pt)
	}

	// Init payment contracts
	for _, pc := range data.PaymentContracts {
		keeper.SetPaymentContract(ctx, pc)
	}

	// Init subscriptions
	for _, s := range data.Subscriptions {
		keeper.SetSubscription(ctx, s)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	// Export payment templates
	var templates []types.PaymentTemplate
	iterator := keeper.GetPaymentTemplateIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		template := keeper.MustGetPaymentTemplateByKey(ctx, iterator.Key())
		templates = append(templates, template)
	}
	iterator.Close()

	// Export payment contracts
	var contracts []types.PaymentContract
	iterator = keeper.GetPaymentContractIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		contract := keeper.MustGetPaymentContractByKey(ctx, iterator.Key())
		contracts = append(contracts, contract)
	}
	iterator.Close()

	// Export subscriptions
	var subscriptions []types.Subscription
	iterator = keeper.GetSubscriptionIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())
		subscriptions = append(subscriptions, subscription)
	}
	iterator.Close()

	return types.NewGenesisState(templates, contracts, subscriptions)
}
