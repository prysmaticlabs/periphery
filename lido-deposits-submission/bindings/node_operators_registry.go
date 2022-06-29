// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// NodeOperatorsRegistryMetaData contains all meta data concerning the NodeOperatorsRegistry contract.
var NodeOperatorsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_rewardAddress\",\"type\":\"address\"},{\"name\":\"_stakingLimit\",\"type\":\"uint64\"}],\"name\":\"addNodeOperator\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"},{\"name\":\"_quantity\",\"type\":\"uint256\"},{\"name\":\"_pubkeys\",\"type\":\"bytes\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"addSigningKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"name\":\"_usedSigningKeys\",\"type\":\"uint64[]\"}],\"name\":\"updateUsedKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SIGNATURE_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_NODE_OPERATOR_ADDRESS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setNodeOperatorName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setNodeOperatorActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_NODE_OPERATOR_NAME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeSigningKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ADD_NODE_OPERATOR_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"},{\"name\":\"_quantity\",\"type\":\"uint256\"},{\"name\":\"_pubkeys\",\"type\":\"bytes\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"addSigningKeysOperatorBH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveNodeOperatorsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"}],\"name\":\"getUnusedSigningKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRewardAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_fullInfo\",\"type\":\"bool\"}],\"name\":\"getNodeOperator\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"rewardAddress\",\"type\":\"address\"},{\"name\":\"stakingLimit\",\"type\":\"uint64\"},{\"name\":\"stoppedValidators\",\"type\":\"uint64\"},{\"name\":\"totalSigningKeys\",\"type\":\"uint64\"},{\"name\":\"usedSigningKeys\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PUBKEY_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeOperatorsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_stakingLimit\",\"type\":\"uint64\"}],\"name\":\"setNodeOperatorStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSigningKey\",\"outputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"name\":\"used\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_stoppedIncrement\",\"type\":\"uint64\"}],\"name\":\"reportStoppedValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REPORT_STOPPED_VALIDATORS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_NODE_OPERATOR_ACTIVE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_NODE_OPERATOR_LIMIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"}],\"name\":\"getTotalSigningKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator_id\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeSigningKeyOperatorBH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_SIGNING_KEYS\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"trimUnusedKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"rewardAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"stakingLimit\",\"type\":\"uint64\"}],\"name\":\"NodeOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"NodeOperatorActiveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NodeOperatorNameSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakingLimit\",\"type\":\"uint64\"}],\"name\":\"NodeOperatorStakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalStopped\",\"type\":\"uint64\"}],\"name\":\"NodeOperatorTotalStoppedValidatorsReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyRemoved\",\"type\":\"event\"}]",
}

// NodeOperatorsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeOperatorsRegistryMetaData.ABI instead.
var NodeOperatorsRegistryABI = NodeOperatorsRegistryMetaData.ABI

// NodeOperatorsRegistry is an auto generated Go binding around an Ethereum contract.
type NodeOperatorsRegistry struct {
	NodeOperatorsRegistryCaller     // Read-only binding to the contract
	NodeOperatorsRegistryTransactor // Write-only binding to the contract
	NodeOperatorsRegistryFilterer   // Log filterer for contract events
}

// NodeOperatorsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeOperatorsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeOperatorsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeOperatorsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeOperatorsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeOperatorsRegistrySession struct {
	Contract     *NodeOperatorsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeOperatorsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeOperatorsRegistryCallerSession struct {
	Contract *NodeOperatorsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// NodeOperatorsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeOperatorsRegistryTransactorSession struct {
	Contract     *NodeOperatorsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// NodeOperatorsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeOperatorsRegistryRaw struct {
	Contract *NodeOperatorsRegistry // Generic contract binding to access the raw methods on
}

// NodeOperatorsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryCallerRaw struct {
	Contract *NodeOperatorsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeOperatorsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeOperatorsRegistryTransactorRaw struct {
	Contract *NodeOperatorsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeOperatorsRegistry creates a new instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistry(address common.Address, backend bind.ContractBackend) (*NodeOperatorsRegistry, error) {
	contract, err := bindNodeOperatorsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistry{NodeOperatorsRegistryCaller: NodeOperatorsRegistryCaller{contract: contract}, NodeOperatorsRegistryTransactor: NodeOperatorsRegistryTransactor{contract: contract}, NodeOperatorsRegistryFilterer: NodeOperatorsRegistryFilterer{contract: contract}}, nil
}

// NewNodeOperatorsRegistryCaller creates a new read-only instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistryCaller(address common.Address, caller bind.ContractCaller) (*NodeOperatorsRegistryCaller, error) {
	contract, err := bindNodeOperatorsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryCaller{contract: contract}, nil
}

// NewNodeOperatorsRegistryTransactor creates a new write-only instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeOperatorsRegistryTransactor, error) {
	contract, err := bindNodeOperatorsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryTransactor{contract: contract}, nil
}

// NewNodeOperatorsRegistryFilterer creates a new log filterer instance of NodeOperatorsRegistry, bound to a specific deployed contract.
func NewNodeOperatorsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeOperatorsRegistryFilterer, error) {
	contract, err := bindNodeOperatorsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryFilterer{contract: contract}, nil
}

// bindNodeOperatorsRegistry binds a generic wrapper to an already deployed contract.
func bindNodeOperatorsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeOperatorsRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeOperatorsRegistry.Contract.NodeOperatorsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.NodeOperatorsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.NodeOperatorsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeOperatorsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.contract.Transact(opts, method, params...)
}

// ADDNODEOPERATORROLE is a free data retrieval call binding the contract method 0x7294d685.
//
// Solidity: function ADD_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) ADDNODEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "ADD_NODE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDNODEOPERATORROLE is a free data retrieval call binding the contract method 0x7294d685.
//
// Solidity: function ADD_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) ADDNODEOPERATORROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.ADDNODEOPERATORROLE(&_NodeOperatorsRegistry.CallOpts)
}

