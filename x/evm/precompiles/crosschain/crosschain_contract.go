// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crosschain

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CrosschainMetaData contains all meta data concerning the Crosschain contract.
var CrosschainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txID\",\"type\":\"uint256\"}],\"name\":\"CancelSendToExternal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receipt\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"target\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"CrossChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"IncreaseBridgeFee\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"}],\"name\":\"bridgeCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_txID\",\"type\":\"uint256\"}],\"name\":\"cancelSendToExternal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_receipt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_receipt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"fip20CrossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_txID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"increaseBridgeFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610646806100206000396000f3fe60806040526004361061004a5760003560e01c80630b56c1901461004f578063160d7c731461008c57806329ec0ff1146100bc5780633c3e7d77146100f9578063c79a6b7b14610136575b600080fd5b34801561005b57600080fd5b506100766004803603810190610071919061035c565b610166565b6040516100839190610459565b60405180910390f35b6100a660048036038101906100a19190610297565b610172565b6040516100b39190610459565b60405180910390f35b3480156100c857600080fd5b506100e360048036038101906100de9190610257565b610182565b6040516100f09190610474565b60405180910390f35b34801561010557600080fd5b50610120600480360381019061011b9190610297565b61018a565b60405161012d9190610459565b60405180910390f35b610150600480360381019061014b91906103b8565b61019a565b60405161015d9190610459565b60405180910390f35b60006001905092915050565b6000600190509695505050505050565b600092915050565b6000600190509695505050505050565b600060019050949350505050565b60006101bb6101b6846104b4565b61048f565b9050828152602081018484840111156101d7576101d66105ab565b5b6101e2848285610537565b509392505050565b6000813590506101f9816105cb565b92915050565b60008135905061020e816105e2565b92915050565b600082601f830112610229576102286105a6565b5b81356102398482602086016101a8565b91505092915050565b600081359050610251816105f9565b92915050565b6000806040838503121561026e5761026d6105b5565b5b600061027c858286016101ea565b925050602061028d858286016101ff565b9150509250929050565b60008060008060008060c087890312156102b4576102b36105b5565b5b60006102c289828a016101ea565b965050602087013567ffffffffffffffff8111156102e3576102e26105b0565b5b6102ef89828a01610214565b955050604061030089828a01610242565b945050606061031189828a01610242565b935050608061032289828a016101ff565b92505060a087013567ffffffffffffffff811115610343576103426105b0565b5b61034f89828a01610214565b9150509295509295509295565b60008060408385031215610373576103726105b5565b5b600083013567ffffffffffffffff811115610391576103906105b0565b5b61039d85828601610214565b92505060206103ae85828601610242565b9150509250929050565b600080600080608085870312156103d2576103d16105b5565b5b600085013567ffffffffffffffff8111156103f0576103ef6105b0565b5b6103fc87828801610214565b945050602061040d87828801610242565b935050604061041e878288016101ea565b925050606061042f87828801610242565b91505092959194509250565b610444816104f7565b82525050565b6104538161052d565b82525050565b600060208201905061046e600083018461043b565b92915050565b6000602082019050610489600083018461044a565b92915050565b60006104996104aa565b90506104a58282610546565b919050565b6000604051905090565b600067ffffffffffffffff8211156104cf576104ce610577565b5b6104d8826105ba565b9050602081019050919050565b60006104f08261050d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b61054f826105ba565b810181811067ffffffffffffffff8211171561056e5761056d610577565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6105d4816104e5565b81146105df57600080fd5b50565b6105eb81610503565b81146105f657600080fd5b50565b6106028161052d565b811461060d57600080fd5b5056fea26469706673582212209de2d2402973c149a92e521bb4ef2561d26e4ad46d029b070c6b77792845546f64736f6c63430008060033",
}

// CrosschainABI is the input ABI used to generate the binding from.
// Deprecated: Use CrosschainMetaData.ABI instead.
var CrosschainABI = CrosschainMetaData.ABI

// CrosschainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrosschainMetaData.Bin instead.
var CrosschainBin = CrosschainMetaData.Bin

