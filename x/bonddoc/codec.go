package bonddoc

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgCreateBond{}, "bonddoc/CreateBond", nil)
	cdc.RegisterConcrete(types.MsgUpdateBondStatus{}, "bonddoc/UpdateBondStatus", nil)
}

var moduleCdc = codec.New()

func init() {
	RegisterCodec(moduleCdc)
}
