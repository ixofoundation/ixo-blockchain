package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&Bond{}, "cosmos-sdk/Bond", nil)
	cdc.RegisterConcrete(&FunctionParam{}, "cosmos-sdk/FunctionParam", nil)
	cdc.RegisterConcrete(&FunctionParams{}, "cosmos-sdk/FunctionParams", nil)
	cdc.RegisterConcrete(&Batch{}, "cosmos-sdk/Batch", nil)
	cdc.RegisterConcrete(&BaseOrder{}, "cosmos-sdk/BaseOrder", nil)
	cdc.RegisterConcrete(&BuyOrder{}, "cosmos-sdk/BuyOrder", nil)
	cdc.RegisterConcrete(&SellOrder{}, "cosmos-sdk/SellOrder", nil)
	cdc.RegisterConcrete(&SwapOrder{}, "cosmos-sdk/SwapOrder", nil)
	cdc.RegisterConcrete(MsgCreateBond{}, "cosmos-sdk/MsgCreateBond", nil)
	cdc.RegisterConcrete(MsgEditBond{}, "cosmos-sdk/MsgEditBond", nil)
	cdc.RegisterConcrete(MsgBuy{}, "cosmos-sdk/MsgBuy", nil)
	cdc.RegisterConcrete(MsgSell{}, "cosmos-sdk/MsgSell", nil)
	cdc.RegisterConcrete(MsgSwap{}, "cosmos-sdk/MsgSwap", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
