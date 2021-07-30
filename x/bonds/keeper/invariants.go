package keeper

// DONTCOVER

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
)

// RegisterInvariants registers all supply invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "bonds-supply",
		SupplyInvariant(k))
	ir.RegisterRoute(types.ModuleName, "bonds-reserve",
		ReserveInvariant(k))
	ir.RegisterRoute(types.ModuleName, "bonds-available-reserve",
		AvailableReserveInvariant(k))
}

// AllInvariants runs all invariants of the bonds module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := SupplyInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = ReserveInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return AvailableReserveInvariant(k)(ctx)
	}
}

func SupplyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		var count int

		// Get supply of coins held in accounts (includes stake token)
		supplyInAccounts := sdk.Coins{}
		k.BankKeeper.IterateAllBalances(ctx, func(_ sdk.AccAddress, balance sdk.Coin) bool {
			supplyInAccounts = supplyInAccounts.Add(balance)
			return false
		})

		iterator := k.GetBondIterator(ctx)
		for ; iterator.Valid(); iterator.Next() {
			bond := k.MustGetBondByKey(ctx, iterator.Key())
			denom := bond.Token
			batch := k.MustGetBatch(ctx, bond.BondDid)
			did := bond.BondDid

			// Add bond current supply
			supplyInBondsAndBatches := bond.CurrentSupply

			// Subtract amount to be burned (this amount was already burned
			// in handleMsgSell but is still a part of bond's CurrentSupply)
			for _, s := range batch.Sells {
				if !s.BaseOrder.Cancelled {
					supplyInBondsAndBatches = supplyInBondsAndBatches.Sub(
						s.BaseOrder.Amount)
				}
			}

			// Check that amount matches supply in accounts
			inAccounts := supplyInAccounts.AmountOf(bond.Token)
			if !supplyInBondsAndBatches.Amount.Equal(inAccounts) {
				count++
				msg += fmt.Sprintf("total %s supply invariance:\n"+
					"\ttotal %s supply: %s\n"+
					"\tsum of %s in accounts: %s\n",
					did, denom, supplyInBondsAndBatches.Amount.String(),
					denom, inAccounts.String())
			}
		}

		broken := count != 0
		return sdk.FormatInvariant(types.ModuleName, "supply", fmt.Sprintf(
			"%d Bonds supply invariants broken\n%s", count, msg)), broken
	}
}

func ReserveInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		var count int

		iterator := k.GetBondIterator(ctx)
		for ; iterator.Valid(); iterator.Next() {
			bond := k.MustGetBondByKey(ctx, iterator.Key())
			denom := bond.Token
			did := bond.BondDid

			if bond.FunctionType == types.AugmentedFunction ||
				bond.FunctionType == types.SwapperFunction {
				continue // Check does not apply to augmented/swapper functions
			}

			expectedReserve, err := bond.ReserveAtSupply(bond.CurrentSupply.Amount)
			if err != nil {
				continue // ignore error
			}
			expectedRounded := expectedReserve.Ceil().TruncateInt()
			actualReserve := k.GetReserveBalances(ctx, did)

			for _, r := range actualReserve {
				if r.Amount.LT(expectedRounded) {
					count++
					msg += fmt.Sprintf("%s reserve invariance:\n"+
						"\texpected(ceil-rounded) %s reserve: %s\n"+
						"\tactual %s reserve: %s\n",
						did, denom, expectedReserve.String(),
						denom, r.String())
				}
			}
		}

		broken := count != 0
		return sdk.FormatInvariant(types.ModuleName, "reserve", fmt.Sprintf(
			"%d Bonds reserve invariants broken\n%s", count, msg)), broken
	}
}

func AvailableReserveInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {

		// Get actual available reserve
		reservePool := k.accountKeeper.GetModuleAddress(types.BondsReserveAccount)
		actualAvailableReserve := k.BankKeeper.GetAllBalances(ctx, reservePool)

		// If no bonds (iterator invalid) then invariant automatically holds
		iterator := k.GetBondIterator(ctx)
		if !iterator.Valid() {
			return "", false
		}

		// Calculate sum of available reserves reported by bonds, including any
		// outcome payment reserve, since this has already reached the reserve
		// but is not considered a part of the available reserve
		availableReserveSum := sdk.NewCoins()
		for ; iterator.Valid(); iterator.Next() {
			bond := k.MustGetBondByKey(ctx, iterator.Key())
			availableReserveSum = availableReserveSum.Add(bond.AvailableReserve...)
			availableReserveSum = availableReserveSum.Add(bond.CurrentOutcomePaymentReserve...)
		}

		broken := !availableReserveSum.IsEqual(actualAvailableReserve)
		return sdk.FormatInvariant(types.ModuleName, "available-reserve",
			"Bonds available reserve invariant broken"), broken
	}
}
