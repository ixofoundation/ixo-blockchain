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

func RedeemFiatCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeemFiat",
		Short: "Redeem fiat from account",
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			toAddrStr := viper.GetString(FlagTo)
			if toAddrStr == "" {
				return sdk.ErrInvalidAddress("To address empty.")
			}
			toAddress, err := sdk.AccAddressFromBech32(toAddrStr)
			if err != nil {
				return err
			}

			amount := viper.GetInt64(FlagAmount)
			if amount <= 0 {
				return sdk.ErrInvalidCoins("Invalid amount.")
			}

			msg := client.BuildRedeemFiatMsg(cliCtx.GetFromAddress(), toAddress, amount)
			return cUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(fsTo)
	cmd.Flags().AddFlagSet(fsAmount)
	return cmd
}
