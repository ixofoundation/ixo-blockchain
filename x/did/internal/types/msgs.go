package types

import (
	"encoding/json"
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AddDidMsg struct {
	DidDoc    BaseDidDoc `json:"didDoc"`
	SignBytes string     `json:"signBytes"`
}

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

var _ sdk.Msg = AddDidMsg{}

func (msg AddDidMsg) Type() string { return "did" }

func (msg AddDidMsg) Route() string { return RouterKey }

func (msg AddDidMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.DidDoc.GetDid())}
}

func (msg AddDidMsg) ValidateBasic() sdk.Error {
	if msg.DidDoc.Did == "" {
		return ErrorInvalidDid(DefaultCodeSpace, "did should not be empty")
	} else if msg.DidDoc.PubKey == "" {
		return ErrorInvalidPubKey(DefaultCodeSpace, "pubKey should not be empty")
	}
	
	for _, credential := range msg.DidDoc.Credentials {
		if credential.Issuer == "" {
			return ErrorInvalidIssuer(DefaultCodeSpace, "issuer should not be empty")
		} else if credential.Claim.Id == "" {
			return ErrorInvalidDid(DefaultCodeSpace, "claim id should not be empty")
		}
	}
	
	return nil
}

func (msg AddDidMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg AddDidMsg) String() string {
	return fmt.Sprintf("AddDidMsg{Did: %v, publicKey: %v}", string(msg.DidDoc.GetDid()), msg.DidDoc.GetPubKey())
}

func (msg AddDidMsg) IsNewDid() bool { return true }

type AddCredentialMsg struct {
	DidCredential DidCredential `json:"credential"`
}

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

var _ sdk.Msg = AddCredentialMsg{}

func (msg AddCredentialMsg) Type() string  { return "did" }
func (msg AddCredentialMsg) Route() string { return RouterKey }
func (msg AddCredentialMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.DidCredential.Issuer)}
}

func (msg AddCredentialMsg) String() string {
	return fmt.Sprintf("AddCredentialMsg{Did: %v, Type: %v, Signer: %v}",
		string(msg.DidCredential.Claim.Id), msg.DidCredential.CredType, string(msg.DidCredential.Issuer))
}

func (msg AddCredentialMsg) ValidateBasic() sdk.Error {
	if msg.DidCredential.Claim.Id == "" {
		return ErrorInvalidDid(DefaultCodeSpace, "claim id should not be nil")
	} else if msg.DidCredential.Issuer == "" {
		return ErrorInvalidIssuer(DefaultCodeSpace, "Issuer should not be nil")
	}
	
	return nil
}

func (msg AddCredentialMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	return b
}

func (msg AddCredentialMsg) IsNewDid() bool { return false }
