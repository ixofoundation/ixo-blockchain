package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	bondsTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateBond(cdc),
		GetCmdEditBond(cdc),
		GetCmdBuy(cdc),
		GetCmdSell(cdc),
		GetCmdSwap(cdc),
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
			_allowSells := viper.GetString(FlagAllowSells)
			_batchBlocks := viper.GetString(FlagBatchBlocks)
			_bondDid := viper.GetString(FlagBondDid)
			_creatorDid := viper.GetString(FlagCreatorDid)

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
				return fmt.Errorf(types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "tx fee percentage").Error())
			}

			// Parse exit fee percentage
			exitFeePercentage, err := sdk.NewDecFromStr(_exitFeePercentage)
			if err != nil {
				return fmt.Errorf(types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "exit fee percentage").Error())
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

			// parse sanity rate
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
				return types.ErrArgumentMissingOrNonUInteger(types.DefaultCodespace, "max batch blocks")
			}

			// Parse creator's ixo DID
			creatorDid, err := did.UnmarshalIxoDid(_creatorDid)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(creatorDid.Address())

			msg := types.NewMsgCreateBond(_token, _name, _description,
				creatorDid.Did, _functionType, functionParams, reserveTokens,
				txFeePercentage, exitFeePercentage, feeAddress, maxSupply,
				orderQuantityLimits, sanityRate, sanityMarginPercentage,
				_allowSells, batchBlocks, _bondDid)

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
	_ = cmd.MarkFlagRequired(FlagAllowSells)
	_ = cmd.MarkFlagRequired(FlagBatchBlocks)
	_ = cmd.MarkFlagRequired(FlagBondDid)
	_ = cmd.MarkFlagRequired(FlagCreatorDid)

	return cmd
}

func GetCmdEditBond(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-bond",
		Short: "Edit bond",
		RunE: func(cmd *cobra.Command, args []string) error {
			_token := viper.GetString(FlagToken)
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

			msg := types.NewMsgEditBond(
				_token, _name, _description, _orderQuantityLimits, _sanityRate,
				_sanityMarginPercentage, editorDid.Did, _bondDid)

			return ixo.GenerateOrBroadcastMsgs(cliCtx, msg, editorDid)
		},
	}

	cmd.Flags().AddFlagSet(fsBondGeneral)
	cmd.Flags().AddFlagSet(fsBondEdit)

	_ = cmd.MarkFlagRequired(FlagToken)
	_ = cmd.MarkFlagRequired(FlagBondDid)
	_ = cmd.MarkFlagRequired(FlagEditorDid)

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
