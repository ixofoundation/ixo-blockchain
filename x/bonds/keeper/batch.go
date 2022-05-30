package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

func (k Keeper) MustGetBatch(ctx sdk.Context, bondDid didexported.Did) types.Batch {
	store := ctx.KVStore(k.storeKey)
	if !k.BatchExists(ctx, bondDid) {
		panic(fmt.Sprintf("batch not found for %s\n", bondDid))
	}

	bz := store.Get(types.GetBatchKey(bondDid))
	var batch types.Batch
	k.cdc.MustUnmarshal(bz, &batch)

	return batch
}

func (k Keeper) MustGetLastBatch(ctx sdk.Context, bondDid didexported.Did) types.Batch {
	store := ctx.KVStore(k.storeKey)
	if !k.LastBatchExists(ctx, bondDid) {
		panic(fmt.Sprintf("last batch not found for %s\n", bondDid))
	}

	bz := store.Get(types.GetLastBatchKey(bondDid))
	var batch types.Batch
	k.cdc.MustUnmarshal(bz, &batch)

	return batch
}

func (k Keeper) BatchExists(ctx sdk.Context, bondDid didexported.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetBatchKey(bondDid))
}

func (k Keeper) LastBatchExists(ctx sdk.Context, bondDid didexported.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetLastBatchKey(bondDid))
}

func (k Keeper) SetBatch(ctx sdk.Context, bondDid didexported.Did, batch types.Batch) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBatchKey(bondDid), k.cdc.MustMarshal(&batch))
}

func (k Keeper) SetLastBatch(ctx sdk.Context, bondDid didexported.Did, batch types.Batch) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetLastBatchKey(bondDid), k.cdc.MustMarshal(&batch))
}

func (k Keeper) AddBuyOrder(ctx sdk.Context, bondDid didexported.Did, bo types.BuyOrder, buyPrices, sellPrices sdk.DecCoins) {
	batch := k.MustGetBatch(ctx, bondDid)
	batch.TotalBuyAmount = batch.TotalBuyAmount.Add(bo.BaseOrder.Amount)
	batch.BuyPrices = buyPrices
	batch.SellPrices = sellPrices
	batch.Buys = append(batch.Buys, bo)
	k.SetBatch(ctx, bondDid, batch)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("added buy order for %s from %s", bo.BaseOrder.Amount.String(), bo.BaseOrder.AccountDid))
}

func (k Keeper) AddSellOrder(ctx sdk.Context, bondDid didexported.Did, so types.SellOrder, buyPrices, sellPrices sdk.DecCoins) {
	batch := k.MustGetBatch(ctx, bondDid)
	batch.TotalSellAmount = batch.TotalSellAmount.Add(so.BaseOrder.Amount)
	batch.BuyPrices = buyPrices
	batch.SellPrices = sellPrices
	batch.Sells = append(batch.Sells, so)
	k.SetBatch(ctx, bondDid, batch)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("added sell order for %s from %s", so.BaseOrder.Amount.String(), so.BaseOrder.AccountDid))
}

func (k Keeper) AddSwapOrder(ctx sdk.Context, bondDid didexported.Did, so types.SwapOrder) {
	batch := k.MustGetBatch(ctx, bondDid)
	batch.Swaps = append(batch.Swaps, so)
	k.SetBatch(ctx, bondDid, batch)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("added swap order for %s to %s from %s", so.BaseOrder.Amount.String(), so.ToToken, so.BaseOrder.AccountDid))
}

func (k Keeper) GetBatchBuySellPrices(ctx sdk.Context, bondDid string, batch types.Batch) (buyPricesPT, sellPricesPT sdk.DecCoins, err error) {
	bond := k.MustGetBond(ctx, bondDid)

	buyAmountDec := batch.TotalBuyAmount.Amount.ToDec()
	sellAmountDec := batch.TotalSellAmount.Amount.ToDec()

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
		curvedValues, err = bond.GetReturnsForBurn(extraSells.Amount, reserveBalances) // sell returns
		if err != nil {
			return nil, nil, err
		}
	}

	// Get (actual) matched values
	matchedValues := types.MultiplyDecCoinsByDec(currentPricesPT, matchedAmount)

	// If buys > sells, totalValues is the total buy prices
	// If sells > buys, totalValues is the total sell returns
	totalValues := matchedValues.Add(curvedValues...)

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

