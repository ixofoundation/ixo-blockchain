// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package project

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BasicProjectWalletABI is the input ABI used to generate the binding from.
const BasicProjectWalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_authoriser\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BasicProjectWalletBin is the compiled bytecode used for deploying new contracts.
const BasicProjectWalletBin = `0x608060405234801561001057600080fd5b5060405160608061029983398101604090815281516020830151919092015160008054600160a060020a03948516600160a060020a03199182161790915560018054949093169316929092179055600255610229806100706000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610050578063a9059cbb14610077575b600080fd5b34801561005c57600080fd5b506100656100bc565b60408051918252519081900360200190f35b34801561008357600080fd5b506100a873ffffffffffffffffffffffffffffffffffffffff600435166024356100c2565b604080519115158252519081900360200190f35b60025481565b60015460009073ffffffffffffffffffffffffffffffffffffffff16331461014b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b505050506040513d60208110156101f457600080fd5b509093925050505600a165627a7a72305820a57f44df56894921b295f5260067fe533d856755c79ff5af9d7e46f1324c122b0029`

// DeployBasicProjectWallet deploys a new Ethereum contract, binding an instance of BasicProjectWallet to it.
func DeployBasicProjectWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _authoriser common.Address, _name [32]byte) (common.Address, *types.Transaction, *BasicProjectWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicProjectWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicProjectWalletBin), backend, _token, _authoriser, _name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicProjectWallet{BasicProjectWalletCaller: BasicProjectWalletCaller{contract: contract}, BasicProjectWalletTransactor: BasicProjectWalletTransactor{contract: contract}, BasicProjectWalletFilterer: BasicProjectWalletFilterer{contract: contract}}, nil
}

// BasicProjectWallet is an auto generated Go binding around an Ethereum contract.
type BasicProjectWallet struct {
	BasicProjectWalletCaller     // Read-only binding to the contract
	BasicProjectWalletTransactor // Write-only binding to the contract
	BasicProjectWalletFilterer   // Log filterer for contract events
}

// BasicProjectWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicProjectWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicProjectWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicProjectWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicProjectWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicProjectWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicProjectWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicProjectWalletSession struct {
	Contract     *BasicProjectWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasicProjectWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicProjectWalletCallerSession struct {
	Contract *BasicProjectWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BasicProjectWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicProjectWalletTransactorSession struct {
	Contract     *BasicProjectWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BasicProjectWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicProjectWalletRaw struct {
	Contract *BasicProjectWallet // Generic contract binding to access the raw methods on
}

// BasicProjectWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicProjectWalletCallerRaw struct {
	Contract *BasicProjectWalletCaller // Generic read-only contract binding to access the raw methods on
}

// BasicProjectWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicProjectWalletTransactorRaw struct {
	Contract *BasicProjectWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicProjectWallet creates a new instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWallet(address common.Address, backend bind.ContractBackend) (*BasicProjectWallet, error) {
	contract, err := bindBasicProjectWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWallet{BasicProjectWalletCaller: BasicProjectWalletCaller{contract: contract}, BasicProjectWalletTransactor: BasicProjectWalletTransactor{contract: contract}, BasicProjectWalletFilterer: BasicProjectWalletFilterer{contract: contract}}, nil
}

// NewBasicProjectWalletCaller creates a new read-only instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWalletCaller(address common.Address, caller bind.ContractCaller) (*BasicProjectWalletCaller, error) {
	contract, err := bindBasicProjectWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWalletCaller{contract: contract}, nil
}

// NewBasicProjectWalletTransactor creates a new write-only instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicProjectWalletTransactor, error) {
	contract, err := bindBasicProjectWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWalletTransactor{contract: contract}, nil
}

// NewBasicProjectWalletFilterer creates a new log filterer instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicProjectWalletFilterer, error) {
	contract, err := bindBasicProjectWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWalletFilterer{contract: contract}, nil
}

