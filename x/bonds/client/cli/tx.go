package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	client2 "github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	bondsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bonds transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bondsTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateBond(cdc),
		GetCmdEditBond(cdc),
		GetCmdSetNextAlpha(cdc),
		GetCmdUpdateBondState(cdc),
		GetCmdBuy(cdc),
		GetCmdSell(cdc),
		GetCmdSwap(cdc),
		GetCmdMakeOutcomePayment(cdc),
		GetCmdWithdrawShare(cdc),
	)...)

	return bondsTxCmd
}

func GetCmdCreateBond(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bond",
		Short: "Create bond",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			_token := viper.GetString(FlagToken)
			_name := viper.GetString(FlagName)
			_description := viper.GetString(FlagDescription)
			_functionType := viper.GetString(FlagFunctionType)
			_functionParameters := viper.GetString(FlagFunctionParameters)
			_reserveTokens := viper.GetString(FlagReserveTokens)
			_txFeePercentage := viper.GetString(FlagTxFeePercentage)
			_exitFeePercentage := viper.GetString(FlagExitFeePercentage)
			_feeAddress := viper.GetString(FlagFeeAddress)
			_maxSupply := viper.GetString(FlagMaxSupply)
			_orderQuantityLimits := viper.GetString(FlagOrderQuantityLimits)
			_sanityRate := viper.GetString(FlagSanityRate)
			_sanityMarginPercentage := viper.GetString(FlagSanityMarginPercentage)
			_allowSells := viper.GetBool(FlagAllowSells)
			_alphaBond := viper.GetBool(FlagAlphaBond)
			_batchBlocks := viper.GetString(FlagBatchBlocks)
			_outcomePayment := viper.GetString(FlagOutcomePayment)
			_bondDid := viper.GetString(FlagBondDid)
			_creatorDid := viper.GetString(FlagCreatorDid)
			_controllerDid := viper.GetString(FlagControllerDid)

			// Parse function parameters
			functionParams, err := client2.ParseFunctionParams(_functionParameters)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			// Parse reserve tokens
			reserveTokens := strings.Split(_reserveTokens, ",")

			// Parse tx fee percentage
			txFeePercentage, err := sdk.NewDecFromStr(_txFeePercentage)
			if err != nil {
				return sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "tx fee percentage")
			}

			// Parse exit fee percentage
			exitFeePercentage, err := sdk.NewDecFromStr(_exitFeePercentage)
			if err != nil {
				return sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "exit fee percentage")
			}

			// Parse fee address
			feeAddress, err := sdk.AccAddressFromBech32(_feeAddress)
			if err != nil {
				return err
			}

			// Parse max supply
			maxSupply, err := sdk.ParseCoin(_maxSupply)
			if err != nil {
				return err
			}

			// Parse order quantity limits
			orderQuantityLimits, err := sdk.ParseCoins(_orderQuantityLimits)
			if err != nil {
				return err
			}

			// Parse sanity rate
			sanityRate, err := sdk.NewDecFromStr(_sanityRate)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			// Parse sanity margin percentage
			sanityMarginPercentage, err := sdk.NewDecFromStr(_sanityMarginPercentage)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			// Parse batch blocks
			batchBlocks, err := sdk.ParseUint(_batchBlocks)
			if err != nil {
				return sdkerrors.Wrap(types.ErrArgumentMissingOrNonUInteger, "max batch blocks")
			}

			// Parse creator's ixo DID
			creatorDid, err := did.UnmarshalIxoDid(_creatorDid)
			if err != nil {
				return err
			}

			// Parse outcome payment
			var outcomePayment sdk.Int
			if len(_outcomePayment) == 0 {
				outcomePayment = sdk.ZeroInt()
			} else {
				var ok bool
				outcomePayment, ok = sdk.NewIntFromString(_outcomePayment)
				if !ok {
					return sdkerrors.Wrap(types.ErrArgumentMustBeInteger, "outcome payment")
				}
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(creatorDid.Address())

			msg := types.NewMsgCreateBond(_token, _name, _description,
				creatorDid.Did, _controllerDid, _functionType, functionParams,
				reserveTokens, txFeePercentage, exitFeePercentage, feeAddress,
				maxSupply, orderQuantityLimits, sanityRate, sanityMarginPercentage,
				_allowSells, _alphaBond, batchBlocks, outcomePayment, _bondDid)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, creatorDid)
		},
	}

	cmd.Flags().AddFlagSet(fsBondGeneral)
	cmd.Flags().AddFlagSet(fsBondCreate)

	_ = cmd.MarkFlagRequired(FlagToken)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagDescription)
	_ = cmd.MarkFlagRequired(FlagFunctionType)
	_ = cmd.MarkFlagRequired(FlagFunctionParameters)
	_ = cmd.MarkFlagRequired(FlagReserveTokens)
	_ = cmd.MarkFlagRequired(FlagTxFeePercentage)
	_ = cmd.MarkFlagRequired(FlagExitFeePercentage)
	_ = cmd.MarkFlagRequired(FlagFeeAddress)
	_ = cmd.MarkFlagRequired(FlagMaxSupply)
	_ = cmd.MarkFlagRequired(FlagOrderQuantityLimits)
	_ = cmd.MarkFlagRequired(FlagSanityRate)
	_ = cmd.MarkFlagRequired(FlagSanityMarginPercentage)
	_ = cmd.MarkFlagRequired(FlagBatchBlocks)
	// _ = cmd.MarkFlagRequired(FlagOutcomePayment) // Optional
	_ = cmd.MarkFlagRequired(FlagBondDid)
	_ = cmd.MarkFlagRequired(FlagCreatorDid)
	_ = cmd.MarkFlagRequired(FlagControllerDid)

	return cmd
}

