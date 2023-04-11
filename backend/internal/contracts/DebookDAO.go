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

// DebookDAOProposal is an auto generated low-level Go binding around an user-defined struct.
type DebookDAOProposal struct {
	Proposer      common.Address
	Title         [32]byte
	Description   string
	VoteStart     TimersTimestamp
	VoteEnd       TimersTimestamp
	UpVoteCount   *big.Int
	DownVoteCount *big.Int
	IsSet         bool
	IsExecuted    bool
	IsCanceled    bool
}

// TimersTimestamp is an auto generated low-level Go binding around an user-defined struct.
type TimersTimestamp struct {
	Deadline uint64
}

// DebookDAOMetaData contains all meta data concerning the DebookDAO contract.
var DebookDAOMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"message\",\"type\":\"address\"}],\"name\":\"Console\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDebookDAO.VoteState\",\"name\":\"selection\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PROPOSAL_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumDebookDAO.VoteState\",\"name\":\"selection\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteWithReasonBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"exeute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"title\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_deadline\",\"type\":\"uint64\"}],\"internalType\":\"structTimers.Timestamp\",\"name\":\"voteStart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_deadline\",\"type\":\"uint64\"}],\"internalType\":\"structTimers.Timestamp\",\"name\":\"voteEnd\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"upVoteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"downVoteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCanceled\",\"type\":\"bool\"}],\"internalType\":\"structDebookDAO.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalVoteStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"title\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interactContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"enumDebookDAO.VoteState\",\"name\":\"selection\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"hashVote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voterId\",\"type\":\"address\"}],\"name\":\"isVoterVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"title\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sp\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"setStakePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sp\",\"outputs\":[{\"internalType\":\"contractIStakePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumDebookDAO.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DebookDAOABI is the input ABI used to generate the binding from.
// Deprecated: Use DebookDAOMetaData.ABI instead.
var DebookDAOABI = DebookDAOMetaData.ABI

// DebookDAO is an auto generated Go binding around an Ethereum contract.
type DebookDAO struct {
	DebookDAOCaller     // Read-only binding to the contract
	DebookDAOTransactor // Write-only binding to the contract
	DebookDAOFilterer   // Log filterer for contract events
}

// DebookDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebookDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebookDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebookDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebookDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebookDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebookDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebookDAOSession struct {
	Contract     *DebookDAO        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebookDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebookDAOCallerSession struct {
	Contract *DebookDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DebookDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebookDAOTransactorSession struct {
	Contract     *DebookDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DebookDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type DebookDAORaw struct {
	Contract *DebookDAO // Generic contract binding to access the raw methods on
}

// DebookDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebookDAOCallerRaw struct {
	Contract *DebookDAOCaller // Generic read-only contract binding to access the raw methods on
}

// DebookDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebookDAOTransactorRaw struct {
	Contract *DebookDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebookDAO creates a new instance of DebookDAO, bound to a specific deployed contract.
func NewDebookDAO(address common.Address, backend bind.ContractBackend) (*DebookDAO, error) {
	contract, err := bindDebookDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebookDAO{DebookDAOCaller: DebookDAOCaller{contract: contract}, DebookDAOTransactor: DebookDAOTransactor{contract: contract}, DebookDAOFilterer: DebookDAOFilterer{contract: contract}}, nil
}

// NewDebookDAOCaller creates a new read-only instance of DebookDAO, bound to a specific deployed contract.
func NewDebookDAOCaller(address common.Address, caller bind.ContractCaller) (*DebookDAOCaller, error) {
	contract, err := bindDebookDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebookDAOCaller{contract: contract}, nil
}

// NewDebookDAOTransactor creates a new write-only instance of DebookDAO, bound to a specific deployed contract.
func NewDebookDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*DebookDAOTransactor, error) {
	contract, err := bindDebookDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebookDAOTransactor{contract: contract}, nil
}

// NewDebookDAOFilterer creates a new log filterer instance of DebookDAO, bound to a specific deployed contract.
func NewDebookDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*DebookDAOFilterer, error) {
	contract, err := bindDebookDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebookDAOFilterer{contract: contract}, nil
}

