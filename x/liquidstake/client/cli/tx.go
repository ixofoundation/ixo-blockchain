package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/v6/x/liquidstake/types"
)

// GetTxCmd returns the cli tx commands for x/liquidstake.
//
// Manually registered because cosmos-sdk's autocli + dynamicpb path
// panics on gogoproto-generated Coin messages with
// `cosmos.base.v1beta1.Coin.denom: field descriptor does not belong to
// this message`. The same root cause as the iid nested-message render
// bug — the protoreflect runtime can't fully resolve gogo's Coin
// descriptor. Manual cobra cmds use the legacy gogo-friendly txCLI
// path and parse Coin via sdk.ParseCoinNormalized.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Liquidstake transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		newLiquidStakeCmd(),
		newLiquidUnstakeCmd(),
		newPausePoolCmd(),
		newBurnCmd(),
	)
	return cmd
}

// newLiquidStakeCmd: `tx liquidstake liquid-stake [pool-id] [amount-uixo]`
func newLiquidStakeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquid-stake [pool-id] [amount]",
		Short: "Liquid-stake native uixo into a pool, receive the pool's LST denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amt, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			msg := &types.MsgLiquidStake{
				PoolId:           args[0],
				DelegatorAddress: clientCtx.GetFromAddress().String(),
				Amount:           amt,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// newLiquidUnstakeCmd: `tx liquidstake liquid-unstake [pool-id] [amount-LST]`
func newLiquidUnstakeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquid-unstake [pool-id] [amount]",
		Short: "Liquid-unstake the pool's LST denom back into native uixo",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amt, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			msg := &types.MsgLiquidUnstake{
				PoolId:           args[0],
				DelegatorAddress: clientCtx.GetFromAddress().String(),
				Amount:           amt,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// newPausePoolCmd: `tx liquidstake pause-pool [pool-id] [paused-bool]`
func newPausePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-pool [pool-id] [paused]",
		Short: "Pause or unpause a single pool (gov or pool admin only)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			paused, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}
			msg := &types.MsgSetPoolPaused{
				Authority: clientCtx.GetFromAddress().String(),
				PoolId:    args[0],
				IsPaused:  paused,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// newBurnCmd: `tx liquidstake burn [amount-uixo]`
func newBurnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount]",
		Short: "Burn native uixo tokens (debug helper)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amt, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			msg := &types.MsgBurn{
				Burner: clientCtx.GetFromAddress().String(),
				Amount: amt,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
