package main

import (
	"os"
	"github.com/cosmos/cosmos-sdk/server"
)

// ixod custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	rootCmd, _ := NewRootCmd()
	if err := Execute(rootCmd); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}

//func main() {
//	cdc := app.MakeCodec()
//
//	config := sdk.GetConfig()
//	config.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
//	config.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
//	config.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
//	config.Seal()
//
//	ctx := server.NewDefaultContext()
//	cobra.EnableCommandSorting = false
//	rootCmd := &cobra.Command{
//		Use:               "ixod",
//		Short:             "ixo Daemon (server)",
//		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
//	}
//
//	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
//	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app.DefaultNodeHome))
//	rootCmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx, cdc))
//	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
//		auth.GenesisAccountIterator{}, app.DefaultNodeHome, app.DefaultCLIHome))
//	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
//	rootCmd.AddCommand(AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
//	rootCmd.AddCommand(oraclesCli.AddGenesisOracleCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
//	rootCmd.AddCommand(flags.NewCompletionCmd(rootCmd, true))
//	rootCmd.AddCommand(replayCmd())
//
//	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)
//
//	// prepare and add flags
//	executor := cli.PrepareBaseCmd(rootCmd, "IXO", app.DefaultNodeHome)
//	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
//		0, "Assert registered invariants every N blocks")
//	err := executor.Execute()
//	if err != nil {
//		panic(err)
//	}
//}

