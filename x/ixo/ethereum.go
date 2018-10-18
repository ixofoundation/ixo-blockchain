package ixo

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rpc"
)

const ETH_URL = "ETH_URL"
const ETH_REGISTRY_CONTRACT = "ETH_REGISTRY_CONTRACT"
const ETH_IXO_ERC20_TOKEN = "ETH_IXO_ERC20_TOKEN"

var FUNDING_METHOD_HASH = GetKeccak("transfer(address,uint256)")[0:8]
var REGISTRY_WALLET_OF = GetKeccak("walletOf(bytes32)")[0:8]

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
	client           *rpc.Client
	registryContract common.Address
	ercContract      common.Address
}

func NewEthClient() (EthClient, error) {
	//url := LookupEnv(ETH_URL, "https://api.infura.io/v1/jsonrpc/ropsten")
	// url := LookupEnv(ETH_URL, "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk")
	url := LookupEnv(ETH_URL, "http://localhost:7545")
	client, err := rpc.DialContext(context.Background(), url)

	if err != nil {
		return EthClient{}, err
	}
	registryContractStr := LookupEnv(ETH_REGISTRY_CONTRACT, "0x1d7d616c01c63c662e676e63008a30845cbcfe1e")
	if len(registryContractStr) == 0 {
		return EthClient{}, errors.New("Ethereum Registry contract not set on env. ETH_REGISTRY_CONTRACT=")
	}
	registryContract := common.HexToAddress(registryContractStr)
	erc20TokenStr := LookupEnv(ETH_IXO_ERC20_TOKEN, "0x026aFf3ab0DaE74d5d85537f78B4dDEcC101C6D6")
	if len(erc20TokenStr) == 0 {
		return EthClient{}, errors.New("Ethereum IXO ERC20 token contract not set on env. ETH_IXO_ERC20_TOKEN=")
	}
	erc20Token := common.HexToAddress(erc20TokenStr)
	return EthClient{
		client,
		registryContract,
		erc20Token,
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
	err := c.client.CallContext(context.Background(), &tx, "eth_getTransactionByHash", hash)

	return tx, err
}

func (c EthClient) IsProjectFundingTx(project Did, tx *EthTransaction) bool {

	fmt.Println(FUNDING_METHOD_HASH)
	fmt.Println("To:", common.HexToAddress(tx.Result.To).String())
	txMethod := tx.Result.Input[2:10]
	fmt.Println("Method Hash", txMethod)
	txProjWallet := common.HexToAddress(tx.Result.Input[10:74]).String()
	fmt.Println("Proj Wallet:", txProjWallet)
	amt := c.GetFundingAmt(tx)
	fmt.Println("Amount:", amt)

	return false
}

func (c EthClient) GetEthProjectWallet(project Did) (string, error) {
	var hex hexutil.Bytes
	fmt.Println("Reg_Wallet_of:", REGISTRY_WALLET_OF)
	dataBytes := make([]byte, 0)
	dataBytes = append(dataBytes, []byte("0x")...)
	dataBytes = append(dataBytes, []byte(REGISTRY_WALLET_OF)...)
	dataBytes = append(dataBytes, []byte("00000000000000000000000078a706edcd907a5e897340724b3b530a5c8dcd9a")...)
	fmt.Println(string(dataBytes))
	//	data := common.FromHex("0xfca3b5aa00000000000000000000000078a706edcd907a5e897340724b3b530a5c8dcd9a")
	arg := map[string]interface{}{
		"from": c.registryContract,
		"to":   c.registryContract,
		"data": dataBytes,
	}
	err := c.client.CallContext(context.Background(), &hex, "eth_call", arg, "latest")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(hex))
	return string(hex), nil
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
