package commands

import (
	"fmt"
	//	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// take the coolness quiz transaction
func CreateProjectCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createProject projectData",
		Short: "Create a new Project for the given data",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide the project data")
			}

			ctx := context.NewCoreContextFromViper()

			// create the message
			msg := ixo.NewIxoMsg("3J56r8ZGfD6ThhwhaDv9iA", args[0])

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

/*
// sign and build the transaction from the msg
func SignWithDidBuildBroadcast(ctx sdk.Context, msg sdk.Msg, cdc *wire.Codec) (*ctypes.ResultBroadcastTxCommit, error) {
	did, err := GetDidFromStdin()
	if err != nil {
		return nil, err
	}

	secret, err := GetSecretKeyFromStdin(did)
	if err != nil {
		return nil, err
	}

	txBytes, err := SignWithDidAndBuild(ctx, did, secret, msg, cdc)
	if err != nil {
		return nil, err
	}

	return ctx.BroadcastTx(txBytes)
}

func SignWithDidAndBuild(ctx core.CoreContext, did string, secret string, msg sdk.Msg, cdc *wire.Codec) ([]byte, error) {

	// sign and build
	bz := msg.Bytes()

	sig, pubkey, err := keybase.Sign(name, passphrase, bz)
	if err != nil {
		return nil, err
	}
	sigs := []sdk.StdSignature{{
		PubKey:    pubkey,
		Signature: sig,
		Sequence:  sequence,
	}}

	// marshal bytes
	tx := sdk.NewStdTx(signMsg.Msg, signMsg.Fee, sigs)

	return cdc.MarshalBinary(tx)
}

// get did from std input
func GetDidFromStdin() (did string, err error) {
	buf := client.BufferStdin()
	prompt := fmt.Sprintf("Sovrin did to sign with:")
	return ixoClient.GetString(prompt, 10, buf)
}

// get secret from std input
func GetSecretKeyFromStdin(did string) (secretKey string, err error) {
	buf := client.BufferStdin()
	prompt := fmt.Sprintf("Sovrin secret key to sign with '%s':", did)
	return ixoClient.GetString(prompt, 32, buf)
}
*/
