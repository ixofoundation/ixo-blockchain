package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// GetQueryCmd returns the cli query commands for this module.
//
// Manually registered (rather than relying on the AutoCLIOptions.Query
// surface in module/autocli.go) because cosmos-sdk's autocli renders
// gogoproto-generated nested message types as empty `{}` objects in its
// `--output json` path: it constructs a `dynamicpb.Message` from the
// proto descriptor, decodes the gRPC reply into it, and marshals via
// `aminojson.NewEncoder`. The dynamicpb / aminojson combination doesn't
// fully resolve gogo-generated nested types like Service,
// VerificationMethod, IidMetadata — they survive the wire-level decode
// but render with empty inner fields.
//
// The legacy clientCtx.PrintProto() path used here goes through gogo's
// jsonpb marshaller, which DOES handle gogo nested types correctly.
//
// This file restores correct CLI JSON rendering for x/iid queries; the
// AutoCLIOptions.Query block in autocli.go is still defined and used as
// a fallback in case a downstream tool depends on it, but cobra prefers
// the manually-registered commands when both exist.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the iid module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		newQueryIidsCmd(),
		newQueryIidCmd(),
	)

	return queryCmd
}

func newQueryIidsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iids",
		Short: "Query for all iid documents",
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
			qc := types.NewQueryClient(clientCtx)
			res, err := qc.IidDocuments(cmd.Context(), &types.QueryIidDocumentsRequest{Pagination: pageReq})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "iids")
	return cmd
}

func newQueryIidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iid [id]",
		Short: "Query for a single iid document by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			qc := types.NewQueryClient(clientCtx)
			res, err := qc.IidDocument(cmd.Context(), &types.QueryIidDocumentRequest{Id: args[0]})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
