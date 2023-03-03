package cli

import (
	"encoding/json"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
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
		NewCmdUpdateEntity(),
		NewCmdUpdateEntityVerified(),
		NewCmdTransferEntity(),
	)

	return entityTxCmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdUpdateEntityParamsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-entity-params [nft-contract-code] [nft-minter-address] [flags]",
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
		Use:   "create [create-entity-doc]",
		Short: "Create a new Entity - flag is raw json with struct of MsgCreateEntity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgCreateEntity
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var verJson iidtypes.VerificationsJSON
			if err := json.Unmarshal([]byte(args[0]), &verJson); err != nil {
				return err
			}

			// Manually generate verifications based of json values
			verifications, err := iidtypes.GenerateVerificationsFromJson(verJson)
			if err != nil {
				return err
			}

			msg.Verification = verifications
			msg.OwnerAddress = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// When using this function it updates all fields, even if dopnt provide fields it will use the proto defaults
func NewCmdUpdateEntity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [update-entity-doc]",
		Short: "Update an Entity - flag is raw json with struct of MsgUpdateEntity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgUpdateEntity
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.ControllerAddress = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateEntityVerified() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-entity-verified [id] [relayer-did] [verified]",
		Short: "Update if an Entity is verified, only the relayer-node can verify",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argId := args[0]
			argRelayerDid := args[1]
			argVerified, err := strconv.ParseBool(args[2])
			if err != nil {
				return sdkerrors.Wrapf(err, "verified must be a boolean value")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateEntityVerified{
				Id:                 argId,
				RelayerNodeDid:     iidtypes.DIDFragment(argRelayerDid),
				EntityVerified:     argVerified,
				RelayerNodeAddress: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdTransferEntity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [id] [owner-did] [recipient-did]",
		Short: "Transfer an Entity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argId := args[0]
			argOwnerDid := args[1]
			argRecipientDid := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgTransferEntity{
				Id:           argId,
				OwnerDid:     iidtypes.DIDFragment(argOwnerDid),
				RecipientDid: iidtypes.DIDFragment(argRecipientDid),
				OwnerAddress: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
