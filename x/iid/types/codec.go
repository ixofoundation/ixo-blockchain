package types

import ( // this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

//nolint
func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
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
		//&MsgAddDidContext{},
		//&MsgDeleteDidContext{},
		//&MsgAddLinkedResource{},
		//&MsgDeleteLinkedResource{},
		//&MsgAddAccordedRight{},
		//&MsgDeleteAccordedRight{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// ModuleCdc codec used by the module (protobuf)
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
