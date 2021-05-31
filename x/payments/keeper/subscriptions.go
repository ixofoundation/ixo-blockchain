package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/ixofoundation/ixo-blockchain/x/payments/types"
)

// -------------------------------------------------------- Subscriptions Get/Set

func (k Keeper) GetSubscriptionIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types2.SubscriptionKeyPrefix)
}

func (k Keeper) MustGetSubscriptionByKey(ctx sdk.Context, key []byte) types2.Subscription {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("subscription not found")
	}

	bz := store.Get(key)
	var subscription types2.Subscription
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &subscription)

	return subscription
}

func (k Keeper) SubscriptionExists(ctx sdk.Context, subscriptionId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types2.GetSubscriptionKey(subscriptionId))
}

func (k Keeper) GetSubscription(ctx sdk.Context, subscriptionId string) (types2.Subscription, error) {
	store := ctx.KVStore(k.storeKey)
	key := types2.GetSubscriptionKey(subscriptionId)

	bz := store.Get(key)
	if bz == nil {
		return types2.Subscription{}, fmt.Errorf("invalid subscription")
	}

	var subscription types2.Subscription
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &subscription)

	return subscription, nil
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types2.Subscription) {
	store := ctx.KVStore(k.storeKey)
	key := types2.GetSubscriptionKey(subscription.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(&subscription))
}

// -------------------------------------------------------- Subscriptions Payment

func (k Keeper) EffectSubscriptionPayment(ctx sdk.Context, subscriptionId string) error {

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return err
	}

	// Check if should effect
	if !subscription.ShouldEffect(ctx) {
		return types2.ErrTriedToEffectSubscriptionPaymentWhenShouldnt
	}

	// Effect payment
	effected, err := k.EffectPayment(ctx, subscription.PaymentContractId)
	if err != nil {
		return err
	}

	// If max number of periods has not been reached, then the payment (if any)
	// was due to current period, so we can move to the next period. Otherwise,
	// it means we're tackling accumulated periods, and if payment was actually
	// effected, then we should deduct one from the accumulated periods.
	if !subscription.MaxPeriodsReached() {
		subscription.NextPeriod(effected)
	} else if effected {
		subscription.PeriodsAccumulated =
			subscription.PeriodsAccumulated.Sub(sdk.OneUint())
	}

	// If the payment was not effected (assuming err == nil) this is because (i)
	// the payer does not have enough coins, or (ii) because the payment *cannot*
	// be effected (i.e. maximum payment reached, contract not authorised, etc.)

	// Update subscription
	k.SetSubscription(ctx, subscription)

	return nil
}