func GetCmdEditBond(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-bond",
		Short: "Edit bond",
		RunE: func(cmd *cobra.Command, args []string) error {
			_name := viper.GetString(FlagName)
			_description := viper.GetString(FlagDescription)
			_orderQuantityLimits := viper.GetString(FlagOrderQuantityLimits)
			_sanityRate := viper.GetString(FlagSanityRate)
			_sanityMarginPercentage := viper.GetString(FlagSanityMarginPercentage)
			_bondDid := viper.GetString(FlagBondDid)
			_editorDid := viper.GetString(FlagEditorDid)

			// Parse editor's ixo DID
			editorDid, err := did.UnmarshalIxoDid(_editorDid)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(editorDid.Address())

			msg := types.NewMsgEditBond(_name, _description, _orderQuantityLimits,
				_sanityRate, _sanityMarginPercentage, editorDid.Did, _bondDid)
			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, editorDid)
		},
	}

	cmd.Flags().AddFlagSet(fsBondGeneral)
	cmd.Flags().AddFlagSet(fsBondEdit)

	_ = cmd.MarkFlagRequired(FlagBondDid)
	_ = cmd.MarkFlagRequired(FlagEditorDid)

	return cmd
}

func GetCmdSetNextAlpha(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-next-alpha [new-alpha] [bond-did] [editor-did]",
		Example: "set-next-alpha 0.5 U7GK8p8rVhJMKhBVRCJJ8c <editor-ixo-did>",
		Short:   "Edit a bond's alpha parameter",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			_alpha := args[0]
			_bondDid := args[1]
			_editorDid := args[2]

			// Parse alpha
			alpha, err := sdk.NewDecFromStr(_alpha)
			if err != nil {
				return sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "alpha")
			}

			// Parse editor's ixo DID
			editorDid, err := did.UnmarshalIxoDid(_editorDid)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(editorDid.Address())

			msg := types.NewMsgSetNextAlpha(alpha, editorDid.Did, _bondDid)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, editorDid)
		},
	}
	return cmd
}

