package types

import (
	"regexp"
	"slices"
	"strings"

	errorsmod "cosmossdk.io/errors"
)

// defaultNameCharset is the v1 character set for normalized names: lowercase
// ASCII letters, digits, hyphen, and underscore. Per-namespace `regex` is
// applied on top of this default.
var defaultNameCharset = regexp.MustCompile(`^[a-z0-9_-]+$`)

// Defensive length caps on user/gov-supplied text fields. These bound
// per-record state size and prevent state-bloat surfaces. Values chosen
// to comfortably fit any realistic use case while keeping a single
// NameRecord well under typical IAVL row budgets.
const (
	// MaxNamespaceNameLength bounds Namespace.Name. 64 chars matches DNS
	// label conventions and is more than enough for any handle namespace.
	MaxNamespaceNameLength = 64

	// MaxNamespaceDescriptionLength bounds Namespace.Description. Generous
	// to allow a paragraph or two of human context per namespace.
	MaxNamespaceDescriptionLength = 4096

	// MaxNamespaceRegexLength bounds Namespace.Regex. 256 chars covers any
	// reasonable ASCII regex pattern.
	MaxNamespaceRegexLength = 256

	// MaxNameLengthCap is the upper bound on Namespace.MaxLength. Caps the
	// effective size of NameRecord.NormalizedName (and indirectly DisplayName
	// up to ~2x via leading/trailing whitespace trim).
	MaxNameLengthCap = 256

	// MaxNameRecordEvidenceHashLength bounds NameRecord.EvidenceHash.
	// SHA-256 hex is 64 chars; multibase IPFS CIDs are ~60; 256 is generous.
	MaxNameRecordEvidenceHashLength = 256

	// MaxNameRecordSourceLength bounds NameRecord.Source. Source tags are
	// short identifiers like "twitter-oauth" or "workos".
	MaxNameRecordSourceLength = 64
)

// NormalizeName produces the canonical lookup form of a name (v1: trim +
// ASCII lowercase). Non-ASCII input is rejected at validation time, so this
// is sufficient.
func NormalizeName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

// ValidateNameAgainstNamespace checks the (normalized) name against the
// namespace's length and regex constraints, plus the chain-wide ASCII charset
// rule.
func ValidateNameAgainstNamespace(ns Namespace, normalized string) error {
	l := uint32(len(normalized))
	if ns.MaxLength == 0 {
		return errorsmod.Wrap(ErrInvalidNamespace, "max_length must be > 0")
	}
	if l < ns.MinLength {
		return errorsmod.Wrapf(ErrInvalidName, "name shorter than namespace min_length %d", ns.MinLength)
	}
	if l > ns.MaxLength {
		return errorsmod.Wrapf(ErrInvalidName, "name longer than namespace max_length %d", ns.MaxLength)
	}
	if !defaultNameCharset.MatchString(normalized) {
		return errorsmod.Wrap(ErrInvalidName, "name must be lowercase ASCII alphanumeric, hyphen or underscore")
	}
	if ns.Regex != "" {
		re, err := regexp.Compile(ns.Regex)
		if err != nil {
			return errorsmod.Wrapf(ErrInvalidNamespace, "namespace regex did not compile: %s", err)
		}
		if !re.MatchString(normalized) {
			return errorsmod.Wrapf(ErrInvalidName, "name does not match namespace regex %q", ns.Regex)
		}
	}
	return nil
}

// ValidateNamespace checks a Namespace value for self-consistency. It does
// not check uniqueness against the store.
func ValidateNamespace(ns Namespace) error {
	if ns.Name == "" {
		return errorsmod.Wrap(ErrInvalidNamespace, "name is required")
	}
	if len(ns.Name) > MaxNamespaceNameLength {
		return errorsmod.Wrapf(ErrInvalidNamespace, "namespace name longer than %d bytes", MaxNamespaceNameLength)
	}
	if !defaultNameCharset.MatchString(ns.Name) {
		return errorsmod.Wrap(ErrInvalidNamespace, "namespace name must be lowercase ASCII alphanumeric, hyphen or underscore")
	}
	if len(ns.Description) > MaxNamespaceDescriptionLength {
		return errorsmod.Wrapf(ErrInvalidNamespace, "description longer than %d bytes", MaxNamespaceDescriptionLength)
	}
	if ns.MaxLength == 0 {
		return errorsmod.Wrap(ErrInvalidNamespace, "max_length must be > 0")
	}
	if ns.MaxLength > MaxNameLengthCap {
		return errorsmod.Wrapf(ErrInvalidNamespace, "max_length exceeds chain cap of %d", MaxNameLengthCap)
	}
	if ns.MinLength > ns.MaxLength {
		return errorsmod.Wrap(ErrInvalidNamespace, "min_length must be <= max_length")
	}
	if len(ns.Regex) > MaxNamespaceRegexLength {
		return errorsmod.Wrapf(ErrInvalidNamespace, "regex longer than %d bytes", MaxNamespaceRegexLength)
	}
	if ns.Regex != "" {
		if _, err := regexp.Compile(ns.Regex); err != nil {
			return errorsmod.Wrapf(ErrInvalidNamespace, "regex did not compile: %s", err)
		}
	}
	if !ns.AllowSelfRegister && len(ns.RegistrarAccounts) == 0 {
		return errorsmod.Wrap(ErrInvalidNamespace, "namespace must allow self-registration or have at least one registrar")
	}
	return nil
}

// ValidateRecordMetadata bounds the registrar-supplied free-form fields on a
// NameRecord (evidence_hash and source). Called from both message
// ValidateBasic and the keeper handler so Wasm sub-message dispatches that
// bypass ante still get the check.
func ValidateRecordMetadata(evidenceHash, source string) error {
	if len(evidenceHash) > MaxNameRecordEvidenceHashLength {
		return errorsmod.Wrapf(ErrInvalidRequest, "evidence_hash longer than %d bytes", MaxNameRecordEvidenceHashLength)
	}
	if len(source) > MaxNameRecordSourceLength {
		return errorsmod.Wrapf(ErrInvalidRequest, "source longer than %d bytes", MaxNameRecordSourceLength)
	}
	return nil
}

// HasRegistrar reports whether addr is one of the namespace's registrars.
func HasRegistrar(ns Namespace, addr string) bool {
	return slices.Contains(ns.RegistrarAccounts, addr)
}
