package ixo

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	ethAuth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	ethProject "github.com/ixofoundation/ixo-go-abi/abi/project"

	"github.com/ixofoundation/ixo-cosmos/x/params"
)

const ETH_URL = "ETH_URL"

var FUNDING_METHOD_HASH = strings.ToLower(GetKeccak("transfer(address,uint256)")[0:8])

type EthTransaction struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Hash             string `json:"hash"`
		Nonce            string `json:"nonce"`
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		TransactionIndex string `json:"transactionIndex"`
		From             string `json:"from"`
		To               string `json:"to"`
		Value            string `json:"value"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Input            string `json:"input"`
	}
}

func (tx *EthTransaction) UnmarshalJSON(msg []byte) error {
	return json.Unmarshal(msg, &tx.Result)
}

type EthClient struct {
	rpcClient *rpc.Client
	client    *ethclient.Client
	callOpts  bind.CallOpts
}

func NewEthClient() (EthClient, error) {
	// TODO: REMEMBER TO GET THE TARGET RPC ENDPOINT FROM THE ENVIRONMENT !!!
	// url := LookupEnv(ETH_URL, "https://api.infura.io/v1/jsonrpc/ropsten")
	// url := LookupEnv(ETH_URL, "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk")

	url := LookupEnv(ETH_URL, "http://localhost:8545")
	rpcClient, err := rpc.DialContext(context.Background(), url)

	if err != nil {
		return EthClient{}, err
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return EthClient{}, err
	}
	validatorWallet := getValidationEthWallet()
	callOpts := bind.CallOpts{
		Pending: false,
		From:    common.HexToAddress(validatorWallet.Address),
		Context: context.Background(),
	}

	return EthClient{
		rpcClient,
		client,
		callOpts,
	}, nil
}

func (c EthClient) GetTransactionByHash(txHash string) (*EthTransaction, error) {
	hash := common.HexToHash(txHash)
	var tx *EthTransaction
	err := c.rpcClient.CallContext(context.Background(), &tx, "eth_getTransactionByHash", hash)

	return tx, err
}

func (c EthClient) IsProjectFundingTx(ctx sdk.Context, projectDid Did, tx *EthTransaction) bool {

	ixoTokenContractAddress := "" // TODO (contracts): c.k.GetContract(ctx, contracts.KeyIxoTokenContractAddress)

	if strings.ToLower(tx.Result.To) != strings.ToLower(ixoTokenContractAddress) {
		ctx.Logger().Error("Token contract mismatch. Got " + tx.Result.To + " should be " + ixoTokenContractAddress)
		return false
	}

	if strings.ToLower(getMethodHashFromInput(tx.Result.Input)) != FUNDING_METHOD_HASH {
		ctx.Logger().Error("Method hash mismatch. Got " + getMethodHashFromInput(tx.Result.Input) + " should be " + FUNDING_METHOD_HASH)
		return false
	}

	_, err := c.ProjectWalletFromProjectRegistry(ctx, projectDid)
	if err != nil {
		ctx.Logger().Error("Could not get Project Wallet for DID " + projectDid)
		return false
	}

	return true
}

func getMethodHashFromInput(input string) string {
	return input[2:10]
}

func getParamFromInput(input string, paramPos int) string {

	start := 10 + 64*(paramPos-1)
	end := 10 + 64*paramPos
	param := input[start:end]

	return param
}

func (c EthClient) ProjectWalletFromProjectRegistry(ctx sdk.Context, did Did) (string, error) {

	regex := regexp.MustCompile("[^:]+$")

	var projectDid [32]byte
	copy(projectDid[:], regex.FindString(did))

	registryContractStr := "" // TODO (contracts): c.k.GetContract(ctx, contracts.KeyProjectRegistryContractAddress)
	registryContract := common.HexToAddress(registryContractStr)

	projectRegistryContact, err := ethProject.NewProjectWalletRegistry(registryContract, c.client)
	if err != nil {
		return "", err
	}

	projectWalletAddress, err := projectRegistryContact.WalletOf(&c.callOpts, projectDid)
	if err != nil {
		return "", err
	}

	return projectWalletAddress.String(), err
}

func (c EthClient) InitiateTokenTransfer(ctx sdk.Context, pk params.Keeper, senderAddr string, receiverAddr string, amount int64) (bool, [32]byte) {
	authContractAddress := common.HexToAddress("") // TODO (contracts): c.k.GetContract(ctx, contracts.KeyAuthContractAddress)
	authContract, err := ethAuth.NewAuthContract(authContractAddress, c.client)
	if err != nil {
		return false, [32]byte{}
	}

	validationEthWallet := getValidationEthWallet()
	privateKey, err := crypto.HexToECDSA(validationEthWallet.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	transOpts := bind.NewKeyedTransactor(privateKey)
	transOpts.GasLimit = uint64(2782100)

	projectWalletAuthoriserAddress := "" // TODO (contracts): c.k.GetContract(ctx, contracts.KeyProjectWalletAuthoriserContractAddress)

	nextTxID := getNextTxID(ctx, pk)
	var debugTxID []byte
	debugTxID = nextTxID[:]
	fmt.Println("-------------------\n\n", common.ToHex(debugTxID))

	txResult, err := authContract.Validate(transOpts, nextTxID, common.HexToAddress(projectWalletAuthoriserAddress),
		common.HexToAddress(senderAddr), common.HexToAddress(receiverAddr), big.NewInt(amount))
	fmt.Println("authContract.Validate: ", txResult, err)
	if err != nil {
		return false, nextTxID
	}

	return true, nextTxID
}

func (c EthClient) GetFundingAmt(tx *EthTransaction) int64 {
	return c.GetInt64FromHexString(tx.Result.Input[74:])
}

func (c EthClient) GetInt64FromHexString(hex string) int64 {
	amtHash := common.HexToHash(hex)
	return amtHash.Big().Int64()
}

func GetKeccak(tx string) string {
	hash := sha3.NewKeccak256()

	var buf []byte
	hash.Write([]byte(tx))
	buf = hash.Sum(buf)

	return hex.EncodeToString(buf)
}

func IxoAppGenEthWallet() (string, error) {
	ethWallet, err := CreateEthWallet()
	if err != nil {
		return "", err
	}

	json, err := json.Marshal(ethWallet)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(getEthWalletFilename(), json, 0644)

	return ethWallet.Address, nil
}

func getValidationEthWallet() EthWallet {
	data, err := ioutil.ReadFile(getEthWalletFilename())
	var ethWallet EthWallet
	err = json.Unmarshal(data, &ethWallet)
	if err != nil {
		panic(err)
	}
	return ethWallet
}

func getEthWalletFilename() string {
	rootDir := os.ExpandEnv("$HOME/.ixod/config")
	return rootDir + "/ethWallet.json"
}

func CreateEthWallet() (ethWallet *EthWallet, err error) {
	key, err := ethCrypto.GenerateKey()
	if err != nil {
		return
	}

	address := ethCrypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())

	ethWallet = &EthWallet{Address: address, PrivateKey: privateKey}
	return
}

func removeDidPrefix(did Did) string {
	const prefix = "did:ixo:"
	didStr := string(did)
	if strings.HasPrefix(didStr, prefix) {
		return didStr[8:]
	}
	return didStr
}

func getNextTxID(ctx sdk.Context, keeper params.Keeper) [32]byte {
	var nextTxID sdk.Dec
	actionID, err := keeper.Getter().GetDec(ctx, "actionID")
	if err == nil {
		nextTxID = actionID.Add(sdk.NewDec(1))
	} else {
		nextTxID = sdk.NewDec(1)
	}
	keeper.Setter().SetDec(ctx, "actionID", nextTxID)

	var result [32]byte
	copy(result[:], nextTxID.Bytes())

	return result
}
