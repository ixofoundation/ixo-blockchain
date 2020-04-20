package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgAddDid struct {
	DidDoc    BaseDidDoc `json:"didDoc" yaml:"didDoc"`
	SignBytes string     `json:"signBytes" yaml:"signBytes"`
}

func NewMsgAddDid(did string, publicKey string) MsgAddDid {
	didDoc := BaseDidDoc{
		Did:         did,
		PubKey:      publicKey,
		Credentials: make([]DidCredential, 0),
	}

	return MsgAddDid{
		DidDoc: didDoc,
	}
}

var _ sdk.Msg = MsgAddDid{}

func (msg MsgAddDid) Type() string { return "did" }

func (msg MsgAddDid) Route() string { return RouterKey }

func (msg MsgAddDid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.DidDoc.GetDid())}
}

func (msg MsgAddDid) ValidateBasic() sdk.Error {
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

func (msg MsgAddDid) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg MsgAddDid) String() string {
	return fmt.Sprintf("MsgAddDid{Did: %v, publicKey: %v}", string(msg.DidDoc.GetDid()), msg.DidDoc.GetPubKey())
}

func (msg MsgAddDid) IsNewDid() bool { return true }

type MsgAddCredential struct {
	DidCredential DidCredential `json:"credential" yaml:"credential"`
}

func NewMsgAddCredential(did string, credType []string, issuer string, issued string) MsgAddCredential {
	didCredential := DidCredential{
		CredType: credType,
		Issuer:   issuer,
		Issued:   issued,
		Claim: Claim{
			Id:           did,
			KYCValidated: true,
		},
	}

	return MsgAddCredential{
		DidCredential: didCredential,
	}
}

var _ sdk.Msg = MsgAddCredential{}

func (msg MsgAddCredential) Type() string  { return "did" }
func (msg MsgAddCredential) Route() string { return RouterKey }
func (msg MsgAddCredential) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.DidCredential.Issuer)}
}

func (msg MsgAddCredential) String() string {
	return fmt.Sprintf("MsgAddCredential{Did: %v, Type: %v, Signer: %v}",
		string(msg.DidCredential.Claim.Id), msg.DidCredential.CredType, string(msg.DidCredential.Issuer))
}

func (msg MsgAddCredential) ValidateBasic() sdk.Error {
	if msg.DidCredential.Claim.Id == "" {
		return ErrorInvalidDid(DefaultCodeSpace, "claim id should not be nil")
	} else if msg.DidCredential.Issuer == "" {
		return ErrorInvalidIssuer(DefaultCodeSpace, "Issuer should not be nil")
	}

	return nil
}

func (msg MsgAddCredential) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return b
}

func (msg MsgAddCredential) IsNewDid() bool { return false }
