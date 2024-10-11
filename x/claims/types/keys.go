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
	CollectionKey = []byte{0x01}
	ClaimKey      = []byte{0x02}
	DisputeKey    = []byte{0x03}
	IntentKey     = []byte{0x04}
)

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
