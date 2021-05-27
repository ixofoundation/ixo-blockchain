package keeper
// TODO uncomment and fix CreateTestInput()
//
//import (
//	"github.com/cosmos/cosmos-sdk/codec"
//	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
//	"github.com/cosmos/cosmos-sdk/store"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/module"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
//	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
//	//"github.com/cosmos/cosmos-sdk/x/bank"
//	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
//	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
//	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
//	"github.com/ixofoundation/ixo-blockchain/x/did"
//	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
//	//abci "github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/crypto"
//	"github.com/tendermint/tendermint/libs/log"
//	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
//	dbm "github.com/tendermint/tm-db"
//)
//
//var (
//	shareAddr1          = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr1")))
//	shareAddr2          = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr2")))
//	templateCreatorAddr = sdk.AccAddress(crypto.AddressHash([]byte("templateCreatorAddr")))
//	payerAddr           = sdk.AccAddress(crypto.AddressHash([]byte("payerAddr")))
//
//	validPaymentAmount, _  = sdk.ParseCoins("1uixo,2res")
//	validPaymentMinimum, _ = sdk.ParseCoins("3res")
//	validPaymentMaximum    = sdk.NewCoins()
//
//	validTemplateId1        = types.PaymentTemplateIdPrefix + "pt1"
//	validTemplateId2        = types.PaymentTemplateIdPrefix + "pt2"
//	validPaymentContractId1 = types.PaymentContractIdPrefix + "pc1"
//	validPaymentContractId2 = types.PaymentContractIdPrefix + "pc2"
//	validSubscriptionId1    = types.SubscriptionIdPrefix + "s1"
//	validSubscriptionId2    = types.SubscriptionIdPrefix + "s2"
//
//	validDoubledPaymentAmount, _ = sdk.ParseCoins("2uixo,4res")
//
//	validDiscounts = types.NewDiscounts(
//		types.NewDiscount(sdk.NewUint(1), sdk.MustNewDecFromStr("10")),
//		types.NewDiscount(sdk.NewUint(2), sdk.MustNewDecFromStr("50")))
//	tenPercentOffId   = validDiscounts[0].Id
//	fiftyPercentOffId = validDiscounts[1].Id
//
//	validRecipients = types.NewDistribution(
//		types.NewDistributionShare(shareAddr1, sdk.NewDec(50)),
//		types.NewDistributionShare(shareAddr2, sdk.NewDec(50)))
//
//	validTemplate = types.NewPaymentTemplate(
//		validTemplateId1,
//		validPaymentAmount,
//		validPaymentMinimum,
//		validPaymentMaximum,
//		validDiscounts)
//
//	validDoublePayTemplate = types.NewPaymentTemplate(
//		validTemplateId1,
//		validDoubledPaymentAmount,
//		validPaymentMinimum,
//		validPaymentMaximum,
//		validDiscounts)
//
//	validContract = types.NewPaymentContractNoDiscount(
//		validPaymentContractId1, validTemplateId1, templateCreatorAddr,
//		payerAddr, validRecipients, false, true)
//)
//
//func ValidateVariables() error {
//	err := validDiscounts.Validate()
//	if err != nil {
//		return err
//	}
//
//	err = validRecipients.Validate()
//	if err != nil {
//		return err
//	}
//
//	err = validTemplate.Validate()
//	if err != nil {
//		return err
//	}
//
//	err = validDoublePayTemplate.Validate()
//	if err != nil {
//		return err
//	}
//
//	err = validContract.Validate()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func CreateTestInput() (sdk.Context, Keeper, *codec.LegacyAmino) {
//	if err := ValidateVariables(); err != nil {
//		panic(err)
//	}
//
//	storeKey := sdk.NewKVStoreKey(types.StoreKey)
//	actStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
//	keyDid := sdk.NewKVStoreKey(did.StoreKey)
//
//	db := dbm.NewMemDB()
//	ms := store.NewCommitMultiStore(db)
//	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(actStoreKey, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(keyDid, sdk.StoreTypeIAVL, nil)
//
//	_ = ms.LoadLatestVersion()
//	ctx := sdk.NewContext(ms, /*abci.Header{}*/tmproto.Header{}, true, log.NewNopLogger())
//
//	cdc := codec.NewLegacyAmino()
//	module.NewBasicManager(auth.AppModuleBasic{}).RegisterLegacyAminoCodec(cdc)
//	sdk.RegisterLegacyAminoCodec(cdc)
//	cryptocodec.RegisterCrypto(cdc)
//	types.RegisterLegacyAminoCodec(cdc)
//	cdc.RegisterConcrete(types.TestPeriod{}, "payments/TestPeriod", nil)
//
//	keyParams := sdk.NewKVStoreKey("subspace")
//	tkeyParams := sdk.NewTransientStoreKey("transient_params")
//
//	pk1 := paramskeeper.NewKeeper(cdc, keyParams, tkeyParams)
//
//	accountKeeper := authkeeper.NewAccountKeeper(cdc, actStoreKey, pk1.Subspace(/*auth.DefaultParamspace*/authtypes.ModuleName), authtypes.ProtoBaseAccount)
//	bankKeeper := bankkeeper.NewBaseKeeper(accountKeeper, pk1.Subspace(/*bank.DefaultParamspace*/banktypes.ModuleName), nil)
//	didKeeper := did.NewKeeper(*cdc, keyDid)
//	keeper := NewKeeper(cdc, storeKey, bankKeeper, didKeeper, nil)
//
//	return ctx, keeper, cdc
//}
