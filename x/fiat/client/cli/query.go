package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-cosmos/types"
	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

func GetFiatAccountCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fiatAccount",
		Short: "Query fiat account details",
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext()

			bech32addr := args[0]

			addr, err := sdk.AccAddressFromBech32(bech32addr)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryStore(addr.Bytes(), fiatTypes.ModuleName)
			if err != nil {
				return err
			}

			if res == nil {
				return sdk.ErrUnknownAddress("No fiatAccount with address " + bech32addr +
					" was found in the state.\nAre you sure there has been a transaction involving it?")
			}

			var fiatAccount types.FiatAccount
			err = cdc.UnmarshalBinaryBare(res, &fiatAccount)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSON(fiatAccount)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(fsAddress)
	return cmd
}
