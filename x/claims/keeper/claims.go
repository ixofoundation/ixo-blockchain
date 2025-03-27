package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/v4/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v4/x/entity/types"
	"github.com/ixofoundation/ixo-blockchain/v4/x/token/types/contracts/cw20"
)

func (k Keeper) SetCollection(ctx sdk.Context, data types.Collection) {
	k.Set(ctx, []byte(data.Id), types.CollectionKey, data, k.Marshal)
}

func (k Keeper) GetCollection(ctx sdk.Context, id string) (types.Collection, error) {
	val, found := k.Get(ctx, []byte(id), types.CollectionKey, k.UnmarshalCollection)
	if !found {
		return types.Collection{}, errorsmod.Wrapf(types.ErrCollectionNotFound, "for %s", id)
	}
	collection, ok := val.(types.Collection)
	if !ok {
		return types.Collection{}, errorsmod.Wrapf(types.ErrCollectionNotFound, "for %s", id)
	}
	return collection, nil
}

func (k Keeper) UnmarshalCollection(value []byte) (interface{}, bool) {
	data := types.Collection{}
	k.Unmarshal(value, &data)
	return data, types.IsValidCollection(&data)
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.Collection:
		bytes = k.cdc.MustMarshal(&value)
	case types.Claim:
		bytes = k.cdc.MustMarshal(&value)
	case types.Dispute:
		bytes = k.cdc.MustMarshal(&value)
	case types.Intent:
		bytes = k.cdc.MustMarshal(&value)
	}
	return
}

// nolint:staticcheck
// Unmarshal unmarshal a byte slice to a struct, return false in case of errors
func (k Keeper) Unmarshal(data []byte, val codec.ProtoMarshaler) bool {
	if len(data) == 0 {
		return false
	}
	if err := k.cdc.Unmarshal(data, val); err != nil {
		return false
	}
	return true
}

func (k Keeper) SetClaim(ctx sdk.Context, data types.Claim) {
	k.Set(ctx, []byte(data.ClaimId), types.ClaimKey, data, k.Marshal)
}

func (k Keeper) GetClaim(ctx sdk.Context, id string) (types.Claim, error) {
	val, found := k.Get(ctx, []byte(id), types.ClaimKey, k.UnmarshalClaim)
	if !found {
		return types.Claim{}, errorsmod.Wrapf(types.ErrClaimNotFound, "for %s", id)
	}
	claim, ok := val.(types.Claim)
	if !ok {
		return types.Claim{}, errorsmod.Wrapf(types.ErrClaimNotFound, "for %s", id)
	}
	return claim, nil
}

func (k Keeper) UnmarshalClaim(value []byte) (interface{}, bool) {
	data := types.Claim{}
	k.Unmarshal(value, &data)
	return data, types.IsValidClaim(&data)
}

func (k Keeper) SetDispute(ctx sdk.Context, data types.Dispute) {
	k.Set(ctx, []byte(data.Data.Proof), types.DisputeKey, data, k.Marshal)
}

func (k Keeper) GetDispute(ctx sdk.Context, proof string) (types.Dispute, error) {
	val, found := k.Get(ctx, []byte(proof), types.DisputeKey, k.UnmarshalDispute)
	if !found {
		return types.Dispute{}, errorsmod.Wrapf(types.ErrDisputeNotFound, "for proof %s", proof)
	}
	dispute, ok := val.(types.Dispute)
	if !ok {
		return types.Dispute{}, errorsmod.Wrapf(types.ErrDisputeNotFound, "for proof %s", proof)
	}
	return dispute, nil
}

