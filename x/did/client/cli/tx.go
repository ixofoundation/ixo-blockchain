package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"time"

	"github.com/spf13/cobra"

	//"github.com/cosmos/cosmos-sdk/client/context"
	//"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func NewCmdAddDidDoc(/*cdc *codec.Codec*/) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-did-doc [ixo-did]",
		Short: "Add a new IxoDid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ixoDid, err := types.UnmarshalIxoDid(args[0])
			if err != nil {
				return err
			}

			//cliCtx := context.NewCLIContext().WithCodec(cdc).
			//	WithFromAddress(ixoDid.Address())
			cliCtx, err := client.GetClientTxContext(cmd) //client.GetClientContextFromCmd(cmd)
			cliCtx = cliCtx.WithFromAddress(ixoDid.Address())
			if err != nil {
				return err
			}

			msg := types.NewMsgAddDid(ixoDid.Did, ixoDid.VerifyKey)
			if err := msg.ValidateBasic() ; err != nil {
				return err
			}
			return ixo.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), ixoDid, msg) //return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
			// TODO we have to prepend an & to msg above because MsgAddDid does not have the gogoproto.nullable = false option set, should it be?
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdAddCredential(/*cdc *codec.Codec*/) *cobra.Command {
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

			//cliCtx := context.NewCLIContext().WithCodec(cdc).
			//	WithFromAddress(ixoDid.Address())

			cliCtx, err := client.GetClientTxContext(cmd) //GetClientContextFromCmd(cmd)
			cliCtx = cliCtx.WithFromAddress(ixoDid.Address())
			if err != nil {
				return err
			}

			msg := types.NewMsgAddCredential(didAddr, credTypes, ixoDid.Did, issued)
			if err := msg.ValidateBasic() ; err != nil {
				return err
			}
			return ixo.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), ixoDid, &msg) //ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
			// TODO we have to prepend an & to msg above because MsgAddCredential does not have the gogoproto.nullable = false option set, should it be?
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
