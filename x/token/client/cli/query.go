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
		CmdListTokens(),
		CmdShowToken(),
	)

	return tokenQueryCmd
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
