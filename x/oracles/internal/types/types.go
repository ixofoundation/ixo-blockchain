package types

import "github.com/ixofoundation/ixo-cosmos/x/ixo"

type (
	Oracle  ixo.Did
	Oracles []Oracle
)

var (
	OraclesKey = []byte{0x00}
)
