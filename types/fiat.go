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
func NewFiatPegWithPegHash(pegHash PegHash) FiatPeg {
	return &BaseFiatPeg{
		PegHash: pegHash,
	}
}

// FiatPegWallet : A wallet of fiat peg tokens
type FiatPegWallet []BaseFiatPeg

// ProtoBaseFiatPeg : create a new interface prototype
func ProtoBaseFiatPeg() FiatPeg {
	return &BaseFiatPeg{}
}

// Sort interface

func (fiatPegWallet FiatPegWallet) Len() int { return len(fiatPegWallet) }

func (fiatPegWallet FiatPegWallet) Less(i, j int) bool {
	if (fiatPegWallet[i].TransactionAmount - fiatPegWallet[j].TransactionAmount) < 0 {
		return true
	}
	return false
}
func (fiatPegWallet FiatPegWallet) Swap(i, j int) {
	fiatPegWallet[i], fiatPegWallet[j] = fiatPegWallet[j], fiatPegWallet[i]
}

var _ sort.Interface = FiatPegWallet{}

// Sort is a helper function to sort the set of fiat pegs inplace
func (fiatPegWallet FiatPegWallet) Sort() FiatPegWallet {
	sort.Sort(fiatPegWallet)
	return fiatPegWallet
}

// ToFiatPeg : convert base fiat peg to interface fiat peg
func ToFiatPeg(baseFiatPeg BaseFiatPeg) FiatPeg {
	fiatI := ProtoBaseFiatPeg()
	fiatI.SetOwners(baseFiatPeg.Owners)
	fiatI.SetPegHash(baseFiatPeg.PegHash)
	fiatI.SetRedeemedAmount(baseFiatPeg.RedeemedAmount)
	fiatI.SetTransactionAmount(baseFiatPeg.TransactionAmount)
	fiatI.SetTransactionID(baseFiatPeg.TransactionID)
	return fiatI
}

// ToBaseFiatPeg : convert interface fiat peg to base fiat peg
func ToBaseFiatPeg(fiatPeg FiatPeg) BaseFiatPeg {
	var baseFiatPeg BaseFiatPeg
	baseFiatPeg.Owners = fiatPeg.GetOwners()
	baseFiatPeg.PegHash = fiatPeg.GetPegHash()
	baseFiatPeg.RedeemedAmount = fiatPeg.GetRedeemedAmount()
	baseFiatPeg.TransactionAmount = fiatPeg.GetTransactionAmount()
	baseFiatPeg.TransactionID = fiatPeg.GetTransactionID()
	return baseFiatPeg
}

// SubtractFiatPegWalletFromWallet : subtract fiat peg  wallet from wallet
func SubtractFiatPegWalletFromWallet(inFiatPegWallet FiatPegWallet, fiatPegWallet FiatPegWallet) FiatPegWallet {
	for _, inFiatPeg := range inFiatPegWallet {
		for i, fiatPeg := range fiatPegWallet {
			if fiatPeg.GetPegHash().String() == inFiatPeg.GetPegHash().String() {
				fiatPegWallet = append(fiatPegWallet[:i], fiatPegWallet[i+1:]...)
				fiatPegWallet = fiatPegWallet.Sort()
				break
			}
		}
	}
	return fiatPegWallet
}

// SubtractAmountFromWallet : subtract fiat peg from wallet
func SubtractAmountFromWallet(amount int64, fiatPegWallet FiatPegWallet) (sentFiatPegWallet FiatPegWallet, oldFiatPegWallet FiatPegWallet) {
	var collected int64
	fiatPegWallet = fiatPegWallet.Sort()
	for _, fiatPeg := range fiatPegWallet {
		if collected < amount {
			if fiatPeg.TransactionAmount <= amount-collected {
				collected += fiatPeg.TransactionAmount
				sentFiatPegWallet = append(sentFiatPegWallet, fiatPeg)
			} else if fiatPeg.TransactionAmount > amount-collected {
				splitFiatPeg := fiatPeg
				splitFiatPeg.TransactionAmount = amount - collected
				fiatPeg.TransactionAmount -= amount - collected
				oldFiatPegWallet = append(oldFiatPegWallet, fiatPeg)
				sentFiatPegWallet = append(sentFiatPegWallet, splitFiatPeg)
				collected += amount - collected
			}
		} else {
			oldFiatPegWallet = append(oldFiatPegWallet, fiatPeg)
		}
	}
	if collected == amount {
		oldFiatPegWallet = oldFiatPegWallet.Sort()
		sentFiatPegWallet = sentFiatPegWallet.Sort()
		return
	}
	return FiatPegWallet{}, FiatPegWallet{}

}

// RedeemAmountFromWallet : subtract fiat peg from wallet
func RedeemAmountFromWallet(amount int64, fiatPegWallet FiatPegWallet) (emptiedFiatPegWallet FiatPegWallet, redeemerFiatPegWallet FiatPegWallet) {
	var collected int64
	for _, fiatPeg := range fiatPegWallet {
		if collected < amount {
			if fiatPeg.TransactionAmount <= amount-collected {
				collected += fiatPeg.TransactionAmount
				emptiedFiatPegWallet = append(emptiedFiatPegWallet, fiatPeg)
			} else if fiatPeg.TransactionAmount > amount-collected {
				fiatPeg.TransactionAmount -= amount - collected
				fiatPeg.RedeemedAmount = amount - collected
				redeemerFiatPegWallet = append(redeemerFiatPegWallet, fiatPeg)
				collected += amount - collected
			}
		} else {
			redeemerFiatPegWallet = append(redeemerFiatPegWallet, fiatPeg)
		}
	}
	if collected == amount {
		redeemerFiatPegWallet = redeemerFiatPegWallet.Sort()
		emptiedFiatPegWallet = emptiedFiatPegWallet.Sort()
		return
	}
	return FiatPegWallet{}, FiatPegWallet{}

}

