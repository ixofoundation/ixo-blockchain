# Messages

In this section we describe the processing of the payment messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](01_state.md) section.

## MsgCreatePaymentTemplate

This message creates and stores the payment template at appropriate indexes.

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

## MsgCreatePaymentContract 

This message creates and stores the payment contract at appropriate indexes.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| CreatorDid        | `string`         | Did of the user
| PaymentTemplateId | `string`         | ID of the paymentTemplate
| PaymentContractId | `string`         | ID of the PaymentContract
| Payer             | `sdk.AccAddress` | Address of the payer
| CanDeauthorise    | `bool`           | Bool of de_authorise
| DiscountId        | ` sdk.Uint`      | Any discount given

```go
type MsgCreatePaymentContract struct {
	CreatorDid        did.Did
	PaymentTemplateId string
	PaymentContractId string
	Payer             sdk.AccAddress
	CanDeauthorise    bool
	DiscountId        sdk.Uint
}
```

## MsgCreateSubscription 

This message creates and stores the subscription at appropriate indexes.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| CreatorDid        | `string`         | Did of the user
| SubscriptionId    | `string`         | ID for the subscription
| PaymentContractId | `string`         | ID for the paymentContract
| MaxPeriods        | `sdk.AccAddress` | Maximum number of times chargeable
| Period            | `bool`           | IF the periods are allowed or not

```go
type MsgCreateSubscription struct {
	CreatorDid        did.Did
	SubscriptionId    string
	PaymentContractId string
	MaxPeriods        sdk.Uint
	Period            Period
}
```

## MsgSetPaymentContractAuthorisation

This message authorises or deauthorises a payment contract, in order to enable or disable effecting of the payment contract.

| **Field**         | **Type** | **Description** |
|:------------------|:---------|:----------------|
| PayerDid          | `string` | Payer's DID
| PaymentContractId | `string` | ID of the paymentContract
| Authorised        | `bool`   | Status of authorisation

```go
type MsgSetPaymentContractAuthorisation struct {
	PayerDid          did.Did
	PaymentContractId string
	Authorised        bool
}
``` 

## MsgGrantDiscount

This message grants a discount to a recipient for a particular payment contract.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| SenderDid         | `string`         | Initiator DID 
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

## MsgRevokeDiscount 

This message revokes a discount of a discount holder for a particular payment contract.

| **Field**         | **Type**         | **Description** |
|:------------------|:-----------------|:----------------|
| SenderDid         | `string`         | Who send the transaction 
| PaymentContractId | `string`         | ID of the payment_Contract
| Holder            | `sdk.AccAddress` | Address of who's holds the discount

```go
type MsgRevokeDiscount struct {
	SenderDid         did.Did
	PaymentContractId string
	Holder            sdk.AccAddress
}
```
