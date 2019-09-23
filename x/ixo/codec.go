package ixo

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func Registercodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(AddEthWalletMsg{}, "ixo/AddEthWallet", nil)
}

var msgCdc = codec.New()

func init() {
	Registercodec(msgCdc)
}
