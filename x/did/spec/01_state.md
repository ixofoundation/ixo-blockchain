# State

## DidDocs

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

The DidDoc interface defines an interface for decentralised identifiers, which are
designed to enable a verifiable and decentralised digital identity. 

Currently, the only implementation of the DidDoc interface in ixo-blockchain is
BaseDidDoc, which stores a DID by which the DidDoc is accessed, a corresponding 
public key, and a list of DID credentials. TODO

```go
type BaseDidDoc struct {
    Did         string
    PubKey      string
    Credentials []*DidCredential
}
```

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
