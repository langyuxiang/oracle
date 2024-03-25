// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package sdk


import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
)

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"NodeId\",\"type\":\"uint256\"}],\"name\":\"GetResult\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NodeId\",\"type\":\"uint256\"},{\"name\":\"_DataImportance\",\"type\":\"uint256\"},{\"name\":\"_NodeFailureRate\",\"type\":\"uint256\"},{\"name\":\"_Source\",\"type\":\"uint256\"},{\"name\":\"_Integrity\",\"type\":\"uint256\"},{\"name\":\"_Utilization\",\"type\":\"uint256\"},{\"name\":\"_Result\",\"type\":\"uint256\"}],\"name\":\"Grade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"IdToNode\",\"outputs\":[{\"name\":\"DataImportance\",\"type\":\"uint256\"},{\"name\":\"NodeFailureRate\",\"type\":\"uint256\"},{\"name\":\"Source\",\"type\":\"uint256\"},{\"name\":\"Integrity\",\"type\":\"uint256\"},{\"name\":\"Utilization\",\"type\":\"uint256\"},{\"name\":\"Result\",\"type\":\"uint256\"},{\"name\":\"NumberOfEvaluated\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NodeId\",\"type\":\"uint256\"}],\"name\":\"DeleteInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NodeId\",\"type\":\"uint256\"}],\"name\":\"AddNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OracleBin is the compiled bytecode used for deploying new contracts.
var OracleBin = "0x608060405234801561001057600080fd5b50610643806100206000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636225a8ad1461007257806376a03551146100b7578063866482c214610120578063c77f52281461018b578063ea7581ea146101b8575b600080fd5b34801561007e57600080fd5b5061009d600480360381019080803590602001909291905050506101e5565b604051808215151515815260200191505060405180910390f35b3480156100c357600080fd5b5061011e60048036038101908080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610219565b005b34801561012c57600080fd5b5061014b600480360381019080803590602001909291905050506104c7565b6040518088815260200187815260200186815260200185815260200184815260200183815260200182815260200197505050505050505060405180910390f35b34801561019757600080fd5b506101b660048036038101908080359060200190929190505050610509565b005b3480156101c457600080fd5b506101e360048036038101908080359060200190929190505050610571565b005b600060336000808481526020019081526020016000206005015410151561020f5760019050610214565b600090505b919050565b6000806000808a81526020019081526020016000206006015414156102f757866000808a815260200190815260200160002060000181905550856000808a815260200190815260200160002060010181905550846000808a815260200190815260200160002060020181905550836000808a815260200190815260200160002060030181905550826000808a815260200190815260200160002060040181905550606482026000808a81526020019081526020016000206005018190555060016000808a8152602001908152602001600020600601819055506104bd565b6000808981526020019081526020016000206006015490506001810187826000808c815260200190815260200160002060000154020181151561033657fe5b046000808a8152602001908152602001600020600001819055506001810186826000808c815260200190815260200160002060010154020181151561037757fe5b046000808a8152602001908152602001600020600101819055506001810185826000808c81526020019081526020016000206002015402018115156103b857fe5b046000808a8152602001908152602001600020600201819055506001810184826000808c81526020019081526020016000206003015402018115156103f957fe5b046000808a8152602001908152602001600020600301819055506001810183826000808c815260200190815260200160002060040154020181151561043a57fe5b046000808a8152602001908152602001600020600401819055506001810160648302826000808c815260200190815260200160002060050154020181151561047e57fe5b046000808a81526020019081526020016000206005018190555060016000808a8152602001908152602001600020600601600082825401925050819055505b5050505050505050565b60006020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154905087565b6105116105d9565b80600080848152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c082015181600601559050505050565b6105796105d9565b80600080848152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c082015181600601559050505050565b60e0604051908101604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815250905600a165627a7a72305820fd599403c4f63a50da4ce2e4de8bcd63a5c5708b20094158b0dbef6c31e7e4b00029"

// DeployOracle deploys a new contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

func AsyncDeployOracle(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Oracle is an auto generated Go binding around a Solidity contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around a Solidity contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around a Solidity contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around a Solidity contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetResult is a free data retrieval call binding the contract method 0x6225a8ad.
//
// Solidity: function GetResult(uint256 NodeId) constant returns(bool)
func (_Oracle *OracleCaller) GetResult(opts *bind.CallOpts, NodeId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "GetResult", NodeId)
	return *ret0, err
}

// GetResult is a free data retrieval call binding the contract method 0x6225a8ad.
//
// Solidity: function GetResult(uint256 NodeId) constant returns(bool)
func (_Oracle *OracleSession) GetResult(NodeId *big.Int) (bool, error) {
	return _Oracle.Contract.GetResult(&_Oracle.CallOpts, NodeId)
}

