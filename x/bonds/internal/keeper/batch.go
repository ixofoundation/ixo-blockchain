package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func (k Keeper) MustGetBatch(ctx sdk.Context, bondDid ixo.Did) types.Batch {
	store := ctx.KVStore(k.storeKey)
	if !k.BatchExists(ctx, bondDid) {
		panic(fmt.Sprintf("batch not found for %s\n", bondDid))
	}
	bz := store.Get(types.GetBatchKey(bondDid))
	var batch types.Batch
	k.cdc.MustUnmarshalBinaryBare(bz, &batch)
	return batch
}

func (k Keeper) MustGetLastBatch(ctx sdk.Context, bondDid ixo.Did) types.Batch {
	store := ctx.KVStore(k.storeKey)
	if !k.LastBatchExists(ctx, bondDid) {
		panic(fmt.Sprintf("last batch not found for %s\n", bondDid))
	}
	bz := store.Get(types.GetLastBatchKey(bondDid))
	var batch types.Batch
	k.cdc.MustUnmarshalBinaryBare(bz, &batch)
	return batch
}

func (k Keeper) BatchExists(ctx sdk.Context, bondDid ixo.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetBatchKey(bondDid))
}

func (k Keeper) LastBatchExists(ctx sdk.Context, bondDid ixo.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetLastBatchKey(bondDid))
}

func (k Keeper) SetBatch(ctx sdk.Context, bondDid ixo.Did, batch types.Batch) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBatchKey(bondDid), k.cdc.MustMarshalBinaryBare(batch))
}

func (k Keeper) SetLastBatch(ctx sdk.Context, bondDid ixo.Did, batch types.Batch) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetLastBatchKey(bondDid), k.cdc.MustMarshalBinaryBare(batch))
}

func (k Keeper) AddBuyOrder(ctx sdk.Context, bondDid ixo.Did, bo types.BuyOrder, buyPrices, sellPrices sdk.DecCoins) {
	batch := k.MustGetBatch(ctx, bondDid)
	batch.TotalBuyAmount = batch.TotalBuyAmount.Add(bo.Amount)
	batch.BuyPrices = buyPrices
	batch.SellPrices = sellPrices
	batch.Buys = append(batch.Buys, bo)
	k.SetBatch(ctx, bondDid, batch)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("added buy order for %s from %s", bo.Amount.String(), bo.Address.String()))
}

func (k Keeper) AddSellOrder(ctx sdk.Context, bondDid ixo.Did, so types.SellOrder, buyPrices, sellPrices sdk.DecCoins) {
	batch := k.MustGetBatch(ctx, bondDid)
	batch.TotalSellAmount = batch.TotalSellAmount.Add(so.Amount)
	batch.BuyPrices = buyPrices
	batch.SellPrices = sellPrices
	batch.Sells = append(batch.Sells, so)
	k.SetBatch(ctx, bondDid, batch)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("added sell order for %s from %s", so.Amount.String(), so.Address.String()))
}

func (k Keeper) AddSwapOrder(ctx sdk.Context, bondDid ixo.Did, so types.SwapOrder) {
	batch := k.MustGetBatch(ctx, bondDid)
	batch.Swaps = append(batch.Swaps, so)
	k.SetBatch(ctx, bondDid, batch)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("added swap order for %s to %s from %s", so.Amount.String(), so.ToToken, so.Address.String()))
}