// ADDNODEOPERATORROLE is a free data retrieval call binding the contract method 0x7294d685.
//
// Solidity: function ADD_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) ADDNODEOPERATORROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.ADDNODEOPERATORROLE(&_NodeOperatorsRegistry.CallOpts)
}

// MANAGESIGNINGKEYS is a free data retrieval call binding the contract method 0xf31bd9c1.
//
// Solidity: function MANAGE_SIGNING_KEYS() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) MANAGESIGNINGKEYS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "MANAGE_SIGNING_KEYS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGESIGNINGKEYS is a free data retrieval call binding the contract method 0xf31bd9c1.
//
// Solidity: function MANAGE_SIGNING_KEYS() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) MANAGESIGNINGKEYS() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.MANAGESIGNINGKEYS(&_NodeOperatorsRegistry.CallOpts)
}

// MANAGESIGNINGKEYS is a free data retrieval call binding the contract method 0xf31bd9c1.
//
// Solidity: function MANAGE_SIGNING_KEYS() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) MANAGESIGNINGKEYS() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.MANAGESIGNINGKEYS(&_NodeOperatorsRegistry.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) PUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) PUBKEYLENGTH() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.PUBKEYLENGTH(&_NodeOperatorsRegistry.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) PUBKEYLENGTH() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.PUBKEYLENGTH(&_NodeOperatorsRegistry.CallOpts)
}

// REPORTSTOPPEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xcb10af07.
//
// Solidity: function REPORT_STOPPED_VALIDATORS_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) REPORTSTOPPEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "REPORT_STOPPED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTSTOPPEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xcb10af07.
//
// Solidity: function REPORT_STOPPED_VALIDATORS_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) REPORTSTOPPEDVALIDATORSROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.REPORTSTOPPEDVALIDATORSROLE(&_NodeOperatorsRegistry.CallOpts)
}

// REPORTSTOPPEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xcb10af07.
//
// Solidity: function REPORT_STOPPED_VALIDATORS_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) REPORTSTOPPEDVALIDATORSROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.REPORTSTOPPEDVALIDATORSROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORACTIVEROLE is a free data retrieval call binding the contract method 0xd6e1c2cc.
//
// Solidity: function SET_NODE_OPERATOR_ACTIVE_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) SETNODEOPERATORACTIVEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "SET_NODE_OPERATOR_ACTIVE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETNODEOPERATORACTIVEROLE is a free data retrieval call binding the contract method 0xd6e1c2cc.
//
// Solidity: function SET_NODE_OPERATOR_ACTIVE_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SETNODEOPERATORACTIVEROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORACTIVEROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORACTIVEROLE is a free data retrieval call binding the contract method 0xd6e1c2cc.
//
// Solidity: function SET_NODE_OPERATOR_ACTIVE_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) SETNODEOPERATORACTIVEROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORACTIVEROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORADDRESSROLE is a free data retrieval call binding the contract method 0x5a9fc07e.
//
// Solidity: function SET_NODE_OPERATOR_ADDRESS_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) SETNODEOPERATORADDRESSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "SET_NODE_OPERATOR_ADDRESS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETNODEOPERATORADDRESSROLE is a free data retrieval call binding the contract method 0x5a9fc07e.
//
// Solidity: function SET_NODE_OPERATOR_ADDRESS_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SETNODEOPERATORADDRESSROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORADDRESSROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORADDRESSROLE is a free data retrieval call binding the contract method 0x5a9fc07e.
//
// Solidity: function SET_NODE_OPERATOR_ADDRESS_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) SETNODEOPERATORADDRESSROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORADDRESSROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORLIMITROLE is a free data retrieval call binding the contract method 0xd8e71cd1.
//
// Solidity: function SET_NODE_OPERATOR_LIMIT_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) SETNODEOPERATORLIMITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "SET_NODE_OPERATOR_LIMIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETNODEOPERATORLIMITROLE is a free data retrieval call binding the contract method 0xd8e71cd1.
//
// Solidity: function SET_NODE_OPERATOR_LIMIT_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SETNODEOPERATORLIMITROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORLIMITROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORLIMITROLE is a free data retrieval call binding the contract method 0xd8e71cd1.
//
// Solidity: function SET_NODE_OPERATOR_LIMIT_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) SETNODEOPERATORLIMITROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORLIMITROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORNAMEROLE is a free data retrieval call binding the contract method 0x69602607.
//
// Solidity: function SET_NODE_OPERATOR_NAME_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) SETNODEOPERATORNAMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "SET_NODE_OPERATOR_NAME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETNODEOPERATORNAMEROLE is a free data retrieval call binding the contract method 0x69602607.
//
// Solidity: function SET_NODE_OPERATOR_NAME_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SETNODEOPERATORNAMEROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORNAMEROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SETNODEOPERATORNAMEROLE is a free data retrieval call binding the contract method 0x69602607.
//
// Solidity: function SET_NODE_OPERATOR_NAME_ROLE() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) SETNODEOPERATORNAMEROLE() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.SETNODEOPERATORNAMEROLE(&_NodeOperatorsRegistry.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SIGNATURELENGTH() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.SIGNATURELENGTH(&_NodeOperatorsRegistry.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) SIGNATURELENGTH() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.SIGNATURELENGTH(&_NodeOperatorsRegistry.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AllowRecoverability(token common.Address) (bool, error) {
	return _NodeOperatorsRegistry.Contract.AllowRecoverability(&_NodeOperatorsRegistry.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _NodeOperatorsRegistry.Contract.AllowRecoverability(&_NodeOperatorsRegistry.CallOpts, token)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AppId() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.AppId(&_NodeOperatorsRegistry.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) AppId() ([32]byte, error) {
	return _NodeOperatorsRegistry.Contract.AppId(&_NodeOperatorsRegistry.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.CanPerform(&_NodeOperatorsRegistry.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _NodeOperatorsRegistry.Contract.CanPerform(&_NodeOperatorsRegistry.CallOpts, _sender, _role, _params)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetActiveNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getActiveNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetActiveNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetActiveNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptExecutor(&_NodeOperatorsRegistry.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptExecutor(&_NodeOperatorsRegistry.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetEVMScriptRegistry() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptRegistry(&_NodeOperatorsRegistry.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetEVMScriptRegistry(&_NodeOperatorsRegistry.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetInitializationBlock() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetInitializationBlock(&_NodeOperatorsRegistry.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetInitializationBlock(&_NodeOperatorsRegistry.CallOpts)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _id, bool _fullInfo) view returns(bool active, string name, address rewardAddress, uint64 stakingLimit, uint64 stoppedValidators, uint64 totalSigningKeys, uint64 usedSigningKeys)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperator(opts *bind.CallOpts, _id *big.Int, _fullInfo bool) (struct {
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperator", _id, _fullInfo)

	outstruct := new(struct {
		Active            bool
		Name              string
		RewardAddress     common.Address
		StakingLimit      uint64
		StoppedValidators uint64
		TotalSigningKeys  uint64
		UsedSigningKeys   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.RewardAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StakingLimit = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.StoppedValidators = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.TotalSigningKeys = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.UsedSigningKeys = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _id, bool _fullInfo) view returns(bool active, string name, address rewardAddress, uint64 stakingLimit, uint64 stoppedValidators, uint64 totalSigningKeys, uint64 usedSigningKeys)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperator(_id *big.Int, _fullInfo bool) (struct {
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperator(&_NodeOperatorsRegistry.CallOpts, _id, _fullInfo)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x9a56983c.
//
// Solidity: function getNodeOperator(uint256 _id, bool _fullInfo) view returns(bool active, string name, address rewardAddress, uint64 stakingLimit, uint64 stoppedValidators, uint64 totalSigningKeys, uint64 usedSigningKeys)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperator(_id *big.Int, _fullInfo bool) (struct {
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperator(&_NodeOperatorsRegistry.CallOpts, _id, _fullInfo)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetNodeOperatorsCount(&_NodeOperatorsRegistry.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetRecoveryVault() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetRecoveryVault(&_NodeOperatorsRegistry.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetRecoveryVault() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.GetRecoveryVault(&_NodeOperatorsRegistry.CallOpts)
}

// GetSigningKey is a free data retrieval call binding the contract method 0xb449402a.
//
// Solidity: function getSigningKey(uint256 _operator_id, uint256 _index) view returns(bytes key, bytes depositSignature, bool used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetSigningKey(opts *bind.CallOpts, _operator_id *big.Int, _index *big.Int) (struct {
	Key              []byte
	DepositSignature []byte
	Used             bool
}, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getSigningKey", _operator_id, _index)

	outstruct := new(struct {
		Key              []byte
		DepositSignature []byte
		Used             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Key = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.DepositSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Used = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetSigningKey is a free data retrieval call binding the contract method 0xb449402a.
//
// Solidity: function getSigningKey(uint256 _operator_id, uint256 _index) view returns(bytes key, bytes depositSignature, bool used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetSigningKey(_operator_id *big.Int, _index *big.Int) (struct {
	Key              []byte
	DepositSignature []byte
	Used             bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetSigningKey(&_NodeOperatorsRegistry.CallOpts, _operator_id, _index)
}

// GetSigningKey is a free data retrieval call binding the contract method 0xb449402a.
//
// Solidity: function getSigningKey(uint256 _operator_id, uint256 _index) view returns(bytes key, bytes depositSignature, bool used)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetSigningKey(_operator_id *big.Int, _index *big.Int) (struct {
	Key              []byte
	DepositSignature []byte
	Used             bool
}, error) {
	return _NodeOperatorsRegistry.Contract.GetSigningKey(&_NodeOperatorsRegistry.CallOpts, _operator_id, _index)
}

// GetTotalSigningKeyCount is a free data retrieval call binding the contract method 0xdb9887ea.
//
// Solidity: function getTotalSigningKeyCount(uint256 _operator_id) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetTotalSigningKeyCount(opts *bind.CallOpts, _operator_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getTotalSigningKeyCount", _operator_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSigningKeyCount is a free data retrieval call binding the contract method 0xdb9887ea.
//
// Solidity: function getTotalSigningKeyCount(uint256 _operator_id) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetTotalSigningKeyCount(_operator_id *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetTotalSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _operator_id)
}

// GetTotalSigningKeyCount is a free data retrieval call binding the contract method 0xdb9887ea.
//
// Solidity: function getTotalSigningKeyCount(uint256 _operator_id) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetTotalSigningKeyCount(_operator_id *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetTotalSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _operator_id)
}

// GetUnusedSigningKeyCount is a free data retrieval call binding the contract method 0x8ca7c052.
//
// Solidity: function getUnusedSigningKeyCount(uint256 _operator_id) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) GetUnusedSigningKeyCount(opts *bind.CallOpts, _operator_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "getUnusedSigningKeyCount", _operator_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnusedSigningKeyCount is a free data retrieval call binding the contract method 0x8ca7c052.
//
// Solidity: function getUnusedSigningKeyCount(uint256 _operator_id) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) GetUnusedSigningKeyCount(_operator_id *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetUnusedSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _operator_id)
}

// GetUnusedSigningKeyCount is a free data retrieval call binding the contract method 0x8ca7c052.
//
// Solidity: function getUnusedSigningKeyCount(uint256 _operator_id) view returns(uint256)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) GetUnusedSigningKeyCount(_operator_id *big.Int) (*big.Int, error) {
	return _NodeOperatorsRegistry.Contract.GetUnusedSigningKeyCount(&_NodeOperatorsRegistry.CallOpts, _operator_id)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) HasInitialized() (bool, error) {
	return _NodeOperatorsRegistry.Contract.HasInitialized(&_NodeOperatorsRegistry.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) HasInitialized() (bool, error) {
	return _NodeOperatorsRegistry.Contract.HasInitialized(&_NodeOperatorsRegistry.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) IsPetrified() (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsPetrified(&_NodeOperatorsRegistry.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) IsPetrified() (bool, error) {
	return _NodeOperatorsRegistry.Contract.IsPetrified(&_NodeOperatorsRegistry.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) Kernel() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.Kernel(&_NodeOperatorsRegistry.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) Kernel() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.Kernel(&_NodeOperatorsRegistry.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeOperatorsRegistry.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) Pool() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.Pool(&_NodeOperatorsRegistry.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryCallerSession) Pool() (common.Address, error) {
	return _NodeOperatorsRegistry.Contract.Pool(&_NodeOperatorsRegistry.CallOpts)
}

// AddNodeOperator is a paid mutator transaction binding the contract method 0x09573fd4.
//
// Solidity: function addNodeOperator(string _name, address _rewardAddress, uint64 _stakingLimit) returns(uint256 id)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) AddNodeOperator(opts *bind.TransactOpts, _name string, _rewardAddress common.Address, _stakingLimit uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "addNodeOperator", _name, _rewardAddress, _stakingLimit)
}

// AddNodeOperator is a paid mutator transaction binding the contract method 0x09573fd4.
//
// Solidity: function addNodeOperator(string _name, address _rewardAddress, uint64 _stakingLimit) returns(uint256 id)
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AddNodeOperator(_name string, _rewardAddress common.Address, _stakingLimit uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _name, _rewardAddress, _stakingLimit)
}

// AddNodeOperator is a paid mutator transaction binding the contract method 0x09573fd4.
//
// Solidity: function addNodeOperator(string _name, address _rewardAddress, uint64 _stakingLimit) returns(uint256 id)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) AddNodeOperator(_name string, _rewardAddress common.Address, _stakingLimit uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddNodeOperator(&_NodeOperatorsRegistry.TransactOpts, _name, _rewardAddress, _stakingLimit)
}

// AddSigningKeys is a paid mutator transaction binding the contract method 0x096b7b35.
//
// Solidity: function addSigningKeys(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) AddSigningKeys(opts *bind.TransactOpts, _operator_id *big.Int, _quantity *big.Int, _pubkeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "addSigningKeys", _operator_id, _quantity, _pubkeys, _signatures)
}

// AddSigningKeys is a paid mutator transaction binding the contract method 0x096b7b35.
//
// Solidity: function addSigningKeys(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AddSigningKeys(_operator_id *big.Int, _quantity *big.Int, _pubkeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeys(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _quantity, _pubkeys, _signatures)
}

// AddSigningKeys is a paid mutator transaction binding the contract method 0x096b7b35.
//
// Solidity: function addSigningKeys(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) AddSigningKeys(_operator_id *big.Int, _quantity *big.Int, _pubkeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeys(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _quantity, _pubkeys, _signatures)
}

// AddSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x805911ae.
//
// Solidity: function addSigningKeysOperatorBH(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) AddSigningKeysOperatorBH(opts *bind.TransactOpts, _operator_id *big.Int, _quantity *big.Int, _pubkeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "addSigningKeysOperatorBH", _operator_id, _quantity, _pubkeys, _signatures)
}

// AddSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x805911ae.
//
// Solidity: function addSigningKeysOperatorBH(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) AddSigningKeysOperatorBH(_operator_id *big.Int, _quantity *big.Int, _pubkeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeysOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _quantity, _pubkeys, _signatures)
}

// AddSigningKeysOperatorBH is a paid mutator transaction binding the contract method 0x805911ae.
//
// Solidity: function addSigningKeysOperatorBH(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) AddSigningKeysOperatorBH(_operator_id *big.Int, _quantity *big.Int, _pubkeys []byte, _signatures []byte) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.AddSigningKeysOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _quantity, _pubkeys, _signatures)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0xa8031a1d.
//
// Solidity: function distributeRewards(address _token, uint256 _totalReward) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) DistributeRewards(opts *bind.TransactOpts, _token common.Address, _totalReward *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "distributeRewards", _token, _totalReward)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0xa8031a1d.
//
// Solidity: function distributeRewards(address _token, uint256 _totalReward) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) DistributeRewards(_token common.Address, _totalReward *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.DistributeRewards(&_NodeOperatorsRegistry.TransactOpts, _token, _totalReward)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0xa8031a1d.
//
// Solidity: function distributeRewards(address _token, uint256 _totalReward) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) DistributeRewards(_token common.Address, _totalReward *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.DistributeRewards(&_NodeOperatorsRegistry.TransactOpts, _token, _totalReward)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _pool) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) Initialize(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "initialize", _pool)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _pool) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) Initialize(_pool common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.Initialize(&_NodeOperatorsRegistry.TransactOpts, _pool)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _pool) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) Initialize(_pool common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.Initialize(&_NodeOperatorsRegistry.TransactOpts, _pool)
}

// RemoveSigningKey is a paid mutator transaction binding the contract method 0x6ef355f1.
//
// Solidity: function removeSigningKey(uint256 _operator_id, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) RemoveSigningKey(opts *bind.TransactOpts, _operator_id *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "removeSigningKey", _operator_id, _index)
}

// RemoveSigningKey is a paid mutator transaction binding the contract method 0x6ef355f1.
//
// Solidity: function removeSigningKey(uint256 _operator_id, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) RemoveSigningKey(_operator_id *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKey(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _index)
}

// RemoveSigningKey is a paid mutator transaction binding the contract method 0x6ef355f1.
//
// Solidity: function removeSigningKey(uint256 _operator_id, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) RemoveSigningKey(_operator_id *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKey(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _index)
}

