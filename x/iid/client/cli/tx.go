package cli

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		NewCreateIidDocumentCmd(),
		NewAddVerificationCmd(),
		NewAddServiceCmd(),
		NewRevokeVerificationCmd(),
		NewDeleteServiceCmd(),
		NewSetVerificationRelationshipCmd(),
		NewLinkAriesAgentCmd(),
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
		NewUpdateIidMetaCmd(),
	)

	return cmd
}

// deriveVMType derive the verification method type from a public key
func deriveVMType(pubKey cryptotypes.PubKey) (vmType types.VerificationMaterialType, err error) {
	switch pubKey.(type) {
	case *ed25519.PubKey:
		vmType = types.DIDVMethodTypeEd25519VerificationKey2018
	case *secp256k1.PubKey:
		vmType = types.DIDVMethodTypeEcdsaSecp256k1VerificationKey2019
	default:
		err = types.ErrKeyFormatNotSupported
	}
	return
}

// NewCreateDidDocumentCmd defines the command to create a new IBC light client.
func NewCreateIidDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-iid [id] [chain name]",
		Short:   "create decentralized did (did) document",
		Example: "creates a did document for users",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// did
			did := types.NewChainDID(args[1], args[0])
			// verification
			signer := clientCtx.GetFromAddress()
			// pubkey
			info, err := clientCtx.Keyring.KeyByAddress(signer)
			if err != nil {
				return err
			}
			pubKey := info.GetPubKey()
			// verification method id
			vmID := did.NewVerificationMethodID(signer.String())
			// understand the vmType
			vmType, err := deriveVMType(pubKey)
			if err != nil {
				return err
			}
			auth := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					did,
					types.NewPublicKeyMultibase(pubKey.Bytes(), vmType),
				),
				[]string{types.Authentication},
				nil,
			)
			// create the message
			msg := types.NewMsgCreateIidDocument(
				did.String(),
				types.Verifications{auth},
				types.Services{},
				types.AccordedRights{},
				types.LinkedResources{},
				types.LinkedEntities{},
				signer.String(),
				types.Contexts{},
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewAddVerificationCmd define the command to add a verification message
func NewAddVerificationCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "add-verification-method [id] [pubkey]",
		Short:   "add an verification method to an (iid) document",
		Example: "adds an verification method for an iid document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// signer address
			signer := clientCtx.GetFromAddress()
			// public key
			var pk cryptotypes.PubKey
			err = clientCtx.Codec.UnmarshalInterfaceJSON([]byte(args[1]), &pk)
			if err != nil {
				return err
			}
			// derive the public key type
			vmType, err := deriveVMType(pk)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// verification method id
			vmID := did.NewVerificationMethodID(sdk.MustBech32ifyAddressBytes(
				sdk.GetConfig().GetBech32AccountAddrPrefix(),
				pk.Address().Bytes(),
			))

			verification := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					did,
					types.NewPublicKeyMultibase(pk.Bytes(), vmType),
				),
				[]string{types.Authentication},
				nil,
			)
			// add verification
			msg := types.NewMsgAddVerification(
				did.String(),
				verification,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAddServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-service [id] [service_id] [type] [endpoint]",
		Short:   "add a service to a decentralized did (did) document",
		Example: "adds a service to a did document",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			serviceID, serviceType, endpoint := args[1], args[2], args[3]
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			service := types.NewService(
				serviceID,
				serviceType,
				endpoint,
			)

			msg := types.NewMsgAddService(
				did.String(),
				service,
				signer.String(),
			)
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewRevokeVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-verification-method [did_id] [verification_method_id_fragment]",
		Short: "revoke a verification method from a decentralized did (did) document",
		Example: `cosmos-cashd tx did revoke-verification-method 575d062c-d110-42a9-9c04-cb1ff8c01f06 \
 Z46DAL1MrJlVW_WmJ19WY8AeIpGeFOWl49Qwhvsnn2M \
 --from alice \
 --node https://rpc.cosmos-cash.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// verification method id
			vmID := did.NewVerificationMethodID(args[1])
			// build the message
			msg := types.NewMsgRevokeVerification(
				did.String(),
				vmID,
				signer.String(),
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewDeleteServiceCmd deletes a service from a DID Document
func NewDeleteServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-service [id] [service-id]",
		Short:   "deletes a service from a decentralized did (did) document",
		Example: "delete a service for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// service id
			sID := args[1]

			msg := types.NewMsgDeleteService(
				did.String(),
				sID,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewAddControllerCmd adds a controller to a did document
func NewAddControllerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-controller [id] [controllerAddress]",
		Short:   "updates a decentralized identifier (did) document to contain a controller",
		Example: "add-controller vasp cosmos1kslgpxklq75aj96cz3qwsczr95vdtrd3p0fslp",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			// did key to use as the controller
			didKey := types.NewKeyDID(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := types.NewMsgAddController(
				did.String(),
				didKey.String(),
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewDeleteControllerCmd adds a controller to a did document
func NewDeleteControllerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-controller [id] [controllerAddress]",
		Short:   "updates a decentralized identifier (did) document removing a controller",
		Example: "delete-controller vasp cosmos1kslgpxklq75aj96cz3qwsczr95vdtrd3p0fslp",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			// did key to use as the controller
			didKey := types.NewKeyDID(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := types.NewMsgDeleteController(
				did.String(),
				didKey.String(),
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewSetVerificationRelationshipCmd adds a verification relationship to a verification method
func NewSetVerificationRelationshipCmd() *cobra.Command {

	// relationships
	var relationships []string
	// if true do not add the default authentication relationship
	var unsafe bool

	cmd := &cobra.Command{
		Use:     "set-verification-relationship [did_id] [verification_method_id_fragment] --relationship NAME [--relationship NAME ...]",
		Short:   "sets one or more verification relationships to a key on a decentralized identifier (did) document.",
		Example: "set-verification-relationship vasp 6f1e0700-6c86-41b6-9e05-ae3cf839cdd0 --relationship capabilityInvocation",
		Args:    cobra.ExactArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			// method id
			vmID := did.NewVerificationMethodID(args[1])

			// signer
			signer := clientCtx.GetFromAddress()

			msg := types.NewMsgSetVerificationRelationships(
				did.String(),
				vmID,
				relationships,
				signer.String(),
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

////////////////////
// IID Extension
////////////////////

func NewAddLinkedResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-linked-resource [id] [resource_id] [type] [description] [media_type] [service_endpoint] [proof] [encrypted] [privacy]",
		Short:   "add a linked resource to a decentralized did (did/IID) document",
		Example: "adds a linked resource to a did document",
		Args:    cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			resourceId, resourceType, desc, mediaType, endpoint, proof, encrypted, privacy := args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8]
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			resource := types.NewLinkedResource(
				resourceId,
				resourceType,
				desc,
				mediaType,
				endpoint,
				proof,
				encrypted,
				privacy,
			)

			msg := types.NewMsgAddLinkedResource(
				did.String(),
				resource,
				signer.String(),
			)
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDeleteLinkedresourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-resource [id] [resource-id]",
		Short:   "deletes a resource from a decentralized did (did) document",
		Example: "delete a resource for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// resource id
			rID := args[1]

			msg := types.NewMsgDeleteLinkedResource(
				did.String(),
				rID,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAddAccordedRightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-accorded-right [id] [right_id] [type] [mechanism] [message] [service_endpoint]",
		Short:   "add an Accorded Right to a decentralized did (did/IID) document",
		Example: "adds an accorded right to a did document",
		Args:    cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// Right parameters
			rightId, rightType, mechanism, message, endpoint := args[1], args[2], args[3], args[4], args[5]
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			right := types.NewAccordedRight(
				rightId,
				rightType,
				mechanism,
				message,
				endpoint,
			)

			msg := types.NewMsgAddAccordedRight(
				did.String(),
				right,
				signer.String(),
			)
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDeleteAccordedRightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-accorded-right [id] [resource-id]",
		Short:   "deletes a right from a decentralized did (did) document",
		Example: "delete a right for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// resource id
			rID := args[1]

			msg := types.NewMsgDeleteAccordedRight(
				did.String(),
				rID,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAddIidContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-iid-context [iid-id] [key] [value]",
		Short:   "add a context item to a decentralized (did/IID) document",
		Example: "adds a context item to a iid document",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			key := args[1]
			val := args[2]
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			didContext := types.NewDidContext(
				key,
				val,
			)

			msg := types.NewMsgAddDidContext(
				did.String(),
				didContext,
				signer.String(),
			)
			fmt.Println(msg)
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDeleteIidContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-context [iid-id] [key]",
		Short:   "deletes a iid context from a decentralized iid document",
		Example: "delete a context for a iid document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// resource id
			cID := args[1]

			msg := types.NewMsgDeleteDidContext(
				did.String(),
				cID,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewUpdateIidMetaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-iid-meta [iid-id] [meta]",
		Short:   "add a context item to a decentralized (did/IID) document",
		Example: "adds a context item to a iid document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			var metaData *types.IidMetadata

			if err := json.Unmarshal([]byte(args[1]), &metaData); err != nil {
				panic(err)
				err = fmt.Errorf(err.Error())
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			msg := types.NewMsgUpdateDidMetaData(
				did.String(),
				metaData,
				signer.String(),
			)
			fmt.Println(msg)

			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

//Linked Entity
func NewAddLinkedEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-linked-resource [id] [relationship id] [relationship]",
		Short:   "add a linked entity to a decentralized did (did/IID) document",
		Example: "adds a linked entity to a did document",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// tx signer
			signer := clientCtx.GetFromAddress()
			// service parameters
			relationshipId, relationship := args[1], args[2]
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])

			entity := types.NewLinkedEntity(
				relationshipId,
				relationship,
			)

			msg := types.NewMsgAddLinkedEntity(
				did.String(),
				entity,
				signer.String(),
			)
			// broadcast
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDeleteLinkedEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-linked-entity [id] [entity-id]",
		Short:   "deletes a entity from a decentralized did (did) document",
		Example: "delete a entity for a did document",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// document did
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// signer
			signer := clientCtx.GetFromAddress()
			// resource id
			eID := args[1]

			msg := types.NewMsgDeleteLinkedEntity(
				did.String(),
				eID,
				signer.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
