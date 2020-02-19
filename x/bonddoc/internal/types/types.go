package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type StoredBondDoc interface {
	GetBondDid() ixo.Did
	GetPubKey() string
}

type BondDoc struct {
	CreatedOn string `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
}

type BondDocDecoder func(bondEntryBytes []byte) (StoredBondDoc, error)

type BondMsg interface {
	sdk.Msg
	IsNewDid() bool
}
