package fees

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	iterator := keeper.GetSubscriptionIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())

		// Skip if should not charge
		if !subscription.ShouldCharge(ctx) {
			continue
		}

		// Charge subscription fee
		err := keeper.ChargeSubscriptionFee(ctx, subscription.Id)
		if err != nil {
			panic(err) // TODO: maybe shouldn't panic?
		}

		// Note: if fee can be re-charged immediately, this should be done in
		// the next block to prevent spending too much time charging fees

		// Get updated subscription
		subscription, err = keeper.GetSubscription(ctx, subscription.Id)
		if err != nil {
			panic(err)
		}

		// Delete subscription if it has completed
		if subscription.IsComplete() {
			// TODO: delete subscription
		}

		// Note: no need to save the subscription, as it is being saved by the
		// functions operating on it, such as ChargeSubscriptionFee()
	}
	return []abci.ValidatorUpdate{}
}

func NewHandler(k Keeper, bk bank.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgSetFeeContractAuthorisation:
			return handleMsgSetFeeContractAuthorisation(ctx, k, msg)
		case MsgCreateFee:
			return handleMsgCreateFee(ctx, k, bk, msg)
		case MsgCreateFeeContract:
			return handleMsgCreateFeeContract(ctx, k, bk, msg)
		case MsgCreateSubscription:
			return handleMsgCreateSubscription(ctx, k, msg)
		case MsgGrantFeeDiscount:
			return handleMsgGrantFeeDiscount(ctx, k, msg)
		case MsgRevokeFeeDiscount:
			return handleMsgRevokeFeeDiscount(ctx, k, msg)
		case MsgChargeFee:
			return handleMsgChargeFee(ctx, k, bk, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgSetFeeContractAuthorisation(ctx sdk.Context, k Keeper, msg MsgSetFeeContractAuthorisation) sdk.Result {

	// Get fee contract
	feeContract, err := k.GetFeeContract(ctx, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Confirm that signer is actually the payer in the fee contract
	payerAddr := ixo.DidToAddr(msg.PayerDid)
	if !payerAddr.Equals(feeContract.Payer) {
		return sdk.ErrInvalidAddress("signer must be fee contract payer").Result()
	}

	// Set authorised status
	err = k.SetFeeContractAuthorised(ctx, msg.FeeContractId, msg.Authorised)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgCreateFee(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgCreateFee) sdk.Result {

	// Ensure that fee doesn't already exist
	if k.FeeExists(ctx, msg.Fee.Id) {
		return types.ErrAlreadyExists(types.DefaultCodespace, fmt.Sprintf(
			"fee '%s' already exists", msg.Fee.Id)).Result()
	}

	// Ensure that fee ID is not reserved
	if k.FeeIdReserved(msg.Fee.Id) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed as it is "+
			"using a reserved prefix", msg.Fee.Id)).Result()
	}

	// Create and validate fee
	if err := msg.Fee.Validate(); err != nil {
		return err.Result()
	}

	// Ensure no blacklisted address in wallet distribution
	for _, share := range msg.Fee.WalletDistribution {
		if bk.BlacklistedAddr(share.Address) {
			return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed "+
				"to receive transactions", share.Address)).Result()
		}
	}

	// Submit fee
	k.SetFee(ctx, msg.Fee)

	return sdk.Result{}
}

func handleMsgCreateFeeContract(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgCreateFeeContract) sdk.Result {

	// Ensure that fee contract doesn't already exist
	if k.FeeContractExists(ctx, msg.FeeContractId) {
		return types.ErrAlreadyExists(types.DefaultCodespace, fmt.Sprintf(
			"fee contract '%s' already exists", msg.FeeContractId)).Result()
	}

	// Ensure that fee contract ID is not reserved
	if k.FeeContractIdReserved(msg.FeeContractId) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed as it is "+
			"using a reserved prefix", msg.FeeContractId)).Result()
	}

	// Ensure payer is not a blacklisted address
	if bk.BlacklistedAddr(msg.Payer) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed "+
			"to receive transactions", msg.Payer)).Result()
	}

	// Confirm that fee exists
	if !k.FeeExists(ctx, msg.FeeId) {
		return sdk.ErrInternal("invalid fee").Result()
	}

	// Create fee contract and validate
	creatorAddr := ixo.DidToAddr(msg.CreatorDid)
	feeContract := NewFeeContract(msg.FeeContractId, msg.FeeId,
		creatorAddr, msg.Payer, msg.CanDeauthorise, false, msg.DiscountId)
	if err := feeContract.Validate(); err != nil {
		return err.Result()
	}

	// Submit fee contract
	k.SetFeeContract(ctx, feeContract)

	return sdk.Result{}
}

