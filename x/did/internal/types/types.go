package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

var _ ixo.DidDoc = (*BaseDidDoc)(nil)

type BaseDidDoc struct {
	Did         ixo.Did         `json:"did" yaml:"did"`
	PubKey      string          `json:"pubKey" yaml:"pubKey"`
	Credentials []DidCredential `json:"credentials" yaml:"credentials"`
}

type DidCredential struct {
	CredType []string `json:"type" yaml:"type"`
	Issuer   ixo.Did  `json:"issuer" yaml:"issuer"`
	Issued   string   `json:"issued" yaml:"issued"`
	Claim    Claim    `json:"claim" yaml:"claim"`
}

type Claim struct {
	Id           ixo.Did `json:"id" yaml:"id"`
	KYCValidated bool    `json:"KYCValidated" yaml:"KYCValidated"`
}

type Credential struct{}

func (dd BaseDidDoc) GetDid() ixo.Did                 { return dd.Did }
func (dd BaseDidDoc) GetPubKey() string               { return dd.PubKey }
func (dd BaseDidDoc) GetCredentials() []DidCredential { return dd.Credentials }

func InitDidDoc(did ixo.Did, pubKey string) BaseDidDoc {
	return BaseDidDoc{
		Did:         did,
		PubKey:      pubKey,
		Credentials: nil,
	}
}

func (dd BaseDidDoc) SetDid(did ixo.Did) error {
	if len(dd.Did) != 0 {
		return errors.New("cannot override BaseDidDoc did")
	}

	dd.Did = did

	return nil
}

func (dd BaseDidDoc) SetPubKey(pubKey string) error {
	if len(dd.PubKey) != 0 {
		return errors.New("cannot override BaseDidDoc pubKey")
	}

	dd.PubKey = pubKey

	return nil
}

func (dd *BaseDidDoc) AddCredential(cred DidCredential) {
	if dd.Credentials == nil {
		dd.Credentials = make([]DidCredential, 0)
	}

	dd.Credentials = append(dd.Credentials, cred)
}

type DidMsg interface {
	IsNewDid() bool
}

func DidToAddr(did ixo.Did) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(did)))
}
