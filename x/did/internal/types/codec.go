package types

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(AddDidMsg{}, "did/AddDid", nil)
	cdc.RegisterConcrete(AddCredentialMsg{}, "did/AddCredential", nil)
	cdc.RegisterInterface((*ixo.DidDoc)(nil), nil)

}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
