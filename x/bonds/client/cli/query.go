package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	bondsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bonds querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bondsQueryCmd.AddCommand(flags.GetCommands(
		GetCmdBondsList(storeKey, cdc),
		GetCmdBondsListDetailed(storeKey, cdc),
		GetCmdBond(storeKey, cdc),
		GetCmdBatch(storeKey, cdc),
		GetCmdLastBatch(storeKey, cdc),
		GetCmdCurrentPrice(storeKey, cdc),
		GetCmdCurrentReserve(storeKey, cdc),
		GetCmdCustomPrice(storeKey, cdc),
		GetCmdBuyPrice(storeKey, cdc),
		GetCmdSellReturn(storeKey, cdc),
		GetCmdSwapReturn(storeKey, cdc),
		GetCmdAlphaMaximums(storeKey, cdc),
		GetParamsRequestHandler(cdc),
	)...)

	return bondsQueryCmd
}

func GetCmdBondsList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bonds-list",
		Short: "List of all bonds",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute,
					keeper.QueryBonds), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBonds
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBondsListDetailed(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bonds-list-detailed",
		Short: "List of all bonds with information about current state",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute,
					keeper.QueryBondsDetailed), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBondsDetailed
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBond(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bond [bond-did]",
		Short: "Query info of a bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryBond, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Bond
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

func GetCmdBatch(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "batch [bond-did]",
		Short: "Query info of a bond's current batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryBatch, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Batch
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

func GetCmdLastBatch(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "last-batch [bond-did]",
		Short: "Query info of a bond's last batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryLastBatch, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Batch
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

func GetCmdCurrentPrice(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "current-price [bond-did]",
		Short: "Query current price(s) of the bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryCurrentPrice, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.DecCoins
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

func GetCmdCurrentReserve(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "current-reserve [bond-did]",
		Example: "current-reserve U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query current balance(s) of the reserve pool",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryCurrentReserve, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.Coins
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

func GetCmdCustomPrice(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "price [bond-token-with-amount] [bond-did]",
		Example: "price 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query price(s) of the bond at a specific supply",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondTokenWithAmount := args[0]
			bondDid := args[1]

			bondCoinWithAmount, err := sdk.ParseCoin(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute,
					keeper.QueryCustomPrice, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.DecCoins
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

func GetCmdBuyPrice(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "buy-price [bond-token-with-amount] [bond-did]",
		Example: "buy-price 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query price(s) of buying an amount of tokens of the bond",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondTokenWithAmount := args[0]
			bondDid := args[1]

			bondCoinWithAmount, err := sdk.ParseCoin(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute,
					keeper.QueryBuyPrice, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBuyPrice
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

func GetCmdSellReturn(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "sell-return [bond-token-with-amount] [bond-did]",
		Example: "sell-return 10abc U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query return(s) on selling an amount of tokens of the bond",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondTokenWithAmount := args[0]
			bondDid := args[1]

			bondCoinWithAmount, err := sdk.ParseCoin(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute,
					keeper.QuerySellReturn, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QuerySellReturn
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

func GetCmdSwapReturn(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "swap-return [bond-did] [from-token-with-amount] [to-token]",
		Example: "swap-return U7GK8p8rVhJMKhBVRCJJ8c 10res1 res2",
		Short:   "Query return(s) on swapping an amount of tokens to another token",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]
			fromTokenWithAmount := args[1]
			toToken := args[2]

			fromCoinWithAmount, err := sdk.ParseCoin(fromTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s/%s/%s/%s", queryRoute,
					keeper.QuerySwapReturn, bondDid, fromCoinWithAmount.Denom,
					fromCoinWithAmount.Amount.String(), toToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QuerySwapReturn
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

func GetCmdAlphaMaximums(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "alpha-maximums [bond-did]",
		Example: "alpha-maximums U7GK8p8rVhJMKhBVRCJJ8c",
		Short:   "Query alpha maximums for an augmented bonding curve",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute,
					keeper.QueryAlphaMaximums, bondDid), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryAlphaMaximums
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

func GetParamsRequestHandler(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query params",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s",
				types.QuerierRoute, keeper.QueryParams), nil)
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
