package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Period)(nil), nil)
	cdc.RegisterConcrete(&BlockPeriod{}, "payments/BlockPeriod", nil)
	cdc.RegisterConcrete(&TimePeriod{}, "payments/TimePeriod", nil)

	cdc.RegisterConcrete(&MsgCreatePaymentTemplate{}, "payments/MsgCreatePaymentTemplate", nil)
	cdc.RegisterConcrete(&MsgCreatePaymentContract{}, "payments/MsgCreatePaymentContract", nil)
	cdc.RegisterConcrete(&MsgCreateSubscription{}, "payments/MsgCreateSubscription", nil)
	cdc.RegisterConcrete(&MsgSetPaymentContractAuthorisation{}, "payments/MsgSetPaymentContractAuthorisation", nil)
	cdc.RegisterConcrete(&MsgGrantDiscount{}, "payments/MsgGrantDiscount", nil)
	cdc.RegisterConcrete(&MsgRevokeDiscount{}, "payments/MsgRevokeDiscount", nil)
	cdc.RegisterConcrete(&MsgEffectPayment{}, "payments/MsgEffectPayment", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetPaymentContractAuthorisation{},
		&MsgCreatePaymentTemplate{},
		&MsgCreatePaymentContract{},
		&MsgCreateSubscription{},
		&MsgGrantDiscount{},
		&MsgRevokeDiscount{},
		&MsgEffectPayment{},
	)
	registry.RegisterInterface(
		"payments.Period",
		(*Period)(nil),
		&BlockPeriod{},
		&TimePeriod{},
		&TestPeriod{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/payments module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/gov and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
