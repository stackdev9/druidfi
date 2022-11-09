// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package druidfinance

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

// DruidfinanceMetaData contains all meta data concerning the Druidfinance contract.
var DruidfinanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"_transferm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"bq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"first\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sec\",\"type\":\"address\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"m_b_multi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addres\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"m_b_s\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addres\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"m_w_s\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taxFee\",\"type\":\"uint256\"}],\"name\":\"setTaxFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"wted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DruidfinanceABI is the input ABI used to generate the binding from.
// Deprecated: Use DruidfinanceMetaData.ABI instead.
var DruidfinanceABI = DruidfinanceMetaData.ABI

// Druidfinance is an auto generated Go binding around an Ethereum contract.
type Druidfinance struct {
	DruidfinanceCaller     // Read-only binding to the contract
	DruidfinanceTransactor // Write-only binding to the contract
	DruidfinanceFilterer   // Log filterer for contract events
}

// DruidfinanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DruidfinanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DruidfinanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DruidfinanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DruidfinanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DruidfinanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DruidfinanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DruidfinanceSession struct {
	Contract     *Druidfinance     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DruidfinanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DruidfinanceCallerSession struct {
	Contract *DruidfinanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DruidfinanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DruidfinanceTransactorSession struct {
	Contract     *DruidfinanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DruidfinanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DruidfinanceRaw struct {
	Contract *Druidfinance // Generic contract binding to access the raw methods on
}

// DruidfinanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DruidfinanceCallerRaw struct {
	Contract *DruidfinanceCaller // Generic read-only contract binding to access the raw methods on
}

// DruidfinanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DruidfinanceTransactorRaw struct {
	Contract *DruidfinanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDruidfinance creates a new instance of Druidfinance, bound to a specific deployed contract.
func NewDruidfinance(address common.Address, backend bind.ContractBackend) (*Druidfinance, error) {
	contract, err := bindDruidfinance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Druidfinance{DruidfinanceCaller: DruidfinanceCaller{contract: contract}, DruidfinanceTransactor: DruidfinanceTransactor{contract: contract}, DruidfinanceFilterer: DruidfinanceFilterer{contract: contract}}, nil
}

// NewDruidfinanceCaller creates a new read-only instance of Druidfinance, bound to a specific deployed contract.
func NewDruidfinanceCaller(address common.Address, caller bind.ContractCaller) (*DruidfinanceCaller, error) {
	contract, err := bindDruidfinance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DruidfinanceCaller{contract: contract}, nil
}

// NewDruidfinanceTransactor creates a new write-only instance of Druidfinance, bound to a specific deployed contract.
func NewDruidfinanceTransactor(address common.Address, transactor bind.ContractTransactor) (*DruidfinanceTransactor, error) {
	contract, err := bindDruidfinance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DruidfinanceTransactor{contract: contract}, nil
}

// NewDruidfinanceFilterer creates a new log filterer instance of Druidfinance, bound to a specific deployed contract.
func NewDruidfinanceFilterer(address common.Address, filterer bind.ContractFilterer) (*DruidfinanceFilterer, error) {
	contract, err := bindDruidfinance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DruidfinanceFilterer{contract: contract}, nil
}

// bindDruidfinance binds a generic wrapper to an already deployed contract.
func bindDruidfinance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DruidfinanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Druidfinance *DruidfinanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Druidfinance.Contract.DruidfinanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Druidfinance *DruidfinanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Druidfinance.Contract.DruidfinanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Druidfinance *DruidfinanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Druidfinance.Contract.DruidfinanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Druidfinance *DruidfinanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Druidfinance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Druidfinance *DruidfinanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Druidfinance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Druidfinance *DruidfinanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Druidfinance.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Druidfinance *DruidfinanceSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Druidfinance.Contract.Allowance(&_Druidfinance.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Druidfinance.Contract.Allowance(&_Druidfinance.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Druidfinance *DruidfinanceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Druidfinance.Contract.BalanceOf(&_Druidfinance.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Druidfinance.Contract.BalanceOf(&_Druidfinance.CallOpts, owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Druidfinance *DruidfinanceSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Druidfinance.Contract.Balances(&_Druidfinance.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Druidfinance.Contract.Balances(&_Druidfinance.CallOpts, arg0)
}

// Bq is a free data retrieval call binding the contract method 0x2605a080.
//
// Solidity: function bq(address owner) view returns(bool)
func (_Druidfinance *DruidfinanceCaller) Bq(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "bq", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bq is a free data retrieval call binding the contract method 0x2605a080.
//
// Solidity: function bq(address owner) view returns(bool)
func (_Druidfinance *DruidfinanceSession) Bq(owner common.Address) (bool, error) {
	return _Druidfinance.Contract.Bq(&_Druidfinance.CallOpts, owner)
}

// Bq is a free data retrieval call binding the contract method 0x2605a080.
//
// Solidity: function bq(address owner) view returns(bool)
func (_Druidfinance *DruidfinanceCallerSession) Bq(owner common.Address) (bool, error) {
	return _Druidfinance.Contract.Bq(&_Druidfinance.CallOpts, owner)
}

// Compare is a free data retrieval call binding the contract method 0xbeb5f658.
//
// Solidity: function compare(address first, address sec) pure returns(bool)
func (_Druidfinance *DruidfinanceCaller) Compare(opts *bind.CallOpts, first common.Address, sec common.Address) (bool, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "compare", first, sec)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Compare is a free data retrieval call binding the contract method 0xbeb5f658.
//
// Solidity: function compare(address first, address sec) pure returns(bool)
func (_Druidfinance *DruidfinanceSession) Compare(first common.Address, sec common.Address) (bool, error) {
	return _Druidfinance.Contract.Compare(&_Druidfinance.CallOpts, first, sec)
}

// Compare is a free data retrieval call binding the contract method 0xbeb5f658.
//
// Solidity: function compare(address first, address sec) pure returns(bool)
func (_Druidfinance *DruidfinanceCallerSession) Compare(first common.Address, sec common.Address) (bool, error) {
	return _Druidfinance.Contract.Compare(&_Druidfinance.CallOpts, first, sec)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Druidfinance *DruidfinanceSession) Decimals() (*big.Int, error) {
	return _Druidfinance.Contract.Decimals(&_Druidfinance.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) Decimals() (*big.Int, error) {
	return _Druidfinance.Contract.Decimals(&_Druidfinance.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Druidfinance *DruidfinanceCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Druidfinance *DruidfinanceSession) IsOwner(account common.Address) (bool, error) {
	return _Druidfinance.Contract.IsOwner(&_Druidfinance.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Druidfinance *DruidfinanceCallerSession) IsOwner(account common.Address) (bool, error) {
	return _Druidfinance.Contract.IsOwner(&_Druidfinance.CallOpts, account)
}

// LockedUntil is a free data retrieval call binding the contract method 0xce0617ec.
//
// Solidity: function lockedUntil() view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) LockedUntil(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "lockedUntil")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedUntil is a free data retrieval call binding the contract method 0xce0617ec.
//
// Solidity: function lockedUntil() view returns(uint256)
func (_Druidfinance *DruidfinanceSession) LockedUntil() (*big.Int, error) {
	return _Druidfinance.Contract.LockedUntil(&_Druidfinance.CallOpts)
}

// LockedUntil is a free data retrieval call binding the contract method 0xce0617ec.
//
// Solidity: function lockedUntil() view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) LockedUntil() (*big.Int, error) {
	return _Druidfinance.Contract.LockedUntil(&_Druidfinance.CallOpts)
}

// Marketing is a free data retrieval call binding the contract method 0x2d3e474a.
//
// Solidity: function marketing() view returns(address)
func (_Druidfinance *DruidfinanceCaller) Marketing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "marketing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketing is a free data retrieval call binding the contract method 0x2d3e474a.
//
// Solidity: function marketing() view returns(address)
func (_Druidfinance *DruidfinanceSession) Marketing() (common.Address, error) {
	return _Druidfinance.Contract.Marketing(&_Druidfinance.CallOpts)
}

// Marketing is a free data retrieval call binding the contract method 0x2d3e474a.
//
// Solidity: function marketing() view returns(address)
func (_Druidfinance *DruidfinanceCallerSession) Marketing() (common.Address, error) {
	return _Druidfinance.Contract.Marketing(&_Druidfinance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Druidfinance *DruidfinanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Druidfinance *DruidfinanceSession) Name() (string, error) {
	return _Druidfinance.Contract.Name(&_Druidfinance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Druidfinance *DruidfinanceCallerSession) Name() (string, error) {
	return _Druidfinance.Contract.Name(&_Druidfinance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Druidfinance *DruidfinanceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Druidfinance *DruidfinanceSession) Symbol() (string, error) {
	return _Druidfinance.Contract.Symbol(&_Druidfinance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Druidfinance *DruidfinanceCallerSession) Symbol() (string, error) {
	return _Druidfinance.Contract.Symbol(&_Druidfinance.CallOpts)
}

// Tax is a free data retrieval call binding the contract method 0x99c8d556.
//
// Solidity: function tax() view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) Tax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "tax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tax is a free data retrieval call binding the contract method 0x99c8d556.
//
// Solidity: function tax() view returns(uint256)
func (_Druidfinance *DruidfinanceSession) Tax() (*big.Int, error) {
	return _Druidfinance.Contract.Tax(&_Druidfinance.CallOpts)
}

// Tax is a free data retrieval call binding the contract method 0x99c8d556.
//
// Solidity: function tax() view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) Tax() (*big.Int, error) {
	return _Druidfinance.Contract.Tax(&_Druidfinance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Druidfinance *DruidfinanceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Druidfinance *DruidfinanceSession) TotalSupply() (*big.Int, error) {
	return _Druidfinance.Contract.TotalSupply(&_Druidfinance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Druidfinance *DruidfinanceCallerSession) TotalSupply() (*big.Int, error) {
	return _Druidfinance.Contract.TotalSupply(&_Druidfinance.CallOpts)
}

// Wted is a free data retrieval call binding the contract method 0x243f61b7.
//
// Solidity: function wted(address owner) view returns(bool)
func (_Druidfinance *DruidfinanceCaller) Wted(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Druidfinance.contract.Call(opts, &out, "wted", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Wted is a free data retrieval call binding the contract method 0x243f61b7.
//
// Solidity: function wted(address owner) view returns(bool)
func (_Druidfinance *DruidfinanceSession) Wted(owner common.Address) (bool, error) {
	return _Druidfinance.Contract.Wted(&_Druidfinance.CallOpts, owner)
}

// Wted is a free data retrieval call binding the contract method 0x243f61b7.
//
// Solidity: function wted(address owner) view returns(bool)
func (_Druidfinance *DruidfinanceCallerSession) Wted(owner common.Address) (bool, error) {
	return _Druidfinance.Contract.Wted(&_Druidfinance.CallOpts, owner)
}

// Transferm is a paid mutator transaction binding the contract method 0x0390de63.
//
// Solidity: function _transferm(address from, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactor) Transferm(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "_transferm", from, value)
}

// Transferm is a paid mutator transaction binding the contract method 0x0390de63.
//
// Solidity: function _transferm(address from, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceSession) Transferm(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Transferm(&_Druidfinance.TransactOpts, from, value)
}

// Transferm is a paid mutator transaction binding the contract method 0x0390de63.
//
// Solidity: function _transferm(address from, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactorSession) Transferm(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Transferm(&_Druidfinance.TransactOpts, from, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Approve(&_Druidfinance.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Approve(&_Druidfinance.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Druidfinance *DruidfinanceTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Druidfinance *DruidfinanceSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Burn(&_Druidfinance.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Druidfinance *DruidfinanceTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Burn(&_Druidfinance.TransactOpts, _value)
}

// MBMulti is a paid mutator transaction binding the contract method 0x62cfe4c9.
//
// Solidity: function m_b_multi(address[] addresses, bool status) returns()
func (_Druidfinance *DruidfinanceTransactor) MBMulti(opts *bind.TransactOpts, addresses []common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "m_b_multi", addresses, status)
}

// MBMulti is a paid mutator transaction binding the contract method 0x62cfe4c9.
//
// Solidity: function m_b_multi(address[] addresses, bool status) returns()
func (_Druidfinance *DruidfinanceSession) MBMulti(addresses []common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.Contract.MBMulti(&_Druidfinance.TransactOpts, addresses, status)
}

// MBMulti is a paid mutator transaction binding the contract method 0x62cfe4c9.
//
// Solidity: function m_b_multi(address[] addresses, bool status) returns()
func (_Druidfinance *DruidfinanceTransactorSession) MBMulti(addresses []common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.Contract.MBMulti(&_Druidfinance.TransactOpts, addresses, status)
}

// MBS is a paid mutator transaction binding the contract method 0x0cb6f8af.
//
// Solidity: function m_b_s(address addres, bool status) returns()
func (_Druidfinance *DruidfinanceTransactor) MBS(opts *bind.TransactOpts, addres common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "m_b_s", addres, status)
}

// MBS is a paid mutator transaction binding the contract method 0x0cb6f8af.
//
// Solidity: function m_b_s(address addres, bool status) returns()
func (_Druidfinance *DruidfinanceSession) MBS(addres common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.Contract.MBS(&_Druidfinance.TransactOpts, addres, status)
}

// MBS is a paid mutator transaction binding the contract method 0x0cb6f8af.
//
// Solidity: function m_b_s(address addres, bool status) returns()
func (_Druidfinance *DruidfinanceTransactorSession) MBS(addres common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.Contract.MBS(&_Druidfinance.TransactOpts, addres, status)
}

// MWS is a paid mutator transaction binding the contract method 0x61d4a118.
//
// Solidity: function m_w_s(address addres, bool status) returns()
func (_Druidfinance *DruidfinanceTransactor) MWS(opts *bind.TransactOpts, addres common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "m_w_s", addres, status)
}

// MWS is a paid mutator transaction binding the contract method 0x61d4a118.
//
// Solidity: function m_w_s(address addres, bool status) returns()
func (_Druidfinance *DruidfinanceSession) MWS(addres common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.Contract.MWS(&_Druidfinance.TransactOpts, addres, status)
}

// MWS is a paid mutator transaction binding the contract method 0x61d4a118.
//
// Solidity: function m_w_s(address addres, bool status) returns()
func (_Druidfinance *DruidfinanceTransactorSession) MWS(addres common.Address, status bool) (*types.Transaction, error) {
	return _Druidfinance.Contract.MWS(&_Druidfinance.TransactOpts, addres, status)
}

// SetTaxFeePercent is a paid mutator transaction binding the contract method 0x061c82d0.
//
// Solidity: function setTaxFeePercent(uint256 taxFee) returns()
func (_Druidfinance *DruidfinanceTransactor) SetTaxFeePercent(opts *bind.TransactOpts, taxFee *big.Int) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "setTaxFeePercent", taxFee)
}

// SetTaxFeePercent is a paid mutator transaction binding the contract method 0x061c82d0.
//
// Solidity: function setTaxFeePercent(uint256 taxFee) returns()
func (_Druidfinance *DruidfinanceSession) SetTaxFeePercent(taxFee *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.SetTaxFeePercent(&_Druidfinance.TransactOpts, taxFee)
}

// SetTaxFeePercent is a paid mutator transaction binding the contract method 0x061c82d0.
//
// Solidity: function setTaxFeePercent(uint256 taxFee) returns()
func (_Druidfinance *DruidfinanceTransactorSession) SetTaxFeePercent(taxFee *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.SetTaxFeePercent(&_Druidfinance.TransactOpts, taxFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Transfer(&_Druidfinance.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.Transfer(&_Druidfinance.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.TransferFrom(&_Druidfinance.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Druidfinance *DruidfinanceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Druidfinance.Contract.TransferFrom(&_Druidfinance.TransactOpts, from, to, value)
}

// DruidfinanceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Druidfinance contract.
type DruidfinanceApprovalIterator struct {
	Event *DruidfinanceApproval // Event containing the contract specifics and raw log

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
func (it *DruidfinanceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DruidfinanceApproval)
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
		it.Event = new(DruidfinanceApproval)
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
func (it *DruidfinanceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DruidfinanceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DruidfinanceApproval represents a Approval event raised by the Druidfinance contract.
type DruidfinanceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DruidfinanceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Druidfinance.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DruidfinanceApprovalIterator{contract: _Druidfinance.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DruidfinanceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Druidfinance.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DruidfinanceApproval)
				if err := _Druidfinance.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) ParseApproval(log types.Log) (*DruidfinanceApproval, error) {
	event := new(DruidfinanceApproval)
	if err := _Druidfinance.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DruidfinanceBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Druidfinance contract.
type DruidfinanceBurnIterator struct {
	Event *DruidfinanceBurn // Event containing the contract specifics and raw log

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
func (it *DruidfinanceBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DruidfinanceBurn)
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
		it.Event = new(DruidfinanceBurn)
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
func (it *DruidfinanceBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DruidfinanceBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DruidfinanceBurn represents a Burn event raised by the Druidfinance contract.
type DruidfinanceBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*DruidfinanceBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Druidfinance.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &DruidfinanceBurnIterator{contract: _Druidfinance.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DruidfinanceBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Druidfinance.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DruidfinanceBurn)
				if err := _Druidfinance.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) ParseBurn(log types.Log) (*DruidfinanceBurn, error) {
	event := new(DruidfinanceBurn)
	if err := _Druidfinance.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DruidfinanceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Druidfinance contract.
type DruidfinanceTransferIterator struct {
	Event *DruidfinanceTransfer // Event containing the contract specifics and raw log

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
func (it *DruidfinanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DruidfinanceTransfer)
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
		it.Event = new(DruidfinanceTransfer)
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
func (it *DruidfinanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DruidfinanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DruidfinanceTransfer represents a Transfer event raised by the Druidfinance contract.
type DruidfinanceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DruidfinanceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Druidfinance.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DruidfinanceTransferIterator{contract: _Druidfinance.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DruidfinanceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Druidfinance.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DruidfinanceTransfer)
				if err := _Druidfinance.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Druidfinance *DruidfinanceFilterer) ParseTransfer(log types.Log) (*DruidfinanceTransfer, error) {
	event := new(DruidfinanceTransfer)
	if err := _Druidfinance.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
