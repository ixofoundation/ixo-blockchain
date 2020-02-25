package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"

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

	bondsQueryCmd.AddCommand(client.GetCommands(
		GetCmdBonds(storeKey, cdc),
		GetCmdBond(storeKey, cdc),
		GetCmdBatch(storeKey, cdc),
		GetCmdLastBatch(storeKey, cdc),
		GetCmdCurrentPrice(storeKey, cdc),
		GetCmdCurrentReserve(storeKey, cdc),
		GetCmdCustomPrice(storeKey, cdc),
		GetCmdBuyPrice(storeKey, cdc),
		GetCmdSellReturn(storeKey, cdc),
		GetCmdSwapReturn(storeKey, cdc),
	)...)

	return bondsQueryCmd
}

func GetCmdBonds(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bonds-list",
		Short: "List of all bonds",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/bonds",
					queryRoute), nil)
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

func GetCmdBond(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bond [bond-did]",
		Short: "Query info of a bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondDid := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/bond/%s",
					queryRoute, bondDid), nil)
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
				fmt.Sprintf("custom/%s/batch/%s",
					queryRoute, bondDid), nil)
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
				fmt.Sprintf("custom/%s/last_batch/%s",
					queryRoute, bondDid), nil)
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
				fmt.Sprintf("custom/%s/current_price/%s",
					queryRoute, bondDid), nil)
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
				fmt.Sprintf("custom/%s/current_reserve/%s",
					queryRoute, bondDid), nil)
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
				fmt.Sprintf("custom/%s/custom_price/%s/%s",
					queryRoute, bondDid,
					bondCoinWithAmount.Amount.String()), nil)
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
				fmt.Sprintf("custom/%s/buy_price/%s/%s",
					queryRoute, bondDid,
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
				fmt.Sprintf("custom/%s/sell_return/%s/%s",
					queryRoute, bondDid,
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
		Example: "swap-return abc 10res1 res2 U7GK8p8rVhJMKhBVRCJJ8c",
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
				fmt.Sprintf("custom/%s/swap_return/%s/%s/%s/%s",
					queryRoute, bondDid, fromCoinWithAmount.Denom,
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
