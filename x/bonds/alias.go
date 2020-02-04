package bonds

import (
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
)

//noinspection GoUnusedConst
const (
	QueryBonds          = keeper.QueryBonds
	QueryBond           = keeper.QueryBond
	QueryCurrentPrice   = keeper.QueryCurrentPrice
	QueryCurrentReserve = keeper.QueryCurrentReserve
	QueryCustomPrice    = keeper.QueryCustomPrice
	QueryBuyPrice       = keeper.QueryBuyPrice
	QuerySellReturn     = keeper.QuerySellReturn

	DefaultCodeSpace = types.DefaultCodespace

	CodeArgumentInvalid                      = types.CodeArgumentInvalid
	CodeArgumentMissingOrIncorrectType       = types.CodeArgumentMissingOrIncorrectType
	CodeIncorrectNumberOfValues              = types.CodeIncorrectNumberOfValues
	CodeBondDoesNotExist                     = types.CodeBondDoesNotExist
	CodeBondAlreadyExists                    = types.CodeBondAlreadyExists
	CodeBondDoesNotAllowSelling              = types.CodeBondDoesNotAllowSelling
	CodeDidNotEditAnything                   = types.CodeDidNotEditAnything
	CodeReserveAddrCannotBeFeeAddr           = types.CodeInvalidBond
	CodeUnrecognizedFunctionType             = types.CodeUnrecognizedFunctionType
	CodeInvalidFunctionParameter             = types.CodeInvalidFunctionParameter
	CodeFunctionNotAvailableForFunctionType  = types.CodeFunctionNotAvailableForFunctionType
	CodeFunctionRequiresNonZeroCurrentSupply = types.CodeFunctionRequiresNonZeroCurrentSupply
	CodeReserveTokenInvalid                  = types.CodeReserveTokenInvalid
	CodeBondTokenInvalid                     = types.CodeBondTokenInvalid
	CodeFromAndToCannotBeTheSameToken        = types.CodeInvalidSwapper
	CodeReserveDenomsMismatch                = types.CodeReserveDenomsMismatch
	CodeDuplicateReserveToken                = types.CodeInvalidBond
	CodeInvalidCoinDenomination              = types.CodeInvalidCoinDenomination
	CodeMaxSupplyExceeded                    = types.CodeInvalidResultantSupply
	CodeMinSupplyExceeded                    = types.CodeInvalidResultantSupply
	CodeMaxPriceExceeded                     = types.CodeMaxPriceExceeded
	CodeSwapAmountInvalid                    = types.CodeSwapAmountInvalid
	CodeOrderQuantityLimitExceeded           = types.CodeOrderLimitExceeded
	CodeSanityRateViolated                   = types.CodeSanityRateViolated
	CodeFeeTooLarge                          = types.CodeFeeTooLarge

	BondsMintBurnAccount       = types.BondsMintBurnAccount
	BatchesIntermediaryAccount = types.BatchesIntermediaryAccount

	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
)

