// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crosschain_test

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

// CrosschainTestMetaData contains all meta data concerning the CrosschainTest contract.
var CrosschainTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txID\",\"type\":\"uint256\"}],\"name\":\"CancelSendToExternal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receipt\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"target\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"CrossChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"IncreaseBridgeFee\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"}],\"name\":\"bridgeCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_txID\",\"type\":\"uint256\"}],\"name\":\"cancelSendToExternal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_receipt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_receipt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"fip20CrossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_txID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"increaseBridgeFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061167a806100206000396000f3fe60806040526004361061004a5760003560e01c80630b56c1901461004f578063160d7c731461008c57806329ec0ff1146100bc5780633c3e7d77146100f9578063c79a6b7b14610136575b600080fd5b34801561005b57600080fd5b5061007660048036038101906100719190610e75565b610166565b6040516100839190611228565b60405180910390f35b6100a660048036038101906100a19190610d3a565b61017a565b6040516100b39190611228565b60405180910390f35b3480156100c857600080fd5b506100e360048036038101906100de9190610cfa565b610300565b6040516100f09190611321565b60405180910390f35b34801561010557600080fd5b50610120600480360381019061011b9190610d3a565b610314565b60405161012d9190611228565b60405180910390f35b610150600480360381019061014b9190610ed1565b610324565b60405161015d9190611228565b60405180910390f35b6000610172838361033c565b905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16146102e6578673ffffffffffffffffffffffffffffffffffffffff166323b872dd333087896101d991906113cf565b6040518463ffffffff1660e01b81526004016101f793929190611130565b602060405180830381600087803b15801561021157600080fd5b505af1158015610225573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102499190610dff565b508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3611004868861027591906113cf565b6040518363ffffffff1660e01b81526004016102929291906111ff565b602060405180830381600087803b1580156102ac57600080fd5b505af11580156102c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e49190610dff565b505b6102f487878787878761040b565b90509695505050505050565b600061030c8383610648565b905092915050565b6000600190509695505050505050565b600061033285858585610715565b9050949350505050565b600080600061100473ffffffffffffffffffffffffffffffffffffffff1661036486866107e8565b60405161037191906110c1565b6000604051808303816000865af19150503d80600081146103ae576040519150601f19603f3d011682016040523d82523d6000602084013e6103b3565b606091505b50915091506103f882826040518060400160405280601e81526020017f63616e63656c2073656e6420746f2065787465726e616c206661696c65640000815250610882565b61040181610949565b9250505092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16146105245760008773ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e306110046040518363ffffffff1660e01b815260040161047f929190611107565b60206040518083038186803b15801561049757600080fd5b505afa1580156104ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cf9190610f54565b905084866104dd91906113cf565b811461051e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610515906112e1565b60405180910390fd5b50610572565b838561053091906113cf565b3414610571576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056890611301565b60405180910390fd5b5b60008061100473ffffffffffffffffffffffffffffffffffffffff163461059d8b8b8b8b8b8b61096b565b6040516105aa91906110c1565b60006040518083038185875af1925050503d80600081146105e7576040519150601f19603f3d011682016040523d82523d6000602084013e6105ec565b606091505b509150915061063182826040518060400160405280601281526020017f63726f73732d636861696e206661696c65640000000000000000000000000000815250610882565b61063a81610a11565b925050509695505050505050565b600080600061100473ffffffffffffffffffffffffffffffffffffffff166106708686610a33565b60405161067d91906110c1565b600060405180830381855afa9150503d80600081146106b8576040519150601f19603f3d011682016040523d82523d6000602084013e6106bd565b606091505b509150915061070282826040518060400160405280601281526020017f62726964676520636f696e206661696c65640000000000000000000000000000815250610882565b61070b81610acd565b9250505092915050565b600080600061100473ffffffffffffffffffffffffffffffffffffffff1661073f88888888610aef565b60405161074c91906110c1565b6000604051808303816000865af19150503d8060008114610789576040519150601f19603f3d011682016040523d82523d6000602084013e61078e565b606091505b50915091506107d382826040518060400160405280601a81526020017f696e6372656173652062726964676520666565206661696c6564000000000000815250610882565b6107dc81610b8f565b92505050949350505050565b606082826040516024016107fd929190611265565b6040516020818303038152906040527feeb3593d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905092915050565b826109445760008280602001905181019061089d9190610e2c565b90506001825110156108e657806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dd9190611243565b60405180910390fd5b81816040516020016108f99291906110d8565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093b9190611243565b60405180910390fd5b505050565b600080828060200190518101906109609190610dff565b905080915050919050565b606086868686868660405160240161098896959493929190611190565b6040516020818303038152906040527f160d7c73000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090509695505050505050565b60008082806020019051810190610a289190610dff565b905080915050919050565b60608282604051602401610a48929190611167565b6040516020818303038152906040527f29ec0ff1000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905092915050565b60008082806020019051810190610ae49190610f54565b905080915050919050565b606084848484604051602401610b089493929190611295565b6040516020818303038152906040527f9b45009d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050949350505050565b60008082806020019051810190610ba69190610dff565b905080915050919050565b6000610bc4610bbf84611361565b61133c565b905082815260208101848484011115610be057610bdf61154d565b5b610beb848285611477565b509392505050565b6000610c06610c0184611361565b61133c565b905082815260208101848484011115610c2257610c2161154d565b5b610c2d848285611486565b509392505050565b600081359050610c44816115e8565b92915050565b600081519050610c59816115ff565b92915050565b600081359050610c6e81611616565b92915050565b600082601f830112610c8957610c88611548565b5b8135610c99848260208601610bb1565b91505092915050565b600082601f830112610cb757610cb6611548565b5b8151610cc7848260208601610bf3565b91505092915050565b600081359050610cdf8161162d565b92915050565b600081519050610cf48161162d565b92915050565b60008060408385031215610d1157610d10611557565b5b6000610d1f85828601610c35565b9250506020610d3085828601610c5f565b9150509250929050565b60008060008060008060c08789031215610d5757610d56611557565b5b6000610d6589828a01610c35565b965050602087013567ffffffffffffffff811115610d8657610d85611552565b5b610d9289828a01610c74565b9550506040610da389828a01610cd0565b9450506060610db489828a01610cd0565b9350506080610dc589828a01610c5f565b92505060a087013567ffffffffffffffff811115610de657610de5611552565b5b610df289828a01610c74565b9150509295509295509295565b600060208284031215610e1557610e14611557565b5b6000610e2384828501610c4a565b91505092915050565b600060208284031215610e4257610e41611557565b5b600082015167ffffffffffffffff811115610e6057610e5f611552565b5b610e6c84828501610ca2565b91505092915050565b60008060408385031215610e8c57610e8b611557565b5b600083013567ffffffffffffffff811115610eaa57610ea9611552565b5b610eb685828601610c74565b9250506020610ec785828601610cd0565b9150509250929050565b60008060008060808587031215610eeb57610eea611557565b5b600085013567ffffffffffffffff811115610f0957610f08611552565b5b610f1587828801610c74565b9450506020610f2687828801610cd0565b9350506040610f3787828801610c35565b9250506060610f4887828801610cd0565b91505092959194509250565b600060208284031215610f6a57610f69611557565b5b6000610f7884828501610ce5565b91505092915050565b610f8a81611425565b82525050565b610f9981611437565b82525050565b610fa881611443565b82525050565b6000610fb982611392565b610fc381856113a8565b9350610fd3818560208601611486565b80840191505092915050565b6000610fea8261139d565b610ff481856113b3565b9350611004818560208601611486565b61100d8161155c565b840191505092915050565b60006110238261139d565b61102d81856113c4565b935061103d818560208601611486565b80840191505092915050565b60006110566020836113b3565b91506110618261156d565b602082019050919050565b60006110796002836113c4565b915061108482611596565b600282019050919050565b600061109c6020836113b3565b91506110a7826115bf565b602082019050919050565b6110bb8161146d565b82525050565b60006110cd8284610fae565b915081905092915050565b60006110e48285611018565b91506110ef8261106c565b91506110fb8284611018565b91508190509392505050565b600060408201905061111c6000830185610f81565b6111296020830184610f81565b9392505050565b60006060820190506111456000830186610f81565b6111526020830185610f81565b61115f60408301846110b2565b949350505050565b600060408201905061117c6000830185610f81565b6111896020830184610f9f565b9392505050565b600060c0820190506111a56000830189610f81565b81810360208301526111b78188610fdf565b90506111c660408301876110b2565b6111d360608301866110b2565b6111e06080830185610f9f565b81810360a08301526111f28184610fdf565b9050979650505050505050565b60006040820190506112146000830185610f81565b61122160208301846110b2565b9392505050565b600060208201905061123d6000830184610f90565b92915050565b6000602082019050818103600083015261125d8184610fdf565b905092915050565b6000604082019050818103600083015261127f8185610fdf565b905061128e60208301846110b2565b9392505050565b600060808201905081810360008301526112af8187610fdf565b90506112be60208301866110b2565b6112cb6040830185610f81565b6112d860608301846110b2565b95945050505050565b600060208201905081810360008301526112fa81611049565b9050919050565b6000602082019050818103600083015261131a8161108f565b9050919050565b600060208201905061133660008301846110b2565b92915050565b6000611346611357565b905061135282826114b9565b919050565b6000604051905090565b600067ffffffffffffffff82111561137c5761137b611519565b5b6113858261155c565b9050602081019050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b60006113da8261146d565b91506113e58361146d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561141a576114196114ea565b5b828201905092915050565b60006114308261144d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156114a4578082015181840152602081019050611489565b838111156114b3576000848401525b50505050565b6114c28261155c565b810181811067ffffffffffffffff821117156114e1576114e0611519565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f616c6c6f77616e6365206e6f7420657175616c20616d6f756e74202b20666565600082015250565b7f3a20000000000000000000000000000000000000000000000000000000000000600082015250565b7f6d73672e76616c7565206e6f7420657175616c20616d6f756e74202b20666565600082015250565b6115f181611425565b81146115fc57600080fd5b50565b61160881611437565b811461161357600080fd5b50565b61161f81611443565b811461162a57600080fd5b50565b6116368161146d565b811461164157600080fd5b5056fea26469706673582212201215a768e37e47ee4a92be1defd890c53990a120d6ae3759c0d2eaf5792ac95c64736f6c63430008060033",
}

