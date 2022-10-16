package keeper_test

/*
import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func CreateTestInput() (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
	if err := ValidateVariables(); err != nil {
		panic(err)
	}

	appl := cmd.Setup(false)
	ctx := appl.BaseApp.NewContext(false, tmproto.Header{})

	return appl.LegacyAmino(), appl, ctx
}
*/
