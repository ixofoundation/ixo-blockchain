package cli

import (
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
)

const (
	TRUE  = "true"
	FALSE = "false"
)

func parseBool(boolStr, boolName string) (bool, sdk.Error) {
	boolStr = strings.ToLower(strings.TrimSpace(boolStr))
	if boolStr == TRUE {
		return true, nil
	} else if boolStr == FALSE {
		return false, nil
	} else {
		return false, types.ErrInvalidArgument(types.DefaultCodespace, ""+
			fmt.Sprintf("%s is not a valid bool (true/false)", boolName))
	}
}

func GetCmdCreatePaymentTemplate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-payment-template [payment-template-json] [creator-sovrin-did]",
		Short: "Create and sign a create-payment-template tx using DIDs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			templateJsonStr := args[0]
			sovrinDidStr := args[1]

			sovrinDid, err := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err != nil {
				return err
			}

			var template types.PaymentTemplate
			err = cdc.UnmarshalJSON([]byte(templateJsonStr), &template)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreatePaymentTemplate(template, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdCreatePaymentContract(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "create-payment-contract [payment-contract-id] [payment-template-id] " +
			"[payer-addr] [can-deauthorise] [discount-id] [creator-sovrin-did]",
		Short: "Create and sign a create-payment-contract tx using DIDs",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			templateIdStr := args[1]
			payerAddrStr := args[2]
			canDeauthoriseStr := args[3]
			discountIdStr := args[4]
			sovrinDidStr := args[5]

			payerAddr, err := sdk.AccAddressFromBech32(payerAddrStr)
			if err != nil {
				return err
			}

			canDeauthorise, err := parseBool(canDeauthoriseStr, "canDeauthorise")
			if err != nil {
				return err
			}

			discountId, err := sdk.ParseUint(discountIdStr)
			if err != nil {
				return err
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreatePaymentContract(
				templateIdStr, contractIdStr, payerAddr,
				canDeauthorise, discountId, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateSubscription(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "create-subscription [subscription-id] [payment-contract-id] " +
			"[max-periods] [period-json] [creator-sovrin-did]",
		Short: "Create and sign a create-subscription tx using DIDs",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			subIdStr := args[0]
			contractIdStr := args[1]
			maxPeriodsStr := args[2]
			periodStr := args[3]
			sovrinDidStr := args[4]

			maxPeriods, err := sdk.ParseUint(maxPeriodsStr)
			if err != nil {
				return err
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err != nil {
				return err
			}

			var period types.Period
			err = cdc.UnmarshalJSON([]byte(periodStr), &period)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreateSubscription(subIdStr,
				contractIdStr, maxPeriods, period, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdSetPaymentContractAuthorisation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "set-payment-contract-authorisation [payment-contract-id] " +
			"[authorised] [payer-sovrin-did]",
		Short: "Create and sign a set-payment-contract-authorisation tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			authorisedStr := args[1]
			sovrinDidStr := args[2]

			authorised, err := parseBool(authorisedStr, "authorised")
			if err != nil {
				return err
			}

			sovrinDid, err2 := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err2 != nil {
				return err2
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgSetPaymentContractAuthorisation(
				contractIdStr, authorised, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdGrantPaymentDiscount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "grant-discount [payment-contract-id] [discount-id] " +
			"[recipient-addr] [creator-sovrin-did]",
		Short: "Create and sign a grant-discount tx using DIDs",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			discountIdStr := args[1]
			recipientAddrStr := args[2]
			sovrinDidStr := args[3]

			discountId, err := sdk.ParseUint(discountIdStr)
			if err != nil {
				return err
			}

			recipientAddr, err := sdk.AccAddressFromBech32(recipientAddrStr)
			if err != nil {
				return err
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgGrantDiscount(
				contractIdStr, discountId, recipientAddr, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdRevokePaymentDiscount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke-discount [payment-contract-id] [holder-addr] [creator-sovrin-did]",
		Short: "Create and sign a revoke-discount tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			holderAddrStr := args[1]
			sovrinDidStr := args[2]

			holderAddr, err := sdk.AccAddressFromBech32(holderAddrStr)
			if err != nil {
				return err
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgRevokeDiscount(
				contractIdStr, holderAddr, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdEffectPayment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "effect-payment [payment-contract-id] [creator-sovrin-did]",
		Short: "Create and sign a effect-payment tx using DIDs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			sovrinDidStr := args[1]

			sovrinDid, err := sovrin.UnmarshalSovrinDid(sovrinDidStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgEffectPayment(contractIdStr, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}