// RemoveSigningKeyOperatorBH is a paid mutator transaction binding the contract method 0xed5cfa41.
//
// Solidity: function removeSigningKeyOperatorBH(uint256 _operator_id, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) RemoveSigningKeyOperatorBH(opts *bind.TransactOpts, _operator_id *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "removeSigningKeyOperatorBH", _operator_id, _index)
}

// RemoveSigningKeyOperatorBH is a paid mutator transaction binding the contract method 0xed5cfa41.
//
// Solidity: function removeSigningKeyOperatorBH(uint256 _operator_id, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) RemoveSigningKeyOperatorBH(_operator_id *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeyOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _index)
}

// RemoveSigningKeyOperatorBH is a paid mutator transaction binding the contract method 0xed5cfa41.
//
// Solidity: function removeSigningKeyOperatorBH(uint256 _operator_id, uint256 _index) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) RemoveSigningKeyOperatorBH(_operator_id *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.RemoveSigningKeyOperatorBH(&_NodeOperatorsRegistry.TransactOpts, _operator_id, _index)
}

// ReportStoppedValidators is a paid mutator transaction binding the contract method 0xbe726da2.
//
// Solidity: function reportStoppedValidators(uint256 _id, uint64 _stoppedIncrement) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) ReportStoppedValidators(opts *bind.TransactOpts, _id *big.Int, _stoppedIncrement uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "reportStoppedValidators", _id, _stoppedIncrement)
}

// ReportStoppedValidators is a paid mutator transaction binding the contract method 0xbe726da2.
//
// Solidity: function reportStoppedValidators(uint256 _id, uint64 _stoppedIncrement) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) ReportStoppedValidators(_id *big.Int, _stoppedIncrement uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ReportStoppedValidators(&_NodeOperatorsRegistry.TransactOpts, _id, _stoppedIncrement)
}

// ReportStoppedValidators is a paid mutator transaction binding the contract method 0xbe726da2.
//
// Solidity: function reportStoppedValidators(uint256 _id, uint64 _stoppedIncrement) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) ReportStoppedValidators(_id *big.Int, _stoppedIncrement uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.ReportStoppedValidators(&_NodeOperatorsRegistry.TransactOpts, _id, _stoppedIncrement)
}

// SetNodeOperatorActive is a paid mutator transaction binding the contract method 0x687ca337.
//
// Solidity: function setNodeOperatorActive(uint256 _id, bool _active) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorActive(opts *bind.TransactOpts, _id *big.Int, _active bool) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorActive", _id, _active)
}

// SetNodeOperatorActive is a paid mutator transaction binding the contract method 0x687ca337.
//
// Solidity: function setNodeOperatorActive(uint256 _id, bool _active) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorActive(_id *big.Int, _active bool) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorActive(&_NodeOperatorsRegistry.TransactOpts, _id, _active)
}

// SetNodeOperatorActive is a paid mutator transaction binding the contract method 0x687ca337.
//
// Solidity: function setNodeOperatorActive(uint256 _id, bool _active) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorActive(_id *big.Int, _active bool) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorActive(&_NodeOperatorsRegistry.TransactOpts, _id, _active)
}

// SetNodeOperatorName is a paid mutator transaction binding the contract method 0x5e57d742.
//
// Solidity: function setNodeOperatorName(uint256 _id, string _name) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorName(opts *bind.TransactOpts, _id *big.Int, _name string) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorName", _id, _name)
}

// SetNodeOperatorName is a paid mutator transaction binding the contract method 0x5e57d742.
//
// Solidity: function setNodeOperatorName(uint256 _id, string _name) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorName(_id *big.Int, _name string) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorName(&_NodeOperatorsRegistry.TransactOpts, _id, _name)
}

// SetNodeOperatorName is a paid mutator transaction binding the contract method 0x5e57d742.
//
// Solidity: function setNodeOperatorName(uint256 _id, string _name) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorName(_id *big.Int, _name string) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorName(&_NodeOperatorsRegistry.TransactOpts, _id, _name)
}

// SetNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x973e9328.
//
// Solidity: function setNodeOperatorRewardAddress(uint256 _id, address _rewardAddress) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorRewardAddress(opts *bind.TransactOpts, _id *big.Int, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorRewardAddress", _id, _rewardAddress)
}

// SetNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x973e9328.
//
// Solidity: function setNodeOperatorRewardAddress(uint256 _id, address _rewardAddress) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorRewardAddress(_id *big.Int, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorRewardAddress(&_NodeOperatorsRegistry.TransactOpts, _id, _rewardAddress)
}

// SetNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x973e9328.
//
// Solidity: function setNodeOperatorRewardAddress(uint256 _id, address _rewardAddress) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorRewardAddress(_id *big.Int, _rewardAddress common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorRewardAddress(&_NodeOperatorsRegistry.TransactOpts, _id, _rewardAddress)
}

// SetNodeOperatorStakingLimit is a paid mutator transaction binding the contract method 0xae962acf.
//
// Solidity: function setNodeOperatorStakingLimit(uint256 _id, uint64 _stakingLimit) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) SetNodeOperatorStakingLimit(opts *bind.TransactOpts, _id *big.Int, _stakingLimit uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "setNodeOperatorStakingLimit", _id, _stakingLimit)
}

