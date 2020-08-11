# Events

The did module emits the following events:

## Handlers

## MsgAddDidDoc

| Type               | Attribute Key | Attribute Value |
|--------------------|---------------|-----------------|
| EventTypeAddDidDoc | did           | {did}           |
| EventTypeAddDidDoc | pub_key       | {pub_key}       |

## MsgAddCredential

| Type                   | Attribute Key | Attribute Value |
|------------------------|---------------|-----------------|
| EventTypeAddCredential | cred_type     | {credType}      |
| EventTypeAddCredential | issuer        | {issuer}        |
| EventTypeAddCredential | issued        | {issued}        |
| EventTypeAddCredential | claim         | {claim}         |
| EventTypeAddCredential | true          | {bool}          |
