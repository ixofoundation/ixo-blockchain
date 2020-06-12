package types

import (
	"regexp"
)

var (
	ValidPaymentTemplateId = regexp.MustCompile(`^payment:template:[a-zA-Z][a-zA-Z0-9/_:-]*$`)
	ValidPaymentContractId = regexp.MustCompile(`^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`)
	ValidSubscriptionId    = regexp.MustCompile(`^payment:subscription:[a-zA-Z][a-zA-Z0-9/_:-]*$`)

	IsValidPaymentTemplateId = ValidPaymentTemplateId.MatchString
	IsValidPaymentContractId = ValidPaymentContractId.MatchString
	IsValidSubscriptionId    = ValidSubscriptionId.MatchString
)
