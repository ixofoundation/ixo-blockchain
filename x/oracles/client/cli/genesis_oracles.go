package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

const (
	flagClientHome = "home-client"
)

func AddGenesisOracleCmd(ctx *server.Context, cdc *codec.Codec,
	defaultNodeHome, defaultClientHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-oracle [oracle-did] [capability][,[capability]]",
		Short: "Add oracle to genesis.json",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			oracleDid := did.Did(args[0])
			if len(oracleDid) == 0 {
				return fmt.Errorf("oracle did cannot be empty")
			}

			// Check that oracle token capabilities are valid
			capabilities, err := types.ParseOracleTokenCaps(args[1])
			if err != nil {
				return err
			}

			// Check that oracle DID is valid
			if !did.IsValidDid(oracleDid) {
				return fmt.Errorf("oracle DID is invalid")
			}

			oracle := types.NewOracle(oracleDid, capabilities)

			// retrieve the app state
			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
			if err != nil {
				return err
			}

			// add genesis oracle to the app state
			var genesisState types.GenesisState

			cdc.MustUnmarshalJSON(appState[types.ModuleName], &genesisState)

			if genesisState.Oracles.Includes(oracle) {
				return fmt.Errorf("cannot add oracle since it already exists")
			}

			genesisState.Oracles = append(genesisState.Oracles, oracle)

			genesisStateBz := cdc.MustMarshalJSON(genesisState)
			appState[types.ModuleName] = genesisStateBz

			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return err
			}

			// export app state
			genDoc.AppState = appStateJSON

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")
	return cmd
}