func (k Keeper) UnmarshalDispute(value []byte) (interface{}, bool) {
	data := types.Dispute{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDispute(&data)
}

func (k Keeper) GetCollectionsIterator(ctx sdk.Context) storetypes.Iterator {
	return k.GetAll(ctx, types.CollectionKey)
}

func (k Keeper) GetCollections(ctx sdk.Context) []types.Collection {
	iterator := k.GetCollectionsIterator(ctx)
	collections := []types.Collection{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.Collection
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		collections = append(collections, c)
	}

	return collections
}

func (k Keeper) GetClaimsIterator(ctx sdk.Context) storetypes.Iterator {
	return k.GetAll(ctx, types.ClaimKey)
}

func (k Keeper) GetClaims(ctx sdk.Context) []types.Claim {
	iterator := k.GetClaimsIterator(ctx)
	claims := []types.Claim{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.Claim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		claims = append(claims, c)
	}

	return claims
}

func (k Keeper) GetDisputesIterator(ctx sdk.Context) storetypes.Iterator {
	return k.GetAll(ctx, types.DisputeKey)
}

func (k Keeper) GetDisputes(ctx sdk.Context) []types.Dispute {
	iterator := k.GetDisputesIterator(ctx)
	disputes := []types.Dispute{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var d types.Dispute
		k.cdc.MustUnmarshal(iterator.Value(), &d)
		disputes = append(disputes, d)
	}

	return disputes
}

// SetIntent stores the intent in the KV store with the generated key format.
func (k Keeper) SetIntent(ctx sdk.Context, data types.Intent) {
	key := types.IntentKeyCreate(data.AgentAddress, data.CollectionId, data.Id)

	k.Set(ctx, key, types.IntentKey, data, k.Marshal)
}

// GetIntent retrieves an intent from the KV store using the generated key.
func (k Keeper) GetIntent(ctx sdk.Context, agentAddress, collectionId, intentID string) (types.Intent, error) {
	key := types.IntentKeyCreate(agentAddress, collectionId, intentID)

	// Retrieve the intent from the store
	val, found := k.Get(ctx, key, types.IntentKey, k.UnmarshalIntent)
	if !found {
		return types.Intent{}, errorsmod.Wrapf(types.ErrIntentNotFound, "for id %s", intentID)
	}
	intent, ok := val.(types.Intent)
	if !ok {
		return types.Intent{}, errorsmod.Wrapf(types.ErrIntentNotFound, "for id %s", intentID)
	}
	return intent, nil
}

func (k Keeper) GetIntents(ctx sdk.Context) []types.Intent {
	var intents []types.Intent
	iterator := k.GetAll(ctx, types.IntentKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var intent types.Intent
		k.Unmarshal(iterator.Value(), &intent)
		intents = append(intents, intent)
	}
	return intents
}

func (k Keeper) UnmarshalIntent(value []byte) (interface{}, bool) {
	data := types.Intent{}
	k.Unmarshal(value, &data)
	return data, types.IsValidIntent(&data)
}

func (k Keeper) GetAllUserCollectionIntents(ctx sdk.Context, agentAddress, collectionId string) []types.Intent {
	var intents []types.Intent

	prefix := []byte(agentAddress + "/" + collectionId + "/")
	iterator := k.GetAll(ctx, append(types.IntentKey, prefix...))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var intent types.Intent
		k.Unmarshal(iterator.Value(), &intent)
		intents = append(intents, intent)
	}

	return intents
}

// GetActiveIntent retrieves agents current intents for the collection, and returns the first intent if exists
func (k Keeper) GetActiveIntent(ctx sdk.Context, agentAddress, collectionId string) (types.Intent, bool) {
	intents := k.GetAllUserCollectionIntents(ctx, agentAddress, collectionId)
	if len(intents) == 0 {
		return types.Intent{}, false
	}
	k.Logger(ctx).Info("intents", "intent", intents)
	k.Logger(ctx).Info("intent", "intent", intents[0])
	return intents[0], true
}

