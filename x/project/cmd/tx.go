package commands

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"

	ixo "github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/project"

	base58 "github.com/btcsuite/btcutil/base58"
)

func ixoSignAndBroadcast(cdc *wire.Codec, ctx core.CoreContext, msg sdk.Msg, sovrinDid sovrin.SovrinDid) error {
	// Force the length to 64
	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

	//Create the Signature
	signature := ixo.SignIxoMessage(msg, sovrinDid.Did, privKey)
	fmt.Println(signature)

	tx := ixo.NewIxoTx(msg, signature)

	fmt.Println("*******TRANSACTION******* \n", tx.String())

	bz, err := cdc.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}

	var txMap map[string]interface{}
	if err := json.Unmarshal(bz, &txMap); err != nil {
		panic(err)
	}
	fmt.Println(">txMap : ", txMap)
	txPayload := txMap["payload"].([]interface{})
	txType := txPayload[0]
	txMessage := txPayload[1]
	txSignature := txMap["signature"]

	txMessageJSON, err := json.Marshal(txMessage)
	if err != nil {
		panic(err)
	}

	newTxPayload := []interface{}{txType, hex.EncodeToString(txMessageJSON)}
	txMap["payload"] = newTxPayload

	newTxMap := map[string]interface{}{"payload": newTxPayload, "signature": txSignature}
	fmt.Println(">newTxMap : ", newTxMap)

	newBroadcast, err := json.Marshal(newTxMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n***\n>newBroadcast : ", string(newBroadcast))

	// Broadcast to Tendermint
	res, err := ctx.BroadcastTx(newBroadcast)
	if err != nil {
		return err
	}

	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
	return nil

}

// Unmarshal sovrinDID
func unmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
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

			//			sig := sovrin.SignMessage([]byte("A"), "BGsfxvVMmzcUEnYmXvso3fEGjmbk5He9HpuuPvEdPNUg", "6xRWNCMc4CiJ2A3kgjpFdkJFT7oboy41dtjB8J1F7U52")
			//			fmt.Println(hex.EncodeToString(sig))
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

// Update Project Status
func UpdateProjectStatusCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "updateProjectStatus status ethFundingTxnID projectDid",
		Short: "Update a a project status and ethereum funding txn id signed by the sovrinDID of the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			if len(args) != 4 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 {
				return errors.New("You must provide the parameters")
			}

			projectStatus := project.ProjectStatus(args[0])
			if projectStatus != project.NullStatus && projectStatus != project.PendingStatus && projectStatus != project.FundedStatus && projectStatus != project.StartedStatus && projectStatus != project.StoppedStatus && projectStatus != project.PaidoutStatus {
				return errors.New("The status must be one of 'CREATED', 'PENDING', 'FUNDED', 'STARTED', 'STOPPED', 'PAIDOUT'")
			}

			ethFundingTxnID := args[1]

			// projectID := args[2]

			updateProjectDoc := project.UpdateProjectStatusDoc{
				Status:          projectStatus,
				EthFundingTxnID: ethFundingTxnID,
			}

			projectDid := unmarshalSovrinDID(args[3])

			// create the message
			msg := project.NewUpdateProjectStatusMsg(updateProjectDoc, projectDid)

			return ixoSignAndBroadcast(cdc, ctx, msg, projectDid)
		},
	}
}

// Create Agent
func CreateAgentCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createAgent txHash senderDid agentDid role projectDid",
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
				AgentDid: agentDid,
				Role:     role,
			}

			sovrinDid := unmarshalSovrinDID(args[4])

			// create the message
			msg := project.NewCreateAgentMsg(txHash, senderDid, createAgentDoc, sovrinDid)

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
			if len(args) != 6 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 || len(args[4]) == 0 || len(args[5]) == 0 {
				return errors.New("You must provide the agentDid, status and the projects private key")
			}

			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			agentRole := args[4]

			agentStatus := project.AgentStatus(args[3])
			if agentStatus != project.PendingAgent && agentStatus != project.ApprovedAgent && agentStatus != project.RevokedAgent {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
			}

			updateAgentDoc := project.UpdateAgentDoc{
				Did:    agentDid,
				Status: agentStatus,
				Role:   agentRole,
			}

			sovrinDid := unmarshalSovrinDID(args[5])

			// create the message
			msg := project.NewUpdateAgentMsg(txHash, senderDid, updateAgentDoc, sovrinDid)

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
				ClaimID: claimId,
			}

			sovrinDid := unmarshalSovrinDID(args[3])

			// create the message
			msg := project.NewCreateClaimMsg(txHash, senderDid, createClaimDoc, sovrinDid)

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

			claimStatus := project.ClaimStatus(args[3])
			if claimStatus != project.PendingClaim && claimStatus != project.ApprovedClaim && claimStatus != project.RejectedClaim {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			}

			createEvaluationDoc := project.CreateEvaluationDoc{
				ClaimID: claimId,
				Status:  claimStatus,
			}

			sovrinDid := unmarshalSovrinDID(args[4])

			// create the message
			msg := project.NewCreateEvaluationMsg(txHash, senderDid, createEvaluationDoc, sovrinDid)

			return ixoSignAndBroadcast(cdc, ctx, msg, sovrinDid)
		},
	}
}

// take the coolness quiz transaction
func FundProjectTxCmd(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "fundProject projectDid ethTx amount sovrinDid",
		Short: "Create tokens to fund project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 4 || len(args[0]) == 0 || len(args[1]) == 0 || len(args[2]) == 0 || len(args[3]) == 0 {
				return errors.New("You must provide a valid projectDid, ethereum transaction hash, amount of ixo and sovrinDid")
			}

			ctx := context.NewCoreContextFromViper()

			// create the message
			fundProjectDoc := project.FundProjectDoc{
				ProjectDid: args[0],
				EthTxHash:  args[1],
				Amount:     args[2],
			}
			sovrinDid := unmarshalSovrinDID(args[3])

			// create the message
			msg := project.NewFundProjectMsg(args[1], "", fundProjectDoc, sovrinDid)

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

func GetProjectAccountsCmd(storeName string) *cobra.Command {
	return &cobra.Command{
		Use:   "getProjectAccounts did",
		Short: "Get a Project accounts for a Did",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()

			if len(args) != 1 || len(args[0]) == 0 {
				return errors.New("You must provide a project did")
			}

			// find the key to look up the account
			projectDid := args[0]
			var buffer bytes.Buffer
			buffer.WriteString("ACC-")
			buffer.WriteString(projectDid)
			key := buffer.Bytes()

			res, err := ctx.Query(key, storeName)
			if err != nil {
				return err
			}

			// the query will return empty if there is no data for this did
			if len(res) == 0 {
				return errors.New("Project does not exist")
			}

			// decode the value
			var f interface{}
			err = json.Unmarshal(res, &f)
			if err != nil {
				return err
			}
			accMap := f.(map[string]interface{})

			// print out whole didDoc
			output, err := json.MarshalIndent(accMap, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
}
