package cli

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	didQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "did query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	didQueryCmd.AddCommand(
		GetCmdAddressFromBase58Pubkey(),
		GetCmdAddressFromDid(),
		GetCmdIxoDidFromMnemonic(),
		GetCmdDidDoc(),
		GetCmdAllDids(),
		GetCmdAllDidDocs(),
	)

	return didQueryCmd
}

func GetCmdAddressFromBase58Pubkey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-address-from-pubkey [base-58-encoded-pubkey]",
		Short: "Get the address for a base-58 encoded ed25519 public key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			pubKey := args[0]

			if !types.IsValidPubKey(pubKey) {
				return errors.New("input is not a valid base-58 encoded pubKey")
			}

			accAddress := exported.VerifyKeyToAddr(pubKey)

			return clientCtx.PrintString(accAddress.String())
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAddressFromDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-address-from-did [did]",
		Short: "Query address for a DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			didAddr := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AddressFromDid(context.Background(),
				&types.QueryAddressFromDidRequest{Did: didAddr})
			if err != nil {
				return err
			}

			if len(res.Address) == 0 {
				return errors.New("response bytes are empty")
			}

			return clientCtx.PrintProto(res)
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
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			mnemonic := args[0]

			if len(strings.Split(mnemonic, " ")) != 12 {
				return errors.New("input is not a 12-word mnemonic")
			}

			ixoDid, err := exported.FromMnemonic(mnemonic)
			if err != nil {
				return err
			}

			output, err := json.Marshal(ixoDid)
			if err != nil {
				panic(err)
			}

			return clientCtx.PrintString(string(output))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdDidDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-did-doc [did]",
		Short: "Query DidDoc for a DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			didAddr := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.DidDoc(context.Background(),
				&types.QueryDidDocRequest{Did: didAddr})
			if err != nil {
				return err
			}

			if len(res.Diddoc.String()) == 0 {
				return errors.New("response bytes are empty")
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAllDids() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-dids",
		Short: "Query all DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AllDids(context.Background(),
				&types.QueryAllDidsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAllDidDocs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-did-docs",
		Short: "Query all DID documents",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AllDidDocs(context.Background(),
				&types.QueryAllDidDocsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
