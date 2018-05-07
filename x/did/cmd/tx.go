package commands

import (
	"encoding/json"
	"fmt"
	//	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// Add a did doc to the ledger
func AddDidDocCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "addDidDoc did pubKey",
		Short: "Add a new DidDoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the did and the publicKey")
			}

			ctx := context.NewCoreContextFromViper()

			// create the message
			msg := did.NewAddDidMsg(args[0], args[1])

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

// Get a did doc to the ledger
func GetDidDocCmd(storeName string, cdc *wire.Codec, decoder did.DidDocDecoder) *cobra.Command {
	return &cobra.Command{
		Use:   "getDidDoc did",
		Short: "Get a new DidDoc for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide an did")
			}

			// find the key to look up the account
			didAddr := args[0]
			key := ixo.Did([]byte(didAddr))

			ctx := context.NewCoreContextFromViper()

			res, err := ctx.Query(key, storeName)
			if err != nil {
				return err
			}

			// decode the value
			didDoc, err := decoder(res)
			if err != nil {
				return err
			}
			// print out whole account
			output, err := json.MarshalIndent(didDoc, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))

			return nil
		},
	}
}
