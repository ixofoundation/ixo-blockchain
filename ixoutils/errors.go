package ixoutils

import "fmt"

type DecNotFoundError struct {
	Key string
}

func (e DecNotFoundError) Error() string {
	return fmt.Sprintf("no math.LegacyDec at key (%s)", e.Key)
}
