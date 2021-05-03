/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Monday May 3rd 2021
**	@Filename:				cairo.go
******************************************************************************/

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CairoProverABI is the input ABI used to generate the binding from.
const CairoProverABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registriesProgramHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_identitiesProgramHash\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cairoVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Prove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UpdateRegistry\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CAIRO_VERIFIER\",\"outputs\":[{\"internalType\":\"contractIFactRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identitiesProgramHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"registryKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registryHash\",\"type\":\"uint256\"}],\"name\":\"proveIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registriesHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registriesProgramHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"registryKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldRegistryHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRegistryHash\",\"type\":\"uint256\"}],\"name\":\"updateRegistry\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CairoProverFuncSigs maps the 4-byte function signature to its string representation.
var CairoProverFuncSigs = map[string]string{
	"502864f0": "CAIRO_VERIFIER()",
	"f9fd4769": "identitiesProgramHash()",
	"8da5cb5b": "owner()",
	"bbcab69b": "proveIdentity(uint256,uint256,uint256)",
	"89cdce0b": "registriesHash(uint256)",
	"90d37f4f": "registriesProgramHash()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"4f7a55cb": "updateRegistry(uint256,uint256,uint256)",
}

// CairoProverBin is the compiled bytecode used for deploying new contracts.
var CairoProverBin = "0x608060405234801561001057600080fd5b5060405161083738038061083783398101604081905261002f916100b7565b60006100396100b3565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600391909155600291909155600480546001600160a01b0319166001600160a01b039092169190911790556100fb565b3390565b6000806000606084860312156100cb578283fd5b83516020850151604086015191945092506001600160a01b03811681146100f0578182fd5b809150509250925092565b61072d8061010a6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b146100f357806390d37f4f146100fb578063bbcab69b14610103578063f2fde38b14610123578063f9fd47691461013657610093565b80634f7a55cb14610098578063502864f0146100c1578063715018a6146100d657806389cdce0b146100e0575b600080fd5b6100ab6100a63660046105de565b61013e565b6040516100b89190610636565b60405180910390f35b6100c96102f0565b6040516100b89190610617565b6100de6102ff565b005b6100ab6100ee3660046105c6565b610388565b6100c961039a565b6100ab6103a9565b6101166101113660046105de565b6103af565b6040516100b8919061062b565b6100de610131366004610578565b6104ae565b6100ab61056e565b6000610148610574565b6001600160a01b031661015961039a565b6001600160a01b0316146101885760405162461bcd60e51b815260040161017f906106ac565b60405180910390fd5b6000838360405160200161019d929190610609565b6040516020818303038152906040528051906020012090506000600254826040516020016101cc929190610609565b60408051601f1981840301815290829052805160209091012060048054636a93856760e01b84529193506001600160a01b0390911691636a9385679161021491859101610636565b60206040518083038186803b15801561022c57600080fd5b505afa158015610240573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026491906105a6565b6102805760405162461bcd60e51b815260040161017f90610685565b600086815260016020526040902054851461029a57600080fd5b60008681526001602052604090819020859055517f9ca6faf35211f19dffd11e766fc77386c2d5e406ebe3c28620dc60288d70a8b7906102df908890889088906106e1565b60405180910390a195945050505050565b6004546001600160a01b031681565b610307610574565b6001600160a01b031661031861039a565b6001600160a01b03161461033e5760405162461bcd60e51b815260040161017f906106ac565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60016020526000908152604090205481565b6000546001600160a01b031690565b60025481565b60008083836040516020016103c5929190610609565b6040516020818303038152906040528051906020012090506000600354826040516020016103f4929190610609565b60408051601f19818403018152918152815160209283012060008981526001909352912054909150841480156104a4575060048054604051636a93856760e01b81526001600160a01b0390911691636a9385679161045491859101610636565b60206040518083038186803b15801561046c57600080fd5b505afa158015610480573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a491906105a6565b9695505050505050565b6104b6610574565b6001600160a01b03166104c761039a565b6001600160a01b0316146104ed5760405162461bcd60e51b815260040161017f906106ac565b6001600160a01b0381166105135760405162461bcd60e51b815260040161017f9061063f565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60035481565b3390565b600060208284031215610589578081fd5b81356001600160a01b038116811461059f578182fd5b9392505050565b6000602082840312156105b7578081fd5b8151801515811461059f578182fd5b6000602082840312156105d7578081fd5b5035919050565b6000806000606084860312156105f2578182fd5b505081359360208301359350604090920135919050565b918252602082015260400190565b6001600160a01b0391909116815260200190565b901515815260200190565b90815260200190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252600d908201526c24a72b20a624a22fa82927a7a360991b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b928352602083019190915260408201526060019056fea2646970667358221220473ca92a64968f2f28bd45c39f42dd8c584a173183cbd8afdd3bd3f8141c59ce64736f6c63430008000033"