// RemoveIntentAndEmitEvents removes the intent from the KV store and emits the event as IntentUpdated
// for offchain indexers, Intent status will always change on removal so safe to always emit UpdateIntent event
func (k Keeper) RemoveIntentAndEmitEvents(ctx sdk.Context, intent types.Intent) error {
	// first remove the intent from the KV store
	key := types.IntentKeyCreate(intent.AgentAddress, intent.CollectionId, intent.Id)
	k.Delete(ctx, key, types.IntentKey)

	// then emit events for intent update for offchain indexers
	if err := ctx.EventManager().EmitTypedEvent(
		&types.IntentUpdatedEvent{
			Intent: &intent,
		},
	); err != nil {
		return err
	}
	return nil
}

// TransferCW20Payment transfers CW20 payments to the recipient address.
func (k Keeper) TransferCW20Payment(ctx sdk.Context, fromAddress, toAddress sdk.AccAddress, payment *types.CW20Payment) error {
	// make the payments if amount is not 0
	if payment.Amount == 0 {
		return nil
	}

	encodedTransferMessage, err := cw20.Marshal(cw20.WasmTransfer{
		Transfer: cw20.Transfer{
			Recipient: toAddress.String(),
			Amount:    fmt.Sprint(payment.Amount),
		},
	})
	if err != nil {
		return err
	}

	contractAddress, err := sdk.AccAddressFromBech32(payment.Address)
	if err != nil {
		return err
	}

	_, err = k.WasmKeeper.Execute(
		ctx,
		contractAddress,
		fromAddress,
		encodedTransferMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return err
	}
	return nil
}

// TransferIntentPayments transfers payments, both native coins and CW20 payments, to the recipient address.
func (k Keeper) TransferIntentPayments(ctx sdk.Context, fromAddress, toAddress sdk.AccAddress, amount sdk.Coins, cw20Payments []*types.CW20Payment) error {
	// transfer native coins
	if len(amount) > 0 {
		// clear any Coin with amount 0, generally validation will already block this,
		// but we allow it to know when to use collection defaults or when to have no payments, aka amount 0.
		cleanedAmount := sdk.Coins{}
		for _, coin := range amount {
			if coin.Amount.IsPositive() {
				cleanedAmount = append(cleanedAmount, coin)
			}
		}

		err := k.BankKeeper.SendCoins(ctx, fromAddress, toAddress, cleanedAmount)
		if err != nil {
			return err
		}
	}

	// transfer CW20 payments
	for _, payment := range cw20Payments {
		if payment.Amount != 0 {
			err := k.TransferCW20Payment(ctx, fromAddress, toAddress, payment)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CollectionPersistAndEmitEvents persists the collection and emits the events.
func (k Keeper) CollectionPersistAndEmitEvents(ctx sdk.Context, collection types.Collection) error {
	// persist the Collection
	k.SetCollection(ctx, collection)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.CollectionUpdatedEvent{
			Collection: &collection,
		},
	); err != nil {
		return err
	}

	return nil
}

// RouteGrantEntityAccountAuthz routes the grant entity account authz message to the correct handler.
// It returns an error if the handler is not found or if the message is invalid.
// It emits the events from the message response.
func (k Keeper) RouteGrantEntityAccountAuthz(ctx sdk.Context, msg *entitytypes.MsgGrantEntityAccountAuthz) error {
	// get handler
	handler := k.router.Handler(msg)
	if handler == nil {
		k.Logger(ctx).Error("failed to find grant entity account authz handler")
		return sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(msg))
	}

	// execute handler
	msgResp, err := handler(ctx, msg)
	if err != nil {
		k.Logger(ctx).Error(
			"failed to execute grant entity account authz message",
			"error", err,
			"msg", msg.String(),
		)

		return err
	}

	ctx.EventManager().EmitEvents(msgResp.GetEvents())

	if len(msgResp.MsgResponses) != 1 {
		return errorsmod.Wrapf(
			types.ErrInvalidResponse,
			"expected msg response should be exactly 1, got: %v, responses: %v",
			len(msgResp.MsgResponses), msgResp.MsgResponses,
		)
	}

	return nil
}
