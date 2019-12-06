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

func IssueFiatCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issueFiat",
		Short: "Issue fiat to account",
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			issuerAddressStr := viper.GetString(FlagIssuerAddress)
			if issuerAddressStr == "" {
				return sdk.ErrInvalidAddress("Issuer address empty.")
			}
			issuerAddress, err := sdk.AccAddressFromBech32(issuerAddressStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContextWithFrom(issuerAddressStr).WithCodec(cdc)

			toStr := viper.GetString(FlagTo)
			if toStr == "" {
				return sdk.ErrInvalidAddress("To address empty.")
			}
			to, err := sdk.AccAddressFromBech32(toStr)
			if err != nil {
				return err
			}

			transactionAmount := viper.GetInt64(FlagTransactionAmount)
			if transactionAmount <= 0 {
				return sdk.ErrInvalidCoins("Invalid amount.")
			}

			tranasctionID := viper.GetString(FlagTransactionID)
			if tranasctionID == "" {
				return sdk.ErrInvalidAddress("Invalid Transaction ID.")
			}

			msg := client.BuildIssueFiatMsg(issuerAddress, to, tranasctionID, transactionAmount)
			return cUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().AddFlagSet(fsIssuerAddress)
	cmd.Flags().AddFlagSet(fsTo)
	cmd.Flags().AddFlagSet(fsTransactionAmount)
	cmd.Flags().AddFlagSet(fsTransactionID)
	return cmd
}
