// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ERC721FactoryMetaData contains all meta data concerning the ERC721Factory contract.
var ERC721FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"BookCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMintAmountPerTx\",\"type\":\"uint256\"}],\"name\":\"deployERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"registryERC721Contract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC721FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721FactoryMetaData.ABI instead.
var ERC721FactoryABI = ERC721FactoryMetaData.ABI

// ERC721Factory is an auto generated Go binding around an Ethereum contract.
type ERC721Factory struct {
	ERC721FactoryCaller     // Read-only binding to the contract
	ERC721FactoryTransactor // Write-only binding to the contract
	ERC721FactoryFilterer   // Log filterer for contract events
}

// ERC721FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721FactorySession struct {
	Contract     *ERC721Factory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721FactoryCallerSession struct {
	Contract *ERC721FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721FactoryTransactorSession struct {
	Contract     *ERC721FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721FactoryRaw struct {
	Contract *ERC721Factory // Generic contract binding to access the raw methods on
}

// ERC721FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721FactoryCallerRaw struct {
	Contract *ERC721FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721FactoryTransactorRaw struct {
	Contract *ERC721FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Factory creates a new instance of ERC721Factory, bound to a specific deployed contract.
func NewERC721Factory(address common.Address, backend bind.ContractBackend) (*ERC721Factory, error) {
	contract, err := bindERC721Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Factory{ERC721FactoryCaller: ERC721FactoryCaller{contract: contract}, ERC721FactoryTransactor: ERC721FactoryTransactor{contract: contract}, ERC721FactoryFilterer: ERC721FactoryFilterer{contract: contract}}, nil
}

// NewERC721FactoryCaller creates a new read-only instance of ERC721Factory, bound to a specific deployed contract.
func NewERC721FactoryCaller(address common.Address, caller bind.ContractCaller) (*ERC721FactoryCaller, error) {
	contract, err := bindERC721Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721FactoryCaller{contract: contract}, nil
}

// NewERC721FactoryTransactor creates a new write-only instance of ERC721Factory, bound to a specific deployed contract.
func NewERC721FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721FactoryTransactor, error) {
	contract, err := bindERC721Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721FactoryTransactor{contract: contract}, nil
}

// NewERC721FactoryFilterer creates a new log filterer instance of ERC721Factory, bound to a specific deployed contract.
func NewERC721FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721FactoryFilterer, error) {
	contract, err := bindERC721Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721FactoryFilterer{contract: contract}, nil
}

// bindERC721Factory binds a generic wrapper to an already deployed contract.
func bindERC721Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Factory *ERC721FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Factory.Contract.ERC721FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Factory *ERC721FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Factory.Contract.ERC721FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Factory *ERC721FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Factory.Contract.ERC721FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Factory *ERC721FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Factory *ERC721FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Factory *ERC721FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Factory.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_ERC721Factory *ERC721FactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Factory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_ERC721Factory *ERC721FactorySession) Exists(arg0 common.Address) (bool, error) {
	return _ERC721Factory.Contract.Exists(&_ERC721Factory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_ERC721Factory *ERC721FactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _ERC721Factory.Contract.Exists(&_ERC721Factory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_ERC721Factory *ERC721FactoryCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Factory.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_ERC721Factory *ERC721FactorySession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _ERC721Factory.Contract.Tokens(&_ERC721Factory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_ERC721Factory *ERC721FactoryCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _ERC721Factory.Contract.Tokens(&_ERC721Factory.CallOpts, arg0)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0xb96b89de.
//
// Solidity: function deployERC721(string _tokenName, string _tokenSymbol, string _customURI, uint256 _cost, uint256 _maxSupply, uint256 _maxMintAmountPerTx) returns(address)
func (_ERC721Factory *ERC721FactoryTransactor) DeployERC721(opts *bind.TransactOpts, _tokenName string, _tokenSymbol string, _customURI string, _cost *big.Int, _maxSupply *big.Int, _maxMintAmountPerTx *big.Int) (*types.Transaction, error) {
	return _ERC721Factory.contract.Transact(opts, "deployERC721", _tokenName, _tokenSymbol, _customURI, _cost, _maxSupply, _maxMintAmountPerTx)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0xb96b89de.
//
// Solidity: function deployERC721(string _tokenName, string _tokenSymbol, string _customURI, uint256 _cost, uint256 _maxSupply, uint256 _maxMintAmountPerTx) returns(address)
func (_ERC721Factory *ERC721FactorySession) DeployERC721(_tokenName string, _tokenSymbol string, _customURI string, _cost *big.Int, _maxSupply *big.Int, _maxMintAmountPerTx *big.Int) (*types.Transaction, error) {
	return _ERC721Factory.Contract.DeployERC721(&_ERC721Factory.TransactOpts, _tokenName, _tokenSymbol, _customURI, _cost, _maxSupply, _maxMintAmountPerTx)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0xb96b89de.
//
// Solidity: function deployERC721(string _tokenName, string _tokenSymbol, string _customURI, uint256 _cost, uint256 _maxSupply, uint256 _maxMintAmountPerTx) returns(address)
func (_ERC721Factory *ERC721FactoryTransactorSession) DeployERC721(_tokenName string, _tokenSymbol string, _customURI string, _cost *big.Int, _maxSupply *big.Int, _maxMintAmountPerTx *big.Int) (*types.Transaction, error) {
	return _ERC721Factory.Contract.DeployERC721(&_ERC721Factory.TransactOpts, _tokenName, _tokenSymbol, _customURI, _cost, _maxSupply, _maxMintAmountPerTx)
}

// RegistryERC721Contract is a paid mutator transaction binding the contract method 0xc24b891f.
//
// Solidity: function registryERC721Contract(address contractAddress, string url) returns()
func (_ERC721Factory *ERC721FactoryTransactor) RegistryERC721Contract(opts *bind.TransactOpts, contractAddress common.Address, url string) (*types.Transaction, error) {
	return _ERC721Factory.contract.Transact(opts, "registryERC721Contract", contractAddress, url)
}

// RegistryERC721Contract is a paid mutator transaction binding the contract method 0xc24b891f.
//
// Solidity: function registryERC721Contract(address contractAddress, string url) returns()
func (_ERC721Factory *ERC721FactorySession) RegistryERC721Contract(contractAddress common.Address, url string) (*types.Transaction, error) {
	return _ERC721Factory.Contract.RegistryERC721Contract(&_ERC721Factory.TransactOpts, contractAddress, url)
}

// RegistryERC721Contract is a paid mutator transaction binding the contract method 0xc24b891f.
//
// Solidity: function registryERC721Contract(address contractAddress, string url) returns()
func (_ERC721Factory *ERC721FactoryTransactorSession) RegistryERC721Contract(contractAddress common.Address, url string) (*types.Transaction, error) {
	return _ERC721Factory.Contract.RegistryERC721Contract(&_ERC721Factory.TransactOpts, contractAddress, url)
}

// ERC721FactoryBookCreatedIterator is returned from FilterBookCreated and is used to iterate over the raw logs and unpacked data for BookCreated events raised by the ERC721Factory contract.
type ERC721FactoryBookCreatedIterator struct {
	Event *ERC721FactoryBookCreated // Event containing the contract specifics and raw log

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
func (it *ERC721FactoryBookCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721FactoryBookCreated)
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
		it.Event = new(ERC721FactoryBookCreated)
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
func (it *ERC721FactoryBookCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721FactoryBookCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721FactoryBookCreated represents a BookCreated event raised by the ERC721Factory contract.
type ERC721FactoryBookCreated struct {
	Owner         common.Address
	TokenContract common.Address
	Url           string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBookCreated is a free log retrieval operation binding the contract event 0xdd1d1ce24f6b921a22afd9f7d77847b7dfaf8c19f506b806943c7bbfaa766948.
//
// Solidity: event BookCreated(address owner, address tokenContract, string url)
func (_ERC721Factory *ERC721FactoryFilterer) FilterBookCreated(opts *bind.FilterOpts) (*ERC721FactoryBookCreatedIterator, error) {

	logs, sub, err := _ERC721Factory.contract.FilterLogs(opts, "BookCreated")
	if err != nil {
		return nil, err
	}
	return &ERC721FactoryBookCreatedIterator{contract: _ERC721Factory.contract, event: "BookCreated", logs: logs, sub: sub}, nil
}

// WatchBookCreated is a free log subscription operation binding the contract event 0xdd1d1ce24f6b921a22afd9f7d77847b7dfaf8c19f506b806943c7bbfaa766948.
//
// Solidity: event BookCreated(address owner, address tokenContract, string url)
func (_ERC721Factory *ERC721FactoryFilterer) WatchBookCreated(opts *bind.WatchOpts, sink chan<- *ERC721FactoryBookCreated) (event.Subscription, error) {

	logs, sub, err := _ERC721Factory.contract.WatchLogs(opts, "BookCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721FactoryBookCreated)
				if err := _ERC721Factory.contract.UnpackLog(event, "BookCreated", log); err != nil {
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

// ParseBookCreated is a log parse operation binding the contract event 0xdd1d1ce24f6b921a22afd9f7d77847b7dfaf8c19f506b806943c7bbfaa766948.
//
// Solidity: event BookCreated(address owner, address tokenContract, string url)
func (_ERC721Factory *ERC721FactoryFilterer) ParseBookCreated(log types.Log) (*ERC721FactoryBookCreated, error) {
	event := new(ERC721FactoryBookCreated)
	if err := _ERC721Factory.contract.UnpackLog(event, "BookCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
