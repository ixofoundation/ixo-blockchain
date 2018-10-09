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

// defaultAppState sets up the app_state for the
// default genesis file
func defaultAppState(args []string, addr sdk.Address, coinDenom string) (json.RawMessage, error) {
	// Add ETH_PEG key for signing
	sovrinDid := sovrin.Gen()
	fmt.Println("********* Note ***********************************************************")
	fmt.Println("This is the Ethereum Peg Key and needs to be used to release Project Funds")
	fmt.Println(sovrinDid.String())
	fmt.Println("**************************************************************************")
	opts := fmt.Sprintf(`{
		"accounts": [{
			"address": "%s",
			"coins": [
				{
					"denom": "%s",
					"amount": 9007199254740992
				}
			],
			"name":"coinbase"
		}],
		"project": {
			"pegPubKey":"%s"
		}
	}`, addr.String(), project.COIN_DENOM, sovrinDid.VerifyKey)
	return json.RawMessage(opts), nil
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
