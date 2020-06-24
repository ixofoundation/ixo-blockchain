package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddDid{}, "did/AddDid", nil)
	cdc.RegisterConcrete(MsgAddCredential{}, "did/AddCredential", nil)

	cdc.RegisterInterface((*exported.DidDoc)(nil), nil)

	// TODO: https://github.com/ixofoundation/ixo-blockchain/issues/76
	cdc.RegisterConcrete(BaseDidDoc{}, "did/BaseDidDoc", nil)
	//cdc.RegisterConcrete(DidCredential{}, "did/DidCredential", nil)
	//cdc.RegisterConcrete(Claim{}, "did/Claim", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