// DeployCrosschain deploys a new Ethereum contract, binding an instance of Crosschain to it.
func DeployCrosschain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Crosschain, error) {
	parsed, err := CrosschainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrosschainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Crosschain{CrosschainCaller: CrosschainCaller{contract: contract}, CrosschainTransactor: CrosschainTransactor{contract: contract}, CrosschainFilterer: CrosschainFilterer{contract: contract}}, nil
}

// Crosschain is an auto generated Go binding around an Ethereum contract.
type Crosschain struct {
	CrosschainCaller     // Read-only binding to the contract
	CrosschainTransactor // Write-only binding to the contract
	CrosschainFilterer   // Log filterer for contract events
}

// CrosschainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrosschainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrosschainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrosschainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrosschainSession struct {
	Contract     *Crosschain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrosschainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrosschainCallerSession struct {
	Contract *CrosschainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrosschainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrosschainTransactorSession struct {
	Contract     *CrosschainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrosschainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrosschainRaw struct {
	Contract *Crosschain // Generic contract binding to access the raw methods on
}

// CrosschainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrosschainCallerRaw struct {
	Contract *CrosschainCaller // Generic read-only contract binding to access the raw methods on
}

// CrosschainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrosschainTransactorRaw struct {
	Contract *CrosschainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrosschain creates a new instance of Crosschain, bound to a specific deployed contract.
func NewCrosschain(address common.Address, backend bind.ContractBackend) (*Crosschain, error) {
	contract, err := bindCrosschain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crosschain{CrosschainCaller: CrosschainCaller{contract: contract}, CrosschainTransactor: CrosschainTransactor{contract: contract}, CrosschainFilterer: CrosschainFilterer{contract: contract}}, nil
}

// NewCrosschainCaller creates a new read-only instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainCaller(address common.Address, caller bind.ContractCaller) (*CrosschainCaller, error) {
	contract, err := bindCrosschain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainCaller{contract: contract}, nil
}

// NewCrosschainTransactor creates a new write-only instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainTransactor(address common.Address, transactor bind.ContractTransactor) (*CrosschainTransactor, error) {
	contract, err := bindCrosschain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainTransactor{contract: contract}, nil
}

// NewCrosschainFilterer creates a new log filterer instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainFilterer(address common.Address, filterer bind.ContractFilterer) (*CrosschainFilterer, error) {
	contract, err := bindCrosschain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrosschainFilterer{contract: contract}, nil
}

// bindCrosschain binds a generic wrapper to an already deployed contract.
func bindCrosschain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrosschainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crosschain *CrosschainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crosschain.Contract.CrosschainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crosschain *CrosschainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.Contract.CrosschainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crosschain *CrosschainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crosschain.Contract.CrosschainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crosschain *CrosschainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crosschain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crosschain *CrosschainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crosschain *CrosschainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crosschain.Contract.contract.Transact(opts, method, params...)
}

// BridgeCoin is a free data retrieval call binding the contract method 0x29ec0ff1.
//
// Solidity: function bridgeCoin(address _token, bytes32 _target) view returns(uint256 _amount)
func (_Crosschain *CrosschainCaller) BridgeCoin(opts *bind.CallOpts, _token common.Address, _target [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "bridgeCoin", _token, _target)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BridgeCoin is a free data retrieval call binding the contract method 0x29ec0ff1.
//
// Solidity: function bridgeCoin(address _token, bytes32 _target) view returns(uint256 _amount)
func (_Crosschain *CrosschainSession) BridgeCoin(_token common.Address, _target [32]byte) (*big.Int, error) {
	return _Crosschain.Contract.BridgeCoin(&_Crosschain.CallOpts, _token, _target)
}

// BridgeCoin is a free data retrieval call binding the contract method 0x29ec0ff1.
//
// Solidity: function bridgeCoin(address _token, bytes32 _target) view returns(uint256 _amount)
func (_Crosschain *CrosschainCallerSession) BridgeCoin(_token common.Address, _target [32]byte) (*big.Int, error) {
	return _Crosschain.Contract.BridgeCoin(&_Crosschain.CallOpts, _token, _target)
}

// CancelSendToExternal is a paid mutator transaction binding the contract method 0x0b56c190.
//
// Solidity: function cancelSendToExternal(string _chain, uint256 _txID) returns(bool _result)
func (_Crosschain *CrosschainTransactor) CancelSendToExternal(opts *bind.TransactOpts, _chain string, _txID *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "cancelSendToExternal", _chain, _txID)
}

