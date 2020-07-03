package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"

	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

func GetCmdSend(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "send [to-did] [amount] [sender-ixo-did]",
		Short: "Create and sign a send tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			toDid := args[0]
			coinsStr := args[1]
			ixoDidStr := args[2]

			coins, err := sdk.ParseCoins(coinsStr)
			if err != nil {
				return err
			}

			ixoDid, err := did.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixoDid.Address())

			msg := types.NewMsgSend(toDid, coins, ixoDid.Did)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
		},
	}
}

func GetCmdOracleTransfer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "oracle-transfer [from-did] [to-did] [amount] [oracle-ixo-did] [proof]",
		Short: "Create and sign an oracle-transfer tx using DIDs",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			fromDid := args[0]
			toDid := args[1]
			coinsStr := args[2]
			ixoDidStr := args[3]
			proof := args[4]

			coins, err := sdk.ParseCoins(coinsStr)
			if err != nil {
				return err
			}

			ixoDid, err := did.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixoDid.Address())

			msg := types.NewMsgOracleTransfer(
				fromDid, toDid, coins, ixoDid.Did, proof)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
		},
	}
}

func GetCmdOracleMint(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "oracle-mint [to-did] [amount] [oracle-ixo-did] [proof]",
		Short: "Create and sign an oracle-mint tx using DIDs",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			toDid := args[0]
			coinsStr := args[1]
			ixoDidStr := args[2]
			proof := args[3]

			coins, err := sdk.ParseCoins(coinsStr)
			if err != nil {
				return err
			}

			ixoDid, err := did.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixoDid.Address())

			msg := types.NewMsgOracleMint(
				toDid, coins, ixoDid.Did, proof)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
		},
	}
}

func GetCmdOracleBurn(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "oracle-burn [from-did] [amount] [oracle-ixo-did] [proof]",
		Short: "Create and sign an oracle-burn tx using DIDs",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			fromDid := args[0]
			coinsStr := args[1]
			ixoDidStr := args[2]
			proof := args[3]

			coins, err := sdk.ParseCoins(coinsStr)
			if err != nil {
				return err
			}

			ixoDid, err := did.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixoDid.Address())

			msg := types.NewMsgOracleBurn(
				fromDid, coins, ixoDid.Did, proof)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, ixoDid)
		},
	}
}
