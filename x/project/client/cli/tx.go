package cli

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	didtypes "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	ixotypes "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	projectTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "project transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	projectTxCmd.AddCommand(
		NewCmdCreateProject(),
		NewCmdCreateAgent(),
		NewCmdUpdateProjectStatus(),
		NewCmdUpdateAgent(),
		NewCmdCreateClaim(),
		NewCmdCreateEvaluation(),
		NewCmdWithdrawFunds(),
		NewCmdUpdateProjectDoc(),
	)

	return projectTxCmd
}

func NewCmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [sender-did] [project-iid-json] [ixo-did]",
		Short: "Create a new ProjectDoc signed by the ixoDid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDataStr := args[1]
			ixoDidStr := args[2]

			ixoDid, err := didtypes.UnmarshalIxoDid(ixoDidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreateProject(
				senderDid, json.RawMessage(projectDataStr), ixoDid.Did, ixoDid.VerifyKey)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			res, err := ixotypes.SignAndBroadcastTxFromStdSignMsg(clientCtx, msg, ixoDid, cmd.Flags())
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateProjectStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-status [sender-did] [status] [ixo-did]",
		Short: "Update the status of a project signed by the ixoDid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			status := args[1]
			ixoDid, err := didtypes.UnmarshalIxoDid(args[2])
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

			updateProjectStatusDoc := types.NewUpdateProjectStatusDoc(
				projectStatus, "")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgUpdateProjectStatus(senderDid, updateProjectStatusDoc, ixoDid.Did)
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

func NewCmdCreateAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-agent [tx-hash] [sender-did] [agent-did] " +
			"[role] [project-did]",
		Short: "Create a new agent on a project signed by the ixoDid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			role := args[3]
			if role != "SA" && role != "EA" && role != "IA" {
				return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
			}

			createAgentDoc := types.NewCreateAgentDoc(agentDid, role)

			ixoDid, err := didtypes.UnmarshalIxoDid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreateAgent(txHash, senderDid, createAgentDoc, ixoDid.Did)
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

func NewCmdUpdateAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-agent [tx-hash] [sender-did] [agent-did] " +
			"[status] [ixo-did]",
		Short: "Update the status of an agent on a project signed by the ixoDid of the project",
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

			updateAgentDoc := types.NewUpdateAgentDoc(
				agentDid, agentStatus, agentRole)

			ixoDid, err := didtypes.UnmarshalIxoDid(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgUpdateAgent(txHash, senderDid, updateAgentDoc, ixoDid.Did)
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

func NewCmdCreateClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-claim [tx-hash] [sender-did] [claim-id] [claim-template-id] [ixo-did]",
		Short: "Create a new claim on a project signed by the ixoDid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimTemplateId := args[3]
			createClaimDoc := types.NewCreateClaimDoc(claimId, claimTemplateId)

			ixoDid, err := didtypes.UnmarshalIxoDid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreateClaim(txHash, senderDid, createClaimDoc, ixoDid.Did)
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

func NewCmdCreateEvaluation() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-evaluation [tx-hash] [sender-did] [claim-id] " +
			"[status] [ixo-did]",
		Short: "Create a new claim evaluation on a project signed by the ixoDid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimStatus := types.ClaimStatus(args[3])
			if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			}

			createEvaluationDoc := types.NewCreateEvaluationDoc(
				claimId, claimStatus)

			ixoDid, err := didtypes.UnmarshalIxoDid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgCreateEvaluation(txHash, senderDid, createEvaluationDoc, ixoDid.Did)
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

func NewCmdWithdrawFunds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-funds [sender-did] [iid]",
		Short: "Withdraw funds.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ixoDid, err := didtypes.UnmarshalIxoDid(args[0])
			if err != nil {
				return err
			}

			var data types.WithdrawFundsDoc
			err = json.Unmarshal([]byte(args[1]), &data)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgWithdrawFunds(ixoDid.Did, data)
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

func NewCmdUpdateProjectDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-doc [sender-did] [project-iid-json] [ixo-did]",
		Short: "Update a project's iid signed by the ixoDid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDataStr := args[1]
			ixoDid, err := didtypes.UnmarshalIxoDid(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(ixoDid.Address())

			msg := types.NewMsgUpdateProjectDoc(senderDid, json.RawMessage(projectDataStr), ixoDid.Did)
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
