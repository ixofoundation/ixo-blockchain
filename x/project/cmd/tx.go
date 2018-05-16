package commands

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"

	ixo "github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project"

	base58 "github.com/btcsuite/btcutil/base58"
)

// Add a project doc to the ledger
func AddProjectDocCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "addProjectDoc did projectData",
		Short: "Add a new ProjectDoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the project data and the projects private key")
			}

			projectDoc := project.ProjectDoc{}
			err := json.Unmarshal([]byte(args[0]), &projectDoc)
			if err != nil {
				panic(err)
			}

			sovrinDid := ixo.SovrinDid{}
			sovrinErr := json.Unmarshal([]byte(args[1]), &sovrinDid)
			if sovrinErr != nil {
				panic(sovrinErr)
			}

			// create the message
			msg := project.NewAddProjectMsg(projectDoc, sovrinDid)

			// Force the length to 64
			privKey := [64]byte{}
			copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
			copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

			//Create the Signature
			signature := ixo.SignIxoMessage(msg, sovrinDid.Did, privKey)

			tx := ixo.NewIxoTx(msg, signature)

			fmt.Println("*******TRANSACTION******* \n", tx.String())

			bz, err := cdc.MarshalJSON(tx)
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
			ctx := context.NewCoreContextFromViper()

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide an did")
			}

			// find the key to look up the account
			didAddr := args[0]
			key := ixo.Did(didAddr)

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
