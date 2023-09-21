package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/v2/x/token/types"
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
		CmdQueryParams(),
		CmdListTokens(),
		CmdShowToken(),
		CmdTokenMetadata(),
	)

	return tokenQueryCmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tokens [minter]",
		Short: "List all token docs for a minter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenListRequest{
				Pagination: pageReq,
				Minter:     args[0],
			}

			res, err := queryClient.TokenList(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdShowToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-token-doc [minter] [contract_address]",
		Short: "Query for a token doc",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenDocRequest{
				Minter:          args[0],
				ContractAddress: args[1],
			}

			res, err := queryClient.TokenDoc(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdTokenMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-token-metadata [id]",
		Short: "Query minted token metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenMetadataRequest{
				Id: args[0],
			}

			res, err := queryClient.TokenMetadata(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
