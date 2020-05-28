package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"testing"
	"time"
)

func TestKeeperSetGet(t *testing.T) {
	startingDiscountHoldersId := uint64(1)

	ctx, k, _ := CreateTestInput()

	// Check Fee, FeeContract, Subscription, DiscountHolders existence
	feeGet, err := k.GetFee(ctx, "dummyId")
	require.NotNil(t, err)
	feeContractGet, err := k.GetFeeContract(ctx, "dummyId")
	require.NotNil(t, err)
	blockSubscriptionGet, err := k.GetSubscription(ctx, "dummyId")
	require.NotNil(t, err)
	discountHoldersIter := k.GetFeeDiscountHoldersIterator(ctx, "dummyId", startingDiscountHoldersId)
	require.False(t, discountHoldersIter.Valid())
	discountHoldersIter = k.GetFeeDiscountsHoldersIterator(ctx, "dummyId")
	require.False(t, discountHoldersIter.Valid())
	discountHoldersIter = k.GetFeesDiscountsHoldersIterator(ctx)
	require.False(t, discountHoldersIter.Valid())

	// Submitted Fee
	fee := types.NewFee(validFeeId1, validFeeContent)
	k.SetFee(ctx, fee)

	// Check Fee
	feeGet, err = k.GetFee(ctx, fee.Id)
	require.Nil(t, err)
	require.Equal(t, fee.Id, feeGet.Id)

	// Submitted FeeContract
	feeContract := types.NewFeeContract(validFeeContractId1, validFeeContractContent)
	k.SetFeeContract(ctx, feeContract)

	// Check FeeContract
	feeContractGet, err = k.GetFeeContract(ctx, feeContract.Id)
	require.Nil(t, err)
	require.Equal(t, feeContract.Id, feeContractGet.Id)

	// Submitted BlockSubscription
	blockSubscriptionContent := types.NewBlockSubscriptionContent(
		feeContract.Id, sdk.NewUint(10), 100, 0)
	blockSubscription := types.NewSubscription(
		validFeeContractId1, &blockSubscriptionContent)
	k.SetSubscription(ctx, blockSubscription)

	// Check BlockSubscription
	blockSubscriptionGet, err = k.GetSubscription(ctx, blockSubscription.Id)
	require.Nil(t, err)
	require.Equal(t, blockSubscription.Id, blockSubscriptionGet.Id)

	// Submitted TimeSubscription
	duration, _ := time.ParseDuration("2h")
	timeSubscriptionContent := types.NewTimeSubscriptionContent(
		feeContract.Id, sdk.NewUint(10), duration, time.Now())
	timeSubscription := types.NewSubscription(
		validSubscriptionId2, &timeSubscriptionContent)
	k.SetSubscription(ctx, timeSubscription)

	// Check TimeSubscription
	timeSubscriptionGet, err := k.GetSubscription(ctx, timeSubscription.Id)
	require.Nil(t, err)
	require.Equal(t, timeSubscription.Id, timeSubscriptionGet.Id)

	// Set Discount Holder
	holder1 := sdk.AccAddress(crypto.AddressHash([]byte("holder1")))
	holder2 := sdk.AccAddress(crypto.AddressHash([]byte("holder2")))
	holder3 := sdk.AccAddress(crypto.AddressHash([]byte("holder3")))
	holder4 := sdk.AccAddress(crypto.AddressHash([]byte("holder4")))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(validFeeId1, 1, holder1))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(validFeeId1, 2, holder2))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(validFeeId2, 3, holder3))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder("someFeeId", 4, holder4))

	// Check that 4 discount holders in general
	discountHoldersIter = k.GetFeesDiscountsHoldersIterator(ctx)
	for i := 0; i < 4; i++ {
		_ = k.MustGetDiscountHolderByKey(ctx, discountHoldersIter.Key())
		discountHoldersIter.Next()
	}
	require.False(t, discountHoldersIter.Valid())

	// Check that 2 discount holders for validFeeId1
	discountHoldersIter = k.GetFeeDiscountsHoldersIterator(ctx, validFeeId1)
	for i := 0; i < 2; i++ {
		_ = k.MustGetDiscountHolderByKey(ctx, discountHoldersIter.Key())
		discountHoldersIter.Next()
	}
	require.False(t, discountHoldersIter.Valid())

	// Check that 1 discount holder for validFeeId1 and discount id 2
	discountId := uint64(2)
	discountHoldersIter = k.GetFeeDiscountHoldersIterator(ctx, validFeeId1, discountId)
	discountholder := k.MustGetDiscountHolderByKey(ctx, discountHoldersIter.Key())
	require.Equal(t, discountId, discountholder.DiscountId)
	discountHoldersIter.Next()
	require.False(t, discountHoldersIter.Valid())

	// Check that 0 discount holders for validFeeId1 and discount id 3
	discountHoldersIter = k.GetFeeDiscountHoldersIterator(ctx, validFeeId1, 3)
	require.False(t, discountHoldersIter.Valid())
}

