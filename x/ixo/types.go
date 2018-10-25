package ixo

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var IxoDecimals = sdk.NewRat(100000000, 1)

const IxoNativeToken = "IXO"

//_______________________________________________________________________
// Define the IxoTx

type GenesisState struct {
	FoundationWallet               string `json:"foundationWallet"`
	AuthContractAddress            string `json:"authContractAddress"`
	IxoTokenContractAddress        string `json:"ixoTokenContractAddress"`
	ProjectRegistryContractAddress string `json:"projectRegistryContractAddress"`
	ProjectWalletAuthoriserAddress string `json:"projectWalletAuthoriserAddress"`
}

func DefaultGenesisState() GenesisState {

	return GenesisState{
		FoundationWallet:               "Enter ETH wallet address to accumulate foundations tokens",
		AuthContractAddress:            "Enter ETH auth contract address",
		IxoTokenContractAddress:        "Enter ETH Ixo Token contract address",
		ProjectRegistryContractAddress: "Enter ETH project registry contract address",
		ProjectWalletAuthoriserAddress: "Enter ETH project wallet authoriser contract address",
	}
}

type IxoTx struct {
	Msgs       []sdk.Msg      `json:"payload"`
	Signatures []IxoSignature `json:"signatures"`
}

type IxoSignature struct {
	SignatureValue [64]byte  `json:"signatureValue"`
	Created        time.Time `json:"created"`
}

func NewSignature(created time.Time, signature [64]byte) IxoSignature {
	return IxoSignature{
		SignatureValue: signature,
		Created:        created,
	}
}

func NewIxoTx(msgs []sdk.Msg, sigs []IxoSignature) IxoTx {
	return IxoTx{
		Msgs:       msgs,
		Signatures: sigs,
	}
}

func NewIxoTxSingleMsg(msg sdk.Msg, signature IxoSignature) IxoTx {
	sigs := make([]IxoSignature, 0)
	sigs = append(sigs, signature)

	msgs := make([]sdk.Msg, 0)
	msgs = append(msgs, msg)

	return IxoTx{
		Msgs:       msgs,
		Signatures: sigs,
	}
}

func (tx IxoTx) GetMsgs() []sdk.Msg { return tx.Msgs }

//nolint
func (tx IxoTx) GetMemo() string { return "" }

func (tx IxoTx) GetSignatures() []IxoSignature {
	return tx.Signatures
}

func (tx IxoTx) String() string {
	output, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}

// FeePayer returns the address responsible for paying the fees
// for the transactions. It's the first address returned by msg.GetSigners().
// If GetSigners() is empty, this panics.
func FeePayer(tx sdk.Tx) sdk.AccAddress {
	return tx.GetMsgs()[0].GetSigners()[0]
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

// Ethereum
type EthWallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

// MESSAGES ************************

//Define AddEthWalletDoc
type AddEthWalletDoc struct {
	Id            string `json:"id"`
	WalletAddress string `json:"walletAddress"`
}

type AddEthWalletMsg struct {
	SignBytes string          `json:"signBytes"`
	SignerDid Did             `json:"signerDid"`
	Data      AddEthWalletDoc `json:"data"`
}

// New Ixo message
func NewAddEthWalletMsg(id string, wallet string) AddEthWalletMsg {
	addEthWalletDoc := AddEthWalletDoc{
		Id:            id,
		WalletAddress: wallet,
	}
	return AddEthWalletMsg{
		Data: addEthWalletDoc,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = AddEthWalletMsg{}

// nolint
func (msg AddEthWalletMsg) Type() string                            { return "ixo" }
func (msg AddEthWalletMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg AddEthWalletMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.SignerDid)}
}
func (msg AddEthWalletMsg) String() string {
	return fmt.Sprintf("AddEthWalletMsg{Wallet: %v}", string(msg.Data.WalletAddress))
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg AddEthWalletMsg) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg AddEthWalletMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
