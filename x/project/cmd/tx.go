package commands

import (
	"fmt"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/ixofoundation/ixo-cosmos/x/project"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// Add a project doc to the ledger
func AddProjectDocCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "addProjectDoc did projectData",
		Short: "Add a new ProjectDoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the did and the project data")
			}
			ctx := context.NewCoreContextFromViper()

			// create the message
			msg := project.NewAddProjectMsg(args[0], args[1])

			tx := ixo.NewIxoTx(msg)

			bz, err := cdc.MarshalBinary(tx)
			if err != nil {
				panic(err)
			}
			// Broadcast to Tendermint
			res, err := ctx.BroadcastTx(bz)
			if err != nil {
				return err
			}

			fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
			return nil
		},
	}
}

// Get a project doc from the ledger
func GetProjectDocCmd(storeName string, cdc *wire.Codec, decoder project.ProjectDocDecoder) *cobra.Command {
	return &cobra.Command{
		Use:   "getProjectDoc did",
		Short: "Get a new ProjectDoc for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide an did")
			}

			// find the key to look up the account
			didAddr := args[0]
			key := ixo.Did(didAddr)

			ctx := context.NewCoreContextFromViper()

			res, err := ctx.Query([]byte(key), storeName)
			if err != nil {
				return err
			}

			// decode the value
			projectDoc, err := decoder(res)
			if err != nil {
				return err
			}
			// print out whole account
			output, err := json.MarshalIndent(projectDoc, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))

			return nil
		},
	}
}
