package cli

import (
	"encoding/json"
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/types"
)

func GetCmdBondDoc(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-bond-doc [did]",
		Short: "Query BondDoc for a DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			didAddr := args[0]
			key := did.Did(didAddr)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryBondDoc, key), nil)
			if err != nil {
				return err
			}

			var bondDoc types.MsgCreateBond
			err = cdc.UnmarshalJSON(res, &bondDoc)
			if err != nil {
				return err
			}

			output, err := json.MarshalIndent(bondDoc, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}
