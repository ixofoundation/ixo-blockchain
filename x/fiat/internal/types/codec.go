package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-cosmos/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*types.FiatPeg)(nil), nil)
	cdc.RegisterConcrete(&types.BaseFiatPeg{}, "ixo-cosmos/FiatPeg", nil)
	cdc.RegisterConcrete(MsgIssueFiats{}, "ixo-cosmos/MsgIssueFiats", nil)
	cdc.RegisterConcrete(MsgSendFiats{}, "ixo-cosmos/MsgSendFiats", nil)
	cdc.RegisterConcrete(MsgRedeemFiats{}, "ixo-cosmos/MsgRedeemFiats", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
