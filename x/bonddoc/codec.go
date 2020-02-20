package bonddoc

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
)

func Registercodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.CreateBondMsg{}, "bonddoc/CreateBond", nil)
	cdc.RegisterConcrete(types.UpdateBondStatusMsg{}, "bonddoc/UpdateBondStatus", nil)
}

var moduleCdc = codec.New()

func init() {
	Registercodec(moduleCdc)
}
