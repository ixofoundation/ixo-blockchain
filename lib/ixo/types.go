package ixo

const IxoNativeToken = "uixo"

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// removes an element at index i from a slice, will change order of slice (fast)
func RemoveUnordered[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// removes an element at index i from a slice, will not change order of slice (slow)
func RemoveOrdered[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
