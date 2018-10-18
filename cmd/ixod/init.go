package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	sovrin "github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"

	crypto "github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

// simple genesis tx
type IxoGenTx struct {
	Addr      sdk.AccAddress `json:"addr"`
	ConfigDid did.BaseDidDoc `json:"configDid"`
}

func IxoAppGenEthWallet() {
	// Create an account
	ethWallet, err := ixo.CreateEthWallet()
	if err != nil {
		return
	}

	json, err := json.Marshal(ethWallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("ethWallet.json", json, 0644)

	return
}

// Generate a genesis transaction
func IxoAppGenTx(cdc *wire.Codec, pk crypto.PubKey, genTxConfig serverconfig.GenTx) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {

	mnemonic := sovrin.GenerateMnemonic()
	sov := sovrin.FromMnemonic(mnemonic)

	didDoc := did.InitDidDoc(did.PrefixDid(sov.Did), sov.VerifyKey)

	var addr sdk.AccAddress
	var secret string
	addr, secret, err = server.GenerateCoinKey()
	if err != nil {
		return
	}

	var bz []byte
	ixoGenTx := IxoGenTx{
		addr,
		didDoc,
	}
	bz, err = cdc.MarshalJSON(ixoGenTx)
	if err != nil {
		return
	}
	appGenTx = json.RawMessage(bz)

	cliPrint = json.RawMessage(fmt.Sprintf(`{
		"secret": "%s",
		"config_did": {
			"did": "%s",
			"mnemonic": "%s"
		}
	}`, secret, didDoc.Did, mnemonic))

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

	appState = json.RawMessage(fmt.Sprintf(`{
  "accounts": [{
    "address": "%s",
    "coins": [
      {
        "denom": "ixo-native",
        "amount": "0"
      }
    ]
	}],
	"fees": %s
}`, genTx.Addr, feesJson))
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
