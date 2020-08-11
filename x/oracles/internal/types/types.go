package types

import (
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"strings"
)

// --------------------------------------- Oracle/s

type (
	Oracle struct {
		OracleDid    did.Did         `json:"oracle_did" yaml:"oracle_did"`
		Capabilities OracleTokenCaps `json:"capabilities" yaml:"capabilities"`
	}
	Oracles []Oracle
)

func NewOracle(oracleDid did.Did, caps OracleTokenCaps) Oracle {
	return Oracle{
		OracleDid:    oracleDid,
		Capabilities: caps,
	}
}

func (os Oracles) Includes(oracle Oracle) bool {
	for _, o := range os {
		if oracle.OracleDid == o.OracleDid {
			return true
		}
	}
	return false
}

// --------------------------------------- OracleTokenCap/s

type (
	OracleTokenCap struct {
		Denom        string    `json:"denom" yaml:"denom"`
		Capabilities TokenCaps `json:"capabilities" yaml:"capabilities"`
	}
	OracleTokenCaps []OracleTokenCap
)

func NewOracleTokenCap(denom string, caps TokenCaps) OracleTokenCap {
	return OracleTokenCap{
		Denom:        denom,
		Capabilities: caps,
	}
}

func (otcs OracleTokenCaps) Includes(denom string) bool {
	for _, oc := range otcs {
		if oc.Denom == denom {
			return true
		}
	}
	return false
}

func (otcs OracleTokenCaps) MustGet(denom string) OracleTokenCap {
	for _, oc := range otcs {
		if oc.Denom == denom {
			return oc
		}
	}
	panic("capability for specified denom not found")
}

func ParseTokenCaps(capsStr string) (TokenCaps, error) {
	capsStr = strings.TrimSpace(capsStr)

	capsStrs := strings.Split(capsStr, "/")
	caps := make(TokenCaps, len(capsStrs))
	for i, capStr := range capsStrs {
		capability := TokenCap(capStr)
		if !capability.IsValid() {
			return nil, fmt.Errorf("invalid capability: %s", capStr)
		}
		caps[i] = capability
	}

	return caps, nil
}

func ParseOracleTokenCap(capStr string) (OracleTokenCap, error) {
	capStr = strings.TrimSpace(capStr)

	capsStrs := strings.Split(capStr, ":")
	if len(capsStrs) != 2 {
		return OracleTokenCap{}, fmt.Errorf("invalid capability: %s", capStr)
	}

	denom := capsStrs[0]
	capsStr := capsStrs[1]

	if len(denom) == 0 {
		return OracleTokenCap{}, fmt.Errorf("invalid empty token: %s", capStr)
	}

	tokenCaps, err := ParseTokenCaps(capsStr)
	if err != nil {
		return OracleTokenCap{}, err
	}

	return NewOracleTokenCap(denom, tokenCaps), nil
}

func ParseOracleTokenCaps(capsStr string) (OracleTokenCaps, error) {
	capsStr = strings.TrimSpace(capsStr)
	if len(capsStr) == 0 {
		return nil, nil
	}

	capsStrs := strings.Split(capsStr, ",")
	caps := make(OracleTokenCaps, len(capsStrs))
	for i, capStr := range capsStrs {
		capability, err := ParseOracleTokenCap(capStr)
		if err != nil {
			return nil, err
		}

		caps[i] = capability
	}

	// TODO: consider sorting and validating

	return caps, nil
}

// --------------------------------------- TokenCap/s

type (
	TokenCap  string
	TokenCaps []TokenCap
)

func (tcs TokenCaps) Includes(cap TokenCap) bool {
	for _, tc := range tcs {
		if tc == cap {
			return true
		}
	}
	return false
}

const (
	MintCap     TokenCap = "mint"
	BurnCap     TokenCap = "burn"
	TransferCap TokenCap = "transfer"
)

func (tc TokenCap) IsValid() bool {
	return tc == MintCap || tc == BurnCap || tc == TransferCap
}
