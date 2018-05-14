package ixo

import (
	
)

// SovrinSecret definition 
type SovrinSecret struct {
	Seed                 string `json:"seed"`
	SignKey              string `json:"signKey"`
	EncryptionPrivateKey string `json:"encryptionPrivateKey"`
}

// SovrinDid definition 
type SovrinDid struct {
	Did                 string       `json:"did"`
	VerifyKey           string       `json:"verifyKey"`
	EncryptionPublicKey string       `json:"encryptionPublicKey"`
	Secret              SovrinSecret `json:"secret"`
}
