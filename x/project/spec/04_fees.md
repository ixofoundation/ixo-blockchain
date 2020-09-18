# Fees

As discussed in the messages page, a project doc's [arbitrary data](02_messages.md#non-arbitrary-project-data) is expected to include a `fees` field which follows a specific format, which is represented by the following struct:

```go
type ProjectFeesMap struct {
	Context string `json:"@context" yaml:"@context"`
	Items   []struct {
		Type              FeeType `json:"@type" yaml:"@type"`
		PaymentTemplateId string  `json:"id" yaml:"id"`
	}
}
```

The following is an example of a `fees` section inside a project's `data`:
```json
"data": {
    ...
    "fees": {
        "@context": "...",
        "items": [
            {
                "@type": "OracleFee",
                "id":"payment:template:oracle-fee-template-1"
            },
            {
                "@type": "FeeForService", 
                "id":"payment:template:fee-for-service-template-1"
            }
        ]
    }
    ...
}
```

The two (optional) project-related fees currently supported are:
- Oracle Fee (`OracleFee`)
- Fee for Service (`FeeForService`)

## General Approach

In general, the project doc's `fees` section will specify payment template IDs of pre-created payment templates. These will specify values such as the payment amounts, minimums, maximums, etc (refer to payments module spec).

However, we require a payment contract to be created in order to actually effect a payment. The general approach is to create this contract on-the-fly whenever the pre-conditions for effecting a payment are satisfied.

A general `processPay` function involves most of the functionality required to create a contract and effect a payment. A simplified version of this function is presented below:
```go
func processPay(keepers, projectDid, senderAddr, recipients, feeType, paymentTemplateId) {

	// Validate recipients
	recipients.Validate()

	// Get project address
	projectAddr := getProjectAccountAddress()

	// Get payment template
	template := paymentsKeeper.GetPaymentTemplate(paymentTemplateId)

	// Contruct payment contract ID
	contractId := "payment:contract:<moduleName>:<projectDid>:<senderAddr>:<feeType>"

	// Get or create payment contract, based on if exists already or not
	if paymentsKeeper.ContractExists(contractId) {
		contract = paymentsKeeper.GetPaymentContract(contractId)
	} else {
		// Create new contract with the constructed contractId, pointing to the
		// specified payment template, with the project as the creator and payer,
		// the specified recipients. The contract cannot be deauthorised and is
		// by default authorised (i.e. can be effected by creator).
		contract = payments.NewPaymentContract(contractId, paymentTemplateId,
			projectAddr, projectAddr, recipients, false, true)
		paymentsKeeper.SavePaymentContract(contract)
	}

	// Effect payment if can effect
	if contract.CanEffectPayment {
		// Check that project has enough tokens to effect contract payment
		if !bankKeeper.HasCoins(projectAddr, template.PaymentAmount) {
			return error("project has insufficient funds")
		}

		// Effect payment
		paymentsKeeper.EffectPayment(contractId)
	} else {
		return error("cannot effect contract payment (max reached?)")
	}
}
```

The following are some key points about the above code:

The payment contract's ID is contructed on-the-fly using the below template. This means that the payment contract created is unique to the projects module, project, sender address, and fee type.

- Template: `payment:contract:<moduleName>:<projectDid>:<senderAddr>:<feeType>`
- Example: `payment:contract:project:did:ixo:U7G...J8c:ixo107...0vx:OracleFee`

Thus, if a new and unique set of the above 4 values is encountered, a new payment contract is created. Otherwise, the existing payment contract is fetched. This means that the project contract can (and is) used to persist information between two or more payments of the same type.

An example use case is when we want to specify a maximum payment. Consider a contract created based on a payment template that specifies a pay amount of 100IXO and a maximum of 300IXO.
1. In the first `processPay()` call, a new payment contract is created and immediately effected (cumulative pay: `100IXO`)
2. In the second `processPay()` call, the payment contract already exists and is fetched and the payment is effected (cumulative pay: `200IXO`).
3. In the third `processPay()` call, the cumulative pay will have reached the maximum `300IXO`.
4. If a fourth `processPay()` call comes through, no further payment is effected, given that the contract cannot be effected due to the maximum being reached.

If the project does not have enough funds to effect the payment contract, an error is returned by the function. An error is also returned if the payment contract payment cannot be effected.

## Specific Fees

### Oracle Fee

If an oracle fee payment template ID was provided during project creation, the oracle fee applies whenever an oracle (a.k.a evaluator) evaluates a claim using `MsgCreateEvaluation`. It is paid by the project, with the oracle and other fee project addresses as the contract recipients.

### Fee for Service

If a fee-for-service payment template ID was provided during project creation, the fee-for-service applies whenever a claimer's claim gets approved (claim status set to approved `= 1`) during `MsgCreateEvaluation` handling. It is paid by the project, with the claimer as the contract recipient.
