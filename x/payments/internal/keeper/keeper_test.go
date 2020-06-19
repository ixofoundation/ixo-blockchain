package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestKeeperIdReserver(t *testing.T) {
	_, k, _ := CreateTestInput()

	testTemplateId1 := types.PaymentTemplateIdPrefix + "test1"
	testTemplateId2 := types.PaymentTemplateIdPrefix + "test2"
	testContractId1 := types.PaymentContractIdPrefix + "test1"
	testContractId2 := types.PaymentContractIdPrefix + "test2"
	testSubId1 := types.SubscriptionIdPrefix + "test1"
	testSubId2 := types.SubscriptionIdPrefix + "test2"

	// Check that currently not reserved
	require.False(t, k.PaymentTemplateIdReserved(testTemplateId1))
	require.False(t, k.PaymentTemplateIdReserved(testTemplateId2))
	require.False(t, k.PaymentContractIdReserved(testContractId1))
	require.False(t, k.PaymentContractIdReserved(testContractId2))
	require.False(t, k.SubscriptionIdReserved(testSubId1))
	require.False(t, k.SubscriptionIdReserved(testSubId2))

	// Reserve 'test1' in general
	k.reservedIdPrefixes = []string{"test1"}

	// Check that 'test1' IDs now reserved and 'test2' IDs still unreserved
	require.True(t, k.PaymentTemplateIdReserved(testTemplateId1))
	require.False(t, k.PaymentTemplateIdReserved(testTemplateId2))
	require.True(t, k.PaymentContractIdReserved(testContractId1))
	require.False(t, k.PaymentContractIdReserved(testContractId2))
	require.True(t, k.SubscriptionIdReserved(testSubId1))
	require.False(t, k.SubscriptionIdReserved(testSubId2))
}

func TestKeeperSetGet(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Check PaymentTemplate, PaymentContract, Subscription, existence
	_, err := k.GetPaymentTemplate(ctx, "dummyId")
	require.NotNil(t, err)
	_, err = k.GetPaymentContract(ctx, "dummyId")
	require.NotNil(t, err)
	_, err = k.GetSubscription(ctx, "dummyId")
	require.NotNil(t, err)

	// Submitted PaymentTemplate
	template := validTemplate
	k.SetPaymentTemplate(ctx, template)

	// Check PaymentTemplate
	templateGet, err := k.GetPaymentTemplate(ctx, template.Id)
	require.Nil(t, err)
	require.Equal(t, template.Id, templateGet.Id)

	// Submitted PaymentContract
	contract := validContract
	k.SetPaymentContract(ctx, contract)

	// Check PaymentContract
	contractGet, err := k.GetPaymentContract(ctx, contract.Id)
	require.Nil(t, err)
	require.Equal(t, contract.Id, contractGet.Id)

	// Create BlockPeriod Subscription
	blockPeriod := types.NewBlockPeriod(100, 0)
	blockSubscription := types.NewSubscription(validSubscriptionId1,
		contract.Id, sdk.NewUint(10), &blockPeriod)
	k.SetSubscription(ctx, blockSubscription)

	// Check BlockPeriod Subscription
	blockSubscriptionGet, err := k.GetSubscription(ctx, blockSubscription.Id)
	require.Nil(t, err)
	require.Equal(t, blockSubscription.Id, blockSubscriptionGet.Id)

	// Submitted TimePeriod Subscription
	duration, _ := time.ParseDuration("2h")
	timePeriod := types.NewTimePeriod(duration, time.Now())
	timeSubscription := types.NewSubscription(validSubscriptionId2,
		contract.Id, sdk.NewUint(10), &timePeriod)
	k.SetSubscription(ctx, timeSubscription)

	// Check TimePeriod Subscription
	timeSubscriptionGet, err := k.GetSubscription(ctx, timeSubscription.Id)
	require.Nil(t, err)
	require.Equal(t, timeSubscription.Id, timeSubscriptionGet.Id)

	// Check that currently discount is set as zero
	contract, err = k.GetPaymentContract(ctx, validPaymentContractId1)
	require.Nil(t, err)
	require.True(t, contract.DiscountId.IsZero())

	// Grant PaymentContract discounts
	err = k.GrantDiscount(ctx, validPaymentContractId1, sdk.NewUint(1))
	require.Nil(t, err)
	err = k.GrantDiscount(ctx, validPaymentContractId2, sdk.NewUint(2))
	require.Error(t, err) // since we used non-existent payment contract ID

	// Check that payment contract now has the discount ID (=1)
	contract, err = k.GetPaymentContract(ctx, validPaymentContractId1)
	require.Nil(t, err)
	require.Equal(t, contract.DiscountId, sdk.NewUint(1))

	// Overwrite grant with a new discount grant
	err = k.GrantDiscount(ctx, validPaymentContractId1, sdk.NewUint(2))
	require.Nil(t, err)

	// Check that payment contract has the new discount ID (=2)
	contract, err = k.GetPaymentContract(ctx, validPaymentContractId1)
	require.Nil(t, err)
	require.Equal(t, contract.DiscountId, sdk.NewUint(2))

	// Revoke PaymentContract discounts
	err = k.RevokeDiscount(ctx, validPaymentContractId1)
	require.Nil(t, err)

	// Check that the discount ID is now back to zero
	contract, err = k.GetPaymentContract(ctx, validPaymentContractId1)
	require.Nil(t, err)
	require.True(t, contract.DiscountId.IsZero())
}

