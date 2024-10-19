# Events

The iid module emits the following typed events:

### IidDocumentCreatedEvent

Emitted after a successful `MsgCreateIidDocument`, `MsgCreateEntity`, since for entity creation an iid doc also gets created.

```go
type IidDocumentCreatedEvent struct {
	IidDocument *IidDocument
}
```

The field's descriptions is as follows:

- `iidDocument` - the full [IidDocument](02_state.md#iiddocument)

### IidDocumentUpdatedEvent

Emitted after a successful Msg of all the rest of iid module Msg types since they all update the iid doc.

```go
type IidDocumentUpdatedEvent struct {
	IidDocument *IidDocument
}
```

The field's descriptions is as follows:

- `iidDocument` - the full [IidDocument](02_state.md#iiddocument)
