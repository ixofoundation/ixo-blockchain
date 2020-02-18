package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	client2 "github.com/ixofoundation/ixo-cosmos/x/bonds/client"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			_signers := viper.GetString(FlagSigners)
			_batchBlocks := viper.GetString(FlagBatchBlocks)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Parse function parameters
			functionParams, err := client2.ParseFunctionParams(_functionParameters, _functionType)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			// Parse reserve tokens
			reserveTokens, err := client2.ParseReserveTokens(_reserveTokens, _functionType, _token)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			txFeePercentage, err := sdk.NewDecFromStr(_txFeePercentage)
			if err != nil {
				return fmt.Errorf(types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "tx fee percentage").Error())
			}

			exitFeePercentage, err := sdk.NewDecFromStr(_exitFeePercentage)
			if err != nil {
				return fmt.Errorf(types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "exit fee percentage").Error())
			}

			if txFeePercentage.Add(exitFeePercentage).GTE(sdk.NewDec(100)) {
				return fmt.Errorf(types.ErrFeesCannotBeOrExceed100Percent(types.DefaultCodespace).Error())
			}

			feeAddress, err := sdk.AccAddressFromBech32(_feeAddress)
			if err != nil {
				return err
			}

			maxSupply, err := client2.ParseMaxSupply(_maxSupply, _token)
			if err != nil {
				return err
			}

			orderQuantityLimits, err := sdk.ParseCoins(_orderQuantityLimits)
			if err != nil {
				return err
			}

			// Parse sanity
			sanityRate, sanityMarginPercentage, err := client2.ParseSanityValues(_sanityRate, _sanityMarginPercentage)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			// Parse signers
			signers, err := client2.ParseSigners(_signers)
			if err != nil {
				return err
			}

			// Parse batch blocks
			batchBlocks, err := client2.ParseBatchBlocks(_batchBlocks)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			msg := types.NewMsgCreateBond(_token, _name, _description,
				cliCtx.GetFromAddress(), _functionType, functionParams,
				reserveTokens, txFeePercentage, exitFeePercentage, feeAddress,
				maxSupply, orderQuantityLimits, sanityRate,
				sanityMarginPercentage, _allowSells, signers, batchBlocks)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(fsBondGeneral)
	cmd.Flags().AddFlagSet(fsBondCreate)

	_ = cmd.MarkFlagRequired(client.FlagFrom)
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
	_ = cmd.MarkFlagRequired(FlagSigners)
	_ = cmd.MarkFlagRequired(FlagBatchBlocks)

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
			_signers := viper.GetString(FlagSigners)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Parse signers
			signers, err := client2.ParseSigners(_signers)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditBond(
				_token, _name, _description, _orderQuantityLimits, _sanityRate,
				_sanityMarginPercentage, cliCtx.GetFromAddress(), signers)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(fsBondGeneral)
	cmd.Flags().AddFlagSet(fsBondEdit)

	_ = cmd.MarkFlagRequired(client.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagToken)
	_ = cmd.MarkFlagRequired(FlagSigners)

	return cmd
}

func GetCmdBuy(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "buy [bond-token-with-amount] [max-prices]",
		Example: "" +
			"buy 10abc 1000res1\n" +
			"buy 10abc 1000res1,1000res2",
		Short: "Buy from a bond",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bondCoinWithAmount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			maxPrices, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgBuy(cliCtx.GetFromAddress(),
				bondCoinWithAmount, maxPrices)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	_ = cmd.MarkFlagRequired(client.FlagFrom)
	return cmd
}

func GetCmdSell(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell [bond-token-with-amount]",
		Example: "sell 10abc",
		Short:   "Sell from a bond",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bondCoinWithAmount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSell(cliCtx.GetFromAddress(), bondCoinWithAmount)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	_ = cmd.MarkFlagRequired(client.FlagFrom)
	return cmd
}

func GetCmdSwap(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "swap [bond_token] [from_amount] [from_token] [to_token]",
		Example: "" +
			"swap abc 100 res1 res2\n" +
			"swap abc 100 res2 res1",
		Short: "Perform a swap between two tokens",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Check that bond_token is a valid token name
			_, err := sdk.ParseCoin("0" + args[0])
			if err != nil {
				return types.ErrInvalidCoinDenomination(types.DefaultCodespace, args[0])
			}

			// Check that from amount and token can be parsed to a coin
			from, err := sdk.ParseCoin(args[1] + args[2])
			if err != nil {
				return err
			}

			// Check that to_token is a valid token name
			_, err = sdk.ParseCoin("0" + args[3])
			if err != nil {
				return types.ErrInvalidCoinDenomination(types.DefaultCodespace, args[3])
			}

			msg := types.NewMsgSwap(cliCtx.GetFromAddress(), args[0], from, args[3])
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	_ = cmd.MarkFlagRequired(client.FlagFrom)
	return cmd
}
