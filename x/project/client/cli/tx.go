package cli

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func GetCmdCreateProject(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-project [sender-did] [project-json] [sovrin-did]",
		Short: "Create a new ProjectDoc signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDocStr := args[1]
			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[2])
			if err != nil {
				return err
			}

			var projectDoc types.ProjectDoc
			err = json.Unmarshal([]byte(projectDocStr), &projectDoc)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreateProject(senderDid, projectDoc, sovrinDid)
			stdSignMsg := msg.ToStdSignMsg(types.MsgCreateProjectFee)

			res, err := ixo.SignAndBroadcastTxFromStdSignMsg(cliCtx, stdSignMsg, sovrinDid)
			if err != nil {
				return err
			}

			fmt.Println(res.String())
			fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.TxHash)
			return nil
		},
	}
}

func GetCmdUpdateProjectStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-project-status [sender-did] [status] [sovrin-did]",
		Short: "Update the status of a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			status := args[1]
			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[2])
			if err != nil {
				return err
			}

			projectStatus := types.ProjectStatus(status)
			if projectStatus != types.CreatedProject &&
				projectStatus != types.PendingStatus &&
				projectStatus != types.FundedStatus &&
				projectStatus != types.StartedStatus &&
				projectStatus != types.StoppedStatus &&
				projectStatus != types.PaidoutStatus {
				return errors.New("The status must be one of 'CREATED', " +
					"'PENDING', 'FUNDED', 'STARTED', 'STOPPED' or 'PAIDOUT'")
			}

			updateProjectStatusDoc := types.UpdateProjectStatusDoc{
				Status: projectStatus,
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgUpdateProjectStatus(senderDid, updateProjectStatusDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateAgent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "create-agent [tx-hash] [sender-did] [agent-did] " +
			"[role] [project-did]",
		Short: "Create a new agent on a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			role := args[3]
			if role != "SA" && role != "EA" && role != "IA" {
				return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
			}

			createAgentDoc := types.CreateAgentDoc{
				AgentDid: agentDid,
				Role:     role,
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[4])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreateAgent(txHash, senderDid, createAgentDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdUpdateAgent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "update-agent [tx-hash] [sender-did] [agent-did] " +
			"[status] [sovrin-did]",
		Short: "Update the status of an agent on a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			agentRole := args[4]
			agentStatus := types.AgentStatus(args[3])
			if agentStatus != types.PendingAgent && agentStatus != types.ApprovedAgent && agentStatus != types.RevokedAgent {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
			}

			updateAgentDoc := types.UpdateAgentDoc{
				Did:    agentDid,
				Status: agentStatus,
				Role:   agentRole,
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[5])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgUpdateAgent(txHash, senderDid, updateAgentDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateClaim(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-claim [tx-hash] [sender-did] [claim-id] [sovrin-did]",
		Short: "Create a new claim on a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			createClaimDoc := types.CreateClaimDoc{
				ClaimID: claimId,
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[3])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreateClaim(txHash, senderDid, createClaimDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateEvaluation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "create-evaluation [tx-hash] [sender-did] [claim-id] " +
			"[status] [sovrin-did]",
		Short: "Create a new claim evaluation on a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimStatus := types.ClaimStatus(args[3])
			if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			}

			createEvaluationDoc := types.CreateEvaluationDoc{
				ClaimID: claimId,
				Status:  claimStatus,
			}

			sovrinDid, err := sovrin.UnmarshalSovrinDid(args[4])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(sovrinDid.Did))

			msg := types.NewMsgCreateEvaluation(txHash, senderDid, createEvaluationDoc, sovrinDid)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, sovrinDid)
		},
	}
}

func GetCmdWithdrawFunds(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-funds [sender-did] [data]",
		Short: "Withdraw funds.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid, err := sovrin.UnmarshalSovrinDid(args[0])
			if err != nil {
				return err
			}

			var data types.WithdrawFundsDoc
			err = json.Unmarshal([]byte(args[1]), &data)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc).
				WithFromAddress(ixo.DidToAddr(senderDid.Did))

			msg := types.NewMsgWithdrawFunds(senderDid.Did, data)

			return ixo.SignAndBroadcastTxCli(cliCtx, msg, senderDid)
		},
	}
}
