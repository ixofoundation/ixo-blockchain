package cli

import (
	"encoding/json"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	tokenTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "token transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tokenTxCmd.AddCommand(
		NewCmdCreateToken(),
		NewCmdMintToken(),
		NewCmdTransferToken(),
		NewCmdCancelToken(),
		NewCmdRetireToken(),
		NewCmdPauseToken(),
		NewCmdStopToken(),
	)

	return tokenTxCmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdUpdateTokenParamsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token-params [ixo1155-code-id] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to update token params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ixo1155CodeId, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			content := types.NewSetTokenContract(ixo1155CodeId)

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func NewCmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [create_token_doc]",
		Short: "Create a new Token - flag is raw json with struct of MsgCreateToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgCreateToken
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.Minter = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdMintToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [mint_token_doc]",
		Short: "Mint new Tokens - flag is raw json with struct of MsgMintToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgMintToken
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.Minter = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdTransferToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [transfer_token_doc]",
		Short: "Transfer Tokens - flag is raw json with struct of MsgTransferToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgTransferToken
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.Owner = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdRetireToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retire-token [retire_token_doc]",
		Short: "Retire Tokens - flag is raw json with struct of MsgRetireToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgRetireToken
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.Owner = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCancelToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-token [cancel_token_doc]",
		Short: "Cancel Tokens - flag is raw json with struct of MsgCancelToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgCancelToken
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.Owner = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdPauseToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-token [contract_address] [paused]",
		Short: "Pause Tokens to temporarily suspend token minting",
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

			msg := types.MsgPauseToken{
				Minter:          clientCtx.GetFromAddress().String(),
				ContractAddress: args[0],
				Paused:          paused,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdStopToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop-token [contract_address]",
		Short: "Stop Tokens to permanently suspend token minting",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgStopToken{
				Minter:          clientCtx.GetFromAddress().String(),
				ContractAddress: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
