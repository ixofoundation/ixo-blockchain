## Negotiation  


### types

#### negotiation.go
```
type Negotiation interface {
	GetNegotiationID() NegotiationID
	SetNegotiationID(NegotiationID) error

	GetBuyerAddress() ctypes.AccAddress
	SetBuyerAddress(ctypes.AccAddress) error

	GetSellerAddress() ctypes.AccAddress
	SetSellerAddress(ctypes.AccAddress) error

	GetTime() int64
	SetTime(int64) error
	
	GetPegHash() types.PegHash
	SetPegHash(types.PegHash) error

	GetBuyerSignature() Signature
	SetBuyerSignature(Signature) error

	GetSellerSignature() Signature
	SetSellerSignature(Signature) error

	GetBuyerBlockHeight() int64
	SetBuyerBlockHeight(int64) error

	GetSellerBlockHeight() int64
	SetSellerBlockHeight(int64) error
}
```

- Base Negotiation interface
### keys.go
- #### negotiationKey 
    -  append(append(buyerAddress,sellerAddress),pegHash)  

## Keeper
```
type Keeper struct {
	storeKey      ctypes.StoreKey
	accountKeeper auth.AccountKeeper
	cdc           *codec.Codec
}
```


### methods
- #### setNegotiation
- #### getNegotiation
- #### getNegotiations 