package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

// -------------------------------------------------------- Fees Get/Set

func (k Keeper) GetFeeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FeeKeyPrefix)
}

func (k Keeper) MustGetFeeByKey(ctx sdk.Context, key []byte) types.Fee {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("fee not found")
	}

	bz := store.Get(key)
	var fee types.Fee
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fee)

	return fee
}

func (k Keeper) FeeExists(ctx sdk.Context, feeId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeKey(feeId))
}

func (k Keeper) GetFee(ctx sdk.Context, feeId string) (types.Fee, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeKey(feeId)

	bz := store.Get(key)
	if bz == nil {
		return types.Fee{}, sdk.ErrInternal("invalid fee")
	}

	var fee types.Fee
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fee)

	return fee, nil
}

func (k Keeper) SetFee(ctx sdk.Context, fee types.Fee) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeKey(fee.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(fee))
}

func (k Keeper) DiscountIdExists(ctx sdk.Context, feeId string, discountId sdk.Uint) (bool, sdk.Error) {
	// Get fee
	fee, err := k.GetFee(ctx, feeId)
	if err != nil {
		return false, err
	}

	// Search for discount ID
	for _, d := range fee.Content.Discounts {
		if d.Id.Equal(discountId) {
			return true, nil
		}
	}
	return false, nil
}

// -------------------------------------------------------- FeeContracts Get/Set

func (k Keeper) GetFeeContractIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FeeContractKeyPrefix)
}

func (k Keeper) MustGetFeeContractByKey(ctx sdk.Context, key []byte) types.FeeContract {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("fee contract not found")
	}

	bz := store.Get(key)
	var feeContract types.FeeContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContract)

	return feeContract
}

func (k Keeper) FeeContractExists(ctx sdk.Context, feeContractId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeContractKey(feeContractId))
}

func (k Keeper) GetFeeContract(ctx sdk.Context, feeContractId string) (types.FeeContract, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractKey(feeContractId)

	bz := store.Get(key)
	if bz == nil {
		return types.FeeContract{}, sdk.ErrInternal("invalid fee contract")
	}

	var feeContract types.FeeContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContract)

	return feeContract, nil
}

func (k Keeper) SetFeeContract(ctx sdk.Context, feeContract types.FeeContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractKey(feeContract.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(feeContract))
}

func (k Keeper) SetFeeContractAuthorised(ctx sdk.Context, feeContractId string,
	authorised bool) sdk.Error {
	feeContract, err := k.GetFeeContract(ctx, feeContractId)
	if err != nil {
		return err
	}

	// If de-authorising, check if can be de-authorised
	if !authorised && !feeContract.Content.CanDeauthorise {
		return types.ErrFeeContractCannotBeDeauthorised(types.DefaultCodespace)
	}

	// Set authorised state
	feeContract.Content.Authorised = authorised
	k.SetFeeContract(ctx, feeContract)

	return nil
}

func (k Keeper) GrantFeeDiscount(ctx sdk.Context, feeContractId string, discountId sdk.Uint) sdk.Error {
	// Get fee contract
	feeContract, err := k.GetFeeContract(ctx, feeContractId)
	if err != nil {
		return err
	}

	// Overwrite previous discount ID
	feeContract.Content.DiscountId = discountId
	k.SetFeeContract(ctx, feeContract)
	return nil
}

func (k Keeper) RevokeFeeDiscount(ctx sdk.Context, feeContractId string) sdk.Error {
	// Get fee contract
	feeContract, err := k.GetFeeContract(ctx, feeContractId)
	if err != nil {
		return err
	}

	// Set discount ID to zero
	feeContract.Content.DiscountId = sdk.ZeroUint()
	k.SetFeeContract(ctx, feeContract)
	return nil
}

// -------------------------------------------------------- FeeContracts Charge

func applyDiscount(ctx sdk.Context, k Keeper, fee types.Fee, feeContract types.FeeContract,
	payer sdk.AccAddress, payAmount sdk.Coins) (sdk.Coins, sdk.Error) {

	// No discounts held
	if feeContract.Content.DiscountId.IsZero() {
		return payAmount, nil
	}

	// Get discount percentage to calculate discount amount. Any rounding
	// when multiplying means the payer receives a slightly smaller discount.
	discountPercent, err := fee.Content.GetDiscountPercent(feeContract.Content.DiscountId)
	if err != nil {
		return nil, err
	}
	discountPercentDec := discountPercent.Quo(sdk.NewDec(100)) // 50 -> 0.5
	discountAmt, _ := sdk.NewDecCoins(payAmount).MulDec(discountPercentDec).TruncateDecimal()

	// Confirm that discount is not greater than the payAmount
	if discountAmt.IsAnyGT(payAmount) {
		return nil, types.ErrDiscountPercentageGreaterThan100(types.DefaultCodespace)
	}

	// Return payAmount with discount deducted
	return payAmount.Sub(discountAmt), nil
}

