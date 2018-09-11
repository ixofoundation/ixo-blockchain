package did

import (
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

//DOC SETUP
// enforce the DidDoc type at compile time
var _ ixo.DidDoc = (*BaseDidDoc)(nil)

type BaseDidDoc struct {
	Did         ixo.Did         `json:"did"`
	PubKey      string          `json:"pubKey"` // May be nil, if not known.
	Credentials []DidCredential `json:"credentials"`
}

type DidCredential struct {
	CredType []string `json:"type"`
	Issuer   ixo.Did  `json:"issuer"`
	Issued   string   `json:"issued"`
	Claim    Claim    `json:"claim"`
}

type Claim struct {
	Id           ixo.Did `json:"id"`
	KYCValidated bool    `json:"KYCValidated"`
}

type Credential struct {
}

//GETTERS
func (dd BaseDidDoc) GetDid() ixo.Did                 { return dd.Did }
func (dd BaseDidDoc) GetPubKey() string               { return dd.PubKey }
func (dd BaseDidDoc) GetCredentials() []DidCredential { return dd.Credentials }

//SETTERS
func (dd *BaseDidDoc) Init(did ixo.Did, pubKey string) *BaseDidDoc {
	dd.SetDid(did)
	dd.SetPubKey(pubKey)
	dd.Credentials = make([]DidCredential, 0)
	return dd
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

//DECODERS
type DidDocDecoder func(didDocBytes []byte) (ixo.DidDoc, error)

func GetDidDocDecoder(cdc *wire.Codec) DidDocDecoder {
	return func(didDocBytes []byte) (res ixo.DidDoc, err error) {
		if len(didDocBytes) == 0 {
			return nil, sdk.ErrTxDecode("didDocBytes are empty")
		}
		didDoc := BaseDidDoc{}
		err = cdc.UnmarshalBinary(didDocBytes, &didDoc)
		if err != nil {
			panic(err)
		}
		return didDoc, err
	}
}

//**************************************************************************************

type DidMsg interface {
	IsNewDid() bool
}

//ADD DIDDOC
type AddDidMsg struct {
	DidDoc BaseDidDoc `json:"didDoc"`
}

// New Ixo message
func NewAddDidMsg(did string, publicKey string) AddDidMsg {
	didDoc := BaseDidDoc{
		Did:         did,
		PubKey:      publicKey,
		Credentials: make([]DidCredential, 0),
	}
	return AddDidMsg{
		DidDoc: didDoc,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = AddDidMsg{}

// nolint
func (msg AddDidMsg) Type() string                            { return "did" }
func (msg AddDidMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg AddDidMsg) GetSigners() []sdk.Address               { return []sdk.Address{[]byte(msg.DidDoc.GetDid())} }
func (msg AddDidMsg) String() string {
	return fmt.Sprintf("AddDidMsg{Did: %v, publicKey: %v}", string(msg.DidDoc.GetDid()), msg.DidDoc.GetPubKey())
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg AddDidMsg) ValidateBasic() sdk.Error {
	// Check that the Did and the PublicKey correspond
	// NOTE: this will fail if the key has been rotated
	/*	pk := msg.DidDoc.GetPubKey()
		decodedPk := base58.Decode(pk)
		decodedDid := decodedPk[0:16]
		calculatedDid := base58.Encode(decodedDid)

		if calculatedDid != msg.DidDoc.GetDid() {
			return sdk.ErrInvalidPubKey("Did does not match publicKey")
		}
	*/
	return nil
}

// Get the bytes for the message signer to sign on
func (msg AddDidMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg AddDidMsg) IsNewDid() bool { return true }

type AddCredentialMsg struct {
	DidCredential DidCredential `json:"credential"`
}

// New Ixo message
func NewAddCredentialMsg(did string, credType []string, issuer string, issued string) AddCredentialMsg {
	didCredential := DidCredential{
		CredType: credType,
		Issuer:   issuer,
		Issued:   issued,
		Claim: Claim{
			Id:           did,
			KYCValidated: true,
		},
	}
	return AddCredentialMsg{
		DidCredential: didCredential,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = AddCredentialMsg{}

// nolint
func (msg AddCredentialMsg) Type() string                            { return "did" }
func (msg AddCredentialMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg AddCredentialMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.DidCredential.Issuer)}
}
func (msg AddCredentialMsg) String() string {
	return fmt.Sprintf("AddCredentialMsg{Did: %v, Type: %v, Signer: %v}", string(msg.DidCredential.Claim.Id), msg.DidCredential.CredType, string(msg.DidCredential.Issuer))
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg AddCredentialMsg) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg AddCredentialMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg AddCredentialMsg) IsNewDid() bool { return false }
