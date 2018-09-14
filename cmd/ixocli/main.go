package main

import (
	"os"

	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/project"

	"github.com/spf13/cobra"

	"github.com/tendermint/tmlibs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/commands"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/commands"
	ibccmd "github.com/cosmos/cosmos-sdk/x/ibc/commands"
	simplestakingcmd "github.com/cosmos/cosmos-sdk/x/simplestake/commands"

	"github.com/ixofoundation/ixo-cosmos/app"
	"github.com/ixofoundation/ixo-cosmos/types"
	didcmd "github.com/ixofoundation/ixo-cosmos/x/did/cmd"
	ixolcd "github.com/ixofoundation/ixo-cosmos/x/ixo/lcd"
	projectcmd "github.com/ixofoundation/ixo-cosmos/x/project/cmd"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "ixocli",
		Short: "Ixo light-client",
	}
)

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := app.MakeCodec()

	// TODO: setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc

	// add standard rpc, and tx commands
	rpc.AddCommands(rootCmd)
	rootCmd.AddCommand(client.LineBreak)
	tx.AddCommands(rootCmd, cdc)
	rootCmd.AddCommand(client.LineBreak)

	// add query/post commands (custom to binary)
	rootCmd.AddCommand(
		client.GetCommands(
			authcmd.GetAccountCmd("main", cdc, types.GetAccountDecoder(cdc)),
		)...)
	rootCmd.AddCommand(
		client.PostCommands(
			bankcmd.SendTxCmd(cdc),
		)...)
	rootCmd.AddCommand(
		client.PostCommands(
			ibccmd.IBCTransferCmd(cdc),
		)...)
	rootCmd.AddCommand(
		client.PostCommands(
			ibccmd.IBCRelayCmd(cdc),
			simplestakingcmd.BondTxCmd(cdc),
		)...)
	rootCmd.AddCommand(
		client.PostCommands(
			simplestakingcmd.UnbondTxCmd(cdc),
		)...)

	// and now ixo specific commands
	rootCmd.AddCommand(
		client.PostCommands(
			didcmd.AddDidDocCmd(cdc),
			didcmd.GetDidDocCmd("did", cdc, did.GetDidDocDecoder(cdc)),
			didcmd.AddCredentialCmd("did", cdc, did.GetDidDocDecoder(cdc)),
			client.LineBreak,
		)...)

	rootCmd.AddCommand(
		client.PostCommands(
			projectcmd.CreateProjectCmd(cdc),
			projectcmd.UpdateProjectStatusCmd(cdc),
			projectcmd.GetProjectDocCmd("project", cdc, project.GetProjectDocDecoder(cdc)),
			projectcmd.GetProjectAccountsCmd("project"),
			projectcmd.CreateAgentCmd(cdc),
			projectcmd.UpdateAgentCmd(cdc),
			projectcmd.CreateClaimCmd(cdc),
			projectcmd.CreateEvaluationCmd(cdc),
		)...)

	rootCmd.AddCommand(
		client.PostCommands(
			projectcmd.FundProjectTxCmd(cdc),
		)...)

	// add proxy, version and key info
	rootCmd.AddCommand(
		client.LineBreak,
		ixolcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", os.ExpandEnv("$HOME/.ixo-node"))
	executor.Execute()
}
