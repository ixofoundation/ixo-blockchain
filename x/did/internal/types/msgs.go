package types

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgAddDid        = "add-did"
	TypeMsgAddCredential = "add-credential"
)

var (
	_ ixo.IxoMsg = MsgAddDid{}
	_ ixo.IxoMsg = MsgAddCredential{}
)

type MsgAddDid struct {
	Did    exported.Did `json:"did" yaml:"did"`
	PubKey string       `json:"pubKey" yaml:"pubKey"`
}

func NewMsgAddDid(did string, publicKey string) MsgAddDid {
	return MsgAddDid{
		Did:    did,
		PubKey: publicKey,
	}
}

func (msg MsgAddDid) Type() string { return TypeMsgAddDid }

func (msg MsgAddDid) Route() string { return RouterKey }

func (msg MsgAddDid) GetSignerDid() exported.Did { return msg.Did }
func (msg MsgAddDid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgAddDid) ValidateBasic() error {
	// Check that not empty
	if strings.TrimSpace(msg.Did) == "" {
		return sdkerrors.Wrap(ErrInvalidDid, "did should not be empty")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return sdkerrors.Wrap(ErrInvalidPubKey, "pubKey should not be empty")
	}

	// Check that DID and PubKey valid
	if !IsValidDid(msg.Did) {
		return sdkerrors.Wrap(ErrInvalidDid, "DID is invalid")
	} else if !IsValidPubKey(msg.PubKey) {
		return sdkerrors.Wrap(ErrInvalidPubKey, "pubKey is invalid")
	}

	// Check that DID matches the PubKey
	unprefixedDid := exported.UnprefixedDid(msg.Did)
	expectedUnprefixedDid := exported.UnprefixedDidFromPubKey(msg.PubKey)
	if unprefixedDid != expectedUnprefixedDid {
		return sdkerrors.Wrapf(ErrDidPubKeyMismatch,
			"did not deducable from pubKey; expected: %s received: %s",
			expectedUnprefixedDid, unprefixedDid)
	}

	return nil
}

func (msg MsgAddDid) GetSignBytes() []byte {
	return sdk.MustSortJSON(amino.MustMarshalJSON(msg))
}

func (msg MsgAddDid) String() string {
	return fmt.Sprintf("MsgAddDid{Did: %v, publicKey: %v}", msg.Did, msg.PubKey)
}

func (msg MsgAddDid) Reset() {
	// TODO Implement Reset() of Msg interface
	panic("implement me")
}

func (msg MsgAddDid) ProtoMessage() {
	// TODO Implement ProtoMessage() of Msg interface
	panic("implement me")
}

type MsgAddCredential struct {
	DidCredential exported.DidCredential `json:"credential" yaml:"credential"`
}

func NewMsgAddCredential(did string, credType []string, issuer string, issued string) MsgAddCredential {
	didCredential := exported.DidCredential{
		CredType: credType,
		Issuer:   issuer,
		Issued:   issued,
		Claim: exported.Claim{
			Id:           did,
			KYCValidated: true,
		},
	}

	return MsgAddCredential{
		DidCredential: didCredential,
	}
}

func (msg MsgAddCredential) Type() string  { return TypeMsgAddCredential }
func (msg MsgAddCredential) Route() string { return RouterKey }

func (msg MsgAddCredential) GetSignerDid() exported.Did { return msg.DidCredential.Issuer }
func (msg MsgAddCredential) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgAddCredential) String() string {
	return fmt.Sprintf("MsgAddCredential{Did: %v, Type: %v, Signer: %v}",
		string(msg.DidCredential.Claim.Id), msg.DidCredential.CredType, string(msg.DidCredential.Issuer))
}

func (msg MsgAddCredential) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.DidCredential.Claim.Id) == "" {
		return sdkerrors.Wrap(ErrInvalidClaimId, "claim ID should not be empty")
	} else if strings.TrimSpace(msg.DidCredential.Issuer) == "" {
		return sdkerrors.Wrap(ErrInvalidIssuer, "issuer should not be empty")
	}

	// Check that DID valid
	if !IsValidDid(msg.DidCredential.Issuer) {
		return sdkerrors.Wrap(ErrInvalidDid, "issuer DID is invalid")
	}

	return nil
}

func (msg MsgAddCredential) GetSignBytes() []byte {
	return sdk.MustSortJSON(amino.MustMarshalJSON(msg))
}

func (msg MsgAddCredential) Reset() {
	// TODO Implement Reset() of Msg interface
	panic("implement me")
}

func (msg MsgAddCredential) ProtoMessage() {
	// TODO Implement ProtoMessage() of Msg interface
	panic("implement me")
}
