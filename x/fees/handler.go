package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	iterator := keeper.GetSubscriptionIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())
		subContent := subscription.Content

		// Skip if should not charge
		if !subContent.ShouldCharge(ctx) {
			continue
		}

		// Charge subscription fee
		err := keeper.ChargeSubscriptionFee(ctx, subscription.Id)
		if err != nil {
			panic(err) // TODO: maybe shouldn't panic?
		}

		// Note: if fee can be re-charged immediately, this should be done in
		// the next block to prevent spending too much time charging fees

		// Delete subscription if it has ended and no more charges
		if subContent.Ended() && !subContent.ShouldCharge(ctx) {
			// TODO: delete subscription
		}

		// Save subscription
		keeper.SetSubscription(ctx, subscription)
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
			return handleMsgCreateFee(ctx, k, msg)
		case MsgCreateFeeContract:
			return handleMsgCreateFeeContract(ctx, k, msg)
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
	payerAddr := types.DidToAddr(msg.PayerDid)
	if !payerAddr.Equals(feeContract.Content.Payer) {
		return sdk.ErrInvalidAddress("signer must be fee contract payer").Result()
	}

	// Set authorised status
	err = k.SetFeeContractAuthorised(ctx, msg.FeeContractId, msg.Authorised)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgCreateFee(ctx sdk.Context, k Keeper, msg MsgCreateFee) sdk.Result {

	// Submit fee
	_, err := k.SubmitFee(ctx, msg.FeeContent)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgCreateFeeContract(ctx sdk.Context, k Keeper, msg MsgCreateFeeContract) sdk.Result {

	// Confirm that fee exists
	if !k.FeeExists(ctx, msg.FeeId) {
		return sdk.ErrInternal("invalid fee").Result()
	}

	// Create fee contract content
	creatorAddr := types.DidToAddr(msg.CreatorDid)
	feeContract := NewFeeContractContent(
		msg.FeeId, creatorAddr, msg.Payer, msg.CanDeauthorise, false)

	// Submit fee
	_, err := k.SubmitFeeContract(ctx, feeContract)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgGrantFeeDiscount(ctx sdk.Context, k Keeper, msg MsgGrantFeeDiscount) sdk.Result {
	panic("not implemented") // TODO: implement
}

func handleMsgRevokeFeeDiscount(ctx sdk.Context, k Keeper, msg MsgRevokeFeeDiscount) sdk.Result {
	panic("not implemented") // TODO: implement
}

func handleMsgChargeFee(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgChargeFee) sdk.Result {

	// Get fee contract
	feeContract, err := k.GetFeeContract(ctx, msg.FeeContractId)
	if err != nil {
		return err.Result()
	}

	// Confirm that signer is actually the creator of the fee contract
	creatorAddr := types.DidToAddr(msg.FeeContractCreatorDid)
	if !creatorAddr.Equals(feeContract.Content.Creator) {
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
