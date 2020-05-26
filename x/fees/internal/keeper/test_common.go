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

	validDoubledChargeAmount, _ = sdk.ParseCoins("2ixo,4res")

	validDiscounts = types.NewDiscounts(
		types.NewDiscount(1, sdk.MustNewDecFromStr("10")),
		types.NewDiscount(2, sdk.MustNewDecFromStr("50")))
	tenPercentOffId   = validDiscounts[0].Id
	fiftyPercentOffId = validDiscounts[1].Id

	validDistribution = types.NewDistribution(
		types.NewDistributionShare(shareAddr1, sdk.NewDec(50)),
		types.NewDistributionShare(shareAddr2, sdk.NewDec(50)))

	validFeeContent = types.NewFeeContent(
		validChargeAmount,
		validChargeMinimum,
		validChargeMaximum,
		validDiscounts,
		validDistribution)

	validDoubleChargeFeeContent = types.NewFeeContent(
		validDoubledChargeAmount,
		validChargeMinimum,
		validChargeMaximum,
		validDiscounts,
		validDistribution)

	validFeeContractContent = types.NewFeeContractContent(
		1, feeCreatorAddr, feePayerAddr, false, true)
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

	err = validFeeContent.Validate()
	if err != nil {
		return err
	}

	err = types.NewFee(1, validFeeContent).Validate()
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
	cdc.RegisterConcrete(types.TestSubscriptionContent{}, "fees/TesSubscriptionContent", nil)

	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	pk1 := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	feesSubspace := pk1.Subspace(types.DefaultParamspace)

	accountKeeper := auth.NewAccountKeeper(cdc, actStoreKey, pk1.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk1.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, nil)

	keeper := NewKeeper(cdc, storeKey, feesSubspace, bankKeeper)

	return ctx, keeper, cdc
}
