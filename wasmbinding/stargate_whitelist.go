package wasmbinding

import (
	"fmt"
	"sort"
	"sync"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v5/x/bonds/types"
	claimstypes "github.com/ixofoundation/ixo-blockchain/v5/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v5/x/entity/types"
	epochtypes "github.com/ixofoundation/ixo-blockchain/v5/x/epochs/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v5/x/iid/types"
	liquidstaketypes "github.com/ixofoundation/ixo-blockchain/v5/x/liquidstake/types"
	minttypes "github.com/ixofoundation/ixo-blockchain/v5/x/mint/types"
	smartaccounttypes "github.com/ixofoundation/ixo-blockchain/v5/x/smart-account/types"
	tokentypes "github.com/ixofoundation/ixo-blockchain/v5/x/token/types"
)

// stargateResponsePools keeps whitelist and its deterministic
// response binding for stargate queries.
// CONTRACT: since results of queries go into blocks, queries being added here should always be
// deterministic or can cause non-determinism in the state machine.
//
// The query is multi-threaded so we're using a sync.Pool
// to manage the allocation and de-allocation of newly created
// pb objects.
var stargateResponsePools = make(map[string]*sync.Pool)

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
	setWhitelistedQuery("/cosmos.auth.v1beta1.Query/ModuleAccounts", &authtypes.QueryModuleAccountsResponse{})

	// bank
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/Balance", &banktypes.QueryBalanceResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/AllBalances", &banktypes.QueryAllBalancesResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/DenomMetadata", &banktypes.QueryDenomMetadataResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/DenomsMetadata", &banktypes.QueryDenomsMetadataResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/Params", &banktypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/SupplyOf", &banktypes.QuerySupplyOfResponse{})

	// distribution
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/Params", &distributiontypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress", &distributiontypes.QueryDelegatorWithdrawAddressResponse{})
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/ValidatorCommission", &distributiontypes.QueryValidatorCommissionResponse{})

	// gov
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Deposit", &govtypesv1.QueryDepositResponse{})
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Params", &govtypesv1.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Vote", &govtypesv1.QueryVoteResponse{})

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
	setWhitelistedQuery("/ixo.claims.v1beta1.Query/Intent", &claimstypes.QueryIntentResponse{})
	setWhitelistedQuery("/ixo.claims.v1beta1.Query/IntentList", &claimstypes.QueryIntentListResponse{})

	// entity
	setWhitelistedQuery("/ixo.entity.v1beta1.Query/Params", &entitytypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.entity.v1beta1.Query/Entity", &entitytypes.QueryEntityResponse{})
	setWhitelistedQuery("/ixo.entity.v1beta1.Query/EntityVerified", &entitytypes.QueryEntityVerifiedResponse{})

	// iid
	setWhitelistedQuery("/ixo.iid.v1beta1.Query/IidDocument", &iidtypes.QueryIidDocumentResponse{})

	// token
	setWhitelistedQuery("/ixo.token.v1beta1.Query/Params", &tokentypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.token.v1beta1.Query/TokenMetadata", &tokentypes.QueryTokenMetadataResponse{})
	setWhitelistedQuery("/ixo.token.v1beta1.Query/TokenDoc", &tokentypes.QueryTokenDocResponse{})

	// smartaccount
	setWhitelistedQuery("/ixo.smartaccount.v1beta1.Query/Params", &smartaccounttypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.smartaccount.v1beta1.Query/GetAuthenticator", &smartaccounttypes.GetAuthenticatorResponse{})
	setWhitelistedQuery("/ixo.smartaccount.v1beta1.Query/GetAuthenticators", &smartaccounttypes.GetAuthenticatorsResponse{})

	// epochs
	setWhitelistedQuery("/ixo.epochs.v1beta1.Query/EpochInfos", &epochtypes.QueryEpochsInfoResponse{})
	setWhitelistedQuery("/ixo.epochs.v1beta1.Query/CurrentEpoch", &epochtypes.QueryCurrentEpochResponse{})

	// mint
	setWhitelistedQuery("/ixo.mint.v1beta1.Query/Params", &minttypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.mint.v1beta1.Query/EpochProvisions", &minttypes.QueryEpochProvisionsResponse{})

	// liquidstake
	setWhitelistedQuery("/ixo.liquidstake.v1beta1.Query/Params", &liquidstaketypes.QueryParamsResponse{})
	setWhitelistedQuery("/ixo.liquidstake.v1beta1.Query/LiquidValidators", &liquidstaketypes.QueryLiquidValidatorsResponse{})
	setWhitelistedQuery("/ixo.liquidstake.v1beta1.Query/States", &liquidstaketypes.QueryStatesResponse{})
}

// IsWhitelistedQuery returns if the query is not whitelisted.
func IsWhitelistedQuery(queryPath string) error {
	_, isWhitelisted := stargateResponsePools[queryPath]
	if !isWhitelisted {
		return wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("'%s' path is not allowed from the contract", queryPath)}
	}
	return nil
}

// getWhitelistedQuery returns the whitelisted query at the provided path.
// If the query does not exist, or it was setup wrong by the chain, this returns an error.
// CONTRACT: must call returnStargateResponseToPool in order to avoid pointless allocs.
func getWhitelistedQuery(queryPath string) (proto.Message, error) {
	protoResponseAny, isWhitelisted := stargateResponsePools[queryPath]
	if !isWhitelisted {
		return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("'%s' path is not allowed from the contract", queryPath)}
	}
	protoMarshaler, ok := protoResponseAny.Get().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to assert type to codec.ProtoMarshaler")
	}
	return protoMarshaler, nil
}

type protoTypeG[T any] interface {
	*T
	proto.Message
}

// setWhitelistedQuery sets the whitelisted query at the provided path.
// This method also creates a sync.Pool for the provided protoMarshaler.
// We use generics so we can properly instantiate an object that the
// queryPath expects as a response.
func setWhitelistedQuery[T any, PT protoTypeG[T]](queryPath string, _ PT) {
	stargateResponsePools[queryPath] = &sync.Pool{
		New: func() any {
			return PT(new(T))
		},
	}
}

// returnStargateResponseToPool returns the provided protoMarshaler to the appropriate pool based on it's query path.
func returnStargateResponseToPool(queryPath string, pb proto.Message) {
	stargateResponsePools[queryPath].Put(pb)
}

func GetStargateWhitelistedPaths() (keys []string) {
	// Iterate over the map and collect the keys
	keys = make([]string, 0, len(stargateResponsePools))
	for k := range stargateResponsePools {
		keys = append(keys, k)
	}

	// Sort the keys to ensure determinism
	sort.Strings(keys)
	return keys
}
