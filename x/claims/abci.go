package claims

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/claims/keeper"
	"github.com/ixofoundation/ixo-blockchain/v5/x/claims/types"
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
			// Get account used for APPROVAL payments on collection
			fromAddress, err := sdk.AccAddressFromBech32(intent.FromAddress)
			if err != nil {
				panic(err)
			}
			// Get escrow address
			escrow, err := sdk.AccAddressFromBech32(intent.EscrowAddress)
			if err != nil {
				panic(err)
			}

			// Transfer funds back to the original account
			err = k.TransferIntentPayments(ctx, escrow, fromAddress, intent.Amount, intent.Cw20Payment)
			if err != nil {
				// if this happens then it means there is funds missing in escrow account, should never happen
				panic(err)
			}

			// Mark intent as expired
			intent.Status = types.IntentStatus_expired
			err = k.RemoveIntentAndEmitEvents(ctx, intent)
			if err != nil {
				panic(err)
			}
		}
	}
}