func adjustForMinimums(fee types.Fee, feeContract types.FeeContract, cumulative sdk.Coins) {
	// If first charge, increase to the minimum charge if the cumulative charge
	// is less than the minimum (applied on each denomination independently)
	if feeContract.IsFirstCharge() {
		for i, coin := range cumulative {
			minAmt := fee.Content.ChargeMinimum.AmountOf(coin.Denom)
			if !minAmt.IsZero() && minAmt.GT(coin.Amount) {
				cumulative[i] = sdk.NewCoin(coin.Denom, minAmt)
			}
		}
	}
}

func adjustForMaximums(fee types.Fee, cumulative sdk.Coins) {
	// Reduce to the maximum charge if the cumulative charge is more than the
	// maximum (applied on each denomination independently)
	for i, coin := range cumulative {
		maxAmt := fee.Content.ChargeMaximum.AmountOf(coin.Denom)
		if !maxAmt.IsZero() && maxAmt.LT(coin.Amount) {
			cumulative[i] = sdk.NewCoin(coin.Denom, maxAmt)
		}
	}
}

func (k Keeper) ChargeFee(ctx sdk.Context, bankKeeper bank.Keeper,
	feeContractId string) (charged bool, err sdk.Error) {

	feeContract, err := k.GetFeeContract(ctx, feeContractId)
	if err != nil {
		return false, err
	}
	fcData := &feeContract.Content

	fee, err := k.GetFee(ctx, fcData.FeeId)
	if err != nil {
		return false, err
	}
	feeData := &fee.Content

	// Check if can charge (this is false if e.g. max charge has been reached)
	if !feeContract.CanCharge(fee) {
		return false, nil
	}

	// Assume payer will pay ChargeAmount, apply discount (if any),
	// and calculate initial cumulative (before adjustments)
	payAmount := feeData.ChargeAmount
	payAmount, err = applyDiscount(ctx, k, fee, feeContract, fcData.Payer, payAmount)
	if err != nil {
		return false, err
	}
	cumulative := fcData.CumulativeCharge.Add(payAmount)

	// In-place cumulative adjustments (i.e. considering minimums and maximums)
	adjustForMinimums(fee, feeContract, cumulative)
	adjustForMaximums(fee, cumulative)

	// Find actual charge from adjusted cumulative:
	//    adjustedCumul = previousCumul + actualCharge
	// => actualCharge = adjustedCumul - previousCumul
	charge := cumulative.Sub(fcData.CumulativeCharge)

	// Stop if payer doesn't have enough coins. However, this is not considered
	// an error but the caller should be looking at the 'charged' bool result
	if !bankKeeper.HasCoins(ctx, fcData.Payer, charge) {
		return false, nil
	}

	// Total input is charge plus current remainder in FeeRemainderPool
	inputFromFeeRemainderPool := fcData.CurrentRemainder
	totalInputAmount := charge.Add(inputFromFeeRemainderPool)

	// Calculate list of outputs and calculate the total output to payees based
	// on the calculated wallet distributions
	var outputToPayees sdk.Coins
	var outputs []bank.Output
	distributions := fee.Content.WalletDistribution.GetDistributionsFor(totalInputAmount)
	for i, share := range distributions {
		// Get integer output
		outputAmt, _ := share.TruncateDecimal()

		// If amount not zero, update total and add as output
		if !outputAmt.IsZero() {
			outputToPayees = outputToPayees.Add(outputAmt)
			address := fee.Content.WalletDistribution[i].Address
			outputs = append(outputs, bank.NewOutput(address, outputAmt))
		}
	}

	// Remainder (not output to payees) goes to FeeRemainderPool if not zero
	outputToFeeRemainderPool := totalInputAmount.Sub(outputToPayees)
	if !outputToFeeRemainderPool.IsZero() {
		feeRemainderPoolAddr := supply.NewModuleAddress(types.FeeRemainderPool)
		outputs = append(outputs, bank.NewOutput(feeRemainderPoolAddr, outputToFeeRemainderPool))
	}

	// Construct list of inputs (charge and from FeeRemainderPool if non zero)
	inputs := []bank.Input{bank.NewInput(fcData.Payer, charge)}
	if !inputFromFeeRemainderPool.IsZero() {
		feeRemainderPoolAddr := supply.NewModuleAddress(types.FeeRemainderPool)
		inputs = append(inputs, bank.NewInput(feeRemainderPoolAddr, inputFromFeeRemainderPool))
	}

	// Distribute the fee charge according to the outputs
	err = bankKeeper.InputOutputCoins(ctx, inputs, outputs)
	if err != nil {
		return false, err
	}

	// Update and save fee contract
	fcData.CumulativeCharge = fcData.CumulativeCharge.Add(charge)
	fcData.CurrentRemainder = fcData.CurrentRemainder.Add(outputToFeeRemainderPool).Sub(inputFromFeeRemainderPool)
	k.SetFeeContract(ctx, feeContract)

	return true, nil
}
