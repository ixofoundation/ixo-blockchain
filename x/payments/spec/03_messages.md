# Messages

In this section we describe the processing of the payment messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreatePaymentTemplate

PaymentTemplate can be created by any address using `MsgCreatePaymentTemplate`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-----------------  |:--------------------------------------------------------------------------------------------------------------|
| CreatorDid             | `string`          | 
| PaymentTemplate        | `string`  | 

```go
type MsgCreatePaymentTemplate struct {
	CreatorDid      did.Did         `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}
```

This message creates and stores the `MsgCreatePaymentTemplate` object at appropriate indexes. 

## MsgCreatePaymentContract 

This will create a PaymentContract using `MsgCreatePaymentContract`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| CreatorDid             | `string`          | Did of the user
| PaymentTemplateId      | `string`           | ID of the paymentTemplate
| PaymentContractId      | `string`           | ID of the PaymentContract
| Payer                  | `sdk.AccAddress`   | Address of the payer
| CanDeauthorise         | `bool`             | Bool of de_authorise
| DiscountId             | ` sdk.Uint`        | Any discount given

```go
type MsgCreatePaymentContract struct {
	CreatorDid        did.Did        `json:"creator_did" yaml:"creator_did"`
	PaymentTemplateId string         `json:"payment_template_id" yaml:"payment_template_id"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Payer             sdk.AccAddress `json:"payer" yaml:"payer"`
	CanDeauthorise    bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
}
```

## MsgCreateSubscription 

The owner of a bond can edit some of the bond's parameters using `MsgCreateSubscription`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| CreatorDid             | `string`          | Did of the user 
| SubscriptionId         | `string`           | ID for the subscription
| PaymentContractId      | `string`           | ID for the paymentContract
| MaxPeriods             | `sdk.AccAddress`   | 
| Period                 | `bool`             | IF the periods are allowed or not


```go
type MsgCreateSubscription struct {
	CreatorDid        did.Did  `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string   `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string   `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint `json:"max_periods" yaml:"max_periods"`
	Period            Period   `json:"period" yaml:"period"`
}
```

## MsgSetPaymentContractAuthorisation 



| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| PayerDid               | `string`          | Payer's DID 
| PaymentContractId      | `string`           | ID of the paymentContract
| Authorised             | `bool`             | Status of authorisation


```go
type MsgSetPaymentContractAuthorisation struct {
	PayerDid          did.Did `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string  `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool    `json:"authorised" yaml:"authorised"`
}
```

## MsgGrantDiscount 


| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| SenderDid              | `string`          | Initiator DID 
| PaymentContractId      | `string`           | ID for the paymentContract
| DiscountId             | `sdk.Uint`         | How much is the discount
| Recipient              | `sdk.AccAddress`   | Who is the recipient 


```go
type MsgGrantDiscount struct {
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress `json:"recipient" yaml:"recipient"`
}
```


## MsgRevokeDiscount 

The owner of a bond can edit some of the bond's parameters using `MsgRevokeDiscount`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| SenderDid              | `string`          | Who send the transaction 
| PaymentContractId      | `string`           | ID of the payment_Contract
| Holder                 | `sdk.AccAddress`   | Address of who's holds the discount



```go
type MsgRevokeDiscount struct {
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress `json:"holder" yaml:"holder"`
}

```
