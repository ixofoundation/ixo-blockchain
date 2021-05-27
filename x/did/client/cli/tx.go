package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
	"time"

	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	"github.com/spf13/cobra"
)

func NewCmdAddDidDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-did-doc [ixo-did]",
		Short: "Add a new IxoDid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ixoDid, err := types.UnmarshalIxoDid(args[0])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgAddDid(ixoDid.Did, ixoDid.VerifyKey)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic() ; err != nil {
				return err
			}
			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdAddCredential() *cobra.Command {
	cmd := &cobra.Command{
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

			cliCtx, err := client.GetClientTxContext(cmd)
			cliCtx = cliCtx.WithFromAddress(ixoDid.Address())
			if err != nil {
				return err
			}

			msg := types.NewMsgAddCredential(didAddr, credTypes, ixoDid.Did, issued)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic() ; err != nil {
				return err
			}
			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
