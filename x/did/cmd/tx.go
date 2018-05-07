package commands

import (
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

			/*	didDoc := did.BaseDidDoc{
					Did:    []byte(args[0]),
					PubKey: args[1],
				}

							bzt, errt := cdc.MarshalBinary(didDoc)
							if errt != nil {
								panic(errt)
							}
							r, n, err3 := bytes.NewBuffer(bzt), new(int), new(error)
							didI := oldwire.ReadBinary(struct{ ixo.DidDoc }{}, r, len(bzt), n, err3)
							if *err3 != nil {
								panic(*err3)
							}

							didDoc2 := didI.(struct{ ixo.DidDoc }).DidDoc
							fmt.Println(didDoc2)
			*/
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
func GetDidDocCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getDidDoc did",
		Short: "Get a new DidDoc for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide the did")
			}

			ctx := context.NewCoreContextFromViper()

			// create the message
			msg := did.NewGetDidMsg(args[0])

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
