package fee

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
)

// genesis info must specify starting difficulty and starting count
type Genesis struct {
	Validators []sdk.Validator `json:"validators"`
	//	Supermaj sdk.Rat
	//	Timeout  int64
}

// DidKeeper manages dids
type Keeper struct {
	key sdk.StoreKey
	cdc *wire.Codec
	//	OracleKeeper oracle.Keeper
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey) Keeper {
	//	return Keeper{key, cdc, handler}
	return Keeper{key, cdc}
}

// InitGenesis for the POW module
func InitGenesis(ctx sdk.Context, keeper Keeper, data Genesis) error {
	for i, validator := range data.Validators {
		fmt.Println("Validator:", i, validator)
	}
	//	k.OracleKeeper = oracle.NewKeeper(k.key, k.cdc, genesis.Valset, genesis.Supermaj, genesis.Timeout)
	return nil
}
