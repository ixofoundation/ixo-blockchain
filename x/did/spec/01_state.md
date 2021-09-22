# State

## DID Docs

```go
type DidDoc interface {
    proto.Message

    SetDid(did Did) error
    GetDid() Did
    SetPubKey(pubkey string) error
    GetPubKey() string
    Address() sdk.AccAddress
}
```

The DidDoc interface defines an interface for decentralised identifiers, which
are designed to enable a verifiable and decentralised digital identity.

Currently, the only implementation of the DidDoc interface in ixo-blockchain is
BaseDidDoc, which stores a DID by which the DidDoc is accessed, a corresponding
public key, and a list of DID credentials.

```go
type BaseDidDoc struct {
    Did         string
    PubKey      string
    Credentials []*DidCredential
}
```

The `Did` part of the `BaseDidDoc` follows the [Sovrin DID spec](
https://sovrin-foundation.github.io/sovrin/spec/did-method-spec-template.html)
but allows for a `did:ixo:` prefix as well as a `did:sov:` prefix, as defined by
the following regular expression:

- `^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}([/][a-zA-Z0-9:]+|)$`

The `PubKey` is a base-58 encoded ED25519 public key (specifically the
"verify key" portion of a Sovrin DID). A valid public key obeys the following
regular expression:

- `^[123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ]{43,44}$`

The DID is the first 16 bytes of the public key base-58 encoded. So in order to
convert the above `PubKey` to the `Did`, the `PubKey` needs to be decoded into
bytes, and the first 16 bytes need to then be base-58 encoded. A function that
does exactly this is implemented in this module:

```go
func UnprefixedDidFromPubKey(pubKey string) string {
	// Assumes that PubKey is valid (check IsValidPubKey regex)
	// Since result is not prefixed (did:ixo:), string returned rather than DID
	pubKeyBz := base58.Decode(pubKey)
	return base58.Encode(pubKeyBz[:16])
}
```

The credentials associated with a DID are defined via the structs below and
contain a claim indicating whether or not the DID is KYC validated. Credentials
can be extended to describe a DID's status and characteristics in more detail.

```go
type DidCredential struct {
    CredType []string
    Issuer   string
    Issued   string
    Claim    *Claim
}
```

```go
type Claim struct {
    Id           string
    KYCValidated bool
}
```
