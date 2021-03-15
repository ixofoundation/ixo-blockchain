# Messages

In this section we describe the processing of the payment messages and the
corresponding updates to the state. All created/modified state objects specified
by each message are defined within the [state](01_state.md) section.

## MsgCreatePaymentTemplate

This message creates and stores the payment template at appropriate indexes.
Refer to [01_state.md](./01_state.md) for information about payment templates.

| **Field**       | **Type**          | **Description** |
|:----------------|:------------------|:----------------|
| CreatorDid      | `did.Did`         | Did of the creator
| PaymentTemplate | `PaymentTemplate` | The payment template being created

```go
type MsgCreatePaymentTemplate struct {
	CreatorDid      did.Did
	PaymentTemplate PaymentTemplate
}
``` 

This message is expected to fail if:

- Creator DID is empty or invalid
- Payment template ID violates `^payment:template:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:template:", followed by a letter, followed by
      any mix of alphanumeric characters and the `/`, `_`, `:`, `-` characters.
    - Valid ID example: `payment:template:abc_012-def/345:ghi`
- Payment template already exists (by ID)
- Payment template ID is reserved
- Payment template validation checks fail (these essentially re-check the above)

## MsgCreatePaymentContract

This message creates and stores the payment contract at appropriate indexes.
Refer to [01_state.md](./01_state.md) for information about payment contracts.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| CreatorDid        | `did.Did`        | Did of the user
| PaymentTemplateId | `string`         | ID of the paymentTemplate
| PaymentContractId | `string`         | ID of the PaymentContract
| Payer             | `sdk.AccAddress` | Address of the payer
| Recipients        | `Distribution`   | List of recipients with percentage shares
| CanDeauthorise    | `bool`           | Bool of de_authorise
| DiscountId        | `sdk.Uint`       | Any discount given

```go
type MsgCreatePaymentContract struct {
	CreatorDid        did.Did
	PaymentTemplateId string
	PaymentContractId string
	Payer             sdk.AccAddress
	Recipients        Distribution
	CanDeauthorise    bool
	DiscountId        sdk.Uint
}
```

This message is expected to fail if:

- Creator DID is empty or invalid
- Payment contract ID violates `^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:contract:", followed by a letter, followed by
      any mix of alphanumeric characters and `/`, `_`, `:`, `-`.
    - Valid ID example: `payment:contract:abc_012-def/345:ghi`
- Payment template ID violates `^payment:template:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:template:", followed by a letter, followed by
      any mix of alphanumeric characters and `/`, `_`, `:`, `-`.
    - Valid ID example: `payment:template:abc_012-def/345:ghi`
- Recipients list is invalid:
    - Recipient distribution shares do not add up to 100%
    - Any distribution share percentage is not positive
    - Any distribution share address is empty
- Payment contract already exists (by ID)
- Payment template does not exist (by ID)
- Payment contract ID is reserved
- Payer or payee addresses are blacklisted
- Creator DID is not registered with the DID module
- Payment contract validation checks fail (these essentially re-check the above)

## MsgCreateSubscription

This message creates and stores the subscription at appropriate indexes. Refer
to [01_state.md](./01_state.md) for information about subscriptions.

| **Field**         | **Type**   | **Description** |
|:------------------|:-----------|:----------------|
| CreatorDid        | `did.Did`  | Did of the user
| SubscriptionId    | `string`   | ID for the subscription
| PaymentContractId | `string`   | ID for the paymentContract
| MaxPeriods        | `sdk.Uint` | Maximum number of times chargeable
| Period            | `Period`   | IF the periods are allowed or not

```go
type MsgCreateSubscription struct {
	CreatorDid        did.Did
	SubscriptionId    string
	PaymentContractId string
	MaxPeriods        sdk.Uint
	Period            Period
}
```

This message is expected to fail if:

- Creator DID is empty or invalid
- Subscription ID violates `^payment:subscription:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:subscription:", followed by a letter, followed
      by any mix of alphanumeric characters and `/`, `_`, `:`, `-`.
    - Valid ID example: `payment:subscription:abc_012-def/345:ghi`
- Payment contract ID violates `^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:contract:", followed by a letter, followed by
      any mix of alphanumeric characters and the `/`, `_`, `:`, `-` characters.
    - Valid ID example: `payment:contract:abc_012-def/345:ghi`
- Subscription period is invalid:
    - Start time/block is after end time/block
    - Period length is zero
- Subscription already exists (by ID)
- Payment contract does not exist (by ID)
- Subscription ID is reserved
- Subscription creator is not payment contract creator
- Creator DID is not registered with the DID module
- Subscription validation checks fail (these essentially re-check the above)

## MsgSetPaymentContractAuthorisation

This message authorises or deauthorises a payment contract, in order to enable
or disable effecting of the payment contract.

| **Field**         | **Type**  | **Description** |
|:------------------|:----------|:----------------|
| PayerDid          | `did.Did` | Payer's DID
| PaymentContractId | `string`  | ID of the paymentContract
| Authorised        | `bool`    | Status of authorisation

```go
type MsgSetPaymentContractAuthorisation struct {
	PayerDid          did.Did
	PaymentContractId string
	Authorised        bool
}
``` 

This message is expected to fail if:

- Payer DID is empty or invalid
- Payment contract ID violates `^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:contract:", followed by a letter, followed by
      any mix of alphanumeric characters and the `/`, `_`, `:`, `-` characters.
    - Valid ID example: `payment:contract:abc_012-def/345:ghi`
- Payment contract does not exist (by ID)
- Payer DID is not registered with the DID module
- Payer does not match the payment contract payer

## MsgGrantDiscount

This message grants a discount to a recipient for a particular payment contract.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| SenderDid         | `did.Did`        | Initiator DID
| PaymentContractId | `string`         | ID for the paymentContract
| DiscountId        | `sdk.Uint`       | How much is the discount
| Recipient         | `sdk.AccAddress` | Who is the recipient

```go
type MsgGrantDiscount struct {
	SenderDid         did.Did
	PaymentContractId string
	DiscountId        sdk.Uint
	Recipient         sdk.AccAddress
}
```

This message is expected to fail if:

- Sender DID is empty or invalid
- Payment contract ID violates `^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:contract:", followed by a letter, followed by
      any mix of alphanumeric characters and the `/`, `_`, `:`, `-` characters.
    - Valid ID example: `payment:contract:abc_012-def/345:ghi`
- Payment contract does not exist (by ID)
- Discount ID must be in the list of payment template discounts
- Sender does not match the payment contract creator

## MsgRevokeDiscount

This message revokes a discount of a discount holder for a particular payment
contract.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| SenderDid         | `did.Did`        | Who send the transaction
| PaymentContractId | `string`         | ID of the payment_Contract
| Holder            | `sdk.AccAddress` | Address of who's holds the discount

```go
type MsgRevokeDiscount struct {
	SenderDid         did.Did
	PaymentContractId string
	Holder            sdk.AccAddress
}
```

This message is expected to fail if:

- Sender DID is empty or invalid
- Payment contract ID violates `^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:contract:", followed by a letter, followed by
      any mix of alphanumeric characters and the `/`, `_`, `:`, `-` characters.
    - Valid ID example: `payment:contract:abc_012-def/345:ghi`
- Payment contract does not exist (by ID)
- Sender does not match the payment contract creator
