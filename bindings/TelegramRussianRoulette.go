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
	_ = abi.ConvertType
)

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bettingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumBet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_revenueBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burnBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_revenueWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"tgChatId\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"playerIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"tgChatId\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"tgChatId\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"playerIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Loss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"tgChatId\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Revenue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"tgChatId\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"playerIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Win\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"abortAllGames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"_tgChatId\",\"type\":\"int64\"}],\"name\":\"abortGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeTgGroups\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bettingToken\",\"outputs\":[{\"internalType\":\"contractBulletGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"_tgChatId\",\"type\":\"int64\"},{\"internalType\":\"uint16\",\"name\":\"_loser\",\"type\":\"uint16\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"name\":\"endGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"revolverSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashedBulletChamberIndex\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"inProgress\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"loser\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"_tgChatId\",\"type\":\"int64\"}],\"name\":\"isGameInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"_tgChatId\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"_revolverSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hashedBulletChamberIndex\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_players\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_bets\",\"type\":\"uint256[]\"}],\"name\":\"newGame\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// ActiveTgGroups is a free data retrieval call binding the contract method 0xd057fc1f.
//
// Solidity: function activeTgGroups(uint256 ) view returns(int64)
func (_Bindings *BindingsCaller) ActiveTgGroups(opts *bind.CallOpts, arg0 *big.Int) (int64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "activeTgGroups", arg0)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// ActiveTgGroups is a free data retrieval call binding the contract method 0xd057fc1f.
//
// Solidity: function activeTgGroups(uint256 ) view returns(int64)
func (_Bindings *BindingsSession) ActiveTgGroups(arg0 *big.Int) (int64, error) {
	return _Bindings.Contract.ActiveTgGroups(&_Bindings.CallOpts, arg0)
}

// ActiveTgGroups is a free data retrieval call binding the contract method 0xd057fc1f.
//
// Solidity: function activeTgGroups(uint256 ) view returns(int64)
func (_Bindings *BindingsCallerSession) ActiveTgGroups(arg0 *big.Int) (int64, error) {
	return _Bindings.Contract.ActiveTgGroups(&_Bindings.CallOpts, arg0)
}

// BettingToken is a free data retrieval call binding the contract method 0x43425e88.
//
// Solidity: function bettingToken() view returns(address)
func (_Bindings *BindingsCaller) BettingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "bettingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BettingToken is a free data retrieval call binding the contract method 0x43425e88.
//
// Solidity: function bettingToken() view returns(address)
func (_Bindings *BindingsSession) BettingToken() (common.Address, error) {
	return _Bindings.Contract.BettingToken(&_Bindings.CallOpts)
}

// BettingToken is a free data retrieval call binding the contract method 0x43425e88.
//
// Solidity: function bettingToken() view returns(address)
func (_Bindings *BindingsCallerSession) BettingToken() (common.Address, error) {
	return _Bindings.Contract.BettingToken(&_Bindings.CallOpts)
}

// BurnBps is a free data retrieval call binding the contract method 0x53deb3d6.
//
// Solidity: function burnBps() view returns(uint256)
func (_Bindings *BindingsCaller) BurnBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "burnBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnBps is a free data retrieval call binding the contract method 0x53deb3d6.
//
// Solidity: function burnBps() view returns(uint256)
func (_Bindings *BindingsSession) BurnBps() (*big.Int, error) {
	return _Bindings.Contract.BurnBps(&_Bindings.CallOpts)
}

// BurnBps is a free data retrieval call binding the contract method 0x53deb3d6.
//
// Solidity: function burnBps() view returns(uint256)
func (_Bindings *BindingsCallerSession) BurnBps() (*big.Int, error) {
	return _Bindings.Contract.BurnBps(&_Bindings.CallOpts)
}

// Games is a free data retrieval call binding the contract method 0xf3619716.
//
// Solidity: function games(int64 ) view returns(uint256 revolverSize, uint256 minBet, bytes32 hashedBulletChamberIndex, bool inProgress, uint16 loser)
func (_Bindings *BindingsCaller) Games(opts *bind.CallOpts, arg0 int64) (struct {
	RevolverSize             *big.Int
	MinBet                   *big.Int
	HashedBulletChamberIndex [32]byte
	InProgress               bool
	Loser                    uint16
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		RevolverSize             *big.Int
		MinBet                   *big.Int
		HashedBulletChamberIndex [32]byte
		InProgress               bool
		Loser                    uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RevolverSize = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinBet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HashedBulletChamberIndex = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.InProgress = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Loser = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0xf3619716.
//
// Solidity: function games(int64 ) view returns(uint256 revolverSize, uint256 minBet, bytes32 hashedBulletChamberIndex, bool inProgress, uint16 loser)
func (_Bindings *BindingsSession) Games(arg0 int64) (struct {
	RevolverSize             *big.Int
	MinBet                   *big.Int
	HashedBulletChamberIndex [32]byte
	InProgress               bool
	Loser                    uint16
}, error) {
	return _Bindings.Contract.Games(&_Bindings.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xf3619716.
//
// Solidity: function games(int64 ) view returns(uint256 revolverSize, uint256 minBet, bytes32 hashedBulletChamberIndex, bool inProgress, uint16 loser)
func (_Bindings *BindingsCallerSession) Games(arg0 int64) (struct {
	RevolverSize             *big.Int
	MinBet                   *big.Int
	HashedBulletChamberIndex [32]byte
	InProgress               bool
	Loser                    uint16
}, error) {
	return _Bindings.Contract.Games(&_Bindings.CallOpts, arg0)
}

// IsGameInProgress is a free data retrieval call binding the contract method 0x63c42460.
//
// Solidity: function isGameInProgress(int64 _tgChatId) view returns(bool)
func (_Bindings *BindingsCaller) IsGameInProgress(opts *bind.CallOpts, _tgChatId int64) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isGameInProgress", _tgChatId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGameInProgress is a free data retrieval call binding the contract method 0x63c42460.
//
// Solidity: function isGameInProgress(int64 _tgChatId) view returns(bool)
func (_Bindings *BindingsSession) IsGameInProgress(_tgChatId int64) (bool, error) {
	return _Bindings.Contract.IsGameInProgress(&_Bindings.CallOpts, _tgChatId)
}

// IsGameInProgress is a free data retrieval call binding the contract method 0x63c42460.
//
// Solidity: function isGameInProgress(int64 _tgChatId) view returns(bool)
func (_Bindings *BindingsCallerSession) IsGameInProgress(_tgChatId int64) (bool, error) {
	return _Bindings.Contract.IsGameInProgress(&_Bindings.CallOpts, _tgChatId)
}

// MinimumBet is a free data retrieval call binding the contract method 0xc38a8afd.
//
// Solidity: function minimumBet() view returns(uint256)
func (_Bindings *BindingsCaller) MinimumBet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "minimumBet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumBet is a free data retrieval call binding the contract method 0xc38a8afd.
//
// Solidity: function minimumBet() view returns(uint256)
func (_Bindings *BindingsSession) MinimumBet() (*big.Int, error) {
	return _Bindings.Contract.MinimumBet(&_Bindings.CallOpts)
}

// MinimumBet is a free data retrieval call binding the contract method 0xc38a8afd.
//
// Solidity: function minimumBet() view returns(uint256)
func (_Bindings *BindingsCallerSession) MinimumBet() (*big.Int, error) {
	return _Bindings.Contract.MinimumBet(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// RevenueBps is a free data retrieval call binding the contract method 0xff08aa49.
//
// Solidity: function revenueBps() view returns(uint256)
func (_Bindings *BindingsCaller) RevenueBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "revenueBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevenueBps is a free data retrieval call binding the contract method 0xff08aa49.
//
// Solidity: function revenueBps() view returns(uint256)
func (_Bindings *BindingsSession) RevenueBps() (*big.Int, error) {
	return _Bindings.Contract.RevenueBps(&_Bindings.CallOpts)
}

// RevenueBps is a free data retrieval call binding the contract method 0xff08aa49.
//
// Solidity: function revenueBps() view returns(uint256)
func (_Bindings *BindingsCallerSession) RevenueBps() (*big.Int, error) {
	return _Bindings.Contract.RevenueBps(&_Bindings.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_Bindings *BindingsCaller) RevenueWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "revenueWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_Bindings *BindingsSession) RevenueWallet() (common.Address, error) {
	return _Bindings.Contract.RevenueWallet(&_Bindings.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_Bindings *BindingsCallerSession) RevenueWallet() (common.Address, error) {
	return _Bindings.Contract.RevenueWallet(&_Bindings.CallOpts)
}

// AbortAllGames is a paid mutator transaction binding the contract method 0x65816731.
//
// Solidity: function abortAllGames() returns()
func (_Bindings *BindingsTransactor) AbortAllGames(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "abortAllGames")
}

// AbortAllGames is a paid mutator transaction binding the contract method 0x65816731.
//
// Solidity: function abortAllGames() returns()
func (_Bindings *BindingsSession) AbortAllGames() (*types.Transaction, error) {
	return _Bindings.Contract.AbortAllGames(&_Bindings.TransactOpts)
}

// AbortAllGames is a paid mutator transaction binding the contract method 0x65816731.
//
// Solidity: function abortAllGames() returns()
func (_Bindings *BindingsTransactorSession) AbortAllGames() (*types.Transaction, error) {
	return _Bindings.Contract.AbortAllGames(&_Bindings.TransactOpts)
}

// AbortGame is a paid mutator transaction binding the contract method 0xcb99e91f.
//
// Solidity: function abortGame(int64 _tgChatId) returns()
func (_Bindings *BindingsTransactor) AbortGame(opts *bind.TransactOpts, _tgChatId int64) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "abortGame", _tgChatId)
}

// AbortGame is a paid mutator transaction binding the contract method 0xcb99e91f.
//
// Solidity: function abortGame(int64 _tgChatId) returns()
func (_Bindings *BindingsSession) AbortGame(_tgChatId int64) (*types.Transaction, error) {
	return _Bindings.Contract.AbortGame(&_Bindings.TransactOpts, _tgChatId)
}

// AbortGame is a paid mutator transaction binding the contract method 0xcb99e91f.
//
// Solidity: function abortGame(int64 _tgChatId) returns()
func (_Bindings *BindingsTransactorSession) AbortGame(_tgChatId int64) (*types.Transaction, error) {
	return _Bindings.Contract.AbortGame(&_Bindings.TransactOpts, _tgChatId)
}

// EndGame is a paid mutator transaction binding the contract method 0x88b7904d.
//
// Solidity: function endGame(int64 _tgChatId, uint16 _loser, string[] ) returns()
func (_Bindings *BindingsTransactor) EndGame(opts *bind.TransactOpts, _tgChatId int64, _loser uint16, arg2 []string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "endGame", _tgChatId, _loser, arg2)
}

// EndGame is a paid mutator transaction binding the contract method 0x88b7904d.
//
// Solidity: function endGame(int64 _tgChatId, uint16 _loser, string[] ) returns()
func (_Bindings *BindingsSession) EndGame(_tgChatId int64, _loser uint16, arg2 []string) (*types.Transaction, error) {
	return _Bindings.Contract.EndGame(&_Bindings.TransactOpts, _tgChatId, _loser, arg2)
}

// EndGame is a paid mutator transaction binding the contract method 0x88b7904d.
//
// Solidity: function endGame(int64 _tgChatId, uint16 _loser, string[] ) returns()
func (_Bindings *BindingsTransactorSession) EndGame(_tgChatId int64, _loser uint16, arg2 []string) (*types.Transaction, error) {
	return _Bindings.Contract.EndGame(&_Bindings.TransactOpts, _tgChatId, _loser, arg2)
}

// NewGame is a paid mutator transaction binding the contract method 0xceb7ec04.
//
// Solidity: function newGame(int64 _tgChatId, uint256 _revolverSize, uint256 _minBet, bytes32 _hashedBulletChamberIndex, address[] _players, uint256[] _bets) returns(uint256[])
func (_Bindings *BindingsTransactor) NewGame(opts *bind.TransactOpts, _tgChatId int64, _revolverSize *big.Int, _minBet *big.Int, _hashedBulletChamberIndex [32]byte, _players []common.Address, _bets []*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "newGame", _tgChatId, _revolverSize, _minBet, _hashedBulletChamberIndex, _players, _bets)
}

// NewGame is a paid mutator transaction binding the contract method 0xceb7ec04.
//
// Solidity: function newGame(int64 _tgChatId, uint256 _revolverSize, uint256 _minBet, bytes32 _hashedBulletChamberIndex, address[] _players, uint256[] _bets) returns(uint256[])
func (_Bindings *BindingsSession) NewGame(_tgChatId int64, _revolverSize *big.Int, _minBet *big.Int, _hashedBulletChamberIndex [32]byte, _players []common.Address, _bets []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.NewGame(&_Bindings.TransactOpts, _tgChatId, _revolverSize, _minBet, _hashedBulletChamberIndex, _players, _bets)
}

// NewGame is a paid mutator transaction binding the contract method 0xceb7ec04.
//
// Solidity: function newGame(int64 _tgChatId, uint256 _revolverSize, uint256 _minBet, bytes32 _hashedBulletChamberIndex, address[] _players, uint256[] _bets) returns(uint256[])
func (_Bindings *BindingsTransactorSession) NewGame(_tgChatId int64, _revolverSize *big.Int, _minBet *big.Int, _hashedBulletChamberIndex [32]byte, _players []common.Address, _bets []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.NewGame(&_Bindings.TransactOpts, _tgChatId, _revolverSize, _minBet, _hashedBulletChamberIndex, _players, _bets)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// BindingsBetIterator is returned from FilterBet and is used to iterate over the raw logs and unpacked data for Bet events raised by the Bindings contract.
type BindingsBetIterator struct {
	Event *BindingsBet // Event containing the contract specifics and raw log

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
func (it *BindingsBetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBet)
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
		it.Event = new(BindingsBet)
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
func (it *BindingsBetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBet represents a Bet event raised by the Bindings contract.
type BindingsBet struct {
	TgChatId    int64
	Player      common.Address
	PlayerIndex uint16
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBet is a free log retrieval operation binding the contract event 0x4ca2b6f8214bfec8b3a7c06707618645a8e77d171b22a4eba1d8811fdc30bfdb.
//
// Solidity: event Bet(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) FilterBet(opts *bind.FilterOpts) (*BindingsBetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Bet")
	if err != nil {
		return nil, err
	}
	return &BindingsBetIterator{contract: _Bindings.contract, event: "Bet", logs: logs, sub: sub}, nil
}

// WatchBet is a free log subscription operation binding the contract event 0x4ca2b6f8214bfec8b3a7c06707618645a8e77d171b22a4eba1d8811fdc30bfdb.
//
// Solidity: event Bet(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) WatchBet(opts *bind.WatchOpts, sink chan<- *BindingsBet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Bet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBet)
				if err := _Bindings.contract.UnpackLog(event, "Bet", log); err != nil {
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

// ParseBet is a log parse operation binding the contract event 0x4ca2b6f8214bfec8b3a7c06707618645a8e77d171b22a4eba1d8811fdc30bfdb.
//
// Solidity: event Bet(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) ParseBet(log types.Log) (*BindingsBet, error) {
	event := new(BindingsBet)
	if err := _Bindings.contract.UnpackLog(event, "Bet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Bindings contract.
type BindingsBurnIterator struct {
	Event *BindingsBurn // Event containing the contract specifics and raw log

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
func (it *BindingsBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBurn)
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
		it.Event = new(BindingsBurn)
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
func (it *BindingsBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBurn represents a Burn event raised by the Bindings contract.
type BindingsBurn struct {
	TgChatId int64
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xbc03807cbae975b0551ce6caa7b86a1ff549b347e16440847a8c03140f59c27c.
//
// Solidity: event Burn(int64 tgChatId, uint256 amount)
func (_Bindings *BindingsFilterer) FilterBurn(opts *bind.FilterOpts) (*BindingsBurnIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &BindingsBurnIterator{contract: _Bindings.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xbc03807cbae975b0551ce6caa7b86a1ff549b347e16440847a8c03140f59c27c.
//
// Solidity: event Burn(int64 tgChatId, uint256 amount)
func (_Bindings *BindingsFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BindingsBurn) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBurn)
				if err := _Bindings.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xbc03807cbae975b0551ce6caa7b86a1ff549b347e16440847a8c03140f59c27c.
//
// Solidity: event Burn(int64 tgChatId, uint256 amount)
func (_Bindings *BindingsFilterer) ParseBurn(log types.Log) (*BindingsBurn, error) {
	event := new(BindingsBurn)
	if err := _Bindings.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsLossIterator is returned from FilterLoss and is used to iterate over the raw logs and unpacked data for Loss events raised by the Bindings contract.
type BindingsLossIterator struct {
	Event *BindingsLoss // Event containing the contract specifics and raw log

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
func (it *BindingsLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsLoss)
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
		it.Event = new(BindingsLoss)
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
func (it *BindingsLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsLoss represents a Loss event raised by the Bindings contract.
type BindingsLoss struct {
	TgChatId    int64
	Player      common.Address
	PlayerIndex uint16
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLoss is a free log retrieval operation binding the contract event 0x0955ed2d49ed21e6c630b38b8f1c1bb2803067844901ad0184d41a34b88bd64a.
//
// Solidity: event Loss(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) FilterLoss(opts *bind.FilterOpts) (*BindingsLossIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Loss")
	if err != nil {
		return nil, err
	}
	return &BindingsLossIterator{contract: _Bindings.contract, event: "Loss", logs: logs, sub: sub}, nil
}

// WatchLoss is a free log subscription operation binding the contract event 0x0955ed2d49ed21e6c630b38b8f1c1bb2803067844901ad0184d41a34b88bd64a.
//
// Solidity: event Loss(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) WatchLoss(opts *bind.WatchOpts, sink chan<- *BindingsLoss) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Loss")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsLoss)
				if err := _Bindings.contract.UnpackLog(event, "Loss", log); err != nil {
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

// ParseLoss is a log parse operation binding the contract event 0x0955ed2d49ed21e6c630b38b8f1c1bb2803067844901ad0184d41a34b88bd64a.
//
// Solidity: event Loss(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) ParseLoss(log types.Log) (*BindingsLoss, error) {
	event := new(BindingsLoss)
	if err := _Bindings.contract.UnpackLog(event, "Loss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
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
		it.Event = new(BindingsOwnershipTransferred)
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
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRevenueIterator is returned from FilterRevenue and is used to iterate over the raw logs and unpacked data for Revenue events raised by the Bindings contract.
type BindingsRevenueIterator struct {
	Event *BindingsRevenue // Event containing the contract specifics and raw log

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
func (it *BindingsRevenueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRevenue)
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
		it.Event = new(BindingsRevenue)
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
func (it *BindingsRevenueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRevenueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRevenue represents a Revenue event raised by the Bindings contract.
type BindingsRevenue struct {
	TgChatId int64
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRevenue is a free log retrieval operation binding the contract event 0x0f771b5d5a6b02378d0d1a6b6b371ac1e69759fb677e46109ae1bb55167ea7ad.
//
// Solidity: event Revenue(int64 tgChatId, uint256 amount)
func (_Bindings *BindingsFilterer) FilterRevenue(opts *bind.FilterOpts) (*BindingsRevenueIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Revenue")
	if err != nil {
		return nil, err
	}
	return &BindingsRevenueIterator{contract: _Bindings.contract, event: "Revenue", logs: logs, sub: sub}, nil
}

// WatchRevenue is a free log subscription operation binding the contract event 0x0f771b5d5a6b02378d0d1a6b6b371ac1e69759fb677e46109ae1bb55167ea7ad.
//
// Solidity: event Revenue(int64 tgChatId, uint256 amount)
func (_Bindings *BindingsFilterer) WatchRevenue(opts *bind.WatchOpts, sink chan<- *BindingsRevenue) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Revenue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRevenue)
				if err := _Bindings.contract.UnpackLog(event, "Revenue", log); err != nil {
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

// ParseRevenue is a log parse operation binding the contract event 0x0f771b5d5a6b02378d0d1a6b6b371ac1e69759fb677e46109ae1bb55167ea7ad.
//
// Solidity: event Revenue(int64 tgChatId, uint256 amount)
func (_Bindings *BindingsFilterer) ParseRevenue(log types.Log) (*BindingsRevenue, error) {
	event := new(BindingsRevenue)
	if err := _Bindings.contract.UnpackLog(event, "Revenue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsWinIterator is returned from FilterWin and is used to iterate over the raw logs and unpacked data for Win events raised by the Bindings contract.
type BindingsWinIterator struct {
	Event *BindingsWin // Event containing the contract specifics and raw log

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
func (it *BindingsWinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsWin)
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
		it.Event = new(BindingsWin)
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
func (it *BindingsWinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsWinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsWin represents a Win event raised by the Bindings contract.
type BindingsWin struct {
	TgChatId    int64
	Player      common.Address
	PlayerIndex uint16
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWin is a free log retrieval operation binding the contract event 0x6b5ed972057bb3f9c6b7b2ea6350bf7abde0e0c5f8a765c5dde8402bb2b6efd3.
//
// Solidity: event Win(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) FilterWin(opts *bind.FilterOpts) (*BindingsWinIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Win")
	if err != nil {
		return nil, err
	}
	return &BindingsWinIterator{contract: _Bindings.contract, event: "Win", logs: logs, sub: sub}, nil
}

// WatchWin is a free log subscription operation binding the contract event 0x6b5ed972057bb3f9c6b7b2ea6350bf7abde0e0c5f8a765c5dde8402bb2b6efd3.
//
// Solidity: event Win(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) WatchWin(opts *bind.WatchOpts, sink chan<- *BindingsWin) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Win")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsWin)
				if err := _Bindings.contract.UnpackLog(event, "Win", log); err != nil {
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

// ParseWin is a log parse operation binding the contract event 0x6b5ed972057bb3f9c6b7b2ea6350bf7abde0e0c5f8a765c5dde8402bb2b6efd3.
//
// Solidity: event Win(int64 tgChatId, address player, uint16 playerIndex, uint256 amount)
func (_Bindings *BindingsFilterer) ParseWin(log types.Log) (*BindingsWin, error) {
	event := new(BindingsWin)
	if err := _Bindings.contract.UnpackLog(event, "Win", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
