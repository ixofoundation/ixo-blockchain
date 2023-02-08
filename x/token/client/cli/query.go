package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	tokenQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "token query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tokenQueryCmd.AddCommand(
		GetCmdTokenDocs(),
	)

	return tokenQueryCmd
}

func GetCmdTokenDocs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-token-doc",
		Short: "Query TokenDocs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			minterDid := args[0]

			res, err := queryClient.TokenList(context.Background(), &types.QueryTokenListRequest{MinterDid: minterDid})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
