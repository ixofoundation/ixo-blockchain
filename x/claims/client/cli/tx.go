package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/claims/types"
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
		NewCmdSetCollectionMembers(),
		NewCmdRemoveCollectionMembers(),
		NewCmdUpdateCollectionState(),
		NewCmdUpdateCollectionDates(),
		NewCmdUpdateCollectionPayments(),
		NewCmdUpdateCollectionIntents(),
		NewCmdUpdateCollectionQuota(),
		NewCmdClaimIntent(),
		NewCmdCreateClaimAuthorization(),
		NewCmdUpdateCollectionDisputeConfig(),
		NewCmdAddPerformanceDeposit(),
		NewCmdWithdrawPerformanceDeposit(),
		NewCmdAdjudicateDispute(),
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

func NewCmdSetCollectionMembers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-collection-members [set-collection-members-doc]",
		Short: "Add or update one or more member budgets on a collection - flag is raw json with struct of MsgSetCollectionMembers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgSetCollectionMembers
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

func NewCmdRemoveCollectionMembers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-collection-members [remove-collection-members-doc]",
		Short: "Remove one or more member budgets from a collection - flag is raw json with struct of MsgRemoveCollectionMembers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgRemoveCollectionMembers
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

// jsonCmd is a tiny helper to keep the rest of the file boilerplate-free.
// Every claims tx command follows the same pattern: take one JSON arg,
// unmarshal into the typed Msg, broadcast.
func jsonCmd[M sdk.Msg](use, short string, newMsg func() M) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			msg := newMsg()
			if err := json.Unmarshal([]byte(args[0]), msg); err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// Previously-unwired commands (existed in proto/keeper but missing from CLI).
func NewCmdUpdateCollectionState() *cobra.Command {
	return jsonCmd("update-collection-state [json]",
		"Update collection state - raw json with struct of MsgUpdateCollectionState",
		func() *types.MsgUpdateCollectionState { return &types.MsgUpdateCollectionState{} })
}

func NewCmdUpdateCollectionDates() *cobra.Command {
	return jsonCmd("update-collection-dates [json]",
		"Update collection start/end dates - raw json with struct of MsgUpdateCollectionDates",
		func() *types.MsgUpdateCollectionDates { return &types.MsgUpdateCollectionDates{} })
}

func NewCmdUpdateCollectionPayments() *cobra.Command {
	return jsonCmd("update-collection-payments [json]",
		"Update collection payments - raw json with struct of MsgUpdateCollectionPayments",
		func() *types.MsgUpdateCollectionPayments { return &types.MsgUpdateCollectionPayments{} })
}

func NewCmdUpdateCollectionIntents() *cobra.Command {
	return jsonCmd("update-collection-intents [json]",
		"Update collection intent policy - raw json with struct of MsgUpdateCollectionIntents",
		func() *types.MsgUpdateCollectionIntents { return &types.MsgUpdateCollectionIntents{} })
}

func NewCmdUpdateCollectionQuota() *cobra.Command {
	return jsonCmd("update-collection-quota [json]",
		"Update collection max-claim quota - raw json with struct of MsgUpdateCollectionQuota",
		func() *types.MsgUpdateCollectionQuota { return &types.MsgUpdateCollectionQuota{} })
}

func NewCmdClaimIntent() *cobra.Command {
	return jsonCmd("claim-intent [json]",
		"Submit a claim intent - raw json with struct of MsgClaimIntent",
		func() *types.MsgClaimIntent { return &types.MsgClaimIntent{} })
}

func NewCmdCreateClaimAuthorization() *cobra.Command {
	return jsonCmd("create-claim-authorization [json]",
		"Create a claim authorization - raw json with struct of MsgCreateClaimAuthorization",
		func() *types.MsgCreateClaimAuthorization { return &types.MsgCreateClaimAuthorization{} })
}

// v7 new commands.

func NewCmdUpdateCollectionDisputeConfig() *cobra.Command {
	return jsonCmd("update-collection-dispute-config [json]",
		"Update dispute / performance-deposit config on a collection - raw json with struct of MsgUpdateCollectionDisputeConfig",
		func() *types.MsgUpdateCollectionDisputeConfig { return &types.MsgUpdateCollectionDisputeConfig{} })
}

func NewCmdAddPerformanceDeposit() *cobra.Command {
	return jsonCmd("add-performance-deposit [json]",
		"Top up an agent's performance-deposit balance on a collection - raw json with struct of MsgAddPerformanceDeposit",
		func() *types.MsgAddPerformanceDeposit { return &types.MsgAddPerformanceDeposit{} })
}

func NewCmdWithdrawPerformanceDeposit() *cobra.Command {
	return jsonCmd("withdraw-performance-deposit [json]",
		"Withdraw some/all of an agent's performance-deposit balance on a collection - raw json with struct of MsgWithdrawPerformanceDeposit",
		func() *types.MsgWithdrawPerformanceDeposit { return &types.MsgWithdrawPerformanceDeposit{} })
}

func NewCmdAdjudicateDispute() *cobra.Command {
	return jsonCmd("adjudicate-dispute [json]",
		"Adjudicate an OPEN dispute (AWARDED or DISMISSED) - raw json with struct of MsgAdjudicateDispute",
		func() *types.MsgAdjudicateDispute { return &types.MsgAdjudicateDispute{} })
}
