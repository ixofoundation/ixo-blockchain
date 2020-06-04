package ixo

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

type PubKeyGetter func(ctx sdk.Context, msg sdk.Msg) ([32]byte, sdk.Result)

func processSig(ctx sdk.Context, acc auth.Account, msg sdk.Msg, pubKey [32]byte,
	sig IxoSignature, params auth.Params) (updatedAcc auth.Account, res sdk.Result) {

	// Consume signature gas
	ctx.GasMeter().ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")

	// Verify signature
	if res := VerifySignature(msg, pubKey, sig); !res {
		return nil, sdk.ErrUnauthorized("Signature Verification failed").Result()
	}

	if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
		panic(err)
	}

	return acc, res
}

func NewAnteHandler(ak auth.AccountKeeper, sk supply.Keeper, pubKeyGetter PubKeyGetter) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res sdk.Result, abort bool) {

		if addr := sk.GetModuleAddress(auth.FeeCollectorName); addr == nil {
			panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
		}

		// all transactions must be of type ixo.IxoTx
		ixoTx, ok := tx.(IxoTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = auth.SetGasMeter(simulate, ctx, 0)
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		params := ak.GetParams(ctx)

		// Ensure that the provided fees meet a minimum threshold for the validator,
		// if this is a CheckTx. This is only for local mempool purposes, and thus
		// is only ran on check tx.
		if ctx.IsCheckTx() && !simulate {
			res := auth.EnsureSufficientMempoolFees(ctx, ixoTx.Fee)
			if !res.IsOK() {
				return newCtx, res, true
			}
		}

		newCtx = auth.SetGasMeter(simulate, ctx, ixoTx.Fee.Gas)

		// AnteHandlers must have their own defer/recover in order for the BaseApp
		// to know how much gas was used! This is because the GasMeter is created in
		// the AnteHandler, but if it panics the context won't be set properly in
		// runTx's recover call.
		defer func() {
			if r := recover(); r != nil {
				switch rType := r.(type) {
				case sdk.ErrorOutOfGas:
					log := fmt.Sprintf(
						"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
						rType.Descriptor, ixoTx.Fee.Gas, newCtx.GasMeter().GasConsumed(),
					)
					res = sdk.ErrOutOfGas(log).Result()

					res.GasWanted = ixoTx.Fee.Gas
					res.GasUsed = newCtx.GasMeter().GasConsumed()
					abort = true
				default:
					panic(r)
				}
			}
		}()

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err.Result(), true
		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if res := auth.ValidateMemo(auth.StdTx{Memo: ixoTx.Memo}, params); !res.IsOK() {
			return newCtx, res, true
		}

		// fetch first (and only) signer, who's going to pay the fees
		signerAddr := ixoTx.GetSigner()
		signerAcc, res := auth.GetSignerAcc(ctx, ak, signerAddr)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// deduct the fees
		if !ixoTx.Fee.Amount.IsZero() {
			res = auth.DeductFees(sk, newCtx, signerAcc, ixoTx.Fee.Amount)
			if !res.IsOK() {
				return newCtx, res, true
			}

			// reload the account as fees have been deducted
			signerAcc = ak.GetAccount(newCtx, signerAcc.GetAddress())
		}

		// Get pubKey
		msg := ixoTx.GetMsgs()[0]
		pubKey, res := pubKeyGetter(ctx, msg)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// check signature, return account with incremented nonce
		ixoSig := ixoTx.GetSignatures()[0]
		signerAcc, res = processSig(newCtx, signerAcc, msg, pubKey, ixoSig, params)
		if !res.IsOK() {
			return newCtx, res, true
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, sdk.Result{GasWanted: ixoTx.Fee.Gas}, false // continue...
	}
}