func (k Keeper) GetBatchBuySellPrices(ctx sdk.Context, bondDid string, batch types.Batch) (buyPricesPT, sellPricesPT sdk.DecCoins, err sdk.Error) {
	bond := k.MustGetBond(ctx, bondDid)

	buyAmountDec := sdk.NewDecFromInt(batch.TotalBuyAmount.Amount)
	sellAmountDec := sdk.NewDecFromInt(batch.TotalSellAmount.Amount)

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	currentPricesPT, err := bond.GetCurrentPricesPT(reserveBalances)
	if err != nil {
		return nil, nil, err
	}

	// Get (amount of) matched and (actual) curve-calculated value for the remaining amount
	// - The matched amount is the least of the buys and sells (i.e. greatest common amount)
	// - The curved values are the prices/returns for the extra unmatched buys/sells
	var matchedAmount sdk.Dec
	var curvedValues sdk.DecCoins
	if batch.EqualBuysAndSells() {
		// Since equal, both prices are current prices
		return currentPricesPT, currentPricesPT, nil
	} else if batch.MoreBuysThanSells() {
		matchedAmount = sellAmountDec // since sells < buys, greatest common amount is sells
		extraBuys := batch.TotalBuyAmount.Sub(batch.TotalSellAmount)
		curvedValues, err = bond.GetPricesToMint(extraBuys.Amount, reserveBalances) // buy prices
		if err != nil {
			return nil, nil, err
		}
	} else {
		matchedAmount = buyAmountDec // since buys < sells, greatest common amount is buys
		extraSells := batch.TotalSellAmount.Sub(batch.TotalBuyAmount)
		curvedValues = bond.GetReturnsForBurn(extraSells.Amount, reserveBalances) // sell returns
	}

	// Get (actual) matched values
	matchedValues := types.MultiplyDecCoinsByDec(currentPricesPT, matchedAmount)

	// If buys > sells, totalValues is the total buy prices
	// If sells > buys, totalValues is the total sell returns
	totalValues := matchedValues.Add(curvedValues)

	// Calculate buy and sell prices per token
	if batch.MoreBuysThanSells() {
		buyPricesPT = types.DivideDecCoinsByDec(totalValues, buyAmountDec)
		sellPricesPT = currentPricesPT
	} else {
		buyPricesPT = currentPricesPT
		sellPricesPT = types.DivideDecCoinsByDec(totalValues, sellAmountDec)
	}
	return buyPricesPT, sellPricesPT, nil
}

func (k Keeper) GetUpdatedBatchPricesAfterBuy(ctx sdk.Context, bondDid ixo.Did, bo types.BuyOrder) (buyPrices, sellPrices sdk.DecCoins, err sdk.Error) {
	bond := k.MustGetBond(ctx, bondDid)
	batch := k.MustGetBatch(ctx, bondDid)

	// Max supply cannot be less than supply (max supply >= supply)
	adjustedSupply := k.GetSupplyAdjustedForBuy(ctx, bondDid)
	if bond.MaxSupply.IsLT(adjustedSupply.Add(bo.Amount)) {
		return nil, nil, types.ErrCannotMintMoreThanMaxSupply(types.DefaultCodespace)
	}

	// Simulate buy by bumping up total buy amount
	batch.TotalBuyAmount = batch.TotalBuyAmount.Add(bo.Amount)
	buyPrices, sellPrices, err = k.GetBatchBuySellPrices(ctx, bondDid, batch)
	if err != nil {
		return nil, nil, err
	}

	err = k.CheckIfBuyOrderFulfillableAtPrice(ctx, bondDid, bo, buyPrices)
	if err != nil {
		return nil, nil, err
	}

	return buyPrices, sellPrices, nil
}

func (k Keeper) GetUpdatedBatchPricesAfterSell(ctx sdk.Context, bondDid ixo.Did, so types.SellOrder) (buyPrices, sellPrices sdk.DecCoins, err sdk.Error) {
	batch := k.MustGetBatch(ctx, bondDid)

	// Cannot burn more tokens than what exists
	adjustedSupply := k.GetSupplyAdjustedForSell(ctx, bondDid)
	if adjustedSupply.IsLT(so.Amount) {
		return nil, nil, types.ErrCannotBurnMoreThanSupply(types.DefaultCodespace)
	}

	// Simulate sell by bumping up total sell amount
	batch.TotalSellAmount = batch.TotalSellAmount.Add(so.Amount)
	buyPrices, sellPrices, err = k.GetBatchBuySellPrices(ctx, bondDid, batch)
	if err != nil {
		return nil, nil, err
	}

	return buyPrices, sellPrices, nil
}

