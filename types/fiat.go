package types

import (
	"bytes"
	"encoding/hex"
	"errors"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/common"
)

// PegHash : reference address of asset peg
type PegHash = common.HexBytes

// FiatPeg : peg issued against each fiat transaction
type FiatPeg interface {
	GetPegHash() PegHash
	SetPegHash(PegHash) error

	GetTransactionID() string
	SetTransactionID(string) error

	GetTransactionAmount() int64
	SetTransactionAmount(int64) error

	GetRedeemedAmount() int64
	SetRedeemedAmount(int64)

	GetOwners() []Owner
	SetOwners([]Owner) error
	SearchOwner(sdk.AccAddress) (Owner, error)
}

// Owner : partial or full owner of a transaction
type Owner struct {
	OwnerAddress sdk.AccAddress `json:"ownerAddress"`
	Amount       int64          `json:"amount"`
}

// BaseFiatPeg : fiat peg basic implementation
type BaseFiatPeg struct {
	PegHash           PegHash `json:"pegHash" `
	TransactionID     string  `json:"transactionID" valid:"required~TxID is mandatory,matches(^[A-Z0-9]+$)~Invalid TransactionId,length(2|40)~TransactionId length between 2-40"`
	TransactionAmount int64   `json:"transactionAmount" valid:"required~TransactionAmount is mandatory,matches(^[1-9]{1}[0-9]*$)~Invalid TransactionAmount"`
	RedeemedAmount    int64   `json:"redeemedAmount"`
	Owners            []Owner `json:"owners"`
}

var _ FiatPeg = (*BaseFiatPeg)(nil)

// GetPegHash : getter
func (baseFiatPeg BaseFiatPeg) GetPegHash() PegHash { return baseFiatPeg.PegHash }

// SetPegHash : setter
func (baseFiatPeg *BaseFiatPeg) SetPegHash(pegHash PegHash) error {
	baseFiatPeg.PegHash = pegHash
	return nil
}

// GetTransactionID : getter
func (baseFiatPeg BaseFiatPeg) GetTransactionID() string { return baseFiatPeg.TransactionID }

// SetTransactionID : setter
func (baseFiatPeg *BaseFiatPeg) SetTransactionID(transactionID string) error {
	baseFiatPeg.TransactionID = transactionID
	return nil
}

// GetTransactionAmount : getter
func (baseFiatPeg BaseFiatPeg) GetTransactionAmount() int64 { return baseFiatPeg.TransactionAmount }

// SetTransactionAmount : setter
func (baseFiatPeg *BaseFiatPeg) SetTransactionAmount(transactionAmount int64) error {
	baseFiatPeg.TransactionAmount = transactionAmount
	return nil
}

// GetRedeemedAmount : getter
func (baseFiatPeg BaseFiatPeg) GetRedeemedAmount() int64 {
	return baseFiatPeg.RedeemedAmount
}

// SetRedeemedAmount : setter
func (baseFiatPeg *BaseFiatPeg) SetRedeemedAmount(redeemedAmount int64) {
	baseFiatPeg.RedeemedAmount = redeemedAmount
}

// GetOwners : getter
func (baseFiatPeg BaseFiatPeg) GetOwners() []Owner { return baseFiatPeg.Owners }

// SetOwners : setter
func (baseFiatPeg *BaseFiatPeg) SetOwners(owners []Owner) error {
	baseFiatPeg.Owners = owners
	return nil
}

// SearchOwner : Find owner for given address in FiatPeg
func (baseFiatPeg *BaseFiatPeg) SearchOwner(ownerAddress sdk.AccAddress) (Owner, error) {
	owners := baseFiatPeg.GetOwners()
	index := sort.Search(len(owners), func(i int) bool {
		result := bytes.Compare(owners[i].OwnerAddress, ownerAddress)
		return result == 0
	})
	if index == len(owners) {
		return Owner{}, sdk.ErrInvalidAddress("Owner not found.")
	}
	return owners[index], nil
}

// GetFiatPegHashHex : convert string to hex peg hash
func GetFiatPegHashHex(pegHashStr string) (pegHash PegHash, err error) {
	if len(pegHashStr) == 0 {
		return pegHash, errors.New("must use provide pegHash")
	}
	bz, err := hex.DecodeString(pegHashStr)
	if err != nil {
		return nil, err
	}
	return PegHash(bz), nil
}

// NewBaseFiatPegWithPegHash a base fiat peg with peg hash
func NewBaseFiatPegWithPegHash(pegHash PegHash) BaseFiatPeg {
	return BaseFiatPeg{
		PegHash: pegHash,
	}
}

type FiatAccount interface {
	GetAddress() sdk.AccAddress
	SetAddress(sdk.AccAddress)

	GetFiatPegWallet() []FiatPeg
	SetFiatPegWallet([]FiatPeg)
}

type BaseFiatAccount struct {
	Address       sdk.AccAddress `json:"address"`
	FiatPegWallet []FiatPeg      `json:"fiatPegWallet"`
}

var _ FiatAccount = (*BaseFiatAccount)(nil)

// GetAddress : getter
func (baseFiatAccount BaseFiatAccount) GetAddress() sdk.AccAddress {
	return baseFiatAccount.Address
}

// SetAddress : setter
func (baseFiatAccount *BaseFiatAccount) SetAddress(address sdk.AccAddress) {
	baseFiatAccount.Address = address
}

// GetFiatPegWallet : getter
func (baseFiatAccount BaseFiatAccount) GetFiatPegWallet() []FiatPeg {
	return baseFiatAccount.FiatPegWallet
}

// SetFiatPegWallet : setter
func (baseFiatAccount *BaseFiatAccount) SetFiatPegWallet(fiatPegWallet []FiatPeg) {
	baseFiatAccount.FiatPegWallet = fiatPegWallet
}

func NewBaseFiatAccountWithAddress(address sdk.AccAddress) BaseFiatAccount {
	return BaseFiatAccount{
		Address: address,
	}
}
