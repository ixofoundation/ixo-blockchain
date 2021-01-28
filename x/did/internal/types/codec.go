package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

/*func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddDid{}, "did/AddDid", nil)
	cdc.RegisterConcrete(MsgAddCredential{}, "did/AddCredential", nil)

	cdc.RegisterInterface((*exported.DidDoc)(nil), nil)

	// TODO: https://github.com/ixofoundation/ixo-blockchain/issues/76
	cdc.RegisterConcrete(BaseDidDoc{}, "did/BaseDidDoc", nil)
	//cdc.RegisterConcrete(DidCredential{}, "did/DidCredential", nil)
	//cdc.RegisterConcrete(Claim{}, "did/Claim", nil)
}*/

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgAddDid{}, "did/AddDid", nil)
	cdc.RegisterConcrete(MsgAddCredential{}, "did/AddCredential", nil)

	cdc.RegisterInterface((*exported.DidDoc)(nil), nil)

	// TODO: https://github.com/ixofoundation/ixo-blockchain/issues/76
	cdc.RegisterConcrete(BaseDidDoc{}, "did/BaseDidDoc", nil)
	//cdc.RegisterConcrete(DidCredential{}, "did/DidCredential", nil)
	//cdc.RegisterConcrete(Claim{}, "did/Claim", nil)
}

//Registers did module's interface types and their concrete implementations as proto.Message.
// TODO Add more types if necessary (look at cosmos x/auth/types/codec.go)
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"did.DidDoc",
		(*exported.DidDoc)(nil),
		&BaseDidDoc{},
	)
}

// ModuleCdc is the codec for the module
//var ModuleCdc *codec.Codec

var (
	amino     = codec.NewLegacyAmino()
	//ModuleCdc = codec.NewAminoCodec(amino)
)

//func init() {
//	ModuleCdc = codec.New()
//	RegisterCodec(ModuleCdc)
//	ModuleCdc.Seal()
//}
