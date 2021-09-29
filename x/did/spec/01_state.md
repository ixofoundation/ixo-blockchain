# State

- [Context](#context)
- [DID Docs](#did-docs)
- [IxoDid](#ixodid)
- [IxoDid to Address and Signing Standard Cosmos Transactions
  ](#ixodid-to-address-and-signing-standard-cosmos-transactions)
- [Signing Custom ixo Transactions](#signing-custom-ixo-transactions)

## Context

In the Cosmos SDK, an account designates a pair of public key `PubKey` and
private key `PrivKey`. The `PubKey` can be derived to generate
various `Addresses`, which are used to identify users (among other parties) in
the application. `Addresses` are also associated with messages to identify the
sender of the message. The `PrivKey` is used to generate digital signatures to
prove that an `Address` associated with the `PrivKey` approved of a given
message.

For HD key derivation the Cosmos SDK uses a standard called BIP32, which allows
users to create an HD wallet - a set of accounts derived from an initial secret
seed. A seed is usually created from a 12- or 24-word mnemonic. A single seed
can derive any number of `PrivKey`s using a one-way cryptographic function.
Then, a `PubKey` can be derived from the `PrivKey`. Naturally, the mnemonic is
the most sensitive information, as private keys can always be re-generated if
the mnemonic is preserved.

The principal way of authenticating a user is done using digital signatures.
Users sign transactions using their own private key. Signature verification is
done with the associated public key.

In Cosmos SDK, the SECP256k1 algorithm is used to generate the public key from
the private key. In ixo, the ED25519 algorithm is also supported. The extent and
details of this ED25519 support will be described in this document.

References:

1. Cosmos
   SDK/Basics/Accounts: https://docs.cosmos.network/master/basics/accounts.html
2. Cosmos SDK inline documentation


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
- Example 1: `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`
- Example 2: `did:sov:CYCc2xaJKrp8Yt947Nc6jd`

The `PubKey` is a base-58 encoded **ED25519** public key (specifically the
"verify key" portion of a Sovrin DID). A valid public key obeys the following
regular expression:

- `^[123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ]{43,44}$`
- Example: `7HjjYKd4SoBv26MqXp1SzmvDiouQxarBZ2ryscZLK22x`

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

## IxoDid

A full ixo DID, which follows the [Sovrin DID spec](
https://sovrin-foundation.github.io/sovrin/spec/did-method-spec-template.html)
consists of the `Did` and `VerifyKey` (described in the previous section), but
also an `EncryptionPublicKey` and a set of private fields; a `Seed`, `SignKey`,
and `EncryptionPrivateKey`.

```go
type IxoDid struct {
    Did                 string
    VerifyKey           string
    EncryptionPublicKey string
    Secret              Secret
}
```

```go
type Secret struct {
    Seed                 string
    SignKey              string
    EncryptionPrivateKey string
}
```

A full ixo DID is never stored on-chain, given that it holds secret information.
Such a DID is derived from a private phrase (12- or 24-word mnemonic) with a
one-to-one mapping. This private phrase should of course be stored somewhere
secure and only entered into trusted wallets and clients, which will then have
the necessary functionality required to derive the ixo DID.

This functionality is also implemented in this codebase for reference:

```go
func FromMnemonic(mnemonic string) (IxoDid, error) {
    seed := sha256.New()
    seed.Write([]byte(mnemonic))

    var seed32 [32]byte
    copy(seed32[:], seed.Sum(nil)[:32])

    return FromSeed(seed32)
}
```

where FromSeed is:

```go
func FromSeed(seed [32]byte) (IxoDid, error) {
publicKeyBytes, privateKeyBytes, err := ed25519Local.GenerateKey(bytes.NewReader(seed[0:32]))
    if err != nil {
        return IxoDid{}, err
    }
    publicKey := []byte(publicKeyBytes)
    privateKey := []byte(privateKeyBytes)

    signKey := base58.Encode(privateKey[:32])
    keyPairPublicKey, keyPairPrivateKey, err := naclBox.GenerateKey(bytes.NewReader(privateKey[:]))
    if err != nil {
        return IxoDid{}, err
    }

    return IxoDid{
        Did:                 DidPrefix + base58.Encode(publicKey[:16]),
        VerifyKey:           base58.Encode(publicKey),
        EncryptionPublicKey: base58.Encode(keyPairPublicKey[:]),
        Secret: Secret{
            Seed:                 hex.EncodeToString(seed[0:32]),
            SignKey:              signKey,
            EncryptionPrivateKey: base58.Encode(keyPairPrivateKey[:]),
        },
    }, nil
}
```

If an ixo DID is registered with the blockchain, its `Did` and `VerifyKey` will
be stored, forming a `BaseDidDoc` as described in the previous section.

## IxoDid to Address and Signing Standard Cosmos Transactions

By making use of ED25519, we've managed to get to an `IxoDid`. From this, we can
now derive a Cosmos SDK address that can hold coins and do anything that a
typical Cosmos SDK address can do, but also more.

Since the `IxoDid` verify key is a base58-encoded public key, parsing this to a
Cosmos SDK ED25519 PubKey object and plugging it into an `AccAddress` gives us a
Cosmos SDK address. This functionality is implemented as follows:

```go
func VerifyKeyToAddr(verifyKey string) sdk.AccAddress {
    var pubKey ed25519.PubKey
    pubKey.Key = base58.Decode(verifyKey)
    return sdk.AccAddress(pubKey.Address())
}
```

Similarly, since the `IxoDid` sign key is a base58-encoded private key, parsing
this to a Cosmos SDK ED25519 PrivKey object gives us a private key that we can
use to sign transactions. This parsing involves appending the sign key to the
verify key as follows:

```go
var privateKey ed25519.PrivKey
privateKey.Key = append(base58.Decode(id.Secret.SignKey), base58.Decode(id.VerifyKey)...)
```

This is how ixo supports signing of standard Cosmos SDK transactions using
ED25519. However, for the chain to be able to verify ED25519 signatures, the
`DefaultSigVerificationGasConsumer` which by default does not accept ED25519
signatures was replaced by the `IxoSigVerificationGasConsumer` which does accept
ED25519. This function was then plugged into the standard Cosmos AnteHandler to
augment its behaviour accordingly.

**Note**: even though ED25519 signatures are supported, there is currently no
way to create an ED25519-based key pair from the ixo CLI. This means that any
secure signing of standard Cosmos SDK transactions using ED25519 and DIDs
currently only happens through clients such as [ixo-keysafe](
https://github.com/ixofoundation/ixo-keysafe).

## Signing and Verifying Custom ixo Transactions

On top of standard Cosmos SDK messages, ixo implements its own set of custom
messages. These messages make use of a DID-focused approach for signing and 
signature verification, most of which has already been explained in the above
sections.

What distinguishes this approach from the standard Cosmos SDK approach is that
the message signer is expected to be a `Did`, rather than an `AccAddress`. For
this to be possible, custom AnteHandlers were defined specifically to handle
and verify signed custom ixo messages.

### PubKeyGetters

A typical AnteHandler needs to get the signer's pubkey to verify signatures,
charge gas fees (to the corresponding address), and for other purposes. The
default Cosmos AnteHandler fetches a signer address' pubkey from the GetPubKey()
function after querying the account from the account keeper.

In the case of custom ixo transactions, since signers are DIDs rather than
addresses, we get the DID Doc containing the pubkey from the `did` or `project`
module (depending if signer is a user or a project, respectively).

To get a pubkey from the did/project, the did/project must have been created.
But during the did/project creation, we also need the pubkeys, which we cannot
get because the did/project does not even exist yet. For this purpose, a special 
`didPubKeyGetter` and `projectPubkeyGetter` were created, which get the pubkey
from the did/project creation message itself, given that the pubkey is one of
the parameters in such messages.

- `did` module msgs are signed by did module DIDs
- `project` module msgs are signed by project module DIDs (a.k.a projects)
- `[[default]]` remaining ixo module msgs are signed by did module DIDs

A special case in the project module is the `MsgWithdrawFunds` message, which is
a project module message signed by a did module DID (instead of a project DID).
The project module PubKeyGetter deals with this inconsistency by using the did 
module's pubkey getter for MsgWithdrawFunds.

```go
defaultPubKeyGetter := did.NewDefaultPubKeyGetter(app.DidKeeper)
didPubKeyGetter := did.NewModulePubKeyGetter(app.DidKeeper)
projectPubKeyGetter := project.NewModulePubKeyGetter(app.ProjectKeeper, app.DidKeeper)
```

### AnteHandlers

Since we have parameterised pubkey getters, we can use the same default ixo 
AnteHandler (ixo.NewDefaultAnteHandler) for all three pubkey getters instead of 
having to implement three unique AnteHandlers.

```go
defaultIxoAnteHandler := ixotypes.NewDefaultAnteHandler(
    app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
    defaultPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())
didAnteHandler := ixotypes.NewDefaultAnteHandler(
    app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
    didPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())
projectAnteHandler := ixotypes.NewDefaultAnteHandler(
    app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
    projectPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())
```

In the case of project creation, besides having a custom pubkey getter, we also
have to use a custom project creation AnteHandler. Recall that one of the
purposes of getting the pubkey is to charge gas fees. So we expect the signer to
have enough tokens to pay for gas fees. Typically, these are sent to the signer
before the signer signs their first message.

However, in the case of a project, we cannot send tokens to it before its
creation since we do not know the project DID (and thus where to send the
tokens) until exactly before its creation (when project creation is done through
ixo-cellnode). The project however does have an original creator!

Thus, the gas fees in the case of project creation are instead charged to the
original creator, which is pointed out in the project doc. For this purpose, a
custom project creation AnteHandler had to be created.

```go
projectCreationAnteHandler := project.NewProjectCreationAnteHandler(
	app.AccountKeeper, app.BankKeeper, app.DidKeeper,
	encodingConfig.TxConfig.SignModeHandler(), projectPubKeyGetter)
```

### The custom default ixo AnteHandler

The custom ixo AnteHandler is very similar to the standard Cosmos AnteHandler:

```go
sdk.ChainAnteDecorators(
	ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
	ante.NewMempoolFeeDecorator(),
	ante.NewValidateBasicDecorator(),
	ante.NewValidateMemoDecorator(ak),
	ante.NewConsumeGasForTxSizeDecorator(ak),
	NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
	ante.NewValidateSigCountDecorator(ak),
	NewDeductFeeDecorator(ak, bk, pubKeyGetter),
	NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
	NewSigVerificationDecorator(ak, signModeHandler, pubKeyGetter),
	NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
)
```

It is clear however that our custom AnteHandler is not completely custom. It
uses various functions from the Cosmos ante module. However, it also uses
customised decorators, without adding completely new decorators. Below we
present the differences in the customised decorators:

In general:

- Enforces messages to be of type `IxoMsg`, to be used with `pubKeyGetter`.
- Does not allow for multiple messages (to be added in the future).
- Does not allow for multiple signatures (to be added in the future).

`NewSetPubKeyDecorator`: as opposed to the Cosmos version...

- Gets signer pubkey from `pubKeyGetter` argument instead of tx signatures.
- Gets signer address from `pubkey` instead of the messages' `GetSigners()`.
- Uses `simEd25519Pubkey` instead of `simSecp256k1Pubkey` for simulations.

`NewDeductFeeDecorator`:

- Gets fee payer address from the `pubkey` obtained from `pubKeyGetter` instead
  of from the first message's `GetSigners()` function.

`NewSigGasConsumeDecorator`:

- Gets the only signer address from the `pubkey` obtained from `pubKeyGetter`
  instead of from the messages' `GetSigners()` function.
- Uses `simEd25519Pubkey` instead of `simSecp256k1Pubkey` for simulations.

`NewSigVerificationDecorator`:

- Gets the only signer address and account from the `pubkey` obtained from
  `pubKeyGetter` instead of from the messages' `GetSigners()` function.

`NewIncrementSequenceDecorator`:

- Gets the only signer address from the `pubkey` obtained from `pubKeyGetter`
  instead of from the messages' `GetSigners()` function.

### Signing Custom ixo Transactions

**Note**: even though ED25519 signatures are supported, there is currently no
way to securely store an ED25519-based key pair from the ixo CLI. This means
that any secure signing of custom ixo transactions using ED25519 and DIDs should 
happen only through clients such as [ixo-keysafe](
https://github.com/ixofoundation/ixo-keysafe).
