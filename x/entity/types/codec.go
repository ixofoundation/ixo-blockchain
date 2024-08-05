package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterCodec registers the necessary x/entity interfaces and concrete types on the provided
// LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateEntity{}, "entity/CreateEntity", nil)
	cdc.RegisterConcrete(&MsgUpdateEntity{}, "entity/UpdateEntity", nil)
	cdc.RegisterConcrete(&MsgUpdateEntityVerified{}, "entity/UpdateEntityVerified", nil)
	cdc.RegisterConcrete(&MsgTransferEntity{}, "entity/TransferEntity", nil)
	cdc.RegisterConcrete(&MsgCreateEntityAccount{}, "entity/CreateEntityAccount", nil)
	cdc.RegisterConcrete(&MsgGrantEntityAccountAuthz{}, "entity/GrantEntityAccountAuthz", nil)
	cdc.RegisterConcrete(&MsgRevokeEntityAccountAuthz{}, "entity/RevokeEntityAccountAuthz", nil)

	// gov proposals
	cdc.RegisterConcrete(&InitializeNftContract{}, "entity/InitializeNftContract", nil)
}

// RegisterInterfaces registers interfaces and implementations of the x/entity module.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEntity{},
		&MsgUpdateEntity{},
		&MsgUpdateEntityVerified{},
		&MsgTransferEntity{},
		&MsgCreateEntityAccount{},
		&MsgGrantEntityAccountAuthz{},
		&MsgRevokeEntityAccountAuthz{},
	)

	registry.RegisterImplementations(
		(*govtypesv1.Content)(nil),
		&InitializeNftContract{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
