package exported

import (
	"encoding/json"
	"fmt"
)

type Did = string

type DidDoc interface {
	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
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

func (id IxoDid) String() string {
	output, err := json.MarshalIndent(id, "", "  ")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", string(output))
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
