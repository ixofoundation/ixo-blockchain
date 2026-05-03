package types

import (
	"bytes"
)

const (
	// ModuleName defines the module name.
	ModuleName = "names"

	// StoreKey is the primary KV store key.
	StoreKey = ModuleName

	// RouterKey is the message route.
	RouterKey = ModuleName

	// QuerierRoute is the legacy querier route name.
	QuerierRoute = ModuleName
)

// keyDelimiter is used to separate variable-length components in composite
// keys. It cannot occur in any validated input (namespaces and normalized
// names are ASCII; DIDs use ASCII without NUL).
var keyDelimiter = []byte{0x00}

var (
	// NamespaceKeyPrefix: 0x01 | namespace_name -> Namespace
	NamespaceKeyPrefix = []byte{0x01}

	// NameRecordKeyPrefix: 0x02 | namespace_name | 0x00 | normalized_name -> NameRecord
	NameRecordKeyPrefix = []byte{0x02}

	// OwnerIndexKeyPrefix: 0x03 | owner_did | 0x00 | namespace_name | 0x00 | normalized_name -> []byte{}
	OwnerIndexKeyPrefix = []byte{0x03}
)

// NamespaceKey builds the storage key for a Namespace.
func NamespaceKey(name string) []byte {
	return append(append([]byte{}, NamespaceKeyPrefix...), []byte(name)...)
}

// NameRecordKey builds the storage key for a NameRecord.
func NameRecordKey(namespace, normalizedName string) []byte {
	out := append([]byte{}, NameRecordKeyPrefix...)
	out = append(out, []byte(namespace)...)
	out = append(out, keyDelimiter...)
	out = append(out, []byte(normalizedName)...)
	return out
}

// NameRecordNamespacePrefix builds the iteration prefix for all records under
// a namespace.
func NameRecordNamespacePrefix(namespace string) []byte {
	out := append([]byte{}, NameRecordKeyPrefix...)
	out = append(out, []byte(namespace)...)
	out = append(out, keyDelimiter...)
	return out
}

// OwnerIndexKey builds the reverse-index key.
func OwnerIndexKey(ownerDid, namespace, normalizedName string) []byte {
	out := append([]byte{}, OwnerIndexKeyPrefix...)
	out = append(out, []byte(ownerDid)...)
	out = append(out, keyDelimiter...)
	out = append(out, []byte(namespace)...)
	out = append(out, keyDelimiter...)
	out = append(out, []byte(normalizedName)...)
	return out
}

// OwnerIndexPrefix builds the iteration prefix for all names owned by a DID.
func OwnerIndexPrefix(ownerDid string) []byte {
	out := append([]byte{}, OwnerIndexKeyPrefix...)
	out = append(out, []byte(ownerDid)...)
	out = append(out, keyDelimiter...)
	return out
}

// ParseOwnerIndexSuffix splits the suffix `namespace | 0x00 | normalized_name`
// from an iterator key past the OwnerIndexPrefix(ownerDid).
func ParseOwnerIndexSuffix(suffix []byte) (namespace, normalizedName string, ok bool) {
	ns, name, found := bytes.Cut(suffix, keyDelimiter)
	if !found {
		return "", "", false
	}
	return string(ns), string(name), true
}
