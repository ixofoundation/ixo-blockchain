package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.DidDoc)(nil), nil)

	cdc.RegisterConcrete(&MsgAddDid{}, "did/AddDid", nil)
	cdc.RegisterConcrete(&MsgAddCredential{}, "did/AddCredential", nil)

	// TODO: https://github.com/ixofoundation/ixo-blockchain/issues/76
	cdc.RegisterConcrete(&BaseDidDoc{}, "did/BaseDidDoc", nil)
	//cdc.RegisterConcrete(DidCredential{}, "did/DidCredential", nil)
	//cdc.RegisterConcrete(Claim{}, "did/Claim", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"did.DidDoc",
		(*exported.DidDoc)(nil),
		&BaseDidDoc{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddDid{},
		&MsgAddCredential{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()

	// ModuleCdc references the global x/did module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/did and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
