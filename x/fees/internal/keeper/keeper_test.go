package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"testing"
	"time"
)

func TestKeeper(t *testing.T) {
	startingFeeId := uint64(1)
	startingFeeContractId := uint64(1)
	startingSubscriptionId := uint64(1)
	startingDiscountHoldersId := uint64(1) // does not increment automatically

	ctx, k, cdc := CreateTestInput()
	cdc.RegisterInterface((*types.SubscriptionContent)(nil), nil)
	cdc.RegisterConcrete(types.BlockSubscriptionContent{}, "fees/BlockSubscriptionContent", nil)
	cdc.RegisterConcrete(types.TimeSubscriptionContent{}, "fees/TimeSubscriptionContent", nil)
	k.SetFeeID(ctx, 1)
	k.SetFeeContractID(ctx, 1)
	k.SetSubscriptionID(ctx, 1)

	// Check Fee, FeeContract, Subscription, DiscountHolders existence
	feeGet, err := k.GetFee(ctx, startingFeeId)
	require.NotNil(t, err)
	feeContractGet, err := k.GetFeeContract(ctx, startingFeeContractId)
	require.NotNil(t, err)
	blockSubscriptionGet, err := k.GetSubscription(ctx, startingSubscriptionId)
	require.NotNil(t, err)
	discountHoldersIter := k.GetFeeDiscountHoldersIterator(ctx, startingFeeId, startingDiscountHoldersId)
	require.False(t, discountHoldersIter.Valid())
	discountHoldersIter = k.GetFeeDiscountsHoldersIterator(ctx, startingFeeId)
	require.False(t, discountHoldersIter.Valid())
	discountHoldersIter = k.GetFeesDiscountsHoldersIterator(ctx)
	require.False(t, discountHoldersIter.Valid())

	// Submitted Fee
	fee, err := k.SubmitFee(ctx, validFeeContent)
	if err != nil {
		panic(err.Error())
	}

	// Check Fee
	require.Equal(t, startingFeeId, fee.Id)
	feeGet, err = k.GetFee(ctx, fee.Id)
	require.Nil(t, err)
	require.Equal(t, fee.Id, feeGet.Id)

	// Submitted FeeContract
	feeContract, err := k.SubmitFeeContract(ctx, validFeeContractContent)
	if err != nil {
		panic(err.Error())
	}

	// Check FeeContract
	require.Equal(t, startingFeeContractId, feeContract.Id)
	feeContractGet, err = k.GetFeeContract(ctx, feeContract.Id)
	require.Nil(t, err)
	require.Equal(t, feeContract.Id, feeContractGet.Id)

	// Submitted BlockSubscription
	blockSubscriptionContent := types.NewBlockSubscriptionContent(
		feeContract.Id, sdk.NewUint(10), 100, 0)
	blockSubscription, err := k.SubmitSubscription(ctx, &blockSubscriptionContent)
	if err != nil {
		panic(err.Error())
	}

	// Check BlockSubscription
	require.Equal(t, startingSubscriptionId, blockSubscription.Id)
	blockSubscriptionGet, err = k.GetSubscription(ctx, blockSubscription.Id)
	require.Nil(t, err)
	require.Equal(t, blockSubscription.Id, blockSubscriptionGet.Id)

	// Submitted TimeSubscription
	duration, _ := time.ParseDuration("2h")
	timeSubscriptionContent := types.NewTimeSubscriptionContent(
		feeContract.Id, sdk.NewUint(10), duration, time.Now())
	timeSubscription, err := k.SubmitSubscription(ctx, &timeSubscriptionContent)
	if err != nil {
		panic(err.Error())
	}

	// Check TimeSubscription
	require.Equal(t, startingSubscriptionId+1, timeSubscription.Id)
	timeSubscriptionGet, err := k.GetSubscription(ctx, timeSubscription.Id)
	require.Nil(t, err)
	require.Equal(t, timeSubscription.Id, timeSubscriptionGet.Id)

	// Set Discount Holder
	holder1 := sdk.AccAddress(crypto.AddressHash([]byte("holder1")))
	holder2 := sdk.AccAddress(crypto.AddressHash([]byte("holder2")))
	holder3 := sdk.AccAddress(crypto.AddressHash([]byte("holder3")))
	holder4 := sdk.AccAddress(crypto.AddressHash([]byte("holder4")))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(fee.Id, 1, holder1))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(fee.Id, 2, holder2))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(fee.Id+1, 3, holder3))
	k.SetDiscountHolder(ctx, types.NewDiscountHolder(fee.Id+2, 4, holder4))

	// Check that 4 discount holders in general
	discountHoldersIter = k.GetFeesDiscountsHoldersIterator(ctx)
	for i := 0; i < 4; i++ {
		_ = k.MustGetDiscountHolderByKey(ctx, discountHoldersIter.Key())
		discountHoldersIter.Next()
	}
	require.False(t, discountHoldersIter.Valid())

	// Check that 2 discount holders for fee.Id
	discountHoldersIter = k.GetFeeDiscountsHoldersIterator(ctx, fee.Id)
	for i := 0; i < 2; i++ {
		_ = k.MustGetDiscountHolderByKey(ctx, discountHoldersIter.Key())
		discountHoldersIter.Next()
	}
	require.False(t, discountHoldersIter.Valid())

	// Check that 1 discount holder for fee.Id and discount id 2
	discountId := uint64(2)
	discountHoldersIter = k.GetFeeDiscountHoldersIterator(ctx, fee.Id, discountId)
	discountholder := k.MustGetDiscountHolderByKey(ctx, discountHoldersIter.Key())
	require.Equal(t, discountId, discountholder.DiscountId)
	discountHoldersIter.Next()
	require.False(t, discountHoldersIter.Valid())

	// Check that 0 discount holders for fee.Id and discount id 3
	discountHoldersIter = k.GetFeeDiscountHoldersIterator(ctx, fee.Id, 3)
	require.False(t, discountHoldersIter.Valid())
}