func (k Keeper) PerformBuyAtPrice(ctx sdk.Context, bondDid ixo.Did, bo types.BuyOrder, prices sdk.DecCoins) (err sdk.Error) {
	bond := k.MustGetBond(ctx, bondDid)

	// Mint bond tokens
	err = k.SupplyKeeper.MintCoins(ctx, types.BondsMintBurnAccount,
		sdk.Coins{bo.Amount})
	if err != nil {
		return err
	}

	// Send bond tokens bought to buyer
	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
		types.BondsMintBurnAccount, bo.Address, sdk.Coins{bo.Amount})
	if err != nil {
		return err
	}

	reservePrices := types.MultiplyDecCoinsByInt(prices, bo.Amount.Amount)
	reservePricesRounded := types.RoundReservePrices(reservePrices)
	txFees := bond.GetTxFees(reservePrices)
	totalPrices := reservePricesRounded.Add(txFees)

	if totalPrices.IsAnyGT(bo.MaxPrices) {
		return types.ErrMaxPriceExceeded(types.DefaultCodespace, totalPrices, bo.MaxPrices)
	}

	// Add new reserve to reserve address (reservePricesRounded should never be zero)
	// TODO: investigate possibility of zero reservePricesRounded
	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
		types.BatchesIntermediaryAccount, bond.ReserveAddress, reservePricesRounded)
	if err != nil {
		return err
	}

	// Add charged fee to fee address
	if !txFees.IsZero() {
		err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, bond.FeeAddress, txFees)
		if err != nil {
			return err
		}
	}

	// Add remainder to buyer address
	returnToBuyer := bo.MaxPrices.Sub(totalPrices)
	if !returnToBuyer.IsZero() {
		err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, bo.Address, returnToBuyer)
		if err != nil {
			return err
		}
	}

	// Update supply (max supply exceeded check done during MsgBuy)
	k.SetCurrentSupply(ctx, bondDid, bond.CurrentSupply.Add(bo.Amount))

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("performed buy order for %s from %s", bo.Amount.String(), bo.Address.String()))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeOrderFulfill,
		sdk.NewAttribute(types.AttributeKeyBond, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueBuyOrder),
		sdk.NewAttribute(types.AttributeKeyAddress, bo.Address.String()),
		sdk.NewAttribute(types.AttributeKeyTokensMinted, bo.Amount.Amount.String()),
		sdk.NewAttribute(types.AttributeKeyChargedPrices, reservePricesRounded.String()),
		sdk.NewAttribute(types.AttributeKeyChargedFees, txFees.String()),
		sdk.NewAttribute(types.AttributeKeyReturnedToAddress, returnToBuyer.String()),
	))

	return nil
}

func (k Keeper) PerformSellAtPrice(ctx sdk.Context, bondDid ixo.Did, so types.SellOrder, prices sdk.DecCoins) (err sdk.Error) {
	bond := k.MustGetBond(ctx, bondDid)

	reserveReturns := types.MultiplyDecCoinsByInt(prices, so.Amount.Amount)
	reserveReturnsRounded := types.RoundReserveReturns(reserveReturns)
	txFees := bond.GetTxFees(reserveReturns)
	exitFees := bond.GetExitFees(reserveReturns)

	totalFees := types.AdjustFees(txFees.Add(exitFees), reserveReturnsRounded) // calculate actual total fees
	totalReturns := reserveReturnsRounded.Sub(totalFees)                       // calculate actual reserveReturns

	// Send total returns to seller (totalReturns should never be zero)
	// TODO: investigate possibility of zero totalReturns
	err = k.CoinKeeper.SendCoins(ctx, bond.ReserveAddress, so.Address, totalReturns)
	if err != nil {
		return err
	}

	// Send total fee to fee address
	if !totalFees.IsZero() {
		err := k.CoinKeeper.SendCoins(ctx, bond.ReserveAddress, bond.FeeAddress, totalFees)
		if err != nil {
			return err
		}
	}

	// Update supply (burn more than supply check done during MsgSell)
	k.SetCurrentSupply(ctx, bondDid, bond.CurrentSupply.Sub(so.Amount))

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("performed sell order for %s from %s", so.Amount.String(), so.Address.String()))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeOrderFulfill,
		sdk.NewAttribute(types.AttributeKeyBond, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueSellOrder),
		sdk.NewAttribute(types.AttributeKeyAddress, so.Address.String()),
		sdk.NewAttribute(types.AttributeKeyTokensBurned, so.Amount.Amount.String()),
		sdk.NewAttribute(types.AttributeKeyChargedFees, txFees.String()),
		sdk.NewAttribute(types.AttributeKeyReturnedToAddress, totalReturns.String()),
	))

	return nil
}

