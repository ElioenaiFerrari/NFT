// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// NFTMetaData contains all meta data concerning the NFT contract.
var NFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenId\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"aa6ca808": "getTokens()",
		"d85d3d27": "mint(string)",
		"07546172": "minter()",
		"fcac41f8": "send(string,address)",
		"4abf825d": "tokens(address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556107d3806100326000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063075461721461005c5780634abf825d1461008c578063aa6ca808146100ac578063d85d3d27146100c1578063fcac41f8146100d6575b600080fd5b60005461006f906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61009f61009a366004610517565b6100e9565b604051610083919061059d565b6100b46101a2565b60405161008391906105b7565b6100d46100cf3660046106bc565b61028f565b005b6100d46100e43660046106f9565b61037a565b6001602052816000526040600020818154811061010557600080fd5b9060005260206000200160009150915050805461012190610747565b80601f016020809104026020016040519081016040528092919081815260200182805461014d90610747565b801561019a5780601f1061016f5761010080835404028352916020019161019a565b820191906000526020600020905b81548152906001019060200180831161017d57829003601f168201915b505050505081565b600080546001600160a01b03168152600160209081526040808320805482518185028101850190935280835260609492939192909184015b828210156102865783829060005260206000200180546101f990610747565b80601f016020809104026020016040519081016040528092919081815260200182805461022590610747565b80156102725780601f1061024757610100808354040283529160200191610272565b820191906000526020600020905b81548152906001019060200180831161025557829003601f168201915b5050505050815260200190600101906101da565b50505050905090565b6000546001600160a01b031633146102ee5760405162461bcd60e51b815260206004820152601b60248201527f4f6e6c79206d696e7465722063616e206d696e7420746f6b656e73000000000060448201526064015b60405180910390fd5b600080546001600160a01b03168152600160208181526040832080549283018155835291829020835161032993919092019190840190610462565b50806040516103389190610781565b6040519081900381206000805491926001600160a01b03909216917f4c2f6dd3c99d950ddda88131093e3e393b8e548551839aa501d799a7e94c2a149190a350565b6000546001600160a01b031633146103d45760405162461bcd60e51b815260206004820152601f60248201527f4f6e6c79206d696e7465722063616e207472616e7366657220746f6b656e730060448201526064016102e5565b6001600160a01b0381166000908152600160208181526040832080549283018155835291829020845161040f93919092019190850190610462565b508160405161041e9190610781565b6040519081900381206000805491926001600160a01b038086169316917f55e10366a5f552746106978b694d7ef3bbddec06bd5f9b9d15ad46f475c653ef91a45050565b82805461046e90610747565b90600052602060002090601f01602090048101928261049057600085556104d6565b82601f106104a957805160ff19168380011785556104d6565b828001600101855582156104d6579182015b828111156104d65782518255916020019190600101906104bb565b506104e29291506104e6565b5090565b5b808211156104e257600081556001016104e7565b80356001600160a01b038116811461051257600080fd5b919050565b6000806040838503121561052a57600080fd5b610533836104fb565b946020939093013593505050565b60005b8381101561055c578181015183820152602001610544565b8381111561056b576000848401525b50505050565b60008151808452610589816020860160208601610541565b601f01601f19169290920160200192915050565b6020815260006105b06020830184610571565b9392505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561060c57603f198886030184526105fa858351610571565b945092850192908501906001016105de565b5092979650505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261064057600080fd5b813567ffffffffffffffff8082111561065b5761065b610619565b604051601f8301601f19908116603f0116810190828211818310171561068357610683610619565b8160405283815286602085880101111561069c57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156106ce57600080fd5b813567ffffffffffffffff8111156106e557600080fd5b6106f18482850161062f565b949350505050565b6000806040838503121561070c57600080fd5b823567ffffffffffffffff81111561072357600080fd5b61072f8582860161062f565b92505061073e602084016104fb565b90509250929050565b600181811c9082168061075b57607f821691505b60208210810361077b57634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610793818460208701610541565b919091019291505056fea2646970667358221220192c015a08c4855a08892833c64f982cfc9dfa00b5775c7d7d07fd91fa7950ac64736f6c634300080d0033",
}

// NFTABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTMetaData.ABI instead.
var NFTABI = NFTMetaData.ABI

// Deprecated: Use NFTMetaData.Sigs instead.
// NFTFuncSigs maps the 4-byte function signature to its string representation.
var NFTFuncSigs = NFTMetaData.Sigs

// NFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NFTMetaData.Bin instead.
var NFTBin = NFTMetaData.Bin

// DeployNFT deploys a new Ethereum contract, binding an instance of NFT to it.
func DeployNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NFT, error) {
	parsed, err := NFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NFT{NFTCaller: NFTCaller{contract: contract}, NFTTransactor: NFTTransactor{contract: contract}, NFTFilterer: NFTFilterer{contract: contract}}, nil
}

// NFT is an auto generated Go binding around an Ethereum contract.
type NFT struct {
	NFTCaller     // Read-only binding to the contract
	NFTTransactor // Write-only binding to the contract
	NFTFilterer   // Log filterer for contract events
}

// NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTSession struct {
	Contract     *NFT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCallerSession struct {
	Contract *NFTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTTransactorSession struct {
	Contract     *NFTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTRaw struct {
	Contract *NFT // Generic contract binding to access the raw methods on
}

// NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCallerRaw struct {
	Contract *NFTCaller // Generic read-only contract binding to access the raw methods on
}

// NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTTransactorRaw struct {
	Contract *NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFT creates a new instance of NFT, bound to a specific deployed contract.
func NewNFT(address common.Address, backend bind.ContractBackend) (*NFT, error) {
	contract, err := bindNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFT{NFTCaller: NFTCaller{contract: contract}, NFTTransactor: NFTTransactor{contract: contract}, NFTFilterer: NFTFilterer{contract: contract}}, nil
}

// NewNFTCaller creates a new read-only instance of NFT, bound to a specific deployed contract.
func NewNFTCaller(address common.Address, caller bind.ContractCaller) (*NFTCaller, error) {
	contract, err := bindNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCaller{contract: contract}, nil
}

// NewNFTTransactor creates a new write-only instance of NFT, bound to a specific deployed contract.
func NewNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTTransactor, error) {
	contract, err := bindNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTTransactor{contract: contract}, nil
}

// NewNFTFilterer creates a new log filterer instance of NFT, bound to a specific deployed contract.
func NewNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTFilterer, error) {
	contract, err := bindNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTFilterer{contract: contract}, nil
}

