package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ixofoundation/ixo-cosmos/x/project"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/tmlibs/cli"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/ixofoundation/ixo-cosmos/app"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
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

	dbDid, err := dbm.NewGoLevelDB("ixo-did", dataDir)
	if err != nil {
		return nil, err
	}

	dbProject, err := dbm.NewGoLevelDB("ixo-project", dataDir)
	if err != nil {
		return nil, err
	}

	dbs := map[string]dbm.DB{
		"main":    dbMain,
		"acc":     dbAcc,
		"ibc":     dbIBC,
		"staking": dbStaking,
		"did":     dbDid,
		"project": dbProject,
	}
	bapp := app.NewIxoApp(logger, dbs)
	return bapp, nil
}

// defaultAppState sets up the app_state for the
// default genesis file
func defaultAppState(args []string, addr sdk.Address, coinDenom string) (json.RawMessage, error) {
	fmt.Println("********DEBUG_MAIN.GO*********")
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
	server.AddCommands(rootCmd, defaultAppState, generateApp, context)

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/.ixo-node")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	executor.Execute()
}
