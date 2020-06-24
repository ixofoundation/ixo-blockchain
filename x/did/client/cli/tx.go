package cli

import (
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func GetCmdAddDidDoc(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-did-doc [ixo-did]",
		Short: "Add a new IxoDid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ixoDid, err := types.UnmarshalIxoDid(args[0])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(types.DidToAddr(ixoDid.Did))

			msg := types.NewMsgAddDid(ixoDid.Did, ixoDid.VerifyKey)
			return ixo.SignAndBroadcastTxCli(cliCtx, msg, ixoDid)
		},
	}
}

func GetCmdAddCredential(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-kyc-credential [did] [signer-did-doc]",
		Short: "Add a new KYC Credential for a Did by the signer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			didAddr := args[0]

			ixoDid, err := types.UnmarshalIxoDid(args[1])
			if err != nil {
				return err
			}

			t := time.Now()
			issued := t.Format(time.RFC3339)

			credTypes := []string{"Credential", "ProofOfKYC"}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(types.DidToAddr(ixoDid.Did))

			msg := types.NewMsgAddCredential(didAddr, credTypes, ixoDid.Did, issued)
			return ixo.SignAndBroadcastTxCli(cliCtx, msg, ixoDid)
		},
	}
}
