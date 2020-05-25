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

func (k Keeper) FeeExists(ctx sdk.Context, feeId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeKey(feeId))
}

func (k Keeper) GetFee(ctx sdk.Context, feeId uint64) (types.Fee, sdk.Error) {
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

func (k Keeper) SubmitFee(ctx sdk.Context, content types.FeeContent) (types.Fee, sdk.Error) {
	feeId, err := k.GetFeeID(ctx)
	if err != nil {
		return types.Fee{}, err
	}

	fee := types.NewFee(feeId, content)

	k.SetFee(ctx, fee)
	k.SetFeeID(ctx, feeId+1)

	return fee, nil
}

func (k Keeper) SetFee(ctx sdk.Context, fee types.Fee) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeKey(fee.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(fee))
}

// GetFeeID gets the highest fee ID
func (k Keeper) GetFeeID(ctx sdk.Context) (feeId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial fee ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeId)
	return feeId, nil
}

// Set the fee ID
func (k Keeper) SetFeeID(ctx sdk.Context, feeId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(feeId)
	store.Set(types.FeeIdKey, bz)
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

func (k Keeper) FeeContractExists(ctx sdk.Context, feeContractId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeContractKey(feeContractId))
}

func (k Keeper) GetFeeContract(ctx sdk.Context, feeContractId uint64) (types.FeeContract, sdk.Error) {
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

func (k Keeper) SubmitFeeContract(ctx sdk.Context, content types.FeeContractContent) (types.FeeContract, sdk.Error) {
	feeContractId, err := k.GetFeeContractID(ctx)
	if err != nil {
		return types.FeeContract{}, err
	}

	feeContract := types.NewFeeContract(feeContractId, content)

	k.SetFeeContract(ctx, feeContract)
	k.SetFeeContractID(ctx, feeContractId+1)

	return feeContract, nil
}

func (k Keeper) SetFeeContract(ctx sdk.Context, feeContract types.FeeContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractKey(feeContract.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(feeContract))
}

// GetFeeContractID gets the highest fee contract ID
func (k Keeper) GetFeeContractID(ctx sdk.Context) (feeContractId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeContractIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial fee contract ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContractId)
	return feeContractId, nil
}

// Set the fee contract ID
func (k Keeper) SetFeeContractID(ctx sdk.Context, feeContractId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(feeContractId)
	store.Set(types.FeeContractIdKey, bz)
}

func (k Keeper) SetFeeContractAuthorised(ctx sdk.Context, feeContractId uint64,
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

// -------------------------------------------------------- FeeContracts Charge

func (k Keeper) ChargeFee(ctx sdk.Context, bankKeeper bank.Keeper, feeContractId uint64) (charged bool, err sdk.Error) {

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

	// First assume that address will pay ChargeAmount, and calculate cumulative
	payAmount := feeData.ChargeAmount
	cumulative := fcData.CumulativeCharge.Add(payAmount)

	// Cumulative adjustment 1:
	// If first charge, increase to the minimum charge if the cumulative charge
	// is less than the minimum (applied on each denomination independently)
	if feeContract.IsFirstCharge() {
		for i, coin := range cumulative {
			minAmt := feeData.ChargeMinimum.AmountOf(coin.Denom)
			if !minAmt.IsZero() && minAmt.GT(coin.Amount) {
				cumulative[i] = sdk.NewCoin(coin.Denom, minAmt)
			}
		}
	}

	// Cumulative adjustment 2:
	// Reduce to the maximum charge if the cumulative charge is more than the
	// maximum (applied on each denomination independently)
	for i, coin := range cumulative {
		maxAmt := feeData.ChargeMaximum.AmountOf(coin.Denom)
		if !maxAmt.IsZero() && maxAmt.LT(coin.Amount) {
			cumulative[i] = sdk.NewCoin(coin.Denom, maxAmt)
		}
	}

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
