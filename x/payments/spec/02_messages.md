# Messages

In this section we describe the processing of the payment messages and the
corresponding updates to the state. All created/modified state objects specified
by each message are defined within the [state](01_state.md) section.

## MsgCreatePaymentTemplate

This message creates and stores the payment template at appropriate indexes.
Refer to [01_state.md](./01_state.md) for information about payment templates.

| **Field**       | **Type**          | **Description** |
|:----------------|:------------------|:----------------|
| CreatorDid      | `did.Did`         | DID of the template creator (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PaymentTemplate | `PaymentTemplate` | Details of the payment template being created

```go
type MsgCreatePaymentTemplate struct {
    CreatorDid      string
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
| CreatorDid        | `did.Did`        | DID of the contract creator (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PaymentTemplateId | `string`         | ID of the payment template on which this contract is based (e.g. `payment:template:template1`)
| PaymentContractId | `string`         | ID of this payment contract (e.g. `payment:contract:contract1`)
| Payer             | `sdk.AccAddress` | Address from where tokens will be deducted (e.g. `ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y`)
| Recipients        | `Distribution`   | List of token recipients with percentage shares
| CanDeauthorise    | `bool`           | Whether or not this contract can be de-authorised
| DiscountId        | `sdk.Uint`       | Any discount assigned to this contract (discounts defined in the template) (e.g. `0`)

```go
type MsgCreatePaymentContract struct {
    CreatorDid        string
    PaymentTemplateId string
    PaymentContractId string
    Payer             string
    Recipients        []DistributionShare
    CanDeauthorise    bool
    DiscountId        sdk.Uint
}
```

This message is expected to fail if:

- Creator DID is empty or invalid
- Payment template ID violates `^payment:template:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:template:", followed by a letter, followed by
      any mix of alphanumeric characters and `/`, `_`, `:`, `-`.
    - Valid ID example: `payment:template:abc_012-def/345:ghi`
- Payment contract ID violates `^payment:contract:[a-zA-Z][a-zA-Z0-9/_:-]*$`
    - ID must start with "payment:contract:", followed by a letter, followed by
      any mix of alphanumeric characters and `/`, `_`, `:`, `-`.
    - Valid ID example: `payment:contract:abc_012-def/345:ghi`
- Recipients list is invalid:
    - Recipient distribution shares do not add up to 100%
    - Any distribution share percentage is not positive
    - Any distribution share address is empty
- Payment template does not exist (by ID)
- Payment contract already exists (by ID)
- Payment contract ID is reserved
- Payer or payee addresses are blacklisted
- Creator DID is not registered with the DID module
- Payment contract validation checks fail (these essentially re-check the above)

## MsgCreateSubscription

This message creates and stores the subscription at appropriate indexes. Refer
to [01_state.md](./01_state.md) for information about subscriptions.

| **Field**         | **Type**   | **Description** |
|:------------------|:-----------|:----------------|
| CreatorDid        | `did.Did`  | DID of the subscription creator (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| SubscriptionId    | `string`   | ID of this subscription (e.g. `payment:subscription:subscription1`)
| PaymentContractId | `string`   | ID of the payment contract on which this subscription is based (e.g. `payment:contract:contract1`)
| MaxPeriods        | `sdk.Uint` | Maximum number of times that the subscription payment is triggered
| Period            | `Period`   | Describes the period (period unit and period length)

```go
type MsgCreateSubscription struct {
    CreatorDid        string
    SubscriptionId    string
    PaymentContractId string
    MaxPeriods        sdk.Uint
    Period            *sdk.Any
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
| PayerDid          | `did.Did` | DID of the payer associated with the contract (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PaymentContractId | `string`  | ID of the payment contract being modified (e.g. `payment:contract:contract1`)
| Authorised        | `bool`    | New status of authorisation for the contract

```go
type MsgSetPaymentContractAuthorisation struct {
    PayerDid          string
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
| SenderDid         | `did.Did`        | DID of the discount granter (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PaymentContractId | `string`         | ID of the contract for which the discount is being granted (e.g. `payment:contract:contract1`)
| DiscountId        | `sdk.Uint`       | ID of the discount being granted (e.g. `0`)
| Recipient         | `sdk.AccAddress` | The recipient of the discount (e.g. `ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y`)

```go
type MsgGrantDiscount struct {
    SenderDid         string
    PaymentContractId string
    DiscountId        sdk.Uint
    Recipient         string
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
| SenderDid         | `did.Did`        | DID of the discount revoker (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PaymentContractId | `string`         | ID of the contract from which the discount is being revoked (e.g. `payment:contract:contract1`)
| Holder            | `sdk.AccAddress` | The current holder of the discount (e.g. `ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y`)

```go
type MsgRevokeDiscount struct {
    SenderDid         string
    PaymentContractId string
    Holder            string
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

## MsgEffectPayment

This message puts into effect a particular payment contract.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| SenderDid         | `did.Did`        | DID of the message sender (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PaymentContractId | `string`         | ID of the contract being effected (e.g. `payment:contract:contract1`)

```go
type MsgEffectPayment struct {
    SenderDid         string
    PaymentContractId string
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
- Payment is not effected (e.g. if max pay has been reached or if payer does
  not have enough coins)
