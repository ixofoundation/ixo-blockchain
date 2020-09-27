package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func Setup(isCheckTx bool) *ixoApp {
	db := dbm.NewMemDB()
	ixoApp := NewIxoApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)
	cdc := MakeCodec()
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := NewDefaultGenesisState()
		stateBytes, err := codec.MarshalJSONIndent(cdc, genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		ixoApp.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
			},
		)
	}

	return ixoApp
}

func createTestApp(isCheckTx bool) (*ixoApp, sdk.Context) {
	app := Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})

	return app, ctx
}
