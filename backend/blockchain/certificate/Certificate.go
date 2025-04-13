// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package certificate

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

// CertificateMetaData contains all meta data concerning the Certificate contract.
var CertificateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"certificateId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"CertificateIssued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"certificateHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"certificateId\",\"type\":\"string\"}],\"name\":\"getCertificateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"certificateId\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"storeCertificateHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061050e8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630499ba6f146100465780630a88deaf1461007657806335b8f365146100a6575b600080fd5b610060600480360381019061005b91906102d1565b6100c2565b60405161006d9190610333565b60405180910390f35b610090600480360381019061008b91906102d1565b6100e9565b60405161009d9190610333565b60405180910390f35b6100c060048036038101906100bb919061037a565b610117565b005b600080826040516100d39190610447565b9081526020016040518091039020549050919050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b806000836040516101289190610447565b9081526020016040518091039020819055507fc1255beb592123a69dbf68648642ecc5fca6e86afc2fd709b8173c8d831b03d2828260405161016b9291906104a8565b60405180910390a15050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101de82610195565b810181811067ffffffffffffffff821117156101fd576101fc6101a6565b5b80604052505050565b6000610210610177565b905061021c82826101d5565b919050565b600067ffffffffffffffff82111561023c5761023b6101a6565b5b61024582610195565b9050602081019050919050565b82818337600083830152505050565b600061027461026f84610221565b610206565b9050828152602081018484840111156102905761028f610190565b5b61029b848285610252565b509392505050565b600082601f8301126102b8576102b761018b565b5b81356102c8848260208601610261565b91505092915050565b6000602082840312156102e7576102e6610181565b5b600082013567ffffffffffffffff81111561030557610304610186565b5b610311848285016102a3565b91505092915050565b6000819050919050565b61032d8161031a565b82525050565b60006020820190506103486000830184610324565b92915050565b6103578161031a565b811461036257600080fd5b50565b6000813590506103748161034e565b92915050565b6000806040838503121561039157610390610181565b5b600083013567ffffffffffffffff8111156103af576103ae610186565b5b6103bb858286016102a3565b92505060206103cc85828601610365565b9150509250929050565b600081519050919050565b600081905092915050565b60005b8381101561040a5780820151818401526020810190506103ef565b60008484015250505050565b6000610421826103d6565b61042b81856103e1565b935061043b8185602086016103ec565b80840191505092915050565b60006104538284610416565b915081905092915050565b600082825260208201905092915050565b600061047a826103d6565b610484818561045e565b93506104948185602086016103ec565b61049d81610195565b840191505092915050565b600060408201905081810360008301526104c2818561046f565b90506104d16020830184610324565b939250505056fea2646970667358221220e19884cf9e3d68ed1f912aa3befaed1b876607f1cef6d2757eaae8ea75d01e9d64736f6c634300081c0033",
}

// CertificateABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateMetaData.ABI instead.
var CertificateABI = CertificateMetaData.ABI

// CertificateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CertificateMetaData.Bin instead.
var CertificateBin = CertificateMetaData.Bin

// DeployCertificate deploys a new Ethereum contract, binding an instance of Certificate to it.
func DeployCertificate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Certificate, error) {
	parsed, err := CertificateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CertificateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Certificate{CertificateCaller: CertificateCaller{contract: contract}, CertificateTransactor: CertificateTransactor{contract: contract}, CertificateFilterer: CertificateFilterer{contract: contract}}, nil
}

// Certificate is an auto generated Go binding around an Ethereum contract.
type Certificate struct {
	CertificateCaller     // Read-only binding to the contract
	CertificateTransactor // Write-only binding to the contract
	CertificateFilterer   // Log filterer for contract events
}

// CertificateCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateSession struct {
	Contract     *Certificate      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertificateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateCallerSession struct {
	Contract *CertificateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CertificateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateTransactorSession struct {
	Contract     *CertificateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CertificateRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateRaw struct {
	Contract *Certificate // Generic contract binding to access the raw methods on
}

// CertificateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateCallerRaw struct {
	Contract *CertificateCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateTransactorRaw struct {
	Contract *CertificateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificate creates a new instance of Certificate, bound to a specific deployed contract.
func NewCertificate(address common.Address, backend bind.ContractBackend) (*Certificate, error) {
	contract, err := bindCertificate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Certificate{CertificateCaller: CertificateCaller{contract: contract}, CertificateTransactor: CertificateTransactor{contract: contract}, CertificateFilterer: CertificateFilterer{contract: contract}}, nil
}

// NewCertificateCaller creates a new read-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateCaller(address common.Address, caller bind.ContractCaller) (*CertificateCaller, error) {
	contract, err := bindCertificate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateCaller{contract: contract}, nil
}

// NewCertificateTransactor creates a new write-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateTransactor, error) {
	contract, err := bindCertificate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateTransactor{contract: contract}, nil
}

// NewCertificateFilterer creates a new log filterer instance of Certificate, bound to a specific deployed contract.
func NewCertificateFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateFilterer, error) {
	contract, err := bindCertificate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateFilterer{contract: contract}, nil
}

// bindCertificate binds a generic wrapper to an already deployed contract.
func bindCertificate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertificateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.CertificateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transact(opts, method, params...)
}

