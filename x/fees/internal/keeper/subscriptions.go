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

func (k Keeper) SubscriptionExists(ctx sdk.Context, subscriptionId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetSubscriptionKey(subscriptionId))
}

func (k Keeper) GetSubscription(ctx sdk.Context, subscriptionId uint64) (types.Subscription, sdk.Error) {
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

func (k Keeper) SubmitSubscription(ctx sdk.Context, content types.SubscriptionContent) (types.Subscription, sdk.Error) {
	subscriptionId, err := k.GetSubscriptionID(ctx)
	if err != nil {
		return types.Subscription{}, err
	}

	subscription := types.NewSubscription(subscriptionId, content)

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionID(ctx, subscriptionId+1)

	return subscription, nil
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubscriptionKey(subscription.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(subscription))
}

// GetSubscriptionID gets the highest subscription ID
func (k Keeper) GetSubscriptionID(ctx sdk.Context) (subscriptionId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SubscriptionIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial subscription ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &subscriptionId)
	return subscriptionId, nil
}

// Set the subscription ID
func (k Keeper) SetSubscriptionID(ctx sdk.Context, subscriptionId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(subscriptionId)
	store.Set(types.SubscriptionIdKey, bz)
}

// -------------------------------------------------------- Subscriptions Charge

func (k Keeper) ChargeSubscriptionFee(ctx sdk.Context, subscriptionId uint64) sdk.Error {

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return err
	}
	sData := subscription.Content

	// Check if should charge
	if !sData.ShouldCharge(ctx) {
		return types.ErrTriedToChargeSubscriptionFeeWhenShouldnt(types.DefaultCodespace)
	}

	// Charge fee
	charged, err := k.ChargeFee(ctx, k.bankKeeper, sData.GetFeeContractId())
	if err != nil {
		return err
	}

	// Update and save subscription
	sData.NextPeriod(charged)
	k.SetSubscription(ctx, subscription)

	return nil
}
