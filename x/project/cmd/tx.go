package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"

	ixo "github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project"

	base58 "github.com/btcsuite/btcutil/base58"
)

func ixoSignAndBroadcast(cdc *wire.Codec, ctx core.CoreContext, msg sdk.Msg, sovrinDid ixo.SovrinDid) error {
	// Force the length to 64
	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

	//Create the Signature
	signature := ixo.SignIxoMessage(msg, sovrinDid.Did, privKey)

	tx := ixo.NewIxoTx(msg, signature)

	fmt.Println("*******TRANSACTION******* \n", tx.String())

	bz, err := cdc.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}
	// Broadcast to Tendermint
	res, err := ctx.BroadcastTx(bz)
	if err != nil {
		return err
	}

	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
	return nil

}

// Unmarshal sovrinDID
func unmarshalSovrinDID(sovrinJson string) ixo.SovrinDid {
	sovrinDid := ixo.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}
	return sovrinDid
}

// Create a project doc to the ledger
func CreateProjectCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createProject projectJson sovrinDiD",
		Short: "Create a new ProjectDoc signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
				return errors.New("You must provide the project data and the projects private key")
			}

			projectDoc := project.ProjectDoc{}
			err := json.Unmarshal([]byte(args[0]), &projectDoc)
			if err != nil {
				panic(err)
			}

			sovrinDid := unmarshalSovrinDID(args[1])

			// create the message
			msg := project.NewCreateProjectMsg(projectDoc, sovrinDid)

			return ixoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

// Create Agent
func CreateAgentCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createAgent txHash senderDid agentDid role sovrinDid",
		Short: "Create a new agent on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 5 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 || len(args[4]) == 0 {
				return errors.New("You must provide the agentDid, role and the projects private key")
			}

			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]

			role := args[3]
			if role != "SA" && role != "EA" && role != "IA" {
				return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
			}

			createAgentDoc := project.CreateAgentDoc{
				TxHash:    txHash,
				SenderDid: senderDid,
				Did:       agentDid,
				Role:      role,
			}

			// create the message
			msg := project.NewCreateAgentMsg(createAgentDoc)

			sovrinDid := unmarshalSovrinDID(args[4])

			return ixoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

// Update Agent
func UpdateAgentCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "updateAgent txHash senderDid agentDid status sovrinDid",
		Short: "Update the status of an agent on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 5 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 || len(args[4]) == 0 {
				return errors.New("You must provide the agentDid, status and the projects private key")
			}

			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]

			status, err := strconv.Atoi(args[3])
			if err != nil {
				panic(err)
			}
			agentStatus := project.AgentStatus(status)
			if agentStatus != project.PendingAgent && agentStatus != project.ApprovedAgent && agentStatus != project.RevokedAgent {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
			}

			updateAgentDoc := project.UpdateAgentDoc{
				TxHash:    txHash,
				SenderDid: senderDid,
				Did:       agentDid,
				Status:    agentStatus,
			}

			// create the message
			msg := project.NewUpdateAgentMsg(updateAgentDoc)

			sovrinDid := unmarshalSovrinDID(args[4])

			return ixoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

// Create Claim
func CreateClaimCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createClaim txHash senderDid claimId sovrinDid",
		Short: "Create a new claim on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 4 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 {
				return errors.New("You must provide the claimId and the projects private key")
			}

			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]

			createClaimDoc := project.CreateClaimDoc{
				TxHash:    txHash,
				SenderDid: senderDid,
				ClaimID:   claimId,
			}

			// create the message
			msg := project.NewCreateClaimMsg(createClaimDoc)

			sovrinDid := unmarshalSovrinDID(args[3])

			return ixoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

// Create Evaluation
func CreateEvaluationCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createEvaluation txHash senderDid claimId status sovrinDid",
		Short: "Create a new claim evaluation on a project signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 5 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 || len(args[4]) == 0 || len(args[2]) == 0 {
				return errors.New("You must provide the claimId, status and the projects private key")
			}

			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]

			status, err := strconv.Atoi(args[3])
			if err != nil {
				panic(err)
			}
			claimStatus := project.ClaimStatus(status)
			if claimStatus != project.PendingClaim && claimStatus != project.ApprovedClaim && claimStatus != project.RejectedClaim {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			}

			createEvaluationDoc := project.CreateEvaluationDoc{
				TxHash:    txHash,
				SenderDid: senderDid,
				ClaimID:   claimId,
				Status:    claimStatus,
			}

			// create the message
			msg := project.NewCreateEvaluationMsg(createEvaluationDoc)

			sovrinDid := unmarshalSovrinDID(args[4])

			return ixoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

// Get a project doc from the ledger
func GetProjectDocCmd(storeName string, cdc *wire.Codec, decoder project.ProjectDocDecoder) *cobra.Command {
	return &cobra.Command{
		Use:   "getProjectDoc did",
		Short: "Get a new ProjectDoc for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide an did")
			}

			// find the key to look up the account
			didAddr := args[0]
			key := ixo.Did(didAddr)

			res, err := ctx.Query([]byte(key), storeName)
			if err != nil {
				return err
			}

			// decode the value
			projectDoc, err := decoder(res)
			if err != nil {
				return err
			}

			// print out whole account
			output, err := json.MarshalIndent(projectDoc, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))

			return nil
		},
	}
}
