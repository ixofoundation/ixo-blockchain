package keeper
//
//import (
//	"github.com/cosmos/cosmos-sdk/codec"
//	"github.com/cosmos/cosmos-sdk/store"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	"github.com/cosmos/cosmos-sdk/x/params"
//	"github.com/cosmos/cosmos-sdk/x/staking"
//	//"github.com/cosmos/cosmos-sdk/x/supply"
//	"github.com/ixofoundation/ixo-blockchain/x/did"
//	"github.com/stretchr/testify/require"
//	abci "github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/libs/log"
//	tmtypes "github.com/tendermint/tendermint/types"
//	tmDB "github.com/tendermint/tm-db"
//	"testing"
//
//	"github.com/ixofoundation/ixo-blockchain/x/payments"
//	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
//)
//
//func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, Keeper, payments.Keeper, bank.Keeper) {
//	keyProject := sdk.NewKVStoreKey(types.StoreKey)
//	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
//	keyParams := sdk.NewKVStoreKey("subspace")
//	tkeyParams := sdk.NewTransientStoreKey("transient_params")
//	keyPayments := sdk.NewKVStoreKey(payments.StoreKey)
//	keyDid := sdk.NewKVStoreKey(did.StoreKey)
//
//	db := tmDB.NewMemDB()
//	ms := store.NewCommitMultiStore(db)
//	ms.MountStoreWithDB(keyProject, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(keyPayments, sdk.StoreTypeIAVL, nil)
//	ms.MountStoreWithDB(keyDid, sdk.StoreTypeIAVL, nil)
//	err := ms.LoadLatestVersion()
//	require.Nil(t, err)
//
//	ctx := sdk.NewContext(ms, abci.Header{ChainID: "foochainid"}, isCheckTx, log.NewNopLogger())
//	ctx = ctx.WithConsensusParams(
//		&abci.ConsensusParams{
//			Validator: &abci.ValidatorParams{
//				PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519},
//			},
//		},
//	)
//	cdc := MakeTestCodec()
//
//	feeCollectorAcc := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
//	notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
//	bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)
//
//	blacklistedAddrs := make(map[string]bool)
//	blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
//	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
//	blacklistedAddrs[bondPool.GetAddress().String()] = true
//
//	reservedIdPrefixes := make([]string, 0)
//
//	pk := params.NewKeeper(cdc, keyParams, tkeyParams)
//	accountKeeper := auth.NewAccountKeeper(cdc, keyAcc, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
//	bankKeeper := bank.NewBaseKeeper(accountKeeper, pk.Subspace(bank.DefaultParamspace), blacklistedAddrs)
//	didKeeper := did.NewKeeper(cdc, keyDid)
//	paymentsKeeper := payments.NewKeeper(cdc, keyPayments, bankKeeper, didKeeper, reservedIdPrefixes)
//
//	keeper := NewKeeper(
//		cdc, keyProject, pk.Subspace(types.DefaultParamspace), accountKeeper, didKeeper, paymentsKeeper,
//	)
//
//	return ctx, keeper, paymentsKeeper, bankKeeper
//}
//
//func MakeTestCodec() *codec.Codec {
//	var cdc = codec.New()
//	types.RegisterCodec(cdc)
//	sdk.RegisterCodec(cdc)
//	codec.RegisterCrypto(cdc)
//	auth.RegisterCodec(cdc)
//
//	return cdc
//}