func TestKeeperEffectPayment(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Create and submit PaymentTemplate and PaymentContract
	template := validTemplate
	contract := validContract
	k.SetPaymentTemplate(ctx, template)
	k.SetPaymentContract(ctx, contract)

	// Set payer balance
	balance, err2 := sdk.ParseCoins("10uixo,10res")
	require.Nil(t, err2)
	err := k.bankKeeper.SetCoins(ctx, contract.Payer, balance)
	require.Nil(t, err)

	// At this point, cumulative: /
	// PayAmt:  1uixo, 2res
	// PayMin:  3res
	// PayMax:  /
	// Next payment expected to be: 1uixo, 3res (3res due to PayMin)
	// Updated balance: 9uixo, 7res

	// Effect payment
	effected, err := k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)

	// Check balance
	newBalance := k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 := sdk.ParseCoins("9uixo,7res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 1uixo, 3res
	// PayAmt:  1uixo, 2res
	// PayMin:  3res
	// PayMax:  /
	// Next payment expected to be: 1uixo, 2res (no effect from PayMin)
	// Updated balance: 8uixo, 5res

	// Effect payment
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 = sdk.ParseCoins("8uixo,5res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 2uixo, 5res
	// Next payment expected to be: 1uixo, 2res
	// Updated balance: 7uixo, 3res

	// Effect payment
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 = sdk.ParseCoins("7uixo,3res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// Two more payments will cause an error
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.False(t, effected)
}

func TestKeeperEffectPaymentWithDiscounts(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Create and submit PaymentTemplate (!!double pay!!) and PaymentContract
	template := validDoublePayTemplate
	contract := validContract
	k.SetPaymentTemplate(ctx, template)
	k.SetPaymentContract(ctx, contract)

	// Set payer balance
	balance, err2 := sdk.ParseCoins("10uixo,10res")
	require.Nil(t, err2)
	err := k.bankKeeper.SetCoins(ctx, contract.Payer, balance)
	require.Nil(t, err)

	// Set discount
	err = k.GrantDiscount(ctx, contract.Id, fiftyPercentOffId)
	require.Nil(t, err)

	// At this point, cumulative: /
	// PayAmt:  2uixo, 4res (discounted to 1uixo, 2res)
	// PayMin:  3res
	// PayMax:  /
	// Next payment expected to be: 1uixo, 3res (3res due to PayMin)
	// Updated balance: 9uixo, 7res

	// Effect payment
	effected, err := k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)

	// Check balance
	newBalance := k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 := sdk.ParseCoins("9uixo,7res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 1uixo, 3res
	// PayAmt:  2uixo, 4res (discounted to 1uixo, 2res)
	// PayMin:  3res
	// PayMax:  /
	// Next payment expected to be: 1uixo, 2res (no effect from PayMin)
	// Updated balance: 8uixo, 5res

	// Effect payment
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 = sdk.ParseCoins("8uixo,5res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 2uixo, 5res
	// Next payment expected to be: 1uixo, 2res
	// Updated balance: 7uixo, 3res

	// Effect payment
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 = sdk.ParseCoins("7uixo,3res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// Two more payments will cause an error
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.True(t, effected)
	effected, err = k.EffectPayment(ctx, k.bankKeeper, contract.Id)
	require.Nil(t, err)
	require.False(t, effected)
}

func TestKeeperEffectSubscriptionPayment(t *testing.T) {
	ctx, k, _ := CreateTestInput()

	// Create and submit PaymentTemplate and PaymentContract
	template := validTemplate
	contract := validContract
	k.SetPaymentTemplate(ctx, template)
	k.SetPaymentContract(ctx, contract)

	// Create and submit Subscription
	testPeriod := types.NewTestPeriod(100, 0)
	testSubscription := types.NewSubscription(validSubscriptionId1,
		contract.Id, sdk.NewUint(10), testPeriod)
	k.SetSubscription(ctx, testSubscription)

	// Set payer balance
	balance, err2 := sdk.ParseCoins("10uixo,10res")
	require.Nil(t, err2)
	err := k.bankKeeper.SetCoins(ctx, contract.Payer, balance)
	require.Nil(t, err)

	// At this point, cumulative: /
	// PayAmt:  1uixo, 2res
	// PayMin:  3res
	// PayMax:  /
	// Next payment expected to be: 1uixo, 3res (3res due to PayMin)
	// Updated balance: 9uixo, 7res

	// Effect subscription payment
	err = k.EffectSubscriptionPayment(ctx, testSubscription.Id)
	require.Nil(t, err)

	// Check balance
	newBalance := k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 := sdk.ParseCoins("9uixo,7res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 1uixo, 3res
	// PayAmt:  1uixo, 2res
	// PayMin:  3res
	// PayMax:  /
	// Next payment expected to be: 1uixo, 2res (no effect from PayMin)
	// Updated balance: 8uixo, 5res

	// Effect subscription payment
	err = k.EffectSubscriptionPayment(ctx, testSubscription.Id)
	require.Nil(t, err)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 = sdk.ParseCoins("8uixo,5res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// At this point, cumulative: 2uixo, 5res
	// Next payment expected to be: 1uixo, 2res
	// Updated balance: 7uixo, 3res

	// Effect subscription payment
	err = k.EffectSubscriptionPayment(ctx, testSubscription.Id)
	require.Nil(t, err)

	// Check balance
	newBalance = k.bankKeeper.GetCoins(ctx, contract.Payer)
	expected, err2 = sdk.ParseCoins("7uixo,3res")
	require.Nil(t, err2)
	require.Equal(t, expected.String(), newBalance.String())

	// Two more payments will cause an error
	err = k.EffectSubscriptionPayment(ctx, testSubscription.Id)
	require.Nil(t, err)
	err = k.EffectSubscriptionPayment(ctx, testSubscription.Id)
	require.Nil(t, err)
}
