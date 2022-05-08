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

// Token is an auto generated low-level Go binding around an user-defined struct.
type Token struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}

// NFTMetaData contains all meta data concerning the NFT contract.
var NFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"price\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"}],\"internalType\":\"structToken[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_price\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"_imageUrl\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"price\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"price\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"aa6ca808": "getTokens()",
		"c2954062": "mint(string,string,uint64,string)",
		"07546172": "minter()",
		"fcac41f8": "send(string,address)",
		"9ead7222": "tokenList(uint256)",
		"600ca0ca": "tokens(address,string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610f31806100326000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630754617214610067578063600ca0ca146100975780639ead7222146100ba578063aa6ca808146100cd578063c2954062146100e2578063fcac41f8146100f7575b600080fd5b60005461007a906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100aa6100a5366004610be8565b61010a565b60405161008e9493929190610c92565b6100aa6100c8366004610ce6565b6102f6565b6100d5610329565b60405161008e9190610cff565b6100f56100f0366004610db5565b610556565b005b6100f5610105366004610e57565b6107bb565b6001602090815260009283526040909220815180830184018051928152908401929093019190912091528054819061014190610ea5565b80601f016020809104026020016040519081016040528092919081815260200182805461016d90610ea5565b80156101ba5780601f1061018f576101008083540402835291602001916101ba565b820191906000526020600020905b81548152906001019060200180831161019d57829003601f168201915b5050505050908060010180546101cf90610ea5565b80601f01602080910402602001604051908101604052809291908181526020018280546101fb90610ea5565b80156102485780601f1061021d57610100808354040283529160200191610248565b820191906000526020600020905b81548152906001019060200180831161022b57829003601f168201915b5050506002840154600385018054949567ffffffffffffffff90921694919350915061027390610ea5565b80601f016020809104026020016040519081016040528092919081815260200182805461029f90610ea5565b80156102ec5780601f106102c1576101008083540402835291602001916102ec565b820191906000526020600020905b8154815290600101906020018083116102cf57829003601f168201915b5050505050905084565b6002818154811061030657600080fd5b906000526020600020906004020160009150905080600001805461014190610ea5565b60606002805480602002602001604051908101604052809291908181526020016000905b8282101561054d578382906000526020600020906004020160405180608001604052908160008201805461038090610ea5565b80601f01602080910402602001604051908101604052809291908181526020018280546103ac90610ea5565b80156103f95780601f106103ce576101008083540402835291602001916103f9565b820191906000526020600020905b8154815290600101906020018083116103dc57829003601f168201915b5050505050815260200160018201805461041290610ea5565b80601f016020809104026020016040519081016040528092919081815260200182805461043e90610ea5565b801561048b5780601f106104605761010080835404028352916020019161048b565b820191906000526020600020905b81548152906001019060200180831161046e57829003601f168201915b5050509183525050600282015467ffffffffffffffff1660208201526003820180546040909201916104bc90610ea5565b80601f01602080910402602001604051908101604052809291908181526020018280546104e890610ea5565b80156105355780601f1061050a57610100808354040283529160200191610535565b820191906000526020600020905b81548152906001019060200180831161051857829003601f168201915b5050505050815250508152602001906001019061034d565b50505050905090565b6000546001600160a01b031633146105b55760405162461bcd60e51b815260206004820152601b60248201527f4f6e6c79206d696e7465722063616e206d696e7420746f6b656e73000000000060448201526064015b60405180910390fd5b60408051608081018252858152602080820186905267ffffffffffffffff85168284015260608201849052600080546001600160a01b0316815260019091528290209151909190610607908790610edf565b908152602001604051809103902060008201518160000190805190602001906106319291906109d8565b50602082810151805161064a92600185019201906109d8565b50604082015160028201805467ffffffffffffffff191667ffffffffffffffff9092169190911790556060820151805161068e9160038401916020909101906109d8565b5050600080546001600160a01b0316815260016020526040908190209051600292506106bb908790610edf565b908152604051602091819003820190208254600181018455600093845291909220825460049092020190819083906106f290610ea5565b6106fd929190610a5c565b50600182018160010190805461071290610ea5565b61071d929190610a5c565b50600282810154908201805467ffffffffffffffff191667ffffffffffffffff909216919091179055600380830180549183019161075a90610ea5565b610765929190610a5c565b505050836040516107769190610edf565b6040519081900381206000805491926001600160a01b03909216917f4c2f6dd3c99d950ddda88131093e3e393b8e548551839aa501d799a7e94c2a149190a350505050565b6000546001600160a01b031633146108155760405162461bcd60e51b815260206004820152601f60248201527f4f6e6c79206d696e7465722063616e207472616e7366657220746f6b656e730060448201526064016105ac565b600080546001600160a01b031681526001602052604090819020905161083c908490610edf565b908152602001604051809103902060016000836001600160a01b03166001600160a01b031681526020019081526020016000208360405161087d9190610edf565b90815260405190819003602001902081548190839061089b90610ea5565b6108a6929190610a5c565b5060018201816001019080546108bb90610ea5565b6108c6929190610a5c565b50600282810154908201805467ffffffffffffffff191667ffffffffffffffff909216919091179055600380830180549183019161090390610ea5565b61090e929190610a5c565b5050600080546001600160a01b031681526001602052604090819020905190915061093a908490610edf565b90815260405190819003602001902060006109558282610ad7565b610963600183016000610ad7565b60028201805467ffffffffffffffff19169055610984600383016000610ad7565b5050816040516109949190610edf565b6040519081900381206000805491926001600160a01b038086169316917f55e10366a5f552746106978b694d7ef3bbddec06bd5f9b9d15ad46f475c653ef91a45050565b8280546109e490610ea5565b90600052602060002090601f016020900481019282610a065760008555610a4c565b82601f10610a1f57805160ff1916838001178555610a4c565b82800160010185558215610a4c579182015b82811115610a4c578251825591602001919060010190610a31565b50610a58929150610b14565b5090565b828054610a6890610ea5565b90600052602060002090601f016020900481019282610a8a5760008555610a4c565b82601f10610a9b5780548555610a4c565b82800160010185558215610a4c57600052602060002091601f016020900482015b82811115610a4c578254825591600101919060010190610abc565b508054610ae390610ea5565b6000825580601f10610af3575050565b601f016020900490600052602060002090810190610b119190610b14565b50565b5b80821115610a585760008155600101610b15565b80356001600160a01b0381168114610b4057600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610b6c57600080fd5b813567ffffffffffffffff80821115610b8757610b87610b45565b604051601f8301601f19908116603f01168101908282118183101715610baf57610baf610b45565b81604052838152866020858801011115610bc857600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610bfb57600080fd5b610c0483610b29565b9150602083013567ffffffffffffffff811115610c2057600080fd5b610c2c85828601610b5b565b9150509250929050565b60005b83811015610c51578181015183820152602001610c39565b83811115610c60576000848401525b50505050565b60008151808452610c7e816020860160208601610c36565b601f01601f19169290920160200192915050565b608081526000610ca56080830187610c66565b8281036020840152610cb78187610c66565b905067ffffffffffffffff851660408401528281036060840152610cdb8185610c66565b979650505050505050565b600060208284031215610cf857600080fd5b5035919050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015610da757603f19898403018552815160808151818652610d4c82870182610c66565b915050888201518582038a870152610d648282610c66565b91505067ffffffffffffffff88830151168886015260608083015192508582038187015250610d938183610c66565b968901969450505090860190600101610d26565b509098975050505050505050565b60008060008060808587031215610dcb57600080fd5b843567ffffffffffffffff80821115610de357600080fd5b610def88838901610b5b565b95506020870135915080821115610e0557600080fd5b610e1188838901610b5b565b9450604087013591508082168214610e2857600080fd5b90925060608601359080821115610e3e57600080fd5b50610e4b87828801610b5b565b91505092959194509250565b60008060408385031215610e6a57600080fd5b823567ffffffffffffffff811115610e8157600080fd5b610e8d85828601610b5b565b925050610e9c60208401610b29565b90509250929050565b600181811c90821680610eb957607f821691505b602082108103610ed957634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610ef1818460208701610c36565b919091019291505056fea26469706673582212204c847d5545f162ebd2afc63cf5f00363a160763fcb27b3785dda0700f63d372a64736f6c634300080d0033",
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
// Solidity: function getTokens() view returns((string,string,uint64,string)[])
func (_NFT *NFTCaller) GetTokens(opts *bind.CallOpts) ([]Token, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]Token), err
	}

	out0 := *abi.ConvertType(out[0], new([]Token)).(*[]Token)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns((string,string,uint64,string)[])