// CrosschainTestABI is the input ABI used to generate the binding from.
// Deprecated: Use CrosschainTestMetaData.ABI instead.
var CrosschainTestABI = CrosschainTestMetaData.ABI

// CrosschainTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrosschainTestMetaData.Bin instead.
var CrosschainTestBin = CrosschainTestMetaData.Bin

// DeployCrosschainTest deploys a new Ethereum contract, binding an instance of CrosschainTest to it.
func DeployCrosschainTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrosschainTest, error) {
	parsed, err := CrosschainTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrosschainTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrosschainTest{CrosschainTestCaller: CrosschainTestCaller{contract: contract}, CrosschainTestTransactor: CrosschainTestTransactor{contract: contract}, CrosschainTestFilterer: CrosschainTestFilterer{contract: contract}}, nil
}

// CrosschainTest is an auto generated Go binding around an Ethereum contract.
type CrosschainTest struct {
	CrosschainTestCaller     // Read-only binding to the contract
	CrosschainTestTransactor // Write-only binding to the contract
	CrosschainTestFilterer   // Log filterer for contract events
}

// CrosschainTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrosschainTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrosschainTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrosschainTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrosschainTestSession struct {
	Contract     *CrosschainTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrosschainTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrosschainTestCallerSession struct {
	Contract *CrosschainTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CrosschainTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrosschainTestTransactorSession struct {
	Contract     *CrosschainTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CrosschainTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrosschainTestRaw struct {
	Contract *CrosschainTest // Generic contract binding to access the raw methods on
}

// CrosschainTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrosschainTestCallerRaw struct {
	Contract *CrosschainTestCaller // Generic read-only contract binding to access the raw methods on
}

// CrosschainTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrosschainTestTransactorRaw struct {
	Contract *CrosschainTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrosschainTest creates a new instance of CrosschainTest, bound to a specific deployed contract.
func NewCrosschainTest(address common.Address, backend bind.ContractBackend) (*CrosschainTest, error) {
	contract, err := bindCrosschainTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrosschainTest{CrosschainTestCaller: CrosschainTestCaller{contract: contract}, CrosschainTestTransactor: CrosschainTestTransactor{contract: contract}, CrosschainTestFilterer: CrosschainTestFilterer{contract: contract}}, nil
}

// NewCrosschainTestCaller creates a new read-only instance of CrosschainTest, bound to a specific deployed contract.
func NewCrosschainTestCaller(address common.Address, caller bind.ContractCaller) (*CrosschainTestCaller, error) {
	contract, err := bindCrosschainTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainTestCaller{contract: contract}, nil
}

// NewCrosschainTestTransactor creates a new write-only instance of CrosschainTest, bound to a specific deployed contract.
func NewCrosschainTestTransactor(address common.Address, transactor bind.ContractTransactor) (*CrosschainTestTransactor, error) {
	contract, err := bindCrosschainTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainTestTransactor{contract: contract}, nil
}

// NewCrosschainTestFilterer creates a new log filterer instance of CrosschainTest, bound to a specific deployed contract.
func NewCrosschainTestFilterer(address common.Address, filterer bind.ContractFilterer) (*CrosschainTestFilterer, error) {
	contract, err := bindCrosschainTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrosschainTestFilterer{contract: contract}, nil
}

// bindCrosschainTest binds a generic wrapper to an already deployed contract.
func bindCrosschainTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrosschainTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrosschainTest *CrosschainTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrosschainTest.Contract.CrosschainTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrosschainTest *CrosschainTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrosschainTest.Contract.CrosschainTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrosschainTest *CrosschainTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrosschainTest.Contract.CrosschainTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrosschainTest *CrosschainTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrosschainTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrosschainTest *CrosschainTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrosschainTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrosschainTest *CrosschainTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrosschainTest.Contract.contract.Transact(opts, method, params...)
}

