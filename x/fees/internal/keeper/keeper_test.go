package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestKeeperSetGet(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Check Fee, FeeContract, Subscription, existence
	feeGet, err := k.GetFee(ctx, "dummyId")
	require.NotNil(t, err)
	feeContractGet, err := k.GetFeeContract(ctx, "dummyId")
	require.NotNil(t, err)
	blockSubscriptionGet, err := k.GetSubscription(ctx, "dummyId")
	require.NotNil(t, err)

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

	// Grant FeeContract discounts
	err = k.GrantFeeDiscount(ctx, validFeeContractId1, 1)
	require.Nil(t, err)
	err = k.GrantFeeDiscount(ctx, validFeeContractId1, 2)
	require.Nil(t, err)
	err = k.GrantFeeDiscount(ctx, validFeeContractId2, 3)
	require.Error(t, err) // since we used non-existent fee contract ID

	// Check that fee contract has the two discount IDs
	feeContract, err = k.GetFeeContract(ctx, validFeeContractId1)
	require.Nil(t, err)
	require.Len(t, feeContract.Content.DiscountIds, 2)
	require.Equal(t, feeContract.Content.DiscountIds[0], uint64(1))
	require.Equal(t, feeContract.Content.DiscountIds[1], uint64(2))

	// Revoke FeeContract discounts
	err = k.RevokeFeeDiscount(ctx, validFeeContractId1, 1)
	require.Nil(t, err)
	err = k.RevokeFeeDiscount(ctx, validFeeContractId1, 4)
	require.Nil(t, err) // invalid discount ID not considered an error
	err = k.RevokeFeeDiscount(ctx, validFeeContractId2, 3)
	require.Error(t, err) // since we used non-existent fee contract ID

	// Check that fee contract has just one discount now
	feeContract, err = k.GetFeeContract(ctx, validFeeContractId1)
	require.Nil(t, err)
	require.Len(t, feeContract.Content.DiscountIds, 1)
	require.Equal(t, feeContract.Content.DiscountIds[0], uint64(2))
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
	err = k.GrantFeeDiscount(ctx, feeContract.Id, fiftyPercentOffId)
	require.Nil(t, err)

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
