package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

var (
	shareAddr1     = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr1")))
	shareAddr2     = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr2")))
	feeCreatorAddr = sdk.AccAddress(crypto.AddressHash([]byte("feeCreatorAddr")))
	feePayerAddr   = sdk.AccAddress(crypto.AddressHash([]byte("feePayerAddr")))

	validChargeAmount, _  = sdk.ParseCoins("1ixo,2res")
	validChargeMinimum, _ = sdk.ParseCoins("3res")
	validChargeMaximum    = sdk.NewCoins()

	validFeeId1          = types.FeeIdPrefix + "f1"
	validFeeId2          = types.FeeIdPrefix + "f2"
	validFeeContractId1  = types.FeeContractIdPrefix + "fc1"
	validFeeContractId2  = types.FeeContractIdPrefix + "fc2"
	validSubscriptionId1 = types.SubscriptionIdPrefix + "s1"
	validSubscriptionId2 = types.SubscriptionIdPrefix + "s2"

	validDoubledChargeAmount, _ = sdk.ParseCoins("2ixo,4res")

	validDiscounts = types.NewDiscounts(
		types.NewDiscount(sdk.NewUint(1), sdk.MustNewDecFromStr("10")),
		types.NewDiscount(sdk.NewUint(2), sdk.MustNewDecFromStr("50")))
	tenPercentOffId   = validDiscounts[0].Id
	fiftyPercentOffId = validDiscounts[1].Id

	validDistribution = types.NewDistribution(
		types.NewDistributionShare(shareAddr1, sdk.NewDec(50)),
		types.NewDistributionShare(shareAddr2, sdk.NewDec(50)))

	validFee = types.NewFee(
		validFeeId1,
		validChargeAmount,
		validChargeMinimum,
		validChargeMaximum,
		validDiscounts,
		validDistribution)

	validDoubleChargeFee = types.NewFee(
		validFeeId1,
		validDoubledChargeAmount,
		validChargeMinimum,
		validChargeMaximum,
		validDiscounts,
		validDistribution)

	validFeeContract = types.NewFeeContractNoDiscount(
		validFeeContractId1, validFeeId1, feeCreatorAddr, feePayerAddr, false, true)
)

func ValidateVariables() sdk.Error {
	err := validDiscounts.Validate()
	if err != nil {
		return err
	}

	err = validDistribution.Validate()
	if err != nil {
		return err
	}

	err = validFee.Validate()
	if err != nil {
		return err
	}

	err = validDoubleChargeFee.Validate()
	if err != nil {
		return err
	}

	err = validFeeContract.Validate()
	if err != nil {
		return err
	}

	return nil
}

func CreateTestInput() (sdk.Context, Keeper, *codec.Codec) {
	if err := ValidateVariables(); err != nil {
		panic(err)
	}

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	actStoreKey := sdk.NewKVStoreKey(auth.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)

	_ = ms.LoadLatestVersion()
	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())

	cdc := codec.New()
	module.NewBasicManager(auth.AppModuleBasic{}).RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterConcrete(types.TestPeriod{}, "fees/TestPeriod", nil)

	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	pk1 := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	feesSubspace := pk1.Subspace(types.DefaultParamspace)

	accountKeeper := auth.NewAccountKeeper(cdc, actStoreKey, pk1.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk1.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, nil)

	keeper := NewKeeper(cdc, storeKey, feesSubspace, bankKeeper, nil)

	return ctx, keeper, cdc
}
