package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec registers the necessary x/iid interfaces and concrete types on the provided
// LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateIidDocument{}, "iid/CreateIidDocument", nil)
	cdc.RegisterConcrete(&MsgUpdateIidDocument{}, "iid/UpdateIidDocument", nil)
	cdc.RegisterConcrete(&MsgAddVerification{}, "iid/AddVerification", nil)
	cdc.RegisterConcrete(&MsgRevokeVerification{}, "iid/RevokeVerification", nil)
	cdc.RegisterConcrete(&MsgSetVerificationRelationships{}, "iid/SetVerificationRelationships", nil)
	cdc.RegisterConcrete(&MsgAddService{}, "iid/AddService", nil)
	cdc.RegisterConcrete(&MsgDeleteService{}, "iid/DeleteService", nil)
	cdc.RegisterConcrete(&MsgAddController{}, "iid/AddController", nil)
	cdc.RegisterConcrete(&MsgDeleteController{}, "iid/DeleteController", nil)
	cdc.RegisterConcrete(&MsgAddLinkedResource{}, "iid/AddLinkedResource", nil)
	cdc.RegisterConcrete(&MsgDeleteLinkedResource{}, "iid/DeleteLinkedResource", nil)
	cdc.RegisterConcrete(&MsgAddAccordedRight{}, "iid/AddAccordedRight", nil)
	cdc.RegisterConcrete(&MsgDeleteAccordedRight{}, "iid/DeleteAccordedRight", nil)
	cdc.RegisterConcrete(&MsgAddLinkedEntity{}, "iid/AddLinkedEntity", nil)
	cdc.RegisterConcrete(&MsgDeleteLinkedEntity{}, "iid/DeleteLinkedEntity", nil)
	cdc.RegisterConcrete(&MsgAddIidContext{}, "iid/AddIidContext", nil)
	cdc.RegisterConcrete(&MsgDeleteIidContext{}, "iid/DeleteIidContext", nil)
	cdc.RegisterConcrete(&MsgDeactivateIID{}, "iid/DeactivateIID", nil)
}

// RegisterInterfaces registers interfaces and implementations of the x/iid module.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateIidDocument{},
		&MsgUpdateIidDocument{},
		&MsgAddVerification{},
		&MsgSetVerificationRelationships{},
		&MsgRevokeVerification{},
		&MsgAddService{},
		&MsgDeleteService{},
		&MsgAddController{},
		&MsgDeleteController{},
		&MsgAddLinkedResource{},
		&MsgDeleteLinkedResource{},
		&MsgAddAccordedRight{},
		&MsgDeleteAccordedRight{},
		&MsgAddLinkedEntity{},
		&MsgDeleteLinkedEntity{},
		&MsgAddIidContext{},
		&MsgDeleteIidContext{},
		&MsgDeactivateIID{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleAminoCdc = codec.NewLegacyAmino()
	ModuleCdc      = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(ModuleAminoCdc)
	cryptocodec.RegisterCrypto(ModuleAminoCdc)
	ModuleAminoCdc.Seal()
}
