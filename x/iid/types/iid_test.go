package types

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"
)

func TestNewChainDID(t *testing.T) {

	yy := NewPublicKeyMultibase([]byte{3, 223, 208, 164, 105, 128, 109, 102, 162, 60, 124, 148, 143, 85, 193, 41, 70, 125, 109, 9, 116, 162, 34, 239, 110, 36, 165, 56, 250, 104, 130, 243, 215}, DIDVMethodTypeEcdsaSecp256k1VerificationKey2019)
	mm := base58.Encode(yy.data)
	fmt.Println(mm)
	fmt.Println(yy.EncodeToString())

	x := "2ApAAUuL6yHYXKUcDQkbbZGCgBBwLZtqrLxPc6ra7vZdv"

	c, err := toAddress(mm[1:])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
	fmt.Println()
	y, err := toAddress(x[1:])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y)
	}

	assert.True(t, false)

}
