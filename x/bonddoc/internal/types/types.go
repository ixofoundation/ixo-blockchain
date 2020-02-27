package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type StoredBondDoc interface {
	GetBondDid() ixo.Did
	GetPubKey() string
	GetStatus() BondStatus
	SetStatus(status BondStatus)
}

type BondStatus string

const (
	NullStatus        BondStatus = ""
	PreIssuanceStatus BondStatus = "PREISSUANCE"
	OpenStatus        BondStatus = "OPEN"
	SuspendedStatus   BondStatus = "SUSPENDED"
	ClosedStatus      BondStatus = "CLOSED"
	SettlementStatus  BondStatus = "SETTLEMENT"
	EndedStatus       BondStatus = "ENDED"
)

var StateTransitions = initStateTransitions()

func initStateTransitions() map[BondStatus][]BondStatus {
	return map[BondStatus][]BondStatus{
		NullStatus:        {PreIssuanceStatus},
		PreIssuanceStatus: {OpenStatus},
		OpenStatus:        {SuspendedStatus, ClosedStatus, SettlementStatus},
		SuspendedStatus:   {OpenStatus, ClosedStatus},
		SettlementStatus:  {EndedStatus},
	}

}

func (nextBondStatus BondStatus) IsValidProgressionFrom(previousBondStatus BondStatus) bool {
	validStatuses := StateTransitions[previousBondStatus]
	for _, v := range validStatuses {
		if v == nextBondStatus {
			return true
		}
	}

	return false
}

type UpdateBondStatusDoc struct {
	Status BondStatus `json:"status"`
}

type BondDoc struct {
	CreatedOn string     `json:"createdOn"`
	CreatedBy string     `json:"createdBy"`
	Status    BondStatus `json:"status"`
}

type BondDocDecoder func(bondEntryBytes []byte) (StoredBondDoc, error)

type BondMsg interface {
	sdk.Msg
	IsNewDid() bool
}
