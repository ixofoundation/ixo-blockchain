# Parameters

The names module has **no module-wide parameters** stored in chain state. All configuration is per-namespace, on each [Namespace](./02_state.md#namespace) record, plus a small set of compile-time **module-wide constants** baked into the binary.

## Module-wide parameters

None. There is no `Params` record, no `MsgUpdateParams`, and no params subspace registered for this module. Adding module-wide params (e.g. a global per-tx fee, a chain-wide reserved-words list) is a forward-compatible addition for a future upgrade.

## Per-namespace fields

Per-namespace fields are not "params" in the Cosmos SDK x/params sense — they are stored on each `Namespace` record (KV prefix `0x01`). Each namespace has its own settings, set at create time and replaceable wholesale via `MsgUpdateNamespace`:

| Field                      | Type       | Mutability                         | Notes                                                                                                                  |
| -------------------------- | ---------- | ---------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `name`                     | `string`   | immutable                          | Set at `MsgCreateNamespace`. Lowercase ASCII alphanumeric + `-`/`_` (`^[a-z0-9_-]+$`). Max `MaxNamespaceNameLength = 64` bytes. |
| `description`              | `string`   | gov via `MsgUpdateNamespace`       | Free-form summary. Max `MaxNamespaceDescriptionLength = 4096` bytes.                                                   |
| `registrar_accounts`       | `[]string` | gov via `MsgUpdateNamespace`       | Each entry validated as a bech32 address. List size unbounded. May be empty when `allow_self_register = true`.         |
| `allow_self_register`      | `bool`     | gov via `MsgUpdateNamespace`       | Required true unless `registrar_accounts` is non-empty (mutually-exclusive ban — see `ValidateNamespace`).             |
| `allow_registrar_override` | `bool`     | gov via `MsgUpdateNamespace`       | Enables registrar-driven transfers without owner consent.                                                              |
| `min_length`               | `uint32`   | gov via `MsgUpdateNamespace`       | Minimum length of the **normalised** name. Range typically `1+`.                                                       |
| `max_length`               | `uint32`   | gov via `MsgUpdateNamespace`       | Maximum length of the normalised name. Must be `> 0` and `<= MaxNameLengthCap = 256`. `min_length` must be `<= max_length`. |
| `regex`                    | `string`   | gov via `MsgUpdateNamespace`       | Optional Go-regexp pattern applied on top of the chain-wide `[a-z0-9_-]+` charset. Empty disables. Max `MaxNamespaceRegexLength = 256` bytes. |
| `allow_expiry`             | `bool`     | gov via `MsgUpdateNamespace`       | **Reserved**. Gates whether records may carry a non-zero `valid_until`. Always `false` in v1.                          |

`MsgUpdateNamespace` does **not** retroactively re-validate existing `NameRecord`s. Names registered under the previous configuration remain valid even if the new configuration would have rejected them. The new configuration applies only to subsequent registration / update messages.

## Module-wide constants

These live in `x/names/types/keys.go` and `x/names/types/names.go`. They cannot be changed through state — only through a chain upgrade.

| Constant                          | Location                | Value                          | Notes                                                                                                              |
| --------------------------------- | ----------------------- | ------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ModuleName`                      | `types/keys.go`         | `"names"`                      | The module identifier; used as `StoreKey`, `RouterKey`, and `QuerierRoute`.                                        |
| `NamespaceKeyPrefix`              | `types/keys.go`         | `0x01`                         | KV prefix for `Namespace` records.                                                                                 |
| `NameRecordKeyPrefix`             | `types/keys.go`         | `0x02`                         | KV prefix for `NameRecord` records (`0x02 \| ns \| 0x00 \| normalized`).                                           |
| `OwnerIndexKeyPrefix`             | `types/keys.go`         | `0x03`                         | KV prefix for the reverse-lookup index (`0x03 \| owner_did \| 0x00 \| ns \| 0x00 \| normalized`).                  |
| `keyDelimiter`                    | `types/keys.go`         | `0x00`                         | Separator byte for variable-length composite key segments. Safe because all validated inputs are ASCII without NUL.|
| `defaultNameCharset`              | `types/names.go`        | `^[a-z0-9_-]+$`                | Chain-wide regex applied to every normalised name regardless of namespace `regex`.                                 |
| `MaxNamespaceNameLength`          | `types/names.go`        | `64`                           | Max bytes for `Namespace.name`.                                                                                    |
| `MaxNamespaceDescriptionLength`   | `types/names.go`        | `4096`                         | Max bytes for `Namespace.description`.                                                                             |
| `MaxNamespaceRegexLength`         | `types/names.go`        | `256`                          | Max bytes for `Namespace.regex`.                                                                                   |
| `MaxNameLengthCap`                | `types/names.go`        | `256`                          | Upper bound on `Namespace.max_length` (caps the effective `NameRecord.normalized_name` size).                       |
| `MaxNameRecordEvidenceHashLength` | `types/names.go`        | `256`                          | Max bytes for `NameRecord.evidence_hash` (also enforced on `MsgRegisterNameByRegistrar` / `MsgUpdateNameByRegistrar`). |
| `MaxNameRecordSourceLength`       | `types/names.go`        | `64`                           | Max bytes for `NameRecord.source` (same enforcement points as `evidence_hash`).                                    |

## Genesis state

Genesis state is the empty list of namespaces and the empty list of name records:

```go
type GenesisState struct {
    Namespaces []Namespace
    Names      []NameRecord
}
```

The default genesis (`DefaultGenesisState`) is two empty slices — no built-in namespaces are seeded at genesis. Production chains add namespaces post-launch via gov proposals.

`GenesisState.Validate()` enforces:

- every `Namespace` is internally valid (`ValidateNamespace`),
- namespace names are unique within the genesis,
- every `NameRecord` references an existing namespace,
- every `NameRecord.normalized_name` is non-empty and equals `NormalizeName(record.normalized_name)` (i.e. the genesis was produced through proper normalisation),
- every `NameRecord` passes `ValidateNameAgainstNamespace` for its referenced namespace,
- `(namespace, normalized_name)` tuples are unique within the genesis,
- `NameRecord.owner_did` is non-empty,
- `NameRecord.valid_until` is `0` for any namespace where `allow_expiry = false`.
