// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cow

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

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// MainMetaData contains all meta data concerning the CowMain contract.
var CowMainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickToSellAt\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockLimit\",\"type\":\"uint256\"}],\"name\":\"OrderPlaced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"brevisCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"int24\",\"name\":\"tickToSellAt\",\"type\":\"int24\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockLimit\",\"type\":\"uint256\"}],\"name\":\"placeOrder\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CowMainABI is the input ABI used to generate the binding from.
// Deprecated: Use CowMainMetaData.ABI instead.
var CowMainABI = CowMainMetaData.ABI

// CowMain is an auto generated Go binding around an Ethereum contract.
type CowMain struct {
	CowMainCaller     // Read-only binding to the contract
	CowMainTransactor // Write-only binding to the contract
	CowMainFilterer   // Log filterer for contract events
}

// CowMainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CowMainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CowMainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CowMainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CowMainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CowMainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CowMainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CowMainSession struct {
	Contract     *CowMain             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CowMainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CowMainCallerSession struct {
	Contract *CowMainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CowMainTransactionSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CowMainTransactionSession struct {
	Contract     *CowMainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CowMainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CowMainRaw struct {
	Contract *CowMain // Generic contract binding to access the raw methods on
}

// CowMainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CowMainCallerRaw struct {
	Contract *CowMainCaller // Generic read-only contract binding to access the raw methods on
}

// CowMainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CowMainTransactorRaw struct {
	Contract *CowMainTransactor // Generic write-only contract binding to access the raw methods on
}

// CowNewMain creates a new instance of CowMain, bound to a specific deployed contract.
func CowNewMain(address common.Address, backend bind.ContractBackend) (*CowMain, error) {
	contract, err := cowBindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CowMain{CowMainCaller: CowMainCaller{contract: contract}, CowMainTransactor: CowMainTransactor{contract: contract}, CowMainFilterer: CowMainFilterer{contract: contract}}, nil
}

// CowNewMainCaller creates a new read-only instance of CowMain, bound to a specific deployed contract.
func CowNewMainCaller(address common.Address, caller bind.ContractCaller) (*CowMainCaller, error) {
	contract, err := cowBindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CowMainCaller{contract: contract}, nil
}

// CowNewMainTransactor creates a new write-only instance of CowMain, bound to a specific deployed contract.
func CowNewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*CowMainTransactor, error) {
	contract, err := cowBindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CowMainTransactor{contract: contract}, nil
}

// CowNewMainFilterer creates a new log filterer instance of CowMain, bound to a specific deployed contract.
func CowNewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*CowMainFilterer, error) {
	contract, err := cowBindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CowMainFilterer{contract: contract}, nil
}

// cowBindMain binds a generic wrapper to an already deployed contract.
func cowBindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CowMainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *CowMainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.CowMainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *CowMainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.CowMainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *CowMainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.CowMainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *CowMainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *CowMainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *CowMainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// BrevisCallback is a paid mutator transaction binding the contract method 0x944556d6.
//
// Solidity: function brevisCallback(bytes callbackData) returns()
func (_Main *CowMainTransactor) BrevisCallback(opts *bind.TransactOpts, callbackData []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "brevisCallback", callbackData)
}

// BrevisCallback is a paid mutator transaction binding the contract method 0x944556d6.
//
// Solidity: function brevisCallback(bytes callbackData) returns()
func (_Main *CowMainSession) BrevisCallback(callbackData []byte) (*types.Transaction, error) {
	return _Main.Contract.BrevisCallback(&_Main.TransactOpts, callbackData)
}

// BrevisCallback is a paid mutator transaction binding the contract method 0x944556d6.
//
// Solidity: function brevisCallback(bytes callbackData) returns()
func (_Main *CowMainTransactionSession) BrevisCallback(callbackData []byte) (*types.Transaction, error) {
	return _Main.Contract.BrevisCallback(&_Main.TransactOpts, callbackData)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x28fc8ff8.
//
// Solidity: function placeOrder((address,address,uint24,int24,address) key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit) returns(int24)
func (_Main *CowMainTransactor) PlaceOrder(opts *bind.TransactOpts, key PoolKey, tickToSellAt *big.Int, zeroForOne bool, inputAmount *big.Int, blockLimit *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "placeOrder", key, tickToSellAt, zeroForOne, inputAmount, blockLimit)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x28fc8ff8.
//
// Solidity: function placeOrder((address,address,uint24,int24,address) key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit) returns(int24)
func (_Main *CowMainSession) PlaceOrder(key PoolKey, tickToSellAt *big.Int, zeroForOne bool, inputAmount *big.Int, blockLimit *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PlaceOrder(&_Main.TransactOpts, key, tickToSellAt, zeroForOne, inputAmount, blockLimit)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x28fc8ff8.
//
// Solidity: function placeOrder((address,address,uint24,int24,address) key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit) returns(int24)
func (_Main *CowMainTransactionSession) PlaceOrder(key PoolKey, tickToSellAt *big.Int, zeroForOne bool, inputAmount *big.Int, blockLimit *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PlaceOrder(&_Main.TransactOpts, key, tickToSellAt, zeroForOne, inputAmount, blockLimit)
}

// MainOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the CowMain contract.
type MainOrderPlacedIterator struct {
	Event *MainOrderPlaced // Event containing the contract specifics and raw log

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
func (it *MainOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOrderPlaced)
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
		it.Event = new(MainOrderPlaced)
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
func (it *MainOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOrderPlaced represents a OrderPlaced event raised by the CowMain contract.
type MainOrderPlaced struct {
	Key          PoolKey
	TickToSellAt *big.Int
	ZeroForOne   bool
	InputAmount  *big.Int
	BlockLimit   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0xa58a8688050a7b03684f7441a3d8647b75b0312db4f95826b5d3f4ea2e517220.
//
// Solidity: event OrderPlaced((address,address,uint24,int24,address) key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit)
func (_Main *CowMainFilterer) FilterOrderPlaced(opts *bind.FilterOpts) (*MainOrderPlacedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "OrderPlaced")
	if err != nil {
		return nil, err
	}
	return &MainOrderPlacedIterator{contract: _Main.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0xa58a8688050a7b03684f7441a3d8647b75b0312db4f95826b5d3f4ea2e517220.
//
// Solidity: event OrderPlaced((address,address,uint24,int24,address) key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit)
func (_Main *CowMainFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *MainOrderPlaced) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "OrderPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOrderPlaced)
				if err := _Main.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0xa58a8688050a7b03684f7441a3d8647b75b0312db4f95826b5d3f4ea2e517220.
//
// Solidity: event OrderPlaced((address,address,uint24,int24,address) key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit)
func (_Main *CowMainFilterer) ParseOrderPlaced(log types.Log) (*MainOrderPlaced, error) {
	event := new(MainOrderPlaced)
	if err := _Main.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
