package cli

import (
	"encoding/json"
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	
	"github.com/ixofoundation/ixo-cosmos/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/fees/internal/types"
)

func GetFeesRequestHandler(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getFees",
		Short: "query fees",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().
				WithCodec(cdc)
			
			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
				keeper.QueryFees), nil)
			if err != nil {
				return err
			}
			
			fees := make(map[string]int64)
			err = cliCtx.Codec.UnmarshalJSON(bz, &fees)
			if err != nil {
				return err
			}
			
			bz, err = json.MarshalIndent(fees, "", "  ")
			if err != nil {
				return err
			}
			
			fmt.Println(string(bz))
			return nil
		},
	}
}
