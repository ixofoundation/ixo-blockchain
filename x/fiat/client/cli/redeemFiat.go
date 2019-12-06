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

			issuerAddressStr := viper.GetString(FlagIssuerAddress)
			if issuerAddressStr == "" {
				return sdk.ErrInvalidAddress("Issuer address empty.")
			}
			issuerAddress, err := sdk.AccAddressFromBech32(issuerAddressStr)
			if err != nil {
				return nil
			}

			cliCtx := context.NewCLIContextWithFrom(issuerAddressStr).WithCodec(cdc)

			redeemerAddressStr := viper.GetString(FlagRedeemerAddress)
			if redeemerAddressStr == "" {
				return sdk.ErrInvalidAddress("Redeemer address empty.")
			}
			redeemerAddress, err := sdk.AccAddressFromBech32(redeemerAddressStr)
			if err != nil {
				return nil
			}

			amount := viper.GetInt64(FlagTransactionAmount)
			if amount <= 0 {
				return sdk.ErrInvalidCoins("Invalid amount.")
			}

			msg := client.BuildRedeemFiatMsg(redeemerAddress, issuerAddress, amount)
			return cUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(fsIssuerAddress)
	cmd.Flags().AddFlagSet(fsRedeemerAddress)
	cmd.Flags().AddFlagSet(fsTransactionAmount)
	cmd.Flags().AddFlagSet(fsTransactionID)
	return cmd
}
