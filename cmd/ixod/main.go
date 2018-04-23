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
<<<<<<< HEAD
	"github.com/cosmos/cosmos-sdk/server"
=======
>>>>>>> 5f17e6181ef009ad7792b089ae46583eaf95894e
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

<<<<<<< HEAD
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
=======
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
>>>>>>> 5f17e6181ef009ad7792b089ae46583eaf95894e
	if err != nil {
		return nil, err
	}

	dbs := map[string]dbm.DB{
<<<<<<< HEAD
		"main":    dbMain,
		"acc":     dbAcc,
		"ibc":     dbIBC,
		"staking": dbStaking,
=======
		"main": dbMain,
>>>>>>> 5f17e6181ef009ad7792b089ae46583eaf95894e
	}
	bapp := app.NewIxoApp(logger, dbs)
	return bapp, nil
}

func main() {
	server.AddCommands(rootCmd, server.DefaultGenAppState, generateApp, context)

	// prepare and add flags
<<<<<<< HEAD
	rootDir := os.ExpandEnv("$HOME/.ixod")
=======
	rootDir := os.ExpandEnv("$HOME/.ixo-node")
>>>>>>> 5f17e6181ef009ad7792b089ae46583eaf95894e
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	executor.Execute()
}
