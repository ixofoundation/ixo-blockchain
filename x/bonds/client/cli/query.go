package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v2/x/bonds/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	bondsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bonds querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bondsQueryCmd.AddCommand(
		GetCmdBondsList(),
		GetCmdBondsListDetailed(),
		GetCmdBond(),
		GetCmdBatch(),
		GetCmdLastBatch(),
		GetCmdCurrentPrice(),
		GetCmdCurrentReserve(),
		GetCmdAvailableReserve(),
		GetCmdCustomPrice(),
		GetCmdBuyPrice(),
		GetCmdSellReturn(),
		GetCmdSwapReturn(),
		GetCmdAlphaMaximums(),
		GetParamsRequestHandler(),
	)

	return bondsQueryCmd
}

func GetCmdBondsList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonds-list",
		Short: "List of all bonds",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Bonds(context.Background(), &types.QueryBondsRequest{})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			if len(res.GetBonds()) == 0 {
				return fmt.Errorf("no bonds found")
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBondsListDetailed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonds-list-detailed",
		Short: "List of all bonds with information about current state",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.BondsDetailed(context.Background(), &types.QueryBondsDetailedRequest{})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			if len(res.GetBondsDetailed()) == 0 {
				return fmt.Errorf("no bonds found")
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond [bond-did]",
		Short: "Query info of a bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Bond(context.Background(), &types.QueryBondRequest{BondDid: bondDid})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res.Bond)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch [bond-did]",
		Short: "Query info of a bond's current batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Batch(context.Background(), &types.QueryBatchRequest{BondDid: bondDid})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res.Batch)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdLastBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last-batch [bond-did]",
		Short: "Query info of a bond's last batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.LastBatch(context.Background(), &types.QueryLastBatchRequest{BondDid: bondDid})
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			return clientCtx.PrintProto(res.LastBatch)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdCurrentPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current-price [bond-did]",
		Short: "Query current price(s) of the bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.CurrentPrice(context.Background(), &types.QueryCurrentPriceRequest{BondDid: bondDid})
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

func GetCmdCurrentReserve() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "current-reserve [bond-did]",
		Example: "current-reserve U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query current balance(s) of the reserve pool",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.CurrentReserve(context.Background(), &types.QueryCurrentReserveRequest{BondDid: bondDid})
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

func GetCmdAvailableReserve() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "available-reserve [bond-did]",
		Example: "available-reserve U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query current available balance(s) of the reserve pool",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AvailableReserve(context.Background(), &types.QueryAvailableReserveRequest{BondDid: bondDid})
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

func GetCmdCustomPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "price [bond-token-with-amount] [bond-did]",
		Example: "price 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query price(s) of the bond at a specific supply",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondTokenWithAmount := args[0]
			bondDid := args[1]

			bondCoinWithAmount, err := sdk.ParseCoinNormalized(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.CustomPrice(
				context.Background(),
				&types.QueryCustomPriceRequest{BondDid: bondDid, BondAmount: bondCoinWithAmount.Amount.String()},
			)
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

func GetCmdBuyPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "buy-price [bond-token-with-amount] [bond-did]",
		Example: "buy-price 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query price(s) of buying an amount of tokens of the bond",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondTokenWithAmount := args[0]
			bondDid := args[1]

			bondCoinWithAmount, err := sdk.ParseCoinNormalized(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.BuyPrice(
				context.Background(),
				&types.QueryBuyPriceRequest{BondDid: bondDid, BondAmount: bondCoinWithAmount.Amount.String()},
			)
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

func GetCmdSellReturn() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell-return [bond-token-with-amount] [bond-did]",
		Example: "sell-return 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query return(s) on selling an amount of tokens of the bond",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondTokenWithAmount := args[0]
			bondDid := args[1]

			bondCoinWithAmount, err := sdk.ParseCoinNormalized(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SellReturn(
				context.Background(),
				&types.QuerySellReturnRequest{BondDid: bondDid, BondAmount: bondCoinWithAmount.Amount.String()},
			)
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

func GetCmdSwapReturn() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "swap-return [bond-did] [from-token-with-amount] [to-token]",
		Example: "swap-return U7GK8p8rVhJMKhBVRCJJ8c 10res1 res2",
		Short:   "Query return(s) on swapping an amount of tokens to another token",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]
			fromTokenWithAmount := args[1]
			toToken := args[2]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SwapReturn(
				context.Background(),
				&types.QuerySwapReturnRequest{BondDid: bondDid, FromTokenWithAmount: fromTokenWithAmount, ToToken: toToken},
			)
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

func GetCmdAlphaMaximums() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alpha-maximums [bond-did]",
		Example: "alpha-maximums U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query alpha maximums for an augmented bonding curve",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AlphaMaximums(context.Background(), &types.QueryAlphaMaximumsRequest{BondDid: bondDid})
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

func GetParamsRequestHandler() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
