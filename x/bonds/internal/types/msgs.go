package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"strings"
)

type MsgCreateBond struct { // signBytes should not be changed to sign_bytes because of ixo.types.DefaultTxDecoder
	SignBytes              string           `json:"signBytes" yaml:"signBytes"`
	BondDid                ixo.Did          `json:"bond_did" yaml:"bond_did"`
	PubKey                 string           `json:"pub_key" yaml:"pub_key"`
	Token                  string           `json:"token" yaml:"token"`
	Name                   string           `json:"name" yaml:"name"`
	Description            string           `json:"description" yaml:"description"`
	FunctionType           string           `json:"function_type" yaml:"function_type"`
	FunctionParameters     FunctionParams   `json:"function_parameters" yaml:"function_parameters"`
	Creator                sdk.AccAddress   `json:"creator" yaml:"creator"`
	ReserveTokens          []string         `json:"reserve_tokens" yaml:"reserve_tokens"`
	TxFeePercentage        sdk.Dec          `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
	ExitFeePercentage      sdk.Dec          `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
	FeeAddress             sdk.AccAddress   `json:"fee_address" yaml:"fee_address"`
	MaxSupply              sdk.Coin         `json:"max_supply" yaml:"max_supply"`
	OrderQuantityLimits    sdk.Coins        `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             sdk.Dec          `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage sdk.Dec          `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	AllowSells             string           `json:"allow_sells" yaml:"allow_sells"`
	Signers                []sdk.AccAddress `json:"signers" yaml:"signers"`
	BatchBlocks            sdk.Uint         `json:"batch_blocks" yaml:"batch_blocks"`
}

func NewMsgCreateBond(token, name, description string, creator sdk.AccAddress,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress sdk.AccAddress, maxSupply sdk.Coin,
	orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage sdk.Dec,
	allowSell string, signers []sdk.AccAddress, batchBlocks sdk.Uint,
	bondDid sovrin.SovrinDid) MsgCreateBond {
	return MsgCreateBond{
		SignBytes:              "",
		BondDid:                bondDid.Did,
		PubKey:                 bondDid.VerifyKey,
		Token:                  token,
		Name:                   name,
		Description:            description,
		Creator:                creator,
		FunctionType:           functionType,
		FunctionParameters:     functionParameters,
		ReserveTokens:          reserveTokens,
		TxFeePercentage:        txFeePercentage,
		ExitFeePercentage:      exitFeePercentage,
		FeeAddress:             feeAddress,
		MaxSupply:              maxSupply,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		AllowSells:             strings.ToLower(allowSell),
		Signers:                signers,
		BatchBlocks:            batchBlocks,
	}
}

func (msg MsgCreateBond) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.Token) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Token")
	} else if strings.TrimSpace(msg.Name) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Description")
	} else if msg.Creator.Empty() {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Creator")
	} else if len(msg.ReserveTokens) == 0 {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Reserve token")
	} else if msg.FeeAddress.Empty() {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Fee address")
	} else if strings.TrimSpace(msg.FunctionType) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Function type")
	} else if strings.TrimSpace(msg.AllowSells) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "AllowSells")
	}
	// Note: FunctionParameters can be empty

	// Check that true or false
	if msg.AllowSells != TRUE && msg.AllowSells != FALSE {
		return ErrArgumentMissingOrNonBoolean(DefaultCodespace, "AllowSells")
	}

	// Check that not negative
	if msg.TxFeePercentage.IsNegative() {
		return ErrArgumentCannotBeNegative(DefaultCodespace, "TxFeePercentage")
	} else if msg.ExitFeePercentage.IsNegative() {
		return ErrArgumentCannotBeNegative(DefaultCodespace, "ExitFeePercentage")
	}

	// Check that not zero
	if msg.BatchBlocks.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "BatchBlocks")
	} else if msg.MaxSupply.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "MaxSupply")
	} else {
		// TODO: consider allowing negative function parameters where possible
		for _, fp := range msg.FunctionParameters {
			if fp.Value.IsZero() {
				return ErrArgumentMustBePositive(DefaultCodespace, "FunctionParams:"+fp.Param)
			}
		}
	}

	// Note: uniqueness of reserve tokens checked when parsing

	return nil
}

func (msg MsgCreateBond) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.BondDid)}
}

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return ModuleName }

func (msg MsgCreateBond) IsNewDid() bool { return true }

type MsgEditBond struct { // signBytes should not be changed to sign_bytes because of ixo.types.DefaultTxDecoder
	SignBytes              string           `json:"signBytes" yaml:"signBytes"`
	BondDid                ixo.Did          `json:"bond_did" yaml:"bond_did"`
	Token                  string           `json:"token" yaml:"token"`
	Name                   string           `json:"name" yaml:"name"`
	Description            string           `json:"description" yaml:"description"`
	OrderQuantityLimits    string           `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             string           `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage string           `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	Editor                 sdk.AccAddress   `json:"editor" yaml:"editor"`
	Signers                []sdk.AccAddress `json:"signers" yaml:"signers"`
}

