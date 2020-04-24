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

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
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

func GetCmdSend(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "send [to-did] [amount] [sender-sovrin-did]",
		Short: "Create and sign a send tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 3 || len(args[0]) == 0 ||
				len(args[1]) == 0 || len(args[2]) == 0 {
				return errors.New("You must provide the receiver DID, " +
					"amount, and sender private key")
			}

			toDid := args[0]

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			sovrinDid := unmarshalSovrinDID(args[2])
			msg := types.NewMsgSend(toDid, coins, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdMint(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "mint [to-did] [amount] [oracle-sovrin-did]",
		Short: "Create and sign a mint tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 3 || len(args[0]) == 0 ||
				len(args[1]) == 0 || len(args[2]) == 0 {
				return errors.New("You must provide the recipient DID, " +
					"amount, and oracle private key")
			}

			toDid := args[0]

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			sovrinDid := unmarshalSovrinDID(args[2])
			msg := types.NewMsgMint(toDid, coins, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdBurn(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "burn [from-did] [amount] [oracle-sovrin-did]",
		Short: "Create and sign a burn tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 3 || len(args[0]) == 0 ||
				len(args[1]) == 0 || len(args[2]) == 0 {
				return errors.New("You must provide the source DID, " +
					"amount, and oracle private key")
			}

			fromDid := args[0]

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			sovrinDid := unmarshalSovrinDID(args[2])
			msg := types.NewMsgBurn(fromDid, coins, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}
