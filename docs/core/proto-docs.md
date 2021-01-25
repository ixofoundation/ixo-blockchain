# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [did/did.proto](#did/did.proto)
    - [Claim](#did.Claim)
    - [DidCredential](#did.DidCredential)
    - [IxoDid](#did.IxoDid)
    - [Secret](#did.Secret)
  
- [did/genesis.proto](#did/genesis.proto)
    - [GenesisState](#did.GenesisState)
  
- [did/query.proto](#did/query.proto)
    - [QueryAddressFromBase58EncodedPubkeyRequest](#did.QueryAddressFromBase58EncodedPubkeyRequest)
    - [QueryAddressFromBase58EncodedPubkeyResponse](#did.QueryAddressFromBase58EncodedPubkeyResponse)
    - [QueryAddressFromDidRequest](#did.QueryAddressFromDidRequest)
    - [QueryAddressFromDidResponse](#did.QueryAddressFromDidResponse)
    - [QueryAllDidDocsRequest](#did.QueryAllDidDocsRequest)
    - [QueryAllDidDocsResponse](#did.QueryAllDidDocsResponse)
    - [QueryAllDidsRequest](#did.QueryAllDidsRequest)
    - [QueryAllDidsResponse](#did.QueryAllDidsResponse)
    - [QueryDidDocRequest](#did.QueryDidDocRequest)
    - [QueryDidDocResponse](#did.QueryDidDocResponse)
    - [QueryIxoDidFromMnemonicRequest](#did.QueryIxoDidFromMnemonicRequest)
    - [QueryIxoDidFromMnemonicResponse](#did.QueryIxoDidFromMnemonicResponse)
  
    - [Query](#did.Query)
  
- [did/tx.proto](#did/tx.proto)
    - [MsgAddCredential](#did.MsgAddCredential)
    - [MsgAddCredentialResponse](#did.MsgAddCredentialResponse)
    - [MsgAddDid](#did.MsgAddDid)
    - [MsgAddDidResponse](#did.MsgAddDidResponse)
  
    - [Msg](#did.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="did/did.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/did.proto



<a name="did.Claim"></a>

### Claim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| KYCvalidated | [bool](#bool) |  |  |






<a name="did.DidCredential"></a>

### DidCredential



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credtype | [string](#string) | repeated |  |
| issuer | [string](#string) |  |  |
| issued | [string](#string) |  |  |
| claim | [Claim](#did.Claim) |  |  |






<a name="did.IxoDid"></a>

### IxoDid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| verifyKey | [string](#string) |  |  |
| encryptionPublicKey | [string](#string) |  |  |
| secret | [Secret](#did.Secret) |  |  |






<a name="did.Secret"></a>

### Secret



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seed | [string](#string) |  |  |
| signKey | [string](#string) |  |  |
| encryptionPrivateKey | [string](#string) |  |  |





 

 

 

 



<a name="did/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/genesis.proto



<a name="did.GenesisState"></a>

### GenesisState
GenesisState defines the did module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddocs | [google.protobuf.Any](#google.protobuf.Any) | repeated | DidDoc is an interface to we use Any here, like evidence GenesisState |





 

 

 

 



<a name="did/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/query.proto



<a name="did.QueryAddressFromBase58EncodedPubkeyRequest"></a>

### QueryAddressFromBase58EncodedPubkeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubKey | [string](#string) |  | pubKey defines the PubKey for the requested address |






<a name="did.QueryAddressFromBase58EncodedPubkeyResponse"></a>

### QueryAddressFromBase58EncodedPubkeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address returns the address for a given PubKey |






<a name="did.QueryAddressFromDidRequest"></a>

### QueryAddressFromDidRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  | did defines the DID for the requested address |






<a name="did.QueryAddressFromDidResponse"></a>

### QueryAddressFromDidResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address returns the address for a given DID |






<a name="did.QueryAllDidDocsRequest"></a>

### QueryAllDidDocsRequest
no input needed






<a name="did.QueryAllDidDocsResponse"></a>

### QueryAllDidDocsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddocs | [google.protobuf.Any](#google.protobuf.Any) | repeated | diddocs returns a list of all DidDocs |






<a name="did.QueryAllDidsRequest"></a>

### QueryAllDidsRequest
no input needed






<a name="did.QueryAllDidsResponse"></a>

### QueryAllDidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dids | [string](#string) | repeated | dids returns a list of all DIDs |






<a name="did.QueryDidDocRequest"></a>

### QueryDidDocRequest
Request/response types from old x/did/client/cli/query.go and x/did/client/rest/query.go


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  | did defines the DID for the requested DidDoc |






<a name="did.QueryDidDocResponse"></a>

### QueryDidDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddoc | [google.protobuf.Any](#google.protobuf.Any) |  | diddoc returns the requested DidDoc |






<a name="did.QueryIxoDidFromMnemonicRequest"></a>

### QueryIxoDidFromMnemonicRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mnemonic | [string](#string) |  | mnemonic defines the 12-word secret mnemonic for a given ixo DID |






<a name="did.QueryIxoDidFromMnemonicResponse"></a>

### QueryIxoDidFromMnemonicResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixodid | [IxoDid](#did.IxoDid) |  | ixodid returns the IxoDid for a given 12-word mnemonic |





 

 

 


<a name="did.Query"></a>

### Query
To get a list of all module queries, go to your module&#39;s keeper/querier.go and check all cases in NewQuerier().
REST endpoints taken from previous did/client/rest/query.go

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DidDoc | [QueryDidDocRequest](#did.QueryDidDocRequest) | [QueryDidDocResponse](#did.QueryDidDocResponse) |  |
| AllDids | [QueryAllDidsRequest](#did.QueryAllDidsRequest) | [QueryAllDidsResponse](#did.QueryAllDidsResponse) |  |
| AllDidDocs | [QueryAllDidDocsRequest](#did.QueryAllDidDocsRequest) | [QueryAllDidDocsResponse](#did.QueryAllDidDocsResponse) |  |
| AddressFromDid | [QueryAddressFromDidRequest](#did.QueryAddressFromDidRequest) | [QueryAddressFromDidResponse](#did.QueryAddressFromDidResponse) |  |
| AddressFromBase58EncodedPubkey | [QueryAddressFromBase58EncodedPubkeyRequest](#did.QueryAddressFromBase58EncodedPubkeyRequest) | [QueryAddressFromBase58EncodedPubkeyResponse](#did.QueryAddressFromBase58EncodedPubkeyResponse) |  |
| IxoDidFromMnemonic | [QueryIxoDidFromMnemonicRequest](#did.QueryIxoDidFromMnemonicRequest) | [QueryIxoDidFromMnemonicResponse](#did.QueryIxoDidFromMnemonicResponse) |  |

 



<a name="did/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/tx.proto



<a name="did.MsgAddCredential"></a>

### MsgAddCredential



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| didCredential | [DidCredential](#did.DidCredential) |  |  |






<a name="did.MsgAddCredentialResponse"></a>

### MsgAddCredentialResponse







<a name="did.MsgAddDid"></a>

### MsgAddDid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| pubKey | [string](#string) |  |  |






<a name="did.MsgAddDidResponse"></a>

### MsgAddDidResponse






 

 

 


<a name="did.Msg"></a>

### Msg
To get a list of all module messages, go to your module&#39;s handler.go and check all cases in NewHandler().

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddDid | [MsgAddDid](#did.MsgAddDid) | [MsgAddDidResponse](#did.MsgAddDidResponse) |  |
| AddCredential | [MsgAddCredential](#did.MsgAddCredential) | [MsgAddCredentialResponse](#did.MsgAddCredentialResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

