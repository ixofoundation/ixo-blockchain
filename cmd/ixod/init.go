package main

import (
	"encoding/hex"
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
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	crypto "github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

// simple genesis tx
type IxoGenTx struct {
	Addr sdk.AccAddress `json:"addr"`
}

type EthWallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

func IxoAppGenEthWallet() {
	// Create an account
	key, err := ethCrypto.GenerateKey()
	if err != nil {
		return
	}

	// Get the address
	address := ethCrypto.PubkeyToAddress(key.PublicKey).Hex()
	// Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())

	ethWallet := &EthWallet{Address: address, PrivateKey: privateKey}
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

	var addr sdk.AccAddress
	var secret string
	addr, secret, err = server.GenerateCoinKey()
	if err != nil {
		return
	}

	var bz []byte
	ixoGenTx := IxoGenTx{addr}
	bz, err = cdc.MarshalJSON(ixoGenTx)
	if err != nil {
		return
	}
	appGenTx = json.RawMessage(bz)

	mm := map[string]string{"secret": secret}
	bz, err = cdc.MarshalJSON(mm)
	if err != nil {
		return
	}
	cliPrint = json.RawMessage(bz)

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
	IxoAppGenEthWallet()
	appState = json.RawMessage(fmt.Sprintf(`{
  "accounts": [{
    "address": "%s",
    "coins": [
      {
        "denom": "ixo-native",
        "amount": "0"
      }
    ]
  }]
}`, genTx.Addr))
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
