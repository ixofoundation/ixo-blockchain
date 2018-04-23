package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/tmlibs/cli"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/app"
)

// rootCmd is the entry point for this binary
var (
	context = server.NewDefaultContext()
	rootCmd = &cobra.Command{
		Use:               "ixod",
		Short:             "ixo Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(context),
	}
)

// defaultAppState sets up the app_state for the
// default genesis file
func defaultAppState(args []string, addr sdk.Address, coinDenom string) (json.RawMessage, error) {
	baseJSON, err := server.DefaultGenAppState(args, addr, coinDenom)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]json.RawMessage
	err = json.Unmarshal(baseJSON, &jsonMap)
	if err != nil {
		return nil, err
	}

	bz, err := json.Marshal(jsonMap)
	return json.RawMessage(bz), err
}

func generateApp(rootDir string, logger log.Logger) (abci.Application, error) {
	dbMain, err := dbm.NewGoLevelDB("ixo-node", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}

	dbs := map[string]dbm.DB{
		"main": dbMain,
	}
	bapp := app.NewIxoNodeApp(logger, dbs)
	return bapp, nil
}

func main() {
	server.AddCommands(rootCmd, defaultAppState, generateApp, context)

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/.ixo-node")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	executor.Execute()
}
