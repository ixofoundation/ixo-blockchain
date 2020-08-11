package did

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func GetPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, res sdk.Result) {

		// Get signer PubKey
		var pubKeyEd25519 ed25519.PubKeyEd25519
		switch msg := msg.(type) {
		case MsgAddDid:
			copy(pubKeyEd25519[:], base58.Decode(msg.PubKey))
		default:
			// For the remaining messages, the did is the signer
			didDoc, _ := keeper.GetDidDoc(ctx, msg.GetSignerDid())
			if didDoc == nil {
				return pubKey, sdk.ErrUnauthorized("Issuer did not found").Result()
			}
			copy(pubKeyEd25519[:], base58.Decode(didDoc.GetPubKey()))
		}
		return pubKeyEd25519, sdk.Result{}
	}
}

func getAddDidSignBytes(chainID string, tx auth.StdTx, acc auth.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), tx.Fee, tx.Msgs, tx.Memo,
	)
}

func NewAddDidAnteHandler(ak auth.AccountKeeper, sk supply.Keeper,
	pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res sdk.Result, abort bool) {

		if addr := sk.GetModuleAddress(auth.FeeCollectorName); addr == nil {
			panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
		}

		// all transactions must be of type auth.StdTx
		stdTx, ok := tx.(auth.StdTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = auth.SetGasMeter(simulate, ctx, 0)
			return newCtx, sdk.ErrInternal("tx must be auth.StdTx").Result(), true
		}

		params := ak.GetParams(ctx)

		// Addding of DID uses an infinite gas meter
		newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err.Result(), true
		}

		// Number of messages in the tx must be 1
		if len(tx.GetMsgs()) != 1 {
			return ctx, sdk.ErrInternal("number of messages must be 1").Result(), true
		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if res := auth.ValidateMemo(auth.StdTx{Memo: stdTx.Memo}, params); !res.IsOK() {
			return newCtx, res, true
		}

		// message must be of type MsgAddDid
		msg, ok := stdTx.GetMsgs()[0].(MsgAddDid)
		if !ok {
			return newCtx, sdk.ErrInternal("msg must be MsgCreateProject").Result(), true
		}

		// Get did pubKey
		didPubKey, res := pubKeyGetter(ctx, msg)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// Fetch signer (account underlying DID ). Account expected to not exist
		signerAddr := sdk.AccAddress(didPubKey.Address())
		_, res = auth.GetSignerAcc(newCtx, ak, signerAddr)
		if res.IsOK() {
			return newCtx, sdk.ErrInternal("expected account underlying DID to not exist").Result(), true
		}

		// Create signer's account
		signerAcc := ak.NewAccountWithAddress(ctx, signerAddr)
		ak.SetAccount(ctx, signerAcc)

		// check signature, return account with incremented nonce
		ixoSig := stdTx.GetSignatures()[0]
		isGenesis := ctx.BlockHeight() == 0
		signBytes := getAddDidSignBytes(newCtx.ChainID(), stdTx, signerAcc, isGenesis)
		signerAcc, res = ixo.ProcessSig(newCtx, signerAcc, ixoSig, signBytes, simulate, params)
		if !res.IsOK() {
			return newCtx, res, true
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, sdk.Result{GasWanted: stdTx.Fee.Gas}, false // continue...
	}
}