// bindDebookDAO binds a generic wrapper to an already deployed contract.
func bindDebookDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DebookDAOMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebookDAO *DebookDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebookDAO.Contract.DebookDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebookDAO *DebookDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebookDAO.Contract.DebookDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebookDAO *DebookDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebookDAO.Contract.DebookDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebookDAO *DebookDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebookDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebookDAO *DebookDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebookDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebookDAO *DebookDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebookDAO.Contract.contract.Transact(opts, method, params...)
}

// PROPOSALTYPEHASH is a free data retrieval call binding the contract method 0x853b94a0.
//
// Solidity: function PROPOSAL_TYPEHASH() view returns(bytes32)
func (_DebookDAO *DebookDAOCaller) PROPOSALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "PROPOSAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSALTYPEHASH is a free data retrieval call binding the contract method 0x853b94a0.
//
// Solidity: function PROPOSAL_TYPEHASH() view returns(bytes32)
func (_DebookDAO *DebookDAOSession) PROPOSALTYPEHASH() ([32]byte, error) {
	return _DebookDAO.Contract.PROPOSALTYPEHASH(&_DebookDAO.CallOpts)
}

// PROPOSALTYPEHASH is a free data retrieval call binding the contract method 0x853b94a0.
//
// Solidity: function PROPOSAL_TYPEHASH() view returns(bytes32)
func (_DebookDAO *DebookDAOCallerSession) PROPOSALTYPEHASH() ([32]byte, error) {
	return _DebookDAO.Contract.PROPOSALTYPEHASH(&_DebookDAO.CallOpts)
}

// VOTETYPEHASH is a free data retrieval call binding the contract method 0x86522973.
//
// Solidity: function VOTE_TYPEHASH() view returns(bytes32)
func (_DebookDAO *DebookDAOCaller) VOTETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "VOTE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VOTETYPEHASH is a free data retrieval call binding the contract method 0x86522973.
//
// Solidity: function VOTE_TYPEHASH() view returns(bytes32)
func (_DebookDAO *DebookDAOSession) VOTETYPEHASH() ([32]byte, error) {
	return _DebookDAO.Contract.VOTETYPEHASH(&_DebookDAO.CallOpts)
}