// bindNFT binds a generic wrapper to an already deployed contract.
func bindNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transact(opts, method, params...)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(string[])
func (_NFT *NFTCaller) GetTokens(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(string[])
func (_NFT *NFTSession) GetTokens() ([]string, error) {
	return _NFT.Contract.GetTokens(&_NFT.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(string[])
func (_NFT *NFTCallerSession) GetTokens() ([]string, error) {
	return _NFT.Contract.GetTokens(&_NFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NFT *NFTCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NFT *NFTSession) Minter() (common.Address, error) {
	return _NFT.Contract.Minter(&_NFT.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NFT *NFTCallerSession) Minter() (common.Address, error) {
	return _NFT.Contract.Minter(&_NFT.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4abf825d.
//
// Solidity: function tokens(address , uint256 ) view returns(string)
func (_NFT *NFTCaller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokens", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4abf825d.
//
// Solidity: function tokens(address , uint256 ) view returns(string)
func (_NFT *NFTSession) Tokens(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _NFT.Contract.Tokens(&_NFT.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x4abf825d.
//
// Solidity: function tokens(address , uint256 ) view returns(string)
func (_NFT *NFTCallerSession) Tokens(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _NFT.Contract.Tokens(&_NFT.CallOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0xd85d3d27.
//
// Solidity: function mint(string _tokenId) returns()
func (_NFT *NFTTransactor) Mint(opts *bind.TransactOpts, _tokenId string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "mint", _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd85d3d27.
//
// Solidity: function mint(string _tokenId) returns()
func (_NFT *NFTSession) Mint(_tokenId string) (*types.Transaction, error) {
	return _NFT.Contract.Mint(&_NFT.TransactOpts, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd85d3d27.
//
// Solidity: function mint(string _tokenId) returns()
func (_NFT *NFTTransactorSession) Mint(_tokenId string) (*types.Transaction, error) {
	return _NFT.Contract.Mint(&_NFT.TransactOpts, _tokenId)
}

// Send is a paid mutator transaction binding the contract method 0xfcac41f8.
//
// Solidity: function send(string _tokenId, address _to) returns()
func (_NFT *NFTTransactor) Send(opts *bind.TransactOpts, _tokenId string, _to common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "send", _tokenId, _to)
}

// Send is a paid mutator transaction binding the contract method 0xfcac41f8.
//
// Solidity: function send(string _tokenId, address _to) returns()
func (_NFT *NFTSession) Send(_tokenId string, _to common.Address) (*types.Transaction, error) {
	return _NFT.Contract.Send(&_NFT.TransactOpts, _tokenId, _to)
}

// Send is a paid mutator transaction binding the contract method 0xfcac41f8.
//
// Solidity: function send(string _tokenId, address _to) returns()
func (_NFT *NFTTransactorSession) Send(_tokenId string, _to common.Address) (*types.Transaction, error) {
	return _NFT.Contract.Send(&_NFT.TransactOpts, _tokenId, _to)
}

// NFTMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the NFT contract.
type NFTMintIterator struct {
	Event *NFTMint // Event containing the contract specifics and raw log

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
func (it *NFTMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMint)
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
		it.Event = new(NFTMint)
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
func (it *NFTMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMint represents a Mint event raised by the NFT contract.
type NFTMint struct {
	To      common.Address
	TokenId common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c2f6dd3c99d950ddda88131093e3e393b8e548551839aa501d799a7e94c2a14.
//
// Solidity: event Mint(address indexed to, string indexed tokenId)
func (_NFT *NFTFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address, tokenId []string) (*NFTMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Mint", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTMintIterator{contract: _NFT.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c2f6dd3c99d950ddda88131093e3e393b8e548551839aa501d799a7e94c2a14.
//
// Solidity: event Mint(address indexed to, string indexed tokenId)
func (_NFT *NFTFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *NFTMint, to []common.Address, tokenId []string) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Mint", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMint)
				if err := _NFT.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c2f6dd3c99d950ddda88131093e3e393b8e548551839aa501d799a7e94c2a14.
//
// Solidity: event Mint(address indexed to, string indexed tokenId)
func (_NFT *NFTFilterer) ParseMint(log types.Log) (*NFTMint, error) {
	event := new(NFTMint)
	if err := _NFT.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFT contract.
type NFTTransferIterator struct {
	Event *NFTTransfer // Event containing the contract specifics and raw log

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
func (it *NFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTTransfer)
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
		it.Event = new(NFTTransfer)
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
func (it *NFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTTransfer represents a Transfer event raised by the NFT contract.
type NFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x55e10366a5f552746106978b694d7ef3bbddec06bd5f9b9d15ad46f475c653ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, string indexed tokenId)
func (_NFT *NFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []string) (*NFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTTransferIterator{contract: _NFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x55e10366a5f552746106978b694d7ef3bbddec06bd5f9b9d15ad46f475c653ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, string indexed tokenId)
func (_NFT *NFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTTransfer, from []common.Address, to []common.Address, tokenId []string) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTTransfer)
				if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x55e10366a5f552746106978b694d7ef3bbddec06bd5f9b9d15ad46f475c653ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, string indexed tokenId)
func (_NFT *NFTFilterer) ParseTransfer(log types.Log) (*NFTTransfer, error) {
	event := new(NFTTransfer)
	if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
