package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	cUtils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ixofoundation/ixo-cosmos/x/fiat/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func SendFiatCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sendFiat",
		Short: "Send fiat from account",
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			fromStr := viper.GetString(FlagFrom)
			if fromStr == "" {
				return sdk.ErrInvalidAddress("From address empty.")
			}
			from, err := sdk.AccAddressFromBech32(fromStr)
			if err != nil {
				return nil
			}

			toStr := viper.GetString(FlagTo)
			if toStr == "" {
				return sdk.ErrInvalidAddress("To address empty.")
			}
			to, err := sdk.AccAddressFromBech32(toStr)
			if err != nil {
				return nil
			}

			amount := viper.GetInt64(FlagAmount)
			if amount <= 0 {
				return sdk.ErrInvalidCoins("Invalid amount.")
			}

			msg := client.BuildSendFiatMsg(from, to, amount)
			return cUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(fsFrom)
	cmd.Flags().AddFlagSet(fsTo)
	cmd.Flags().AddFlagSet(fsAmount)
	return cmd
}
