package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/types"
)

func GetOraclesRequestHandler(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-oracles",
		Short: "Query oracles",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
				keeper.QueryOracles), nil)
			if err != nil {
				return err
			}

			var oracles types.Oracles
			if err := cdc.UnmarshalJSON(bz, &oracles); err != nil {
				return err
			}

			fmt.Println(string(bz))
			return nil
		},
	}
}