func (k Keeper) GetUpdatedBatchPricesAfterBuy(ctx sdk.Context, bondDid didexported.Did, bo types.BuyOrder) (buyPrices, sellPrices sdk.DecCoins, err error) {
	bond := k.MustGetBond(ctx, bondDid)
	batch := k.MustGetBatch(ctx, bondDid)

	// Max supply cannot be less than supply (max supply >= supply)
	adjustedSupply := k.GetSupplyAdjustedForBuy(ctx, bondDid)
	adjustedSupplyWithBuy := adjustedSupply.Add(bo.BaseOrder.Amount)
	if bond.MaxSupply.IsLT(adjustedSupplyWithBuy) {
		return nil, nil, types.ErrCannotMintMoreThanMaxSupply
	}

	// If augmented in hatch phase and adjusted supply exceeds S0, disallow buy
	// since it is not allowed for a batch to cross over to the open phase.
	//
	// S0 is rounded to ceil for the case that it has a decimal, otherwise it
	// cannot be reached without being exceeded, when using integer buy amounts
	// (e.g. if supply is 100 and S0=100.5, we cannot reach S0 by performing
	// the minimum buy of 1 token [101>100.5], so S0 is rounded to ceil; S0=101)
	if bond.FunctionType == types.AugmentedFunction &&
		bond.State == types.HatchState.String() {
		args := bond.FunctionParameters.AsMap()
		if adjustedSupplyWithBuy.Amount.ToDec().GT(args["S0"].Ceil()) {
			return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins,
				"Buy exceeds initial supply S0. Consider buying less tokens.")
		}
	}

	// Simulate buy by bumping up total buy amount
	batch.TotalBuyAmount = batch.TotalBuyAmount.Add(bo.BaseOrder.Amount)
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

func (k Keeper) GetUpdatedBatchPricesAfterSell(ctx sdk.Context, bondDid didexported.Did, so types.SellOrder) (buyPrices, sellPrices sdk.DecCoins, err error) {
	batch := k.MustGetBatch(ctx, bondDid)

	// Cannot burn more tokens than what exists
	adjustedSupply := k.GetSupplyAdjustedForSell(ctx, bondDid)
	if adjustedSupply.IsLT(so.BaseOrder.Amount) {
		return nil, nil, types.ErrCannotBurnMoreThanSupply
	}

	// Simulate sell by bumping up total sell amount
	batch.TotalSellAmount = batch.TotalSellAmount.Add(so.BaseOrder.Amount)
	buyPrices, sellPrices, err = k.GetBatchBuySellPrices(ctx, bondDid, batch)
	if err != nil {
		return nil, nil, err
	}

	return buyPrices, sellPrices, nil
}

