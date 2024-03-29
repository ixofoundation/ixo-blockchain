package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateEntity{}, "entity/CreateEntity", nil)
	cdc.RegisterConcrete(&MsgUpdateEntity{}, "entity/UpdateEntity", nil)
	cdc.RegisterConcrete(&MsgUpdateEntityVerified{}, "entity/UpdateEntityVerified", nil)
	cdc.RegisterConcrete(&MsgTransferEntity{}, "entity/TransferEntity", nil)
	cdc.RegisterConcrete(&MsgCreateEntityAccount{}, "entity/CreateEntityAccount", nil)
	cdc.RegisterConcrete(&MsgGrantEntityAccountAuthz{}, "entity/GrantEntityAccountAuthz", nil)
	cdc.RegisterConcrete(&MsgRevokeEntityAccountAuthz{}, "entity/RevokeEntityAccountAuthz", nil)
}

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
		(*govtypes.Content)(nil),
		&InitializeNftContract{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/gov module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/gov and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
