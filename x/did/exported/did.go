package exported

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

type Did = string

type DidDoc interface {
	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
	Address() sdk.AccAddress
}

type Secret struct {
	Seed                 string `json:"seed" yaml:"seed"`
	SignKey              string `json:"signKey" yaml:"signKey"`
	EncryptionPrivateKey string `json:"encryptionPrivateKey" yaml:"encryptionPrivateKey"`
}

func (s Secret) String() string {
	output, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", string(output))
}

type IxoDid struct {
	Did                 string `json:"did" yaml:"did"`
	VerifyKey           string `json:"verifyKey" yaml:"verifyKey"`
	EncryptionPublicKey string `json:"encryptionPublicKey" yaml:"encryptionPublicKey"`
	Secret              Secret `json:"secret" yaml:"secret"`
}

// Above IxoDid modelled after Sovrin documents
// Ref: https://www.npmjs.com/package/sovrin-did
// {
//    did: "<base58 did>",
//    verifyKey: "<base58 publicKey>",
//    publicKey: "<base58 publicKey>",
//
//    secret: {
//        seed: "<hex encoded 32-byte seed>",
//        signKey: "<base58 secretKey>",
//        privateKey: "<base58 privateKey>"
//    }
// }

func (id IxoDid) String() string {
	output, err := json.MarshalIndent(id, "", "  ")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", string(output))
}

func VerifyKeyToAddr(verifyKey string) sdk.AccAddress {
	var pubKey ed25519.PubKeyEd25519
	copy(pubKey[:], base58.Decode(verifyKey))
	return sdk.AccAddress(pubKey.Address())
}

func (id IxoDid) Address() sdk.AccAddress {
	return VerifyKeyToAddr(id.VerifyKey)
}

type Claim struct {
	Id           Did  `json:"id" yaml:"id"`
	KYCValidated bool `json:"KYCValidated" yaml:"KYCValidated"`
}

type DidCredential struct {
	CredType []string `json:"type" yaml:"type"`
	Issuer   Did      `json:"issuer" yaml:"issuer"`
	Issued   string   `json:"issued" yaml:"issued"`
	Claim    Claim    `json:"claim" yaml:"claim"`
}