func (k Keeper) PerformBuyAtPrice(ctx sdk.Context, bondDid didexported.Did, bo types.BuyOrder, prices sdk.DecCoins) (err error) {
	bond := k.MustGetBond(ctx, bondDid)
	var extraEventAttributes []sdk.Attribute

	// Get buyer address
	buyerDidDoc, err := k.didKeeper.GetDidDoc(ctx, bo.BaseOrder.AccountDid)
	if err != nil {
		return err
	}
	buyerAddr := buyerDidDoc.Address()

	feeAddr, err := sdk.AccAddressFromBech32(bond.FeeAddress)
	if err != nil {
		return err
	}

	// Mint bond tokens
	err = k.BankKeeper.MintCoins(ctx, types.BondsMintBurnAccount,
		sdk.Coins{bo.BaseOrder.Amount})
	if err != nil {
		return err
	}

	// Send bond tokens bought to buyer
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
		types.BondsMintBurnAccount, buyerAddr, sdk.Coins{bo.BaseOrder.Amount})
	if err != nil {
		return err
	}

	reservePrices := types.MultiplyDecCoinsByInt(prices, bo.BaseOrder.Amount.Amount)
	reservePricesRounded := types.RoundReservePrices(reservePrices)
	txFees := bond.GetTxFees(reservePrices)
	totalPrices := reservePricesRounded.Add(txFees...)

	if totalPrices.IsAnyGT(bo.MaxPrices) {
		return sdkerrors.Wrapf(types.ErrMaxPriceExceeded,
			"actual prices %s exceed max prices %s",
			totalPrices.String(), bo.MaxPrices.String())
	}

	// Add new reserve to reserve (reservePricesRounded should never be zero)
	// TODO: investigate possibility of zero reservePricesRounded
	if bond.FunctionType == types.AugmentedFunction &&
		bond.State == types.HatchState.String() {
		args := bond.FunctionParameters.AsMap()
		theta := args["theta"]

		// Get current reserve
		var currentReserve sdk.Int
		if bond.CurrentReserve.Empty() {
			currentReserve = sdk.ZeroInt()
		} else {
			// Reserve balances should all be equal given that we are always
			// applying the same additions/subtractions to all reserve balances.
			// Thus we can pick the first reserve balance as the global balance.
			currentReserve = k.GetReserveBalances(ctx, bondDid)[0].Amount
		}

		// Calculate expected new reserve (as fraction 1-theta of new total raise)
		newSupply := bond.CurrentSupply.Add(bo.BaseOrder.Amount).Amount
		newTotalRaise := args["p0"].Mul(newSupply.ToDec())
		newReserve := newTotalRaise.Mul(
			sdk.OneDec().Sub(theta)).Ceil().TruncateInt()

		// Calculate amount that should go into initial reserve
		toInitialReserve := newReserve.Sub(currentReserve)
		if reservePricesRounded[0].Amount.LT(toInitialReserve) {
			// Reserve supplied by buyer is insufficient
			return types.ErrInsufficientReserveToBuy
		}
		coinsToInitialReserve, _ := bond.GetNewReserveDecCoins(
			toInitialReserve.ToDec()).TruncateDecimal()

		// Calculate amount that should go into funding pool
		coinsToFundingPool := reservePricesRounded.Sub(coinsToInitialReserve)

		// Send reserve tokens to initial reserve
		err = k.DepositReserveFromModule(ctx, bond.BondDid,
			types.BatchesIntermediaryAccount, coinsToInitialReserve)
		if err != nil {
			return err
		}

		// Send reserve tokens to funding pool
		err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, feeAddr, coinsToFundingPool)
		if err != nil {
			return err
		}

		extraEventAttributes = append(extraEventAttributes,
			sdk.NewAttribute(types.AttributeKeyChargedPricesReserve, toInitialReserve.String()),
			sdk.NewAttribute(types.AttributeKeyChargedPricesFunding, coinsToFundingPool.String()),
		)
	} else {
		err = k.DepositReserveFromModule(
			ctx, bond.BondDid, types.BatchesIntermediaryAccount, reservePricesRounded)
		if err != nil {
			return err
		}
	}

	// Add charged fee to fee address
	if !txFees.IsZero() {
		err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, feeAddr, txFees)
		if err != nil {
			return err
		}
	}

	// Add remainder to buyer address
	returnToBuyer := bo.MaxPrices.Sub(totalPrices)
	if !returnToBuyer.IsZero() {
		err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, buyerAddr, returnToBuyer)
		if err != nil {
			return err
		}
	}

	// Update supply (max supply exceeded check done during MsgBuy)
	k.SetCurrentSupply(ctx, bondDid, bond.CurrentSupply.Add(bo.BaseOrder.Amount))

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("performed buy order for %s from %s", bo.BaseOrder.Amount.String(), bo.BaseOrder.AccountDid))

	// Get new bond token balance
	bondTokenBalance := k.BankKeeper.GetBalance(ctx, buyerAddr, bond.Token).Amount

	event := sdk.NewEvent(
		types.EventTypeOrderFulfill,
		sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueBuyOrder),
		sdk.NewAttribute(types.AttributeKeyAddress, bo.BaseOrder.AccountDid),
		sdk.NewAttribute(types.AttributeKeyTokensMinted, bo.BaseOrder.Amount.Amount.String()),
		sdk.NewAttribute(types.AttributeKeyChargedPrices, reservePricesRounded.String()),
		sdk.NewAttribute(types.AttributeKeyChargedFees, txFees.String()),
		sdk.NewAttribute(types.AttributeKeyReturnedToAddress, returnToBuyer.String()),
		sdk.NewAttribute(types.AttributeKeyNewBondTokenBalance, bondTokenBalance.String()),
	)
	if len(extraEventAttributes) > 0 {
		event = event.AppendAttributes(extraEventAttributes...)
	}
	ctx.EventManager().EmitEvent(event)

	return nil
}

