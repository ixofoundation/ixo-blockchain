package types

import (
	"fmt"
)

// DefaultGenesisState returns an empty genesis: no namespaces, no names.
// Namespaces are seeded post-launch via governance.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Namespaces: []Namespace{},
		Names:      []NameRecord{},
	}
}

// Validate performs basic genesis state validation:
//   - every namespace is internally valid
//   - namespace names are unique
//   - every NameRecord references an existing namespace and is internally
//     consistent
//   - (namespace, normalized_name) tuples are unique
func (gs GenesisState) Validate() error {
	nsByName := make(map[string]Namespace, len(gs.Namespaces))
	for _, ns := range gs.Namespaces {
		if err := ValidateNamespace(ns); err != nil {
			return fmt.Errorf("namespace %q invalid: %w", ns.Name, err)
		}
		if _, dup := nsByName[ns.Name]; dup {
			return fmt.Errorf("duplicate namespace %q", ns.Name)
		}
		nsByName[ns.Name] = ns
	}

	type recordKey struct{ ns, name string }
	seen := make(map[recordKey]struct{}, len(gs.Names))
	for _, r := range gs.Names {
		ns, ok := nsByName[r.Namespace]
		if !ok {
			return fmt.Errorf("name %q references unknown namespace %q", r.NormalizedName, r.Namespace)
		}
		if r.NormalizedName == "" {
			return fmt.Errorf("name with empty normalized_name in namespace %q", r.Namespace)
		}
		if r.NormalizedName != NormalizeName(r.NormalizedName) {
			return fmt.Errorf("name %q is not in normalized form", r.NormalizedName)
		}
		if err := ValidateNameAgainstNamespace(ns, r.NormalizedName); err != nil {
			return fmt.Errorf("name %q in namespace %q invalid: %w", r.NormalizedName, r.Namespace, err)
		}
		if r.OwnerDid == "" {
			return fmt.Errorf("name %q has empty owner_did", r.NormalizedName)
		}
		if r.ValidUntil != 0 && !ns.AllowExpiry {
			return fmt.Errorf("name %q has valid_until set but namespace %q forbids expiry", r.NormalizedName, r.Namespace)
		}
		k := recordKey{ns: r.Namespace, name: r.NormalizedName}
		if _, dup := seen[k]; dup {
			return fmt.Errorf("duplicate name %q in namespace %q", r.NormalizedName, r.Namespace)
		}
		seen[k] = struct{}{}
	}
	return nil
}
