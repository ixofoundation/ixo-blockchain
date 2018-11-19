// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auth

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

// AuthContractABI is the input ABI used to generate the binding from.
const AuthContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"memberCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_quorum\",\"type\":\"uint256\"}],\"name\":\"setQuorum\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tx\",\"type\":\"bytes32\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"validate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"actions\",\"outputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amt\",\"type\":\"uint256\"},{\"name\":\"confirmations\",\"type\":\"uint256\"},{\"name\":\"triggered\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_members\",\"type\":\"address[]\"},{\"name\":\"_quorum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"member\",\"type\":\"address\"}],\"name\":\"Confirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Triggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AuthContractBin is the compiled bytecode used for deploying new contracts.
const AuthContractBin = `0x608060405234801561001057600080fd5b50604051610c03380380610c0383398101604052805160208083015160008054600160a060020a0319163317815592909301805190939291610057916001918601906100c4565b5050600281905560005b6001548110156100bc5760016004600060018481548110151561008057fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff1916911515919091179055600101610061565b505050610150565b828054828255906000526020600020908101928215610119579160200282015b828111156101195782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906100e4565b50610125929150610129565b5090565b61014d91905b80821115610125578054600160a060020a031916815560010161012f565b90565b610aa48061015f6000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166311aee38081146100b35780631703a018146100da5780634e831d7a146100ef5780635daf08ca14610127578063715018a61461015b5780638da5cb5b14610172578063a230c52414610187578063c1ba4e59146101a8578063d53764e0146101c0578063f2fde38b146101f3578063f3abde3214610214575b600080fd5b3480156100bf57600080fd5b506100c861026e565b60408051918252519081900360200190f35b3480156100e657600080fd5b506100c8610274565b3480156100fb57600080fd5b50610113600435600160a060020a036024351661027a565b604080519115158252519081900360200190f35b34801561013357600080fd5b5061013f60043561029a565b60408051600160a060020a039092168252519081900360200190f35b34801561016757600080fd5b506101706102c2565b005b34801561017e57600080fd5b5061013f610379565b34801561019357600080fd5b50610113600160a060020a0360043516610388565b3480156101b457600080fd5b5061017060043561039d565b3480156101cc57600080fd5b50610113600435600160a060020a0360243581169060443581169060643516608435610404565b3480156101ff57600080fd5b50610170600160a060020a0360043516610756565b34801561022057600080fd5b5061022c6004356107c4565b60408051600160a060020a039788168152958716602087015293909516848401526060840191909152608083015291151560a082015290519081900360c00190f35b60015490565b60025481565b600360209081526000928352604080842090915290825290205460ff1681565b60018054829081106102a857fe5b600091825260209091200154600160a060020a0316905081565b600054600160a060020a03163314610324576040805160e560020a62461bcd02815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b60046020526000908152604090205460ff1681565b600054600160a060020a031633146103ff576040805160e560020a62461bcd02815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b600255565b3360009081526004602052604081205460ff16151561046d576040805160e560020a62461bcd02815260206004820152600c60248201527f4e6f742061206d656d6265720000000000000000000000000000000000000000604482015290519081900360640190fd5b8515156104c4576040805160e560020a62461bcd02815260206004820152601660248201527f496e76616c6964207472616e73616374696f6e20696400000000000000000000604482015290519081900360640190fd5b600160a060020a0385161515610524576040805160e560020a62461bcd02815260206004820152600e60248201527f496e76616c696420746172676574000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0384161515610584576040805160e560020a62461bcd02815260206004820152600e60248201527f496e76616c69642073656e646572000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03831615156105e4576040805160e560020a62461bcd02815260206004820152601060248201527f496e76616c696420726563656976657200000000000000000000000000000000604482015290519081900360640190fd5b600082101561063d576040805160e560020a62461bcd02815260206004820152600e60248201527f496e76616c696420616d6f756e74000000000000000000000000000000000000604482015290519081900360640190fd5b6000868152600560208181526040928390208054600160a060020a03808b1673ffffffffffffffffffffffffffffffffffffffff19928316178355600180840180548c8416908516179055600284018054928b16929093169190911790915560038201879055928101805460ff1916905560040180549092019091558151888152339181019190915281517fd4964a7cd99f5c1fa8f2420fb5e1d3bd26eadf16e2658cf2e29a67dfda38601e929181900390910190a16002546000878152600560205260409020600401541061074a576107168661080b565b6040805187815290517f8dec26062ce2d31c0d76915d8ae104afcbd6bd4c80c98e58f2441fa66ab07b0c9181900360200190a15b50600195945050505050565b600054600160a060020a031633146107b8576040805160e560020a62461bcd02815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b6107c1816109b0565b50565b6005602081905260009182526040909120805460018201546002830154600384015460048501549490950154600160a060020a039384169592841694919093169260ff1686565b3360009081526004602052604090205460ff161515610874576040805160e560020a62461bcd02815260206004820152600c60248201527f4e6f742061206d656d6265720000000000000000000000000000000000000000604482015290519081900360640190fd5b60008181526005602081905260409091200154819060ff16156108e1576040805160e560020a62461bcd02815260206004820152601d60248201527f5472616e73616374696f6e20616c726561647920747269676765726564000000604482015290519081900360640190fd5b6000828152600560208181526040808420928301805460ff19166001908117909155835490840154600285015460039095015483517fbeabacc8000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015295821660248701526044860152915191169363beabacc8936064808201949392918390030190829087803b15801561098057600080fd5b505af1158015610994573d6000803e3d6000fd5b505050506040513d60208110156109aa57600080fd5b50505050565b600160a060020a0381161515610a10576040805160e560020a62461bcd02815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820dfa428ee20e9748ba3392af3fca9ea6e984dcf7d9ceadf7210467bad217ffdb50029`

