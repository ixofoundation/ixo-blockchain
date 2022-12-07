package cw20

import (
	"bytes"
	"encoding/json"
)

type MinterResponse struct {
	Minter string  `json:"minter"`
	Cap    *uint64 `json:"cap,omitempty"`
}

func (m MinterResponse) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type InstantiateMarketingInfo struct {
	Project     string `json:"project"`
	Description string `json:"description"`
	Marketing   string `json:"marketing"`
	// Logo        Logo
}

type Cw20Coin struct {
	Address string
	Amount  uint64
}

type InstantiateMsg struct {
	Name            string                   `json:"name"`
	Symbol          string                   `json:"symbol"`
	Decimals        uint32                   `json:"decimals"`
	Mint            MinterResponse           `json:"mint"`
	Marketing       InstantiateMarketingInfo `json:"marketing,omitempty"`
	InitialBalances []Cw20Coin               `json:"initial_balances"`
}

func (m InstantiateMsg) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type Mint struct {
	Recipient string `json:"recipient"`
	Amount    uint   `json:"amount"`
}

func (m Mint) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