// SetNodeOperatorStakingLimit is a paid mutator transaction binding the contract method 0xae962acf.
//
// Solidity: function setNodeOperatorStakingLimit(uint256 _id, uint64 _stakingLimit) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) SetNodeOperatorStakingLimit(_id *big.Int, _stakingLimit uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorStakingLimit(&_NodeOperatorsRegistry.TransactOpts, _id, _stakingLimit)
}

// SetNodeOperatorStakingLimit is a paid mutator transaction binding the contract method 0xae962acf.
//
// Solidity: function setNodeOperatorStakingLimit(uint256 _id, uint64 _stakingLimit) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) SetNodeOperatorStakingLimit(_id *big.Int, _stakingLimit uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.SetNodeOperatorStakingLimit(&_NodeOperatorsRegistry.TransactOpts, _id, _stakingLimit)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.TransferToVault(&_NodeOperatorsRegistry.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.TransferToVault(&_NodeOperatorsRegistry.TransactOpts, _token)
}

// TrimUnusedKeys is a paid mutator transaction binding the contract method 0xf778021e.
//
// Solidity: function trimUnusedKeys() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) TrimUnusedKeys(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "trimUnusedKeys")
}

// TrimUnusedKeys is a paid mutator transaction binding the contract method 0xf778021e.
//
// Solidity: function trimUnusedKeys() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) TrimUnusedKeys() (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.TrimUnusedKeys(&_NodeOperatorsRegistry.TransactOpts)
}

// TrimUnusedKeys is a paid mutator transaction binding the contract method 0xf778021e.
//
// Solidity: function trimUnusedKeys() returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) TrimUnusedKeys() (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.TrimUnusedKeys(&_NodeOperatorsRegistry.TransactOpts)
}

// UpdateUsedKeys is a paid mutator transaction binding the contract method 0x1fcf0938.
//
// Solidity: function updateUsedKeys(uint256[] _ids, uint64[] _usedSigningKeys) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactor) UpdateUsedKeys(opts *bind.TransactOpts, _ids []*big.Int, _usedSigningKeys []uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.contract.Transact(opts, "updateUsedKeys", _ids, _usedSigningKeys)
}

// UpdateUsedKeys is a paid mutator transaction binding the contract method 0x1fcf0938.
//
// Solidity: function updateUsedKeys(uint256[] _ids, uint64[] _usedSigningKeys) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistrySession) UpdateUsedKeys(_ids []*big.Int, _usedSigningKeys []uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateUsedKeys(&_NodeOperatorsRegistry.TransactOpts, _ids, _usedSigningKeys)
}

// UpdateUsedKeys is a paid mutator transaction binding the contract method 0x1fcf0938.
//
// Solidity: function updateUsedKeys(uint256[] _ids, uint64[] _usedSigningKeys) returns()
func (_NodeOperatorsRegistry *NodeOperatorsRegistryTransactorSession) UpdateUsedKeys(_ids []*big.Int, _usedSigningKeys []uint64) (*types.Transaction, error) {
	return _NodeOperatorsRegistry.Contract.UpdateUsedKeys(&_NodeOperatorsRegistry.TransactOpts, _ids, _usedSigningKeys)
}

