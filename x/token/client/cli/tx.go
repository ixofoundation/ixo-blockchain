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
	)

	return tokenTxCmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdUpdateTokenParamsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token-params [nft_contract_code] [nft_minter_address] [flags]",
		Args:  cobra.ExactArgs(3),
		Short: "Submit a proposal to update token params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cw20CodeId, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			cw721CodeId, err := strconv.ParseUint(args[1], 0, 64)
			if err != nil {
				return err
			}

			ixo1155CodeId, err := strconv.ParseUint(args[2], 0, 64)
			if err != nil {
				return err
			}

			content := types.NewSetTokenContract(cw20CodeId, cw721CodeId, ixo1155CodeId)

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
		Use:   "create-token [token-iid]",
		Short: "Create a new TokenDoc",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var msg types.MsgMint
			err := json.Unmarshal([]byte(args[0]), &msg)
			if err != nil {
				return err
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.MinterAddress = clientCtx.GetFromAddress().String()

			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
