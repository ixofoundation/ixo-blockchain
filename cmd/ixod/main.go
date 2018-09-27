package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/app"
)

// rootCmd is the entry point for this binary
var (
	context = server.NewDefaultContext()
	rootCmd = &cobra.Command{
		Use:               "ixod",
		Short:             "Ixo Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(context),
	}
)

// init parameters
var IxoAppInit = server.AppInit{
	AppGenState: IxoAppGenState,
	AppGenTx:    server.SimpleAppGenTx,
}

// ixoGenAppParams sets up the app_state and appends the ixo app state
func IxoAppGenState(cdc *wire.Codec, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {
	appState, err = server.SimpleAppGenState(cdc, appGenTxs)
	if err != nil {
		return
	}

	// key := "cool"
	// value := json.RawMessage(`{
	//       "trend": "ice-cold"
	//     }`)

	// appState, err = server.InsertKeyJSON(cdc, appState, key, value)
	// if err != nil {
	// 	return
	// }

	return
}

func newApp(logger log.Logger, db dbm.DB, _ io.Writer) abci.Application {
	return app.NewIxoApp(logger, db)
}

func exportAppStateAndTMValidators(logger log.Logger, db dbm.DB, _ io.Writer) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	dapp := app.NewIxoApp(logger, db)
	return dapp.ExportAppStateAndValidators()
}

func main() {
	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "ixod",
		Short:             "ixo Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	server.AddCommands(ctx, cdc, rootCmd, IxoAppInit,
		server.ConstructAppCreator(newApp, "ixo"),
		server.ConstructAppExporter(exportAppStateAndTMValidators, "ixo"))

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/.ixo-node")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}
