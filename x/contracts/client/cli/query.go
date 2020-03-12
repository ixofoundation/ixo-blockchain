package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/types"
)

func GetContractCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getContracts",
		Short: "Get all contract details",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
				keeper.QueryAllContracts), nil)

			contracts := make(map[string]string)
			err = json.Unmarshal(bz, &contracts)
			if err != nil {
				return err
			}

			output, err := json.MarshalIndent(contracts, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))

			return nil
		},
	}
}
