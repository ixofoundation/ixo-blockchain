package did

import (
	"github.com/cosmos/cosmos-sdk/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(AddDidMsg{}, "did/AddDid", nil)
	cdc.RegisterConcrete(AddCredentialMsg{}, "did/AddCredential", nil)
}

var msgCdc = wire.NewCodec()

func init() {
	RegisterWire(msgCdc)
}
