# Messages

In this section we describe the processing of the treasury messages and the corresponding updates to the state. The treasury module does not store any state itself. Whenever conversion from DID to address is mentioned, this is being performed as follows:

```go
func DidToAddr(did ixo.Did) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(did)))
}
```

## TreasuryMessage

A general `TreasuryMessage` interface is defind to be implemented by every message in the Treasury module. Each message must implement the `GetPubKey()` and `GetSenderDid()` functions, both of which are used only in the Treasury module's AnteHandler. The embedded sdk.Msg ensures that the Treasury messages also follow the standard Cosmos SDK message interface.

```go
type TreasuryMessage interface {
	sdk.Msg
	GetPubKey() string
	GetSenderDid() ixo.Did
}
``` 

## MsgSend

Sending of tokens between two addresses identified by DIDs and signed by the sender is done using `MsgSend`. The handler for this message converts the FromDid and ToDid to `sdk.AccAddress` and then uses the Cosmos SDK `Bank` module keeper to perform the send. This message is expected to fail only if the address to which the FromDid maps to does not have enough tokens.

| **Field**              | **Type**         | **Description**                                                                                               |
|:-----------------------|:-----------------|:--------------------------------------------------------------------------------------------------------------|
| PubKey    | string    | PubKey of the message signer |
| FromDid   | ixo.Did   | DID of the sender (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| ToDid     | ixo.Did   | DID of the recipient (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| Amount    | sdk.Coins | The tokens being sent (e.g. `100ixo,200stake`) |

```go
type MsgSend struct {
	PubKey    string
	FromDid   ixo.Did
	ToDid     ixo.Did
	Amount    sdk.Coins
}
``` 

## MsgOracleTransfer

Sending of tokens between two addresses identified by DIDs and signed by an oracle is done using `MsgOracleTransfer`. The handler for this message confirms that the oracle exists and has the required capabilities to transfer _all_ the token denominations specified in the amount, using the `Oracle` module. The rest of the handling is identical to `MsgSend`. This message is expected to fail on the same failing cases of `MsgSend` but also if the oracle does not exist or does not have the required capabilities.

| **Field**              | **Type**         | **Description**                                                                                               |
|:-----------------------|:-----------------|:--------------------------------------------------------------------------------------------------------------|
| PubKey    | string    | PubKey of the message signer |
| OracleDid | ixo.Did   | DID of the oracle (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| FromDid   | ixo.Did   | DID of the sender (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| ToDid     | ixo.Did   | DID of the recipient (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| Amount    | sdk.Coins | The tokens being sent (e.g. `100ixo,200stake`) |
| Proof     | string    | Arbitrary proof backing up this operation (presently unused) |

```go
type MsgOracleTransfer struct {
	PubKey    string
	OracleDid ixo.Did
	FromDid   ixo.Did
	ToDid     ixo.Did
	Amount    sdk.Coins
    Proof     string
}
``` 

## MsgOracleMint

Minting of tokens to an address identified by a DID and signed by an oracle is done using `MsgOracleMint`. The handler for this message confirms that the oracle exists and has the required capabilities to mint _all_ the token denominations specified in the amount, using the `Oracle` module. The handler then uses the Cosmos SDK `supply` module to mint tokens to the module account address (identified by the module name), which are then transferred to the recipient address (from the `ToDid`). This message is expected to fail if the oracle does not exist or does not have the required capabilities.

| **Field**              | **Type**         | **Description**                                                                                               |
|:-----------------------|:-----------------|:--------------------------------------------------------------------------------------------------------------|
| PubKey    | string    | PubKey of the message signer |
| OracleDid | ixo.Did   | DID of the oracle (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| ToDid     | ixo.Did   | DID of the recipient (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| Amount    | sdk.Coins | The tokens being sent (e.g. `100ixo,200stake`) |
| Proof     | string    | Arbitrary proof backing up this operation (presently unused) |

```go
type MsgOracleMint struct {
	PubKey    string
	OracleDid ixo.Did
	ToDid     ixo.Did
	Amount    sdk.Coins
    Proof     string
}
``` 

## MsgOracleBurn

Burning of tokens from an address identified by a DID and signed by an oracle is done using `MsgOracleBurn`. The handler for this message confirms that the oracle exists and has the required capabilities to burn _all_ the token denominations specified in the amount, using the `Oracle` module. The handler then transfers the tokens from the sender address (from the `FromDid`) to the module account address (identified by the module name), which are then burned using the Cosmos SDK `supply` module. This message is expected to fail if the oracle does not exist or does not have the required capabilities, but also if the sender address does not have enough tokens.

| **Field**              | **Type**         | **Description**                                                                                               |
|:-----------------------|:-----------------|:--------------------------------------------------------------------------------------------------------------|
| PubKey    | string    | PubKey of the message signer |
| OracleDid | ixo.Did   | DID of the oracle (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| FromDid   | ixo.Did   | DID of the sender (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`) |
| Amount    | sdk.Coins | The tokens being sent (e.g. `100ixo,200stake`) |
| Proof     | string    | Arbitrary proof backing up this operation (presently unused) |

```go
type MsgOracleBurn struct {
	PubKey    string
	OracleDid ixo.Did
	FromDid   ixo.Did
	Amount    sdk.Coins
    Proof     string
}
``` 