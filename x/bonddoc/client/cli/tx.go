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

func unmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}

	return sovrinDid
}

func GetCmdCreateBond(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-bond [sender-did] [bond-json] [sovrin-did]",
		Short: "Create a new BondDoc signed by the sovrinDID of the bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			senderDid := args[0]
			bondDocStr := args[1]
			sovrinDid := unmarshalSovrinDID(args[2])

			bondDoc := types.BondDoc{}
			err := json.Unmarshal([]byte(bondDocStr), &bondDoc)
			if err != nil {
				panic(err)
			}

			msg := types.NewMsgCreateBond(senderDid, bondDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdUpdateBondStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-bond-status [sender-did] [status] [sovrin-did]",
		Short: "Update the status of a bond signed by the sovrinDID of the bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			senderDid := args[0]
			status := args[1]
			sovrinDid := unmarshalSovrinDID(args[2])

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

			msg := types.NewMsgUpdateBondStatus(senderDid, updateBondStatusDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}