// CancelSendToExternal is a paid mutator transaction binding the contract method 0x0b56c190.
//
// Solidity: function cancelSendToExternal(string _chain, uint256 _txID) returns(bool _result)
func (_Crosschain *CrosschainSession) CancelSendToExternal(_chain string, _txID *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.CancelSendToExternal(&_Crosschain.TransactOpts, _chain, _txID)
}

// CancelSendToExternal is a paid mutator transaction binding the contract method 0x0b56c190.
//
// Solidity: function cancelSendToExternal(string _chain, uint256 _txID) returns(bool _result)
func (_Crosschain *CrosschainTransactorSession) CancelSendToExternal(_chain string, _txID *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.CancelSendToExternal(&_Crosschain.TransactOpts, _chain, _txID)
}

// CrossChain is a paid mutator transaction binding the contract method 0x160d7c73.
//
// Solidity: function crossChain(address _token, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) payable returns(bool _result)
func (_Crosschain *CrosschainTransactor) CrossChain(opts *bind.TransactOpts, _token common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "crossChain", _token, _receipt, _amount, _fee, _target, _memo)
}

// CrossChain is a paid mutator transaction binding the contract method 0x160d7c73.
//
// Solidity: function crossChain(address _token, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) payable returns(bool _result)
func (_Crosschain *CrosschainSession) CrossChain(_token common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _Crosschain.Contract.CrossChain(&_Crosschain.TransactOpts, _token, _receipt, _amount, _fee, _target, _memo)
}

// CrossChain is a paid mutator transaction binding the contract method 0x160d7c73.
//
// Solidity: function crossChain(address _token, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) payable returns(bool _result)
func (_Crosschain *CrosschainTransactorSession) CrossChain(_token common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _Crosschain.Contract.CrossChain(&_Crosschain.TransactOpts, _token, _receipt, _amount, _fee, _target, _memo)
}

// Fip20CrossChain is a paid mutator transaction binding the contract method 0x3c3e7d77.
//
// Solidity: function fip20CrossChain(address _sender, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) returns(bool _result)
func (_Crosschain *CrosschainTransactor) Fip20CrossChain(opts *bind.TransactOpts, _sender common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "fip20CrossChain", _sender, _receipt, _amount, _fee, _target, _memo)
}

// Fip20CrossChain is a paid mutator transaction binding the contract method 0x3c3e7d77.
//
// Solidity: function fip20CrossChain(address _sender, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) returns(bool _result)
func (_Crosschain *CrosschainSession) Fip20CrossChain(_sender common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _Crosschain.Contract.Fip20CrossChain(&_Crosschain.TransactOpts, _sender, _receipt, _amount, _fee, _target, _memo)
}

// Fip20CrossChain is a paid mutator transaction binding the contract method 0x3c3e7d77.
//
// Solidity: function fip20CrossChain(address _sender, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) returns(bool _result)
func (_Crosschain *CrosschainTransactorSession) Fip20CrossChain(_sender common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _Crosschain.Contract.Fip20CrossChain(&_Crosschain.TransactOpts, _sender, _receipt, _amount, _fee, _target, _memo)
}

// IncreaseBridgeFee is a paid mutator transaction binding the contract method 0xc79a6b7b.
//
// Solidity: function increaseBridgeFee(string _chain, uint256 _txID, address _token, uint256 _fee) payable returns(bool _result)
func (_Crosschain *CrosschainTransactor) IncreaseBridgeFee(opts *bind.TransactOpts, _chain string, _txID *big.Int, _token common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "increaseBridgeFee", _chain, _txID, _token, _fee)
}

// IncreaseBridgeFee is a paid mutator transaction binding the contract method 0xc79a6b7b.
//
// Solidity: function increaseBridgeFee(string _chain, uint256 _txID, address _token, uint256 _fee) payable returns(bool _result)
func (_Crosschain *CrosschainSession) IncreaseBridgeFee(_chain string, _txID *big.Int, _token common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.IncreaseBridgeFee(&_Crosschain.TransactOpts, _chain, _txID, _token, _fee)
}

