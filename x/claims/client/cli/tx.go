package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ixofoundation/ixo-blockchain/v4/x/claims/types"
)

// GetTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Claims transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCmdCreateCollection(),
		NewCmdSubmitClaim(),
		NewCmdEvaluateClaim(),
		NewCmdDisputeClaim(),
		NewCmdWithdrawPayment(),
	)

	return cmd
}

func NewCmdCreateCollection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-collection [create-collection-doc]",
		Short: "Create a new collection - flag is raw json with struct of MsgCreateCollection",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgCreateCollection
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdSubmitClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-claim [submit-claim-doc]",
		Short: "Submit a new Claim - flag is raw json with struct of MsgSubmitClaim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgSubmitClaim
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdEvaluateClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "evaluate-claim [evaluate-claim-doc]",
		Short: "Evaluate a Claim - flag is raw json with struct of MsgEvaluateClaim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgEvaluateClaim
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdDisputeClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dispute-claim [dispute-claim-doc]",
		Short: "Dispute a Claim - flag is raw json with struct of MsgDisputeClaim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgDisputeClaim
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdWithdrawPayment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-payment [withdraw-payment-doc]",
		Short: "Withdraw a payment - flag is raw json with struct of MsgWithdrawPayment",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgWithdrawPayment
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdSubmitIntent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-intent [claim-intent-doc]",
		Short: "Submit a claim - flag is raw json with struct of MsgClaimIntent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgClaimIntent
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
