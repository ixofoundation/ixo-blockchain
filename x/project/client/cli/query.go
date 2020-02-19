package cli

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func GetProjectDocCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getProjectDoc did",
		Short: "Get a new ProjectDoc for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide a did")
			}

			didAddr := args[0]
			key := ixo.Did(didAddr)

			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryProjectDoc, key), nil)
			if err != nil {
				return err
			}

			var projectDoc types.CreateProjectMsg
			err = cdc.UnmarshalJSON(res, &projectDoc)
			if err != nil {
				return err
			}

			output, err := json.MarshalIndent(projectDoc, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetProjectAccountsCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getProjectAccounts did",
		Short: "Get a Project accounts for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide a project did")
			}

			projectDid := args[0]

			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryProjectAccount, projectDid), nil)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return errors.New("Project does not exist")
			}

			var f interface{}
			err = json.Unmarshal(res, &f)
			if err != nil {
				return err
			}
			accMap := f.(map[string]interface{})

			output, err := json.MarshalIndent(accMap, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}

func GetProjectTxsCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getProjectTxs projectDid",
		Short: "Get a Project txs for a projectDid",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide a project did")
			}
			projectDid := args[0]

			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
				keeper.QueryProjectTx, projectDid), nil)
			if err != nil {
				return err
			}

			txs := []types.WithdrawalInfo{}
			if len(res) == 0 {
				return errors.New("projectTxs does not exist for a projectDid")
			} else {
				err = cdc.UnmarshalJSON(res, &txs)
				if err != nil {
					return err
				}
			}

			output, err := json.MarshalIndent(txs, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}