func TestKeeperChargeFee(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Create and submit Fee and FeeContract
	fee := types.NewFee(validFeeId1, validFeeContent)
	feeContract := types.NewFeeContract(validFeeContractId1, validFeeContractContent)
	k.SetFee(ctx, fee)
	k.SetFeeContract(ctx, feeContract)

	// Set payer balance
	balance, err2 := sdk.ParseCoins("10ixo,10res")
	require.Nil(t, err2)
	err := k.bankKeeper.SetCoins(ctx, feeContract.Content.Payer, balance)
	require.Nil(t, err)

	// At this point, cumulative: /
	// ChargeAmt:  1ixo, 2res
	// ChargeMin:  3res
	// ChargeMax:  /
	// Next charge expected to be: 1ixo, 3res (3res due to ChargeMin)
	// Updated balance: 9ixo, 7res

	// Charge fee
	charged, err := k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)

	// Check balance
	newBalance := k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 := sdk.ParseCoins("9ixo,7res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 1ixo, 3res
	// ChargeAmt:  1ixo, 2res
	// ChargeMin:  3res
	// ChargeMax:  /
	// Next charge expected to be: 1ixo, 2res (no effect from ChargeMin)
	// Updated balance: 8ixo, 5res

	// Charge fee
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 = sdk.ParseCoins("8ixo,5res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 2ixo, 5res
	// Next charge expected to be: 1ixo, 2res
	// Updated balance: 7ixo, 3res

	// Charge fee
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 = sdk.ParseCoins("7ixo,3res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// Two more charges will cause an error
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.False(t, charged)
}

func TestKeeperChargeFeeWithDiscounts(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Create and submit Fee (!!double charge!!) and FeeContract
	fee := types.NewFee(validFeeId1, validDoubleChargeFeeContent)
	feeContract := types.NewFeeContract(validFeeContractId1, validFeeContractContent)
	k.SetFee(ctx, fee)
	k.SetFeeContract(ctx, feeContract)

	// Set payer balance
	balance, err2 := sdk.ParseCoins("10ixo,10res")
	require.Nil(t, err2)
	err := k.bankKeeper.SetCoins(ctx, feeContract.Content.Payer, balance)
	require.Nil(t, err)

	// Set discount
	discountHolder := types.NewDiscountHolder(fee.Id, fiftyPercentOffId, feeContract.Content.Payer)
	k.SetDiscountHolder(ctx, discountHolder)

	// At this point, cumulative: /
	// ChargeAmt:  2ixo, 4res (discounted to 1ixo, 2res)
	// ChargeMin:  3res
	// ChargeMax:  /
	// Next charge expected to be: 1ixo, 3res (3res due to ChargeMin)
	// Updated balance: 9ixo, 7res

	// Charge fee
	charged, err := k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)

	// Check balance
	newBalance := k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 := sdk.ParseCoins("9ixo,7res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 1ixo, 3res
	// ChargeAmt:  2ixo, 4res (discounted to 1ixo, 2res)
	// ChargeMin:  3res
	// ChargeMax:  /
	// Next charge expected to be: 1ixo, 2res (no effect from ChargeMin)
	// Updated balance: 8ixo, 5res

	// Charge fee
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 = sdk.ParseCoins("8ixo,5res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 2ixo, 5res
	// Next charge expected to be: 1ixo, 2res
	// Updated balance: 7ixo, 3res

	// Charge fee
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 = sdk.ParseCoins("7ixo,3res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// Two more charges will cause an error
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.True(t, charged)
	charged, err = k.ChargeFee(ctx, k.bankKeeper, feeContract.Id)
	require.Nil(t, err)
	require.False(t, charged)
}

func TestKeeperChargeSubscriptionFee(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Create and submit Fee and FeeContract
	fee := types.NewFee(validFeeId1, validFeeContent)
	feeContract := types.NewFeeContract(validFeeContractId1, validFeeContractContent)
	k.SetFee(ctx, fee)
	k.SetFeeContract(ctx, feeContract)

	// Create and submit TestSubscription
	testSubscriptionContent := types.NewTestSubscriptionContent(
		feeContract.Id, sdk.NewUint(10), 100, 0)
	testSubscription := types.NewSubscription(validSubscriptionId1, testSubscriptionContent)
	k.SetSubscription(ctx, testSubscription)

	// Set payer balance
	balance, err2 := sdk.ParseCoins("10ixo,10res")
	require.Nil(t, err2)
	err := k.bankKeeper.SetCoins(ctx, feeContract.Content.Payer, balance)
	require.Nil(t, err)

	// At this point, cumulative: /
	// ChargeAmt:  1ixo, 2res
	// ChargeMin:  3res
	// ChargeMax:  /
	// Next charge expected to be: 1ixo, 3res (3res due to ChargeMin)
	// Updated balance: 9ixo, 7res

	// Charge subscription fee
	err = k.ChargeSubscriptionFee(ctx, testSubscription.Id)
	require.Nil(t, err)

	// Check balance
	newBalance := k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 := sdk.ParseCoins("9ixo,7res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 1ixo, 3res
	// ChargeAmt:  1ixo, 2res
	// ChargeMin:  3res
	// ChargeMax:  /
	// Next charge expected to be: 1ixo, 2res (no effect from ChargeMin)
	// Updated balance: 8ixo, 5res

	// Charge subscription fee
	err = k.ChargeSubscriptionFee(ctx, testSubscription.Id)
	require.Nil(t, err)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 = sdk.ParseCoins("8ixo,5res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 2ixo, 5res
	// Next charge expected to be: 1ixo, 2res
	// Updated balance: 7ixo, 3res

	// Charge subscription fee
	err = k.ChargeSubscriptionFee(ctx, testSubscription.Id)
	require.Nil(t, err)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, feeContract.Content.Payer)
	expected, err2 = sdk.ParseCoins("7ixo,3res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// Two more charges will cause an error
	err = k.ChargeSubscriptionFee(ctx, testSubscription.Id)
	require.Nil(t, err)
	err = k.ChargeSubscriptionFee(ctx, testSubscription.Id)
	require.Nil(t, err)
}
