package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Period)(nil), nil)
	cdc.RegisterConcrete(BlockPeriod{}, "payments/BlockPeriod", nil)
	cdc.RegisterConcrete(TimePeriod{}, "payments/TimePeriod", nil)

	cdc.RegisterConcrete(MsgCreatePaymentTemplate{}, "payments/MsgCreatePaymentTemplate", nil)
	cdc.RegisterConcrete(MsgCreatePaymentContract{}, "payments/MsgCreatePaymentContract", nil)
	cdc.RegisterConcrete(MsgCreateSubscription{}, "payments/MsgCreateSubscription", nil)
	cdc.RegisterConcrete(MsgSetPaymentContractAuthorisation{}, "payments/MsgSetPaymentContractAuthorisation", nil)
	cdc.RegisterConcrete(MsgGrantDiscount{}, "payments/MsgGrantDiscount", nil)
	cdc.RegisterConcrete(MsgRevokeDiscount{}, "payments/MsgRevokeDiscount", nil)
	cdc.RegisterConcrete(MsgEffectPayment{}, "payments/MsgEffectPayment", nil)
}

// ModuleCdc is the codec for the module
//var ModuleCdc *codec.Codec

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

//func init() {
//	ModuleCdc = codec.New()
//	RegisterCodec(ModuleCdc)
//	ModuleCdc.Seal()
//}

func init() {
	RegisterLegacyAminoCodec(amino)
	ModuleCdc.Seal()
}