// DeployAuthContract deploys a new Ethereum contract, binding an instance of AuthContract to it.
func DeployAuthContract(auth *bind.TransactOpts, backend bind.ContractBackend, _members []common.Address, _quorum *big.Int) (common.Address, *types.Transaction, *AuthContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthContractBin), backend, _members, _quorum)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuthContract{AuthContractCaller: AuthContractCaller{contract: contract}, AuthContractTransactor: AuthContractTransactor{contract: contract}, AuthContractFilterer: AuthContractFilterer{contract: contract}}, nil
}

// AuthContract is an auto generated Go binding around an Ethereum contract.
type AuthContract struct {
	AuthContractCaller     // Read-only binding to the contract
	AuthContractTransactor // Write-only binding to the contract
	AuthContractFilterer   // Log filterer for contract events
}

// AuthContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthContractSession struct {
	Contract     *AuthContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthContractCallerSession struct {
	Contract *AuthContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AuthContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthContractTransactorSession struct {
	Contract     *AuthContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AuthContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthContractRaw struct {
	Contract *AuthContract // Generic contract binding to access the raw methods on
}

// AuthContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthContractCallerRaw struct {
	Contract *AuthContractCaller // Generic read-only contract binding to access the raw methods on
}

// AuthContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthContractTransactorRaw struct {
	Contract *AuthContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthContract creates a new instance of AuthContract, bound to a specific deployed contract.
func NewAuthContract(address common.Address, backend bind.ContractBackend) (*AuthContract, error) {
	contract, err := bindAuthContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthContract{AuthContractCaller: AuthContractCaller{contract: contract}, AuthContractTransactor: AuthContractTransactor{contract: contract}, AuthContractFilterer: AuthContractFilterer{contract: contract}}, nil
}

// NewAuthContractCaller creates a new read-only instance of AuthContract, bound to a specific deployed contract.
func NewAuthContractCaller(address common.Address, caller bind.ContractCaller) (*AuthContractCaller, error) {
	contract, err := bindAuthContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthContractCaller{contract: contract}, nil
}

// NewAuthContractTransactor creates a new write-only instance of AuthContract, bound to a specific deployed contract.
func NewAuthContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthContractTransactor, error) {
	contract, err := bindAuthContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthContractTransactor{contract: contract}, nil
}

