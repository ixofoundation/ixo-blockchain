package cli

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ixotypes "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	didtypes "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"github.com/spf13/cobra"
)

const (
	TRUE  = "true"
	FALSE = "false"
)

// Period types
const (
	BlockPeriodType string = "payments/BlockPeriod"
	TimePeriodType  string = "payments/TimePeriod"
)

type period struct {
	Type  string
	Value map[string]string
}

func parsePeriodString(periodStr string) (*period, error) {
	period := &period{}

	err := json.Unmarshal([]byte(periodStr), period)
	if err != nil {
		return nil, err
	}

	return period, nil
}

func parseBool(boolStr, boolName string) (bool, error) {
	boolStr = strings.ToLower(strings.TrimSpace(boolStr))
	if boolStr == TRUE {
		return true, nil
	} else if boolStr == FALSE {
		return false, nil
	} else {
		return false, sdkerrors.Wrapf(types.ErrInvalidArgument, "%s is not a valid bool (true/false)", boolName)
	}
}

func NewTxCmd() *cobra.Command {
	paymentsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "payments transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	paymentsTxCmd.AddCommand(
		NewCmdCreatePaymentTemplate(),
		NewCmdCreatePaymentContract(),
		NewCmdCreateSubscription(),
		NewCmdSetPaymentContractAuthorisation(),
		NewCmdGrantPaymentDiscount(),
		NewCmdRevokePaymentDiscount(),
		NewCmdEffectPayment(),
	)

	return paymentsTxCmd
}

func NewCmdCreatePaymentTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-payment-template [payment-template-json] [creator-ixo-did]",
		Short: "Create and sign a create-payment-template tx using DIDs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			templateJsonStr := args[0]
			ixoDidStr := args[1]

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			var template types.PaymentTemplate
			err = json.Unmarshal([]byte(templateJsonStr), &template)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreatePaymentTemplate(template, ixoDid.Did)

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreatePaymentContract() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-payment-contract [payment-contract-id] [payment-template-id] " +
			"[payer-addr] [recipients] [can-deauthorise] [discount-id] [creator-ixo-did]",
		Short: "Create and sign a create-payment-contract tx using DIDs",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			templateIdStr := args[1]
			payerAddrStr := args[2]
			recipientsStr := args[3]
			canDeauthoriseStr := args[4]
			discountIdStr := args[5]
			ixoDidStr := args[6]

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

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			var recipients types.Distribution
			err = json.Unmarshal([]byte(recipientsStr), &recipients)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreatePaymentContract(templateIdStr,
				contractIdStr, payerAddr, recipients, canDeauthorise,
				discountId, ixoDid.Did)

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateSubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-subscription [subscription-id] [payment-contract-id] " +
			"[max-periods] [period-json] [creator-ixo-did]",
		Short: "Create and sign a create-subscription tx using DIDs",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			subIdStr := args[0]
			contractIdStr := args[1]
			maxPeriodsStr := args[2]
			periodStr := args[3]
			ixoDidStr := args[4]

			maxPeriods, err := sdk.ParseUint(maxPeriodsStr)
			if err != nil {
				return err
			}

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			//var p *period
			p, err := parsePeriodString(periodStr)

			var period types.Period
			switch p.Type {
			case BlockPeriodType:
				periodLength, _ := strconv.ParseInt(p.Value["period_length"], 10, 64)
				periodStartBlock, _ := strconv.ParseInt(p.Value["period_start_block"], 10, 64)
				period = &types.BlockPeriod{PeriodLength: periodLength, PeriodStartBlock: periodStartBlock}
			case TimePeriodType:
				periodDuration, _ := strconv.ParseInt(p.Value["period_duration_ns"], 10, 64)
				periodStartTime, _ := time.Parse(time.RFC3339, p.Value["period_start_time"])
				period = &types.TimePeriod{PeriodDurationNs: time.Duration(periodDuration), PeriodStartTime: periodStartTime}
			default:
				return sdkerrors.Wrapf(types.ErrInvalidArgument, "%s is not a valid Period", periodStr)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreateSubscription(subIdStr,
				contractIdStr, maxPeriods, period, ixoDid.Did)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdSetPaymentContractAuthorisation() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set-payment-contract-authorisation [payment-contract-id] " +
			"[authorised] [payer-ixo-did]",
		Short: "Create and sign a set-payment-contract-authorisation tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			authorisedStr := args[1]
			ixoDidStr := args[2]

			authorised, err := parseBool(authorisedStr, "authorised")
			if err != nil {
				return err
			}

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgSetPaymentContractAuthorisation(
				contractIdStr, authorised, ixoDid.Did)

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdGrantPaymentDiscount() *cobra.Command {
	cmd := &cobra.Command{
		Use: "grant-discount [payment-contract-id] [discount-id] " +
			"[recipient-addr] [creator-ixo-did]",
		Short: "Create and sign a grant-discount tx using DIDs",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			discountIdStr := args[1]
			recipientAddrStr := args[2]
			ixoDidStr := args[3]

			discountId, err := sdk.ParseUint(discountIdStr)
			if err != nil {
				return err
			}

			recipientAddr, err := sdk.AccAddressFromBech32(recipientAddrStr)
			if err != nil {
				return err
			}

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgGrantDiscount(
				contractIdStr, discountId, recipientAddr, ixoDid.Did)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdRevokePaymentDiscount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-discount [payment-contract-id] [holder-addr] [creator-ixo-did]",
		Short: "Create and sign a revoke-discount tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			holderAddrStr := args[1]
			ixoDidStr := args[2]

			holderAddr, err := sdk.AccAddressFromBech32(holderAddrStr)
			if err != nil {
				return err
			}

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgRevokeDiscount(
				contractIdStr, holderAddr, ixoDid.Did)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdEffectPayment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "effect-payment [payment-contract-id] [creator-ixo-did]",
		Short: "Create and sign a effect-payment tx using DIDs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractIdStr := args[0]
			ixoDidStr := args[1]

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgEffectPayment(contractIdStr, ixoDid.Did)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return ixotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), ixoDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
