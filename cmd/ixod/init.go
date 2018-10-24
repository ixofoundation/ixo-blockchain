package main

import (
	"encoding/json"
	"errors"
	"fmt"

	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"

	crypto "github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

// simple genesis tx
type IxoGenTx struct {
	Addr sdk.AccAddress `json:"addr"`
}

// Generate a genesis transaction
func IxoAppGenTx(cdc *wire.Codec, pk crypto.PubKey, genTxConfig serverconfig.GenTx) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {

	var addr sdk.AccAddress
	var secret string
	addr, secret, err = server.GenerateCoinKey()
	if err != nil {
		return
	}

	ixo.IxoAppGenEthWallet()

	var bz []byte
	ixoGenTx := IxoGenTx{
		addr,
	}
	bz, err = cdc.MarshalJSON(ixoGenTx)
	if err != nil {
		return
	}
	appGenTx = json.RawMessage(bz)

	cliPrint = json.RawMessage(fmt.Sprintf(`{
		"secret": "%s"
	}`, secret))

	validator = tmtypes.GenesisValidator{
		PubKey: pk,
		Power:  10,
	}
	return
}

// create the genesis app state
func IxoAppGenState(cdc *wire.Codec, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	if len(appGenTxs) != 1 {
		err = errors.New("must provide a single genesis transaction")
		return
	}

	var genTx IxoGenTx
	err = cdc.UnmarshalJSON(appGenTxs[0], &genTx)
	if err != nil {
		return
	}

	feesState := fees.DefaultGenesis()
	feesJson, err := json.MarshalIndent(feesState, "", "  ")

	ixoConfig := ixo.DefaultGenesisState()
	ixoConfigJson, err := json.MarshalIndent(ixoConfig, "", "  ")

	appState = json.RawMessage(fmt.Sprintf(`{
  "accounts": [{
    "address": "%s",
    "coins": [
      {
        "denom": "%s",
        "amount": "0"
      }
    ]
	}],
	"fees": %s,
	"config": %s
}`, genTx.Addr, ixo.IxoNativeToken, feesJson, ixoConfigJson))
	return
}

//___________________________________________________________________________________________

// GenerateSaveCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func GenerateSaveCoinKey(clientRoot, keyName, keyPass string, overwrite bool) (sdk.AccAddress, string, error) {

	// get the keystore from the client
	keybase, err := clkeys.GetKeyBaseFromDir(clientRoot)
	if err != nil {
		return sdk.AccAddress([]byte{}), "", err
	}

	// ensure no overwrite
	if !overwrite {
		_, err := keybase.Get(keyName)
		if err == nil {
			return sdk.AccAddress([]byte{}), "", errors.New("key already exists, overwrite is disabled")
		}
	}

	// generate a private key, with recovery phrase
	info, secret, err := keybase.CreateMnemonic(keyName, keys.English, keyPass, keys.Secp256k1)
	if err != nil {
		return sdk.AccAddress([]byte{}), "", err
	}
	addr := info.GetPubKey().Address()
	return sdk.AccAddress(addr), secret, nil
}
