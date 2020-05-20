package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(Fee{}, "fees/Fee", nil)
	cdc.RegisterConcrete(FeeContract{}, "fees/FeeContract", nil)
	cdc.RegisterConcrete(Distribution{}, "fees/Distribution", nil)
	cdc.RegisterConcrete(DistributionShare{}, "fees/DistributionShare", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
