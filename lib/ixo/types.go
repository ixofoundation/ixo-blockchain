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