func NewMsgEditBond(token, name, description, orderQuantityLimits, sanityRate,
	sanityMarginPercentage string, editor sdk.AccAddress,
	signers []sdk.AccAddress, bondDid sovrin.SovrinDid) MsgEditBond {
	return MsgEditBond{
		SignBytes:              "",
		BondDid:                bondDid.Did,
		Token:                  token,
		Name:                   name,
		Description:            description,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		Editor:                 editor,
		Signers:                signers,
	}
}

func (msg MsgEditBond) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.Token) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Token")
	} else if strings.TrimSpace(msg.Name) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Description")
	} else if strings.TrimSpace(msg.SanityRate) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SanityRate")
	} else if strings.TrimSpace(msg.SanityMarginPercentage) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SanityMarginPercentage")
	}
	// Note: order quantity limits can be blank

	// Check that at least one editable was edited. Fields that will not
	// be edited should be "DoNotModifyField", and not an empty string
	inputList := []string{
		msg.Name, msg.Description, msg.OrderQuantityLimits,
		msg.SanityRate, msg.SanityMarginPercentage,
	}
	atLeaseOneEdit := false
	for _, e := range inputList {
		if e != DoNotModifyField {
			atLeaseOneEdit = true
			break
		}
	}
	if !atLeaseOneEdit {
		return ErrDidNotEditAnything(DefaultCodespace)
	}

	return nil
}

func (msg MsgEditBond) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg MsgEditBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.BondDid)}
}

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return ModuleName }

func (msg MsgEditBond) IsNewDid() bool { return false }

type MsgBuy struct { // signBytes should not be changed to sign_bytes because of ixo.types.DefaultTxDecoder
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	BuyerDid  ixo.Did   `json:"buyer_did" yaml:"buyer_did"`
	PubKey    string    `json:"pub_key" yaml:"pub_key"`
	Amount    sdk.Coin  `json:"amount" yaml:"amount"`
	MaxPrices sdk.Coins `json:"max_prices" yaml:"max_prices"`
	BondDid   ixo.Did   `json:"bond_did" yaml:"bond_did"`
}

func NewMsgBuy(buyerDid sovrin.SovrinDid, amount sdk.Coin, maxPrices sdk.Coins,
	bondDid ixo.Did) MsgBuy {
	return MsgBuy{
		SignBytes: "",
		BuyerDid:  buyerDid.Did,
		PubKey:    buyerDid.VerifyKey,
		Amount:    amount,
		MaxPrices: maxPrices,
		BondDid:   bondDid,
	}
}

func (msg MsgBuy) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BuyerDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BuyerDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that non zero
	if msg.Amount.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "Amount")
	}

	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.BuyerDid)}
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return ModuleName }

func (msg MsgBuy) IsNewDid() bool { return false }

type MsgSell struct { // signBytes should not be changed to sign_bytes because of ixo.types.DefaultTxDecoder
	SignBytes string   `json:"signBytes" yaml:"signBytes"`
	SellerDid ixo.Did  `json:"seller_did" yaml:"seller_did"`
	PubKey    string   `json:"pub_key" yaml:"pub_key"`
	Amount    sdk.Coin `json:"amount" yaml:"amount"`
	BondDid   ixo.Did  `json:"bond_did" yaml:"bond_did"`
}

func NewMsgSell(sellerDid sovrin.SovrinDid, amount sdk.Coin, bondDid ixo.Did) MsgSell {
	return MsgSell{
		SignBytes: "",
		SellerDid: sellerDid.Did,
		PubKey:    sellerDid.VerifyKey,
		Amount:    amount,
		BondDid:   bondDid,
	}
}

func (msg MsgSell) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SellerDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SellerDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that non zero
	if msg.Amount.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "Amount")
	}

	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.SellerDid)}
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return ModuleName }

func (msg MsgSell) IsNewDid() bool { return false }

type MsgSwap struct { // signBytes should not be changed to sign_bytes because of ixo.types.DefaultTxDecoder
	SignBytes  string   `json:"signBytes" yaml:"signBytes"`
	SwapperDid ixo.Did  `json:"swapper_did" yaml:"swapper_did"`
	PubKey     string   `json:"pub_key" yaml:"pub_key"`
	BondDid    ixo.Did  `json:"bond_did" yaml:"bond_did"`
	From       sdk.Coin `json:"from" yaml:"from"`
	ToToken    string   `json:"to_token" yaml:"to_token"`
}

func NewMsgSwap(swapperDid sovrin.SovrinDid, from sdk.Coin, toToken string,
	bondDid ixo.Did) MsgSwap {
	return MsgSwap{
		SignBytes:  "",
		SwapperDid: swapperDid.Did,
		PubKey:     swapperDid.VerifyKey,
		From:       from,
		ToToken:    toToken,
		BondDid:    bondDid,
	}
}

func (msg MsgSwap) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SwapperDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SwapperDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.ToToken) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "ToToken")
	}

	// Check if from and to the same token
	if msg.From.Denom == msg.ToToken {
		return ErrFromAndToCannotBeTheSameToken(DefaultCodespace)
	}

	// Check that non zero
	if msg.From.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "FromAmount")
	}

	// Note: From denom and amount must be valid since sdk.Coin
	return nil
}

func (msg MsgSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.SwapperDid)}
}

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return ModuleName }

func (msg MsgSwap) IsNewDid() bool { return false }
