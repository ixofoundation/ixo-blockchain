package cli

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func IxoSignAndBroadcast(cdc *codec.Codec, ctx context.CLIContext, msg sdk.Msg, sovrinDid sovrin.SovrinDid) error {
	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
	tx := ixo.NewIxoTxSingleMsg(msg, signature)

	bz, err := cdc.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}

	res, err := ctx.BroadcastTx(bz)
	if err != nil {
		return err
	}

	fmt.Println(res.String())
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.TxHash)
	return nil

}

func unmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}

	return sovrinDid
}

func CreateBondCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createBond [bond-json] [sovrin-did]",
		Short: "Create a new BondDoc signed by the sovrinDID of the bond",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the bond data and the bonds private key")
			}

			bondDoc := types.BondDoc{}
			err := json.Unmarshal([]byte(args[0]), &bondDoc)
			if err != nil {
				panic(err)
			}

			sovrinDid := unmarshalSovrinDID(args[1])
			msg := types.NewCreateBondMsg(bondDoc, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func UpdateBondStatusCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "updateBondStatus [tx-hash] [sender-did] [status] [sovrin-did]",
		Short: "Update the status of a bond signed by the sovrinDID of the bond",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 4 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 {
				return errors.New("You must provide the status and the bonds private key")
			}

			txHash := args[0]
			senderDid := args[1]

			bondStatus := types.BondStatus(args[2])
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

			sovrinDid := unmarshalSovrinDID(args[3])
			msg := types.NewUpdateBondStatusMsg(txHash, senderDid, updateBondStatusDoc, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}
