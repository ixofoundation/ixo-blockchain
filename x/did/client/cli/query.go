package cli

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"strings"

	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GetCmdAddressFromBase58Pubkey() *cobra.Command {
	cmd :=  &cobra.Command{
		Use:   "get-address-from-pubkey [base-58-encoded-pubkey]",
		Short: "Get the address for a base-58 encoded ed25519 public key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !types.IsValidPubKey(args[0]) {
				return errors.New("input is not a valid base-58 encoded pubKey")
			}

			accAddress := exported.VerifyKeyToAddr(args[0])
			fmt.Println(accAddress.String())
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Referred to cosmos-sdk x/upgrade/client/cli/query.go for refactoring GetCmds which took amino as argument
func GetCmdAddressFromDid(/*cdc *codec.Codec*/) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-address-from-did [did]",
		Short: "Query address for a DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(legacyQuerierCdc)

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags()) //client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			didAddr := args[0]
			key := exported.Did(didAddr)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryDidDoc, key), nil)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return errors.New("response bytes are empty")
			}

			var didDoc types.BaseDidDoc
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &didDoc) //cdc.UnmarshalJSON(res, &didDoc)
			if err != nil {
				return err
			}
			addressFromDid := didDoc.Address()

			fmt.Println(addressFromDid.String())
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdIxoDidFromMnemonic() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-ixo-did-from-mnemonic [mnemonic]",
		Short: "Get an ixo DID from a 12-word secret mnemonic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(strings.Split(args[0], " ")) != 12 {
				return errors.New("input is not a 12-word mnemonic")
			}

			ixoDid, err := exported.FromMnemonic(args[0])
			if err != nil {
				return err
			}

			output, err := json.Marshal(ixoDid)
			if err != nil {
				panic(err)
			}

			fmt.Println(fmt.Sprintf("%v", string(output)))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdDidDoc(/*cdc *codec.Codec*/) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-did-doc [did]",
		Short: "Query DidDoc for a DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags()) //client. ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			didAddr := args[0]
			key := exported.Did(didAddr)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryDidDoc, key), nil)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return errors.New("response bytes are empty")
			}

			var didDoc types.BaseDidDoc
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &didDoc) //cdc.UnmarshalJSON(res, &didDoc)
			if err != nil {
				return err
			}

			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(didDoc, "", "  ") //cdc.MarshalJSONIndent(didDoc, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAllDids(/*cdc *codec.Codec*/) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-dids",
		Short: "Query all DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags()) //client. ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			res, _, err := /*cliCtx*/ clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryAllDids, "ALL"), nil)
			if err != nil {
				return err
			}

			var didDids []exported.Did
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &didDids) //cdc.UnmarshalJSON(res, &didDids)
			if err != nil {
				return err
			}

			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(didDids, "", "  ") //cdc.MarshalJSONIndent(didDids, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAllDidDocs(/*cdc *codec.Codec*/) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-did-docs",
		Short: "Query all DID documents",
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags()) //client. ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			res, _, err := /*cliCtx*/ clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryAllDidDocs, "ALL"), nil)
			if err != nil {
				return err
			}

			var didDocs []types.BaseDidDoc
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &didDocs) //cdc.UnmarshalJSON(res, &didDocs)
			if err != nil {
				return err
			}

			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(didDocs, "", "  ") //cdc.MarshalJSONIndent(didDocs, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
