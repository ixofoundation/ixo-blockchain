package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"

	//"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
)

func GetCmdPaymentTemplate(cdc *codec.LegacyAmino) *cobra.Command {
	return &cobra.Command{
		Use:   "payment-template [payment-template-id]",
		Short: "Query info of a payment template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			cliCtx := client.GetClientContextFromCmd(cmd)
			templateId := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s",
					types.QuerierRoute, keeper.QueryPaymentTemplate, templateId), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.PaymentTemplate
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetCmdPaymentContract(cdc *codec.LegacyAmino) *cobra.Command {
	return &cobra.Command{
		Use:   "payment-contract [payment-contract-id]",
		Short: "Query info of a payment contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			cliCtx := client.GetClientContextFromCmd(cmd)
			contractId := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
					keeper.QueryPaymentContract, contractId), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.PaymentContract
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetCmdSubscription(cdc *codec.LegacyAmino) *cobra.Command {
	return &cobra.Command{
		Use:   "subscription [subscription-id]",
		Short: "Query info of a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			cliCtx := client.GetClientContextFromCmd(cmd)
			subscriptionId := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
					keeper.QuerySubscription, subscriptionId), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Subscription
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}
