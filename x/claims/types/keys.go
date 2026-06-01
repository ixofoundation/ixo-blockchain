package types

import (
	fmt "fmt"
	"strings"
)

const (
	// ModuleName defines the module name
	ModuleName = "claims"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	CollectionKey   = []byte{0x01}
	ClaimKey        = []byte{0x02}
	DisputeKey      = []byte{0x03}
	IntentKey       = []byte{0x04}
	MemberBudgetKey = []byte{0x05}
	// AgentDepositBalanceKey indexes rolling per-(collection, agent) performance-deposit
	// balances. Funds live in collection.escrow_account (shared with intent escrow);
	// this key tracks accounting only.
	AgentDepositBalanceKey = []byte{0x06}
	// ActiveDisputeKey is a presence-only secondary index for "is agent X
	// currently targeted by any OPEN dispute on collection Y?". Key format:
	// collectionId + "/" + agentAddress + "/" + subjectId. The check is a
	// prefix scan with limit 1, O(1) gas, used by SubmitClaim / EvaluateClaim
	// / WithdrawPerformanceDeposit to gate the actor.
	ActiveDisputeKey = []byte{0x07}
	// DisputeSubjectIndexKey maps (subjectId, targetRole) -> dispute proof
	// (the dispute's own proof CID, which is its primary-store key). Stores
	// the LATEST dispute for that pair; the primary record is fetched to
	// determine its status. Used to enforce one-OPEN-at-a-time and to
	// permanently block new disputes when the latest is AWARDED.
	DisputeSubjectIndexKey = []byte{0x08}
)

// MemberBudgetKeyCreate creates key for MemberBudget KV Store: collectionId+"/"+memberAddress
func MemberBudgetKeyCreate(collectionId, memberAddress string) []byte {
	return []byte(collectionId + "/" + memberAddress)
}

// AgentDepositBalanceKeyCreate builds the AgentDepositBalance KV key from
// collectionId and agentAddress.
func AgentDepositBalanceKeyCreate(collectionId, agentAddress string) []byte {
	return []byte(collectionId + "/" + agentAddress)
}

// ActiveDisputeKeyCreate builds the active-dispute presence-index key.
func ActiveDisputeKeyCreate(collectionId, agentAddress, subjectId string) []byte {
	return []byte(collectionId + "/" + agentAddress + "/" + subjectId)
}

// ActiveDisputeAgentPrefix builds the prefix used to scan "does agent X have
// any OPEN dispute against them on collection Y?".
func ActiveDisputeAgentPrefix(collectionId, agentAddress string) []byte {
	return []byte(collectionId + "/" + agentAddress + "/")
}

// DisputeSubjectIndexKeyCreate builds the (subjectId, targetRole) index key.
// targetRole is included so disputes against submitter and evaluator on the
// same claim live in distinct slots.
func DisputeSubjectIndexKeyCreate(subjectId string, targetRole DisputeTargetRole) []byte {
	return []byte(subjectId + "/" + fmt.Sprint(int32(targetRole)))
}

// Create the key for Intents KV Store: agentAddress+"/"+collectionID+"/"+intentID
func IntentKeyCreate(agentAddress, collectionId, intentID string) []byte {
	return []byte(agentAddress + "/" + collectionId + "/" + intentID)
}

// Unused for now
// ParseActiveIntentKey parses the ActiveIntentKey into agentAddress and collectionId
func ParseActiveIntentKey(key []byte) (agentAddress string, collectionId string, err error) {
	// Convert byte array to string and split by "/"
	parts := strings.Split(string(key), "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid active intent key: %s", string(key))
	}
	agentAddress = parts[0]
	collectionId = parts[1]
	return agentAddress, collectionId, nil
}
