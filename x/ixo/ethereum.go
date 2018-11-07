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
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	contracts "github.com/ixofoundation/ixo-cosmos/x/contracts"
	ethAuth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	ethProject "github.com/ixofoundation/ixo-go-abi/abi/project"
)

const ETH_URL = "ETH_URL"

var FUNDING_METHOD_HASH = GetKeccak("transfer(address,uint256)")[0:8]

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
	k         contracts.Keeper
	callOpts  bind.CallOpts
}

func NewEthClient(k contracts.Keeper) (EthClient, error) {
	//url := LookupEnv(ETH_URL, "https://api.infura.io/v1/jsonrpc/ropsten")
	// url := LookupEnv(ETH_URL, "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk")

	url := LookupEnv(ETH_URL, "http://localhost:7545")
	fmt.Printf("PROJECT_FUNDING | url: %s \n", url)
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
		k,
		callOpts,
	}, nil
}

func (c EthClient) GetTransactionByHash(txHash string) (*EthTransaction, error) {
	hash := common.HexToHash(txHash)
	// // * We don't use the utility methods because of a problem in Ganache with hex values prefixed with zero
	// tx, _, err := c.ethClient.TransactionByHash(context.Background(), hash)
	// if err != nil {
	// 	fmt.Println(err)
	//	panic(err)
	// }
	// fmt.Println("tx 1")
	// return &types.Transaction{}, nil
	//	var res interface{}
	var tx *EthTransaction
	fmt.Printf("PROJECT_FUNDING | Just before call to \n")
	err := c.rpcClient.CallContext(context.Background(), &tx, "eth_getTransactionByHash", hash)
	fmt.Printf("PROJECT_FUNDING | err: %s \n", err)

	return tx, err
}

// func (c EthClient) DebugEthAddress(ctx sdk.Context) {
// 	ixoTokenContractAddress := c.k.GetEthAddress(ctx, KeyIxoTokenContractAddress)
// 	fmt.Printf("xxxxxxxxxxxxxxxxxxxxx | ixoTokenContractAddress: %s\n", ixoTokenContractAddress)
// }

// checks whether this is a funding transaction on this project
func (c EthClient) IsProjectFundingTx(ctx sdk.Context, projectDid Did, tx *EthTransaction) bool {
	fmt.Printf("PROJECT_FUNDING | func IsProjectFundingTx.\n")

	ixoTokenContractAddress := c.k.GetContract(ctx, contracts.KeyIxoTokenContractAddress)
	fmt.Printf("PROJECT_FUNDING | ercContractStr: %s\n", ixoTokenContractAddress)

	// Check To is the ERC20 Token
	fmt.Printf("PROJECT_FUNDING | tx.To: %s\n", tx.Result.To)
	if tx.Result.To != ixoTokenContractAddress {
		fmt.Printf("PROJECT_FUNDING | debug: fail 1\n")
		return false
	}

	// Check it is the transfer method
	fmt.Printf("PROJECT_FUNDING | FUNDING_METHOD_HASH: %s\n", FUNDING_METHOD_HASH)
	fmt.Printf("PROJECT_FUNDING | tx.Result.Input[2:10]: %s\n", tx.Result.Input[2:10])
	if tx.Result.Input[2:10] != FUNDING_METHOD_HASH {
		fmt.Printf("PROJECT_FUNDING | debug fail 2\n")
		return false
	}

	// Check the project wallet on the registry matches the wallet in the transaction
	txProjWallet := common.HexToAddress(tx.Result.Input[10:74]).String()
	fmt.Printf("PROJECT_FUNDING | txProjWallet: %s\n", txProjWallet)
	// Check it is the transfer method
	projWallet, err := c.GetEthProjectWallet(ctx, projectDid)
	fmt.Printf("PROJECT_FUNDING | projWallet: %s\n", projWallet)
	if err != nil {
		fmt.Printf("PROJECT_FUNDING | debug fail 3\n")
		return false
	}
	if txProjWallet != projWallet {
		fmt.Printf("PROJECT_FUNDING | debug fail 4\n")
		return false
	}

	return true
}

// Retrieves the Project wallet address from the Ethereum registry project conteact
func (c EthClient) GetEthProjectWallet(ctx sdk.Context, projectDid Did) (string, error) {

	registryContractStr := c.k.GetContract(ctx, contracts.KeyProjectRegistryContractAddress)
	registryContract := common.HexToAddress(registryContractStr)

	hexEncodedProjectDid := hex.EncodeToString([]byte(removeDidPrefix(projectDid)))
	var projectDidParam [32]byte
	copy(projectDidParam[:], []byte("0x"+hexEncodedProjectDid))

	projectRegistryContact, err := ethProject.NewProjectWalletRegistry(registryContract, c.client)
	if err != nil {
		return "", err
	}

	projectWalletAddress, err := projectRegistryContact.WalletOf(&c.callOpts, projectDidParam)
	return projectWalletAddress.String(), err
}

// InitiateTokenTransfer initiates the transfer of tokens from a source wallet to a destination wallet
func (c EthClient) InitiateTokenTransfer(ctx sdk.Context, senderAddr string, receiverAddr string, amount int64) bool {
	authContractAddress := common.HexToAddress(c.k.GetContract(ctx, contracts.KeyAuthContractAddress))
	authContract, err := ethAuth.NewAuthContract(authContractAddress, c.client)
	if err != nil {
		return false
	}

	validationEthWallet := getValidationEthWallet()
	privateKey, err := crypto.HexToECDSA(validationEthWallet.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	transOpts := bind.NewKeyedTransactor(privateKey)

	var txBytes [32]byte
	projectWalletAuthoriserAddress := c.k.GetContract(ctx, contracts.KeyProjectWalletAuthoriserContractAddress)

	authContract.Validate(transOpts, txBytes, common.HexToAddress(projectWalletAuthoriserAddress), common.HexToAddress(senderAddr), common.HexToAddress(receiverAddr), big.NewInt(amount))

	return true
}

// Gets the Funding amount out of the transcation data
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
	// Create an account
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
	rootDir := os.ExpandEnv("$HOME/.ixo-node/config")
	return rootDir + "/ethWallet.json"
}

func CreateEthWallet() (ethWallet *EthWallet, err error) {
	// Create an account
	key, err := ethCrypto.GenerateKey()
	if err != nil {
		return
	}

	// Get the address
	address := ethCrypto.PubkeyToAddress(key.PublicKey).Hex()
	// Get the private key
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
