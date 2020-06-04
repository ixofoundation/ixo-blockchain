package types

import (
	"regexp"
)

var (
	ValidFeeId          = regexp.MustCompile(`^fee:[a-zA-Z][a-zA-Z0-9/_:-]*$`)
	ValidFeeContractId  = regexp.MustCompile(`^fee:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`)
	ValidSubscriptionId = regexp.MustCompile(`^fee:subscription:[a-zA-Z][a-zA-Z0-9/_:-]*$`)

	IsValidFeeId          = ValidFeeId.MatchString
	IsValidFeeContractId  = ValidFeeContractId.MatchString
	IsValidSubscriptionId = ValidSubscriptionId.MatchString
)
