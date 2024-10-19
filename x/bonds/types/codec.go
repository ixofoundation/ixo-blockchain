package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec registers the necessary x/bonds interfaces and concrete types on the provided
// LegacyAmino codec. These types are used for Amino JSON serialization.
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
	cdc.RegisterConcrete(&MsgWithdrawReserve{}, "bonds/MsgWithdrawReserve", nil)
}

// RegisterInterfaces registers interfaces and implementations of the x/bonds module.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
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
		&MsgWithdrawReserve{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
