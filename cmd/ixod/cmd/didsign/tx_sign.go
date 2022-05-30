package didsign

import (
	"os"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	// "github.com/cosmos/cosmos-sdk/x/auth/ante"
	// authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	// bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	didtypes "github.com/ixofoundation/ixo-blockchain/x/did/types"
)

const (
	flagJson            = "file"
	flagFile            = "json"
	flagDid             = "did"
	flagMultisig        = "multisig"
	flagOverwrite       = "overwrite"
	flagSigOnly         = "signature-only"
	flagAmino           = "amino"
	flagNoAutoIncrement = "no-auto-increment"
)

type BroadcastReq struct {
	Tx   legacytx.StdTx `json:"tx" yaml:"tx"`
	Mode string         `json:"mode" yaml:"mode"`
}

func readTxAndInitContexts(clientCtx client.Context, cmd *cobra.Command, filename string) (client.Context, tx.Factory, sdk.Tx, error) {
	stdTx, err := authclient.ReadTxFromFile(clientCtx, filename)
	if err != nil {
		return clientCtx, tx.Factory{}, nil, err
	}

	txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
	return clientCtx, txFactory, stdTx, nil
}

func setOutputFile(cmd *cobra.Command) (func(), error) {
	outputDoc, _ := cmd.Flags().GetString(flags.FlagOutputDocument)
	if outputDoc == "" {
		return func() {}, nil
	}

	fp, err := os.OpenFile(outputDoc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return func() {}, err
	}

	cmd.SetOut(fp)

	return func() { fp.Close() }, nil
}

// GetSignCommand returns the transaction sign command.
func GetSignCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-with-did [did] [file]",
		Short: "Sign a transaction generated offline",
		Long: `Sign a transaction created with the --generate-only flag.
It will read a transaction from [file], sign it, and print its JSON encoding.

If the --signature-only flag is set, it will output the signature parts only.

The --offline flag makes sure that the client will not reach out to full node.
As a result, the account and sequence number queries will not be performed and
it is required to set such parameters manually. Note, invalid values will cause
the transaction to fail.

The --multisig=<multisig_key> flag generates a signature on behalf of a multisig account
key. It implies --signature-only. Full multisig signed transactions may eventually
be generated via the 'multisign' command.
`,
		PreRun: preSignCmd,
		RunE:   makeSignCmd(),
		Args:   cobra.ExactArgs(2),
	}

	cmd.Flags().String(flagMultisig, "", "Address or key name of the multisig account on behalf of which the transaction shall be signed")
	cmd.Flags().Bool(flagOverwrite, false, "Overwrite existing signatures with a new one. If disabled, new signature will be appended")
	cmd.Flags().Bool(flagSigOnly, false, "Print only the signatures")
	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")
	cmd.Flags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.Flags().Bool(flagAmino, false, "Generate Amino encoded JSON suitable for submiting to the txs REST endpoint")
	// cmd.MarkFlagRequired(flags.FlagFrom)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func preSignCmd(cmd *cobra.Command, _ []string) {
	// Conditionally mark the account and sequence numbers required as no RPC
	// query will be done.
	if offline, _ := cmd.Flags().GetBool(flags.FlagOffline); offline {
		cmd.MarkFlagRequired(flags.FlagAccountNumber)
		cmd.MarkFlagRequired(flags.FlagSequence)
	}
}

