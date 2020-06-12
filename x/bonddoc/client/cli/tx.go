package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
)

func GetCmdCreateBond(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-bond [sender-did] [bond-json] [sovrin-did]",
		Short: "Create a new BondDoc signed by the sovrinDID of the bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			bondDocStr := args[1]
			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[2])
			if err != nil {
				return err
			}

			var bondDoc types.BondDoc
			err = json.Unmarshal([]byte(bondDocStr), &bondDoc)
			if err != nil {
				panic(err)
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreateBond(senderDid, bondDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdUpdateBondStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-bond-status [sender-did] [status] [sovrin-did]",
		Short: "Update the status of a bond signed by the sovrinDID of the bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			status := args[1]
			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[2])
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
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgUpdateBondStatus(senderDid, updateBondStatusDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}
