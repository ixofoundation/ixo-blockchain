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
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	entityTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "entity transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	entityTxCmd.AddCommand(
		NewCmdCreateEntity(),
	)

	return entityTxCmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdUpdateEntityParamsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-entity-params [nft_contract_code] [nft_minter_address] [flags]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit a proposal to update entity params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeId, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			content := types.NewInitializeNftContract(codeId, args[1])

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

func NewCmdCreateEntity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-entity [entity-iid]",
		Short: "Create a new EntityDoc",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var msg types.MsgCreateEntity
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

			msg.OwnerAddress = clientCtx.GetFromAddress().String()

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