// IncreaseBridgeFee is a paid mutator transaction binding the contract method 0xc79a6b7b.
//
// Solidity: function increaseBridgeFee(string _chain, uint256 _txID, address _token, uint256 _fee) payable returns(bool _result)
func (_Crosschain *CrosschainTransactorSession) IncreaseBridgeFee(_chain string, _txID *big.Int, _token common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.IncreaseBridgeFee(&_Crosschain.TransactOpts, _chain, _txID, _token, _fee)
}

// CrosschainCancelSendToExternalIterator is returned from FilterCancelSendToExternal and is used to iterate over the raw logs and unpacked data for CancelSendToExternal events raised by the Crosschain contract.
type CrosschainCancelSendToExternalIterator struct {
	Event *CrosschainCancelSendToExternal // Event containing the contract specifics and raw log

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
func (it *CrosschainCancelSendToExternalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainCancelSendToExternal)
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
		it.Event = new(CrosschainCancelSendToExternal)
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
func (it *CrosschainCancelSendToExternalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainCancelSendToExternalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainCancelSendToExternal represents a CancelSendToExternal event raised by the Crosschain contract.
type CrosschainCancelSendToExternal struct {
	Sender common.Address
	Chain  string
	TxID   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelSendToExternal is a free log retrieval operation binding the contract event 0xe2ae965fb5b8e4c7da962424292951c18e0e9c1905b87c78cf0186fa70382535.
//
// Solidity: event CancelSendToExternal(address indexed sender, string chain, uint256 txID)
func (_Crosschain *CrosschainFilterer) FilterCancelSendToExternal(opts *bind.FilterOpts, sender []common.Address) (*CrosschainCancelSendToExternalIterator, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "CancelSendToExternal", senderRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainCancelSendToExternalIterator{contract: _Crosschain.contract, event: "CancelSendToExternal", logs: logs, sub: sub}, nil
}

// WatchCancelSendToExternal is a free log subscription operation binding the contract event 0xe2ae965fb5b8e4c7da962424292951c18e0e9c1905b87c78cf0186fa70382535.
//
// Solidity: event CancelSendToExternal(address indexed sender, string chain, uint256 txID)
func (_Crosschain *CrosschainFilterer) WatchCancelSendToExternal(opts *bind.WatchOpts, sink chan<- *CrosschainCancelSendToExternal, sender []common.Address) (event.Subscription, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "CancelSendToExternal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainCancelSendToExternal)
				if err := _Crosschain.contract.UnpackLog(event, "CancelSendToExternal", log); err != nil {
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

// ParseCancelSendToExternal is a log parse operation binding the contract event 0xe2ae965fb5b8e4c7da962424292951c18e0e9c1905b87c78cf0186fa70382535.
//
// Solidity: event CancelSendToExternal(address indexed sender, string chain, uint256 txID)
func (_Crosschain *CrosschainFilterer) ParseCancelSendToExternal(log types.Log) (*CrosschainCancelSendToExternal, error) {
	event := new(CrosschainCancelSendToExternal)
	if err := _Crosschain.contract.UnpackLog(event, "CancelSendToExternal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainCrossChainIterator is returned from FilterCrossChain and is used to iterate over the raw logs and unpacked data for CrossChain events raised by the Crosschain contract.
type CrosschainCrossChainIterator struct {
	Event *CrosschainCrossChain // Event containing the contract specifics and raw log

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
func (it *CrosschainCrossChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainCrossChain)
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
		it.Event = new(CrosschainCrossChain)
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
func (it *CrosschainCrossChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainCrossChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainCrossChain represents a CrossChain event raised by the Crosschain contract.
type CrosschainCrossChain struct {
	Sender  common.Address
	Token   common.Address
	Denom   string
	Receipt string
	Amount  *big.Int
	Fee     *big.Int
	Target  [32]byte
	Memo    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossChain is a free log retrieval operation binding the contract event 0xb783df819ac99ca709650d67d9237a00b553c6ef941dceabeed6f4bc990d31ba.
//
// Solidity: event CrossChain(address indexed sender, address indexed token, string denom, string receipt, uint256 amount, uint256 fee, bytes32 target, string memo)
func (_Crosschain *CrosschainFilterer) FilterCrossChain(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*CrosschainCrossChainIterator, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "CrossChain", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainCrossChainIterator{contract: _Crosschain.contract, event: "CrossChain", logs: logs, sub: sub}, nil
}

// WatchCrossChain is a free log subscription operation binding the contract event 0xb783df819ac99ca709650d67d9237a00b553c6ef941dceabeed6f4bc990d31ba.
//
// Solidity: event CrossChain(address indexed sender, address indexed token, string denom, string receipt, uint256 amount, uint256 fee, bytes32 target, string memo)
func (_Crosschain *CrosschainFilterer) WatchCrossChain(opts *bind.WatchOpts, sink chan<- *CrosschainCrossChain, sender []common.Address, token []common.Address) (event.Subscription, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "CrossChain", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainCrossChain)
				if err := _Crosschain.contract.UnpackLog(event, "CrossChain", log); err != nil {
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

// ParseCrossChain is a log parse operation binding the contract event 0xb783df819ac99ca709650d67d9237a00b553c6ef941dceabeed6f4bc990d31ba.
//
// Solidity: event CrossChain(address indexed sender, address indexed token, string denom, string receipt, uint256 amount, uint256 fee, bytes32 target, string memo)
func (_Crosschain *CrosschainFilterer) ParseCrossChain(log types.Log) (*CrosschainCrossChain, error) {
	event := new(CrosschainCrossChain)
	if err := _Crosschain.contract.UnpackLog(event, "CrossChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainIncreaseBridgeFeeIterator is returned from FilterIncreaseBridgeFee and is used to iterate over the raw logs and unpacked data for IncreaseBridgeFee events raised by the Crosschain contract.
type CrosschainIncreaseBridgeFeeIterator struct {
	Event *CrosschainIncreaseBridgeFee // Event containing the contract specifics and raw log

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
func (it *CrosschainIncreaseBridgeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainIncreaseBridgeFee)
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
		it.Event = new(CrosschainIncreaseBridgeFee)
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
func (it *CrosschainIncreaseBridgeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainIncreaseBridgeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainIncreaseBridgeFee represents a IncreaseBridgeFee event raised by the Crosschain contract.
type CrosschainIncreaseBridgeFee struct {
	Sender common.Address
	Token  common.Address
	Chain  string
	TxID   *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBridgeFee is a free log retrieval operation binding the contract event 0x4b4d0e64eb77c0f61892107908295f09b3e381c50c655f4a73a4ad61c07350a0.
//
// Solidity: event IncreaseBridgeFee(address indexed sender, address indexed token, string chain, uint256 txID, uint256 fee)
func (_Crosschain *CrosschainFilterer) FilterIncreaseBridgeFee(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*CrosschainIncreaseBridgeFeeIterator, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "IncreaseBridgeFee", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainIncreaseBridgeFeeIterator{contract: _Crosschain.contract, event: "IncreaseBridgeFee", logs: logs, sub: sub}, nil
}

// WatchIncreaseBridgeFee is a free log subscription operation binding the contract event 0x4b4d0e64eb77c0f61892107908295f09b3e381c50c655f4a73a4ad61c07350a0.
//
// Solidity: event IncreaseBridgeFee(address indexed sender, address indexed token, string chain, uint256 txID, uint256 fee)
func (_Crosschain *CrosschainFilterer) WatchIncreaseBridgeFee(opts *bind.WatchOpts, sink chan<- *CrosschainIncreaseBridgeFee, sender []common.Address, token []common.Address) (event.Subscription, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "IncreaseBridgeFee", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainIncreaseBridgeFee)
				if err := _Crosschain.contract.UnpackLog(event, "IncreaseBridgeFee", log); err != nil {
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

// ParseIncreaseBridgeFee is a log parse operation binding the contract event 0x4b4d0e64eb77c0f61892107908295f09b3e381c50c655f4a73a4ad61c07350a0.
//
// Solidity: event IncreaseBridgeFee(address indexed sender, address indexed token, string chain, uint256 txID, uint256 fee)
func (_Crosschain *CrosschainFilterer) ParseIncreaseBridgeFee(log types.Log) (*CrosschainIncreaseBridgeFee, error) {
	event := new(CrosschainIncreaseBridgeFee)
	if err := _Crosschain.contract.UnpackLog(event, "IncreaseBridgeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
