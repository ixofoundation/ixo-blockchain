# Messages

In this section we describe the processing of the token messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreateToken

A `MsgCreateToken` creates and stores a new Token struct (which is a class overview for minting tokens). Sets up the Token by initiating a cw1155 smart contract, and saving the address in the Token as well as a unique name (namespace).

```go
type MsgCreateToken struct {
	Minter string
	Class DIDFragment
	Name string
	Description string
	Image string
	TokenType string
	Cap github_com_cosmos_cosmos_sdk_types.Uint
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter
- `class` - a string containing the token protocol entity DID (validated)
- `name` - a string containing the token name, which must be unique (namespace)
- `description` - a string containing any arbitrary description
- `image` - a string containing the image url for the token
- `tokenType` - a string containing the token type eg. ixo1155
- `cap` - a integer containing the maximum number of tokens with this name that can be minted, 0 is unlimited

## MsgMintToken

A `MsgMintToken` mints tokens on the ixo1155 contract that was instantiate on the [Token](#token) creation. There will also be [TokenProperties](02_state.md#tokenproperties) created for each unique id.
Allows a Minter to directly mint the token with a name that is specified in the Token, and a unique index. The minted tokens `id` is a hex encoded md5 hash of the token `name` plus `index`, with no separators.

```go
type MsgMintToken struct {
	Minter          string
	ContractAddress string
	Owner     string
	MintBatch []*MintBatch
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter
- `contractAddress` - a string containing the contract address, same as the `ContractAddress` of the [Token](#token)
- `owner` - a string containing the address of the owner to mint the tokens for
- `mintBatch` - a list of [MintBatch](#mintbatch)

## MsgTransferToken

A `MsgTransferToken` transfers tokens on the ixo1155 contract that was instantiate on the [Token](#token) creation and where the tokens was minted.

```go
type MsgTransferToken struct {
	Owner string
	Recipient string
	Tokens []*TokenBatch
}
```

The field's descriptions is as follows:

- `owner` - a string containing the address of the owner to transfer the tokens from
- `recipient` - a string containing the address of the owner to transfer the tokens to
- `tokens` - a list of [TokenBatch](#tokenbatch). All the tokens must be on the smart contract.

## MsgRetireToken

A `MsgRetireToken` burns the tokens on the ixo1155 contract that was instantiate on the [Token](#token) creation and where the tokens was minted. It also adds the retire details to the [Token](#token) `Retired` field to keep a record of the retire event. Retiring credits is permanent and cannot be undone.

```go
type MsgRetireToken struct {
	Owner string
	Tokens []*TokenBatch
	Jurisdiction string
	Reason string
}
```

The field's descriptions is as follows:

- `owner` - a string containing the address of the owner of the tokens to retire
- `tokens` - a list of [TokenBatch](#tokenbatch). All the tokens must be on the smart contract.
- `reason` - a string containing any arbitrary reason that specifies the reason for retiring tokens.
- `jurisdiction` - a string containing the jurisdiction of the token owner. A jurisdiction has the format: <country-code>[-<sub-national-code>[ <postal-code>]] The country-code must be 2 alphabetic characters, the sub-national-code can be 1-3 alphanumeric characters, and the postal-code can be up to 64 alphanumeric characters. Only the country-code is required, while the sub-national-code and postal-code are optional and can be added for increased precision.

## MsgCancelToken

A `MsgCancelToken` burns the tokens on the ixo1155 contract that was instantiate on the [Token](#token) creation and where the tokens was minted. It also adds the cancel details to the [Token](#token) `Cancelled` field to keep a record of the cancel event.
Cancels a specified amount of tradable tokens, removing the amount from the token owner's tradable balance(by burning it) and removing the amount from the token’s tradable supply. Cancelling credits is permanent and cannot be undone.

```go
type MsgCancelToken struct {
	Owner string
	Tokens []*TokenBatch
	Reason string
}
```

The field's descriptions is as follows:

- `owner` - a string containing the address of the owner of the tokens to retire
- `tokens` - a list of [TokenBatch](#tokenbatch). All the tokens must be on the smart contract.
- `reason` - a string containing any arbitrary reason that specifies the reason for retiring tokens.

## MsgPauseToken

A `MsgPauseToken` changes the `Paused` field on the [Token](#token) which stops allowance of token minting temporarily if set to `true` or allow minting again if set to `false`

```go
type MsgPauseToken struct {
	Minter          string
	ContractAddress string
	Paused bool
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter, must be same as `Minter` of the [Token](#token), must also be the signers address
- `contractAddress` - a string containing the contract address, same as the `ContractAddress` of the [Token](#token)
- `paused` - a boolean indicating whether to stop allowance of token minting temporarily

## MsgStopToken

A `MsgPauseToken` changes the `Cancelled` field on the [Token](#token) which stops allowance of token minting permanently

```go
type MsgStopToken struct {
	Minter          string
	ContractAddress string
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter, must be same as `Minter` of the [Token](#token), must also be the signers address
- `contractAddress` - a string containing the contract address, same as the `ContractAddress` of the [Token](#token)

## Generic types

### MintBatch

A MintBatch allows batching mint token info

```go
type MintBatch struct {
	Name        string
	Index       string
	Amount      github_com_cosmos_cosmos_sdk_types.Uint
	Collection  string
	TokenData   []*TokenData
}
```

The field's descriptions is as follows:

- `name` - a string containing the token name, which is same as [Token](02_state.md#token) `Name`
- `index` - a string containing the unique identifier hexstring that identifies the token
- `amount` - a integer indicating how many tokens must be minted for this unique `index` and generate `id`
- `collection` - a string containing the did of the "collection" the token was minted towards (eg. Supamoto Malawi Did)
- `tokenData` - a list of [TokenData](02_state.md#tokendata)

### TokenBatch

A TokenBatch allows batching tokens for transfering and retiring

```go
type TokenBatch struct {
	Id      string
	Amount  github_com_cosmos_cosmos_sdk_types.Uint
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the token ([TokenProperties](#tokenproperties))
- `amount` - a integer indicating how many tokens is in the batch for the specified `id`
