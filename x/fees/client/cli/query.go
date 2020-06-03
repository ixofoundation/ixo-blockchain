package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

func GetParamsRequestHandler(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query params",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
				keeper.QueryParams), nil)
			if err != nil {
				return err
			}

			var params types.Params
			if err := cdc.UnmarshalJSON(bz, &params); err != nil {
				return err
			}

			fmt.Println(string(bz))
			return nil
		},
	}
}

func GetCmdFee(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "fee [fee-id]",
		Short: "Query info of a fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			feeId := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s",
					types.QuerierRoute, keeper.QueryFee, feeId), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Fee
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

func GetCmdFeeContract(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "fee-contract [fee-contract-id]",
		Short: "Query info of a fee contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			feeContractId := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
					keeper.QueryFeeContract, feeContractId), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.FeeContract
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

func GetCmdSubscription(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "subscription [subscription-id]",
		Short: "Query info of a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
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
