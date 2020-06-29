package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GetCmdAddressFromDid() *cobra.Command {
	return &cobra.Command{
		Use:   "get-address-from-did [did]",
		Short: "Query for an account address by DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !types.IsValidDid(args[0]) {
				return errors.New("input is not a valid did")
			}

			accAddress := types.DidToAddr(args[0])
			fmt.Println(accAddress.String())
			return nil
		},
	}
}

func GetCmdDidDoc(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-did-doc [did]",
		Short: "Query DidDoc for a DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			didAddr := args[0]
			key := exported.Did(didAddr)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryDidDoc, key), nil)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return errors.New("response bytes are empty")
			}

			var didDoc types.BaseDidDoc
			err = cdc.UnmarshalJSON(res, &didDoc)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSONIndent(didDoc, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetCmdAllDids(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-all-dids",
		Short: "Query all DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryAllDids, "ALL"), nil)
			if err != nil {
				return err
			}

			var didDids []exported.Did
			err = cdc.UnmarshalJSON(res, &didDids)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSONIndent(didDids, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetCmdAllDidDocs(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-all-did-docs",
		Short: "Query all DID documents",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryAllDidDocs, "ALL"), nil)
			if err != nil {
				return err
			}

			var didDocs []types.BaseDidDoc
			err = cdc.UnmarshalJSON(res, &didDocs)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSONIndent(didDocs, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}