// NewAuthContractFilterer creates a new log filterer instance of AuthContract, bound to a specific deployed contract.
func NewAuthContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthContractFilterer, error) {
	contract, err := bindAuthContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthContractFilterer{contract: contract}, nil
}

// bindAuthContract binds a generic wrapper to an already deployed contract.
func bindAuthContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthContract *AuthContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthContract.Contract.AuthContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthContract *AuthContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthContract.Contract.AuthContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthContract *AuthContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthContract.Contract.AuthContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthContract *AuthContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthContract *AuthContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthContract *AuthContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthContract.Contract.contract.Transact(opts, method, params...)
}

// Actions is a free data retrieval call binding the contract method 0xf3abde32.
//
// Solidity: function actions( bytes32) constant returns(target address, sender address, receiver address, amt uint256, confirmations uint256, triggered bool)
func (_AuthContract *AuthContractCaller) Actions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Target        common.Address
	Sender        common.Address
	Receiver      common.Address
	Amt           *big.Int
	Confirmations *big.Int
	Triggered     bool
}, error) {
	ret := new(struct {
		Target        common.Address
		Sender        common.Address
		Receiver      common.Address
		Amt           *big.Int
		Confirmations *big.Int
		Triggered     bool
	})
	out := ret
	err := _AuthContract.contract.Call(opts, out, "actions", arg0)
	return *ret, err
}

// Actions is a free data retrieval call binding the contract method 0xf3abde32.
//
// Solidity: function actions( bytes32) constant returns(target address, sender address, receiver address, amt uint256, confirmations uint256, triggered bool)
func (_AuthContract *AuthContractSession) Actions(arg0 [32]byte) (struct {
	Target        common.Address
	Sender        common.Address
	Receiver      common.Address
	Amt           *big.Int
	Confirmations *big.Int
	Triggered     bool
}, error) {
	return _AuthContract.Contract.Actions(&_AuthContract.CallOpts, arg0)
}

// Actions is a free data retrieval call binding the contract method 0xf3abde32.
//
// Solidity: function actions( bytes32) constant returns(target address, sender address, receiver address, amt uint256, confirmations uint256, triggered bool)
func (_AuthContract *AuthContractCallerSession) Actions(arg0 [32]byte) (struct {
	Target        common.Address
	Sender        common.Address
	Receiver      common.Address
	Amt           *big.Int
	Confirmations *big.Int
	Triggered     bool
}, error) {
	return _AuthContract.Contract.Actions(&_AuthContract.CallOpts, arg0)
}

// ConfirmedBy is a free data retrieval call binding the contract method 0x4e831d7a.
//
// Solidity: function confirmedBy( uint256,  address) constant returns(bool)
func (_AuthContract *AuthContractCaller) ConfirmedBy(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthContract.contract.Call(opts, out, "confirmedBy", arg0, arg1)
	return *ret0, err
}

// ConfirmedBy is a free data retrieval call binding the contract method 0x4e831d7a.
//
// Solidity: function confirmedBy( uint256,  address) constant returns(bool)
func (_AuthContract *AuthContractSession) ConfirmedBy(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _AuthContract.Contract.ConfirmedBy(&_AuthContract.CallOpts, arg0, arg1)
}

