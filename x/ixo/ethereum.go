package ixo

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ETH_URL = "ETH_URL"
const ETH_REGISTRY_CONTRACT = "ETH_REGISTRY_CONTRACT"
const ETH_IXO_ERC20_TOKEN = "ETH_IXO_ERC20_TOKEN"

type EthClient struct {
	ethClient        *ethclient.Client
	registryContract string
	ercContract      string
}

func NewEthClient() (EthClient, error) {
	//url := LookupEnv(ETH_URL, "https://api.infura.io/v1/jsonrpc/ropsten")
	// url := LookupEnv(ETH_URL, "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk")
	url := LookupEnv(ETH_URL, "http://127.0.0.1:7545")
	ethClient, err := ethclient.Dial(url)

	if err != nil {
		return EthClient{}, err
	}
	registryContract := LookupEnv(ETH_REGISTRY_CONTRACT, "")
	// if len(registryContract) == 0 {
	// 	return EthClient{}, errors.New("Ethereum Registry contract not set on env. ETH_REGISTRY_CONTRACT=")
	// }
	erc20Token := LookupEnv(ETH_IXO_ERC20_TOKEN, "")
	// if len(erc20Token) == 0 {
	// 	return EthClient{}, errors.New("Ethereum IXO ERC20 token contract not set on env. ETH_IXO_ERC20_TOKEN=")
	// }

	return EthClient{
		ethClient,
		registryContract,
		erc20Token,
	}, nil
}

func (c EthClient) GetTransactionByHash(txHash string) (*types.Transaction, error) {
	fmt.Println("tx -2")
	hash := common.HexToHash(txHash)
	fmt.Println("tx -1")
	tx, _, err := c.ethClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tx 1")
	fmt.Println(tx)
	fmt.Println(hex.EncodeToString(tx.Data()))
	fmt.Println("tx 2")
	return tx, err
}

func (c EthClient) IsProjectFundingTx(project Did, input []byte) bool {
	return false
}

func (c EthClient) GetFundingAmt(input []byte) int64 {
	return 0
}
