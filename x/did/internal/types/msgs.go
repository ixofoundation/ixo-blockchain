package types

import (
	"encoding/json"
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ ixo.IxoMsg = MsgAddDid{}
	_ ixo.IxoMsg = MsgAddCredential{}
)

type MsgAddDid struct {
	DidDoc BaseDidDoc `json:"didDoc" yaml:"didDoc"`
}

func NewMsgAddDid(did string, publicKey string) MsgAddDid {
	didDoc := BaseDidDoc{
		Did:         did,
		PubKey:      publicKey,
		Credentials: nil,
	}

	return MsgAddDid{
		DidDoc: didDoc,
	}
}

func (msg MsgAddDid) Type() string { return "did" }

func (msg MsgAddDid) Route() string { return RouterKey }

func (msg MsgAddDid) GetSignerDid() ixo.Did {
	return msg.DidDoc.Did
}

func (msg MsgAddDid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgAddDid) ValidateBasic() sdk.Error {
	// Check that not empty
	if strings.TrimSpace(msg.DidDoc.Did) == "" {
		return ErrorInvalidDid(DefaultCodespace, "did should not be empty")
	} else if strings.TrimSpace(msg.DidDoc.PubKey) == "" {
		return ErrorInvalidPubKey(DefaultCodespace, "pubKey should not be empty")
	}

	// Check DidDoc credentials for empty fields
	for _, cred := range msg.DidDoc.Credentials {
		if strings.TrimSpace(cred.Issuer) == "" {
			return ErrorInvalidIssuer(DefaultCodespace, "issuer should not be empty")
		} else if strings.TrimSpace(cred.Claim.Id) == "" {
			return ErrorInvalidDid(DefaultCodespace, "claim id should not be empty")
		}
	}

	// Check that DID valid
	if !ixo.IsValidDid(msg.DidDoc.Did) {
		return ErrorInvalidDid(DefaultCodespace, "did is invalid")
	}

	return nil
}

func (msg MsgAddDid) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgAddDid) String() string {
	return fmt.Sprintf("MsgAddDid{Did: %v, publicKey: %v}", string(msg.DidDoc.GetDid()), msg.DidDoc.GetPubKey())
}

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

func (msg MsgAddCredential) Type() string  { return "did" }
func (msg MsgAddCredential) Route() string { return RouterKey }

func (msg MsgAddCredential) GetSignerDid() ixo.Did {
	return msg.DidCredential.Issuer
}

func (msg MsgAddCredential) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgAddCredential) String() string {
	return fmt.Sprintf("MsgAddCredential{Did: %v, Type: %v, Signer: %v}",
		string(msg.DidCredential.Claim.Id), msg.DidCredential.CredType, string(msg.DidCredential.Issuer))
}

func (msg MsgAddCredential) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.DidCredential.Claim.Id) == "" {
		return ErrorInvalidDid(DefaultCodespace, "claim id should not be empty")
	} else if strings.TrimSpace(msg.DidCredential.Issuer) == "" {
		return ErrorInvalidIssuer(DefaultCodespace, "issuer should not be empty")
	}

	// Check that DID valid
	if !ixo.IsValidDid(msg.DidCredential.Issuer) {
		return ErrorInvalidDid(DefaultCodespace, "issuer did is invalid")
	}

	return nil
}

func (msg MsgAddCredential) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}
