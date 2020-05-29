package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

// -------------------------------------------------------- Subscriptions Get/Set

func (k Keeper) GetSubscriptionIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
}

func (k Keeper) MustGetSubscriptionByKey(ctx sdk.Context, key []byte) types.Subscription {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("subscription not found")
	}

	bz := store.Get(key)
	var subscription types.Subscription
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &subscription)

	return subscription
}

func (k Keeper) SubscriptionExists(ctx sdk.Context, subscriptionId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetSubscriptionKey(subscriptionId))
}

func (k Keeper) GetSubscription(ctx sdk.Context, subscriptionId string) (types.Subscription, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubscriptionKey(subscriptionId)

	bz := store.Get(key)
	if bz == nil {
		return types.Subscription{}, sdk.ErrInternal("invalid subscription")
	}

	var subscription types.Subscription
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &subscription)

	return subscription, nil
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubscriptionKey(subscription.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(subscription))
}

// -------------------------------------------------------- Subscriptions Charge

func (k Keeper) ChargeSubscriptionFee(ctx sdk.Context, subscriptionId string) sdk.Error {

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return err
	}

	// Check if should charge
	if !subscription.ShouldCharge(ctx) {
		return types.ErrTriedToChargeSubscriptionFeeWhenShouldnt(types.DefaultCodespace)
	}

	// Charge fee
	charged, err := k.ChargeFee(ctx, k.bankKeeper, subscription.FeeContractId)
	if err != nil {
		return err
	}

	// If the max number of periods has not been reached, then the ?charge?
	// was due to the current period, so we can move to the next period.
	// Otherwise, it means we're tackling accumulated periods. If the fee
	// was charged, then we should deduct an accumulated period
	if !subscription.MaxPeriodsReached() {
		subscription.NextPeriod(charged)
	} else if charged {
		subscription.PeriodsAccumulated =
			subscription.PeriodsAccumulated.Sub(sdk.OneUint())
	}

	k.SetSubscription(ctx, subscription)

	return nil
}
