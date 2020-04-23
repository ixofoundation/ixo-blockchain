package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	ModuleCdc.Seal()
}
