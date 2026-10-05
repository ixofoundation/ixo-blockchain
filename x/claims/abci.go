package claims

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v8/x/claims/keeper"
	"github.com/ixofoundation/ixo-blockchain/v8/x/claims/types"
)

// NOTE: if performance becomes an issue, we can consider using a similar approach to cosmos sdk grants queue
// for active intents

// EndBlocker is the end blocker function for the claims module
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// Get iterator for active intents
	iterator := k.GetAll(ctx, types.IntentKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var intent types.Intent
		k.Unmarshal(iterator.Value(), &intent)

		// Check if the intent is past its expiration date
		if ctx.BlockTime().After(*intent.ExpireAt) {
			expireIntent := func() {
				// Mark intent as expired
				intent.Status = types.IntentStatus_expired
				if err := k.RemoveIntentAndEmitEvents(ctx, intent); err != nil {
					k.Logger(ctx).Error("failed to remove expired intent", "intent_id", intent.Id, "collection_id", intent.CollectionId, "error", err)
				}
			}

			// Get account used for APPROVAL payments on collection
			fromAddress, err := sdk.AccAddressFromBech32(intent.FromAddress)
			if err != nil {
				k.Logger(ctx).Error("failed to parse expired intent from address", "intent_id", intent.Id, "collection_id", intent.CollectionId, "from_address", intent.FromAddress, "error", err)
				expireIntent()
				continue
			}
			// Get escrow address
			escrow, err := sdk.AccAddressFromBech32(intent.EscrowAddress)
			if err != nil {
				k.Logger(ctx).Error("failed to parse expired intent escrow address", "intent_id", intent.Id, "collection_id", intent.CollectionId, "escrow_address", intent.EscrowAddress, "error", err)
				expireIntent()
				continue
			}

			// Transfer funds back to the original account
			_, err = k.TransferIntentPayments(ctx, escrow, fromAddress, intent.Amount, intent.Cw20Payment, intent.Cw1155Payment, intent.Cw1155IntentPayment)
			if err != nil {
				k.Logger(ctx).Error("failed to refund expired intent payments", "intent_id", intent.Id, "collection_id", intent.CollectionId, "error", err)
			}

			// Restore member budget if this intent was on behalf of a team member
			if intent.MemberAddress != "" {
				if err := k.RestoreMemberBudget(ctx, intent.CollectionId, intent.MemberAddress, intent.Amount, intent.Cw20Payment); err != nil {
					k.Logger(ctx).Error("failed to restore member budget for expired intent", "intent_id", intent.Id, "collection_id", intent.CollectionId, "member_address", intent.MemberAddress, "error", err)
				}
			}

			expireIntent()
		}
	}
}