func (k Keeper) PerformSellAtPrice(ctx sdk.Context, bondDid didexported.Did, so types.SellOrder, prices sdk.DecCoins) (err error) {
	bond := k.MustGetBond(ctx, bondDid)

	// Get seller address
	sellerDidDoc, err := k.didKeeper.GetDidDoc(ctx, so.BaseOrder.AccountDid)
	if err != nil {
		return err
	}
	sellerAddr := sellerDidDoc.Address()

	reserveReturns := types.MultiplyDecCoinsByInt(prices, so.BaseOrder.Amount.Amount)
	reserveReturnsRounded := types.RoundReserveReturns(reserveReturns)
	txFees := bond.GetTxFees(reserveReturns)
	exitFees := bond.GetExitFees(reserveReturns)

	totalFees := types.AdjustFees(txFees.Add(exitFees...), reserveReturnsRounded) // calculate actual total fees
	totalReturns := reserveReturnsRounded.Sub(totalFees)                          // calculate actual reserveReturns

	// Send total returns to seller (totalReturns should never be zero)
	// TODO: investigate possibility of zero totalReturns
	err = k.WithdrawFromReserve(ctx, bond.BondDid, sellerAddr, totalReturns)
	if err != nil {
		return err
	}

	feeAddr, err := sdk.AccAddressFromBech32(bond.FeeAddress)
	if err != nil {
		return err
	}

	// Send total fee to fee address
	if !totalFees.IsZero() {
		err = k.WithdrawFromReserve(ctx, bond.BondDid, feeAddr, totalFees)
		if err != nil {
			return err
		}
	}

	// Update supply (burn more than supply check done during MsgSell)
	k.SetCurrentSupply(ctx, bondDid, bond.CurrentSupply.Sub(so.BaseOrder.Amount))

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("performed sell order for %s from %s", so.BaseOrder.Amount.String(), so.BaseOrder.AccountDid))

	// Get new bond token balance
	bondTokenBalance := k.BankKeeper.GetBalance(ctx, sellerAddr, bond.Token).Amount

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeOrderFulfill,
		sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueSellOrder),
		sdk.NewAttribute(types.AttributeKeyAddress, so.BaseOrder.AccountDid),
		sdk.NewAttribute(types.AttributeKeyTokensBurned, so.BaseOrder.Amount.Amount.String()),
		sdk.NewAttribute(types.AttributeKeyChargedFees, txFees.String()),
		sdk.NewAttribute(types.AttributeKeyReturnedToAddress, totalReturns.String()),
		sdk.NewAttribute(types.AttributeKeyNewBondTokenBalance, bondTokenBalance.String()),
	))

	return nil
}

