# Messages

In this section we describe the processing of the project messages and the
corresponding updates to the state. All created/modified state objects specified
by each message are defined within the [state](01_state.md) section.

## MsgCreateProject

This message creates and stores a new entity doc as an iid along with its nft at
appropriate indexes. Refer to [01_state.md](./01_state.md) for information about
project docs.

| **Field**            | **Type**                  | **Description**                                                                       |
| :------------------- | :------------------------ | :------------------------------------------------------------------------------------ |
| EntityType           | `string`                  | |
| EntityStatus         | `int32`                   | |
| Controllers          | `[]string`                | |
| Context              | `[]*types.Context`        | |
| Verifications        | `[]*types.Verification`   | |
| Services             | `[]*types.Service`        | |
| AccordedRight        | `[]*types.AccordedRight`  | |
| LinkedResource       | `[]*types.LinkedResource` | |
| Deactivated          | `bool`                    | |
| StartDate            | `time.Time`               | |
| EndDate              | `time.Time`               | |
| Stage                | `string`                  | |
| RelayerNode          | `string`                  | |
| VerifiableCredential | `string`                  | |
| Credentials          | `[]string`                | |
| Signer               | `string`                  | |

```go
type MsgCreateEntity struct {
  EntityType     string
  EntityStatus   int32
  Controllers    []string
  Context        []*types.Context
  Verifications  []*types.Verification
  Services       []*types.Service
  AccordedRight  []*types.AccordedRight
  LinkedResource []*types.LinkedResource
  Deactivated          bool
  StartDate            *time.Time
  EndDate              *time.Time
  Stage                string
  RelayerNode          string
  VerifiableCredential string
  Credentials          []string
  Signer               string
}
```

This message is expected to fail if:
