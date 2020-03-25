package ixo

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddEthWallet{}, "ixo/AddEthWallet", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
