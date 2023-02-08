package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreateIidDocumentCmd(),
		NewUpdateIidDocumentCmd(),
		NewAddVerificationCmd(),
		NewAddServiceCmd(),
		NewRevokeVerificationCmd(),
		NewDeleteServiceCmd(),
		NewSetVerificationRelationshipCmd(),
		// TODO check if need aries agent creation
		// NewLinkAriesAgentCmd(),
		NewAddControllerCmd(),
		NewDeleteControllerCmd(),
		NewAddLinkedResourceCmd(),
		NewDeleteLinkedresourceCmd(),
		NewAddLinkedEntityCmd(),
		NewDeleteLinkedEntityCmd(),
		NewAddAccordedRightCmd(),
		NewDeleteAccordedRightCmd(),
		NewAddIidContextCmd(),
		NewDeleteIidContextCmd(),
		NewDeactivateIIDCmd(),
		NewCreateIidDocumentFormLegacyDidCmd(),
	)

	return cmd
}

// deriveVMType derive the verification method type from a public key
// func deriveVMType(pubKey cryptotypes.PubKey) (vmType types.VerificationMaterialType, err error) {
// 	switch pubKey.(type) {
// 	case *ed25519.PubKey:
// 		vmType = types.DIDVMethodTypeEd25519VerificationKey2018
// 	case *secp256k1.PubKey:
// 		vmType = types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019
// 	default:
// 		err = types.ErrKeyFormatNotSupported
// 	}
// 	return
// }

// NewCreateDidDocumentCmd defines the command to create a new IBC light client.
func NewCreateIidDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-iid [did-doc]",
		Short: "Create decentralized iid (did) document - flag is raw json with struct of MsgCreateIidDocument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgCreateIidDocument
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}
			var verJson types.VerificationsJSON
			if err := json.Unmarshal([]byte(args[0]), &verJson); err != nil {
				return err
			}

			// Manually gnerate verifications based of json values
			verifications, err := types.GenerateVerificationsFromJson(verJson)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.Verifications = verifications
			msg.Signer = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateIidDocumentCmd updates and iid document
// When using this function it updates all fields, even if dopnt provide fields it will use the proto defaults
func NewUpdateIidDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-iid [did-doc]",
		Short: "updates and iid document - flag is raw json with struct of MsgUpdateIidDocument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg types.MsgUpdateIidDocument
			if err := json.Unmarshal([]byte(args[0]), &msg); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var verJson types.VerificationsJSON
			if err := json.Unmarshal([]byte(args[0]), &verJson); err != nil {
				return err
			}

			// Manually generate verifications based of json values
			verifications, err := types.GenerateVerificationsFromJson(verJson)
			if err != nil {
				return err
			}

			msg.Verifications = verifications
			msg.Signer = clientCtx.GetFromAddress().String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewAddVerificationCmd define the command to add a verification message
func NewAddVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-verification-method [id] [verification]",
		Short: "add an verification method to an iid document - verification is raw json of struct Verification",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var verJson types.VerificationsJSON
			if err := json.Unmarshal([]byte(args[0]), &verJson); err != nil {
				return err
			}

			// Manually gnerate verifications based of json values
			verifications, err := types.GenerateVerificationsFromJson(verJson)
			if err != nil {
				return err
			}
			if len(verifications) == 0 {
				return fmt.Errorf("no verification provided")
			}

			msg := types.NewMsgAddVerification(
				args[0],
				verifications[0],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-service [id] [service-id] [type] [endpoint]",
		Short: "add a service to an iid document",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, serviceID, serviceType, endpoint := args[0], args[1], args[2], args[3]

			service := types.NewService(
				serviceID,
				serviceType,
				endpoint,
			)

			msg := types.NewMsgAddService(
				id,
				service,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRevokeVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-verification-method [id] [method-id]",
		Short: "revoke a verification method from an iid document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeVerification(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewDeleteServiceCmd deletes a service from a DID Document
func NewDeleteServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-service [id] [service-id]",
		Short: "deletes a service from an iid document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteService(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewAddControllerCmd adds a controller to a did document
func NewAddControllerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-controller [id] [controller-did]",
		Short: "updates an iid document to contain a controller",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddController(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewDeleteControllerCmd removes a controller from a did document
func NewDeleteControllerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-controller [id] [controller-did]",
		Short: "updates aan iid document by removing a controller",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteController(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewSetVerificationRelationshipCmd adds a verification relationship to a verification method
func NewSetVerificationRelationshipCmd() *cobra.Command {
	var relationships []string
	// if true do not add the default authentication relationship
	var unsafe bool

	cmd := &cobra.Command{
		Use:   "set-verification-relationship [id] [method-id] --relationship NAME [--relationship NAME ...]",
		Short: "sets one or more verification relationships to a key on an iid document.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetVerificationRelationships(
				args[0],
				args[1],
				relationships,
				clientCtx.GetFromAddress().String(),
			)

			// make sure that the authentication relationship is preserved
			if !unsafe {
				msg.Relationships = append(msg.Relationships, types.Authentication)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags to set did relationships
	cmd.Flags().StringSliceVarP(&relationships, "relationship", "r", []string{}, "the relationships to set for the verification method in the DID")
	cmd.Flags().BoolVar(&unsafe, "unsafe", false, fmt.Sprint("do not ensure that '", types.Authentication, "' relationship is set"))

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddLinkedResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-linked-resource [id] [resource-id] [type] [description] [media-type] [service-endpoint] [proof] [encrypted] [privacy]",
		Short: "add a linked resource to an iid document",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, resourceId, serviceType, desc, mediaType, endpoint, proof, encrypted, privacy := args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8]

			resource := types.NewLinkedResource(
				resourceId,
				serviceType,
				desc,
				mediaType,
				endpoint,
				proof,
				encrypted,
				privacy,
			)

			msg := types.NewMsgAddLinkedResource(
				id,
				resource,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDeleteLinkedresourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-resource [id] [resource-id]",
		Short: "deletes a resource from an iid document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteLinkedResource(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddAccordedRightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-accorded-right [id] [right-id] [type] [mechanism] [message] [service-endpoint]",
		Short: "add an Accorded Right to an iid document",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, rightId, rightType, mechanism, message, endpoint := args[0], args[1], args[2], args[3], args[4], args[5]

			right := types.NewAccordedRight(
				rightId,
				rightType,
				mechanism,
				message,
				endpoint,
			)

			msg := types.NewMsgAddAccordedRight(
				id,
				right,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDeleteAccordedRightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-accorded-right [id] [resource-id]",
		Short: "deletes a right from an iid document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAccordedRight(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddIidContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-iid-context [id] [key] [value]",
		Short: "add a context item to an iid document",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, key, value := args[0], args[1], args[2]

			didContext := types.NewDidContext(
				key,
				value,
			)

			msg := types.NewMsgAddDidContext(
				id,
				didContext,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDeleteIidContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-context [id] [key]",
		Short: "deletes a iid context from an iid document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDidContext(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddLinkedEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-linked-entity [id] [entity-id] [type] [relationship]",
		Short: "add a linked entity to an iid document",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, entityId, entityType, relationship := args[0], args[1], args[2], args[3]

			entity := types.NewLinkedEntity(
				entityId,
				entityType,
				relationship,
			)

			msg := types.NewMsgAddLinkedEntity(
				id,
				entity,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDeleteLinkedEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-linked-entity [id] [entity-id]",
		Short: "deletes an entity from an iid document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteLinkedEntity(
				args[0],
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDeactivateIIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-iid [id] [state]",
		Short: "changes (deactivates) the deactivated field off iid metadata",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newState, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeactivateIID(
				args[0],
				newState,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
