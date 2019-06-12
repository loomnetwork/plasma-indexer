// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethcontract

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

// LoomStoreABI is the input ABI used to generate the binding from.
const LoomStoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"NewValueSet\",\"type\":\"event\",\"signature\":\"0x7e8ee33c8615178a01c0dbd6263ef0af255dcaba4dde4f387384567abfab718f\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x8a42ebe9\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x20965255\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x17d7de7c\"}]"

// LoomStore is an auto generated Go binding around an Ethereum contract.
type LoomStore struct {
	LoomStoreCaller     // Read-only binding to the contract
	LoomStoreTransactor // Write-only binding to the contract
	LoomStoreFilterer   // Log filterer for contract events
}

// LoomStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoomStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoomStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoomStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoomStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoomStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoomStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoomStoreSession struct {
	Contract     *LoomStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoomStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoomStoreCallerSession struct {
	Contract *LoomStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LoomStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoomStoreTransactorSession struct {
	Contract     *LoomStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LoomStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoomStoreRaw struct {
	Contract *LoomStore // Generic contract binding to access the raw methods on
}

// LoomStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoomStoreCallerRaw struct {
	Contract *LoomStoreCaller // Generic read-only contract binding to access the raw methods on
}

// LoomStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoomStoreTransactorRaw struct {
	Contract *LoomStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoomStore creates a new instance of LoomStore, bound to a specific deployed contract.
func NewLoomStore(address common.Address, backend bind.ContractBackend) (*LoomStore, error) {
	contract, err := bindLoomStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoomStore{LoomStoreCaller: LoomStoreCaller{contract: contract}, LoomStoreTransactor: LoomStoreTransactor{contract: contract}, LoomStoreFilterer: LoomStoreFilterer{contract: contract}}, nil
}

// NewLoomStoreCaller creates a new read-only instance of LoomStore, bound to a specific deployed contract.
func NewLoomStoreCaller(address common.Address, caller bind.ContractCaller) (*LoomStoreCaller, error) {
	contract, err := bindLoomStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoomStoreCaller{contract: contract}, nil
}

// NewLoomStoreTransactor creates a new write-only instance of LoomStore, bound to a specific deployed contract.
func NewLoomStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*LoomStoreTransactor, error) {
	contract, err := bindLoomStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoomStoreTransactor{contract: contract}, nil
}

// NewLoomStoreFilterer creates a new log filterer instance of LoomStore, bound to a specific deployed contract.
func NewLoomStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*LoomStoreFilterer, error) {
	contract, err := bindLoomStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoomStoreFilterer{contract: contract}, nil
}

// bindLoomStore binds a generic wrapper to an already deployed contract.
func bindLoomStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoomStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoomStore *LoomStoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoomStore.Contract.LoomStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoomStore *LoomStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoomStore.Contract.LoomStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoomStore *LoomStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoomStore.Contract.LoomStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoomStore *LoomStoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoomStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoomStore *LoomStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoomStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoomStore *LoomStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoomStore.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_LoomStore *LoomStoreCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoomStore.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_LoomStore *LoomStoreSession) GetName() (string, error) {
	return _LoomStore.Contract.GetName(&_LoomStore.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_LoomStore *LoomStoreCallerSession) GetName() (string, error) {
	return _LoomStore.Contract.GetName(&_LoomStore.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() constant returns(uint256)
func (_LoomStore *LoomStoreCaller) GetValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoomStore.contract.Call(opts, out, "getValue")
	return *ret0, err
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() constant returns(uint256)
func (_LoomStore *LoomStoreSession) GetValue() (*big.Int, error) {
	return _LoomStore.Contract.GetValue(&_LoomStore.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() constant returns(uint256)
func (_LoomStore *LoomStoreCallerSession) GetValue() (*big.Int, error) {
	return _LoomStore.Contract.GetValue(&_LoomStore.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(_name string, _value uint256) returns()
func (_LoomStore *LoomStoreTransactor) Set(opts *bind.TransactOpts, _name string, _value *big.Int) (*types.Transaction, error) {
	return _LoomStore.contract.Transact(opts, "set", _name, _value)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(_name string, _value uint256) returns()
func (_LoomStore *LoomStoreSession) Set(_name string, _value *big.Int) (*types.Transaction, error) {
	return _LoomStore.Contract.Set(&_LoomStore.TransactOpts, _name, _value)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(_name string, _value uint256) returns()
func (_LoomStore *LoomStoreTransactorSession) Set(_name string, _value *big.Int) (*types.Transaction, error) {
	return _LoomStore.Contract.Set(&_LoomStore.TransactOpts, _name, _value)
}

// LoomStoreNewValueSetIterator is returned from FilterNewValueSet and is used to iterate over the raw logs and unpacked data for NewValueSet events raised by the LoomStore contract.
type LoomStoreNewValueSetIterator struct {
	Event *LoomStoreNewValueSet // Event containing the contract specifics and raw log

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
func (it *LoomStoreNewValueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoomStoreNewValueSet)
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
		it.Event = new(LoomStoreNewValueSet)
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
func (it *LoomStoreNewValueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoomStoreNewValueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoomStoreNewValueSet represents a NewValueSet event raised by the LoomStore contract.
type LoomStoreNewValueSet struct {
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewValueSet is a free log retrieval operation binding the contract event 0x7e8ee33c8615178a01c0dbd6263ef0af255dcaba4dde4f387384567abfab718f.
//
// Solidity: e NewValueSet(_name string, _value uint256)
func (_LoomStore *LoomStoreFilterer) FilterNewValueSet(opts *bind.FilterOpts) (*LoomStoreNewValueSetIterator, error) {

	logs, sub, err := _LoomStore.contract.FilterLogs(opts, "NewValueSet")
	if err != nil {
		return nil, err
	}
	return &LoomStoreNewValueSetIterator{contract: _LoomStore.contract, event: "NewValueSet", logs: logs, sub: sub}, nil
}

// WatchNewValueSet is a free log subscription operation binding the contract event 0x7e8ee33c8615178a01c0dbd6263ef0af255dcaba4dde4f387384567abfab718f.
//
// Solidity: e NewValueSet(_name string, _value uint256)
func (_LoomStore *LoomStoreFilterer) WatchNewValueSet(opts *bind.WatchOpts, sink chan<- *LoomStoreNewValueSet) (event.Subscription, error) {

	logs, sub, err := _LoomStore.contract.WatchLogs(opts, "NewValueSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoomStoreNewValueSet)
				if err := _LoomStore.contract.UnpackLog(event, "NewValueSet", log); err != nil {
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
