package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestKeeperIdReserver(t *testing.T) {
	_, k, _ := CreateTestInput()

	testFeeId1 := types.FeeIdPrefix + "test1"
	testFeeId2 := types.FeeIdPrefix + "test2"
	testFeeConId1 := types.FeeContractIdPrefix + "test1"
	testFeeConId2 := types.FeeContractIdPrefix + "test2"
	testSubId1 := types.SubscriptionIdPrefix + "test1"
	testSubId2 := types.SubscriptionIdPrefix + "test2"

	// Check that currently not reserved
	require.False(t, k.FeeIdReserved(testFeeId1))
	require.False(t, k.FeeIdReserved(testFeeId2))
	require.False(t, k.FeeContractIdReserved(testFeeConId1))
	require.False(t, k.FeeContractIdReserved(testFeeConId2))
	require.False(t, k.SubscriptionIdReserved(testSubId1))
	require.False(t, k.SubscriptionIdReserved(testSubId2))

	// Reserve 'test1' in general
	k.reservedIdPrefixes = []string{"test1"}

	// Check that 'test1' IDs now reserved and 'test2' IDs still unreserved
	require.True(t, k.FeeIdReserved(testFeeId1))
	require.False(t, k.FeeIdReserved(testFeeId2))
	require.True(t, k.FeeContractIdReserved(testFeeConId1))
	require.False(t, k.FeeContractIdReserved(testFeeConId2))
	require.True(t, k.SubscriptionIdReserved(testSubId1))
	require.False(t, k.SubscriptionIdReserved(testSubId2))
}

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

	// Check that currently discount is set as zero
	feeContract, err = k.GetFeeContract(ctx, validFeeContractId1)
	require.Nil(t, err)
	require.True(t, feeContract.Content.DiscountId.IsZero())

	// Grant FeeContract discounts
	err = k.GrantFeeDiscount(ctx, validFeeContractId1, sdk.NewUint(1))
	require.Nil(t, err)
	err = k.GrantFeeDiscount(ctx, validFeeContractId2, sdk.NewUint(2))
	require.Error(t, err) // since we used non-existent fee contract ID

	// Check that fee contract now has the discount ID (=1)
	feeContract, err = k.GetFeeContract(ctx, validFeeContractId1)
	require.Nil(t, err)
	require.Equal(t, feeContract.Content.DiscountId, sdk.NewUint(1))

	// Overwrite grant with a new discount grant
	err = k.GrantFeeDiscount(ctx, validFeeContractId1, sdk.NewUint(2))
	require.Nil(t, err)

	// Check that fee contract has the new discount ID (=2)
	feeContract, err = k.GetFeeContract(ctx, validFeeContractId1)
	require.Nil(t, err)
	require.Equal(t, feeContract.Content.DiscountId, sdk.NewUint(2))

	// Revoke FeeContract discounts
	err = k.RevokeFeeDiscount(ctx, validFeeContractId1)
	require.Nil(t, err)

	// Check that the discount ID is now back to zero
	feeContract, err = k.GetFeeContract(ctx, validFeeContractId1)
	require.Nil(t, err)
	require.True(t, feeContract.Content.DiscountId.IsZero())
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
