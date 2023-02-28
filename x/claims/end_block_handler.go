package claims

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/claims/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {
	// iterator := keeper.GetSubscriptionIterator(ctx)
	// defer iterator.Close()
	// for ; iterator.Valid(); iterator.Next() {
	// 	subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())

	// 	// Skip if should not effect
	// 	if !subscription.ShouldEffect(ctx) {
	// 		continue
	// 	}

	// 	// Effect subscription payment
	// 	err := keeper.EffectSubscriptionPayment(ctx, subscription.Id)
	// 	if err != nil {
	// 		panic(err) // TODO: maybe shouldn't panic?
	// 	}

	// 	// Note: if payment can be re-effected immediately, this should be done
	// 	// in the next block to prevent spending too much time effecting payments

	// 	// Get updated subscription
	// 	subscription, err = keeper.GetSubscription(ctx, subscription.Id)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Delete subscription if it has completed
	// 	if subscription.IsComplete() {
	// 		// TODO: delete subscription
	// 	}

	// 	// Note: no need to save the subscription, as it is being saved by the
	// 	// functions operating on it, such as EffectSubscriptionPayment()
	// }
	return []abci.ValidatorUpdate{}
}
