package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&Bond{}, "bonds/Bond", nil)
	cdc.RegisterConcrete(&FunctionParam{}, "bonds/FunctionParam", nil)
	cdc.RegisterConcrete(&Batch{}, "bonds/Batch", nil)
	cdc.RegisterConcrete(&BaseOrder{}, "bonds/BaseOrder", nil)
	cdc.RegisterConcrete(&BuyOrder{}, "bonds/BuyOrder", nil)
	cdc.RegisterConcrete(&SellOrder{}, "bonds/SellOrder", nil)
	cdc.RegisterConcrete(&SwapOrder{}, "bonds/SwapOrder", nil)
	cdc.RegisterConcrete(MsgCreateBond{}, "bonds/MsgCreateBond", nil)
	cdc.RegisterConcrete(MsgEditBond{}, "bonds/MsgEditBond", nil)
	cdc.RegisterConcrete(MsgBuy{}, "bonds/MsgBuy", nil)
	cdc.RegisterConcrete(MsgSell{}, "bonds/MsgSell", nil)
	cdc.RegisterConcrete(MsgSwap{}, "bonds/MsgSwap", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