func (_NFT *NFTSession) GetTokens() ([]Token, error) {
	return _NFT.Contract.GetTokens(&_NFT.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns((string,string,uint64,string)[])
func (_NFT *NFTCallerSession) GetTokens() ([]Token, error) {
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

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(string id, string name, uint64 price, string imageUrl)
func (_NFT *NFTCaller) TokenList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenList", arg0)

	outstruct := new(struct {
		Id       string
		Name     string
		Price    uint64
		ImageUrl string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ImageUrl = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(string id, string name, uint64 price, string imageUrl)
func (_NFT *NFTSession) TokenList(arg0 *big.Int) (struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}, error) {
	return _NFT.Contract.TokenList(&_NFT.CallOpts, arg0)
}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(string id, string name, uint64 price, string imageUrl)
func (_NFT *NFTCallerSession) TokenList(arg0 *big.Int) (struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}, error) {
	return _NFT.Contract.TokenList(&_NFT.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x600ca0ca.
//
// Solidity: function tokens(address , string ) view returns(string id, string name, uint64 price, string imageUrl)
func (_NFT *NFTCaller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 string) (struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokens", arg0, arg1)

	outstruct := new(struct {
		Id       string
		Name     string
		Price    uint64
		ImageUrl string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ImageUrl = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0x600ca0ca.
//
// Solidity: function tokens(address , string ) view returns(string id, string name, uint64 price, string imageUrl)
func (_NFT *NFTSession) Tokens(arg0 common.Address, arg1 string) (struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}, error) {
	return _NFT.Contract.Tokens(&_NFT.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x600ca0ca.
//
// Solidity: function tokens(address , string ) view returns(string id, string name, uint64 price, string imageUrl)
func (_NFT *NFTCallerSession) Tokens(arg0 common.Address, arg1 string) (struct {
	Id       string
	Name     string
	Price    uint64
	ImageUrl string
}, error) {
	return _NFT.Contract.Tokens(&_NFT.CallOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0xc2954062.
//
// Solidity: function mint(string _id, string _name, uint64 _price, string _imageUrl) returns()
func (_NFT *NFTTransactor) Mint(opts *bind.TransactOpts, _id string, _name string, _price uint64, _imageUrl string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "mint", _id, _name, _price, _imageUrl)
}

// Mint is a paid mutator transaction binding the contract method 0xc2954062.
//
// Solidity: function mint(string _id, string _name, uint64 _price, string _imageUrl) returns()
func (_NFT *NFTSession) Mint(_id string, _name string, _price uint64, _imageUrl string) (*types.Transaction, error) {
	return _NFT.Contract.Mint(&_NFT.TransactOpts, _id, _name, _price, _imageUrl)
}

// Mint is a paid mutator transaction binding the contract method 0xc2954062.
//
// Solidity: function mint(string _id, string _name, uint64 _price, string _imageUrl) returns()
func (_NFT *NFTTransactorSession) Mint(_id string, _name string, _price uint64, _imageUrl string) (*types.Transaction, error) {
	return _NFT.Contract.Mint(&_NFT.TransactOpts, _id, _name, _price, _imageUrl)
}

// Send is a paid mutator transaction binding the contract method 0xfcac41f8.
//
// Solidity: function send(string _id, address _to) returns()
func (_NFT *NFTTransactor) Send(opts *bind.TransactOpts, _id string, _to common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "send", _id, _to)
}

// Send is a paid mutator transaction binding the contract method 0xfcac41f8.
//
// Solidity: function send(string _id, address _to) returns()
func (_NFT *NFTSession) Send(_id string, _to common.Address) (*types.Transaction, error) {
	return _NFT.Contract.Send(&_NFT.TransactOpts, _id, _to)
}

// Send is a paid mutator transaction binding the contract method 0xfcac41f8.
//
// Solidity: function send(string _id, address _to) returns()
func (_NFT *NFTTransactorSession) Send(_id string, _to common.Address) (*types.Transaction, error) {
	return _NFT.Contract.Send(&_NFT.TransactOpts, _id, _to)
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