func GetCmdUpdateBondState(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-bond-state [new-state] [bond-did] [editor-did]",
		Example: "update-bond-state SETTLE U7GK8p8rVhJMKhBVRCJJ8c <editor-ixo-did>",
		Short:   "Edit a bond's current state",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			_state := args[0]
			_bondDid := args[1]
			_editorDid := args[2]

			// Parse editor's ixo DID
			editorDid, err := did.UnmarshalIxoDid(_editorDid)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(editorDid.Address())

			msg := types.NewMsgUpdateBondState(types.BondState(_state), editorDid.Did, _bondDid)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, editorDid)
		},
	}
	return cmd
}

func GetCmdBuy(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "buy [bond-token-with-amount] [max-prices] [bond-did] [buyer-did]",
		Example: "" +
			"buy 10abc 1000res1 U7GK8p8rVhJMKhBVRCJJ8c <buyer-ixo-did>\n" +
			"buy 10abc 1000res1,1000res2 U7GK8p8rVhJMKhBVRCJJ8c <buyer-ixo-did>",
		Short: "Buy from a bond",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			bondCoinWithAmount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			maxPrices, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// Parse buyer's ixo DID
			buyerDid, err := did.UnmarshalIxoDid(args[3])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(buyerDid.Address())

			msg := types.NewMsgBuy(
				buyerDid.Did, bondCoinWithAmount, maxPrices, args[2])

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, buyerDid)
		},
	}
	return cmd
}

func GetCmdSell(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell [bond-token-with-amount] [bond-did] [seller-did]",
		Example: "sell 10abc U7GK8p8rVhJMKhBVRCJJ8c <seller-ixo-did>",
		Short:   "Sell from a bond",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			bondCoinWithAmount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			// Parse seller's ixo DID
			sellerDid, err := did.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(sellerDid.Address())

			msg := types.NewMsgSell(sellerDid.Did, bondCoinWithAmount, args[1])

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, sellerDid)
		},
	}
	return cmd
}

func GetCmdSwap(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "swap [from-amount] [from-token] [to-token] [bond-did] [swapper-did]",
		Example: "" +
			"swap 100 res1 res2 U7GK8p8rVhJMKhBVRCJJ8c <swapper-ixo-did>\n" +
			"swap 100 res2 res1 U7GK8p8rVhJMKhBVRCJJ8c <swapper-ixo-did>",
		Short: "Perform a swap between two tokens",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Check that from amount and token can be parsed to a coin
			from, err := client2.ParseTwoPartCoin(args[0], args[1])
			if err != nil {
				return err
			}

			// Parse swapper's ixo DID
			swapperDid, err := did.UnmarshalIxoDid(args[4])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(swapperDid.Address())

			msg := types.NewMsgSwap(swapperDid.Did, from, args[2], args[3])

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, swapperDid)
		},
	}
	return cmd
}

func GetCmdMakeOutcomePayment(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "make-outcome-payment [bond-did] [amount] [sender-did]",
		Example: "make-outcome-payment U7GK8p8rVhJMKhBVRCJJ8c 100 <sender-ixo-did>",
		Short:   "Make an outcome payment to a bond",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return sdkerrors.Wrap(types.ErrArgumentMustBeInteger, "outcome payment")
			}

			// Parse sender's ixo DID
			sender, err := did.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(sender.Address())

			msg := types.NewMsgMakeOutcomePayment(sender.Did, amount, args[0])

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, sender)
		},
	}
	return cmd
}

func GetCmdWithdrawShare(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw-share [bond-did] [recipient-did]",
		Example: "withdraw-share U7GK8p8rVhJMKhBVRCJJ8c <recipient-ixo-did>",
		Short:   "Withdraw share from a bond that is in settlement state",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Parse recipient's ixo DID
			recipientDid, err := did.UnmarshalIxoDid(args[1])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(recipientDid.Address())

			msg := types.NewMsgWithdrawShare(recipientDid.Did, args[0])

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, recipientDid)
		},
	}
	return cmd
}
