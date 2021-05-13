package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"github.com/spf13/cobra"
)

func GetCmdPaymentTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment-template [payment-template-id]",
		Short: "Query info of a payment template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			
			templateId := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PaymentTemplate(context.Background(),
				&types.QueryPaymentTemplateRequest{PaymentTemplateId: templateId})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdPaymentContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment-contract [payment-contract-id]",
		Short: "Query info of a payment contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			contractId := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PaymentContract(context.Background(),
				&types.QueryPaymentContractRequest{PaymentContractId: contractId})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdPaymentContractsByIdPrefix() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment-contracts-by-id-prefix [payment-contracts-id-prefix]",
		Short: "Query info of list of payment contracts by ID prefix",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			contractIdPrefix := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PaymentContractsByIdPrefix(context.Background(),
				&types.QueryPaymentContractsByIdPrefixRequest{PaymentContractsIdPrefix: contractIdPrefix})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription [subscription-id]",
		Short: "Query info of a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			subscriptionId := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Subscription(context.Background(),
				&types.QuerySubscriptionRequest{SubscriptionId: subscriptionId})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
