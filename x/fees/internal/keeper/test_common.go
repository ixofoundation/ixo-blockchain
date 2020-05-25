package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

var (
	shareAddr1     = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr1")))
	shareAddr2     = sdk.AccAddress(crypto.AddressHash([]byte("shareAddr2")))
	feeCreatorAddr = sdk.AccAddress(crypto.AddressHash([]byte("feeCreatorAddr")))
	feePayerAddr   = sdk.AccAddress(crypto.AddressHash([]byte("feePayerAddr")))

	validChargeAmount = sdk.NewCoins(
		sdk.NewCoin(ixo.IxoNativeToken, sdk.OneInt()),
		sdk.NewCoin("rez", sdk.NewInt(2)))

	validChargeMinimum    = sdk.NewCoins(sdk.NewCoin("res", sdk.NewInt(3)))
	validChargeMaximum    = sdk.NewCoins()
	validCumulativeCharge = sdk.NewCoins(sdk.NewCoin("ixo", sdk.OneInt()))

	validDiscounts = types.NewDiscounts(
		types.NewDiscount(1, sdk.MustNewDecFromStr("0.5")),
		types.NewDiscount(2, sdk.MustNewDecFromStr("0.5")))

	validDistribution = types.NewDistribution(
		types.NewDistributionShare(shareAddr1, sdk.NewDec(50)),
		types.NewDistributionShare(shareAddr2, sdk.NewDec(50)))

	validFeeContent = types.NewFeeContent(
		validChargeAmount,
		validChargeMinimum,
		validChargeMaximum,
		validDiscounts,
		validDistribution)

	validFeeContractContent = types.NewFeeContractContent(
		0, feeCreatorAddr, feePayerAddr,
		validCumulativeCharge, false, true)
)

func CreateTestInput() (sdk.Context, Keeper, *codec.Codec) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)
	_ = ms.LoadLatestVersion()
	ctx := sdk.NewContext(ms, abci.Header{}, true, log.NewNopLogger())
	cdc := codec.New()
	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")
	pk1 := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	feesSubspace := pk1.Subspace(types.DefaultParamspace)
	keeper := NewKeeper(cdc, storeKey, feesSubspace)

	return ctx, keeper, cdc
}
