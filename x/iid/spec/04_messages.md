 # Messages

In this section we describe the processing of the staking messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state_transitions.md) section.


### Verification 

A verification message represent a combination of a verification method and a set of verification relationships. It has the following fields:

- `relationships` - a list of strings identifying the verification relationship for the verification method
- `method` - a [verification method object](02_state.md#verification_method) 
- `context` - a list of strings identifying additional [json ld contexts](https://json-ld.org/spec/latest/json-ld/#the-context)


#### Source 
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L32



### MsgCreateIidDocument

A `MsgCreateIidDocument` is used to create a new iid document, sit has the following fields

- `id` - the iid string identifying the iid document
- `controller` - a list of did that are controllers of the iid document
- `verifications` - a list of [verification](04_messages.md#verification) for the iid document
- `services` - a list of [services](02_state.md#service) for the iid document
- `rights` - a list of [Accorded Rights](02_state.md#Accorded Right) for the iid document
- `resources` - a list of [Linked Resource](02_state.md#Linked Resource) for the iid document
- `entities` - a list of [Linked Entity](02_state.md#Linked Entity) for the iid document
- `controllers` - a list of [Controller](02_state.md#Controller) for the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source

https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L64

### MsgUpdateIidDocument

The `MsgUpdateIidDocument` is used to update a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `controller` - a list of iid's that are controllers of the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L58
### MsgAddVerification

The `MsgAddVerification` is used to add new [verification methods](https://w3c.github.io/did-core/#verification-methods) and [verification relationships](https://w3c.github.io/did-core/#verification-relationships) to a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `verification` - the [verification](04_messages.md#verification) to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L73

### MsgSetVerificationRelationships

The `MsgSetVerificationRelationships` is used to overwrite the [verification relationships](https://w3c.github.io/did-core/#verification-relationships) for a [verification methods](https://w3c.github.io/did-core/#verification-methods) of a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `method_id` - a string containing the unique identifier of the verification method within the iid document.
- `relationships` - a list of strings identifying the verification relationship for the verification method
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L84
### MsgRevokeVerification

The `MsgRevokeVerification` is used to remove a [verification method](https://w3c.github.io/did-core/#verification-methods) and related [verification relationships](https://w3c.github.io/did-core/#verification-relationships) from a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `method_id` - a string containing the unique identifier of the verification method within the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L96
### MsgAddService

The `MsgAddService` is used to add a [service](https://w3c.github.io/did-core/#services) to a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `service_data` - the [service](02_state.md#service) object to add to the iid document 
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L111
### MsgDeleteService

The `MsgDeleteService` is used to remove a [service](https://w3c.github.io/did-core/#services) from a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `service_id` - the unique id of the [service](02_state.md#service) in the iid document 
- `signer` - a string containing the cosmos address of the private key signing the transaction 

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L122

### MsgAddAccordedRight

The `MsgAddAccordedRight` is used to add an [AccordedRight](#) to a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `accordedRight ` - the [AccordedRight ](02_state.md#Accorded Right ) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L223
### MsgDeleteAccordedRight

The `MsgDeleteAccordedRight` is used to remove a [AccordedRight](#) from a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `right_id ` - the unique id of the [AccordedRight](02_state.md#Accorded Right) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L231

### MsgAddLinkedEntity

The `MsgAddLinkedEntity` is used to add a [LinkedEntity](#) to a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `linkedEntity ` - the [LinkedEntity](02_state.md#Linked Entity) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L205
### MsgDeleteLinkedEntity

The `MsgDeleteLinkedEntity` is used to remove a [LinkedEntity](#) from a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `entity_id ` - the unique id of the [LinkedEntity](02_state.md#Linked Entity) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L214

### MsgAddLinkedResource

The `MsgAddLinkedResource` is used to add a [LinkedResource](#) to a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `linkedResource` - the [LinkedResource](02_state.md#Linked Resource) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L187
### MsgDeleteLinkedResource

The `MsgDeleteLinkedResource` is used to remove a [LinkedResource](#) from a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `resource_id` - the unique id of the [LinkedResource](02_state.md#Linked Resource) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L196

### MsgAddIidContext

The `MsgAddIidContext` is used to add a [service](02_state.md#Iid Context) to a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `context ` - the [IidContext](02_state.md#Iid Context) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L240
### MsgDeleteIidContext

The `MsgDeleteIidContext` is used to remove a [IidContext](02_state.md#Iid Context) from a iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `contextKey ` - the unique id of the [IidContext](02_state.md#Iid Context) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L248
### MsgDeleteIidContext

The `MsgUpdateIidMeta ` is used to REPLACE the [IidMetadata](02_state.md#IidMetadata) in an iid document. It has the following fields:

- `id` - the iid string identifying the iid document
- `meta` - the [IidMetadata](02_state.md#IidMetadata) object to replace the existing metadata on an iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

#### Source:
https://github.com/ixofoundation/ixo-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L257

### QueryIidDocumentRequest

The `QueryIidDocumentRequest` is used to resolve a iid document. That is, to retrieve  a iid document from its id. It has the following fields:

- `id` - the iid string identifying the iid document

#### Source: 
https://github.com/allinbits/cosmos-cash/blob/v1.0.0/proto/did/query.proto#L45