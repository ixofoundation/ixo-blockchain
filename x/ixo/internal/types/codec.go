package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	crypto "github.com/cosmos/cosmos-sdk/crypto/codec"
)

// module wide codec
var ModuleCdc *codec.LegacyAmino

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*IxoMsg)(nil), nil)
}

func init() {
	ModuleCdc = codec.NewLegacyAmino()
	RegisterCodec(ModuleCdc)
	crypto.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