func (k Keeper) PerformSwap(ctx sdk.Context, bondDid ixo.Did, so types.SwapOrder) (err sdk.Error, ok bool) {
	bond := k.MustGetBond(ctx, bondDid)

	// WARNING: do not return ok=true if money has already been transferred when error occurs

	// Get return for swap
	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(so.Amount, so.ToToken, reserveBalances)
	if err != nil {
		return err, true
	}
	adjustedInput := so.Amount.Sub(txFee) // same as during GetReturnsForSwap

	// Check if new rates violate sanity rate
	newReserveBalances := reserveBalances.Add(sdk.Coins{adjustedInput}).Sub(reserveReturns)
	if bond.ReservesViolateSanityRate(newReserveBalances) {
		return types.ErrValuesViolateSanityRate(types.DefaultCodespace), true
	}

	// Give resultant tokens to swapper (reserveReturns should never be zero)
	err = k.CoinKeeper.SendCoins(ctx, bond.ReserveAddress, so.Address, reserveReturns)
	if err != nil {
		return err, false
	}

	// Add fee-reduced coins to be swapped to reserve (adjustedInput should never be zero)
	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
		types.BatchesIntermediaryAccount, bond.ReserveAddress, sdk.Coins{adjustedInput})
	if err != nil {
		return err, false
	}

	// Add fee (taken from swapper) to fee address
	if !txFee.IsZero() {
		err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, bond.FeeAddress, sdk.Coins{txFee})
		if err != nil {
			return err, false
		}
	}

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("performed swap order for %s to %s from %s",
		so.Amount.String(), reserveReturns, so.Address.String()))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeOrderFulfill,
		sdk.NewAttribute(types.AttributeKeyBond, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueSwapOrder),
		sdk.NewAttribute(types.AttributeKeyAddress, so.Address.String()),
		sdk.NewAttribute(types.AttributeKeyTokensSwapped, adjustedInput.String()),
		sdk.NewAttribute(types.AttributeKeyChargedFees, txFee.String()),
		sdk.NewAttribute(types.AttributeKeyReturnedToAddress, reserveReturns.String()),
	))

	return nil, true
}

func (k Keeper) PerformBuyOrders(ctx sdk.Context, bondDid ixo.Did) {
	batch := k.MustGetBatch(ctx, bondDid)

	// Perform buys or return to buyer
	for _, bo := range batch.Buys {
		if !bo.IsCancelled() {
			err := k.PerformBuyAtPrice(ctx, bondDid, bo, batch.BuyPrices)
			if err != nil {
				// Panic here since all calculations should have been done
				// correctly to prevent any errors during the buy
				panic(err)
			}
		}
	}

	// Update batch with any new changes (shouldn't be any)
	k.SetBatch(ctx, bondDid, batch)
}

func (k Keeper) PerformSellOrders(ctx sdk.Context, bondDid ixo.Did) {
	batch := k.MustGetBatch(ctx, bondDid)

	// Perform sells or return to seller
	for _, so := range batch.Sells {
		if !so.IsCancelled() {
			err := k.PerformSellAtPrice(ctx, bondDid, so, batch.SellPrices)
			if err != nil {
				// Panic here since all calculations should have been done
				// correctly to prevent any errors during the sell
				panic(err)
			}
		}
	}

	// Update batch with any new changes (shouldn't be any)
	k.SetBatch(ctx, bondDid, batch)
}

func (k Keeper) PerformSwapOrders(ctx sdk.Context, bondDid ixo.Did) {
	logger := ctx.Logger()
	batch := k.MustGetBatch(ctx, bondDid)

	// Perform swaps
	// TODO: implement swaps front-running prevention
	for i, so := range batch.Swaps {
		if !so.IsCancelled() {
			err, ok := k.PerformSwap(ctx, bondDid, so)
			if err != nil {
				if ok {
					batch.Swaps[i].Cancelled = types.TRUE
					batch.Swaps[i].CancelReason = err.Error()

					logger.Info(fmt.Sprintf("cancelled swap order for %s to %s from %s", so.Amount.String(), so.ToToken, so.Address.String()))
					logger.Debug(fmt.Sprintf("cancellation reason: %s", err.Error()))

					// Return from amount to swapper
					err := k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
						types.BatchesIntermediaryAccount, so.Address, sdk.Coins{so.Amount})
					if err != nil {
						panic(err)
					}
				} else {
					// Panic here since all calculations should have been done
					// correctly to prevent any errors during the swap
					panic(err)
				}
			}
		}
	}

	// Update batch with any new cancellations
	k.SetBatch(ctx, bondDid, batch)
}

