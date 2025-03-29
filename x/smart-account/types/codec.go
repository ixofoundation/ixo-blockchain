package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ixofoundation/ixo-blockchain/v5/x/smart-account/crypto"
)

// AuthenticatorTxOptions
type AuthenticatorTxOptions interface {
	GetSelectedAuthenticators() []uint64
}

func RegisterCodec(cdc *codec.LegacyAmino) {
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*tx.TxExtensionOptionI)(nil), &TxExtension{})

	registry.RegisterImplementations(
		(*AuthenticatorTxOptions)(nil),
		&TxExtension{},
	)

	registry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil),
		&crypto.AuthnPubKey{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
