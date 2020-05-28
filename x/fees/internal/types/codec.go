package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(Fee{}, "fees/Fee", nil)
	cdc.RegisterConcrete(FeeContract{}, "fees/FeeContract", nil)

	cdc.RegisterConcrete(Distribution{}, "fees/Distribution", nil)
	cdc.RegisterConcrete(DistributionShare{}, "fees/DistributionShare", nil)

	cdc.RegisterConcrete(Discounts{}, "fees/Discounts", nil)
	cdc.RegisterConcrete(Discount{}, "fees/Discount", nil)

	cdc.RegisterConcrete(Subscription{}, "fees/Subscription", nil)
	cdc.RegisterInterface((*SubscriptionContent)(nil), nil)
	cdc.RegisterConcrete(BlockSubscriptionContent{}, "fees/BlockSubscriptionContent", nil)
	cdc.RegisterConcrete(TimeSubscriptionContent{}, "fees/TimeSubscriptionContent", nil)

	cdc.RegisterConcrete(MsgSetFeeContractAuthorisation{}, "fees/MsgSetFeeContractAuthorisation", nil)
	cdc.RegisterConcrete(MsgCreateFee{}, "fees/MsgCreateFee", nil)
	cdc.RegisterConcrete(MsgCreateFeeContract{}, "fees/MsgCreateFeeContract", nil)
	cdc.RegisterConcrete(MsgGrantFeeDiscount{}, "fees/MsgGrantFeeDiscount", nil)
	cdc.RegisterConcrete(MsgRevokeFeeDiscount{}, "fees/MsgRevokeFeeDiscount", nil)
	cdc.RegisterConcrete(MsgChargeFee{}, "fees/MsgChargeFee", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