func (k Keeper) PerformSwap(ctx sdk.Context, bondDid didexported.Did, so types.SwapOrder) (err error, ok bool) {
	bond := k.MustGetBond(ctx, bondDid)

	// WARNING: do not return ok=true if money has already been transferred when error occurs

	// Get swapper address
	swapperDidDoc, err := k.didKeeper.GetDidDoc(ctx, so.BaseOrder.AccountDid)
	if err != nil {
		return err, true
	}
	swapperAddr := swapperDidDoc.Address()

	// Get return for swap
	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(so.BaseOrder.Amount, so.ToToken, reserveBalances)
	if err != nil {
		return err, true
	}
	adjustedInput := so.BaseOrder.Amount.Sub(txFee) // same as during GetReturnsForSwap

	// Check if new rates violate sanity rate
	newReserveBalances := reserveBalances.Add(sdk.Coins{adjustedInput}...).Sub(reserveReturns)
	if bond.ReservesViolateSanityRate(newReserveBalances) {
		return types.ErrValuesViolateSanityRate, true
	}

	// Give resultant tokens to swapper (reserveReturns should never be zero)
	err = k.WithdrawFromReserve(ctx, bond.BondDid, swapperAddr, reserveReturns)
	if err != nil {
		return err, false
	}

	// Add fee-reduced coins to be swapped to reserve (adjustedInput should never be zero)
	err = k.DepositReserveFromModule(
		ctx, bond.BondDid, types.BatchesIntermediaryAccount, sdk.Coins{adjustedInput})
	if err != nil {
		return err, false
	}

	feeAddr, err := sdk.AccAddressFromBech32(bond.FeeAddress)
	if err != nil {
		return err, false
	}

	// Add fee (taken from swapper) to fee address
	if !txFee.IsZero() {
		err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
			types.BatchesIntermediaryAccount, feeAddr, sdk.Coins{txFee})
		if err != nil {
			return err, false
		}
	}

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("performed swap order for %s to %s from %s",
		so.BaseOrder.Amount.String(), reserveReturns, so.BaseOrder.AccountDid))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeOrderFulfill,
		sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueSwapOrder),
		sdk.NewAttribute(types.AttributeKeyAddress, so.BaseOrder.AccountDid),
		sdk.NewAttribute(types.AttributeKeyTokensSwapped, adjustedInput.String()),
		sdk.NewAttribute(types.AttributeKeyChargedFees, txFee.String()),
		sdk.NewAttribute(types.AttributeKeyReturnedToAddress, reserveReturns.String()),
	))

	return nil, true
}

