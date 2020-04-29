package cli

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func IxoSignAndBroadcast(cdc *codec.Codec, ctx context.CLIContext, msg sdk.Msg, sovrinDid sovrin.SovrinDid) error {
	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
	tx := ixo.NewIxoTxSingleMsg(msg, signature)

	bz, err := cdc.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}

	res, err := ctx.BroadcastTx(bz)
	if err != nil {
		return err
	}

	fmt.Println(res.String())
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.TxHash)
	return nil

}

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
		Use:   "createProject [project-json] [sovrin-did]",
		Short: "Create a new ProjectDoc signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the project data and the projects private key")
			}

			projectDoc := types.ProjectDoc{}
			err := json.Unmarshal([]byte(args[0]), &projectDoc)
			if err != nil {
				panic(err)
			}

			sovrinDid := unmarshalSovrinDID(args[1])
			msg := types.NewMsgCreateProject(projectDoc, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdUpdateProjectStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "updateProjectStatus [sender-did] [status] [sovrin-did]",
		Short: "Update the status of a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 3 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 {
				return errors.New("You must provide the status and the projects private key")
			}

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

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateAgent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createAgent [tx-hash] [sender-did] [agent-did] [role] [project-did]",
		Short: "Create a new agent on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 5 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 ||
				len(args[4]) == 0 {
				return errors.New("You must provide the agentDid, role and the projects private key")
			}

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

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdUpdateAgent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "updateAgent [tx-hash] [sender-did] [agent-did] [status] [sovrin-did]",
		Short: "Update the status of an agent on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 6 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 ||
				len(args[4]) == 0 || len(args[5]) == 0 {
				return errors.New("You must provide the agentDid, status and the projects private key")
			}

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

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateClaim(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createClaim [tx-hash] [sender-did] [claim-id] [sovrin-did]",
		Short: "Create a new claim on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 4 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 {
				return errors.New("You must provide the claimId and the projects private key")
			}

			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			createClaimDoc := types.CreateClaimDoc{
				ClaimID: claimId,
			}

			sovrinDid := unmarshalSovrinDID(args[3])
			msg := types.NewMsgCreateClaim(txHash, senderDid, createClaimDoc, sovrinDid)

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdCreateEvaluation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createEvaluation [tx-hash] [sender-did] [claim-id] [status] [sovrin-did]",
		Short: "Create a new claim evaluation on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 5 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 ||
				len(args[4]) == 0 {
				return errors.New("You must provide the claimId, status and the projects private key")
			}

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

			return IxoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

func GetCmdWithdrawFunds(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdrawFunds [sender-did] [data]",
		Short: "withdraw funds.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().
				WithCodec(cdc)

			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the sender did and data.")
			}

			senderDid := unmarshalSovrinDID(args[0])
			var data types.WithdrawFundsDoc
			err := json.Unmarshal([]byte(args[1]), &data)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawFunds(senderDid.Did, data)

			return IxoSignAndBroadcast(cdc, ctx, msg, senderDid)
		},
	}
}
