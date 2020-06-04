package cli

import (
	"encoding/json"
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
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

func unmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}

	return sovrinDid
}

func GetCmdCreateFee(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-fee [fee-json] [creator-sovrin-did]",
		Short: "Create and sign a create-fee tx using DIDs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			feeStr := args[0]
			sovrinDidStr := args[1]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			var fee types.Fee
			err := cdc.UnmarshalJSON([]byte(feeStr), &fee)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateFee(fee, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateFeeContract(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "create-fee-contract [fee-contract-id] [fee-id] [payer-addr] " +
			"[can-deauthorise] [discount-id] [creator-sovrin-did]",
		Short: "Create and sign a create-fee-contract tx using DIDs",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			feeContractIdStr := args[0]
			feeIdStr := args[1]
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

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			msg := types.NewMsgCreateFeeContract(
				feeIdStr, feeContractIdStr, payerAddr,
				canDeauthorise, discountId, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateSubscription(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "create-subscription [subscription-id] [fee-contract-id] " +
			"[max-periods] [period-json] [creator-sovrin-did]",
		Short: "Create and sign a create-subscription tx using DIDs",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			subIdStr := args[0]
			feeContractIdStr := args[1]
			maxPeriodsStr := args[2]
			periodStr := args[3]
			sovrinDidStr := args[4]

			maxPeriods, err := sdk.ParseUint(maxPeriodsStr)
			if err != nil {
				return err
			}

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			var period types.Period
			err = cdc.UnmarshalJSON([]byte(periodStr), &period)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSubscription(subIdStr,
				feeContractIdStr, maxPeriods, period, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdSetFeeContractAuthorisation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-fee-contract-authorisation [fee-contract-id] [authorised] [payer-sovrin-did]",
		Short: "Create and sign a set-fee-contract-authorisation tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			feeContractIdStr := args[0]
			authorisedStr := args[1]
			sovrinDidStr := args[2]

			authorised, err := parseBool(authorisedStr, "authorised")
			if err != nil {
				return err
			}

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			msg := types.NewMsgSetFeeContractAuthorisation(
				feeContractIdStr, authorised, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdGrantFeeDiscount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "grant-fee-discount [fee-contract-id] [discount-id] " +
			"[recipient-addr] [creator-sovrin-did]",
		Short: "Create and sign a grant-fee-discount tx using DIDs",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			feeContractIdStr := args[0]
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

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			msg := types.NewMsgGrantFeeDiscount(
				feeContractIdStr, discountId, recipientAddr, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdRevokeFeeDiscount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke-fee-discount [fee-contract-id] [holder-addr] [creator-sovrin-did]",
		Short: "Create and sign a revoke-fee-discount tx using DIDs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			feeContractIdStr := args[0]
			holderAddrStr := args[1]
			sovrinDidStr := args[2]

			holderAddr, err := sdk.AccAddressFromBech32(holderAddrStr)
			if err != nil {
				return err
			}

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			msg := types.NewMsgRevokeFeeDiscount(
				feeContractIdStr, holderAddr, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdChargeFee(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "charge-fee [fee-contract-id] [creator-sovrin-did]",
		Short: "Create and sign a charge-fee tx using DIDs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			feeContractIdStr := args[0]
			sovrinDidStr := args[1]

			sovrinDid := unmarshalSovrinDID(sovrinDidStr)

			msg := types.NewMsgChargeFee(feeContractIdStr, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}