// BridgeCoin is a free data retrieval call binding the contract method 0x29ec0ff1.
//
// Solidity: function bridgeCoin(address _token, bytes32 _target) view returns(uint256)
func (_CrosschainTest *CrosschainTestCaller) BridgeCoin(opts *bind.CallOpts, _token common.Address, _target [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CrosschainTest.contract.Call(opts, &out, "bridgeCoin", _token, _target)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BridgeCoin is a free data retrieval call binding the contract method 0x29ec0ff1.
//
// Solidity: function bridgeCoin(address _token, bytes32 _target) view returns(uint256)
func (_CrosschainTest *CrosschainTestSession) BridgeCoin(_token common.Address, _target [32]byte) (*big.Int, error) {
	return _CrosschainTest.Contract.BridgeCoin(&_CrosschainTest.CallOpts, _token, _target)
}

// BridgeCoin is a free data retrieval call binding the contract method 0x29ec0ff1.
//
// Solidity: function bridgeCoin(address _token, bytes32 _target) view returns(uint256)
func (_CrosschainTest *CrosschainTestCallerSession) BridgeCoin(_token common.Address, _target [32]byte) (*big.Int, error) {
	return _CrosschainTest.Contract.BridgeCoin(&_CrosschainTest.CallOpts, _token, _target)
}

// CancelSendToExternal is a paid mutator transaction binding the contract method 0x0b56c190.
//
// Solidity: function cancelSendToExternal(string _chain, uint256 _txID) returns(bool)
func (_CrosschainTest *CrosschainTestTransactor) CancelSendToExternal(opts *bind.TransactOpts, _chain string, _txID *big.Int) (*types.Transaction, error) {
	return _CrosschainTest.contract.Transact(opts, "cancelSendToExternal", _chain, _txID)
}

// CancelSendToExternal is a paid mutator transaction binding the contract method 0x0b56c190.
//
// Solidity: function cancelSendToExternal(string _chain, uint256 _txID) returns(bool)
func (_CrosschainTest *CrosschainTestSession) CancelSendToExternal(_chain string, _txID *big.Int) (*types.Transaction, error) {
	return _CrosschainTest.Contract.CancelSendToExternal(&_CrosschainTest.TransactOpts, _chain, _txID)
}

// CancelSendToExternal is a paid mutator transaction binding the contract method 0x0b56c190.
//
// Solidity: function cancelSendToExternal(string _chain, uint256 _txID) returns(bool)
func (_CrosschainTest *CrosschainTestTransactorSession) CancelSendToExternal(_chain string, _txID *big.Int) (*types.Transaction, error) {
	return _CrosschainTest.Contract.CancelSendToExternal(&_CrosschainTest.TransactOpts, _chain, _txID)
}

// CrossChain is a paid mutator transaction binding the contract method 0x160d7c73.
//
// Solidity: function crossChain(address _token, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) payable returns(bool)
func (_CrosschainTest *CrosschainTestTransactor) CrossChain(opts *bind.TransactOpts, _token common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _CrosschainTest.contract.Transact(opts, "crossChain", _token, _receipt, _amount, _fee, _target, _memo)
}

// CrossChain is a paid mutator transaction binding the contract method 0x160d7c73.
//
// Solidity: function crossChain(address _token, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) payable returns(bool)
func (_CrosschainTest *CrosschainTestSession) CrossChain(_token common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _CrosschainTest.Contract.CrossChain(&_CrosschainTest.TransactOpts, _token, _receipt, _amount, _fee, _target, _memo)
}

// CrossChain is a paid mutator transaction binding the contract method 0x160d7c73.
//
// Solidity: function crossChain(address _token, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) payable returns(bool)
func (_CrosschainTest *CrosschainTestTransactorSession) CrossChain(_token common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _CrosschainTest.Contract.CrossChain(&_CrosschainTest.TransactOpts, _token, _receipt, _amount, _fee, _target, _memo)
}

// Fip20CrossChain is a paid mutator transaction binding the contract method 0x3c3e7d77.
//
// Solidity: function fip20CrossChain(address _sender, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) returns(bool _result)
func (_CrosschainTest *CrosschainTestTransactor) Fip20CrossChain(opts *bind.TransactOpts, _sender common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _CrosschainTest.contract.Transact(opts, "fip20CrossChain", _sender, _receipt, _amount, _fee, _target, _memo)
}

// Fip20CrossChain is a paid mutator transaction binding the contract method 0x3c3e7d77.
//
// Solidity: function fip20CrossChain(address _sender, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) returns(bool _result)
func (_CrosschainTest *CrosschainTestSession) Fip20CrossChain(_sender common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _CrosschainTest.Contract.Fip20CrossChain(&_CrosschainTest.TransactOpts, _sender, _receipt, _amount, _fee, _target, _memo)
}

// Fip20CrossChain is a paid mutator transaction binding the contract method 0x3c3e7d77.
//
// Solidity: function fip20CrossChain(address _sender, string _receipt, uint256 _amount, uint256 _fee, bytes32 _target, string _memo) returns(bool _result)
func (_CrosschainTest *CrosschainTestTransactorSession) Fip20CrossChain(_sender common.Address, _receipt string, _amount *big.Int, _fee *big.Int, _target [32]byte, _memo string) (*types.Transaction, error) {
	return _CrosschainTest.Contract.Fip20CrossChain(&_CrosschainTest.TransactOpts, _sender, _receipt, _amount, _fee, _target, _memo)
}

// IncreaseBridgeFee is a paid mutator transaction binding the contract method 0xc79a6b7b.
//
// Solidity: function increaseBridgeFee(string _chain, uint256 _txID, address _token, uint256 _fee) payable returns(bool)
func (_CrosschainTest *CrosschainTestTransactor) IncreaseBridgeFee(opts *bind.TransactOpts, _chain string, _txID *big.Int, _token common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _CrosschainTest.contract.Transact(opts, "increaseBridgeFee", _chain, _txID, _token, _fee)
}

// IncreaseBridgeFee is a paid mutator transaction binding the contract method 0xc79a6b7b.
//
// Solidity: function increaseBridgeFee(string _chain, uint256 _txID, address _token, uint256 _fee) payable returns(bool)
func (_CrosschainTest *CrosschainTestSession) IncreaseBridgeFee(_chain string, _txID *big.Int, _token common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _CrosschainTest.Contract.IncreaseBridgeFee(&_CrosschainTest.TransactOpts, _chain, _txID, _token, _fee)
}

// IncreaseBridgeFee is a paid mutator transaction binding the contract method 0xc79a6b7b.
//
// Solidity: function increaseBridgeFee(string _chain, uint256 _txID, address _token, uint256 _fee) payable returns(bool)
func (_CrosschainTest *CrosschainTestTransactorSession) IncreaseBridgeFee(_chain string, _txID *big.Int, _token common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _CrosschainTest.Contract.IncreaseBridgeFee(&_CrosschainTest.TransactOpts, _chain, _txID, _token, _fee)
}

// CrosschainTestCancelSendToExternalIterator is returned from FilterCancelSendToExternal and is used to iterate over the raw logs and unpacked data for CancelSendToExternal events raised by the CrosschainTest contract.
type CrosschainTestCancelSendToExternalIterator struct {
	Event *CrosschainTestCancelSendToExternal // Event containing the contract specifics and raw log

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
func (it *CrosschainTestCancelSendToExternalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainTestCancelSendToExternal)
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
		it.Event = new(CrosschainTestCancelSendToExternal)
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
func (it *CrosschainTestCancelSendToExternalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainTestCancelSendToExternalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainTestCancelSendToExternal represents a CancelSendToExternal event raised by the CrosschainTest contract.
type CrosschainTestCancelSendToExternal struct {
	Sender common.Address
	Chain  string
	TxID   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelSendToExternal is a free log retrieval operation binding the contract event 0xe2ae965fb5b8e4c7da962424292951c18e0e9c1905b87c78cf0186fa70382535.
//
// Solidity: event CancelSendToExternal(address indexed sender, string chain, uint256 txID)
func (_CrosschainTest *CrosschainTestFilterer) FilterCancelSendToExternal(opts *bind.FilterOpts, sender []common.Address) (*CrosschainTestCancelSendToExternalIterator, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrosschainTest.contract.FilterLogs(opts, "CancelSendToExternal", senderRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainTestCancelSendToExternalIterator{contract: _CrosschainTest.contract, event: "CancelSendToExternal", logs: logs, sub: sub}, nil
}

// WatchCancelSendToExternal is a free log subscription operation binding the contract event 0xe2ae965fb5b8e4c7da962424292951c18e0e9c1905b87c78cf0186fa70382535.
//
// Solidity: event CancelSendToExternal(address indexed sender, string chain, uint256 txID)
func (_CrosschainTest *CrosschainTestFilterer) WatchCancelSendToExternal(opts *bind.WatchOpts, sink chan<- *CrosschainTestCancelSendToExternal, sender []common.Address) (event.Subscription, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrosschainTest.contract.WatchLogs(opts, "CancelSendToExternal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainTestCancelSendToExternal)
				if err := _CrosschainTest.contract.UnpackLog(event, "CancelSendToExternal", log); err != nil {
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
func (_CrosschainTest *CrosschainTestFilterer) ParseCancelSendToExternal(log types.Log) (*CrosschainTestCancelSendToExternal, error) {
	event := new(CrosschainTestCancelSendToExternal)
	if err := _CrosschainTest.contract.UnpackLog(event, "CancelSendToExternal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainTestCrossChainIterator is returned from FilterCrossChain and is used to iterate over the raw logs and unpacked data for CrossChain events raised by the CrosschainTest contract.
type CrosschainTestCrossChainIterator struct {
	Event *CrosschainTestCrossChain // Event containing the contract specifics and raw log

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
func (it *CrosschainTestCrossChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainTestCrossChain)
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
		it.Event = new(CrosschainTestCrossChain)
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
func (it *CrosschainTestCrossChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainTestCrossChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainTestCrossChain represents a CrossChain event raised by the CrosschainTest contract.
type CrosschainTestCrossChain struct {
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
func (_CrosschainTest *CrosschainTestFilterer) FilterCrossChain(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*CrosschainTestCrossChainIterator, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrosschainTest.contract.FilterLogs(opts, "CrossChain", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainTestCrossChainIterator{contract: _CrosschainTest.contract, event: "CrossChain", logs: logs, sub: sub}, nil
}

// WatchCrossChain is a free log subscription operation binding the contract event 0xb783df819ac99ca709650d67d9237a00b553c6ef941dceabeed6f4bc990d31ba.
//
// Solidity: event CrossChain(address indexed sender, address indexed token, string denom, string receipt, uint256 amount, uint256 fee, bytes32 target, string memo)
func (_CrosschainTest *CrosschainTestFilterer) WatchCrossChain(opts *bind.WatchOpts, sink chan<- *CrosschainTestCrossChain, sender []common.Address, token []common.Address) (event.Subscription, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrosschainTest.contract.WatchLogs(opts, "CrossChain", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainTestCrossChain)
				if err := _CrosschainTest.contract.UnpackLog(event, "CrossChain", log); err != nil {
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
func (_CrosschainTest *CrosschainTestFilterer) ParseCrossChain(log types.Log) (*CrosschainTestCrossChain, error) {
	event := new(CrosschainTestCrossChain)
	if err := _CrosschainTest.contract.UnpackLog(event, "CrossChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainTestIncreaseBridgeFeeIterator is returned from FilterIncreaseBridgeFee and is used to iterate over the raw logs and unpacked data for IncreaseBridgeFee events raised by the CrosschainTest contract.
type CrosschainTestIncreaseBridgeFeeIterator struct {
	Event *CrosschainTestIncreaseBridgeFee // Event containing the contract specifics and raw log

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
func (it *CrosschainTestIncreaseBridgeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainTestIncreaseBridgeFee)
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
		it.Event = new(CrosschainTestIncreaseBridgeFee)
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
func (it *CrosschainTestIncreaseBridgeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainTestIncreaseBridgeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainTestIncreaseBridgeFee represents a IncreaseBridgeFee event raised by the CrosschainTest contract.
type CrosschainTestIncreaseBridgeFee struct {
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
func (_CrosschainTest *CrosschainTestFilterer) FilterIncreaseBridgeFee(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*CrosschainTestIncreaseBridgeFeeIterator, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrosschainTest.contract.FilterLogs(opts, "IncreaseBridgeFee", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainTestIncreaseBridgeFeeIterator{contract: _CrosschainTest.contract, event: "IncreaseBridgeFee", logs: logs, sub: sub}, nil
}

// WatchIncreaseBridgeFee is a free log subscription operation binding the contract event 0x4b4d0e64eb77c0f61892107908295f09b3e381c50c655f4a73a4ad61c07350a0.
//
// Solidity: event IncreaseBridgeFee(address indexed sender, address indexed token, string chain, uint256 txID, uint256 fee)
func (_CrosschainTest *CrosschainTestFilterer) WatchIncreaseBridgeFee(opts *bind.WatchOpts, sink chan<- *CrosschainTestIncreaseBridgeFee, sender []common.Address, token []common.Address) (event.Subscription, error) {
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrosschainTest.contract.WatchLogs(opts, "IncreaseBridgeFee", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainTestIncreaseBridgeFee)
				if err := _CrosschainTest.contract.UnpackLog(event, "IncreaseBridgeFee", log); err != nil {
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
func (_CrosschainTest *CrosschainTestFilterer) ParseIncreaseBridgeFee(log types.Log) (*CrosschainTestIncreaseBridgeFee, error) {
	event := new(CrosschainTestIncreaseBridgeFee)
	if err := _CrosschainTest.contract.UnpackLog(event, "IncreaseBridgeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
