package ante

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
// )

// type FeePayer struct {
// 	feePayerAccount  authtypes.AccountI
// 	recipientAddress sdk.AccAddress
// }

// func NewFeePayer(feePayerAccount authtypes.AccountI, recipientAddress sdk.AccAddress) FeePayer {
// 	return FeePayer{
// 		feePayerAccount:  feePayerAccount,
// 		recipientAddress: recipientAddress,
// 	}
// }

// func (fp *FeePayer) GetFeePayerAccount() authtypes.AccountI { return fp.feePayerAccount }
// func (fp *FeePayer) GetRecipientAddress() sdk.AccAddress    { return fp.recipientAddress }

// type IxoFeeTxMsg interface {
// 	FeePayerFromIid(ctx sdk.Context, accountKeeper authante.AccountKeeper, iidKeeper iidkeeper.Keeper) (FeePayer, error)
// }

// type IxoFeeTx struct {
// 	sdk.FeeTx
// }

// func (tx *IxoFeeTx) GetFeePayerMsgs() []IxoFeeTxMsg {
// 	var msgs []IxoFeeTxMsg

// 	for _, msg := range tx.GetMsgs() {
// 		if msg, ok := msg.(IxoFeeTxMsg); ok {
// 			msgs = append(msgs, msg)
// 		}
// 	}

// 	return msgs
// }
