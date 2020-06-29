package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/types"
)

func GetCmdCreateBond(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-bond [sender-did] [bond-json] [ixo-did]",
		Short: "Create a new BondDoc signed by the ixoDid of the bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			bondDocStr := args[1]
			ixoDid, err := did.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			var bondDoc types.BondDoc
			err = json.Unmarshal([]byte(bondDocStr), &bondDoc)
			if err != nil {
				panic(err)
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(did.DidToAddr(ixoDid.Did))

			msg := types.NewMsgCreateBond(senderDid, bondDoc, ixoDid)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
		},
	}
}

func GetCmdUpdateBondStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-bond-status [sender-did] [status] [ixo-did]",
		Short: "Update the status of a bond signed by the ixoDid of the bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			status := args[1]
			ixoDid, err := did.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			bondStatus := types.BondStatus(status)
			if bondStatus != types.PreIssuanceStatus &&
				bondStatus != types.OpenStatus &&
				bondStatus != types.SuspendedStatus &&
				bondStatus != types.ClosedStatus &&
				bondStatus != types.SettlementStatus &&
				bondStatus != types.EndedStatus {
				return errors.New("The status must be one of 'PREISSUANCE', " +
					"'OPEN', 'SUSPENDED', 'CLOSED', 'SETTLEMENT' or 'ENDED'")
			}

			updateBondStatusDoc := types.UpdateBondStatusDoc{
				Status: bondStatus,
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(did.DidToAddr(ixoDid.Did))

			msg := types.NewMsgUpdateBondStatus(senderDid, updateBondStatusDoc, ixoDid)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
		},
	}
}
