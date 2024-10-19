package bonds

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds/keeper"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	iterator := keeper.GetBondIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bond := keeper.MustGetBondByKey(ctx, iterator.Key())
		batch := keeper.MustGetBatch(ctx, bond.BondDid)

		// Subtract one block
		batch.BlocksRemaining = batch.BlocksRemaining.SubUint64(1)
		keeper.SetBatch(ctx, bond.BondDid, batch)

		// If blocks remaining > 0 do not perform orders
		if !batch.BlocksRemaining.IsZero() {
			continue
		}

		// Store current reserve to check if this has changed later on
		reserveBeforeOrderProcessing := bond.CurrentReserve

		// Perform orders
		keeper.PerformOrders(ctx, bond.BondDid)

		// Get bond again just in case current supply was updated
		// Get batch again just in case orders were cancelled
		bond = keeper.MustGetBond(ctx, bond.BondDid)
		batch = keeper.MustGetBatch(ctx, bond.BondDid)

		// For augmented, if hatch phase and newSupply >= S0, go to open phase
		if bond.FunctionType == types.AugmentedFunction &&
			bond.State == types.HatchState.String() {
			args := bond.FunctionParameters.AsMap()
			if bond.CurrentSupply.Amount.ToLegacyDec().GTE(args["S0"]) {
				keeper.SetBondState(ctx, bond.BondDid, types.OpenState.String())
				bond = keeper.MustGetBond(ctx, bond.BondDid) // get updated bond
			}
		}

		// For augmented, if hatch phase and newSupply >= S0, go to open phase
		if bond.FunctionType == types.BondingFunction &&
			bond.State == types.HatchState.String() {
			args := bond.FunctionParameters.AsMap()
			if bond.CurrentSupply.Amount.ToLegacyDec().GTE(args["Hatch_Supply"]) {
				keeper.SetBondState(ctx, bond.BondDid, types.OpenState.String())
				bond = keeper.MustGetBond(ctx, bond.BondDid) // get updated bond
			}
		}

		// Update alpha value if in open state and next alpha is not null
		if bond.State == types.OpenState.String() && batch.HasNextAlpha() {
			if bond.FunctionType == types.BondingFunction {
				keeper.HandleBondingFunctionAlphaUpdate(ctx, bond.BondDid)
			} else {
				keeper.UpdateAlpha(ctx, bond.BondDid)
			}
		}

		// Save current batch as last batch and reset current batch
		keeper.SetLastBatch(ctx, bond.BondDid, batch)
		keeper.SetBatch(ctx, bond.BondDid, types.NewBatch(bond.BondDid, bond.Token, bond.BatchBlocks))

		// If reserve has not changed, no need to recalculate I0; rest of function can be skipped
		if bond.CurrentReserve.Equal(reserveBeforeOrderProcessing) && !batch.HasNextAlpha() {
			continue
		}

		// Recalculate and re-set I0 if alpha bond
		if bond.AlphaBond && bond.FunctionType == types.AugmentedFunction {
			bond = keeper.MustGetBond(ctx, bond.BondDid)
			bondFunctions := bond.FunctionParameters.AsMap()

			// I0 := bondFunctions["I0"]
			currentSystemAlpha := bondFunctions["systemAlpha"]
			S := bond.CurrentSupply.Amount
			//fmt.Println("S: ", S)
			R := bond.CurrentReserve[0].Amount
			C := bond.OutcomePayment
			// Kappa := bondFunctions["kappa"]
			// Recalculate kappa and V0 using new alpha

			// Set new function parameters
			if bond.State == types.OpenState.String() {
				//fmt.Println("Updating I0 -------------------")
				//fmt.Println("Current Supply: ", S)
				//fmt.Println("Current Reserve: ", R)
				//fmt.Println("Current I0: ", I0)
				//fmt.Println("Current Kappa: ", Kappa)
				//fmt.Println("Current SystemAlpha: ", currentSystemAlpha)

				newI0 := types.InvariantI(C, currentSystemAlpha, R)
				//fmt.Println("New I0: ", newI0)

				newKappa := types.Kappa(newI0, C, currentSystemAlpha)
				//fmt.Println("newKappa: ", newKappa)

				newV0, err := types.Invariant(R.ToLegacyDec(), S.ToLegacyDec(), newKappa)
				if err != nil {
					err := ctx.EventManager().EmitTypedEvents(
						&types.BondEditAlphaFailedEvent{
							BondDid:      bond.BondDid,
							Token:        bond.Token,
							CancelReason: err.Error(),
						},
					)
					if err != nil {
						panic(err)
					}
					continue
				}
				//fmt.Println("new V0: ", newV0)
				bond.FunctionParameters.ReplaceParam("V0", newV0)
				bond.FunctionParameters.ReplaceParam("kappa", newKappa)
				bond.FunctionParameters.ReplaceParam("I0", newI0)
			}
		}

		// Save bond
		keeper.SetBond(ctx, bond.BondDid, bond)
	}
}
