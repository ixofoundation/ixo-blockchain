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
var _ ixo.DidDoc = (*BaseDidDoc)(nil)

type BaseDidDoc struct {
	Did    ixo.Did `json:"did"`
	PubKey string  `json:"pub_key"` // May be nil, if not known.
}

//GETTERS
func (dd BaseDidDoc) GetDid() ixo.Did   { return dd.Did }
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

//ADD DidDoc
type AddDidMsg struct {
	DidDoc BaseDidDoc
}

func NewAddDidMsg(did string, publicKey string) AddDidMsg {
	didDoc := BaseDidDoc{
		Did:    did,
		PubKey: publicKey,
	}
	return AddDidMsg{
		DidDoc: didDoc,
	}
}

var _ sdk.Msg = AddDidMsg{}

func (msg AddDidMsg) Type() string                            { return "did" }
func (msg AddDidMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg AddDidMsg) GetSigners() []sdk.Address               { return []sdk.Address{[]byte(msg.DidDoc.GetDid())} }
func (msg AddDidMsg) String() string {
	return fmt.Sprintf("AddDidMsg{Did: %v, publicKey: %v}", string(msg.DidDoc.GetDid()), msg.DidDoc.GetPubKey())
}

func (msg AddDidMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg AddDidMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

//**************************************************************************************

//GET DidDoc

type GetDidMsg struct {
	Did ixo.Did `json:"did"`
}

func NewGetDidMsg(did string) GetDidMsg {
	return GetDidMsg{
		Did: did,
	}
}

var _ sdk.Msg = GetDidMsg{}

func (msg GetDidMsg) Type() string                            { return "did" }
func (msg GetDidMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg GetDidMsg) GetSigners() []sdk.Address               { return nil }
func (msg GetDidMsg) String() string {
	return fmt.Sprintf("DidMsg{Did: %v}", msg.Did)
}

func (msg GetDidMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg GetDidMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
