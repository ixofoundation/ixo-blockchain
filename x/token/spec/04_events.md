# Events

The token module emits the following typed events:

### TokenCreatedEvent

Emitted after a successfull `MsgCreateToken`

```go
type TokenCreatedEvent struct {
	Token *Token
}
```

The field's descriptions is as follows:

- `token` - the full [Token](02_state.md#token)

### TokenUpdatedEvent

Emitted after a successfull `MsgMintToken`, `RetireToken`, `CancelToken`, `PauseToken`, `StopToken`

```go
type TokenUpdatedEvent struct {
	Token *Token
}
```

The field's descriptions is as follows:

- `token` - the full [Token](02_state.md#token)

### TokenMintedEvent

Emitted for every batch in `MintBatch` after a successfull `MsgMintToken`

```go
type TokenMintedEvent struct {
	ContractAddress string
	Minter          string
	Owner           string
	Amount          github_com_cosmos_cosmos_sdk_types.Uint
	TokenProperties *TokenProperties
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter
- `contractAddress` - a string containing the contract address, same as the `ContractAddress` of the [Token](#token)
- `owner` - a string containing the address of the owner to mint the tokens for
- `amount` - a integer indicating how many tokens has been minted for the specific id in `TokenProperties`
- `tokenProperties` - the [TokenProperties](02_state.md#tokenproperties) that was just created for the minted tokens

### TokenTransferredEvent

Emitted after a successfull `TransferToken`

```go
type TokenTransferredEvent struct {
	Owner     string
	Recipient string
	Tokens    []*TokenBatch
}
```

The field's descriptions is as follows:

- `owner` - a string containing the address of the owner the tokens was transferred from
- `recipient` - a string containing the address of the new owner the tokens was transferred to
- `tokens` - a list of [TokenBatch](#tokenbatch). All the tokens that was trasnferred

### TokenCancelledEvent

Emitted after a successfull `MsgCancelToken`

```go
type TokenCancelledEvent struct {
	Owner     string
	Tokens    []*TokenBatch
}
```

The field's descriptions is as follows:

- `owner` - a string containing the address of the owner who cancelled the tokens
- `tokens` - a list of [TokenBatch](#tokenbatch). All the tokens that was cancelled

### TokenRetiredEvent

Emitted after a successfull `MsgRetireToken`

```go
type TokenRetiredEvent struct {
	Owner     string
	Tokens    []*TokenBatch
}
```

The field's descriptions is as follows:

- `owner` - a string containing the address of the owner who retired the tokens
- `tokens` - a list of [TokenBatch](#tokenbatch). All the tokens that was retired

### TokenPausedEvent

Emitted after a successfull `MsgPauseToken`

```go
type TokenPausedEvent struct {
	Minter          string
	ContractAddress string
	Paused          bool
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter
- `contractAddress` - a string containing the contract address
- `paused` - a boolean indicating whether token minting has been temporarily stopped or not, also the field `Paused` on [Token](#token)

### TokenStoppedEvent

Emitted after a successfull `MsgStopToken`

```go
type TokenStoppedEvent struct {
	Minter          string
	ContractAddress string
	Stopped          bool
}
```

The field's descriptions is as follows:

- `minter` - a string containing the address of the minter
- `contractAddress` - a string containing the contract address
- `stopped` - a boolean indicating whether token minting has been permanently stopped, also the field `Stopped` on [Token](#token)
