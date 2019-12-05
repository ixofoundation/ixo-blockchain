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

func GetFiatPegCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "[fiatPeg-id]",
		Short: "Query fiatPeg details",
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext()

			pegHash, err := types.GetPegHashFromString(args[0])
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryStore(pegHash, fiatTypes.ModuleName)
			if err != nil {
				return err
			}

			if res == nil {
				return sdk.ErrUnknownAddress("No fiatPeg with pegHash " + args[0] +
					" was found in the state.\nAre you sure there has been a transaction involving it?")
			}

			var _fiatPeg types.FiatPeg
			err = cdc.UnmarshalBinaryBare(res, &_fiatPeg)
			if err != nil {
				return err
			}

			output, err := cdc.MarshalJSON(_fiatPeg)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(fsPegHash)
	return cmd
}