// GetResult is a free data retrieval call binding the contract method 0x6225a8ad.
//
// Solidity: function GetResult(uint256 NodeId) constant returns(bool)
func (_Oracle *OracleCallerSession) GetResult(NodeId *big.Int) (bool, error) {
	return _Oracle.Contract.GetResult(&_Oracle.CallOpts, NodeId)
}

// IdToNode is a free data retrieval call binding the contract method 0x866482c2.
//
// Solidity: function IdToNode(uint256 ) constant returns(uint256 DataImportance, uint256 NodeFailureRate, uint256 Source, uint256 Integrity, uint256 Utilization, uint256 Result, uint256 NumberOfEvaluated)
func (_Oracle *OracleCaller) IdToNode(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DataImportance    *big.Int
	NodeFailureRate   *big.Int
	Source            *big.Int
	Integrity         *big.Int
	Utilization       *big.Int
	Result            *big.Int
	NumberOfEvaluated *big.Int
}, error) {
	ret := new(struct {
		DataImportance    *big.Int
		NodeFailureRate   *big.Int
		Source            *big.Int
		Integrity         *big.Int
		Utilization       *big.Int
		Result            *big.Int
		NumberOfEvaluated *big.Int
	})
	out := ret
	err := _Oracle.contract.Call(opts, out, "IdToNode", arg0)
	return *ret, err
}

// IdToNode is a free data retrieval call binding the contract method 0x866482c2.
//
// Solidity: function IdToNode(uint256 ) constant returns(uint256 DataImportance, uint256 NodeFailureRate, uint256 Source, uint256 Integrity, uint256 Utilization, uint256 Result, uint256 NumberOfEvaluated)
func (_Oracle *OracleSession) IdToNode(arg0 *big.Int) (struct {
	DataImportance    *big.Int
	NodeFailureRate   *big.Int
	Source            *big.Int
	Integrity         *big.Int
	Utilization       *big.Int
	Result            *big.Int
	NumberOfEvaluated *big.Int
}, error) {
	return _Oracle.Contract.IdToNode(&_Oracle.CallOpts, arg0)
}

// IdToNode is a free data retrieval call binding the contract method 0x866482c2.
//
// Solidity: function IdToNode(uint256 ) constant returns(uint256 DataImportance, uint256 NodeFailureRate, uint256 Source, uint256 Integrity, uint256 Utilization, uint256 Result, uint256 NumberOfEvaluated)
func (_Oracle *OracleCallerSession) IdToNode(arg0 *big.Int) (struct {
	DataImportance    *big.Int
	NodeFailureRate   *big.Int
	Source            *big.Int
	Integrity         *big.Int
	Utilization       *big.Int
	Result            *big.Int
	NumberOfEvaluated *big.Int
}, error) {
	return _Oracle.Contract.IdToNode(&_Oracle.CallOpts, arg0)
}

// AddNode is a paid mutator transaction binding the contract method 0xea7581ea.
//
// Solidity: function AddNode(uint256 NodeId) returns()
func (_Oracle *OracleTransactor) AddNode(opts *bind.TransactOpts, NodeId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.contract.Transact(opts, "AddNode", NodeId)
}

func (_Oracle *OracleTransactor) AsyncAddNode(handler func(*types.Receipt, error), opts *bind.TransactOpts, NodeId *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.AsyncTransact(opts, handler, "AddNode", NodeId)
}

// AddNode is a paid mutator transaction binding the contract method 0xea7581ea.
//
// Solidity: function AddNode(uint256 NodeId) returns()
func (_Oracle *OracleSession) AddNode(NodeId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.AddNode(&_Oracle.TransactOpts, NodeId)
}

func (_Oracle *OracleSession) AsyncAddNode(handler func(*types.Receipt, error), NodeId *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AsyncAddNode(handler, &_Oracle.TransactOpts, NodeId)
}

// AddNode is a paid mutator transaction binding the contract method 0xea7581ea.
//
// Solidity: function AddNode(uint256 NodeId) returns()
func (_Oracle *OracleTransactorSession) AddNode(NodeId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.AddNode(&_Oracle.TransactOpts, NodeId)
}

func (_Oracle *OracleTransactorSession) AsyncAddNode(handler func(*types.Receipt, error), NodeId *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AsyncAddNode(handler, &_Oracle.TransactOpts, NodeId)
}

// DeleteInformation is a paid mutator transaction binding the contract method 0xc77f5228.
//
// Solidity: function DeleteInformation(uint256 NodeId) returns()
func (_Oracle *OracleTransactor) DeleteInformation(opts *bind.TransactOpts, NodeId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.contract.Transact(opts, "DeleteInformation", NodeId)
}

