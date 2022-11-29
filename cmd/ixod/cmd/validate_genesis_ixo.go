package cmd

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"

	tmtypes "github.com/tendermint/tendermint/types"
)

const chainUpgradeGuide = "https://docs.cosmos.network/master/migrations/chain-upgrade-guide-040.html"

func ValidateGenesisCmd(mbm module.BasicManager) *cobra.Command {
	return &cobra.Command{
		Use:   "validate-genesis [file]",
		Args:  cobra.RangeArgs(0, 1),
		Short: "validates the genesis file at the default location or at the location passed as an arg",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			serverCtx := server.GetServerContextFromCmd(cmd)
			clientCtx := client.GetClientContextFromCmd(cmd)

			cdc := clientCtx.Codec

			// Load default if passed no args, otherwise load passed file
			var genesis string
			if len(args) == 0 {
				genesis = serverCtx.Config.GenesisFile()
			} else {
				genesis = args[0]
			}

			genDoc, err := validateGenDoc(genesis)
			if err != nil {
				debug.PrintStack()
				return err
			}

			var genState map[string]json.RawMessage
			if err = json.Unmarshal(genDoc.AppState, &genState); err != nil {
				debug.PrintStack()
				return fmt.Errorf("error unmarshalling genesis doc %s: %s", genesis, err.Error())
			}

			for _, b := range mbm {
				if err := b.ValidateGenesis(cdc, clientCtx.TxConfig, genState[b.Name()]); err != nil {

					fmt.Println(b.Name())
					debug.PrintStack()
					return fmt.Errorf("error validating genesis file %s: %s", genesis, err.Error())
				}
			}

			// if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, genState); err != nil {
			// 	debug.PrintStack()
			// 	return fmt.Errorf("error validating genesis file %s: %s", genesis, err.Error())
			// }

			fmt.Printf("File at %s is a valid genesis file\n", genesis)
			return nil
		},
	}
}

// validateGenDoc reads a genesis file and validates that it is a correct
// Tendermint GenesisDoc. This function does not do any cosmos-related
// validation.
func validateGenDoc(importGenesisFile string) (*tmtypes.GenesisDoc, error) {
	genDoc, err := tmtypes.GenesisDocFromFile(importGenesisFile)
	if err != nil {
		return nil, fmt.Errorf("%s. Make sure that"+
			" you have correctly migrated all Tendermint consensus params, please see the"+
			" chain migration guide at %s for more info",
			err.Error(), chainUpgradeGuide,
		)
	}

	return genDoc, nil
}
