package did

import (
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

//DOC SETUP
// enforce the DidDoc type at compile time
var _ ixo.DidDoc = (*BaseDidDoc)(nil)

type BaseDidDoc struct {
	Did    ixo.Did `json:"did"`
	PubKey string  `json:"pubKey"` // May be nil, if not known.
}

//GETTERS
func (dd BaseDidDoc) GetDid() ixo.Did { return dd.Did }
func (dd BaseDidDoc) GetPubKey() string { return dd.PubKey }


//SETTERS
func (dd BaseDidDoc) SetDid(did ixo.Did) error {
	if len(dd.Did) != 0 {
		return errors.New("cannot override BaseDidDoc did")
	}
	dd.Did = did
	return nil
}
func (dd BaseDidDoc) SetPubKey(pubKey string) error {
	if len(dd.PubKey) != 0 {
		return errors.New("cannot override BaseDidDoc pubKey")
	}
	dd.PubKey = pubKey
	return nil
}

//DECODERS
type DidDocDecoder func(didDocBytes []byte) (ixo.DidDoc, error)

func GetDidDocDecoder(cdc *wire.Codec) DidDocDecoder {
	return func(didDocBytes []byte) (res ixo.DidDoc, err error) {
		if len(didDocBytes) == 0 {
			return nil, sdk.ErrTxDecode("didDocBytes are empty")
		}
		didDoc := BaseDidDoc{}
		err = cdc.UnmarshalBinary(didDocBytes, &didDoc)
		if err != nil {
			panic(err)
		}
		return didDoc, err
	}
}

//**************************************************************************************

//ADD DIDDOC
type AddDidMsg struct {
	DidDoc BaseDidDoc `json:"didDoc"`
}

// New Ixo message
func NewAddDidMsg(did string, publicKey string) AddDidMsg {
	didDoc := BaseDidDoc{
		Did:    did,
		PubKey: publicKey,
	}
	return AddDidMsg{
		DidDoc: didDoc,
	}
}
// enforce the msg type at compile time
var _ sdk.Msg = AddDidMsg{}

// nolint
func (msg AddDidMsg) Type() string                            { return "did" }
func (msg AddDidMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg AddDidMsg) GetSigners() []sdk.Address               { return []sdk.Address{[]byte(msg.DidDoc.GetDid())} }
func (msg AddDidMsg) String() string {
	return fmt.Sprintf("AddDidMsg{Did: %v, publicKey: %v}", string(msg.DidDoc.GetDid()), msg.DidDoc.GetPubKey())
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg AddDidMsg) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg AddDidMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("SignBytes")
	fmt.Println(string(b))
	return b
}