func (k Keeper) PerformOrders(ctx sdk.Context, bondDid ixo.Did) {
	k.PerformBuyOrders(ctx, bondDid)
	k.PerformSellOrders(ctx, bondDid)
	k.PerformSwapOrders(ctx, bondDid)
}

func (k Keeper) CheckIfBuyOrderFulfillableAtPrice(ctx sdk.Context, bondDid ixo.Did, bo types.BuyOrder, prices sdk.DecCoins) sdk.Error {
	bond := k.MustGetBond(ctx, bondDid)

	reservePrices := types.MultiplyDecCoinsByInt(prices, bo.Amount.Amount)
	reserveRounded := types.RoundReservePrices(reservePrices)
	txFees := bond.GetTxFees(reservePrices)
	totalPrices := reserveRounded.Add(txFees)

	// Check that max prices not exceeded
	if totalPrices.IsAnyGT(bo.MaxPrices) {
		return types.ErrMaxPriceExceeded(types.DefaultCodespace, totalPrices, bo.MaxPrices)
	}

	return nil
}

func (k Keeper) CancelUnfulfillableBuys(ctx sdk.Context, bondDid ixo.Did) (cancelledOrders int) {
	logger := k.Logger(ctx)
	batch := k.MustGetBatch(ctx, bondDid)

	// Cancel unfulfillable buys
	for i, bo := range batch.Buys {
		if !bo.IsCancelled() {
			err := k.CheckIfBuyOrderFulfillableAtPrice(ctx, bondDid, bo, batch.BuyPrices)
			if err != nil {
				// Cancel (important to use batch.Buys[i] and not bo!)
				batch.Buys[i].Cancelled = types.TRUE
				batch.Buys[i].CancelReason = err.Error()
				batch.TotalBuyAmount = batch.TotalBuyAmount.Sub(bo.Amount)
				cancelledOrders += 1

				logger.Info(fmt.Sprintf("cancelled buy order for %s from %s", bo.Amount.String(), bo.Address.String()))
				logger.Debug(fmt.Sprintf("cancellation reason: %s", err.Error()))

				ctx.EventManager().EmitEvent(sdk.NewEvent(
					types.EventTypeOrderCancel,
					sdk.NewAttribute(types.AttributeKeyBond, bondDid),
					sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueBuyOrder),
					sdk.NewAttribute(types.AttributeKeyAddress, bo.Address.String()),
					sdk.NewAttribute(types.AttributeKeyCancelReason, bo.CancelReason),
				))

				// Return reserve to buyer
				err := k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
					types.BatchesIntermediaryAccount, bo.Address, bo.MaxPrices)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	// Save batch and return number of cancelled orders
	k.SetBatch(ctx, bondDid, batch)
	return cancelledOrders
}

func (k Keeper) CancelUnfulfillableOrders(ctx sdk.Context, bondDid ixo.Did) (cancelledOrders int) {
	batch := k.MustGetBatch(ctx, bondDid)
	cancelledOrders = 0

	cancelledOrders += k.CancelUnfulfillableBuys(ctx, bondDid)
	//cancelledOrders += k.CancelUnfulfillableSells(ctx, bondDid) // Sells always fulfillable
	//cancelledOrders += k.CancelUnfulfillableSwaps(ctx, bondDid) // Swaps only cancelled while they are being performed

	// Update buy and sell prices if any cancellation took place
	if cancelledOrders > 0 {
		batch = k.MustGetBatch(ctx, bondDid) // get batch again
		buyPrices, sellPrices, err := k.GetBatchBuySellPrices(ctx, bondDid, batch)
		if err != nil {
			panic(err)
		}
		batch.BuyPrices = buyPrices
		batch.SellPrices = sellPrices
	}

	// Save batch and return number of cancelled orders
	k.SetBatch(ctx, bondDid, batch)
	return cancelledOrders
}
