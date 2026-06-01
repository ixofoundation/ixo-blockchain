package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers names module types for legacy Amino JSON
// serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateNamespace{}, "names/CreateNamespace", nil)
	cdc.RegisterConcrete(&MsgUpdateNamespace{}, "names/UpdateNamespace", nil)
	cdc.RegisterConcrete(&MsgRegisterName{}, "names/RegisterName", nil)
	cdc.RegisterConcrete(&MsgRegisterNameByRegistrar{}, "names/RegisterNameByRegistrar", nil)
	cdc.RegisterConcrete(&MsgUpdateNameByRegistrar{}, "names/UpdateNameByRegistrar", nil)
	cdc.RegisterConcrete(&MsgTransferName{}, "names/TransferName", nil)
	cdc.RegisterConcrete(&MsgSetNameStatus{}, "names/SetNameStatus", nil)
}

// RegisterInterfaces registers the module's sdk.Msg implementations.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateNamespace{},
		&MsgUpdateNamespace{},
		&MsgRegisterName{},
		&MsgRegisterNameByRegistrar{},
		&MsgUpdateNameByRegistrar{},
		&MsgTransferName{},
		&MsgSetNameStatus{},
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
