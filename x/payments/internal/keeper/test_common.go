package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

var (
	shareAddr1          = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr1")))
	shareAddr2          = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr2")))
	templateCreatorAddr = sdk.AccAddress(crypto.AddressHash([]byte("templateCreatorAddr")))
	payerAddr           = sdk.AccAddress(crypto.AddressHash([]byte("payerAddr")))

	validPaymentAmount, _  = sdk.ParseCoins("1uixo,2res")
	validPaymentMinimum, _ = sdk.ParseCoins("3res")
	validPaymentMaximum    = sdk.NewCoins()

	validTemplateId1        = types.PaymentTemplateIdPrefix + "pt1"
	validTemplateId2        = types.PaymentTemplateIdPrefix + "pt2"
	validPaymentContractId1 = types.PaymentContractIdPrefix + "pc1"
	validPaymentContractId2 = types.PaymentContractIdPrefix + "pc2"
	validSubscriptionId1    = types.SubscriptionIdPrefix + "s1"
	validSubscriptionId2    = types.SubscriptionIdPrefix + "s2"

	validDoubledPaymentAmount, _ = sdk.ParseCoins("2uixo,4res")

	validDiscounts = types.NewDiscounts(
		types.NewDiscount(sdk.NewUint(1), sdk.MustNewDecFromStr("10")),
		types.NewDiscount(sdk.NewUint(2), sdk.MustNewDecFromStr("50")))
	tenPercentOffId   = validDiscounts[0].Id
	fiftyPercentOffId = validDiscounts[1].Id

	validDistribution = types.NewDistribution(
		types.NewDistributionShare(shareAddr1, sdk.NewDec(50)),
		types.NewDistributionShare(shareAddr2, sdk.NewDec(50)))

	validTemplate = types.NewPaymentTemplate(
		validTemplateId1,
		validPaymentAmount,
		validPaymentMinimum,
		validPaymentMaximum,
		validDiscounts,
		validDistribution)

	validDoublePayTemplate = types.NewPaymentTemplate(
		validTemplateId1,
		validDoubledPaymentAmount,
		validPaymentMinimum,
		validPaymentMaximum,
		validDiscounts,
		validDistribution)

	validContract = types.NewPaymentContractNoDiscount(
		validPaymentContractId1, validTemplateId1,
		templateCreatorAddr, payerAddr, false, true)
)

func ValidateVariables() error {
	err := validDiscounts.Validate()
	if err != nil {
		return err
	}

	err = validDistribution.Validate()
	if err != nil {
		return err
	}

	err = validTemplate.Validate()
	if err != nil {
		return err
	}

	err = validDoublePayTemplate.Validate()
	if err != nil {
		return err
	}

	err = validContract.Validate()
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
	keyDid := sdk.NewKVStoreKey(did.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)
	ms.MountStoreWithDB(keyDid, sdk.StoreTypeIAVL, nil)

	_ = ms.LoadLatestVersion()
	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())

	cdc := codec.New()
	module.NewBasicManager(auth.AppModuleBasic{}).RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterConcrete(types.TestPeriod{}, "payments/TestPeriod", nil)

	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	pk1 := params.NewKeeper(cdc, keyParams, tkeyParams)
	paymentsSubspace := pk1.Subspace(types.DefaultParamspace)

	accountKeeper := auth.NewAccountKeeper(cdc, actStoreKey, pk1.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk1.Subspace(bank.DefaultParamspace), nil)
	didKeeper := did.NewKeeper(cdc, keyDid)
	keeper := NewKeeper(cdc, storeKey, paymentsSubspace, bankKeeper, didKeeper, nil)

	return ctx, keeper, cdc
}
