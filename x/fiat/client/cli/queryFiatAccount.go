package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/spf13/cobra"

	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

func GetFiatAccountCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fiatAccount",
		Short: "Query fiat account details",
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext()

			bech32addr := args[0]

			if bech32addr == "" {
				return sdk.ErrInvalidAddress("address is empty.")
			}

			addr, err := sdk.AccAddressFromBech32(bech32addr)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryStore(fiatTypes.FiatAccountStoreKey(addr), fiatTypes.ModuleName)
			if err != nil {
				return err
			}
			if res == nil {
				return sdk.ErrUnknownAddress("No fiatAccount with address " + bech32addr + " was found in the state.\n Are you sure there has been a transaction involving it?")
			}

			var fiatAccount types.BaseFiatAccount
			err = cdc.UnmarshalBinaryLengthPrefixed(res, &fiatAccount)
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
