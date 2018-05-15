package ixo

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/go-crypto"
)

//_______________________________________________________________________
// Define the IxoTx

type IxoTx struct {
	Msg       sdk.Msg      `json:"payload"`
	Signature IxoSignature `json:"signature"`
}

type IxoSignature struct {
	SignatureValue crypto.Signature `json:"signatureValue"`
	Created        time.Time        `json:"created"`
	Creator        Did              `json:"creator"`
}

func NewSignature(created time.Time, did string, signature crypto.Signature) IxoSignature {
	return IxoSignature{
		SignatureValue: signature,
		Created:        created,
		Creator:        did,
	}
}

func NewIxoTx(msg sdk.Msg, signature IxoSignature) IxoTx {
	return IxoTx{
		Msg:       msg,
		Signature: signature,
	}
}

func (tx IxoTx) GetMsg() sdk.Msg { return tx.Msg }
func (tx IxoTx) GetSignatures() []sdk.StdSignature {
	//TODO: Is is necesasary to create a new one
	stdSig := sdk.StdSignature{
		Signature: tx.Signature.SignatureValue,
	}
	return []sdk.StdSignature{stdSig}
}
func (tx IxoTx) String() string {
	output, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}

// enforce the msg type at compile time
var _ sdk.Tx = IxoTx{}

// Define Did as an Address
type Did = string

type DidDoc interface {
	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
}

// Define Project
type Project = string

type ProjectDoc interface {
	SetCreatedBy(did Did) error
	GetCreatedBy() Project
	SetDid(did Did) error
	GetDid() Project
}
