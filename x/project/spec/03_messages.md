# Messages

In this section we describe the processing of the bonds messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreateProject

Bonds can be created by any address using `MsgCreateProject`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | The denomination of the bond's tokens |
| SenderDid              | `did.Did`          | 
| ProjectDid             | `did.Did`          | 
| PubKey                 | `string`           | 
| Data                   | `json.RawMessage`  | 


```go
type MsgCreateProject struct {
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did         `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did         `json:"projectDid" yaml:"projectDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}
```

This message is expected to fail if:

- another bond with this token is already registered, the token is the staking token, or the token is not a valid denomination
- name or description is an empty string
- function type is not one of the defined function types (`power_function`, `sigmoid_function`, `swapper_function`)

This message creates and stores the `Bond` object at appropriate indexes. Note that the sanity rate and sanity margin percentage are only used in the case of the `swapper_function`, but no error is raised if these are set for other function types.

## MsgUpdateProjectStatus

The owner of a bond can edit some of the bond's parameters using `MsgEditBond`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | The denomination of the bond's tokens |
| SenderDid              | `did.Did`          | 
| ProjectDid             | `did.Did`          | 
| Data                   | `UpdateProjectStatusDoc`  | 

This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list

```go
type MsgUpdateProjectStatus struct {
	TxHash     string                 `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                `json:"projectDid" yaml:"projectDid"`
	Data       UpdateProjectStatusDoc `json:"data" yaml:"data"`
}
```

This message stores the updated `MsgUpdateProjectStatus` object.

## MsgUpdateAgent

Any address that holds tokens that a bond uses as its reserve can buy tokens from that bond in exchange for reserve tokens. Rather than performing the buy itself, the `MsgBuy` handler registers a buy order in the current orders batch and cancels any other orders that become unfulfillable. Any order in that batch gets fulfilled at the end of the batch's lifespan. The `MsgBuy` handler also locks away the `MaxPrices` value (`< Balance`) indicated by the address so that these are not used elsewhere whilst the batch is being processed.


| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | The denomination of the bond's tokens |
| SenderDid              | `did.Did`          | 
| ProjectDid             | `did.Did`          | 
| Data                   | `UpdateAgentDoc`  | 

This message is expected to fail if:
- amount is not an amount of an existing bond
- max prices is greater than the balance of the buyer
The batch-adjusted current supply in the case of buys is the current supply of the bond plus any uncancelled buy amounts in the current batch. 

```go
type MsgUpdateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       UpdateAgentDoc `json:"data" yaml:"data"`
}
```

This message adds the buy order to the current batch.

### MsgBuy for Swapper Function Bonds

In general, but especially in the case of swapper function bonds, buying tokens from a bond can be seen as adding liquidity to that bond's token. To add liquidity to a swapper function, the current exchange rate is used to determine how much of each reserve token makes up the price. Otherwise, the price is an equal number of each of the reserve tokens according to the function type.

Moreover, in the case of the swapper function, the first `MsgBuy` performed is special and plays a very important role in specifying the price of the bond token. Since we have no price reference for the first buy in a swapper function, the `MaxPrices` specified are used as the actual price, with no fees charged.

This effectively means that if the user requested `n` bond tokens with max prices `aR1` and `bR2` (for reserve tokens `R1` and `R2`), the next buyers will have to pay `(a/n)R1` and `(b/n)R2` tokens per bond token requested. Specifying high `a` and `b` prices for a small `n` (say `n=1`) means that the next buyers will have to pay at most `aR1` and `bR2` per bond token. **Thus, it is important that the first buy is well-calculated and performed carefully.**

## MsgCreateClaim

Any address that holds previously bought bond tokens can, at any point, sell the tokens back to the bond in exchange for reserve tokens. Similar to the `MsgBuy`, the `MsgSell` handler just registers a sell order in the current orders batch which then gets fulfilled at the end of the batch's lifespan.

Once the sell order is fulfilled, the number of tokens to be sold are burned on the fly and the address gets reserve tokens in return, minus the transaction and exit fees specified by the bond. The actual number of reserve tokens given to the address in return is determined from the bond function, but is also influenced by any other buys and sells in the same orders batch, as a means to prevent front-running. A sell order cannot be cancelled.

In general, but especially in the case of swapper function bonds, buying tokens from a bond can be seen as adding liquidity for that bond. To add liquidity to a swapper function, the current exchange rate is used to determine how much of each reserve token makes up the price. Otherwise, the price is an equal number of each of the reserve tokens according to the function type.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | The denomination of the bond's tokens |
| SenderDid              | `did.Did`          | 
| ProjectDid             | `did.Did`          | 
| Data                   | `CreateClaimDoc`  | 

This message is expected to fail if:
- amount is not an amount of an existing bond
- amount is greater than the balance of the seller
- amount is greater than the bond's current supply
- amount causes the bond's batch-adjusted current supply to become negative
- amount violates an order quantity limit defined by the bond

The batch-adjusted current supply in the case of sells is the current supply of the bond minus any uncancelled sell amounts in the current batch.

```go
type MsgCreateClaim struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateClaimDoc `json:"data" yaml:"data"`
}
```

This message adds the sell order to the current batch.

## MsgCreateEvaluation

Any address that holds tokens (_t1_) that a swapper function bond uses as one of its two reserves (_t1_ and _t2_) can swap the tokens in exchange for reserve tokens of the other type (_t2_). Similar to the `MsgBuy` and `MsgSell`, the `MsgSwap` handler just registers a swap order in the current orders batch which then gets fulfilled at the end of the batch's lifespan.

Once the swap order is fulfilled, 

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | The denomination of the bond's tokens |
| SenderDid              | `did.Did`          | 
| ProjectDid             | `did.Did`          | 
| Data                   | `CreateEvaluationDoc`  | 

This message is expected to fail if:
- bond does not exist or is not swapper function
- from amount is greater than the balance of the swapper
- from and to tokens are the same token
- from and to tokens are not the swapper function's reserve tokens
- from amount violates an order quantity limit defined by the bond

```go
type MsgCreateEvaluation struct {
	TxHash     string              `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did             `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did             `json:"projectDid" yaml:"projectDid"`
	Data       CreateEvaluationDoc `json:"data" yaml:"data"`
}
```

This message adds the swap order to the current batch.

## MsgWithdrawFunds

Any address that holds tokens (_t1_) that a swapper function bond uses as one of its two reserves (_t1_ and _t2_) can swap the tokens in exchange for reserve tokens of the other type (_t2_). Similar to the `MsgBuy` and `MsgSell`, the `MsgSwap` handler just registers a swap order in the current orders batch which then gets fulfilled at the end of the batch's lifespan.

Once the swap order is fulfilled, 

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| SenderDid              | `did.Did`          | 
| Data                   | `WithdrawFundsDoc`  | 

This message is expected to fail if:
- bond does not exist or is not swapper function
- from amount is greater than the balance of the swapper
- from and to tokens are the same token
- from and to tokens are not the swapper function's reserve tokens
- from amount violates an order quantity limit defined by the bond

```go
type MsgWithdrawFunds struct {
	SenderDid did.Did          `json:"senderDid" yaml:"senderDid"`
	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
}
```

This message adds the swap order to the current batch.


