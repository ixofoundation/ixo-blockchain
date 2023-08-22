package wasmbinding

import (
	"fmt"
	"sync"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"

	bondstypes "github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	claimstypes "github.com/ixofoundation/ixo-blockchain/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	tokentypes "github.com/ixofoundation/ixo-blockchain/x/token/types"
)

// stargateWhitelist keeps whitelist and its deterministic
// response binding for stargate queries.
//
// The query can be multi-thread, so we have to use
// thread safe sync.Map.
var stargateWhitelist sync.Map

// Note: When adding a migration here, we should also add it to the Async ICQ params in the upgrade.
// In the future we may want to find a better way to keep these in sync

func init() {
	// ibc queries
	setWhitelistedQuery("/ibc.applications.transfer.v1.Query/DenomTrace", &ibctransfertypes.QueryDenomTraceResponse{})
	setWhitelistedQuery("/ibc.applications.transfer.v1.Query/DenomTraces", &ibctransfertypes.QueryDenomTracesResponse{})
	setWhitelistedQuery("/ibc.applications.transfer.v1.Query/DenomHash", &ibctransfertypes.QueryDenomHashResponse{})

	// cosmos-sdk queries
	// =============================

	// auth
	setWhitelistedQuery("/cosmos.auth.v1beta1.Query/Account", &authtypes.QueryAccountResponse{})
	setWhitelistedQuery("/cosmos.auth.v1beta1.Query/Params", &authtypes.QueryParamsResponse{})

	// bank
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/Balance", &banktypes.QueryBalanceResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/AllBalances", &banktypes.QueryAllBalancesResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/DenomMetadata", &banktypes.QueryDenomsMetadataResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/Params", &banktypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/SupplyOf", &banktypes.QuerySupplyOfResponse{})

	// distribution
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/Params", &distributiontypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress", &distributiontypes.QueryDelegatorWithdrawAddressResponse{})
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/ValidatorCommission", &distributiontypes.QueryValidatorCommissionResponse{})

	// gov
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Deposit", &govtypes.QueryDepositResponse{})
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Params", &govtypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Vote", &govtypes.QueryVoteResponse{})

	// slashing
	setWhitelistedQuery("/cosmos.slashing.v1beta1.Query/Params", &slashingtypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.slashing.v1beta1.Query/SigningInfo", &slashingtypes.QuerySigningInfoResponse{})

	// staking
	setWhitelistedQuery("/cosmos.staking.v1beta1.Query/Delegation", &stakingtypes.QueryDelegationResponse{})
	setWhitelistedQuery("/cosmos.staking.v1beta1.Query/Params", &stakingtypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.staking.v1beta1.Query/Validator", &stakingtypes.QueryValidatorResponse{})

	// ixo queries
	// =============================

	// bonds
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/Params", &bondstypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/Bond", &bondstypes.QueryBondResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/CurrentPrice", &bondstypes.QueryCurrentPriceResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/Batch", &bondstypes.QueryBatchResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/LastBatch", &bondstypes.QueryLastBatchResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/CurrentReserve", &bondstypes.QueryCurrentReserveResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/AvailableReserve", &bondstypes.QueryAvailableReserveResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/CustomPrice", &bondstypes.QueryCustomPriceResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/BuyPrice", &bondstypes.QueryBuyPriceResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/SellReturn", &bondstypes.QuerySellReturnResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/SwapReturn", &bondstypes.QuerySwapReturnResponse{})
	setWhitelistedQuery("/ixo.bonds.v1beta1.Query/AlphaMaximums", &bondstypes.QueryAlphaMaximumsResponse{})

	// claims
	setWhitelistedQuery("/ixo.claims.v1beta1.Query/Params", &claimstypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.claims.v1beta1.Query/Collection", &claimstypes.QueryCollectionResponse{})
	setWhitelistedQuery("/ixo.claims.v1beta1.Query/Claim", &claimstypes.QueryClaimResponse{})
	setWhitelistedQuery("/ixo.claims.v1beta1.Query/Dispute", &claimstypes.QueryDisputeResponse{})

	// entity
	setWhitelistedQuery("/ixo.entity.v1beta1.Query/Params", &entitytypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.entity.v1beta1.Query/Entity", &entitytypes.QueryEntityResponse{})
	setWhitelistedQuery("/ixo.entity.v1beta1.Query/EntityVerified", &entitytypes.QueryEntityVerifiedResponse{})

	// iid
	setWhitelistedQuery("/ixo.iid.v1beta1.Query/IidDocument", &iidtypes.QueryIidDocumentResponse{})

	// token
	setWhitelistedQuery("/ixo.token.v1beta1.Query/Params", &tokentypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.token.v1beta1.Query/TokenMetadata", &tokentypes.QueryTokenMetadataResponse{})
}

// GetWhitelistedQuery returns the whitelisted query at the provided path.
// If the query does not exist, or it was setup wrong by the chain, this returns an error.
func GetWhitelistedQuery(queryPath string) (codec.ProtoMarshaler, error) {
	protoResponseAny, isWhitelisted := stargateWhitelist.Load(queryPath)
	if !isWhitelisted {
		return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("'%s' path is not allowed from the contract", queryPath)}
	}
	protoResponseType, ok := protoResponseAny.(codec.ProtoMarshaler)
	if !ok {
		return nil, wasmvmtypes.Unknown{}
	}
	return protoResponseType, nil
}

func setWhitelistedQuery(queryPath string, protoType codec.ProtoMarshaler) {
	stargateWhitelist.Store(queryPath, protoType)
}

func GetStargateWhitelistedPaths() (keys []string) {
	// Iterate over the map and collect the keys
	stargateWhitelist.Range(func(key, value interface{}) bool {
		keyStr, ok := key.(string)
		if !ok {
			panic("key is not a string")
		}
		keys = append(keys, keyStr)
		return true
	})

	return keys
}
