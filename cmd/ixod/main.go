package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ixofoundation/ixo-blockchain/v5/app"
	"github.com/ixofoundation/ixo-blockchain/v5/app/params"
	"github.com/ixofoundation/ixo-blockchain/v5/cmd/ixod/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "IXO", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
