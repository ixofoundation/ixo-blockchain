package types

import "github.com/ixofoundation/ixo-cosmos/x/ixo"

type (
	Oracle  ixo.Did
	Oracles []Oracle
)

func (os Oracles) Contains(oracle Oracle) bool {
	for _, o := range os {
		if oracle == o {
			return true
		}
	}
	return false
}