func handleMsgCreateSubscription(ctx sdk.Context, k Keeper, msg MsgCreateSubscription) sdk.Result {

	// Ensure that subscription doesn't already exist
	if k.SubscriptionExists(ctx, msg.SubscriptionId) {
		return types.ErrAlreadyExists(types.DefaultCodespace, fmt.Sprintf(
			"subscription '%s' already exists", msg.SubscriptionId)).Result()
	}

	// Ensure that subscription ID is not reserved
	if k.SubscriptionIdReserved(msg.SubscriptionId) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed as it is "+
			"using a reserved prefix", msg.SubscriptionId)).Result()
	}

	// Get fee contract
	feeContract, err := k.GetFeeContract(ctx, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Confirm that signer is actually the creator of the fee contract
	creatorAddr := ixo.DidToAddr(msg.CreatorDid)
	if !creatorAddr.Equals(feeContract.Creator) {
		return sdk.ErrInvalidAddress("signer must be fee contract creator").Result()
	}

	// Create subscription and validate
	subscription := NewSubscription(msg.SubscriptionId,
		msg.FeeContractId, msg.MaxPeriods, msg.Period)
	if err := subscription.Validate(); err != nil {
		return err.Result()
	}

	// Submit subscription
	k.SetSubscription(ctx, subscription)

	return sdk.Result{}
}

func handleMsgGrantFeeDiscount(ctx sdk.Context, k Keeper, msg MsgGrantFeeDiscount) sdk.Result {

	// Get FeeContract
	feeContract, err := k.GetFeeContract(ctx, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Confirm that signer is actually the creator of the fee contract
	creatorAddr := ixo.DidToAddr(msg.SenderDid)
	if !creatorAddr.Equals(feeContract.Creator) {
		return sdk.ErrInvalidAddress("signer must be fee contract creator").Result()
	}

	// Confirm that discount ID is in the fee (to avoid invalid discount IDs)
	found, err := k.DiscountIdExists(ctx, feeContract.FeeId, msg.DiscountId)
	if err != nil {
		return err.Result()
	} else if !found {
		return types.ErrInvalidId(types.DefaultCodespace, "discount ID not in fee's discount list").Result()
	}

	// Grant the fee discount
	err = k.GrantFeeDiscount(ctx, feeContract.Id, msg.DiscountId)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgRevokeFeeDiscount(ctx sdk.Context, k Keeper, msg MsgRevokeFeeDiscount) sdk.Result {

	// Get FeeContract
	feeContract, err := k.GetFeeContract(ctx, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Confirm that signer is actually the creator of the fee contract
	creatorAddr := ixo.DidToAddr(msg.SenderDid)
	if !creatorAddr.Equals(feeContract.Creator) {
		return sdk.ErrInvalidAddress("signer must be fee contract creator").Result()
	}

	// Revoke the fee discount
	err = k.RevokeFeeDiscount(ctx, feeContract.Id)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgChargeFee(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgChargeFee) sdk.Result {

	// Get fee contract
	feeContract, err := k.GetFeeContract(ctx, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Confirm that signer is actually the creator of the fee contract
	creatorAddr := ixo.DidToAddr(msg.SenderDid)
	if !creatorAddr.Equals(feeContract.Creator) {
		return sdk.ErrInvalidAddress("signer must be fee contract creator").Result()
	}

	// Charge fee
	charged, err := k.ChargeFee(ctx, bk, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Fee not charged but no error, meaning that fee should have been charged
	if !charged {
		return sdk.ErrInternal("fee not charged due to unknown rason").Result()
	}

	return sdk.Result{}
}
