package fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/oracle"
)

//DOC SETUP
var _ oracle.Payload = Payload{}

// This is the payload for updating fees
type Payload struct {
	FeeFactor int
	Nonce     int
}

func (payload Payload) Type() string { return "fee" }
func (payload Payload) ValidateBasic() sdk.Error {
	return nil
}
