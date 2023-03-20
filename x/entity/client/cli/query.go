package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	entityQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "entity query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	entityQueryCmd.AddCommand(
		CmdQueryParams(),
		CmdListEntity(),
		CmdShowEntity(),
		CmdShowEntityMetadata(),
		CmdShowEntityIidDocument(),
		CmdShowEntityVerified(),
	)

	return entityQueryCmd
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

func CmdListEntity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-entity",
		Short: "list all entity",
		Args:  cobra.ExactArgs(0),
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

			params := &types.QueryEntityListRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EntityList(context.Background(), params)
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

func CmdShowEntity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-entity [id]",
		Short: "Query for an entity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			argId := args[0]

			params := &types.QueryEntityRequest{
				Id: argId,
			}

			res, err := queryClient.Entity(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEntityMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-entity-metadata [id]",
		Short: "Query for an entity metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			argId := args[0]

			params := &types.QueryEntityMetadataRequest{
				Id: argId,
			}

			res, err := queryClient.EntityMetaData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEntityIidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-entity-iid-docuemnt [id]",
		Short: "Query for an entity iid document",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			argId := args[0]

			params := &types.QueryEntityIidDocumentRequest{
				Id: argId,
			}

			res, err := queryClient.EntityIidDocument(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEntityVerified() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-entity-verified [id]",
		Short: "Query for an entity verified",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			argId := args[0]

			params := &types.QueryEntityVerifiedRequest{
				Id: argId,
			}

			res, err := queryClient.EntityVerified(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