// VOTETYPEHASH is a free data retrieval call binding the contract method 0x86522973.
//
// Solidity: function VOTE_TYPEHASH() view returns(bytes32)
func (_DebookDAO *DebookDAOCallerSession) VOTETYPEHASH() ([32]byte, error) {
	return _DebookDAO.Contract.VOTETYPEHASH(&_DebookDAO.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0xe4cdd80e.
//
// Solidity: function _votingDelay() view returns(uint256)
func (_DebookDAO *DebookDAOCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "_votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0xe4cdd80e.
//
// Solidity: function _votingDelay() view returns(uint256)
func (_DebookDAO *DebookDAOSession) VotingDelay() (*big.Int, error) {
	return _DebookDAO.Contract.VotingDelay(&_DebookDAO.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0xe4cdd80e.
//
// Solidity: function _votingDelay() view returns(uint256)
func (_DebookDAO *DebookDAOCallerSession) VotingDelay() (*big.Int, error) {
	return _DebookDAO.Contract.VotingDelay(&_DebookDAO.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x9199907a.
//
// Solidity: function _votingPeriod() view returns(uint256)
func (_DebookDAO *DebookDAOCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "_votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x9199907a.
//
// Solidity: function _votingPeriod() view returns(uint256)
func (_DebookDAO *DebookDAOSession) VotingPeriod() (*big.Int, error) {
	return _DebookDAO.Contract.VotingPeriod(&_DebookDAO.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x9199907a.
//
// Solidity: function _votingPeriod() view returns(uint256)
func (_DebookDAO *DebookDAOCallerSession) VotingPeriod() (*big.Int, error) {
	return _DebookDAO.Contract.VotingPeriod(&_DebookDAO.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns((address,bytes32,string,(uint64),(uint64),uint256,uint256,bool,bool,bool))
func (_DebookDAO *DebookDAOCaller) GetProposal(opts *bind.CallOpts, proposalId *big.Int) (DebookDAOProposal, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "getProposal", proposalId)

	if err != nil {
		return *new(DebookDAOProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(DebookDAOProposal)).(*DebookDAOProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns((address,bytes32,string,(uint64),(uint64),uint256,uint256,bool,bool,bool))
func (_DebookDAO *DebookDAOSession) GetProposal(proposalId *big.Int) (DebookDAOProposal, error) {
	return _DebookDAO.Contract.GetProposal(&_DebookDAO.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns((address,bytes32,string,(uint64),(uint64),uint256,uint256,bool,bool,bool))
func (_DebookDAO *DebookDAOCallerSession) GetProposal(proposalId *big.Int) (DebookDAOProposal, error) {
	return _DebookDAO.Contract.GetProposal(&_DebookDAO.CallOpts, proposalId)
}

// GetProposalDeadline is a free data retrieval call binding the contract method 0x8f363200.
//
// Solidity: function getProposalDeadline(uint256 proposalId) view returns(uint256)
func (_DebookDAO *DebookDAOCaller) GetProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "getProposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalDeadline is a free data retrieval call binding the contract method 0x8f363200.
//
// Solidity: function getProposalDeadline(uint256 proposalId) view returns(uint256)
func (_DebookDAO *DebookDAOSession) GetProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.GetProposalDeadline(&_DebookDAO.CallOpts, proposalId)
}

// GetProposalDeadline is a free data retrieval call binding the contract method 0x8f363200.
//
// Solidity: function getProposalDeadline(uint256 proposalId) view returns(uint256)
func (_DebookDAO *DebookDAOCallerSession) GetProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.GetProposalDeadline(&_DebookDAO.CallOpts, proposalId)
}

// GetProposalList is a free data retrieval call binding the contract method 0x346750f3.
//
// Solidity: function getProposalList() view returns(uint256[])
func (_DebookDAO *DebookDAOCaller) GetProposalList(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "getProposalList")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetProposalList is a free data retrieval call binding the contract method 0x346750f3.
//
// Solidity: function getProposalList() view returns(uint256[])
func (_DebookDAO *DebookDAOSession) GetProposalList() ([]*big.Int, error) {
	return _DebookDAO.Contract.GetProposalList(&_DebookDAO.CallOpts)
}

// GetProposalList is a free data retrieval call binding the contract method 0x346750f3.
//
// Solidity: function getProposalList() view returns(uint256[])
func (_DebookDAO *DebookDAOCallerSession) GetProposalList() ([]*big.Int, error) {
	return _DebookDAO.Contract.GetProposalList(&_DebookDAO.CallOpts)
}

// GetProposalVoteStart is a free data retrieval call binding the contract method 0xba017fd9.
//
// Solidity: function getProposalVoteStart(uint256 proposalId) view returns(uint256)
func (_DebookDAO *DebookDAOCaller) GetProposalVoteStart(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "getProposalVoteStart", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalVoteStart is a free data retrieval call binding the contract method 0xba017fd9.
//
// Solidity: function getProposalVoteStart(uint256 proposalId) view returns(uint256)
func (_DebookDAO *DebookDAOSession) GetProposalVoteStart(proposalId *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.GetProposalVoteStart(&_DebookDAO.CallOpts, proposalId)
}

// GetProposalVoteStart is a free data retrieval call binding the contract method 0xba017fd9.
//
// Solidity: function getProposalVoteStart(uint256 proposalId) view returns(uint256)
func (_DebookDAO *DebookDAOCallerSession) GetProposalVoteStart(proposalId *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.GetProposalVoteStart(&_DebookDAO.CallOpts, proposalId)
}

// HashProposal is a free data retrieval call binding the contract method 0xad959e90.
//
// Solidity: function hashProposal(address proposer, bytes32 title, uint256 createdAt) pure returns(uint256)
func (_DebookDAO *DebookDAOCaller) HashProposal(opts *bind.CallOpts, proposer common.Address, title [32]byte, createdAt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "hashProposal", proposer, title, createdAt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xad959e90.
//
// Solidity: function hashProposal(address proposer, bytes32 title, uint256 createdAt) pure returns(uint256)
func (_DebookDAO *DebookDAOSession) HashProposal(proposer common.Address, title [32]byte, createdAt *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.HashProposal(&_DebookDAO.CallOpts, proposer, title, createdAt)
}

// HashProposal is a free data retrieval call binding the contract method 0xad959e90.
//
// Solidity: function hashProposal(address proposer, bytes32 title, uint256 createdAt) pure returns(uint256)
func (_DebookDAO *DebookDAOCallerSession) HashProposal(proposer common.Address, title [32]byte, createdAt *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.HashProposal(&_DebookDAO.CallOpts, proposer, title, createdAt)
}

// HashVote is a free data retrieval call binding the contract method 0xeb2bbc2f.
//
// Solidity: function hashVote(address interactContract, uint256 proposalId, address voter, uint8 selection, string reason) pure returns(bytes32)
func (_DebookDAO *DebookDAOCaller) HashVote(opts *bind.CallOpts, interactContract common.Address, proposalId *big.Int, voter common.Address, selection uint8, reason string) ([32]byte, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "hashVote", interactContract, proposalId, voter, selection, reason)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashVote is a free data retrieval call binding the contract method 0xeb2bbc2f.
//
// Solidity: function hashVote(address interactContract, uint256 proposalId, address voter, uint8 selection, string reason) pure returns(bytes32)
func (_DebookDAO *DebookDAOSession) HashVote(interactContract common.Address, proposalId *big.Int, voter common.Address, selection uint8, reason string) ([32]byte, error) {
	return _DebookDAO.Contract.HashVote(&_DebookDAO.CallOpts, interactContract, proposalId, voter, selection, reason)
}

// HashVote is a free data retrieval call binding the contract method 0xeb2bbc2f.
//
// Solidity: function hashVote(address interactContract, uint256 proposalId, address voter, uint8 selection, string reason) pure returns(bytes32)
func (_DebookDAO *DebookDAOCallerSession) HashVote(interactContract common.Address, proposalId *big.Int, voter common.Address, selection uint8, reason string) ([32]byte, error) {
	return _DebookDAO.Contract.HashVote(&_DebookDAO.CallOpts, interactContract, proposalId, voter, selection, reason)
}

// IsVoterVoted is a free data retrieval call binding the contract method 0x2b8ef699.
//
// Solidity: function isVoterVoted(uint256 proposalId, address voterId) view returns(bool)
func (_DebookDAO *DebookDAOCaller) IsVoterVoted(opts *bind.CallOpts, proposalId *big.Int, voterId common.Address) (bool, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "isVoterVoted", proposalId, voterId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterVoted is a free data retrieval call binding the contract method 0x2b8ef699.
//
// Solidity: function isVoterVoted(uint256 proposalId, address voterId) view returns(bool)
func (_DebookDAO *DebookDAOSession) IsVoterVoted(proposalId *big.Int, voterId common.Address) (bool, error) {
	return _DebookDAO.Contract.IsVoterVoted(&_DebookDAO.CallOpts, proposalId, voterId)
}

// IsVoterVoted is a free data retrieval call binding the contract method 0x2b8ef699.
//
// Solidity: function isVoterVoted(uint256 proposalId, address voterId) view returns(bool)
func (_DebookDAO *DebookDAOCallerSession) IsVoterVoted(proposalId *big.Int, voterId common.Address) (bool, error) {
	return _DebookDAO.Contract.IsVoterVoted(&_DebookDAO.CallOpts, proposalId, voterId)
}

// ProposalsList is a free data retrieval call binding the contract method 0xe5957023.
//
// Solidity: function proposalsList(uint256 ) view returns(uint256)
func (_DebookDAO *DebookDAOCaller) ProposalsList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "proposalsList", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalsList is a free data retrieval call binding the contract method 0xe5957023.
//
// Solidity: function proposalsList(uint256 ) view returns(uint256)
func (_DebookDAO *DebookDAOSession) ProposalsList(arg0 *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.ProposalsList(&_DebookDAO.CallOpts, arg0)
}

// ProposalsList is a free data retrieval call binding the contract method 0xe5957023.
//
// Solidity: function proposalsList(uint256 ) view returns(uint256)
func (_DebookDAO *DebookDAOCallerSession) ProposalsList(arg0 *big.Int) (*big.Int, error) {
	return _DebookDAO.Contract.ProposalsList(&_DebookDAO.CallOpts, arg0)
}

// Sp is a free data retrieval call binding the contract method 0xe1dadf3b.
//
// Solidity: function sp() view returns(address)
func (_DebookDAO *DebookDAOCaller) Sp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "sp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sp is a free data retrieval call binding the contract method 0xe1dadf3b.
//
// Solidity: function sp() view returns(address)
func (_DebookDAO *DebookDAOSession) Sp() (common.Address, error) {
	return _DebookDAO.Contract.Sp(&_DebookDAO.CallOpts)
}

// Sp is a free data retrieval call binding the contract method 0xe1dadf3b.
//
// Solidity: function sp() view returns(address)
func (_DebookDAO *DebookDAOCallerSession) Sp() (common.Address, error) {
	return _DebookDAO.Contract.Sp(&_DebookDAO.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_DebookDAO *DebookDAOCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_DebookDAO *DebookDAOSession) State(proposalId *big.Int) (uint8, error) {
	return _DebookDAO.Contract.State(&_DebookDAO.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_DebookDAO *DebookDAOCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _DebookDAO.Contract.State(&_DebookDAO.CallOpts, proposalId)
}

// TeamMerkleRoot is a free data retrieval call binding the contract method 0x29140819.
//
// Solidity: function teamMerkleRoot() view returns(bytes32)
func (_DebookDAO *DebookDAOCaller) TeamMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DebookDAO.contract.Call(opts, &out, "teamMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TeamMerkleRoot is a free data retrieval call binding the contract method 0x29140819.
//
// Solidity: function teamMerkleRoot() view returns(bytes32)
func (_DebookDAO *DebookDAOSession) TeamMerkleRoot() ([32]byte, error) {
	return _DebookDAO.Contract.TeamMerkleRoot(&_DebookDAO.CallOpts)
}

// TeamMerkleRoot is a free data retrieval call binding the contract method 0x29140819.
//
// Solidity: function teamMerkleRoot() view returns(bytes32)
func (_DebookDAO *DebookDAOCallerSession) TeamMerkleRoot() ([32]byte, error) {
	return _DebookDAO.Contract.TeamMerkleRoot(&_DebookDAO.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns(uint256)
func (_DebookDAO *DebookDAOTransactor) Cancel(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "cancel", proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns(uint256)
func (_DebookDAO *DebookDAOSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _DebookDAO.Contract.Cancel(&_DebookDAO.TransactOpts, proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns(uint256)
func (_DebookDAO *DebookDAOTransactorSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _DebookDAO.Contract.Cancel(&_DebookDAO.TransactOpts, proposalId)
}

// CastVoteWithReasonBySig is a paid mutator transaction binding the contract method 0xcee87708.
//
// Solidity: function castVoteWithReasonBySig(uint256 proposalId, uint8 selection, string reason, uint8 v, bytes32 r, bytes32 s) returns()
func (_DebookDAO *DebookDAOTransactor) CastVoteWithReasonBySig(opts *bind.TransactOpts, proposalId *big.Int, selection uint8, reason string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "castVoteWithReasonBySig", proposalId, selection, reason, v, r, s)
}

// CastVoteWithReasonBySig is a paid mutator transaction binding the contract method 0xcee87708.
//
// Solidity: function castVoteWithReasonBySig(uint256 proposalId, uint8 selection, string reason, uint8 v, bytes32 r, bytes32 s) returns()
func (_DebookDAO *DebookDAOSession) CastVoteWithReasonBySig(proposalId *big.Int, selection uint8, reason string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DebookDAO.Contract.CastVoteWithReasonBySig(&_DebookDAO.TransactOpts, proposalId, selection, reason, v, r, s)
}

// CastVoteWithReasonBySig is a paid mutator transaction binding the contract method 0xcee87708.
//
// Solidity: function castVoteWithReasonBySig(uint256 proposalId, uint8 selection, string reason, uint8 v, bytes32 r, bytes32 s) returns()
func (_DebookDAO *DebookDAOTransactorSession) CastVoteWithReasonBySig(proposalId *big.Int, selection uint8, reason string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DebookDAO.Contract.CastVoteWithReasonBySig(&_DebookDAO.TransactOpts, proposalId, selection, reason, v, r, s)
}

// Exeute is a paid mutator transaction binding the contract method 0xc8044d20.
//
// Solidity: function exeute(uint256 proposalId, bytes32[] merkleProof) returns(uint256)
func (_DebookDAO *DebookDAOTransactor) Exeute(opts *bind.TransactOpts, proposalId *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "exeute", proposalId, merkleProof)
}

// Exeute is a paid mutator transaction binding the contract method 0xc8044d20.
//
// Solidity: function exeute(uint256 proposalId, bytes32[] merkleProof) returns(uint256)
func (_DebookDAO *DebookDAOSession) Exeute(proposalId *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _DebookDAO.Contract.Exeute(&_DebookDAO.TransactOpts, proposalId, merkleProof)
}

// Exeute is a paid mutator transaction binding the contract method 0xc8044d20.
//
// Solidity: function exeute(uint256 proposalId, bytes32[] merkleProof) returns(uint256)
func (_DebookDAO *DebookDAOTransactorSession) Exeute(proposalId *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _DebookDAO.Contract.Exeute(&_DebookDAO.TransactOpts, proposalId, merkleProof)
}

// Propose is a paid mutator transaction binding the contract method 0x042e9a43.
//
// Solidity: function propose(bytes32 title, string description) returns(uint256)
func (_DebookDAO *DebookDAOTransactor) Propose(opts *bind.TransactOpts, title [32]byte, description string) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "propose", title, description)
}

// Propose is a paid mutator transaction binding the contract method 0x042e9a43.
//
// Solidity: function propose(bytes32 title, string description) returns(uint256)
func (_DebookDAO *DebookDAOSession) Propose(title [32]byte, description string) (*types.Transaction, error) {
	return _DebookDAO.Contract.Propose(&_DebookDAO.TransactOpts, title, description)
}

// Propose is a paid mutator transaction binding the contract method 0x042e9a43.
//
// Solidity: function propose(bytes32 title, string description) returns(uint256)
func (_DebookDAO *DebookDAOTransactorSession) Propose(title [32]byte, description string) (*types.Transaction, error) {
	return _DebookDAO.Contract.Propose(&_DebookDAO.TransactOpts, title, description)
}

// SetStakePool is a paid mutator transaction binding the contract method 0x2b536a17.
//
// Solidity: function setStakePool(address _sp, bytes32[] merkleProof) returns()
func (_DebookDAO *DebookDAOTransactor) SetStakePool(opts *bind.TransactOpts, _sp common.Address, merkleProof [][32]byte) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "setStakePool", _sp, merkleProof)
}

// SetStakePool is a paid mutator transaction binding the contract method 0x2b536a17.
//
// Solidity: function setStakePool(address _sp, bytes32[] merkleProof) returns()
func (_DebookDAO *DebookDAOSession) SetStakePool(_sp common.Address, merkleProof [][32]byte) (*types.Transaction, error) {
	return _DebookDAO.Contract.SetStakePool(&_DebookDAO.TransactOpts, _sp, merkleProof)
}

// SetStakePool is a paid mutator transaction binding the contract method 0x2b536a17.
//
// Solidity: function setStakePool(address _sp, bytes32[] merkleProof) returns()
func (_DebookDAO *DebookDAOTransactorSession) SetStakePool(_sp common.Address, merkleProof [][32]byte) (*types.Transaction, error) {
	return _DebookDAO.Contract.SetStakePool(&_DebookDAO.TransactOpts, _sp, merkleProof)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 delay) returns()
func (_DebookDAO *DebookDAOTransactor) SetVotingDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "setVotingDelay", delay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 delay) returns()
func (_DebookDAO *DebookDAOSession) SetVotingDelay(delay *big.Int) (*types.Transaction, error) {
	return _DebookDAO.Contract.SetVotingDelay(&_DebookDAO.TransactOpts, delay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 delay) returns()
func (_DebookDAO *DebookDAOTransactorSession) SetVotingDelay(delay *big.Int) (*types.Transaction, error) {
	return _DebookDAO.Contract.SetVotingDelay(&_DebookDAO.TransactOpts, delay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 period) returns()
func (_DebookDAO *DebookDAOTransactor) SetVotingPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _DebookDAO.contract.Transact(opts, "setVotingPeriod", period)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 period) returns()
func (_DebookDAO *DebookDAOSession) SetVotingPeriod(period *big.Int) (*types.Transaction, error) {
	return _DebookDAO.Contract.SetVotingPeriod(&_DebookDAO.TransactOpts, period)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 period) returns()
func (_DebookDAO *DebookDAOTransactorSession) SetVotingPeriod(period *big.Int) (*types.Transaction, error) {
	return _DebookDAO.Contract.SetVotingPeriod(&_DebookDAO.TransactOpts, period)
}

// DebookDAOConsoleIterator is returned from FilterConsole and is used to iterate over the raw logs and unpacked data for Console events raised by the DebookDAO contract.
type DebookDAOConsoleIterator struct {
	Event *DebookDAOConsole // Event containing the contract specifics and raw log

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
func (it *DebookDAOConsoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebookDAOConsole)
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
		it.Event = new(DebookDAOConsole)
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
func (it *DebookDAOConsoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebookDAOConsoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebookDAOConsole represents a Console event raised by the DebookDAO contract.
type DebookDAOConsole struct {
	Title   string
	Message common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConsole is a free log retrieval operation binding the contract event 0x7c6b5c54d9b046b65a76e6926d069c12ee88e654e49e603870f118925096177d.
//
// Solidity: event Console(string title, address message)
func (_DebookDAO *DebookDAOFilterer) FilterConsole(opts *bind.FilterOpts) (*DebookDAOConsoleIterator, error) {

	logs, sub, err := _DebookDAO.contract.FilterLogs(opts, "Console")
	if err != nil {
		return nil, err
	}
	return &DebookDAOConsoleIterator{contract: _DebookDAO.contract, event: "Console", logs: logs, sub: sub}, nil
}

// WatchConsole is a free log subscription operation binding the contract event 0x7c6b5c54d9b046b65a76e6926d069c12ee88e654e49e603870f118925096177d.
//
// Solidity: event Console(string title, address message)
func (_DebookDAO *DebookDAOFilterer) WatchConsole(opts *bind.WatchOpts, sink chan<- *DebookDAOConsole) (event.Subscription, error) {

	logs, sub, err := _DebookDAO.contract.WatchLogs(opts, "Console")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebookDAOConsole)
				if err := _DebookDAO.contract.UnpackLog(event, "Console", log); err != nil {
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

// ParseConsole is a log parse operation binding the contract event 0x7c6b5c54d9b046b65a76e6926d069c12ee88e654e49e603870f118925096177d.
//
// Solidity: event Console(string title, address message)
func (_DebookDAO *DebookDAOFilterer) ParseConsole(log types.Log) (*DebookDAOConsole, error) {
	event := new(DebookDAOConsole)
	if err := _DebookDAO.contract.UnpackLog(event, "Console", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebookDAOProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the DebookDAO contract.
type DebookDAOProposalCanceledIterator struct {
	Event *DebookDAOProposalCanceled // Event containing the contract specifics and raw log

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
func (it *DebookDAOProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebookDAOProposalCanceled)
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
		it.Event = new(DebookDAOProposalCanceled)
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
func (it *DebookDAOProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebookDAOProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebookDAOProposalCanceled represents a ProposalCanceled event raised by the DebookDAO contract.
type DebookDAOProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*DebookDAOProposalCanceledIterator, error) {

	logs, sub, err := _DebookDAO.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &DebookDAOProposalCanceledIterator{contract: _DebookDAO.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *DebookDAOProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _DebookDAO.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebookDAOProposalCanceled)
				if err := _DebookDAO.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) ParseProposalCanceled(log types.Log) (*DebookDAOProposalCanceled, error) {
	event := new(DebookDAOProposalCanceled)
	if err := _DebookDAO.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebookDAOProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the DebookDAO contract.
type DebookDAOProposalCreatedIterator struct {
	Event *DebookDAOProposalCreated // Event containing the contract specifics and raw log

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
func (it *DebookDAOProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebookDAOProposalCreated)
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
		it.Event = new(DebookDAOProposalCreated)
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
func (it *DebookDAOProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebookDAOProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebookDAOProposalCreated represents a ProposalCreated event raised by the DebookDAO contract.
type DebookDAOProposalCreated struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*DebookDAOProposalCreatedIterator, error) {

	logs, sub, err := _DebookDAO.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &DebookDAOProposalCreatedIterator{contract: _DebookDAO.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *DebookDAOProposalCreated) (event.Subscription, error) {

	logs, sub, err := _DebookDAO.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebookDAOProposalCreated)
				if err := _DebookDAO.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) ParseProposalCreated(log types.Log) (*DebookDAOProposalCreated, error) {
	event := new(DebookDAOProposalCreated)
	if err := _DebookDAO.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebookDAOProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the DebookDAO contract.
type DebookDAOProposalExecutedIterator struct {
	Event *DebookDAOProposalExecuted // Event containing the contract specifics and raw log

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
func (it *DebookDAOProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebookDAOProposalExecuted)
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
		it.Event = new(DebookDAOProposalExecuted)
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
func (it *DebookDAOProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebookDAOProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebookDAOProposalExecuted represents a ProposalExecuted event raised by the DebookDAO contract.
type DebookDAOProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*DebookDAOProposalExecutedIterator, error) {

	logs, sub, err := _DebookDAO.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &DebookDAOProposalExecutedIterator{contract: _DebookDAO.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *DebookDAOProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _DebookDAO.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebookDAOProposalExecuted)
				if err := _DebookDAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_DebookDAO *DebookDAOFilterer) ParseProposalExecuted(log types.Log) (*DebookDAOProposalExecuted, error) {
	event := new(DebookDAOProposalExecuted)
	if err := _DebookDAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebookDAOVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the DebookDAO contract.
type DebookDAOVotedIterator struct {
	Event *DebookDAOVoted // Event containing the contract specifics and raw log

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
func (it *DebookDAOVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebookDAOVoted)
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
		it.Event = new(DebookDAOVoted)
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
func (it *DebookDAOVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebookDAOVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebookDAOVoted represents a Voted event raised by the DebookDAO contract.
type DebookDAOVoted struct {
	Voter      common.Address ""
	ProposalId *big.Int
	Selection  uint8
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x8a8fe2a34c02d0b2ebedec236a11ddb52f4fec099b7a3beee8f137f388d63c73.
//
// Solidity: event Voted(address indexed voter, uint256 proposalId, uint8 selection, string reason)
func (_DebookDAO *DebookDAOFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address) (*DebookDAOVotedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DebookDAO.contract.FilterLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return &DebookDAOVotedIterator{contract: _DebookDAO.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x8a8fe2a34c02d0b2ebedec236a11ddb52f4fec099b7a3beee8f137f388d63c73.
//
// Solidity: event Voted(address indexed voter, uint256 proposalId, uint8 selection, string reason)
func (_DebookDAO *DebookDAOFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *DebookDAOVoted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DebookDAO.contract.WatchLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebookDAOVoted)
				if err := _DebookDAO.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x8a8fe2a34c02d0b2ebedec236a11ddb52f4fec099b7a3beee8f137f388d63c73.
//
// Solidity: event Voted(address indexed voter, uint256 proposalId, uint8 selection, string reason)
func (_DebookDAO *DebookDAOFilterer) ParseVoted(log types.Log) (*DebookDAOVoted, error) {
	event := new(DebookDAOVoted)
	if err := _DebookDAO.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
