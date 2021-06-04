package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/ixofoundation/ixo-blockchain/cmd"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"github.com/tendermint/tendermint/crypto"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	shareAddr1          = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr1")))
	shareAddr2          = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr2")))
	templateCreatorAddr = sdk.AccAddress(crypto.AddressHash([]byte("templateCreatorAddr")))
	payerAddr           = sdk.AccAddress(crypto.AddressHash([]byte("payerAddr")))

	validPaymentAmount, _  = sdk.ParseCoinsNormalized("1uixo,2res")
	validPaymentMinimum, _ = sdk.ParseCoinsNormalized("3res")
	validPaymentMaximum    = sdk.NewCoins()

	validTemplateId1        = types.PaymentTemplateIdPrefix + "pt1"
	validTemplateId2        = types.PaymentTemplateIdPrefix + "pt2"
	validPaymentContractId1 = types.PaymentContractIdPrefix + "pc1"
	validPaymentContractId2 = types.PaymentContractIdPrefix + "pc2"
	validSubscriptionId1    = types.SubscriptionIdPrefix + "s1"
	validSubscriptionId2    = types.SubscriptionIdPrefix + "s2"

	validDoubledPaymentAmount, _ = sdk.ParseCoinsNormalized("2uixo,4res")

	validDiscounts = types.NewDiscounts(
		types.NewDiscount(sdk.NewUint(1), sdk.MustNewDecFromStr("10")),
		types.NewDiscount(sdk.NewUint(2), sdk.MustNewDecFromStr("50")))
	tenPercentOffId   = validDiscounts[0].Id
	fiftyPercentOffId = validDiscounts[1].Id

	validRecipients = types.NewDistribution(
		types.NewDistributionShare(shareAddr1, sdk.NewDec(50)),
		types.NewDistributionShare(shareAddr2, sdk.NewDec(50)))

	validTemplate = types.NewPaymentTemplate(
		validTemplateId1,
		validPaymentAmount,
		validPaymentMinimum,
		validPaymentMaximum,
		validDiscounts)

	validDoublePayTemplate = types.NewPaymentTemplate(
		validTemplateId1,
		validDoubledPaymentAmount,
		validPaymentMinimum,
		validPaymentMaximum,
		validDiscounts)

	validContract = types.NewPaymentContractNoDiscount(
		validPaymentContractId1, validTemplateId1, templateCreatorAddr,
		payerAddr, validRecipients, false, true)
)

func ValidateVariables() error {
	err := validDiscounts.Validate()
	if err != nil {
		return err
	}

	err = validRecipients.Validate()
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

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*app.IxoApp, sdk.Context) {
	appl := cmd.Setup(isCheckTx)
	ctx := appl.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	appl.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return appl, ctx
}

func CreateTestInput() (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
	if err := ValidateVariables(); err != nil {
		panic(err)
	}

	appl := cmd.Setup(false)
	ctx := appl.BaseApp.NewContext(false, tmproto.Header{})

	appl.PaymentsKeeper = NewKeeper(
		appl.AppCodec(),
		appl.GetKey(types.StoreKey),
		appl.BankKeeper,
		appl.DidKeeper,
		[]string{types.ModuleName},
	)

	return appl.LegacyAmino(), appl, ctx

	//storeKey := sdk.NewKVStoreKey(types.StoreKey)
	//actStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	//keyDid := sdk.NewKVStoreKey(didtypes.StoreKey)
	//
	//db := dbm.NewMemDB()
	//ms := store.NewCommitMultiStore(db)
	//ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)
	//ms.MountStoreWithDB(keyDid, sdk.StoreTypeIAVL, nil)
	//
	//_ = ms.LoadLatestVersion()
	//ctx := sdk.NewContext(ms, /*abci.Header{}*/tmproto.Header{}, true, log.NewNopLogger())
	//
	//
	//
	//cdc := codec.NewLegacyAmino()
	//module.NewBasicManager(auth.AppModuleBasic{}).RegisterLegacyAminoCodec(cdc)
	//sdk.RegisterLegacyAminoCodec(cdc)
	//cryptocodec.RegisterCrypto(cdc)
	//types.RegisterLegacyAminoCodec(cdc)
	//cdc.RegisterConcrete(types.TestPeriod{}, "payments/TestPeriod", nil)
	//
	//keyParams := sdk.NewKVStoreKey("subspace")
	//tkeyParams := sdk.NewTransientStoreKey("transient_params")
	//
	//pk1 := paramskeeper.NewKeeper(cdc, keyParams, tkeyParams)
	//
	//accountKeeper := authkeeper.NewAccountKeeper(cdc, actStoreKey, pk1.Subspace(/*auth.DefaultParamspace*/authtypes.ModuleName), authtypes.ProtoBaseAccount)
	//bankKeeper := bankkeeper.NewBaseKeeper(accountKeeper, pk1.Subspace(/*bank.DefaultParamspace*/banktypes.ModuleName), nil)
	//didKeeper := did.NewKeeper(*cdc, keyDid)
	//keeper := NewKeeper(cdc, storeKey, bankKeeper, didKeeper, nil)

	//return ctx, keeper, cdc
}
