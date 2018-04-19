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

	"github.com/ixo.foundation/ixo-cosmos/app"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	jsonMap["cool"] = json.RawMessage(`{
        "trend": "ice-cold"
      }`)
	bz, err := json.Marshal(jsonMap)
	return json.RawMessage(bz), err
}

func generateApp(rootDir string, logger log.Logger) (abci.Application, error) {
	dbMain, err := dbm.NewGoLevelDB("ixo", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbAcc, err := dbm.NewGoLevelDB("ixo-acc", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbPow, err := dbm.NewGoLevelDB("ixo-pow", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbIBC, err := dbm.NewGoLevelDB("ixo-ibc", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbStaking, err := dbm.NewGoLevelDB("ixo-staking", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbs := map[string]dbm.DB{
		"main":    dbMain,
		"acc":     dbAcc,
		"pow":     dbPow,
		"ibc":     dbIBC,
		"staking": dbStaking,
	}
	bapp := app.NewIxoApp(logger, dbs)
	return bapp, nil
}

func main() {
	server.AddCommands(rootCmd, defaultAppState, generateApp, context)

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/.ixod")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	executor.Execute()
}