func (k Keeper) PerformBuyOrders(ctx sdk.Context, bondDid didexported.Did) {
	batch := k.MustGetBatch(ctx, bondDid)

	// Perform buys or return to buyer
	for _, bo := range batch.Buys {
		if !bo.BaseOrder.IsCancelled() {
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

func (k Keeper) PerformSellOrders(ctx sdk.Context, bondDid didexported.Did) {
	batch := k.MustGetBatch(ctx, bondDid)

	// Perform sells or return to seller
	for _, so := range batch.Sells {
		if !so.BaseOrder.IsCancelled() {
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

func (k Keeper) PerformSwapOrders(ctx sdk.Context, bondDid didexported.Did) {
	logger := ctx.Logger()
	batch := k.MustGetBatch(ctx, bondDid)

	// Perform swaps
	// TODO: implement swaps front-running prevention
	for i, so := range batch.Swaps {
		if !so.BaseOrder.IsCancelled() {
			err, ok := k.PerformSwap(ctx, bondDid, so)
			if err != nil {
				if ok {
					batch.Swaps[i].BaseOrder.Cancelled = true
					batch.Swaps[i].BaseOrder.CancelReason = err.Error()

					logger.Info(fmt.Sprintf("cancelled swap order for %s to %s from %s", so.BaseOrder.Amount.String(), so.ToToken, so.BaseOrder.AccountDid))
					logger.Debug(fmt.Sprintf("cancellation reason: %s", err.Error()))

					// Return from amount to swapper
					swapperAddr := k.didKeeper.MustGetDidDoc(ctx, so.BaseOrder.AccountDid).Address()
					err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
						types.BatchesIntermediaryAccount, swapperAddr, sdk.Coins{so.BaseOrder.Amount})
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

func (k Keeper) PerformOrders(ctx sdk.Context, bondDid didexported.Did) {
	k.PerformBuyOrders(ctx, bondDid)
	k.PerformSellOrders(ctx, bondDid)
	k.PerformSwapOrders(ctx, bondDid)
}

func (k Keeper) CheckIfBuyOrderFulfillableAtPrice(ctx sdk.Context, bondDid didexported.Did, bo types.BuyOrder, prices sdk.DecCoins) error {
	bond := k.MustGetBond(ctx, bondDid)

	reservePrices := types.MultiplyDecCoinsByInt(prices, bo.BaseOrder.Amount.Amount)
	reserveRounded := types.RoundReservePrices(reservePrices)
	txFees := bond.GetTxFees(reservePrices)
	totalPrices := reserveRounded.Add(txFees...)

	// Check that max prices not exceeded
	if totalPrices.IsAnyGT(bo.MaxPrices) {
		return sdkerrors.Wrapf(types.ErrMaxPriceExceeded,
			"actual prices %s exceed max prices %s",
			totalPrices.String(), bo.MaxPrices.String())
	}

	return nil
}

func (k Keeper) CancelUnfulfillableBuys(ctx sdk.Context, bondDid didexported.Did) (cancelledOrders int) {
	logger := k.Logger(ctx)
	batch := k.MustGetBatch(ctx, bondDid)

	// Cancel unfulfillable buys
	for i, bo := range batch.Buys {
		if !bo.BaseOrder.IsCancelled() {
			err := k.CheckIfBuyOrderFulfillableAtPrice(ctx, bondDid, bo, batch.BuyPrices)
			if err != nil {
				// Cancel (important to use batch.Buys[i] and not bo!)
				batch.Buys[i].BaseOrder.Cancelled = true
				batch.Buys[i].BaseOrder.CancelReason = err.Error()
				batch.TotalBuyAmount = batch.TotalBuyAmount.Sub(bo.BaseOrder.Amount)
				cancelledOrders += 1

				logger.Info(fmt.Sprintf("cancelled buy order for %s from %s", bo.BaseOrder.Amount.String(), bo.BaseOrder.AccountDid))
				logger.Debug(fmt.Sprintf("cancellation reason: %s", err.Error()))

				ctx.EventManager().EmitEvent(sdk.NewEvent(
					types.EventTypeOrderCancel,
					sdk.NewAttribute(types.AttributeKeyBondDid, bondDid),
					sdk.NewAttribute(types.AttributeKeyOrderType, types.AttributeValueBuyOrder),
					sdk.NewAttribute(types.AttributeKeyAddress, bo.BaseOrder.AccountDid),
					sdk.NewAttribute(types.AttributeKeyCancelReason, bo.BaseOrder.CancelReason),
				))

				// Return reserve to buyer
				buyerAddr := k.didKeeper.MustGetDidDoc(ctx, bo.BaseOrder.AccountDid).Address()
				err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx,
					types.BatchesIntermediaryAccount, buyerAddr, bo.MaxPrices)
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

func (k Keeper) CancelUnfulfillableOrders(ctx sdk.Context, bondDid didexported.Did) (cancelledOrders int) {
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

func (k Keeper) UpdateAlpha(ctx sdk.Context, bondDid didexported.Did) {
	//fmt.Println("Updating alpha-------------------------")
	bond := k.MustGetBond(ctx, bondDid)
	//fmt.Println("bond did: ", bondDid)
	batch := k.MustGetBatch(ctx, bondDid)
	newPublicAlpha := batch.NextPublicAlpha

	// Get supply, reserve, outcome payment
	// S := bond.CurrentSupply.Amount.ToDec()
	//fmt.Println("S: ", S)
	R := bond.CurrentReserve[0].Amount.ToDec()
	//fmt.Println("R: ", R)
	C := bond.OutcomePayment
	//fmt.Println("C: ", C)

	// Get current parameters
	paramsMap := bond.FunctionParameters.AsMap()

	// Calculate scaled delta public alpha, to calculate new system alpha
	prevPublicAlpha := paramsMap["publicAlpha"]
	//fmt.Println("prevPublicAlpha: ", prevPublicAlpha)
	deltaPublicAlpha := newPublicAlpha.Sub(prevPublicAlpha)
	//fmt.Println("deltaPublicAlpha: ", deltaPublicAlpha)
	temp, err := types.ApproxPower(
		prevPublicAlpha.Mul(sdk.OneDec().Sub(types.StartingPublicAlpha)),
		sdk.MustNewDecFromStr("2"))
	if err != nil {
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeEditAlphaFailed,
			sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
			sdk.NewAttribute(types.AttributeKeyToken, bond.Token),
			sdk.NewAttribute(types.AttributeKeyCancelReason, err.Error()),
		))
		return
	}
	//fmt.Println("temp: ", temp)
	scaledDeltaPublicAlpha := deltaPublicAlpha.Mul(temp)
	//fmt.Println("scaledDeltaPublicAlpha: ", scaledDeltaPublicAlpha)

	// Calculate new system alpha
	prevSystemAlpha := paramsMap["systemAlpha"]
	//fmt.Println("prevSystemAlpha: ", prevSystemAlpha)
	var newSystemAlpha sdk.Dec
	if deltaPublicAlpha.IsPositive() {
		//fmt.Println("deltaPublicAlpha is positive")
		// 1 - (1 - scaled_delta_public_alpha) * (1 - previous_alpha)
		temp1 := sdk.OneDec().Sub(scaledDeltaPublicAlpha)
		//fmt.Println("temp1: ", temp1)
		temp2 := sdk.OneDec().Sub(prevSystemAlpha)
		//fmt.Println("temp2: ", temp2)
		newSystemAlpha = sdk.OneDec().Sub(temp1.Mul(temp2))
	} else {
		//fmt.Println("deltaPublicAlpha is negative")
		// (1 - scaled_delta_public_alpha) * (previous_alpha)
		temp1 := sdk.OneDec().Sub(scaledDeltaPublicAlpha)
		//fmt.Println("temp1: ", temp1)
		temp2 := prevSystemAlpha
		//fmt.Println("temp2: ", temp2)
		newSystemAlpha = temp1.Mul(temp2)
	}
	//fmt.Println("newSystemAlpha: ", newSystemAlpha)

	// Check 1 (newSystemAlpha != prevSystemAlpha)
	if newSystemAlpha.Equal(prevSystemAlpha) {
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeEditAlphaFailed,
			sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
			sdk.NewAttribute(types.AttributeKeyToken, bond.Token),
			sdk.NewAttribute(types.AttributeKeyCancelReason,
				"resultant system alpha based on public alpha is unchanged"),
		))
		return
	}
	// Check 2 (I > C * newSystemAlpha)
	if paramsMap["I0"].LTE(newSystemAlpha.MulInt(C)) {
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeEditAlphaFailed,
			sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
			sdk.NewAttribute(types.AttributeKeyToken, bond.Token),
			sdk.NewAttribute(types.AttributeKeyCancelReason,
				"cannot change alpha to that value due to violated restriction [1]"),
		))
		return
	}
	// Check 3 (R / C > newSystemAlpha - prevSystemAlpha)
	if R.QuoInt(C).LTE(newSystemAlpha.Sub(prevSystemAlpha)) {
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeEditAlphaFailed,
			sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
			sdk.NewAttribute(types.AttributeKeyToken, bond.Token),
			sdk.NewAttribute(types.AttributeKeyCancelReason,
				"cannot change alpha to that value due to violated restriction [2]"),
		))
		return
	}

	// Recalculate kappa and V0 using new alpha
	// I0 := paramsMap["I0"]
	// //fmt.Println("I0: ", I0)
	// newKappa := types.Kappa(I0, C, newSystemAlpha)
	// //fmt.Println("newKappa: ", newKappa)
	// newV0, err := types.Invariant(R, S, newKappa)
	// //fmt.Println("newV0: ", newV0)
	// if err != nil {
	// 	ctx.EventManager().EmitEvent(sdk.NewEvent(
	// 		types.EventTypeEditAlphaFailed,
	// 		sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
	// 		sdk.NewAttribute(types.AttributeKeyToken, bond.Token),
	// 		sdk.NewAttribute(types.AttributeKeyCancelReason, err.Error()),
	// 	))
	// 	return
	// }

	// Get batch to reset alpha
	batch = k.MustGetBatch(ctx, bond.BondDid)
	batch.NextPublicAlpha = sdk.OneDec().Neg()
	k.SetBatch(ctx, bond.BondDid, batch)

	// Set new function parameters
	// bond.FunctionParameters.ReplaceParam("kappa", newKappa)
	// bond.FunctionParameters.ReplaceParam("V0", newV0)
	bond.FunctionParameters.ReplaceParam("publicAlpha", newPublicAlpha)
	bond.FunctionParameters.ReplaceParam("systemAlpha", newSystemAlpha)
	k.SetBond(ctx, bond.BondDid, bond)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEditAlphaSuccess,
		sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
		sdk.NewAttribute(types.AttributeKeyToken, bond.Token),
		sdk.NewAttribute(types.AttributeKeyPublicAlpha, newPublicAlpha.String()),
		sdk.NewAttribute(types.AttributeKeySystemAlpha, newSystemAlpha.String()),
	))
}
