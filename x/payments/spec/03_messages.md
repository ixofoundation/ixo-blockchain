# Messages

In this section we describe the processing of the payment messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreatePaymentTemplate

Bonds can be created by any address using `MsgAddDid`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-----------------  |:--------------------------------------------------------------------------------------------------------------|
| CreatorDid             | `did.DID`          | 
| PaymentTemplate        | `PaymentTemplate`  | 

```go
type MsgCreatePaymentTemplate struct {
	CreatorDid      did.Did         `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}
```

This message creates and stores the `reatePaymentTemplate` object at appropriate indexes. Note that the sanity rate and sanity margin percentage are only used in the case of the `swapper_function`, but no error is raised if these are set for other function types.

## MsgCreatePaymentContract 

The owner of a bond can edit some of the bond's parameters using `MsgAddCredential`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| CreatorDid             | `did.Did`          | The bond to be edited 
| PaymentTemplateId      | `string`           | 
| PaymentContractId      | `string`           | 
| Payer                  | `sdk.AccAddress`   | 
| CanDeauthorise         | `bool`             | 
| DiscountId             | ` sdk.Uint`        | 

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
This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list

## MsgCreateSubscription 

The owner of a bond can edit some of the bond's parameters using `MsgAddCredential`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| CreatorDid             | `did.Did`          | edited 
| SubscriptionId         | `string`           | 
| PaymentContractId      | `string`           | 
| MaxPeriods             | `sdk.AccAddress`   | 
| Period                 | `bool`             | 


```go
type MsgCreateSubscription struct {
	CreatorDid        did.Did  `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string   `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string   `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint `json:"max_periods" yaml:"max_periods"`
	Period            Period   `json:"period" yaml:"period"`
}
```
This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list

## MsgSetPaymentContractAuthorisation 

The owner of a bond can edit some of the bond's parameters using `MsgSetPaymentContractAuthorisation`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| PayerDid               | `did.Did`          | edited 
| PaymentContractId      | `string`           | 
| Authorised             | `bool`             | 


```go
type MsgSetPaymentContractAuthorisation struct {
	PayerDid          did.Did `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string  `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool    `json:"authorised" yaml:"authorised"`
}
```
This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list

## MsgGrantDiscount 

The owner of a bond can edit some of the bond's parameters using `MsgSetPaymentContractAuthorisation`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| SenderDid              | `did.Did`          | edited 
| PaymentContractId      | `string`           | 
| DiscountId             | `sdk.Uint`         | 
| Recipient              | `sdk.AccAddress`   |


```go
type MsgGrantDiscount struct {
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress `json:"recipient" yaml:"recipient"`
}
```
This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list

## MsgRevokeDiscount 

The owner of a bond can edit some of the bond's parameters using `MsgSetPaymentContractAuthorisation`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| SenderDid              | `did.Did`          | edited 
| PaymentContractId      | `string`           | 
| Holder                 | `sdk.AccAddress`   | 



```go
type MsgRevokeDiscount struct {
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress `json:"holder" yaml:"holder"`
}

```
This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list

