package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bondsclient "github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	didtypes "github.com/ixofoundation/ixo-blockchain/x/did/types"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
)

func NewTxCmd() *cobra.Command {
	bondsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bonds transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bondsTxCmd.AddCommand(
		NewCmdCreateBond(),
		NewCmdEditBond(),
		NewCmdSetNextAlpha(),
		NewCmdUpdateBondState(),
		NewCmdBuy(),
		NewCmdSell(),
		NewCmdSwap(),
		NewCmdMakeOutcomePayment(),
		NewCmdWithdrawShare(),
		NewCmdWithdrawReserve(),
	)

	return bondsTxCmd
}

func NewCmdCreateBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bond",
		Short: "Create bond",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			_token, _ := cmd.Flags().GetString(FlagToken)
			_name, _ := cmd.Flags().GetString(FlagName)
			_description, _ := cmd.Flags().GetString(FlagDescription)
			_functionType, _ := cmd.Flags().GetString(FlagFunctionType)
			_functionParameters, _ := cmd.Flags().GetString(FlagFunctionParameters)
			_reserveTokens, _ := cmd.Flags().GetString(FlagReserveTokens)
			_txFeePercentage, _ := cmd.Flags().GetString(FlagTxFeePercentage)
			_exitFeePercentage, _ := cmd.Flags().GetString(FlagExitFeePercentage)
			_feeAddress, _ := cmd.Flags().GetString(FlagFeeAddress)
			_reserveWithdrawalAddress, _ := cmd.Flags().GetString(FlagReserveWithdrawalAddress)
			_maxSupply, _ := cmd.Flags().GetString(FlagMaxSupply)
			_orderQuantityLimits, _ := cmd.Flags().GetString(FlagOrderQuantityLimits)
			_sanityRate, _ := cmd.Flags().GetString(FlagSanityRate)
			_sanityMarginPercentage, _ := cmd.Flags().GetString(FlagSanityMarginPercentage)
			_allowSells, _ := cmd.Flags().GetBool(FlagAllowSells)
			_allowReserveWithdrawals, _ := cmd.Flags().GetBool(FlagAllowReserveWithdrawals)
			_alphaBond, _ := cmd.Flags().GetBool(FlagAlphaBond)
			_batchBlocks, _ := cmd.Flags().GetString(FlagBatchBlocks)
			_outcomePayment, _ := cmd.Flags().GetString(FlagOutcomePayment)
			_bondDid, _ := cmd.Flags().GetString(FlagBondDid)
			_creatorDid, _ := cmd.Flags().GetString(FlagCreatorDid)
			_controllerDid, _ := cmd.Flags().GetString(FlagControllerDid)

			// Parse function parameters
			functionParams, err := bondsclient.ParseFunctionParams(_functionParameters)
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

			// Parse reserve withdrawal address
			reserveWithdrawalAddress, err := sdk.AccAddressFromBech32(_reserveWithdrawalAddress)
			if err != nil {
				return err
			}

			// Parse max supply
			maxSupply, err := sdk.ParseCoinNormalized(_maxSupply) //sdk.ParseCoin(_maxSupply)
			if err != nil {
				return err
			}

			// Parse order quantity limits
			orderQuantityLimits, err := sdk.ParseCoinsNormalized(_orderQuantityLimits) //sdk.ParseCoins(_orderQuantityLimits)
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
			creatorDid, err := didtypes.UnmarshalIxoDid(_creatorDid)
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

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(creatorDid.Address())

			msg := types.NewMsgCreateBond(_token, _name, _description,
				creatorDid.Did, _controllerDid, _functionType, functionParams,
				reserveTokens, txFeePercentage, exitFeePercentage, feeAddress,
				reserveWithdrawalAddress, maxSupply, orderQuantityLimits,
				sanityRate, sanityMarginPercentage, _allowSells,
				_allowReserveWithdrawals, _alphaBond, batchBlocks,
				outcomePayment, _bondDid)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), creatorDid, msg)
		},
	}

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
	_ = cmd.MarkFlagRequired(FlagReserveWithdrawalAddress)
	_ = cmd.MarkFlagRequired(FlagMaxSupply)
	_ = cmd.MarkFlagRequired(FlagOrderQuantityLimits)
	_ = cmd.MarkFlagRequired(FlagSanityRate)
	_ = cmd.MarkFlagRequired(FlagSanityMarginPercentage)
	//_ = cmd.MarkFlagRequired(FlagAllowSells) // If not specified, assume False
	//_ = cmd.MarkFlagRequired(FlagAllowReserveWithdrawals) // If not specified, assume False
	_ = cmd.MarkFlagRequired(FlagBatchBlocks)
	// _ = cmd.MarkFlagRequired(FlagOutcomePayment) // Optional
	_ = cmd.MarkFlagRequired(FlagBondDid)
	_ = cmd.MarkFlagRequired(FlagCreatorDid)
	_ = cmd.MarkFlagRequired(FlagControllerDid)

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdEditBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-bond",
		Short: "Edit bond",
		RunE: func(cmd *cobra.Command, args []string) error {
			_name, _ := cmd.Flags().GetString(FlagName)
			_description, _ := cmd.Flags().GetString(FlagDescription)
			_orderQuantityLimits, _ := cmd.Flags().GetString(FlagOrderQuantityLimits)
			_sanityRate, _ := cmd.Flags().GetString(FlagSanityRate)
			_sanityMarginPercentage, _ := cmd.Flags().GetString(FlagSanityMarginPercentage)
			_bondDid, _ := cmd.Flags().GetString(FlagBondDid)
			_editorDid, _ := cmd.Flags().GetString(FlagEditorDid)

			// Parse editor's ixo DID
			editorDid, err := didtypes.UnmarshalIxoDid(_editorDid)
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(editorDid.Address())

			msg := types.NewMsgEditBond(_name, _description, _orderQuantityLimits,
				_sanityRate, _sanityMarginPercentage, editorDid.Did, _bondDid)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), editorDid, msg)
		},
	}

	cmd.Flags().AddFlagSet(fsBondEdit)

	_ = cmd.MarkFlagRequired(FlagBondDid)
	_ = cmd.MarkFlagRequired(FlagEditorDid)

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdSetNextAlpha() *cobra.Command {
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
			editorDid, err := didtypes.UnmarshalIxoDid(_editorDid)
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(editorDid.Address())

			msg := types.NewMsgSetNextAlpha(alpha, editorDid.Did, _bondDid)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), editorDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateBondState() *cobra.Command {
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
			editorDid, err := didtypes.UnmarshalIxoDid(_editorDid)
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(editorDid.Address())

			msg := types.NewMsgUpdateBondState(types.BondState(_state), editorDid.Did, _bondDid)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), editorDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdBuy() *cobra.Command {
	cmd := &cobra.Command{
		Use: "buy [bond-token-with-amount] [max-prices] [bond-did] [buyer-did]",
		Example: "" +
			"buy 10abc 1000res1 U7GK8p8rVhJMKhBVRCJJ8c <buyer-ixo-did>\n" +
			"buy 10abc 1000res1,1000res2 U7GK8p8rVhJMKhBVRCJJ8c <buyer-ixo-did>",
		Short: "Buy from a bond",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			bondCoinWithAmount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			maxPrices, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			// Parse buyer's ixo DID
			buyerDid, err := didtypes.UnmarshalIxoDid(args[3])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(buyerDid.Address())

			msg := types.NewMsgBuy(
				buyerDid.Did, bondCoinWithAmount, maxPrices, args[2])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), buyerDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdSell() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell [bond-token-with-amount] [bond-did] [seller-did]",
		Example: "sell 10abc U7GK8p8rVhJMKhBVRCJJ8c <seller-ixo-did>",
		Short:   "Sell from a bond",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			bondCoinWithAmount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			// Parse seller's ixo DID
			sellerDid, err := didtypes.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(sellerDid.Address())

			msg := types.NewMsgSell(sellerDid.Did, bondCoinWithAmount, args[1])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), sellerDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use: "swap [from-amount] [from-token] [to-token] [bond-did] [swapper-did]",
		Example: "" +
			"swap 100 res1 res2 U7GK8p8rVhJMKhBVRCJJ8c <swapper-ixo-did>\n" +
			"swap 100 res2 res1 U7GK8p8rVhJMKhBVRCJJ8c <swapper-ixo-did>",
		Short: "Perform a swap between two tokens",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Check that from amount and token can be parsed to a coin
			from, err := bondsclient.ParseTwoPartCoin(args[0], args[1])
			if err != nil {
				return err
			}

			// Parse swapper's ixo DID
			swapperDid, err := didtypes.UnmarshalIxoDid(args[4])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(swapperDid.Address())

			msg := types.NewMsgSwap(swapperDid.Did, from, args[2], args[3])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), swapperDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdMakeOutcomePayment() *cobra.Command {
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
			sender, err := didtypes.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(sender.Address())

			msg := types.NewMsgMakeOutcomePayment(sender.Did, amount, args[0])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), sender, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdWithdrawShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw-share [bond-did] [recipient-did]",
		Example: "withdraw-share U7GK8p8rVhJMKhBVRCJJ8c <recipient-ixo-did>",
		Short:   "Withdraw share from a bond that is in settlement state",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Parse recipient's ixo DID
			recipientDid, err := didtypes.UnmarshalIxoDid(args[1])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(recipientDid.Address())

			msg := types.NewMsgWithdrawShare(recipientDid.Did, args[0])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), recipientDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdWithdrawReserve() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw-reserve [bond-did] [amount] [withdrawer-did]",
		Example: "withdraw-reserve U7GK8p8rVhJMKhBVRCJJ8c 1000res <withdrawer-ixo-did>",
		Short:   "Withdraw reserve from a bond",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			// Parse withdrawer's ixo DID
			withdrawerDid, err := didtypes.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithFromAddress(withdrawerDid.Address())

			msg := types.NewMsgWithdrawReserve(withdrawerDid.Did, amount, args[0])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), withdrawerDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
