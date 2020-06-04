package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func unmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}

	return sovrinDid
}

func GetCmdCreateProject(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-project [sender-did] [project-json] [sovrin-did]",
		Short: "Create a new ProjectDoc signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			senderDid := args[0]
			projectDocStr := args[1]
			sovrinDid := unmarshalSovrinDID(args[2])

			projectDoc := types.ProjectDoc{}
			err := json.Unmarshal([]byte(projectDocStr), &projectDoc)
			if err != nil {
				panic(err)
			}

			msg := types.NewMsgCreateProject(senderDid, projectDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdUpdateProjectStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-project-status [sender-did] [status] [sovrin-did]",
		Short: "Update the status of a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			senderDid := args[0]
			status := args[1]
			sovrinDid := unmarshalSovrinDID(args[2])

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

			msg := types.NewMsgUpdateProjectStatus(senderDid, updateProjectStatusDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
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
			ctx := context.NewCLIContext().WithCodec(cdc)

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

			sovrinDid := unmarshalSovrinDID(args[4])
			msg := types.NewMsgCreateAgent(txHash, senderDid, createAgentDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
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
			ctx := context.NewCLIContext().WithCodec(cdc)

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

			sovrinDid := unmarshalSovrinDID(args[5])
			msg := types.NewMsgUpdateAgent(txHash, senderDid, updateAgentDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateClaim(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-claim [tx-hash] [sender-did] [claim-id] [sovrin-did]",
		Short: "Create a new claim on a project signed by the sovrinDID of the project",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			createClaimDoc := types.CreateClaimDoc{
				ClaimID: claimId,
			}

			sovrinDid := unmarshalSovrinDID(args[3])
			msg := types.NewMsgCreateClaim(txHash, senderDid, createClaimDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
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
			ctx := context.NewCLIContext().WithCodec(cdc)

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

			sovrinDid := unmarshalSovrinDID(args[4])
			msg := types.NewMsgCreateEvaluation(txHash, senderDid, createEvaluationDoc, sovrinDid)

			return ixo.SignAndBroadcastCli(ctx, msg, sovrinDid)
		},
	}
}

func GetCmdWithdrawFunds(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-funds [sender-did] [data]",
		Short: "Withdraw funds.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			senderDid := unmarshalSovrinDID(args[0])
			var data types.WithdrawFundsDoc
			err := json.Unmarshal([]byte(args[1]), &data)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawFunds(senderDid.Did, data)

			return ixo.SignAndBroadcastCli(ctx, msg, senderDid)
		},
	}
}
