package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type CreateBondMsg struct {
	SignBytes string  `json:"signBytes"`
	TxHash    string  `json:"txHash"`
	SenderDid ixo.Did `json:"senderDid"`
	BondDid   ixo.Did `json:"bondDid"`
	PubKey    string  `json:"pubKey"`
	Data      BondDoc `json:"data"`
}

var _ sdk.Msg = CreateBondMsg{}

func (msg CreateBondMsg) Type() string                            { return ModuleName }
func (msg CreateBondMsg) Route() string                           { return RouterKey }
func (msg CreateBondMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateBondMsg) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}

	valid, err = CheckNotEmpty(msg.BondDid, "BondDid")
	if !valid {
		return err
	}

	valid, err = CheckNotEmpty(msg.Data.CreatedBy, "CreatedBy")
	if !valid {
		return err
	}

	return nil
}

func (msg CreateBondMsg) GetBondDid() ixo.Did   { return msg.BondDid }
func (msg CreateBondMsg) GetSenderDid() ixo.Did { return msg.SenderDid }
func (msg CreateBondMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetBondDid())}
}

func (msg CreateBondMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg CreateBondMsg) GetPubKey() string { return msg.PubKey }

func (msg CreateBondMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg CreateBondMsg) IsNewDid() bool { return true }

var _ StoredBondDoc = (*CreateBondMsg)(nil)
