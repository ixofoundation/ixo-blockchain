package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&Bond{}, "bonds/Bond", nil)
	cdc.RegisterConcrete(&FunctionParam{}, "bonds/FunctionParam", nil)
	cdc.RegisterConcrete(&Batch{}, "bonds/Batch", nil)
	cdc.RegisterConcrete(&BaseOrder{}, "bonds/BaseOrder", nil)
	cdc.RegisterConcrete(&BuyOrder{}, "bonds/BuyOrder", nil)
	cdc.RegisterConcrete(&SellOrder{}, "bonds/SellOrder", nil)
	cdc.RegisterConcrete(&SwapOrder{}, "bonds/SwapOrder", nil)
	cdc.RegisterConcrete(&MsgCreateBond{}, "bonds/MsgCreateBond", nil)
	cdc.RegisterConcrete(&MsgEditBond{}, "bonds/MsgEditBond", nil)
	cdc.RegisterConcrete(&MsgSetNextAlpha{}, "bonds/MsgSetNextAlpha", nil)
	cdc.RegisterConcrete(&MsgUpdateBondState{}, "bonds/MsgUpdateBondState", nil)
	cdc.RegisterConcrete(&MsgBuy{}, "bonds/MsgBuy", nil)
	cdc.RegisterConcrete(&MsgSell{}, "bonds/MsgSell", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "bonds/MsgSwap", nil)
	cdc.RegisterConcrete(&MsgMakeOutcomePayment{}, "bonds/MsgMakeOutcomePayment", nil)
	cdc.RegisterConcrete(&MsgWithdrawShare{}, "bonds/MsgWithdrawShare", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBond{},
		&MsgEditBond{},
		&MsgSetNextAlpha{},
		&MsgUpdateBondState{},
		&MsgBuy{},
		&MsgSell{},
		&MsgSwap{},
		&MsgMakeOutcomePayment{},
		&MsgWithdrawShare{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var(
	amino     = codec.NewLegacyAmino()

	// ModuleCdc references the global x/did module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/did and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
