package types
//
//import (
//	"github.com/cosmos/cosmos-sdk/codec"
//)
//
//// ModuleCdc is the codec for the module
//var ModuleCdc *codec.Codec
//
//func RegisterCodec(cdc *codec.Codec) {
//	cdc.RegisterConcrete(Oracle{}, "oracles/Oracle", nil)
//	cdc.RegisterConcrete(OracleTokenCap{}, "oracles/OracleTokenCap", nil)
//}
//
//func init() {
//	ModuleCdc = codec.New()
//	RegisterCodec(ModuleCdc)
//	ModuleCdc.Seal()
//}
