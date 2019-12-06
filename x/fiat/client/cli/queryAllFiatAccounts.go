package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/spf13/cobra"

	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

func GetAllFiatAccountsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allFiatAccounts",
		Short: "Query all fiat accounts",
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext()

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/", fiatTypes.QuerierRoute, "queryAllFiatAccounts"), nil)
			if err != nil {
				return err
			}

			fmt.Println(string(res))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(fsAddress)
	return cmd
}