// ConfirmedBy is a free data retrieval call binding the contract method 0x4e831d7a.
//
// Solidity: function confirmedBy( uint256,  address) constant returns(bool)
func (_AuthContract *AuthContractCallerSession) ConfirmedBy(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _AuthContract.Contract.ConfirmedBy(&_AuthContract.CallOpts, arg0, arg1)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember( address) constant returns(bool)
func (_AuthContract *AuthContractCaller) IsMember(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthContract.contract.Call(opts, out, "isMember", arg0)
	return *ret0, err
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember( address) constant returns(bool)
func (_AuthContract *AuthContractSession) IsMember(arg0 common.Address) (bool, error) {
	return _AuthContract.Contract.IsMember(&_AuthContract.CallOpts, arg0)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember( address) constant returns(bool)
func (_AuthContract *AuthContractCallerSession) IsMember(arg0 common.Address) (bool, error) {
	return _AuthContract.Contract.IsMember(&_AuthContract.CallOpts, arg0)
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() constant returns(uint256)
func (_AuthContract *AuthContractCaller) MemberCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuthContract.contract.Call(opts, out, "memberCount")
	return *ret0, err
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() constant returns(uint256)
func (_AuthContract *AuthContractSession) MemberCount() (*big.Int, error) {
	return _AuthContract.Contract.MemberCount(&_AuthContract.CallOpts)
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() constant returns(uint256)
func (_AuthContract *AuthContractCallerSession) MemberCount() (*big.Int, error) {
	return _AuthContract.Contract.MemberCount(&_AuthContract.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members( uint256) constant returns(address)
func (_AuthContract *AuthContractCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuthContract.contract.Call(opts, out, "members", arg0)
	return *ret0, err
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members( uint256) constant returns(address)
func (_AuthContract *AuthContractSession) Members(arg0 *big.Int) (common.Address, error) {
	return _AuthContract.Contract.Members(&_AuthContract.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members( uint256) constant returns(address)
func (_AuthContract *AuthContractCallerSession) Members(arg0 *big.Int) (common.Address, error) {
	return _AuthContract.Contract.Members(&_AuthContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AuthContract *AuthContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuthContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AuthContract *AuthContractSession) Owner() (common.Address, error) {
	return _AuthContract.Contract.Owner(&_AuthContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AuthContract *AuthContractCallerSession) Owner() (common.Address, error) {
	return _AuthContract.Contract.Owner(&_AuthContract.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint256)
func (_AuthContract *AuthContractCaller) Quorum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuthContract.contract.Call(opts, out, "quorum")
	return *ret0, err
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint256)
func (_AuthContract *AuthContractSession) Quorum() (*big.Int, error) {
	return _AuthContract.Contract.Quorum(&_AuthContract.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint256)
func (_AuthContract *AuthContractCallerSession) Quorum() (*big.Int, error) {
	return _AuthContract.Contract.Quorum(&_AuthContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthContract *AuthContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthContract *AuthContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuthContract.Contract.RenounceOwnership(&_AuthContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthContract *AuthContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuthContract.Contract.RenounceOwnership(&_AuthContract.TransactOpts)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(_quorum uint256) returns()
func (_AuthContract *AuthContractTransactor) SetQuorum(opts *bind.TransactOpts, _quorum *big.Int) (*types.Transaction, error) {
	return _AuthContract.contract.Transact(opts, "setQuorum", _quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(_quorum uint256) returns()
func (_AuthContract *AuthContractSession) SetQuorum(_quorum *big.Int) (*types.Transaction, error) {
	return _AuthContract.Contract.SetQuorum(&_AuthContract.TransactOpts, _quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0xc1ba4e59.
//
// Solidity: function setQuorum(_quorum uint256) returns()
func (_AuthContract *AuthContractTransactorSession) SetQuorum(_quorum *big.Int) (*types.Transaction, error) {
	return _AuthContract.Contract.SetQuorum(&_AuthContract.TransactOpts, _quorum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AuthContract *AuthContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _AuthContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AuthContract *AuthContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AuthContract.Contract.TransferOwnership(&_AuthContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_AuthContract *AuthContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AuthContract.Contract.TransferOwnership(&_AuthContract.TransactOpts, _newOwner)
}

// Validate is a paid mutator transaction binding the contract method 0xd53764e0.
//
// Solidity: function validate(_tx bytes32, _target address, _sender address, _receiver address, _amt uint256) returns(bool)
func (_AuthContract *AuthContractTransactor) Validate(opts *bind.TransactOpts, _tx [32]byte, _target common.Address, _sender common.Address, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _AuthContract.contract.Transact(opts, "validate", _tx, _target, _sender, _receiver, _amt)
}

// Validate is a paid mutator transaction binding the contract method 0xd53764e0.
//
// Solidity: function validate(_tx bytes32, _target address, _sender address, _receiver address, _amt uint256) returns(bool)
func (_AuthContract *AuthContractSession) Validate(_tx [32]byte, _target common.Address, _sender common.Address, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _AuthContract.Contract.Validate(&_AuthContract.TransactOpts, _tx, _target, _sender, _receiver, _amt)
}

// Validate is a paid mutator transaction binding the contract method 0xd53764e0.
//
// Solidity: function validate(_tx bytes32, _target address, _sender address, _receiver address, _amt uint256) returns(bool)
func (_AuthContract *AuthContractTransactorSession) Validate(_tx [32]byte, _target common.Address, _sender common.Address, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _AuthContract.Contract.Validate(&_AuthContract.TransactOpts, _tx, _target, _sender, _receiver, _amt)
}

// AuthContractConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the AuthContract contract.
type AuthContractConfirmedIterator struct {
	Event *AuthContractConfirmed // Event containing the contract specifics and raw log

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
func (it *AuthContractConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthContractConfirmed)
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
		it.Event = new(AuthContractConfirmed)
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
func (it *AuthContractConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthContractConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthContractConfirmed represents a Confirmed event raised by the AuthContract contract.
type AuthContractConfirmed struct {
	Id     [32]byte
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xd4964a7cd99f5c1fa8f2420fb5e1d3bd26eadf16e2658cf2e29a67dfda38601e.
//
// Solidity: e Confirmed(id bytes32, member address)
func (_AuthContract *AuthContractFilterer) FilterConfirmed(opts *bind.FilterOpts) (*AuthContractConfirmedIterator, error) {

	logs, sub, err := _AuthContract.contract.FilterLogs(opts, "Confirmed")
	if err != nil {
		return nil, err
	}
	return &AuthContractConfirmedIterator{contract: _AuthContract.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xd4964a7cd99f5c1fa8f2420fb5e1d3bd26eadf16e2658cf2e29a67dfda38601e.
//
// Solidity: e Confirmed(id bytes32, member address)
func (_AuthContract *AuthContractFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *AuthContractConfirmed) (event.Subscription, error) {

	logs, sub, err := _AuthContract.contract.WatchLogs(opts, "Confirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthContractConfirmed)
				if err := _AuthContract.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// AuthContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the AuthContract contract.
type AuthContractOwnershipRenouncedIterator struct {
	Event *AuthContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AuthContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthContractOwnershipRenounced)
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
		it.Event = new(AuthContractOwnershipRenounced)
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
func (it *AuthContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthContractOwnershipRenounced represents a OwnershipRenounced event raised by the AuthContract contract.
type AuthContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AuthContract *AuthContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AuthContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AuthContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthContractOwnershipRenouncedIterator{contract: _AuthContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AuthContract *AuthContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AuthContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AuthContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthContractOwnershipRenounced)
				if err := _AuthContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// AuthContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuthContract contract.
type AuthContractOwnershipTransferredIterator struct {
	Event *AuthContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuthContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthContractOwnershipTransferred)
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
		it.Event = new(AuthContractOwnershipTransferred)
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
func (it *AuthContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthContractOwnershipTransferred represents a OwnershipTransferred event raised by the AuthContract contract.
type AuthContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AuthContract *AuthContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuthContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthContractOwnershipTransferredIterator{contract: _AuthContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AuthContract *AuthContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuthContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthContractOwnershipTransferred)
				if err := _AuthContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AuthContractTriggeredIterator is returned from FilterTriggered and is used to iterate over the raw logs and unpacked data for Triggered events raised by the AuthContract contract.
type AuthContractTriggeredIterator struct {
	Event *AuthContractTriggered // Event containing the contract specifics and raw log

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
func (it *AuthContractTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthContractTriggered)
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
		it.Event = new(AuthContractTriggered)
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
func (it *AuthContractTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthContractTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthContractTriggered represents a Triggered event raised by the AuthContract contract.
type AuthContractTriggered struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTriggered is a free log retrieval operation binding the contract event 0x8dec26062ce2d31c0d76915d8ae104afcbd6bd4c80c98e58f2441fa66ab07b0c.
//
// Solidity: e Triggered(id bytes32)
func (_AuthContract *AuthContractFilterer) FilterTriggered(opts *bind.FilterOpts) (*AuthContractTriggeredIterator, error) {

	logs, sub, err := _AuthContract.contract.FilterLogs(opts, "Triggered")
	if err != nil {
		return nil, err
	}
	return &AuthContractTriggeredIterator{contract: _AuthContract.contract, event: "Triggered", logs: logs, sub: sub}, nil
}

// WatchTriggered is a free log subscription operation binding the contract event 0x8dec26062ce2d31c0d76915d8ae104afcbd6bd4c80c98e58f2441fa66ab07b0c.
//
// Solidity: e Triggered(id bytes32)
func (_AuthContract *AuthContractFilterer) WatchTriggered(opts *bind.WatchOpts, sink chan<- *AuthContractTriggered) (event.Subscription, error) {

	logs, sub, err := _AuthContract.contract.WatchLogs(opts, "Triggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthContractTriggered)
				if err := _AuthContract.contract.UnpackLog(event, "Triggered", log); err != nil {
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

// ProjectWalletAuthoriserABI is the input ABI used to generate the binding from.
const ProjectWalletAuthoriserABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authoriser\",\"type\":\"address\"}],\"name\":\"setAuthoriser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ProjectWalletAuthoriserBin is the compiled bytecode used for deploying new contracts.
const ProjectWalletAuthoriserBin = `0x608060405260008054600160a060020a031916331790556104e2806100256000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a681146100715780638da5cb5b14610088578063beabacc8146100b9578063f2fde38b146100f7578063f7da0e0414610118575b600080fd5b34801561007d57600080fd5b50610086610139565b005b34801561009457600080fd5b5061009d6101de565b60408051600160a060020a039092168252519081900360200190f35b3480156100c557600080fd5b506100e3600160a060020a03600435811690602435166044356101ed565b604080519115158252519081900360200190f35b34801561010357600080fd5b50610086600160a060020a03600435166102f0565b34801561012457600080fd5b506100e3600160a060020a036004351661034c565b600054600160a060020a03163314610189576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610497833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600154600090600160a060020a03163314610240576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610497833981519152604482015290519081900360640190fd5b83600160a060020a031663a9059cbb84846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156102bc57600080fd5b505af11580156102d0573d6000803e3d6000fd5b505050506040513d60208110156102e657600080fd5b5090949350505050565b600054600160a060020a03163314610340576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610497833981519152604482015290519081900360640190fd5b610349816103ce565b50565b60008054600160a060020a0316331461039d576040805160e560020a62461bcd0281526020600482015260116024820152600080516020610497833981519152604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03939093169290921790915590565b600160a060020a038116151561042e576040805160e560020a62461bcd02815260206004820152601560248201527f43616e2774207472616e7366657220746f203078300000000000000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556005065726d697373696f6e2064656e696564000000000000000000000000000000a165627a7a723058205e169f37b3baf44de188592c09134246f9cff6d4c8edb08dad307bde59e3a0a10029`

// DeployProjectWalletAuthoriser deploys a new Ethereum contract, binding an instance of ProjectWalletAuthoriser to it.
func DeployProjectWalletAuthoriser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProjectWalletAuthoriser, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletAuthoriserABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProjectWalletAuthoriserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProjectWalletAuthoriser{ProjectWalletAuthoriserCaller: ProjectWalletAuthoriserCaller{contract: contract}, ProjectWalletAuthoriserTransactor: ProjectWalletAuthoriserTransactor{contract: contract}, ProjectWalletAuthoriserFilterer: ProjectWalletAuthoriserFilterer{contract: contract}}, nil
}

// ProjectWalletAuthoriser is an auto generated Go binding around an Ethereum contract.
type ProjectWalletAuthoriser struct {
	ProjectWalletAuthoriserCaller     // Read-only binding to the contract
	ProjectWalletAuthoriserTransactor // Write-only binding to the contract
	ProjectWalletAuthoriserFilterer   // Log filterer for contract events
}

// ProjectWalletAuthoriserCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectWalletAuthoriserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletAuthoriserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectWalletAuthoriserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletAuthoriserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectWalletAuthoriserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletAuthoriserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectWalletAuthoriserSession struct {
	Contract     *ProjectWalletAuthoriser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProjectWalletAuthoriserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectWalletAuthoriserCallerSession struct {
	Contract *ProjectWalletAuthoriserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ProjectWalletAuthoriserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectWalletAuthoriserTransactorSession struct {
	Contract     *ProjectWalletAuthoriserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ProjectWalletAuthoriserRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectWalletAuthoriserRaw struct {
	Contract *ProjectWalletAuthoriser // Generic contract binding to access the raw methods on
}

// ProjectWalletAuthoriserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectWalletAuthoriserCallerRaw struct {
	Contract *ProjectWalletAuthoriserCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectWalletAuthoriserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectWalletAuthoriserTransactorRaw struct {
	Contract *ProjectWalletAuthoriserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectWalletAuthoriser creates a new instance of ProjectWalletAuthoriser, bound to a specific deployed contract.
func NewProjectWalletAuthoriser(address common.Address, backend bind.ContractBackend) (*ProjectWalletAuthoriser, error) {
	contract, err := bindProjectWalletAuthoriser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletAuthoriser{ProjectWalletAuthoriserCaller: ProjectWalletAuthoriserCaller{contract: contract}, ProjectWalletAuthoriserTransactor: ProjectWalletAuthoriserTransactor{contract: contract}, ProjectWalletAuthoriserFilterer: ProjectWalletAuthoriserFilterer{contract: contract}}, nil
}

// NewProjectWalletAuthoriserCaller creates a new read-only instance of ProjectWalletAuthoriser, bound to a specific deployed contract.
func NewProjectWalletAuthoriserCaller(address common.Address, caller bind.ContractCaller) (*ProjectWalletAuthoriserCaller, error) {
	contract, err := bindProjectWalletAuthoriser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletAuthoriserCaller{contract: contract}, nil
}

// NewProjectWalletAuthoriserTransactor creates a new write-only instance of ProjectWalletAuthoriser, bound to a specific deployed contract.
func NewProjectWalletAuthoriserTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectWalletAuthoriserTransactor, error) {
	contract, err := bindProjectWalletAuthoriser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletAuthoriserTransactor{contract: contract}, nil
}

// NewProjectWalletAuthoriserFilterer creates a new log filterer instance of ProjectWalletAuthoriser, bound to a specific deployed contract.
func NewProjectWalletAuthoriserFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectWalletAuthoriserFilterer, error) {
	contract, err := bindProjectWalletAuthoriser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletAuthoriserFilterer{contract: contract}, nil
}

// bindProjectWalletAuthoriser binds a generic wrapper to an already deployed contract.
func bindProjectWalletAuthoriser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletAuthoriserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletAuthoriser.Contract.ProjectWalletAuthoriserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.ProjectWalletAuthoriserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.ProjectWalletAuthoriserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletAuthoriser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProjectWalletAuthoriser.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserSession) Owner() (common.Address, error) {
	return _ProjectWalletAuthoriser.Contract.Owner(&_ProjectWalletAuthoriser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserCallerSession) Owner() (common.Address, error) {
	return _ProjectWalletAuthoriser.Contract.Owner(&_ProjectWalletAuthoriser.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.RenounceOwnership(&_ProjectWalletAuthoriser.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.RenounceOwnership(&_ProjectWalletAuthoriser.TransactOpts)
}

// SetAuthoriser is a paid mutator transaction binding the contract method 0xf7da0e04.
//
// Solidity: function setAuthoriser(_authoriser address) returns(bool)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactor) SetAuthoriser(opts *bind.TransactOpts, _authoriser common.Address) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.contract.Transact(opts, "setAuthoriser", _authoriser)
}

// SetAuthoriser is a paid mutator transaction binding the contract method 0xf7da0e04.
//
// Solidity: function setAuthoriser(_authoriser address) returns(bool)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserSession) SetAuthoriser(_authoriser common.Address) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.SetAuthoriser(&_ProjectWalletAuthoriser.TransactOpts, _authoriser)
}

// SetAuthoriser is a paid mutator transaction binding the contract method 0xf7da0e04.
//
// Solidity: function setAuthoriser(_authoriser address) returns(bool)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactorSession) SetAuthoriser(_authoriser common.Address) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.SetAuthoriser(&_ProjectWalletAuthoriser.TransactOpts, _authoriser)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_sender address, _receiver address, _amt uint256) returns(bool)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactor) Transfer(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.contract.Transact(opts, "transfer", _sender, _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_sender address, _receiver address, _amt uint256) returns(bool)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserSession) Transfer(_sender common.Address, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.Transfer(&_ProjectWalletAuthoriser.TransactOpts, _sender, _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_sender address, _receiver address, _amt uint256) returns(bool)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactorSession) Transfer(_sender common.Address, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.Transfer(&_ProjectWalletAuthoriser.TransactOpts, _sender, _receiver, _amt)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.TransferOwnership(&_ProjectWalletAuthoriser.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ProjectWalletAuthoriser.Contract.TransferOwnership(&_ProjectWalletAuthoriser.TransactOpts, _newOwner)
}

// ProjectWalletAuthoriserOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ProjectWalletAuthoriser contract.
type ProjectWalletAuthoriserOwnershipRenouncedIterator struct {
	Event *ProjectWalletAuthoriserOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ProjectWalletAuthoriserOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectWalletAuthoriserOwnershipRenounced)
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
		it.Event = new(ProjectWalletAuthoriserOwnershipRenounced)
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
func (it *ProjectWalletAuthoriserOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectWalletAuthoriserOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectWalletAuthoriserOwnershipRenounced represents a OwnershipRenounced event raised by the ProjectWalletAuthoriser contract.
type ProjectWalletAuthoriserOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ProjectWalletAuthoriserOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ProjectWalletAuthoriser.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletAuthoriserOwnershipRenouncedIterator{contract: _ProjectWalletAuthoriser.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ProjectWalletAuthoriserOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ProjectWalletAuthoriser.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectWalletAuthoriserOwnershipRenounced)
				if err := _ProjectWalletAuthoriser.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ProjectWalletAuthoriserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProjectWalletAuthoriser contract.
type ProjectWalletAuthoriserOwnershipTransferredIterator struct {
	Event *ProjectWalletAuthoriserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProjectWalletAuthoriserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectWalletAuthoriserOwnershipTransferred)
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
		it.Event = new(ProjectWalletAuthoriserOwnershipTransferred)
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
func (it *ProjectWalletAuthoriserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectWalletAuthoriserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectWalletAuthoriserOwnershipTransferred represents a OwnershipTransferred event raised by the ProjectWalletAuthoriser contract.
type ProjectWalletAuthoriserOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProjectWalletAuthoriserOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProjectWalletAuthoriser.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletAuthoriserOwnershipTransferredIterator{contract: _ProjectWalletAuthoriser.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProjectWalletAuthoriser *ProjectWalletAuthoriserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProjectWalletAuthoriserOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProjectWalletAuthoriser.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectWalletAuthoriserOwnershipTransferred)
				if err := _ProjectWalletAuthoriser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