//noinspection GoUnusedGlobalVariable,GoNameStartsWithPackageName
var (
	// function aliases
	RegisterInvariants = keeper.RegisterInvariants
	AllInvariants      = keeper.AllInvariants
	SupplyInvariant    = keeper.SupplyInvariant
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	RegisterCodec      = types.RegisterCodec

	ErrArgumentCannotBeEmpty                = types.ErrArgumentCannotBeEmpty
	ErrArgumentCannotBeNegative             = types.ErrArgumentCannotBeNegative
	ErrFunctionParameterMissingOrNonInteger = types.ErrFunctionParameterMissingOrNonInteger
	ErrArgumentMissingOrNonFloat            = types.ErrArgumentMissingOrNonFloat
	ErrArgumentMissingOrNonInteger          = types.ErrArgumentMissingOrNonInteger
	ErrArgumentMissingOrNonUInteger         = types.ErrArgumentMissingOrNonUInteger
	ErrArgumentMissingOrNonBoolean          = types.ErrArgumentMissingOrNonBoolean
	ErrIncorrectNumberOfReserveTokens       = types.ErrIncorrectNumberOfReserveTokens
	ErrIncorrectNumberOfFunctionParameters  = types.ErrIncorrectNumberOfFunctionParameters
	ErrBondDoesNotExist                     = types.ErrBondDoesNotExist
	ErrBondAlreadyExists                    = types.ErrBondAlreadyExists
	ErrBondDoesNotAllowSelling              = types.ErrBondDoesNotAllowSelling
	ErrDidNotEditAnything                   = types.ErrDidNotEditAnything
	ErrReserveAddrCannotBeFeeAddr           = types.ErrReserveAddrCannotBeFeeAddr
	ErrUnrecognizedFunctionType             = types.ErrUnrecognizedFunctionType
	ErrInvalidFunctionParameter             = types.ErrInvalidFunctionParameter
	ErrFunctionNotAvailableForFunctionType  = types.ErrFunctionNotAvailableForFunctionType
	ErrFunctionRequiresNonZeroCurrentSupply = types.ErrFunctionRequiresNonZeroCurrentSupply
	ErrTokenIsNotAValidReserveToken         = types.ErrTokenIsNotAValidReserveToken
	ErrBondTokenCannotAlsoBeReserveToken    = types.ErrBondTokenCannotAlsoBeReserveToken
	ErrBondTokenCannotBeStakingToken        = types.ErrBondTokenCannotBeStakingToken
	ErrFromAndToCannotBeTheSameToken        = types.ErrFromAndToCannotBeTheSameToken
	ErrReserveDenomsMismatch                = types.ErrReserveDenomsMismatch
	ErrDuplicateReserveToken                = types.ErrDuplicateReserveToken
	ErrInvalidCoinDenomination              = types.ErrInvalidCoinDenomination
	ErrCannotMintMoreThanMaxSupply          = types.ErrCannotMintMoreThanMaxSupply
	ErrCannotBurnMoreThanSupply             = types.ErrCannotBurnMoreThanSupply
	ErrMaxPriceExceeded                     = types.ErrMaxPriceExceeded
	ErrSwapAmountTooSmallToGiveAnyReturn    = types.ErrSwapAmountTooSmallToGiveAnyReturn
	ErrSwapAmountCausesReserveDepletion     = types.ErrSwapAmountCausesReserveDepletion
	ErrOrderQuantityLimitExceeded           = types.ErrOrderQuantityLimitExceeded
	ErrValuesViolateSanityRate              = types.ErrValuesViolateSanityRate
	ErrFeesCannotBeOrExceed100Percent       = types.ErrFeesCannotBeOrExceed100Percent

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	SquareRootDec       = types.SquareRootDec
	SquareRootInt       = types.SquareRootInt
	RoundReservePrice   = types.RoundReservePrice
	RoundReserveReturn  = types.RoundReserveReturn
	RoundFee            = types.RoundFee
	RoundReservePrices  = types.RoundReservePrices
	RoundReserveReturns = types.RoundReserveReturns

	NewFunctionParam = types.NewFunctionParam
	NewBond          = types.NewBond
	NewBatch         = types.NewBatch
	NewBaseOrder     = types.NewBaseOrder
	NewBuyOrder      = types.NewBuyOrder
	NewSellOrder     = types.NewSellOrder
	NewSwapOrder     = types.NewSwapOrder
	NewMsgCreateBond = types.NewMsgCreateBond
	NewMsgEditBond   = types.NewMsgEditBond
	NewMsgBuy        = types.NewMsgBuy
	NewMsgSell       = types.NewMsgSell
	NewMsgSwap       = types.NewMsgSwap

	// variable aliases
	ModuleCdc            = types.ModuleCdc
	BondsKeyPrefix       = types.BondsKeyPrefix
	BatchesKeyPrefix     = types.BatchesKeyPrefix
	LastBatchesKeyPrefix = types.LastBatchesKeyPrefix
)

type (
	Keeper       = keeper.Keeper
	CodeType     = types.CodeType
	GenesisState = types.GenesisState

	MsgCreateBond = types.MsgCreateBond
	MsgEditBond   = types.MsgEditBond
	MsgBuy        = types.MsgBuy
	MsgSell       = types.MsgSell
	MsgSwap       = types.MsgSwap

	FunctionParam  = types.FunctionParam
	FunctionParams = types.FunctionParams
	Bond           = types.Bond
	Batch          = types.Batch
	Order          = types.BaseOrder
	BuyOrder       = types.BuyOrder
	SellOrder      = types.SellOrder
	SwapOrder      = types.SwapOrder

	QueryResBonds      = types.QueryBonds
	QueryResBuyPrice   = types.QueryBuyPrice
	QueryResSellReturn = types.QuerySellReturn
	QueryResSwapReturn = types.QuerySwapReturn
)
