# State

The project module stores four lists of the following four types of data, as
well as module parameters (see the [Params spec page](06_params.md)):

1. [Entity docs](#entity-docs)

## Entity Docs

```go
type EntityDoc struct {
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
