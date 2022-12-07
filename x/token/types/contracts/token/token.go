package token

// type Mint struct {
// 	OwnerDid
// }

type TokenContract interface {
	Initiate()
	Mint()
	Transfer()
}
