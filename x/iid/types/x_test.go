package types

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"
)

func TestPubkey(t *testing.T) {

	pubKeyString := "FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG"
	pubKeyBytes := base58.Decode(pubKeyString)
	result := len(pubKeyBytes) == 256

	fmt.Println(len(pubKeyBytes))
	assert.True(t, result)
}
