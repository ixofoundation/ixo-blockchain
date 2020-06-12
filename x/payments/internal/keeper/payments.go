package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
)

// -------------------------------------------------------- PaymentTemplates Get/Set

func (k Keeper) GetPaymentTemplateIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.PaymentTemplateKeyPrefix)
}

func (k Keeper) MustGetPaymentTemplateByKey(ctx sdk.Context, key []byte) types.PaymentTemplate {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("payment template not found")
	}

	bz := store.Get(key)
	var template types.PaymentTemplate
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &template)

	return template
}

func (k Keeper) PaymentTemplateExists(ctx sdk.Context, templateId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetPaymentTemplateKey(templateId))
}

func (k Keeper) GetPaymentTemplate(ctx sdk.Context, templateId string) (types.PaymentTemplate, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPaymentTemplateKey(templateId)

	bz := store.Get(key)
	if bz == nil {
		return types.PaymentTemplate{}, sdk.ErrInternal("invalid payment template")
	}

	var template types.PaymentTemplate
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &template)

	return template, nil
}

func (k Keeper) SetPaymentTemplate(ctx sdk.Context, template types.PaymentTemplate) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPaymentTemplateKey(template.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(template))
}

func (k Keeper) DiscountIdExists(ctx sdk.Context, templateId string, discountId sdk.Uint) (bool, sdk.Error) {
	// Get payment template
	template, err := k.GetPaymentTemplate(ctx, templateId)
	if err != nil {
		return false, err
	}

	// Search for discount ID
	for _, d := range template.Discounts {
		if d.Id.Equal(discountId) {
			return true, nil
		}
	}
	return false, nil
}

// -------------------------------------------------------- PaymentContracts Get/Set

func (k Keeper) GetPaymentContractIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.PaymentContractKeyPrefix)
}

func (k Keeper) MustGetPaymentContractByKey(ctx sdk.Context, key []byte) types.PaymentContract {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("payment contract not found")
	}

	bz := store.Get(key)
	var contract types.PaymentContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &contract)

	return contract
}

func (k Keeper) PaymentContractExists(ctx sdk.Context, contractId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetPaymentContractKey(contractId))
}

func (k Keeper) GetPaymentContract(ctx sdk.Context, contractId string) (types.PaymentContract, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPaymentContractKey(contractId)

	bz := store.Get(key)
	if bz == nil {
		return types.PaymentContract{}, sdk.ErrInternal("invalid payment contract")
	}

	var contract types.PaymentContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &contract)

	return contract, nil
}

func (k Keeper) SetPaymentContract(ctx sdk.Context, contract types.PaymentContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPaymentContractKey(contract.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(contract))
}

func (k Keeper) SetPaymentContractAuthorised(ctx sdk.Context, contractId string,
	authorised bool) sdk.Error {
	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return err
	}

	// If de-authorising, check if can be de-authorised
	if !authorised && !contract.CanDeauthorise {
		return types.ErrPaymentContractCannotBeDeauthorised(types.DefaultCodespace)
	}

	// Set authorised state
	contract.Authorised = authorised
	k.SetPaymentContract(ctx, contract)

	return nil
}

func (k Keeper) GrantDiscount(ctx sdk.Context, contractId string, discountId sdk.Uint) sdk.Error {
	// Get payment contract
	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return err
	}

	// Overwrite previous discount ID
	contract.DiscountId = discountId
	k.SetPaymentContract(ctx, contract)
	return nil
}

func (k Keeper) RevokeDiscount(ctx sdk.Context, contractId string) sdk.Error {
	// Get payment contract
	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return err
	}

	// Set discount ID to zero
	contract.DiscountId = sdk.ZeroUint()
	k.SetPaymentContract(ctx, contract)
	return nil
}

// -------------------------------------------------------- PaymentContracts payment

