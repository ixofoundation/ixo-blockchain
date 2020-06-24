package types

import "github.com/ixofoundation/ixo-blockchain/x/did"

type StoredBondDoc interface {
	GetBondDid() did.Did
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
	Status BondStatus `json:"status" yaml:"status"`
}

type BondDoc struct {
	CreatedOn string     `json:"created_on" yaml:"created_on"`
	CreatedBy string     `json:"created_by" yaml:"created_by"`
	Status    BondStatus `json:"status" yaml:"status"`
}
