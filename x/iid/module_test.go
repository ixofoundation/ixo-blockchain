package did_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/allinbits/cosmos-cash/v3/app"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp"

	dbm "github.com/tendermint/tm-db"
)

func TestCreateModuleInApp(t *testing.T) {
	app := app.New(
		"cosmos-cash",
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		make(map[int64]bool),
		app.DefaultNodeHome("cosmoscash"),
		0,
		app.MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
	)

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	require.NotNil(t, app.DidDocumentKeeper)
}
