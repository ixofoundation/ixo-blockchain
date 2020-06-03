package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Period)(nil), nil)
	cdc.RegisterConcrete(BlockPeriod{}, "fees/BlockPeriod", nil)
	cdc.RegisterConcrete(TimePeriod{}, "fees/TimePeriod", nil)

	cdc.RegisterConcrete(MsgSetFeeContractAuthorisation{}, "fees/MsgSetFeeContractAuthorisation", nil)
	cdc.RegisterConcrete(MsgCreateFee{}, "fees/MsgCreateFee", nil)
	cdc.RegisterConcrete(MsgCreateFeeContract{}, "fees/MsgCreateFeeContract", nil)
	cdc.RegisterConcrete(MsgCreateSubscription{}, "fees/MsgCreateSubscription", nil)
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