// DeployCairoProver deploys a new Ethereum contract, binding an instance of CairoProver to it.
func DeployCairoProver(auth *bind.TransactOpts, backend bind.ContractBackend, _registriesProgramHash *big.Int, _identitiesProgramHash *big.Int, cairoVerifier common.Address) (common.Address, *types.Transaction, *CairoProver, error) {
	parsed, err := abi.JSON(strings.NewReader(CairoProverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CairoProverBin), backend, _registriesProgramHash, _identitiesProgramHash, cairoVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CairoProver{CairoProverCaller: CairoProverCaller{contract: contract}, CairoProverTransactor: CairoProverTransactor{contract: contract}, CairoProverFilterer: CairoProverFilterer{contract: contract}}, nil
}

// CairoProver is an auto generated Go binding around an Ethereum contract.
type CairoProver struct {
	CairoProverCaller     // Read-only binding to the contract
	CairoProverTransactor // Write-only binding to the contract
	CairoProverFilterer   // Log filterer for contract events
}

// CairoProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type CairoProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CairoProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CairoProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CairoProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CairoProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CairoProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CairoProverSession struct {
	Contract     *CairoProver      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CairoProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CairoProverCallerSession struct {
	Contract *CairoProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CairoProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CairoProverTransactorSession struct {
	Contract     *CairoProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CairoProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type CairoProverRaw struct {
	Contract *CairoProver // Generic contract binding to access the raw methods on
}

// CairoProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CairoProverCallerRaw struct {
	Contract *CairoProverCaller // Generic read-only contract binding to access the raw methods on
}

// CairoProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CairoProverTransactorRaw struct {
	Contract *CairoProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCairoProver creates a new instance of CairoProver, bound to a specific deployed contract.
func NewCairoProver(address common.Address, backend bind.ContractBackend) (*CairoProver, error) {
	contract, err := bindCairoProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CairoProver{CairoProverCaller: CairoProverCaller{contract: contract}, CairoProverTransactor: CairoProverTransactor{contract: contract}, CairoProverFilterer: CairoProverFilterer{contract: contract}}, nil
}

// NewCairoProverCaller creates a new read-only instance of CairoProver, bound to a specific deployed contract.
func NewCairoProverCaller(address common.Address, caller bind.ContractCaller) (*CairoProverCaller, error) {
	contract, err := bindCairoProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CairoProverCaller{contract: contract}, nil
}

// NewCairoProverTransactor creates a new write-only instance of CairoProver, bound to a specific deployed contract.
func NewCairoProverTransactor(address common.Address, transactor bind.ContractTransactor) (*CairoProverTransactor, error) {
	contract, err := bindCairoProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CairoProverTransactor{contract: contract}, nil
}

// NewCairoProverFilterer creates a new log filterer instance of CairoProver, bound to a specific deployed contract.
func NewCairoProverFilterer(address common.Address, filterer bind.ContractFilterer) (*CairoProverFilterer, error) {
	contract, err := bindCairoProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CairoProverFilterer{contract: contract}, nil
}

// bindCairoProver binds a generic wrapper to an already deployed contract.
func bindCairoProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CairoProverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CairoProver *CairoProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CairoProver.Contract.CairoProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CairoProver *CairoProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CairoProver.Contract.CairoProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CairoProver *CairoProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CairoProver.Contract.CairoProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CairoProver *CairoProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CairoProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CairoProver *CairoProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CairoProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CairoProver *CairoProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CairoProver.Contract.contract.Transact(opts, method, params...)
}

// CAIROVERIFIER is a free data retrieval call binding the contract method 0x502864f0.
//
// Solidity: function CAIRO_VERIFIER() view returns(address)
func (_CairoProver *CairoProverCaller) CAIROVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CairoProver.contract.Call(opts, &out, "CAIRO_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CAIROVERIFIER is a free data retrieval call binding the contract method 0x502864f0.
//
// Solidity: function CAIRO_VERIFIER() view returns(address)
func (_CairoProver *CairoProverSession) CAIROVERIFIER() (common.Address, error) {
	return _CairoProver.Contract.CAIROVERIFIER(&_CairoProver.CallOpts)
}

// CAIROVERIFIER is a free data retrieval call binding the contract method 0x502864f0.
//
// Solidity: function CAIRO_VERIFIER() view returns(address)
func (_CairoProver *CairoProverCallerSession) CAIROVERIFIER() (common.Address, error) {
	return _CairoProver.Contract.CAIROVERIFIER(&_CairoProver.CallOpts)
}

// IdentitiesProgramHash is a free data retrieval call binding the contract method 0xf9fd4769.
//
// Solidity: function identitiesProgramHash() view returns(uint256)
func (_CairoProver *CairoProverCaller) IdentitiesProgramHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CairoProver.contract.Call(opts, &out, "identitiesProgramHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdentitiesProgramHash is a free data retrieval call binding the contract method 0xf9fd4769.
//
// Solidity: function identitiesProgramHash() view returns(uint256)
func (_CairoProver *CairoProverSession) IdentitiesProgramHash() (*big.Int, error) {
	return _CairoProver.Contract.IdentitiesProgramHash(&_CairoProver.CallOpts)
}

// IdentitiesProgramHash is a free data retrieval call binding the contract method 0xf9fd4769.
//
// Solidity: function identitiesProgramHash() view returns(uint256)
func (_CairoProver *CairoProverCallerSession) IdentitiesProgramHash() (*big.Int, error) {
	return _CairoProver.Contract.IdentitiesProgramHash(&_CairoProver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CairoProver *CairoProverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CairoProver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CairoProver *CairoProverSession) Owner() (common.Address, error) {
	return _CairoProver.Contract.Owner(&_CairoProver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CairoProver *CairoProverCallerSession) Owner() (common.Address, error) {
	return _CairoProver.Contract.Owner(&_CairoProver.CallOpts)
}

// ProveIdentity is a free data retrieval call binding the contract method 0xbbcab69b.
//
// Solidity: function proveIdentity(uint256 registryKey, uint256 hash, uint256 registryHash) view returns(bool)
func (_CairoProver *CairoProverCaller) ProveIdentity(opts *bind.CallOpts, registryKey *big.Int, hash *big.Int, registryHash *big.Int) (bool, error) {
	var out []interface{}
	err := _CairoProver.contract.Call(opts, &out, "proveIdentity", registryKey, hash, registryHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveIdentity is a free data retrieval call binding the contract method 0xbbcab69b.
//
// Solidity: function proveIdentity(uint256 registryKey, uint256 hash, uint256 registryHash) view returns(bool)
func (_CairoProver *CairoProverSession) ProveIdentity(registryKey *big.Int, hash *big.Int, registryHash *big.Int) (bool, error) {
	return _CairoProver.Contract.ProveIdentity(&_CairoProver.CallOpts, registryKey, hash, registryHash)
}

// ProveIdentity is a free data retrieval call binding the contract method 0xbbcab69b.
//
// Solidity: function proveIdentity(uint256 registryKey, uint256 hash, uint256 registryHash) view returns(bool)
func (_CairoProver *CairoProverCallerSession) ProveIdentity(registryKey *big.Int, hash *big.Int, registryHash *big.Int) (bool, error) {
	return _CairoProver.Contract.ProveIdentity(&_CairoProver.CallOpts, registryKey, hash, registryHash)
}

// RegistriesHash is a free data retrieval call binding the contract method 0x89cdce0b.
//
// Solidity: function registriesHash(uint256 ) view returns(uint256)
func (_CairoProver *CairoProverCaller) RegistriesHash(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CairoProver.contract.Call(opts, &out, "registriesHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistriesHash is a free data retrieval call binding the contract method 0x89cdce0b.
//
// Solidity: function registriesHash(uint256 ) view returns(uint256)
func (_CairoProver *CairoProverSession) RegistriesHash(arg0 *big.Int) (*big.Int, error) {
	return _CairoProver.Contract.RegistriesHash(&_CairoProver.CallOpts, arg0)
}

// RegistriesHash is a free data retrieval call binding the contract method 0x89cdce0b.
//
// Solidity: function registriesHash(uint256 ) view returns(uint256)
func (_CairoProver *CairoProverCallerSession) RegistriesHash(arg0 *big.Int) (*big.Int, error) {
	return _CairoProver.Contract.RegistriesHash(&_CairoProver.CallOpts, arg0)
}

// RegistriesProgramHash is a free data retrieval call binding the contract method 0x90d37f4f.
//
// Solidity: function registriesProgramHash() view returns(uint256)
func (_CairoProver *CairoProverCaller) RegistriesProgramHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CairoProver.contract.Call(opts, &out, "registriesProgramHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistriesProgramHash is a free data retrieval call binding the contract method 0x90d37f4f.
//
// Solidity: function registriesProgramHash() view returns(uint256)
func (_CairoProver *CairoProverSession) RegistriesProgramHash() (*big.Int, error) {
	return _CairoProver.Contract.RegistriesProgramHash(&_CairoProver.CallOpts)
}

// RegistriesProgramHash is a free data retrieval call binding the contract method 0x90d37f4f.
//
// Solidity: function registriesProgramHash() view returns(uint256)
func (_CairoProver *CairoProverCallerSession) RegistriesProgramHash() (*big.Int, error) {
	return _CairoProver.Contract.RegistriesProgramHash(&_CairoProver.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CairoProver *CairoProverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CairoProver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CairoProver *CairoProverSession) RenounceOwnership() (*types.Transaction, error) {
	return _CairoProver.Contract.RenounceOwnership(&_CairoProver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CairoProver *CairoProverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CairoProver.Contract.RenounceOwnership(&_CairoProver.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CairoProver *CairoProverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CairoProver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CairoProver *CairoProverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CairoProver.Contract.TransferOwnership(&_CairoProver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CairoProver *CairoProverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CairoProver.Contract.TransferOwnership(&_CairoProver.TransactOpts, newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x4f7a55cb.
//
// Solidity: function updateRegistry(uint256 registryKey, uint256 oldRegistryHash, uint256 newRegistryHash) returns(bytes32)
func (_CairoProver *CairoProverTransactor) UpdateRegistry(opts *bind.TransactOpts, registryKey *big.Int, oldRegistryHash *big.Int, newRegistryHash *big.Int) (*types.Transaction, error) {
	return _CairoProver.contract.Transact(opts, "updateRegistry", registryKey, oldRegistryHash, newRegistryHash)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x4f7a55cb.
//
// Solidity: function updateRegistry(uint256 registryKey, uint256 oldRegistryHash, uint256 newRegistryHash) returns(bytes32)
func (_CairoProver *CairoProverSession) UpdateRegistry(registryKey *big.Int, oldRegistryHash *big.Int, newRegistryHash *big.Int) (*types.Transaction, error) {
	return _CairoProver.Contract.UpdateRegistry(&_CairoProver.TransactOpts, registryKey, oldRegistryHash, newRegistryHash)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x4f7a55cb.
//
// Solidity: function updateRegistry(uint256 registryKey, uint256 oldRegistryHash, uint256 newRegistryHash) returns(bytes32)
func (_CairoProver *CairoProverTransactorSession) UpdateRegistry(registryKey *big.Int, oldRegistryHash *big.Int, newRegistryHash *big.Int) (*types.Transaction, error) {
	return _CairoProver.Contract.UpdateRegistry(&_CairoProver.TransactOpts, registryKey, oldRegistryHash, newRegistryHash)
}

// CairoProverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CairoProver contract.
type CairoProverOwnershipTransferredIterator struct {
	Event *CairoProverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CairoProverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CairoProverOwnershipTransferred)
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
		it.Event = new(CairoProverOwnershipTransferred)
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
func (it *CairoProverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CairoProverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CairoProverOwnershipTransferred represents a OwnershipTransferred event raised by the CairoProver contract.
type CairoProverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CairoProver *CairoProverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CairoProverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CairoProver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CairoProverOwnershipTransferredIterator{contract: _CairoProver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CairoProver *CairoProverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CairoProverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CairoProver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CairoProverOwnershipTransferred)
				if err := _CairoProver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CairoProver *CairoProverFilterer) ParseOwnershipTransferred(log types.Log) (*CairoProverOwnershipTransferred, error) {
	event := new(CairoProverOwnershipTransferred)
	if err := _CairoProver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CairoProverProveIterator is returned from FilterProve and is used to iterate over the raw logs and unpacked data for Prove events raised by the CairoProver contract.
type CairoProverProveIterator struct {
	Event *CairoProverProve // Event containing the contract specifics and raw log

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
func (it *CairoProverProveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CairoProverProve)
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
		it.Event = new(CairoProverProve)
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
func (it *CairoProverProveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CairoProverProveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CairoProverProve represents a Prove event raised by the CairoProver contract.
type CairoProverProve struct {
	Arg0 *big.Int
	Arg1 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProve is a free log retrieval operation binding the contract event 0x0211c34451638aa441e5171e6efe280ce65af9f54bfd3e5c093ec2ab0287a1c2.
//
// Solidity: event Prove(uint256 arg0, bytes32 arg1)
func (_CairoProver *CairoProverFilterer) FilterProve(opts *bind.FilterOpts) (*CairoProverProveIterator, error) {

	logs, sub, err := _CairoProver.contract.FilterLogs(opts, "Prove")
	if err != nil {
		return nil, err
	}
	return &CairoProverProveIterator{contract: _CairoProver.contract, event: "Prove", logs: logs, sub: sub}, nil
}

// WatchProve is a free log subscription operation binding the contract event 0x0211c34451638aa441e5171e6efe280ce65af9f54bfd3e5c093ec2ab0287a1c2.
//
// Solidity: event Prove(uint256 arg0, bytes32 arg1)
func (_CairoProver *CairoProverFilterer) WatchProve(opts *bind.WatchOpts, sink chan<- *CairoProverProve) (event.Subscription, error) {

	logs, sub, err := _CairoProver.contract.WatchLogs(opts, "Prove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CairoProverProve)
				if err := _CairoProver.contract.UnpackLog(event, "Prove", log); err != nil {
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

// ParseProve is a log parse operation binding the contract event 0x0211c34451638aa441e5171e6efe280ce65af9f54bfd3e5c093ec2ab0287a1c2.
//
// Solidity: event Prove(uint256 arg0, bytes32 arg1)
func (_CairoProver *CairoProverFilterer) ParseProve(log types.Log) (*CairoProverProve, error) {
	event := new(CairoProverProve)
	if err := _CairoProver.contract.UnpackLog(event, "Prove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CairoProverUpdateRegistryIterator is returned from FilterUpdateRegistry and is used to iterate over the raw logs and unpacked data for UpdateRegistry events raised by the CairoProver contract.
type CairoProverUpdateRegistryIterator struct {
	Event *CairoProverUpdateRegistry // Event containing the contract specifics and raw log

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
func (it *CairoProverUpdateRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CairoProverUpdateRegistry)
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
		it.Event = new(CairoProverUpdateRegistry)
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
func (it *CairoProverUpdateRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CairoProverUpdateRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CairoProverUpdateRegistry represents a UpdateRegistry event raised by the CairoProver contract.
type CairoProverUpdateRegistry struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateRegistry is a free log retrieval operation binding the contract event 0x9ca6faf35211f19dffd11e766fc77386c2d5e406ebe3c28620dc60288d70a8b7.
//
// Solidity: event UpdateRegistry(uint256 arg0, uint256 arg1, uint256 arg2)
func (_CairoProver *CairoProverFilterer) FilterUpdateRegistry(opts *bind.FilterOpts) (*CairoProverUpdateRegistryIterator, error) {

	logs, sub, err := _CairoProver.contract.FilterLogs(opts, "UpdateRegistry")
	if err != nil {
		return nil, err
	}
	return &CairoProverUpdateRegistryIterator{contract: _CairoProver.contract, event: "UpdateRegistry", logs: logs, sub: sub}, nil
}

// WatchUpdateRegistry is a free log subscription operation binding the contract event 0x9ca6faf35211f19dffd11e766fc77386c2d5e406ebe3c28620dc60288d70a8b7.
//
// Solidity: event UpdateRegistry(uint256 arg0, uint256 arg1, uint256 arg2)
func (_CairoProver *CairoProverFilterer) WatchUpdateRegistry(opts *bind.WatchOpts, sink chan<- *CairoProverUpdateRegistry) (event.Subscription, error) {

	logs, sub, err := _CairoProver.contract.WatchLogs(opts, "UpdateRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CairoProverUpdateRegistry)
				if err := _CairoProver.contract.UnpackLog(event, "UpdateRegistry", log); err != nil {
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

// ParseUpdateRegistry is a log parse operation binding the contract event 0x9ca6faf35211f19dffd11e766fc77386c2d5e406ebe3c28620dc60288d70a8b7.
//
// Solidity: event UpdateRegistry(uint256 arg0, uint256 arg1, uint256 arg2)
func (_CairoProver *CairoProverFilterer) ParseUpdateRegistry(log types.Log) (*CairoProverUpdateRegistry, error) {
	event := new(CairoProverUpdateRegistry)
	if err := _CairoProver.contract.UnpackLog(event, "UpdateRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// IFactRegistryABI is the input ABI used to generate the binding from.
const IFactRegistryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fact\",\"type\":\"bytes32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFactRegistryFuncSigs maps the 4-byte function signature to its string representation.
var IFactRegistryFuncSigs = map[string]string{
	"6a938567": "isValid(bytes32)",
}

// IFactRegistry is an auto generated Go binding around an Ethereum contract.
type IFactRegistry struct {
	IFactRegistryCaller     // Read-only binding to the contract
	IFactRegistryTransactor // Write-only binding to the contract
	IFactRegistryFilterer   // Log filterer for contract events
}

// IFactRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFactRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFactRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFactRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFactRegistrySession struct {
	Contract     *IFactRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFactRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFactRegistryCallerSession struct {
	Contract *IFactRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IFactRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFactRegistryTransactorSession struct {
	Contract     *IFactRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IFactRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFactRegistryRaw struct {
	Contract *IFactRegistry // Generic contract binding to access the raw methods on
}

// IFactRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFactRegistryCallerRaw struct {
	Contract *IFactRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IFactRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFactRegistryTransactorRaw struct {
	Contract *IFactRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFactRegistry creates a new instance of IFactRegistry, bound to a specific deployed contract.
func NewIFactRegistry(address common.Address, backend bind.ContractBackend) (*IFactRegistry, error) {
	contract, err := bindIFactRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFactRegistry{IFactRegistryCaller: IFactRegistryCaller{contract: contract}, IFactRegistryTransactor: IFactRegistryTransactor{contract: contract}, IFactRegistryFilterer: IFactRegistryFilterer{contract: contract}}, nil
}

// NewIFactRegistryCaller creates a new read-only instance of IFactRegistry, bound to a specific deployed contract.
func NewIFactRegistryCaller(address common.Address, caller bind.ContractCaller) (*IFactRegistryCaller, error) {
	contract, err := bindIFactRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFactRegistryCaller{contract: contract}, nil
}

// NewIFactRegistryTransactor creates a new write-only instance of IFactRegistry, bound to a specific deployed contract.
func NewIFactRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFactRegistryTransactor, error) {
	contract, err := bindIFactRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFactRegistryTransactor{contract: contract}, nil
}

// NewIFactRegistryFilterer creates a new log filterer instance of IFactRegistry, bound to a specific deployed contract.
func NewIFactRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFactRegistryFilterer, error) {
	contract, err := bindIFactRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFactRegistryFilterer{contract: contract}, nil
}

// bindIFactRegistry binds a generic wrapper to an already deployed contract.
func bindIFactRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFactRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactRegistry *IFactRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFactRegistry.Contract.IFactRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactRegistry *IFactRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactRegistry.Contract.IFactRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFactRegistry *IFactRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFactRegistry.Contract.IFactRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactRegistry *IFactRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFactRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactRegistry *IFactRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFactRegistry *IFactRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFactRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsValid is a free data retrieval call binding the contract method 0x6a938567.
//
// Solidity: function isValid(bytes32 fact) view returns(bool)
func (_IFactRegistry *IFactRegistryCaller) IsValid(opts *bind.CallOpts, fact [32]byte) (bool, error) {
	var out []interface{}
	err := _IFactRegistry.contract.Call(opts, &out, "isValid", fact)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0x6a938567.
//
// Solidity: function isValid(bytes32 fact) view returns(bool)
func (_IFactRegistry *IFactRegistrySession) IsValid(fact [32]byte) (bool, error) {
	return _IFactRegistry.Contract.IsValid(&_IFactRegistry.CallOpts, fact)
}

// IsValid is a free data retrieval call binding the contract method 0x6a938567.
//
// Solidity: function isValid(bytes32 fact) view returns(bool)
func (_IFactRegistry *IFactRegistryCallerSession) IsValid(fact [32]byte) (bool, error) {
	return _IFactRegistry.Contract.IsValid(&_IFactRegistry.CallOpts, fact)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
