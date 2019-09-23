package cli

import (
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	
	"github.com/ixofoundation/ixo-cosmos/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/did/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func GetDidDocCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getDidDoc did",
		Short: "Get a new DidDoc for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide a did")
			}
			
			didAddr := args[0]
			key := ixo.Did(didAddr)
			
			ctx := context.NewCLIContext().
				WithCodec(cdc)
			
			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
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

func GetAllDidsCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getAllDids",
		Short: "query all dids",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryAllDids, "ALL"), nil)
			if err != nil {
				return err
			}
			
			didDids := []ixo.Did{}
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

func GetAllDidDocsCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getAllDidDocs",
		Short: "get all did docs",
		RunE: func(cmd *cobra.Command, args []string) error {
			
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
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
