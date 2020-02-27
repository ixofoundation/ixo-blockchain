package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(CreateBondMsg{}, "ixo-cosmos/MsgCreateBond", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	ModuleCdc.Seal()
}
