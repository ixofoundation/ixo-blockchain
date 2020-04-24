package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "treasury/MsgSend", nil)
	cdc.RegisterConcrete(MsgMint{}, "treasury/MsgMint", nil)
	cdc.RegisterConcrete(MsgBurn{}, "treasury/MsgBurn", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