func applyDiscount(template types.PaymentTemplate, contract types.PaymentContract,
	payAmount sdk.Coins) (sdk.Coins, sdk.Error) {

	// No discounts held
	if contract.DiscountId.IsZero() {
		return payAmount, nil
	}

	// Get discount percentage to calculate discount amount. Any rounding
	// when multiplying means the payer receives a slightly smaller discount.
	discountPercent, err := template.GetDiscountPercent(contract.DiscountId)
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

func adjustForMinimums(template types.PaymentTemplate, contract types.PaymentContract,
	cumulative sdk.Coins) {
	// If first payment, increase to the minimum pay if the cumulative pay
	// is less than the minimum (applied on each denomination independently)
	if contract.IsFirstPayment() {
		for i, coin := range cumulative {
			minAmt := template.PaymentMinimum.AmountOf(coin.Denom)
			if !minAmt.IsZero() && minAmt.GT(coin.Amount) {
				cumulative[i] = sdk.NewCoin(coin.Denom, minAmt)
			}
		}
	}
}

func adjustForMaximums(template types.PaymentTemplate, cumulative sdk.Coins) {
	// Reduce to the maximum pay if the cumulative pay is more than the
	// maximum (applied on each denomination independently)
	for i, coin := range cumulative {
		maxAmt := template.PaymentMaximum.AmountOf(coin.Denom)
		if !maxAmt.IsZero() && maxAmt.LT(coin.Amount) {
			cumulative[i] = sdk.NewCoin(coin.Denom, maxAmt)
		}
	}
}

func (k Keeper) EffectPayment(ctx sdk.Context, bankKeeper bank.Keeper,
	contractId string) (effected bool, err sdk.Error) {

	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return false, err
	}

	template, err := k.GetPaymentTemplate(ctx, contract.PaymentTemplateId)
	if err != nil {
		return false, err
	}

	// Check if can effect payment (this is false if e.g. max pay has been reached)
	if !contract.CanEffectPayment(template) {
		return false, nil
	}

	// Assume payer will pay PaymentAmount, apply discount (if any),
	// and calculate initial cumulative (before adjustments)
	payAmount := template.PaymentAmount
	payAmount, err = applyDiscount(template, contract, payAmount)
	if err != nil {
		return false, err
	}
	cumulative := contract.CumulativePay.Add(payAmount)

	// In-place cumulative adjustments (i.e. considering minimums and maximums)
	adjustForMinimums(template, contract, cumulative)
	adjustForMaximums(template, cumulative)

	// Find actual pay from adjusted cumulative:
	//    adjustedCumul = previousCumul + actualPay
	// => actualPay = adjustedCumul - previousCumul
	pay := cumulative.Sub(contract.CumulativePay)

	// Stop if payer doesn't have enough coins. However, this is not considered
	// an error but the caller should be looking at the 'effected' bool result
	if !bankKeeper.HasCoins(ctx, contract.Payer, pay) {
		return false, nil
	}

	// Total input is pay plus current remainder in PayRemainderPool
	inputFromPayRemainderPool := contract.CurrentRemainder
	totalInputAmount := pay.Add(inputFromPayRemainderPool)

	// Calculate list of outputs and calculate the total output to payees based
	// on the calculated wallet distributions
	var outputToPayees sdk.Coins
	var outputs []bank.Output
	distributions := template.WalletDistribution.GetDistributionsFor(totalInputAmount)
	for i, share := range distributions {
		// Get integer output
		outputAmt, _ := share.TruncateDecimal()

		// If amount not zero, update total and add as output
		if !outputAmt.IsZero() {
			outputToPayees = outputToPayees.Add(outputAmt)
			address := template.WalletDistribution[i].Address
			outputs = append(outputs, bank.NewOutput(address, outputAmt))
		}
	}

	// Remainder (not output to payees) goes to PayRemainderPool if not zero
	outputToPayRemainderPool := totalInputAmount.Sub(outputToPayees)
	if !outputToPayRemainderPool.IsZero() {
		payRemainderPoolAddr := supply.NewModuleAddress(types.PayRemainderPool)
		outputs = append(outputs, bank.NewOutput(payRemainderPoolAddr, outputToPayRemainderPool))
	}

	// Construct list of inputs (pay and from PayRemainderPool if non zero)
	inputs := []bank.Input{bank.NewInput(contract.Payer, pay)}
	if !inputFromPayRemainderPool.IsZero() {
		payRemainderPoolAddr := supply.NewModuleAddress(types.PayRemainderPool)
		inputs = append(inputs, bank.NewInput(payRemainderPoolAddr, inputFromPayRemainderPool))
	}

	// Distribute the payment according to the outputs
	err = bankKeeper.InputOutputCoins(ctx, inputs, outputs)
	if err != nil {
		return false, err
	}

	// Update and save payment contract
	contract.CumulativePay = contract.CumulativePay.Add(pay)
	contract.CurrentRemainder = contract.CurrentRemainder.Add(
		outputToPayRemainderPool).Sub(inputFromPayRemainderPool)
	k.SetPaymentContract(ctx, contract)

	return true, nil
}
