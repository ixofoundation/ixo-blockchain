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

	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
)

func IxoSignAndBroadcast(cdc *codec.Codec, ctx context.CLIContext, msg sdk.Msg,
	sovrinDid sovrin.SovrinDid) error {
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

func GetCmdSetFeeContractAuthorisation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-fee-contract-authorisation [payer-sovrin-did]",
		Short: "Create and sign a set-fee-contract-authorisation tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide payer private key")
			}

			sovrinDidStr := args[0]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)
			msg := types.MsgSetFeeContractAuthorisation{}

			// TODO: implement

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateFee(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-fee [fee-id] [fee-content] [creator-sovrin-did]",
		Short: "Create and sign a create-fee tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 3 || len(args[0]) == 0 ||
				len(args[1]) == 0 || len(args[2]) == 0 {
				return errors.New("You must provide the fee id, fee content " +
					"json, and creator private key")
			}

			feeIdStr := args[0]
			feeContentStr := args[1]
			sovrinDidStr := args[2]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			var feeContent types.FeeContent
			err := json.Unmarshal([]byte(feeContentStr), &feeContent)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateFee(feeIdStr, feeContent, sovrinDid)

			// TODO: implement properly

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateFeeContract(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-fee-contract [creator-sovrin-did]",
		Short: "Create and sign a create-fee-contract tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide creator private key")
			}

			sovrinDidStr := args[0]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)
			msg := types.MsgCreateFeeContract{}

			// TODO: implement

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdGrantFeeDiscount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "grant-fee-discount [creator-sovrin-did]",
		Short: "Create and sign a grant-fee-discount tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide fee creator private key")
			}

			sovrinDidStr := args[0]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)
			msg := types.MsgGrantFeeDiscount{}

			// TODO: implement

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdRevokeFeeDiscount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke-fee-discount [creator-sovrin-did]",
		Short: "Create and sign a revoke-fee-discount tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide fee creator private key")
			}

			sovrinDidStr := args[0]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)
			msg := types.MsgRevokeFeeDiscount{}

			// TODO: implement

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdChargeFee(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "charge-fee [creator-sovrin-did]",
		Short: "Create and sign a charge-fee tx using DIDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide fee creator private key")
			}

			sovrinDidStr := args[0]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)
			msg := types.MsgChargeFee{}

			// TODO: implement

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}