// NodeOperatorsRegistryNodeOperatorActiveSetIterator is returned from FilterNodeOperatorActiveSet and is used to iterate over the raw logs and unpacked data for NodeOperatorActiveSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorActiveSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorActiveSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorActiveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorActiveSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorActiveSet)
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
func (it *NodeOperatorsRegistryNodeOperatorActiveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorActiveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorActiveSet represents a NodeOperatorActiveSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorActiveSet struct {
	Id     *big.Int
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorActiveSet is a free log retrieval operation binding the contract event 0xecdf08e8a6c4493efb460f6abc7d14532074fa339c3a6410623a1d3ee0fb2cac.
//
// Solidity: event NodeOperatorActiveSet(uint256 indexed id, bool active)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorActiveSet(opts *bind.FilterOpts, id []*big.Int) (*NodeOperatorsRegistryNodeOperatorActiveSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorActiveSet", idRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorActiveSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorActiveSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorActiveSet is a free log subscription operation binding the contract event 0xecdf08e8a6c4493efb460f6abc7d14532074fa339c3a6410623a1d3ee0fb2cac.
//
// Solidity: event NodeOperatorActiveSet(uint256 indexed id, bool active)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorActiveSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorActiveSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorActiveSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorActiveSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorActiveSet", log); err != nil {
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

// ParseNodeOperatorActiveSet is a log parse operation binding the contract event 0xecdf08e8a6c4493efb460f6abc7d14532074fa339c3a6410623a1d3ee0fb2cac.
//
// Solidity: event NodeOperatorActiveSet(uint256 indexed id, bool active)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorActiveSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorActiveSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorActiveSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorActiveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorAddedIterator is returned from FilterNodeOperatorAdded and is used to iterate over the raw logs and unpacked data for NodeOperatorAdded events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorAddedIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorAdded)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorAdded)
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
func (it *NodeOperatorsRegistryNodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorAdded represents a NodeOperatorAdded event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorAdded struct {
	Id            *big.Int
	Name          string
	RewardAddress common.Address
	StakingLimit  uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorAdded is a free log retrieval operation binding the contract event 0xc52ec0ad7872dae440d886040390c13677df7bf3cca136d8d81e5e5e7dd62ff1.
//
// Solidity: event NodeOperatorAdded(uint256 id, string name, address rewardAddress, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts) (*NodeOperatorsRegistryNodeOperatorAddedIterator, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorAddedIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorAdded is a free log subscription operation binding the contract event 0xc52ec0ad7872dae440d886040390c13677df7bf3cca136d8d81e5e5e7dd62ff1.
//
// Solidity: event NodeOperatorAdded(uint256 id, string name, address rewardAddress, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorAdded)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

// ParseNodeOperatorAdded is a log parse operation binding the contract event 0xc52ec0ad7872dae440d886040390c13677df7bf3cca136d8d81e5e5e7dd62ff1.
//
// Solidity: event NodeOperatorAdded(uint256 id, string name, address rewardAddress, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorAdded(log types.Log) (*NodeOperatorsRegistryNodeOperatorAdded, error) {
	event := new(NodeOperatorsRegistryNodeOperatorAdded)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorNameSetIterator is returned from FilterNodeOperatorNameSet and is used to iterate over the raw logs and unpacked data for NodeOperatorNameSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorNameSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorNameSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorNameSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorNameSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorNameSet)
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
func (it *NodeOperatorsRegistryNodeOperatorNameSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorNameSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorNameSet represents a NodeOperatorNameSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorNameSet struct {
	Id   *big.Int
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorNameSet is a free log retrieval operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 indexed id, string name)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorNameSet(opts *bind.FilterOpts, id []*big.Int) (*NodeOperatorsRegistryNodeOperatorNameSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorNameSet", idRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorNameSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorNameSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorNameSet is a free log subscription operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 indexed id, string name)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorNameSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorNameSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorNameSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorNameSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorNameSet", log); err != nil {
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

// ParseNodeOperatorNameSet is a log parse operation binding the contract event 0xcb16868f4831cc58a28d413f658752a2958bd1f50e94ed6391716b936c48093b.
//
// Solidity: event NodeOperatorNameSet(uint256 indexed id, string name)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorNameSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorNameSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorNameSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorNameSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator is returned from FilterNodeOperatorRewardAddressSet and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorRewardAddressSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
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
func (it *NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorRewardAddressSet represents a NodeOperatorRewardAddressSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorRewardAddressSet struct {
	Id            *big.Int
	RewardAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressSet is a free log retrieval operation binding the contract event 0x9a52205165d510fc1e428886d52108725dc01ed544da1702dc7bd3fdb3f243b2.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 indexed id, address rewardAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorRewardAddressSet(opts *bind.FilterOpts, id []*big.Int) (*NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorRewardAddressSet", idRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorRewardAddressSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorRewardAddressSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressSet is a free log subscription operation binding the contract event 0x9a52205165d510fc1e428886d52108725dc01ed544da1702dc7bd3fdb3f243b2.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 indexed id, address rewardAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorRewardAddressSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorRewardAddressSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorRewardAddressSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorRewardAddressSet", log); err != nil {
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

// ParseNodeOperatorRewardAddressSet is a log parse operation binding the contract event 0x9a52205165d510fc1e428886d52108725dc01ed544da1702dc7bd3fdb3f243b2.
//
// Solidity: event NodeOperatorRewardAddressSet(uint256 indexed id, address rewardAddress)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorRewardAddressSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorRewardAddressSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorRewardAddressSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorRewardAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator is returned from FilterNodeOperatorStakingLimitSet and is used to iterate over the raw logs and unpacked data for NodeOperatorStakingLimitSet events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorStakingLimitSet)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorStakingLimitSet)
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
func (it *NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorStakingLimitSet represents a NodeOperatorStakingLimitSet event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorStakingLimitSet struct {
	Id           *big.Int
	StakingLimit uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorStakingLimitSet is a free log retrieval operation binding the contract event 0x59d11713a8801e3ba782d59e757fbcef75ca2ecabce8ccd06a0827941230b9f2.
//
// Solidity: event NodeOperatorStakingLimitSet(uint256 indexed id, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorStakingLimitSet(opts *bind.FilterOpts, id []*big.Int) (*NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorStakingLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorStakingLimitSetIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorStakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorStakingLimitSet is a free log subscription operation binding the contract event 0x59d11713a8801e3ba782d59e757fbcef75ca2ecabce8ccd06a0827941230b9f2.
//
// Solidity: event NodeOperatorStakingLimitSet(uint256 indexed id, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorStakingLimitSet(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorStakingLimitSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorStakingLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorStakingLimitSet)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorStakingLimitSet", log); err != nil {
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

// ParseNodeOperatorStakingLimitSet is a log parse operation binding the contract event 0x59d11713a8801e3ba782d59e757fbcef75ca2ecabce8ccd06a0827941230b9f2.
//
// Solidity: event NodeOperatorStakingLimitSet(uint256 indexed id, uint64 stakingLimit)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorStakingLimitSet(log types.Log) (*NodeOperatorsRegistryNodeOperatorStakingLimitSet, error) {
	event := new(NodeOperatorsRegistryNodeOperatorStakingLimitSet)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorStakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator is returned from FilterNodeOperatorTotalStoppedValidatorsReported and is used to iterate over the raw logs and unpacked data for NodeOperatorTotalStoppedValidatorsReported events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator struct {
	Event *NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported)
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
		it.Event = new(NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported)
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
func (it *NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported represents a NodeOperatorTotalStoppedValidatorsReported event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported struct {
	Id           *big.Int
	TotalStopped uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorTotalStoppedValidatorsReported is a free log retrieval operation binding the contract event 0xe6452c223b953d8ab5cb26c014095615322891268537911ba6fe1c939689703d.
//
// Solidity: event NodeOperatorTotalStoppedValidatorsReported(uint256 indexed id, uint64 totalStopped)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterNodeOperatorTotalStoppedValidatorsReported(opts *bind.FilterOpts, id []*big.Int) (*NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "NodeOperatorTotalStoppedValidatorsReported", idRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReportedIterator{contract: _NodeOperatorsRegistry.contract, event: "NodeOperatorTotalStoppedValidatorsReported", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorTotalStoppedValidatorsReported is a free log subscription operation binding the contract event 0xe6452c223b953d8ab5cb26c014095615322891268537911ba6fe1c939689703d.
//
// Solidity: event NodeOperatorTotalStoppedValidatorsReported(uint256 indexed id, uint64 totalStopped)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchNodeOperatorTotalStoppedValidatorsReported(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "NodeOperatorTotalStoppedValidatorsReported", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorTotalStoppedValidatorsReported", log); err != nil {
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

// ParseNodeOperatorTotalStoppedValidatorsReported is a log parse operation binding the contract event 0xe6452c223b953d8ab5cb26c014095615322891268537911ba6fe1c939689703d.
//
// Solidity: event NodeOperatorTotalStoppedValidatorsReported(uint256 indexed id, uint64 totalStopped)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseNodeOperatorTotalStoppedValidatorsReported(log types.Log) (*NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported, error) {
	event := new(NodeOperatorsRegistryNodeOperatorTotalStoppedValidatorsReported)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "NodeOperatorTotalStoppedValidatorsReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryRecoverToVaultIterator struct {
	Event *NodeOperatorsRegistryRecoverToVault // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryRecoverToVault)
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
		it.Event = new(NodeOperatorsRegistryRecoverToVault)
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
func (it *NodeOperatorsRegistryRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryRecoverToVault represents a RecoverToVault event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*NodeOperatorsRegistryRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryRecoverToVaultIterator{contract: _NodeOperatorsRegistry.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryRecoverToVault)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseRecoverToVault(log types.Log) (*NodeOperatorsRegistryRecoverToVault, error) {
	event := new(NodeOperatorsRegistryRecoverToVault)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistryScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryScriptResultIterator struct {
	Event *NodeOperatorsRegistryScriptResult // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistryScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistryScriptResult)
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
		it.Event = new(NodeOperatorsRegistryScriptResult)
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
func (it *NodeOperatorsRegistryScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistryScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistryScriptResult represents a ScriptResult event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistryScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*NodeOperatorsRegistryScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistryScriptResultIterator{contract: _NodeOperatorsRegistry.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistryScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistryScriptResult)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseScriptResult(log types.Log) (*NodeOperatorsRegistryScriptResult, error) {
	event := new(NodeOperatorsRegistryScriptResult)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistrySigningKeyAddedIterator is returned from FilterSigningKeyAdded and is used to iterate over the raw logs and unpacked data for SigningKeyAdded events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistrySigningKeyAddedIterator struct {
	Event *NodeOperatorsRegistrySigningKeyAdded // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistrySigningKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistrySigningKeyAdded)
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
		it.Event = new(NodeOperatorsRegistrySigningKeyAdded)
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
func (it *NodeOperatorsRegistrySigningKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistrySigningKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistrySigningKeyAdded represents a SigningKeyAdded event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistrySigningKeyAdded struct {
	OperatorId *big.Int
	Pubkey     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyAdded is a free log retrieval operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed operatorId, bytes pubkey)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterSigningKeyAdded(opts *bind.FilterOpts, operatorId []*big.Int) (*NodeOperatorsRegistrySigningKeyAddedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "SigningKeyAdded", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistrySigningKeyAddedIterator{contract: _NodeOperatorsRegistry.contract, event: "SigningKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSigningKeyAdded is a free log subscription operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed operatorId, bytes pubkey)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchSigningKeyAdded(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistrySigningKeyAdded, operatorId []*big.Int) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "SigningKeyAdded", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistrySigningKeyAdded)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
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

// ParseSigningKeyAdded is a log parse operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed operatorId, bytes pubkey)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseSigningKeyAdded(log types.Log) (*NodeOperatorsRegistrySigningKeyAdded, error) {
	event := new(NodeOperatorsRegistrySigningKeyAdded)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOperatorsRegistrySigningKeyRemovedIterator is returned from FilterSigningKeyRemoved and is used to iterate over the raw logs and unpacked data for SigningKeyRemoved events raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistrySigningKeyRemovedIterator struct {
	Event *NodeOperatorsRegistrySigningKeyRemoved // Event containing the contract specifics and raw log

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
func (it *NodeOperatorsRegistrySigningKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOperatorsRegistrySigningKeyRemoved)
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
		it.Event = new(NodeOperatorsRegistrySigningKeyRemoved)
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
func (it *NodeOperatorsRegistrySigningKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOperatorsRegistrySigningKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOperatorsRegistrySigningKeyRemoved represents a SigningKeyRemoved event raised by the NodeOperatorsRegistry contract.
type NodeOperatorsRegistrySigningKeyRemoved struct {
	OperatorId *big.Int
	Pubkey     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyRemoved is a free log retrieval operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed operatorId, bytes pubkey)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) FilterSigningKeyRemoved(opts *bind.FilterOpts, operatorId []*big.Int) (*NodeOperatorsRegistrySigningKeyRemovedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.FilterLogs(opts, "SigningKeyRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodeOperatorsRegistrySigningKeyRemovedIterator{contract: _NodeOperatorsRegistry.contract, event: "SigningKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchSigningKeyRemoved is a free log subscription operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed operatorId, bytes pubkey)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) WatchSigningKeyRemoved(opts *bind.WatchOpts, sink chan<- *NodeOperatorsRegistrySigningKeyRemoved, operatorId []*big.Int) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _NodeOperatorsRegistry.contract.WatchLogs(opts, "SigningKeyRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOperatorsRegistrySigningKeyRemoved)
				if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
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

// ParseSigningKeyRemoved is a log parse operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed operatorId, bytes pubkey)
func (_NodeOperatorsRegistry *NodeOperatorsRegistryFilterer) ParseSigningKeyRemoved(log types.Log) (*NodeOperatorsRegistrySigningKeyRemoved, error) {
	event := new(NodeOperatorsRegistrySigningKeyRemoved)
	if err := _NodeOperatorsRegistry.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