// AddFiatPegToWallet : add fiat peg to wallet
func AddFiatPegToWallet(fiatPegWallet FiatPegWallet, inFiatPegWallet FiatPegWallet) FiatPegWallet {
	for _, inFiatPeg := range inFiatPegWallet {
		added := false
		for i, fiatPeg := range fiatPegWallet {
			if fiatPeg.PegHash.String() == inFiatPeg.PegHash.String() {
				inFiatPeg.TransactionAmount += fiatPeg.TransactionAmount
				fiatPegWallet[i] = inFiatPeg
				added = true
				break
			}
		}
		if !added {
			fiatPegWallet = append(fiatPegWallet, inFiatPeg)
		}
	}
	fiatPegWallet = fiatPegWallet.Sort()
	return fiatPegWallet
}

// GetFiatPegWalletBalance :  gets the total sum of all fiat pegs in a wallet
func GetFiatPegWalletBalance(fiatPegWallet FiatPegWallet) int64 {
	var balance int64
	for _, fiatPeg := range fiatPegWallet {
		balance += fiatPeg.TransactionAmount
	}
	return balance
}

// TransferFiatPegsToWallet : subtracts and changes owners of fiat peg in fiat chain
func TransferFiatPegsToWallet(fiatPegWallet FiatPegWallet, oldFiatPegWallet FiatPegWallet, fromAddress sdk.AccAddress, toAddress sdk.AccAddress) FiatPegWallet {
	for _, fiatPeg := range fiatPegWallet {
		transfered := false
		for j, oldFiatPeg := range oldFiatPegWallet {
			if fiatPeg.GetPegHash().String() == oldFiatPeg.GetPegHash().String() {
				subtracted := 0
				added := 0
				for i, owner := range oldFiatPeg.Owners {
					if owner.OwnerAddress.String() == fromAddress.String() && owner.Amount >= fiatPeg.TransactionAmount {
						owner.Amount -= fiatPeg.TransactionAmount
						oldFiatPeg.Owners[i] = owner
						subtracted++
					} else if owner.OwnerAddress.String() == toAddress.String() {
						owner.Amount += fiatPeg.TransactionAmount
						oldFiatPeg.Owners[i] = owner
						added++
					}
				}
				if added == 0 {
					owner := Owner{toAddress, fiatPeg.TransactionAmount}
					oldFiatPeg.Owners = append(oldFiatPeg.Owners, owner)
					added++
				}
				if subtracted != 1 || added != 1 {
					return nil
				}
				transfered = true
				oldFiatPegWallet[j] = oldFiatPeg
				break
			}
		}
		if !transfered {
			return nil
		}
	}
	return oldFiatPegWallet
}

// RedeemFiatPegsFromWallet : subtracts and changes owners of fiat peg in fiat chain
func RedeemFiatPegsFromWallet(fiatPegWallet FiatPegWallet, oldFiatPegWallet FiatPegWallet, fromAddress sdk.AccAddress) FiatPegWallet {
	for _, fiatPeg := range fiatPegWallet {
		transfered := false
		for j, oldFiatPeg := range oldFiatPegWallet {
			if fiatPeg.GetPegHash().String() == oldFiatPeg.GetPegHash().String() {
				subtracted := 0

				for i, owner := range oldFiatPeg.Owners {
					if owner.OwnerAddress.String() == fromAddress.String() && owner.Amount > fiatPeg.RedeemedAmount {
						owner.Amount -= fiatPeg.RedeemedAmount
						oldFiatPeg.Owners[i] = owner
						subtracted++
					} else if owner.OwnerAddress.String() == fromAddress.String() && owner.Amount == fiatPeg.RedeemedAmount {
						oldFiatPeg.Owners = append(oldFiatPeg.Owners[:i], oldFiatPeg.Owners[i+1:]...)
						subtracted++
					}
				}

				if subtracted != 1 {
					return nil
				}
				oldFiatPeg.TransactionAmount -= fiatPeg.RedeemedAmount
				oldFiatPeg.RedeemedAmount += fiatPeg.RedeemedAmount

				transfered = true
				oldFiatPegWallet[j] = oldFiatPeg
				break
			}
		}
		if !transfered {
			return nil
		}
	}
	return oldFiatPegWallet
}

type FiatAccount interface {
	GetAddress() sdk.AccAddress
	SetAddress(sdk.AccAddress)

	GetFiatPegWallet() FiatPegWallet
	SetFiatPegWallet(FiatPegWallet)
}

type BaseFiatAccount struct {
	Address       sdk.AccAddress `json:"address"`
	FiatPegWallet FiatPegWallet  `json:"fiatPegWallet"`
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
func (baseFiatAccount BaseFiatAccount) GetFiatPegWallet() FiatPegWallet {
	return baseFiatAccount.FiatPegWallet
}

// SetFiatPegWallet : setter
func (baseFiatAccount *BaseFiatAccount) SetFiatPegWallet(fiatPegWallet FiatPegWallet) {
	baseFiatAccount.FiatPegWallet = fiatPegWallet
}

func NewFiatAccountWithAddress(address sdk.AccAddress) FiatAccount {
	return &BaseFiatAccount{
		Address: address,
	}
}
