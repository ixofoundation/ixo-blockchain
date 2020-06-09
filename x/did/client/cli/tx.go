package cli

import (
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
)

func GetCmdAddDidDoc(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-did-doc [sovrin-did]",
		Short: "Add a new SovrinDid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddDid(sovrinDid.Did, sovrinDid.VerifyKey)
			return ixo.SignAndBroadcastTxCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdAddCredential(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-kyc-credential [did] [signer-did-doc]",
		Short: "Add a new KYC Credential for a Did by the signer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			didAddr := args[0]

			_, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, keeper.QueryDidDoc, didAddr), nil)
			if err != nil {
				return errors.New("The did is not on the blockchain")
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[1])
			if err != nil {
				return err
			}

			t := time.Now()
			issued := t.Format(time.RFC3339)

			credTypes := []string{"Credential", "ProofOfKYC"}

			msg := types.NewMsgAddCredential(didAddr, credTypes, sovrinDid.Did, issued)
			return ixo.SignAndBroadcastTxCli(ctx, msg, sovrinDid)
		},
	}
}
