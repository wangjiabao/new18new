// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// BuyMetaData contains all meta data concerning the Buy contract.
var BuyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersAmountByIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usersAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BuyABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyMetaData.ABI instead.
var BuyABI = BuyMetaData.ABI

// Buy is an auto generated Go binding around an Ethereum contract.
type Buy struct {
	BuyCaller     // Read-only binding to the contract
	BuyTransactor // Write-only binding to the contract
	BuyFilterer   // Log filterer for contract events
}

// BuyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuySession struct {
	Contract     *Buy              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyCallerSession struct {
	Contract *BuyCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyTransactorSession struct {
	Contract     *BuyTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyRaw struct {
	Contract *Buy // Generic contract binding to access the raw methods on
}

// BuyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyCallerRaw struct {
	Contract *BuyCaller // Generic read-only contract binding to access the raw methods on
}

// BuyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyTransactorRaw struct {
	Contract *BuyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuy creates a new instance of Buy, bound to a specific deployed contract.
func NewBuy(address common.Address, backend bind.ContractBackend) (*Buy, error) {
	contract, err := bindBuy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buy{BuyCaller: BuyCaller{contract: contract}, BuyTransactor: BuyTransactor{contract: contract}, BuyFilterer: BuyFilterer{contract: contract}}, nil
}

// NewBuyCaller creates a new read-only instance of Buy, bound to a specific deployed contract.
func NewBuyCaller(address common.Address, caller bind.ContractCaller) (*BuyCaller, error) {
	contract, err := bindBuy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyCaller{contract: contract}, nil
}

// NewBuyTransactor creates a new write-only instance of Buy, bound to a specific deployed contract.
func NewBuyTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyTransactor, error) {
	contract, err := bindBuy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyTransactor{contract: contract}, nil
}

// NewBuyFilterer creates a new log filterer instance of Buy, bound to a specific deployed contract.
func NewBuyFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyFilterer, error) {
	contract, err := bindBuy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyFilterer{contract: contract}, nil
}

// bindBuy binds a generic wrapper to an already deployed contract.
func bindBuy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.BuyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transact(opts, method, params...)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_Buy *BuyCaller) GetUserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getUserLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_Buy *BuySession) GetUserLength() (*big.Int, error) {
	return _Buy.Contract.GetUserLength(&_Buy.CallOpts)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_Buy *BuyCallerSession) GetUserLength() (*big.Int, error) {
	return _Buy.Contract.GetUserLength(&_Buy.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Buy *BuyCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Buy *BuySession) GetUsers() ([]common.Address, error) {
	return _Buy.Contract.GetUsers(&_Buy.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Buy *BuyCallerSession) GetUsers() ([]common.Address, error) {
	return _Buy.Contract.GetUsers(&_Buy.CallOpts)
}

// GetUsersAmountByIndex is a free data retrieval call binding the contract method 0xadaf9e71.
//
// Solidity: function getUsersAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Buy *BuyCaller) GetUsersAmountByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getUsersAmountByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUsersAmountByIndex is a free data retrieval call binding the contract method 0xadaf9e71.
//
// Solidity: function getUsersAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Buy *BuySession) GetUsersAmountByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Buy.Contract.GetUsersAmountByIndex(&_Buy.CallOpts, startIndex, endIndex)
}

// GetUsersAmountByIndex is a free data retrieval call binding the contract method 0xadaf9e71.
//
// Solidity: function getUsersAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Buy *BuyCallerSession) GetUsersAmountByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Buy.Contract.GetUsersAmountByIndex(&_Buy.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_Buy *BuyCaller) GetUsersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getUsersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_Buy *BuySession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _Buy.Contract.GetUsersByIndex(&_Buy.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_Buy *BuyCallerSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _Buy.Contract.GetUsersByIndex(&_Buy.CallOpts, startIndex, endIndex)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Buy *BuyCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Buy *BuySession) Usdt() (common.Address, error) {
	return _Buy.Contract.Usdt(&_Buy.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Buy *BuyCallerSession) Usdt() (common.Address, error) {
	return _Buy.Contract.Usdt(&_Buy.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Buy *BuyCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Buy *BuySession) Users(arg0 *big.Int) (common.Address, error) {
	return _Buy.Contract.Users(&_Buy.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Buy *BuyCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Buy.Contract.Users(&_Buy.CallOpts, arg0)
}

// UsersAmount is a free data retrieval call binding the contract method 0x0963b51e.
//
// Solidity: function usersAmount(uint256 ) view returns(uint256)
func (_Buy *BuyCaller) UsersAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "usersAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersAmount is a free data retrieval call binding the contract method 0x0963b51e.
//
// Solidity: function usersAmount(uint256 ) view returns(uint256)
func (_Buy *BuySession) UsersAmount(arg0 *big.Int) (*big.Int, error) {
	return _Buy.Contract.UsersAmount(&_Buy.CallOpts, arg0)
}

// UsersAmount is a free data retrieval call binding the contract method 0x0963b51e.
//
// Solidity: function usersAmount(uint256 ) view returns(uint256)
func (_Buy *BuyCallerSession) UsersAmount(arg0 *big.Int) (*big.Int, error) {
	return _Buy.Contract.UsersAmount(&_Buy.CallOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 num) returns()
func (_Buy *BuyTransactor) Buy(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "buy", num)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 num) returns()
func (_Buy *BuySession) Buy(num *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.Buy(&_Buy.TransactOpts, num)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 num) returns()
func (_Buy *BuyTransactorSession) Buy(num *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.Buy(&_Buy.TransactOpts, num)
}
