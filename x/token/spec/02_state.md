# State

## Tokens

A Token is stored in the state and is accessed by the concatenation of the `minter + contract_address` as each `Token` creation instantiates a new contract with contract code form `Ixo1155ContractCode`.

- Tokens: `0x01 | minter + contract_address -> ProtocolBuffer(Token)`

## TokenProperties

A TokenProperties is stored in the state and is accessed by the identity of the TokenProperties(user provided).

- Claims: `0x02 | tokenPropertiesId -> ProtocolBuffer(TokenProperties)`

# Types

### Token

A Token is the base Token associated with a specific Minter, Contract Address and Token name (namespace)

```go
type Token struct {
	Minter            string
	ContractAddress   string
	Class             string
	Name              string
	Description       string
	Image             string
	Type              string
	Cap               github_com_cosmos_cosmos_sdk_types.Uint
	Supply            github_com_cosmos_cosmos_sdk_types.Uint
	Paused            bool
	Stopped           bool
	Retired           []*TokensRetired
	Cancelled         []*TokensCancelled
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter
- `contractAddress` - a string containing the address of the smart contract(cw1155) that was instantiated on creation of the `Token` and is used to mint the 1155 tokens.
- `class` - a string containing the token protocol entity DID (validated)
- `name` - a string containing the token name, which must be unique (namespace)
- `description` - a string containing any arbitrary description
- `image` - a string containing the image url for the token
- `type` - a string containing the token type eg. ixo1155
- `cap` - a integer containing the maximum number of tokens with this name that can be minted, 0 is unlimited
- `supply` - a integer containing how much has already been minted for this Token name
- `paused` - a boolean indicating wheter to stop allowance of token minting temporarily
- `stopped` - a boolean indicating wheter to stop allowance of token minting permanently
- `retired` - a list of [EntityAccount](#entityaccount)
- `cancelled` - a list of [EntityAccount](#entityaccount)

### TokensRetired

A TokensCancelled stores information about the retired tokens that was burned on the ixo1155 smart contract with the address at `ContractAddress` of the [Token](#token)

```go
type TokensRetired struct {
	Id           string
	Reason       string
	Jurisdiction string
	Amount       github_com_cosmos_cosmos_sdk_types.Uint
	Owner        string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the token ([TokenProperties](#tokenproperties))
- `reason` - a string containing any arbitrary reason that specifies the reason for retiring tokens.
- `jurisdiction` - a string containing the jurisdiction of the token owner. A jurisdiction has the format: <country-code>[-<sub-national-code>[ <postal-code>]] The country-code must be 2 alphabetic characters, the sub-national-code can be 1-3 alphanumeric characters, and the postal-code can be up to 64 alphanumeric characters. Only the country-code is required, while the sub-national-code and postal-code are optional and can be added for increased precision.
- `amount` - a integer indicating how many tokens of the token with `id` has been retired.
- `owner` - a string containing the cosmos address of the owner who did the retiring.

### TokensCancelled

A TokensCancelled stores information about the cancelled tokens that was burned on the ixo1155 smart contract with the address at `ContractAddress` of the [Token](#token)

```go
type TokensCancelled struct {
	Id     string
	Reason string
	Amount github_com_cosmos_cosmos_sdk_types.Uint
	Owner  string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the token ([TokenProperties](#tokenproperties))
- `reason` - a string containing any arbitrary reason that specifies the reason for cancelling tokens.
- `amount` - a integer indicating how many tokens of the token with `id` has been cancelled.
- `owner` - a string containing the cosmos address of the owner who did the cancelling.

### TokenProperties

A TokenProperties stores information about a specific minted token that will be minted on the ixo1155 smart contract with the address at `ContractAddress` of the [Token](#token)

```go
type TokenProperties struct {
	Id          string
	Index       string
	Name        string
	Collection  string
	TokenData   []*TokenData
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the token
- `index` - a string containing the unique identifier hexstring that identifies the token
- `name` - a string containing the token name, which is same as [Token](#token) `Name`
- `collection` - a string containing the did of the "collection" the token was minted towards (eg. Supamoto Malawi Did)
- `tokenData` - a list of [TokenData](#tokendata)

### TokenData

A TokenData stores information that is the `linkedResources` added to `tokenMetadata` when queried

```go
type TokenData struct {
	Uri       string
	Encrypted bool
	Proof     string
	Type      string
	Id        string
}
```

The field's descriptions is as follows:

- `uri` - a string containing the uri where the resource can be fetched (eg. a credential link abc123.ipfs)
- `encrypted` - a boolean indicating whether the resource data has been encrypted or not
- `proof` - a string containing the proof to verify the resource if a proof exists
- `type` - a string containing the resource type, media type value should always be `application/json`
- `id` - a string containing the did of entity to map token to, can be empty!

## Authz Types

### MintAuthorization

A MintAuthorization is an authz authorization that can be granted to allow the grantee to mint tokens for the specified [Token](02_state.md#token) with its constraints

```go
type WithdrawPaymentAuthorization struct {
	Minter         string
	Constraints   []*MintConstraints
}
```

The field's descriptions is as follows:

- `minter` - a string containing the account address defined in the [Token](02_state.md#token) `Minter` field
- `constraints` - a list of [MintConstraints](#mintconstraints)

### MintConstraints

A MintConstraints stores information about authorization given to make a withdrawal payment through a [WithdrawPaymentAuthorization](#withdrawpaymentauthorization)

```go
type MintConstraints struct {
	ContractAddress string
	Amount          github_com_cosmos_cosmos_sdk_types.Uint
	Name            string
	Index           string
	Collection      string
	TokenData       []*TokenData
}
```

The field's descriptions is as follows:

- `claimId` - a string containing the `id` of the claim the withdrawal is for
- `inputs` - a list of cosmos defined `Input` to pass to the the multisend tx to run to withdraw payment
- `outputs` - a list of cosmos defined `Output` to pass to the the multisend tx to run to withdraw payment
- `paymentType` - a [PaymentType](02_state.md#paymenttype)
- `contract_1155Payment` - a [Contract1155Payment](02_state.md#contract1155payment)
- `toAddress` - a string containing the account address to make the payment to
- `fromAddress` - a string containing the account address to make the payment from
- `releaseDate` - a timestamp of the date that grantee can execute authorization to make the withdrawal payment, calculated from created date plus the timeout on [Collection](02_state.md#collection) `Payments`
- `adminAddress` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