func (_Oracle *OracleTransactor) AsyncDeleteInformation(handler func(*types.Receipt, error), opts *bind.TransactOpts, NodeId *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.AsyncTransact(opts, handler, "DeleteInformation", NodeId)
}

// DeleteInformation is a paid mutator transaction binding the contract method 0xc77f5228.
//
// Solidity: function DeleteInformation(uint256 NodeId) returns()
func (_Oracle *OracleSession) DeleteInformation(NodeId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.DeleteInformation(&_Oracle.TransactOpts, NodeId)
}

func (_Oracle *OracleSession) AsyncDeleteInformation(handler func(*types.Receipt, error), NodeId *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AsyncDeleteInformation(handler, &_Oracle.TransactOpts, NodeId)
}

// DeleteInformation is a paid mutator transaction binding the contract method 0xc77f5228.
//
// Solidity: function DeleteInformation(uint256 NodeId) returns()
func (_Oracle *OracleTransactorSession) DeleteInformation(NodeId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.DeleteInformation(&_Oracle.TransactOpts, NodeId)
}

func (_Oracle *OracleTransactorSession) AsyncDeleteInformation(handler func(*types.Receipt, error), NodeId *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AsyncDeleteInformation(handler, &_Oracle.TransactOpts, NodeId)
}

// Grade is a paid mutator transaction binding the contract method 0x76a03551.
//
// Solidity: function Grade(uint256 NodeId, uint256 _DataImportance, uint256 _NodeFailureRate, uint256 _Source, uint256 _Integrity, uint256 _Utilization, uint256 _Result) returns()
func (_Oracle *OracleTransactor) Grade(opts *bind.TransactOpts, NodeId *big.Int, _DataImportance *big.Int, _NodeFailureRate *big.Int, _Source *big.Int, _Integrity *big.Int, _Utilization *big.Int, _Result *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.contract.Transact(opts, "Grade", NodeId, _DataImportance, _NodeFailureRate, _Source, _Integrity, _Utilization, _Result)
}

func (_Oracle *OracleTransactor) AsyncGrade(handler func(*types.Receipt, error), opts *bind.TransactOpts, NodeId *big.Int, _DataImportance *big.Int, _NodeFailureRate *big.Int, _Source *big.Int, _Integrity *big.Int, _Utilization *big.Int, _Result *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.AsyncTransact(opts, handler, "Grade", NodeId, _DataImportance, _NodeFailureRate, _Source, _Integrity, _Utilization, _Result)
}

// Grade is a paid mutator transaction binding the contract method 0x76a03551.
//
// Solidity: function Grade(uint256 NodeId, uint256 _DataImportance, uint256 _NodeFailureRate, uint256 _Source, uint256 _Integrity, uint256 _Utilization, uint256 _Result) returns()
func (_Oracle *OracleSession) Grade(NodeId *big.Int, _DataImportance *big.Int, _NodeFailureRate *big.Int, _Source *big.Int, _Integrity *big.Int, _Utilization *big.Int, _Result *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.Grade(&_Oracle.TransactOpts, NodeId, _DataImportance, _NodeFailureRate, _Source, _Integrity, _Utilization, _Result)
}

func (_Oracle *OracleSession) AsyncGrade(handler func(*types.Receipt, error), NodeId *big.Int, _DataImportance *big.Int, _NodeFailureRate *big.Int, _Source *big.Int, _Integrity *big.Int, _Utilization *big.Int, _Result *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AsyncGrade(handler, &_Oracle.TransactOpts, NodeId, _DataImportance, _NodeFailureRate, _Source, _Integrity, _Utilization, _Result)
}

// Grade is a paid mutator transaction binding the contract method 0x76a03551.
//
// Solidity: function Grade(uint256 NodeId, uint256 _DataImportance, uint256 _NodeFailureRate, uint256 _Source, uint256 _Integrity, uint256 _Utilization, uint256 _Result) returns()
func (_Oracle *OracleTransactorSession) Grade(NodeId *big.Int, _DataImportance *big.Int, _NodeFailureRate *big.Int, _Source *big.Int, _Integrity *big.Int, _Utilization *big.Int, _Result *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Oracle.Contract.Grade(&_Oracle.TransactOpts, NodeId, _DataImportance, _NodeFailureRate, _Source, _Integrity, _Utilization, _Result)
}

func (_Oracle *OracleTransactorSession) AsyncGrade(handler func(*types.Receipt, error), NodeId *big.Int, _DataImportance *big.Int, _NodeFailureRate *big.Int, _Source *big.Int, _Integrity *big.Int, _Utilization *big.Int, _Result *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.AsyncGrade(handler, &_Oracle.TransactOpts, NodeId, _DataImportance, _NodeFailureRate, _Source, _Integrity, _Utilization, _Result)
}
