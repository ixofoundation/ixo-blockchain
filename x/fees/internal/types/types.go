package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	"regexp"
)

var (
	ValidFeeId          = regexp.MustCompile(`^fee:[a-zA-Z][a-zA-Z0-9/_]*$`)
	ValidFeeContractId  = regexp.MustCompile(`^fee:contract:[a-zA-Z][a-zA-Z0-9/_]*$`)
	ValidSubscriptionId = regexp.MustCompile(`^fee:subscription:[a-zA-Z][a-zA-Z0-9/_]*$`)

	IsValidFeeId          = ValidFeeId.MatchString
	IsValidFeeContractId  = ValidFeeContractId.MatchString
	IsValidSubscriptionId = ValidSubscriptionId.MatchString
)

func DidToAddr(did ixo.Did) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(did)))
}
