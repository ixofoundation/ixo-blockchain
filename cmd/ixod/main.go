package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/tmlibs/cli"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/ixofoundation/ixo-cosmos/app"
	"github.com/cosmos/cosmos-sdk/server"
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

func generateApp(rootDir string, logger log.Logger) (abci.Application, error) {
	dataDir := filepath.Join(rootDir, "data")
	dbMain, err := dbm.NewGoLevelDB("ixo", dataDir)
	if err != nil {
		return nil, err
	}
	dbAcc, err := dbm.NewGoLevelDB("ixo-acc", dataDir)
	if err != nil {
		return nil, err
	}
	dbIBC, err := dbm.NewGoLevelDB("ixo-ibc", dataDir)
	if err != nil {
		return nil, err
	}
	dbStaking, err := dbm.NewGoLevelDB("ixo-staking", dataDir)
	if err != nil {
		return nil, err
	}
	dbs := map[string]dbm.DB{
		"main":    dbMain,
		"acc":     dbAcc,
		"ibc":     dbIBC,
		"staking": dbStaking,
	}
	bapp := app.NewIxoApp(logger, dbs)
	return bapp, nil
}

func main() {
	server.AddCommands(rootCmd, server.DefaultGenAppState, generateApp, context)

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/.ixod")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	executor.Execute()
}
