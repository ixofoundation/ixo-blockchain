package types

import (
	"encoding/json"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"regexp"
)

var (
	ValidDid      = regexp.MustCompile(`^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}([/][a-zA-Z0-9:]+|)$`)
	ValidPubKey   = regexp.MustCompile(`^[123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ]{44}$`)
	IsValidDid    = ValidDid.MatchString
	IsValidPubKey = ValidPubKey.MatchString
	// https://sovrin-foundation.github.io/sovrin/spec/did-method-spec-template.html
	// IsValidDid adapted from the above link but assumes no sub-namespaces
	// TODO: ValidDid needs to be updated once we no longer want to be able
	//   to consider project accounts as DIDs (especially in treasury module),
	//   possibly should just be `^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}$`.
)

var _ exported.DidDoc = (*BaseDidDoc)(nil)

type BaseDidDoc struct {
	Did         exported.Did             `json:"did" yaml:"did"`
	PubKey      string                   `json:"pubKey" yaml:"pubKey"`
	Credentials []exported.DidCredential `json:"credentials" yaml:"credentials"`
}

func NewBaseDidDoc(did exported.Did, pubKey string) BaseDidDoc {
	return BaseDidDoc{
		Did:         did,
		PubKey:      pubKey,
		Credentials: []exported.DidCredential{},
	}
}

func (dd BaseDidDoc) GetDid() exported.Did                     { return dd.Did }
func (dd BaseDidDoc) GetPubKey() string                        { return dd.PubKey }
func (dd BaseDidDoc) GetCredentials() []exported.DidCredential { return dd.Credentials }

func (dd BaseDidDoc) SetDid(did exported.Did) error {
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

func (dd BaseDidDoc) Address() sdk.AccAddress {
	return exported.VerifyKeyToAddr(dd.GetPubKey())
}

func (dd *BaseDidDoc) AddCredential(cred exported.DidCredential) {
	if dd.Credentials == nil {
		dd.Credentials = make([]exported.DidCredential, 0)
	}

	dd.Credentials = append(dd.Credentials, cred)
}

type Credential struct{}

func fromJsonString(jsonIxoDid string) (exported.IxoDid, error) {
	var did exported.IxoDid
	err := json.Unmarshal([]byte(jsonIxoDid), &did)
	if err != nil {
		err := fmt.Errorf("could not unmarshal did into struct Error: %s", err.Error())
		return exported.IxoDid{}, err
	}

	return did, nil
}

func UnmarshalIxoDid(jsonIxoDid string) (exported.IxoDid, error) {
	return fromJsonString(jsonIxoDid)
}
