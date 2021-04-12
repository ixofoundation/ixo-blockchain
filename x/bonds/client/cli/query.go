package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// TODO (Stef) Remove storeKey

func GetQueryCmd(storeKey string) *cobra.Command {
	bondsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bonds querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bondsQueryCmd.AddCommand(
		GetCmdBondsList(storeKey),
		GetCmdBondsListDetailed(storeKey),
		GetCmdBond(storeKey),
		GetCmdBatch(storeKey),
		GetCmdLastBatch(storeKey),
		GetCmdCurrentPrice(storeKey),
		GetCmdCurrentReserve(storeKey),
		GetCmdCustomPrice(storeKey),
		GetCmdBuyPrice(storeKey),
		GetCmdSellReturn(storeKey),
		GetCmdSwapReturn(storeKey),
		GetCmdAlphaMaximums(storeKey),
		GetParamsRequestHandler(),
	)

	return bondsQueryCmd
}

func GetCmdBondsList(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonds-list",
		Short: "List of all bonds",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			// TODO (Stef) Should we use types.NewQueryClient(clientCtx) here? (Look at staking and gov query.go)

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute,
					keeper.QueryBonds), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBonds
			//cdc.MustUnmarshalJSON(res, &out)
			// TODO (Stef) Or use below
			// clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			// return clientCtx.PrintObjectLegacy(out)
			clientCtx.JSONMarshaler.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintProto(&out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBondsListDetailed(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonds-list-detailed",
		Short: "List of all bonds with information about current state",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute,
					keeper.QueryBondsDetailed), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBondsDetailed
			//cdc.MustUnmarshalJSON(res, &out)
			clientCtx.JSONMarshaler.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintProto(&out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBond(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond [bond-did]",
		Short: "Query info of a bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryBond, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Bond
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBatch(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch [bond-did]",
		Short: "Query info of a bond's current batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryBatch, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Batch
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdLastBatch(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "last-batch [bond-did]",
		Short: "Query info of a bond's last batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryLastBatch, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Batch
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetCmdCurrentPrice(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current-price [bond-did]",
		Short: "Query current price(s) of the bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryCurrentPrice, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.DecCoins
			// TODO (Stef) Replace LegacyAmino
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &out) //supposed to use: clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdCurrentReserve(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "current-reserve [bond-did]",
		Example: "current-reserve U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query current balance(s) of the reserve pool",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryCurrentReserve, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.Coins
			// TODO (Stef) Replace LegacyAmino
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &out) //supposed to use: clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdCustomPrice(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "price [bond-token-with-amount] [bond-did]",
		Example: "price 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query price(s) of the bond at a specific supply",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
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

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute,
					keeper.QueryCustomPrice, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.DecCoins
			// TODO (Stef) Replace LegacyAmino
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.LegacyAmino.UnmarshalJSON(res, &out) //supposed to use: clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBuyPrice(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "buy-price [bond-token-with-amount] [bond-did]",
		Example: "buy-price 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query price(s) of buying an amount of tokens of the bond",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
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

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute,
					keeper.QueryBuyPrice, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBuyPrice
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdSellReturn(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell-return [bond-token-with-amount] [bond-did]",
		Example: "sell-return 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query return(s) on selling an amount of tokens of the bond",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
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

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute,
					keeper.QuerySellReturn, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QuerySellReturn
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdSwapReturn(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "swap-return [bond-did] [from-token-with-amount] [to-token]",
		Example: "swap-return U7GK8p8rVhJMKhBVRCJJ8c 10res1 res2",
		Short:   "Query return(s) on swapping an amount of tokens to another token",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]
			fromTokenWithAmount := args[1]
			toToken := args[2]

			fromCoinWithAmount, err := sdk.ParseCoinNormalized(fromTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s/%s/%s", queryRoute,
					keeper.QuerySwapReturn, bondDid, fromCoinWithAmount.Denom,
					fromCoinWithAmount.Amount.String(), toToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QuerySwapReturn
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAlphaMaximums(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:     "alpha-maximums [bond-did]",
		Example: "alpha-maximums U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query alpha maximums for an augmented bonding curve",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bondDid := args[0]

			res, _, err := clientCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryAlphaMaximums, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryAlphaMaximums
			//err = cdc.UnmarshalJSON(res, &out)
			err = clientCtx.JSONMarshaler.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			// TODO (Stef) Replace LegacyAmino
			output, err := clientCtx.LegacyAmino.MarshalJSONIndent(out, "", "  ") //cdc.MarshalJSONIndent(out, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetParamsRequestHandler() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query params",
		RunE: func(cmd *cobra.Command, args []string) error {
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s",
				types.QuerierRoute, keeper.QueryParams), nil)
			if err != nil {
				return err
			}

			var params types.Params
			if err := clientCtx.JSONMarshaler.UnmarshalJSON(bz, &params) /*cdc.UnmarshalJSON(bz, &params)*/; err != nil {
				return err
			}

			fmt.Println(string(bz))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