func sign(txf tx.Factory, clientCtx client.Context, txBuilder client.TxBuilder, overwriteSig bool, ixoDid exported.IxoDid) error {
	var privateKey ed25519.PrivKey
	privateKey.Key = append(base58.Decode(ixoDid.Secret.SignKey), base58.Decode(ixoDid.VerifyKey)...)

	signMode := txf.SignMode()
	if signMode == signing.SignMode_SIGN_MODE_UNSPECIFIED {
		// use the SignModeHandler's default mode if unspecified
		signMode = clientCtx.TxConfig.SignModeHandler().DefaultMode()
	}
	// err := checkMultipleSigners(signMode, txBuilder.GetTx())
	// if err != nil {
	// 	return err
	// }

	signerData := authsigning.SignerData{
		ChainID:       txf.ChainID(),
		AccountNumber: txf.AccountNumber(),
		Sequence:      txf.Sequence(),
	}

	// For SIGN_MODE_DIRECT, calling SetSignatures calls setSignerInfos on
	// TxBuilder under the hood, and SignerInfos is needed to generated the
	// sign bytes. This is the reason for setting SetSignatures here, with a
	// nil signature.
	//
	// Note: this line is not needed for SIGN_MODE_LEGACY_AMINO, but putting it
	// also doesn't affect its generated sign bytes, so for code's simplicity
	// sake, we put it here.
	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   privateKey.PubKey(),
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}
	var prevSignatures []signing.SignatureV2
	var err error
	if !overwriteSig {
		prevSignatures, err = txBuilder.GetTx().GetSignaturesV2()
		if err != nil {
			return err
		}
	}
	if err := txBuilder.SetSignatures(sig); err != nil {
		return err
	}

	bytesToSign, err := clientCtx.TxConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	if err != nil {
		return err
	}

	sigBytes, err := privateKey.Sign(bytesToSign)
	if err != nil {
		return err
	}

	// Construct the SignatureV2 struct
	sigData = signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: sigBytes,
	}
	sig = signing.SignatureV2{
		PubKey:   privateKey.PubKey(),
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}

	if overwriteSig {
		return txBuilder.SetSignatures(sig)
	}
	prevSignatures = append(prevSignatures, sig)
	return txBuilder.SetSignatures(prevSignatures...)
}

func makeSignCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) (err error) {

		singerDid, err := didtypes.UnmarshalIxoDid(args[0])
		if err != nil {
			return err
		}

		var clientCtx client.Context

		clientCtx, err = client.GetClientTxContext(cmd)

		if err != nil {
			return err
		}
		f := cmd.Flags()

		clientCtx, txF, newTx, err := readTxAndInitContexts(clientCtx, cmd, args[1])
		clientCtx = clientCtx.WithFromAddress(singerDid.Address())

		txF, err = tx.PrepareFactory(clientCtx, txF)

		if err != nil {
			return err
		}

		txCfg := clientCtx.TxConfig
		txBuilder, err := txCfg.WrapTxBuilder(newTx)
		if err != nil {
			return err
		}

		// printSignatureOnly, _ := cmd.Flags().GetBool(flagSigOnly)
		// if err != nil {
		// 	return err
		// }
		// from, _ := cmd.Flags().GetString(flags.FlagFrom)
		// _, fromName, _, err := client.GetFromFields(txF.Keybase(), from, clientCtx.GenerateOnly)
		// if err != nil {
		// 	return fmt.Errorf("error getting account from keybase: %w", err)
		// }

		overwrite, _ := f.GetBool(flagOverwrite)

		err = sign(txF, clientCtx, txBuilder, overwrite, singerDid)
		if err != nil {
			return err
		}

		aminoJSON, err := f.GetBool(flagAmino)
		if err != nil {
			return err
		}

		var json []byte
		if aminoJSON {
			stdTx, err := tx.ConvertTxToStdTx(clientCtx.LegacyAmino, txBuilder.GetTx())
			if err != nil {
				return err
			}
			req := BroadcastReq{
				Tx:   stdTx,
				Mode: "block|sync|async",
			}
			json, err = clientCtx.LegacyAmino.MarshalJSON(req)
			if err != nil {
				return err
			}
		} else {
			json, err = marshalSignatureJSON(txCfg, txBuilder, false)
			if err != nil {
				return err
			}
		}

		outputDoc, _ := cmd.Flags().GetString(flags.FlagOutputDocument)
		if outputDoc == "" {
			cmd.Printf("%s\n", json)
			return nil
		}

		fp, err := os.OpenFile(outputDoc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer func() {
			err2 := fp.Close()
			if err == nil {
				err = err2
			}
		}()

		_, err = fp.Write(append(json, '\n'))

		// txBytes, err := clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
		// if err != nil {
		// 	return err
		// }

		// // broadcast to a Tendermint node
		// res, err := clientCtx.BroadcastTx(txBytes)
		// if err != nil {
		// 	return err
		// }
		// return clientCtx.PrintProto(res)
		return err
	}
}

func marshalSignatureJSON(txConfig client.TxConfig, txBldr client.TxBuilder, signatureOnly bool) ([]byte, error) {
	parsedTx := txBldr.GetTx()
	if signatureOnly {
		sigs, err := parsedTx.GetSignaturesV2()
		if err != nil {
			return nil, err
		}
		return txConfig.MarshalSignatureJSON(sigs)
	}

	return txConfig.TxJSONEncoder()(parsedTx)
}