// CertificateHashes is a free data retrieval call binding the contract method 0x0a88deaf.
//
// Solidity: function certificateHashes(string ) view returns(bytes32)
func (_Certificate *CertificateCaller) CertificateHashes(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "certificateHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CertificateHashes is a free data retrieval call binding the contract method 0x0a88deaf.
//
// Solidity: function certificateHashes(string ) view returns(bytes32)
func (_Certificate *CertificateSession) CertificateHashes(arg0 string) ([32]byte, error) {
	return _Certificate.Contract.CertificateHashes(&_Certificate.CallOpts, arg0)
}

// CertificateHashes is a free data retrieval call binding the contract method 0x0a88deaf.
//
// Solidity: function certificateHashes(string ) view returns(bytes32)
func (_Certificate *CertificateCallerSession) CertificateHashes(arg0 string) ([32]byte, error) {
	return _Certificate.Contract.CertificateHashes(&_Certificate.CallOpts, arg0)
}

// GetCertificateHash is a free data retrieval call binding the contract method 0x0499ba6f.
//
// Solidity: function getCertificateHash(string certificateId) view returns(bytes32)
func (_Certificate *CertificateCaller) GetCertificateHash(opts *bind.CallOpts, certificateId string) ([32]byte, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "getCertificateHash", certificateId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCertificateHash is a free data retrieval call binding the contract method 0x0499ba6f.
//
// Solidity: function getCertificateHash(string certificateId) view returns(bytes32)
func (_Certificate *CertificateSession) GetCertificateHash(certificateId string) ([32]byte, error) {
	return _Certificate.Contract.GetCertificateHash(&_Certificate.CallOpts, certificateId)
}

// GetCertificateHash is a free data retrieval call binding the contract method 0x0499ba6f.
//
// Solidity: function getCertificateHash(string certificateId) view returns(bytes32)
func (_Certificate *CertificateCallerSession) GetCertificateHash(certificateId string) ([32]byte, error) {
	return _Certificate.Contract.GetCertificateHash(&_Certificate.CallOpts, certificateId)
}

// StoreCertificateHash is a paid mutator transaction binding the contract method 0x35b8f365.
//
// Solidity: function storeCertificateHash(string certificateId, bytes32 hash) returns()
func (_Certificate *CertificateTransactor) StoreCertificateHash(opts *bind.TransactOpts, certificateId string, hash [32]byte) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "storeCertificateHash", certificateId, hash)
}

// StoreCertificateHash is a paid mutator transaction binding the contract method 0x35b8f365.
//
// Solidity: function storeCertificateHash(string certificateId, bytes32 hash) returns()
func (_Certificate *CertificateSession) StoreCertificateHash(certificateId string, hash [32]byte) (*types.Transaction, error) {
	return _Certificate.Contract.StoreCertificateHash(&_Certificate.TransactOpts, certificateId, hash)
}

// StoreCertificateHash is a paid mutator transaction binding the contract method 0x35b8f365.
//
// Solidity: function storeCertificateHash(string certificateId, bytes32 hash) returns()
func (_Certificate *CertificateTransactorSession) StoreCertificateHash(certificateId string, hash [32]byte) (*types.Transaction, error) {
	return _Certificate.Contract.StoreCertificateHash(&_Certificate.TransactOpts, certificateId, hash)
}

// CertificateCertificateIssuedIterator is returned from FilterCertificateIssued and is used to iterate over the raw logs and unpacked data for CertificateIssued events raised by the Certificate contract.
type CertificateCertificateIssuedIterator struct {
	Event *CertificateCertificateIssued // Event containing the contract specifics and raw log

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
func (it *CertificateCertificateIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateCertificateIssued)
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
		it.Event = new(CertificateCertificateIssued)
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
func (it *CertificateCertificateIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateCertificateIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateCertificateIssued represents a CertificateIssued event raised by the Certificate contract.
type CertificateCertificateIssued struct {
	CertificateId string
	Hash          [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCertificateIssued is a free log retrieval operation binding the contract event 0xc1255beb592123a69dbf68648642ecc5fca6e86afc2fd709b8173c8d831b03d2.
//
// Solidity: event CertificateIssued(string certificateId, bytes32 hash)
func (_Certificate *CertificateFilterer) FilterCertificateIssued(opts *bind.FilterOpts) (*CertificateCertificateIssuedIterator, error) {

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "CertificateIssued")
	if err != nil {
		return nil, err
	}
	return &CertificateCertificateIssuedIterator{contract: _Certificate.contract, event: "CertificateIssued", logs: logs, sub: sub}, nil
}

// WatchCertificateIssued is a free log subscription operation binding the contract event 0xc1255beb592123a69dbf68648642ecc5fca6e86afc2fd709b8173c8d831b03d2.
//
// Solidity: event CertificateIssued(string certificateId, bytes32 hash)
func (_Certificate *CertificateFilterer) WatchCertificateIssued(opts *bind.WatchOpts, sink chan<- *CertificateCertificateIssued) (event.Subscription, error) {

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "CertificateIssued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateCertificateIssued)
				if err := _Certificate.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
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

// ParseCertificateIssued is a log parse operation binding the contract event 0xc1255beb592123a69dbf68648642ecc5fca6e86afc2fd709b8173c8d831b03d2.
//
// Solidity: event CertificateIssued(string certificateId, bytes32 hash)
func (_Certificate *CertificateFilterer) ParseCertificateIssued(log types.Log) (*CertificateCertificateIssued, error) {
	event := new(CertificateCertificateIssued)
	if err := _Certificate.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