// bindBasicProjectWallet binds a generic wrapper to an already deployed contract.
func bindBasicProjectWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicProjectWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicProjectWallet *BasicProjectWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicProjectWallet.Contract.BasicProjectWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicProjectWallet *BasicProjectWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.BasicProjectWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicProjectWallet *BasicProjectWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.BasicProjectWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicProjectWallet *BasicProjectWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicProjectWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicProjectWallet *BasicProjectWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicProjectWallet *BasicProjectWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_BasicProjectWallet *BasicProjectWalletCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasicProjectWallet.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_BasicProjectWallet *BasicProjectWalletSession) Name() ([32]byte, error) {
	return _BasicProjectWallet.Contract.Name(&_BasicProjectWallet.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_BasicProjectWallet *BasicProjectWalletCallerSession) Name() ([32]byte, error) {
	return _BasicProjectWallet.Contract.Name(&_BasicProjectWallet.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_BasicProjectWallet *BasicProjectWalletTransactor) Transfer(opts *bind.TransactOpts, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _BasicProjectWallet.contract.Transact(opts, "transfer", _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_BasicProjectWallet *BasicProjectWalletSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.Transfer(&_BasicProjectWallet.TransactOpts, _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_BasicProjectWallet *BasicProjectWalletTransactorSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.Transfer(&_BasicProjectWallet.TransactOpts, _receiver, _amt)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IxoERC20TokenABI is the input ABI used to generate the binding from.
const IxoERC20TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newMinter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IxoERC20TokenBin is the compiled bytecode used for deploying new contracts.
const IxoERC20TokenBin = `0x60c0604052600960808190527f49584f20546f6b656e000000000000000000000000000000000000000000000060a090815261003e91600491906100ac565b506040805180820190915260038082527f49584f00000000000000000000000000000000000000000000000000000000006020909201918252610083916005916100ac565b506008600655670de0b6b3a764000060075560008054600160a060020a03191633179055610147565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ed57805160ff191683800117855561011a565b8280016001018555821561011a579182015b8281111561011a5782518255916020019190600101906100ff565b5061012692915061012a565b5090565b61014491905b808211156101265760008155600101610130565b90565b610fe280620001576000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610100578063075461721461018a578063095ea7b3146101bb57806318160ddd146101f357806323b872dd1461021a578063313ce5671461024457806340c10f1914610259578063661884631461027d57806370a08231146102a1578063715018a6146102c25780638da5cb5b146102d957806395d89b41146102ee578063a9059cbb14610303578063d73dd62314610327578063dd62ed3e1461034b578063ec81b48314610372578063f2fde38b14610387578063fca3b5aa146103a8575b600080fd5b34801561010c57600080fd5b506101156103c9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014f578181015183820152602001610137565b50505050905090810190601f16801561017c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019657600080fd5b5061019f610457565b60408051600160a060020a039092168252519081900360200190f35b3480156101c757600080fd5b506101df600160a060020a0360043516602435610466565b604080519115158252519081900360200190f35b3480156101ff57600080fd5b506102086104cc565b60408051918252519081900360200190f35b34801561022657600080fd5b506101df600160a060020a03600435811690602435166044356104d2565b34801561025057600080fd5b5061020861072a565b34801561026557600080fd5b506101df600160a060020a0360043516602435610730565b34801561028957600080fd5b506101df600160a060020a03600435166024356108ca565b3480156102ad57600080fd5b50610208600160a060020a03600435166109b9565b3480156102ce57600080fd5b506102d76109d4565b005b3480156102e557600080fd5b5061019f610a79565b3480156102fa57600080fd5b50610115610a88565b34801561030f57600080fd5b506101df600160a060020a0360043516602435610ae3565b34801561033357600080fd5b506101df600160a060020a0360043516602435610c5a565b34801561035757600080fd5b50610208600160a060020a0360043581169060243516610cf3565b34801561037e57600080fd5b50610208610d1e565b34801561039357600080fd5b506102d7600160a060020a0360043516610d24565b3480156103b457600080fd5b506102d7600160a060020a0360043516610d80565b6004805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561044f5780601f106104245761010080835404028352916020019161044f565b820191906000526020600020905b81548152906001019060200180831161043257829003601f168201915b505050505081565b600154600160a060020a031681565b336000818152600360209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60085490565b600160a060020a038316600090815260026020526040812054821115610542576040805160e560020a62461bcd02815260206004820152601060248201527f4e6f7420656e6f7567682066756e647300000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03841660009081526003602090815260408083203384529091529020548211156105bd576040805160e560020a62461bcd02815260206004820152600c60248201527f4e6f7420617070726f7665640000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038316151561061d576040805160e560020a62461bcd02815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b600160a060020a038416600090815260026020526040902054610646908363ffffffff610dd916565b600160a060020a03808616600090815260026020526040808220939093559085168152205461067b908363ffffffff610e3b16565b600160a060020a0380851660009081526002602090815260408083209490945591871681526003825282812033825290915220546106bf908363ffffffff610dd916565b600160a060020a03808616600081815260036020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b60065481565b600154600090600160a060020a03163314610783576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610f97833981519152604482015290519081900360640190fd5b600754600854610799908463ffffffff610e3b16565b11156107ef576040805160e560020a62461bcd02815260206004820152600b60248201527f4578636565647320636170000000000000000000000000000000000000000000604482015290519081900360640190fd5b600854610802908363ffffffff610e3b16565b600855600160a060020a03831660009081526002602052604090205461082e908363ffffffff610e3b16565b600160a060020a038416600081815260026020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b336000908152600360209081526040808320600160a060020a038616845290915281205480831061091e57336000908152600360209081526040808320600160a060020a0388168452909152812055610953565b61092e818463ffffffff610dd916565b336000908152600360209081526040808320600160a060020a03891684529091529020555b336000818152600360209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526002602052604090205490565b600054600160a060020a03163314610a24576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610f97833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561044f5780601f106104245761010080835404028352916020019161044f565b33600090815260026020526040812054821115610b4a576040805160e560020a62461bcd02815260206004820152601060248201527f4e6f7420656e6f7567682066756e647300000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0383161515610baa576040805160e560020a62461bcd02815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b33600090815260026020526040902054610bca908363ffffffff610dd916565b3360009081526002602052604080822092909255600160a060020a03851681522054610bfc908363ffffffff610e3b16565b600160a060020a0384166000818152600260209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600360209081526040808320600160a060020a0386168452909152812054610c8e908363ffffffff610e3b16565b336000818152600360209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b60075481565b600054600160a060020a03163314610d74576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610f97833981519152604482015290519081900360640190fd5b610d7d81610e9f565b50565b600054600160a060020a03163314610dd0576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610f97833981519152604482015290519081900360640190fd5b610d7d81610f67565b60008083831115610e34576040805160e560020a62461bcd02815260206004820152601060248201527f536166654d617468206661696c75726500000000000000000000000000000000604482015290519081900360640190fd5b5050900390565b600082820183811015610e98576040805160e560020a62461bcd02815260206004820152601060248201527f536166654d617468206661696c75726500000000000000000000000000000000604482015290519081900360640190fd5b9392505050565b600160a060020a0381161515610eff576040805160e560020a62461bcd02815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556005065726d697373696f6e2064656e696564000000000000000000000000000000a165627a7a72305820a3730af89dfc030c1c52c5459dd18eb84e1bdd222ca3aa97c2ca0a279a4927e00029`

// DeployIxoERC20Token deploys a new Ethereum contract, binding an instance of IxoERC20Token to it.
func DeployIxoERC20Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IxoERC20Token, error) {
	parsed, err := abi.JSON(strings.NewReader(IxoERC20TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IxoERC20TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IxoERC20Token{IxoERC20TokenCaller: IxoERC20TokenCaller{contract: contract}, IxoERC20TokenTransactor: IxoERC20TokenTransactor{contract: contract}, IxoERC20TokenFilterer: IxoERC20TokenFilterer{contract: contract}}, nil
}

// IxoERC20Token is an auto generated Go binding around an Ethereum contract.
type IxoERC20Token struct {
	IxoERC20TokenCaller     // Read-only binding to the contract
	IxoERC20TokenTransactor // Write-only binding to the contract
	IxoERC20TokenFilterer   // Log filterer for contract events
}

// IxoERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IxoERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IxoERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IxoERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IxoERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IxoERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IxoERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IxoERC20TokenSession struct {
	Contract     *IxoERC20Token    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IxoERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IxoERC20TokenCallerSession struct {
	Contract *IxoERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IxoERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IxoERC20TokenTransactorSession struct {
	Contract     *IxoERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IxoERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IxoERC20TokenRaw struct {
	Contract *IxoERC20Token // Generic contract binding to access the raw methods on
}

// IxoERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IxoERC20TokenCallerRaw struct {
	Contract *IxoERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// IxoERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IxoERC20TokenTransactorRaw struct {
	Contract *IxoERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIxoERC20Token creates a new instance of IxoERC20Token, bound to a specific deployed contract.
func NewIxoERC20Token(address common.Address, backend bind.ContractBackend) (*IxoERC20Token, error) {
	contract, err := bindIxoERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IxoERC20Token{IxoERC20TokenCaller: IxoERC20TokenCaller{contract: contract}, IxoERC20TokenTransactor: IxoERC20TokenTransactor{contract: contract}, IxoERC20TokenFilterer: IxoERC20TokenFilterer{contract: contract}}, nil
}

// NewIxoERC20TokenCaller creates a new read-only instance of IxoERC20Token, bound to a specific deployed contract.
func NewIxoERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*IxoERC20TokenCaller, error) {
	contract, err := bindIxoERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenCaller{contract: contract}, nil
}

// NewIxoERC20TokenTransactor creates a new write-only instance of IxoERC20Token, bound to a specific deployed contract.
func NewIxoERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IxoERC20TokenTransactor, error) {
	contract, err := bindIxoERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenTransactor{contract: contract}, nil
}

// NewIxoERC20TokenFilterer creates a new log filterer instance of IxoERC20Token, bound to a specific deployed contract.
func NewIxoERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IxoERC20TokenFilterer, error) {
	contract, err := bindIxoERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenFilterer{contract: contract}, nil
}

// bindIxoERC20Token binds a generic wrapper to an already deployed contract.
func bindIxoERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IxoERC20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IxoERC20Token *IxoERC20TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IxoERC20Token.Contract.IxoERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IxoERC20Token *IxoERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.IxoERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IxoERC20Token *IxoERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.IxoERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IxoERC20Token *IxoERC20TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IxoERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IxoERC20Token *IxoERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IxoERC20Token *IxoERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.contract.Transact(opts, method, params...)
}

// CAP is a free data retrieval call binding the contract method 0xec81b483.
//
// Solidity: function CAP() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCaller) CAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "CAP")
	return *ret0, err
}

// CAP is a free data retrieval call binding the contract method 0xec81b483.
//
// Solidity: function CAP() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenSession) CAP() (*big.Int, error) {
	return _IxoERC20Token.Contract.CAP(&_IxoERC20Token.CallOpts)
}

// CAP is a free data retrieval call binding the contract method 0xec81b483.
//
// Solidity: function CAP() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCallerSession) CAP() (*big.Int, error) {
	return _IxoERC20Token.Contract.CAP(&_IxoERC20Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IxoERC20Token.Contract.Allowance(&_IxoERC20Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IxoERC20Token.Contract.Allowance(&_IxoERC20Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _IxoERC20Token.Contract.BalanceOf(&_IxoERC20Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _IxoERC20Token.Contract.BalanceOf(&_IxoERC20Token.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenSession) Decimals() (*big.Int, error) {
	return _IxoERC20Token.Contract.Decimals(&_IxoERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCallerSession) Decimals() (*big.Int, error) {
	return _IxoERC20Token.Contract.Decimals(&_IxoERC20Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_IxoERC20Token *IxoERC20TokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "minter")
	return *ret0, err
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_IxoERC20Token *IxoERC20TokenSession) Minter() (common.Address, error) {
	return _IxoERC20Token.Contract.Minter(&_IxoERC20Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_IxoERC20Token *IxoERC20TokenCallerSession) Minter() (common.Address, error) {
	return _IxoERC20Token.Contract.Minter(&_IxoERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IxoERC20Token *IxoERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IxoERC20Token *IxoERC20TokenSession) Name() (string, error) {
	return _IxoERC20Token.Contract.Name(&_IxoERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IxoERC20Token *IxoERC20TokenCallerSession) Name() (string, error) {
	return _IxoERC20Token.Contract.Name(&_IxoERC20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IxoERC20Token *IxoERC20TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IxoERC20Token *IxoERC20TokenSession) Owner() (common.Address, error) {
	return _IxoERC20Token.Contract.Owner(&_IxoERC20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IxoERC20Token *IxoERC20TokenCallerSession) Owner() (common.Address, error) {
	return _IxoERC20Token.Contract.Owner(&_IxoERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IxoERC20Token *IxoERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IxoERC20Token *IxoERC20TokenSession) Symbol() (string, error) {
	return _IxoERC20Token.Contract.Symbol(&_IxoERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IxoERC20Token *IxoERC20TokenCallerSession) Symbol() (string, error) {
	return _IxoERC20Token.Contract.Symbol(&_IxoERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IxoERC20Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenSession) TotalSupply() (*big.Int, error) {
	return _IxoERC20Token.Contract.TotalSupply(&_IxoERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IxoERC20Token *IxoERC20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IxoERC20Token.Contract.TotalSupply(&_IxoERC20Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.Approve(&_IxoERC20Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.Approve(&_IxoERC20Token.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.DecreaseApproval(&_IxoERC20Token.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.DecreaseApproval(&_IxoERC20Token.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.IncreaseApproval(&_IxoERC20Token.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.IncreaseApproval(&_IxoERC20Token.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.Mint(&_IxoERC20Token.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.Mint(&_IxoERC20Token.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IxoERC20Token *IxoERC20TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IxoERC20Token *IxoERC20TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _IxoERC20Token.Contract.RenounceOwnership(&_IxoERC20Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IxoERC20Token *IxoERC20TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IxoERC20Token.Contract.RenounceOwnership(&_IxoERC20Token.TransactOpts)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(_newMinter address) returns()
func (_IxoERC20Token *IxoERC20TokenTransactor) SetMinter(opts *bind.TransactOpts, _newMinter common.Address) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "setMinter", _newMinter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(_newMinter address) returns()
func (_IxoERC20Token *IxoERC20TokenSession) SetMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.SetMinter(&_IxoERC20Token.TransactOpts, _newMinter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(_newMinter address) returns()
func (_IxoERC20Token *IxoERC20TokenTransactorSession) SetMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.SetMinter(&_IxoERC20Token.TransactOpts, _newMinter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.Transfer(&_IxoERC20Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.Transfer(&_IxoERC20Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.TransferFrom(&_IxoERC20Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_IxoERC20Token *IxoERC20TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.TransferFrom(&_IxoERC20Token.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_IxoERC20Token *IxoERC20TokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _IxoERC20Token.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_IxoERC20Token *IxoERC20TokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.TransferOwnership(&_IxoERC20Token.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_IxoERC20Token *IxoERC20TokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _IxoERC20Token.Contract.TransferOwnership(&_IxoERC20Token.TransactOpts, _newOwner)
}

// IxoERC20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IxoERC20Token contract.
type IxoERC20TokenApprovalIterator struct {
	Event *IxoERC20TokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IxoERC20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IxoERC20TokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IxoERC20TokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IxoERC20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IxoERC20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IxoERC20TokenApproval represents a Approval event raised by the IxoERC20Token contract.
type IxoERC20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IxoERC20Token *IxoERC20TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IxoERC20TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IxoERC20Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenApprovalIterator{contract: _IxoERC20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IxoERC20Token *IxoERC20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IxoERC20TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IxoERC20Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IxoERC20TokenApproval)
				if err := _IxoERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IxoERC20TokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IxoERC20Token contract.
type IxoERC20TokenMintIterator struct {
	Event *IxoERC20TokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IxoERC20TokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IxoERC20TokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IxoERC20TokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IxoERC20TokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IxoERC20TokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IxoERC20TokenMint represents a Mint event raised by the IxoERC20Token contract.
type IxoERC20TokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_IxoERC20Token *IxoERC20TokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*IxoERC20TokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IxoERC20Token.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenMintIterator{contract: _IxoERC20Token.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_IxoERC20Token *IxoERC20TokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IxoERC20TokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IxoERC20Token.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IxoERC20TokenMint)
				if err := _IxoERC20Token.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IxoERC20TokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the IxoERC20Token contract.
type IxoERC20TokenOwnershipRenouncedIterator struct {
	Event *IxoERC20TokenOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IxoERC20TokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IxoERC20TokenOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IxoERC20TokenOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IxoERC20TokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IxoERC20TokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IxoERC20TokenOwnershipRenounced represents a OwnershipRenounced event raised by the IxoERC20Token contract.
type IxoERC20TokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_IxoERC20Token *IxoERC20TokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*IxoERC20TokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _IxoERC20Token.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenOwnershipRenouncedIterator{contract: _IxoERC20Token.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_IxoERC20Token *IxoERC20TokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *IxoERC20TokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _IxoERC20Token.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IxoERC20TokenOwnershipRenounced)
				if err := _IxoERC20Token.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IxoERC20TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IxoERC20Token contract.
type IxoERC20TokenOwnershipTransferredIterator struct {
	Event *IxoERC20TokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IxoERC20TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IxoERC20TokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IxoERC20TokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IxoERC20TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IxoERC20TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IxoERC20TokenOwnershipTransferred represents a OwnershipTransferred event raised by the IxoERC20Token contract.
type IxoERC20TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_IxoERC20Token *IxoERC20TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IxoERC20TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IxoERC20Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenOwnershipTransferredIterator{contract: _IxoERC20Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_IxoERC20Token *IxoERC20TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IxoERC20TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IxoERC20Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IxoERC20TokenOwnershipTransferred)
				if err := _IxoERC20Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IxoERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IxoERC20Token contract.
type IxoERC20TokenTransferIterator struct {
	Event *IxoERC20TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IxoERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IxoERC20TokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IxoERC20TokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IxoERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IxoERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IxoERC20TokenTransfer represents a Transfer event raised by the IxoERC20Token contract.
type IxoERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IxoERC20Token *IxoERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IxoERC20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IxoERC20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IxoERC20TokenTransferIterator{contract: _IxoERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IxoERC20Token *IxoERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IxoERC20TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IxoERC20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IxoERC20TokenTransfer)
				if err := _IxoERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610331806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610192565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a03600435166101a1565b600054600160a060020a0316331461013d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461021a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b61022381610226565b50565b600160a060020a038116151561029d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582071daa85482209acd4a991db4de5b26d80c47fbb1ade66e6fcf57112ce3b6b23a0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ProjectWalletABI is the input ABI used to generate the binding from.
const ProjectWalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProjectWalletBin is the compiled bytecode used for deploying new contracts.
const ProjectWalletBin = `0x`

// DeployProjectWallet deploys a new Ethereum contract, binding an instance of ProjectWallet to it.
func DeployProjectWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProjectWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProjectWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProjectWallet{ProjectWalletCaller: ProjectWalletCaller{contract: contract}, ProjectWalletTransactor: ProjectWalletTransactor{contract: contract}, ProjectWalletFilterer: ProjectWalletFilterer{contract: contract}}, nil
}

// ProjectWallet is an auto generated Go binding around an Ethereum contract.
type ProjectWallet struct {
	ProjectWalletCaller     // Read-only binding to the contract
	ProjectWalletTransactor // Write-only binding to the contract
	ProjectWalletFilterer   // Log filterer for contract events
}

// ProjectWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectWalletSession struct {
	Contract     *ProjectWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectWalletCallerSession struct {
	Contract *ProjectWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProjectWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectWalletTransactorSession struct {
	Contract     *ProjectWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProjectWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectWalletRaw struct {
	Contract *ProjectWallet // Generic contract binding to access the raw methods on
}

// ProjectWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectWalletCallerRaw struct {
	Contract *ProjectWalletCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectWalletTransactorRaw struct {
	Contract *ProjectWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectWallet creates a new instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWallet(address common.Address, backend bind.ContractBackend) (*ProjectWallet, error) {
	contract, err := bindProjectWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectWallet{ProjectWalletCaller: ProjectWalletCaller{contract: contract}, ProjectWalletTransactor: ProjectWalletTransactor{contract: contract}, ProjectWalletFilterer: ProjectWalletFilterer{contract: contract}}, nil
}

// NewProjectWalletCaller creates a new read-only instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWalletCaller(address common.Address, caller bind.ContractCaller) (*ProjectWalletCaller, error) {
	contract, err := bindProjectWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletCaller{contract: contract}, nil
}

// NewProjectWalletTransactor creates a new write-only instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectWalletTransactor, error) {
	contract, err := bindProjectWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletTransactor{contract: contract}, nil
}

// NewProjectWalletFilterer creates a new log filterer instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectWalletFilterer, error) {
	contract, err := bindProjectWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFilterer{contract: contract}, nil
}

// bindProjectWallet binds a generic wrapper to an already deployed contract.
func bindProjectWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWallet *ProjectWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWallet.Contract.ProjectWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWallet *ProjectWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWallet.Contract.ProjectWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWallet *ProjectWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWallet.Contract.ProjectWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWallet *ProjectWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWallet *ProjectWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWallet *ProjectWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWallet.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_ProjectWallet *ProjectWalletTransactor) Transfer(opts *bind.TransactOpts, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWallet.contract.Transact(opts, "transfer", _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_ProjectWallet *ProjectWalletSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWallet.Contract.Transfer(&_ProjectWallet.TransactOpts, _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_ProjectWallet *ProjectWalletTransactorSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWallet.Contract.Transfer(&_ProjectWallet.TransactOpts, _receiver, _amt)
}

// ProjectWalletFactoryABI is the input ABI used to generate the binding from.
const ProjectWalletFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_authoriser\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProjectWalletFactoryBin is the compiled bytecode used for deploying new contracts.
const ProjectWalletFactoryBin = `0x608060405234801561001057600080fd5b50610491806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e8eef2708114610045575b600080fd5b34801561005157600080fd5b5061007c73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356100a5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000807fff000000000000000000000000000000000000000000000000000000000000007f010000000000000000000000000000000000000000000000000000000000000084831a0216151561015c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f496e76616c6964206e616d650000000000000000000000000000000000000000604482015290519081900360640190fd5b8484846101676101bc565b73ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152604080820192909252905190819003606001906000f0801580156101b2573d6000803e3d6000fd5b5095945050505050565b604051610299806101cd833901905600608060405234801561001057600080fd5b5060405160608061029983398101604090815281516020830151919092015160008054600160a060020a03948516600160a060020a03199182161790915560018054949093169316929092179055600255610229806100706000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610050578063a9059cbb14610077575b600080fd5b34801561005c57600080fd5b506100656100bc565b60408051918252519081900360200190f35b34801561008357600080fd5b506100a873ffffffffffffffffffffffffffffffffffffffff600435166024356100c2565b604080519115158252519081900360200190f35b60025481565b60015460009073ffffffffffffffffffffffffffffffffffffffff16331461014b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b505050506040513d60208110156101f457600080fd5b509093925050505600a165627a7a72305820a57f44df56894921b295f5260067fe533d856755c79ff5af9d7e46f1324c122b0029a165627a7a723058206b16a3976414825fc4ed35e8a22423602f47b8e1f8d4090ecbb2782879a6f0b40029`

// DeployProjectWalletFactory deploys a new Ethereum contract, binding an instance of ProjectWalletFactory to it.
func DeployProjectWalletFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProjectWalletFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProjectWalletFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProjectWalletFactory{ProjectWalletFactoryCaller: ProjectWalletFactoryCaller{contract: contract}, ProjectWalletFactoryTransactor: ProjectWalletFactoryTransactor{contract: contract}, ProjectWalletFactoryFilterer: ProjectWalletFactoryFilterer{contract: contract}}, nil
}

// ProjectWalletFactory is an auto generated Go binding around an Ethereum contract.
type ProjectWalletFactory struct {
	ProjectWalletFactoryCaller     // Read-only binding to the contract
	ProjectWalletFactoryTransactor // Write-only binding to the contract
	ProjectWalletFactoryFilterer   // Log filterer for contract events
}

// ProjectWalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectWalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectWalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectWalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectWalletFactorySession struct {
	Contract     *ProjectWalletFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProjectWalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectWalletFactoryCallerSession struct {
	Contract *ProjectWalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ProjectWalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectWalletFactoryTransactorSession struct {
	Contract     *ProjectWalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ProjectWalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectWalletFactoryRaw struct {
	Contract *ProjectWalletFactory // Generic contract binding to access the raw methods on
}

// ProjectWalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectWalletFactoryCallerRaw struct {
	Contract *ProjectWalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectWalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectWalletFactoryTransactorRaw struct {
	Contract *ProjectWalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectWalletFactory creates a new instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactory(address common.Address, backend bind.ContractBackend) (*ProjectWalletFactory, error) {
	contract, err := bindProjectWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactory{ProjectWalletFactoryCaller: ProjectWalletFactoryCaller{contract: contract}, ProjectWalletFactoryTransactor: ProjectWalletFactoryTransactor{contract: contract}, ProjectWalletFactoryFilterer: ProjectWalletFactoryFilterer{contract: contract}}, nil
}

// NewProjectWalletFactoryCaller creates a new read-only instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*ProjectWalletFactoryCaller, error) {
	contract, err := bindProjectWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactoryCaller{contract: contract}, nil
}

// NewProjectWalletFactoryTransactor creates a new write-only instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectWalletFactoryTransactor, error) {
	contract, err := bindProjectWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactoryTransactor{contract: contract}, nil
}

// NewProjectWalletFactoryFilterer creates a new log filterer instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectWalletFactoryFilterer, error) {
	contract, err := bindProjectWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactoryFilterer{contract: contract}, nil
}

// bindProjectWalletFactory binds a generic wrapper to an already deployed contract.
func bindProjectWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletFactory *ProjectWalletFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletFactory.Contract.ProjectWalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletFactory *ProjectWalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.ProjectWalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletFactory *ProjectWalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.ProjectWalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletFactory *ProjectWalletFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletFactory *ProjectWalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletFactory *ProjectWalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(_token address, _authoriser address, _name bytes32) returns(address)
func (_ProjectWalletFactory *ProjectWalletFactoryTransactor) CreateWallet(opts *bind.TransactOpts, _token common.Address, _authoriser common.Address, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletFactory.contract.Transact(opts, "createWallet", _token, _authoriser, _name)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(_token address, _authoriser address, _name bytes32) returns(address)
func (_ProjectWalletFactory *ProjectWalletFactorySession) CreateWallet(_token common.Address, _authoriser common.Address, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.CreateWallet(&_ProjectWalletFactory.TransactOpts, _token, _authoriser, _name)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(_token address, _authoriser address, _name bytes32) returns(address)
func (_ProjectWalletFactory *ProjectWalletFactoryTransactorSession) CreateWallet(_token common.Address, _authoriser common.Address, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.CreateWallet(&_ProjectWalletFactory.TransactOpts, _token, _authoriser, _name)
}

// ProjectWalletRegistryABI is the input ABI used to generate the binding from.
const ProjectWalletRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"walletOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"ensureWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_authoriser\",\"type\":\"address\"},{\"name\":\"_factory\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ProjectWalletRegistryBin is the compiled bytecode used for deploying new contracts.
const ProjectWalletRegistryBin = `0x608060405234801561001057600080fd5b506040516060806106c08339810160409081528151602083015191909201516000805433600160a060020a0319918216178255600180548216600160a060020a0396871617905560028054821694861694909417909355600380549093169390911692909217905561063890819061008890396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166309521458811461007c5780634ea06936146100b05780635bb47808146100c8578063715018a6146100eb5780638da5cb5b14610100578063f2fde38b14610115575b600080fd5b34801561008857600080fd5b50610094600435610136565b60408051600160a060020a039092168252519081900360200190f35b3480156100bc57600080fd5b50610094600435610151565b3480156100d457600080fd5b506100e9600160a060020a0360043516610235565b005b3480156100f757600080fd5b506100e9610328565b34801561010c57600080fd5b506100946103df565b34801561012157600080fd5b506100e9600160a060020a03600435166103ee565b600090815260046020526040902054600160a060020a031690565b6000807fff000000000000000000000000000000000000000000000000000000000000007f010000000000000000000000000000000000000000000000000000000000000084831a021615156101f1576040805160e560020a62461bcd02815260206004820152600c60248201527f496e76616c6964206e616d650000000000000000000000000000000000000000604482015290519081900360640190fd5b600083815260046020526040902054600160a060020a03161515610218576102188361045c565b5050600090815260046020526040902054600160a060020a031690565b600054600160a060020a03163314610297576040805160e560020a62461bcd02815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b600354600160a060020a031615156102f9576040805160e560020a62461bcd02815260206004820152600f60248201527f496e76616c696420666163746f72790000000000000000000000000000000000604482015290519081900360640190fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461038a576040805160e560020a62461bcd02815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a03163314610450576040805160e560020a62461bcd02815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b61045981610544565b50565b600354600154600254604080517fe8eef270000000000000000000000000000000000000000000000000000000008152600160a060020a03938416600482015291831660248301526044820185905251600093929092169163e8eef2709160648082019260209290919082900301818787803b1580156104db57600080fd5b505af11580156104ef573d6000803e3d6000fd5b505050506040513d602081101561050557600080fd5b5051600092835260046020526040909220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039093169290921790915550565b600160a060020a03811615156105a4576040805160e560020a62461bcd02815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820a593ba3fa7deac056725200706eb1c670c4337651567b370f17c5ad92d87c1850029`

// DeployProjectWalletRegistry deploys a new Ethereum contract, binding an instance of ProjectWalletRegistry to it.
func DeployProjectWalletRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _authoriser common.Address, _factory common.Address) (common.Address, *types.Transaction, *ProjectWalletRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProjectWalletRegistryBin), backend, _token, _authoriser, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProjectWalletRegistry{ProjectWalletRegistryCaller: ProjectWalletRegistryCaller{contract: contract}, ProjectWalletRegistryTransactor: ProjectWalletRegistryTransactor{contract: contract}, ProjectWalletRegistryFilterer: ProjectWalletRegistryFilterer{contract: contract}}, nil
}

// ProjectWalletRegistry is an auto generated Go binding around an Ethereum contract.
type ProjectWalletRegistry struct {
	ProjectWalletRegistryCaller     // Read-only binding to the contract
	ProjectWalletRegistryTransactor // Write-only binding to the contract
	ProjectWalletRegistryFilterer   // Log filterer for contract events
}

// ProjectWalletRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectWalletRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectWalletRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectWalletRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectWalletRegistrySession struct {
	Contract     *ProjectWalletRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ProjectWalletRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectWalletRegistryCallerSession struct {
	Contract *ProjectWalletRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ProjectWalletRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectWalletRegistryTransactorSession struct {
	Contract     *ProjectWalletRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ProjectWalletRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectWalletRegistryRaw struct {
	Contract *ProjectWalletRegistry // Generic contract binding to access the raw methods on
}

// ProjectWalletRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectWalletRegistryCallerRaw struct {
	Contract *ProjectWalletRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectWalletRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectWalletRegistryTransactorRaw struct {
	Contract *ProjectWalletRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectWalletRegistry creates a new instance of ProjectWalletRegistry, bound to a specific deployed contract.
func NewProjectWalletRegistry(address common.Address, backend bind.ContractBackend) (*ProjectWalletRegistry, error) {
	contract, err := bindProjectWalletRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletRegistry{ProjectWalletRegistryCaller: ProjectWalletRegistryCaller{contract: contract}, ProjectWalletRegistryTransactor: ProjectWalletRegistryTransactor{contract: contract}, ProjectWalletRegistryFilterer: ProjectWalletRegistryFilterer{contract: contract}}, nil
}

// NewProjectWalletRegistryCaller creates a new read-only instance of ProjectWalletRegistry, bound to a specific deployed contract.
func NewProjectWalletRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProjectWalletRegistryCaller, error) {
	contract, err := bindProjectWalletRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletRegistryCaller{contract: contract}, nil
}

// NewProjectWalletRegistryTransactor creates a new write-only instance of ProjectWalletRegistry, bound to a specific deployed contract.
func NewProjectWalletRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectWalletRegistryTransactor, error) {
	contract, err := bindProjectWalletRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletRegistryTransactor{contract: contract}, nil
}

// NewProjectWalletRegistryFilterer creates a new log filterer instance of ProjectWalletRegistry, bound to a specific deployed contract.
func NewProjectWalletRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectWalletRegistryFilterer, error) {
	contract, err := bindProjectWalletRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletRegistryFilterer{contract: contract}, nil
}

// bindProjectWalletRegistry binds a generic wrapper to an already deployed contract.
func bindProjectWalletRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletRegistry *ProjectWalletRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletRegistry.Contract.ProjectWalletRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletRegistry *ProjectWalletRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.ProjectWalletRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletRegistry *ProjectWalletRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.ProjectWalletRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletRegistry *ProjectWalletRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProjectWalletRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistrySession) Owner() (common.Address, error) {
	return _ProjectWalletRegistry.Contract.Owner(&_ProjectWalletRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistryCallerSession) Owner() (common.Address, error) {
	return _ProjectWalletRegistry.Contract.Owner(&_ProjectWalletRegistry.CallOpts)
}

// WalletOf is a free data retrieval call binding the contract method 0x09521458.
//
// Solidity: function walletOf(_name bytes32) constant returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistryCaller) WalletOf(opts *bind.CallOpts, _name [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProjectWalletRegistry.contract.Call(opts, out, "walletOf", _name)
	return *ret0, err
}

// WalletOf is a free data retrieval call binding the contract method 0x09521458.
//
// Solidity: function walletOf(_name bytes32) constant returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistrySession) WalletOf(_name [32]byte) (common.Address, error) {
	return _ProjectWalletRegistry.Contract.WalletOf(&_ProjectWalletRegistry.CallOpts, _name)
}

// WalletOf is a free data retrieval call binding the contract method 0x09521458.
//
// Solidity: function walletOf(_name bytes32) constant returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistryCallerSession) WalletOf(_name [32]byte) (common.Address, error) {
	return _ProjectWalletRegistry.Contract.WalletOf(&_ProjectWalletRegistry.CallOpts, _name)
}

// EnsureWallet is a paid mutator transaction binding the contract method 0x4ea06936.
//
// Solidity: function ensureWallet(_name bytes32) returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactor) EnsureWallet(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletRegistry.contract.Transact(opts, "ensureWallet", _name)
}

// EnsureWallet is a paid mutator transaction binding the contract method 0x4ea06936.
//
// Solidity: function ensureWallet(_name bytes32) returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistrySession) EnsureWallet(_name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.EnsureWallet(&_ProjectWalletRegistry.TransactOpts, _name)
}

// EnsureWallet is a paid mutator transaction binding the contract method 0x4ea06936.
//
// Solidity: function ensureWallet(_name bytes32) returns(address)
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactorSession) EnsureWallet(_name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.EnsureWallet(&_ProjectWalletRegistry.TransactOpts, _name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectWalletRegistry *ProjectWalletRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.RenounceOwnership(&_ProjectWalletRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.RenounceOwnership(&_ProjectWalletRegistry.TransactOpts)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(_factory address) returns()
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactor) SetFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _ProjectWalletRegistry.contract.Transact(opts, "setFactory", _factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(_factory address) returns()
func (_ProjectWalletRegistry *ProjectWalletRegistrySession) SetFactory(_factory common.Address) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.SetFactory(&_ProjectWalletRegistry.TransactOpts, _factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(_factory address) returns()
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactorSession) SetFactory(_factory common.Address) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.SetFactory(&_ProjectWalletRegistry.TransactOpts, _factory)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ProjectWalletRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ProjectWalletRegistry *ProjectWalletRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.TransferOwnership(&_ProjectWalletRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ProjectWalletRegistry *ProjectWalletRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ProjectWalletRegistry.Contract.TransferOwnership(&_ProjectWalletRegistry.TransactOpts, _newOwner)
}

// ProjectWalletRegistryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ProjectWalletRegistry contract.
type ProjectWalletRegistryOwnershipRenouncedIterator struct {
	Event *ProjectWalletRegistryOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProjectWalletRegistryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectWalletRegistryOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProjectWalletRegistryOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProjectWalletRegistryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectWalletRegistryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectWalletRegistryOwnershipRenounced represents a OwnershipRenounced event raised by the ProjectWalletRegistry contract.
type ProjectWalletRegistryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ProjectWalletRegistry *ProjectWalletRegistryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ProjectWalletRegistryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ProjectWalletRegistry.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletRegistryOwnershipRenouncedIterator{contract: _ProjectWalletRegistry.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ProjectWalletRegistry *ProjectWalletRegistryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ProjectWalletRegistryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ProjectWalletRegistry.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectWalletRegistryOwnershipRenounced)
				if err := _ProjectWalletRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ProjectWalletRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProjectWalletRegistry contract.
type ProjectWalletRegistryOwnershipTransferredIterator struct {
	Event *ProjectWalletRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProjectWalletRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectWalletRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProjectWalletRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProjectWalletRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectWalletRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectWalletRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ProjectWalletRegistry contract.
type ProjectWalletRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProjectWalletRegistry *ProjectWalletRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProjectWalletRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProjectWalletRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletRegistryOwnershipTransferredIterator{contract: _ProjectWalletRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProjectWalletRegistry *ProjectWalletRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProjectWalletRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProjectWalletRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectWalletRegistryOwnershipTransferred)
				if err := _ProjectWalletRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a7c6ce7325c87326dd4cfc2a49281bb7d225122fcd883c7c37bf3fc34b51282d0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
