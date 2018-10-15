// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethBind

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FileShareContractABI is the input ABI used to generate the binding from.
const FileShareContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shares\",\"outputs\":[{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fileHash\",\"type\":\"string\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"publishFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_shareID\",\"type\":\"uint256\"}],\"name\":\"acceptRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_shareID\",\"type\":\"uint256\"}],\"name\":\"acceptPossible\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"sharesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FileShareContractBin is the compiled bytecode used for deploying new contracts.
const FileShareContractBin = `0x608060405234801561001057600080fd5b50604051602080610898833981016040525160028054600160a060020a031916600160a060020a03909216919091179055610848806100506000396000f3006080604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663259ddefc811461007c57806373e01d4f146101855780638616abd814610224578063b55064ef14610248578063fbf7d3c514610280578063fc0c546a146102b3575b600080fd5b34801561008857600080fd5b506100a0600160a060020a03600435166024356102e4565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156100e75781810151838201526020016100cf565b50505050905090810190601f1680156101145780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561014757818101518382015260200161012f565b50505050905090810190601f1680156101745780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561019157600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261022294369492936024939284019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a9099940197509195509182019350915081908401838280828437509497506104469650505050505050565b005b34801561023057600080fd5b50610222600160a060020a03600435166024356104e2565b34801561025457600080fd5b5061026c600160a060020a036004351660243561063c565b604080519115158252519081900360200190f35b34801561028c57600080fd5b506102a1600160a060020a036004351661073e565b60408051918252519081900360200190f35b3480156102bf57600080fd5b506102c8610750565b60408051600160a060020a039092168252519081900360200190f35b6001602052816000526040600020818154811015156102ff57fe5b60009182526020918290206003919091020180546040805160026001841615610100026000190190931692909204601f8101859004850283018501909152808252919450925083918301828280156103985780601f1061036d57610100808354040283529160200191610398565b820191906000526020600020905b81548152906001019060200180831161037b57829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104365780601f1061040b57610100808354040283529160200191610436565b820191906000526020600020905b81548152906001019060200180831161041957829003601f168201915b5050505050908060020154905083565b61044e61075f565b506040805160608101825284815260208082018490528183018590523360009081526001808352938120805494850180825590825290829020835180519495929486946003909402909201926104a992849290910190610781565b5060208281015180516104c29260018501920190610781565b506040918201516002909101553360009081526020819052205550505050565b60006104ee838361063c565b15156104f957600080fd5b33600090815260016020526040902080548390811061051457fe5b60009182526020822060028054600390930290910190810154604080517f1d3aa9ce000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015260248201939093529051929550921692631d3aa9ce926044808201939182900301818387803b15801561059657600080fd5b505af11580156105aa573d6000803e3d6000fd5b50506002805490840154604080517f9e1a00aa000000000000000000000000000000000000000000000000000000008152336004820152602481019290925251600160a060020a039092169350639e1a00aa925060448082019260009290919082900301818387803b15801561061f57600080fd5b505af1158015610633573d6000803e3d6000fd5b50505050505050565b336000818152600160205260408120549091908290841061065c57600080fd5b600160a060020a038216600090815260016020526040902080548590811061068057fe5b600091825260208083206002600390930201828101549254604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a038c8116600483015289811660248301529151939750949591169363dd62ed3e93604480830194919391928390030190829087803b15801561070757600080fd5b505af115801561071b573d6000803e3d6000fd5b505050506040513d602081101561073157600080fd5b5051101595945050505050565b60006020819052908152604090205481565b600254600160a060020a031681565b6060604051908101604052806060815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107c257805160ff19168380011785556107ef565b828001600101855582156107ef579182015b828111156107ef5782518255916020019190600101906107d4565b506107fb9291506107ff565b5090565b61081991905b808211156107fb5760008155600101610805565b905600a165627a7a7230582018e41acac50d43a333c2d601a4fe8ce698189327e1225eb4f85a317e3d67514c0029`

// DeployFileShareContract deploys a new Ethereum contract, binding an instance of FileShareContract to it.
func DeployFileShareContract(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *FileShareContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileShareContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FileShareContractBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileShareContract{FileShareContractCaller: FileShareContractCaller{contract: contract}, FileShareContractTransactor: FileShareContractTransactor{contract: contract}, FileShareContractFilterer: FileShareContractFilterer{contract: contract}}, nil
}

// FileShareContract is an auto generated Go binding around an Ethereum contract.
type FileShareContract struct {
	FileShareContractCaller     // Read-only binding to the contract
	FileShareContractTransactor // Write-only binding to the contract
	FileShareContractFilterer   // Log filterer for contract events
}

// FileShareContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileShareContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileShareContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileShareContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileShareContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileShareContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileShareContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileShareContractSession struct {
	Contract     *FileShareContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FileShareContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileShareContractCallerSession struct {
	Contract *FileShareContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FileShareContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileShareContractTransactorSession struct {
	Contract     *FileShareContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FileShareContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileShareContractRaw struct {
	Contract *FileShareContract // Generic contract binding to access the raw methods on
}

// FileShareContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileShareContractCallerRaw struct {
	Contract *FileShareContractCaller // Generic read-only contract binding to access the raw methods on
}

// FileShareContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileShareContractTransactorRaw struct {
	Contract *FileShareContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileShareContract creates a new instance of FileShareContract, bound to a specific deployed contract.
func NewFileShareContract(address common.Address, backend bind.ContractBackend) (*FileShareContract, error) {
	contract, err := bindFileShareContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileShareContract{FileShareContractCaller: FileShareContractCaller{contract: contract}, FileShareContractTransactor: FileShareContractTransactor{contract: contract}, FileShareContractFilterer: FileShareContractFilterer{contract: contract}}, nil
}

// NewFileShareContractCaller creates a new read-only instance of FileShareContract, bound to a specific deployed contract.
func NewFileShareContractCaller(address common.Address, caller bind.ContractCaller) (*FileShareContractCaller, error) {
	contract, err := bindFileShareContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileShareContractCaller{contract: contract}, nil
}

// NewFileShareContractTransactor creates a new write-only instance of FileShareContract, bound to a specific deployed contract.
func NewFileShareContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FileShareContractTransactor, error) {
	contract, err := bindFileShareContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileShareContractTransactor{contract: contract}, nil
}

// NewFileShareContractFilterer creates a new log filterer instance of FileShareContract, bound to a specific deployed contract.
func NewFileShareContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FileShareContractFilterer, error) {
	contract, err := bindFileShareContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileShareContractFilterer{contract: contract}, nil
}

// bindFileShareContract binds a generic wrapper to an already deployed contract.
func bindFileShareContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileShareContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileShareContract *FileShareContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileShareContract.Contract.FileShareContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileShareContract *FileShareContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileShareContract.Contract.FileShareContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileShareContract *FileShareContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileShareContract.Contract.FileShareContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileShareContract *FileShareContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileShareContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileShareContract *FileShareContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileShareContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileShareContract *FileShareContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileShareContract.Contract.contract.Transact(opts, method, params...)
}

// AcceptPossible is a free data retrieval call binding the contract method 0xb55064ef.
//
// Solidity: function acceptPossible(_recipient address, _shareID uint256) constant returns(bool)
func (_FileShareContract *FileShareContractCaller) AcceptPossible(opts *bind.CallOpts, _recipient common.Address, _shareID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FileShareContract.contract.Call(opts, out, "acceptPossible", _recipient, _shareID)
	return *ret0, err
}

// AcceptPossible is a free data retrieval call binding the contract method 0xb55064ef.
//
// Solidity: function acceptPossible(_recipient address, _shareID uint256) constant returns(bool)
func (_FileShareContract *FileShareContractSession) AcceptPossible(_recipient common.Address, _shareID *big.Int) (bool, error) {
	return _FileShareContract.Contract.AcceptPossible(&_FileShareContract.CallOpts, _recipient, _shareID)
}

// AcceptPossible is a free data retrieval call binding the contract method 0xb55064ef.
//
// Solidity: function acceptPossible(_recipient address, _shareID uint256) constant returns(bool)
func (_FileShareContract *FileShareContractCallerSession) AcceptPossible(_recipient common.Address, _shareID *big.Int) (bool, error) {
	return _FileShareContract.Contract.AcceptPossible(&_FileShareContract.CallOpts, _recipient, _shareID)
}

// Shares is a free data retrieval call binding the contract method 0x259ddefc.
//
// Solidity: function shares( address,  uint256) constant returns(fileHash string, description string, price uint256)
func (_FileShareContract *FileShareContractCaller) Shares(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FileHash    string
	Description string
	Price       *big.Int
}, error) {
	ret := new(struct {
		FileHash    string
		Description string
		Price       *big.Int
	})
	out := ret
	err := _FileShareContract.contract.Call(opts, out, "shares", arg0, arg1)
	return *ret, err
}

// Shares is a free data retrieval call binding the contract method 0x259ddefc.
//
// Solidity: function shares( address,  uint256) constant returns(fileHash string, description string, price uint256)
func (_FileShareContract *FileShareContractSession) Shares(arg0 common.Address, arg1 *big.Int) (struct {
	FileHash    string
	Description string
	Price       *big.Int
}, error) {
	return _FileShareContract.Contract.Shares(&_FileShareContract.CallOpts, arg0, arg1)
}

// Shares is a free data retrieval call binding the contract method 0x259ddefc.
//
// Solidity: function shares( address,  uint256) constant returns(fileHash string, description string, price uint256)
func (_FileShareContract *FileShareContractCallerSession) Shares(arg0 common.Address, arg1 *big.Int) (struct {
	FileHash    string
	Description string
	Price       *big.Int
}, error) {
	return _FileShareContract.Contract.Shares(&_FileShareContract.CallOpts, arg0, arg1)
}

// SharesCount is a free data retrieval call binding the contract method 0xfbf7d3c5.
//
// Solidity: function sharesCount( address) constant returns(uint256)
func (_FileShareContract *FileShareContractCaller) SharesCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FileShareContract.contract.Call(opts, out, "sharesCount", arg0)
	return *ret0, err
}

// SharesCount is a free data retrieval call binding the contract method 0xfbf7d3c5.
//
// Solidity: function sharesCount( address) constant returns(uint256)
func (_FileShareContract *FileShareContractSession) SharesCount(arg0 common.Address) (*big.Int, error) {
	return _FileShareContract.Contract.SharesCount(&_FileShareContract.CallOpts, arg0)
}

// SharesCount is a free data retrieval call binding the contract method 0xfbf7d3c5.
//
// Solidity: function sharesCount( address) constant returns(uint256)
func (_FileShareContract *FileShareContractCallerSession) SharesCount(arg0 common.Address) (*big.Int, error) {
	return _FileShareContract.Contract.SharesCount(&_FileShareContract.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FileShareContract *FileShareContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FileShareContract.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FileShareContract *FileShareContractSession) Token() (common.Address, error) {
	return _FileShareContract.Contract.Token(&_FileShareContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FileShareContract *FileShareContractCallerSession) Token() (common.Address, error) {
	return _FileShareContract.Contract.Token(&_FileShareContract.CallOpts)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x8616abd8.
//
// Solidity: function acceptRequest(_recipient address, _shareID uint256) returns()
func (_FileShareContract *FileShareContractTransactor) AcceptRequest(opts *bind.TransactOpts, _recipient common.Address, _shareID *big.Int) (*types.Transaction, error) {
	return _FileShareContract.contract.Transact(opts, "acceptRequest", _recipient, _shareID)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x8616abd8.
//
// Solidity: function acceptRequest(_recipient address, _shareID uint256) returns()
func (_FileShareContract *FileShareContractSession) AcceptRequest(_recipient common.Address, _shareID *big.Int) (*types.Transaction, error) {
	return _FileShareContract.Contract.AcceptRequest(&_FileShareContract.TransactOpts, _recipient, _shareID)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x8616abd8.
//
// Solidity: function acceptRequest(_recipient address, _shareID uint256) returns()
func (_FileShareContract *FileShareContractTransactorSession) AcceptRequest(_recipient common.Address, _shareID *big.Int) (*types.Transaction, error) {
	return _FileShareContract.Contract.AcceptRequest(&_FileShareContract.TransactOpts, _recipient, _shareID)
}

// PublishFile is a paid mutator transaction binding the contract method 0x73e01d4f.
//
// Solidity: function publishFile(_fileHash string, _price uint256, _description string) returns()
func (_FileShareContract *FileShareContractTransactor) PublishFile(opts *bind.TransactOpts, _fileHash string, _price *big.Int, _description string) (*types.Transaction, error) {
	return _FileShareContract.contract.Transact(opts, "publishFile", _fileHash, _price, _description)
}

// PublishFile is a paid mutator transaction binding the contract method 0x73e01d4f.
//
// Solidity: function publishFile(_fileHash string, _price uint256, _description string) returns()
func (_FileShareContract *FileShareContractSession) PublishFile(_fileHash string, _price *big.Int, _description string) (*types.Transaction, error) {
	return _FileShareContract.Contract.PublishFile(&_FileShareContract.TransactOpts, _fileHash, _price, _description)
}

// PublishFile is a paid mutator transaction binding the contract method 0x73e01d4f.
//
// Solidity: function publishFile(_fileHash string, _price uint256, _description string) returns()
func (_FileShareContract *FileShareContractTransactorSession) PublishFile(_fileHash string, _price *big.Int, _description string) (*types.Transaction, error) {
	return _FileShareContract.Contract.PublishFile(&_FileShareContract.TransactOpts, _fileHash, _price, _description)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"claimAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"renounceAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"sendTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(from address, to address) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, from common.Address, to common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", from, to)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(from address, to address) constant returns(uint256)
func (_Token *TokenSession) Allowance(from common.Address, to common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, from, to)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(from address, to address) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(from common.Address, to common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, from, to)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Allow is a paid mutator transaction binding the contract method 0x6c6f31f2.
//
// Solidity: function allow(dest address, value uint256) returns()
func (_Token *TokenTransactor) Allow(opts *bind.TransactOpts, dest common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "allow", dest, value)
}

// Allow is a paid mutator transaction binding the contract method 0x6c6f31f2.
//
// Solidity: function allow(dest address, value uint256) returns()
func (_Token *TokenSession) Allow(dest common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Allow(&_Token.TransactOpts, dest, value)
}

// Allow is a paid mutator transaction binding the contract method 0x6c6f31f2.
//
// Solidity: function allow(dest address, value uint256) returns()
func (_Token *TokenTransactorSession) Allow(dest common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Allow(&_Token.TransactOpts, dest, value)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Token *TokenTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Token *TokenSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeOwner(&_Token.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Token *TokenTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeOwner(&_Token.TransactOpts, newOwner)
}

// ClaimAllowance is a paid mutator transaction binding the contract method 0x1d3aa9ce.
//
// Solidity: function claimAllowance(from address, value uint256) returns()
func (_Token *TokenTransactor) ClaimAllowance(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimAllowance", from, value)
}

// ClaimAllowance is a paid mutator transaction binding the contract method 0x1d3aa9ce.
//
// Solidity: function claimAllowance(from address, value uint256) returns()
func (_Token *TokenSession) ClaimAllowance(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAllowance(&_Token.TransactOpts, from, value)
}

// ClaimAllowance is a paid mutator transaction binding the contract method 0x1d3aa9ce.
//
// Solidity: function claimAllowance(from address, value uint256) returns()
func (_Token *TokenTransactorSession) ClaimAllowance(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ClaimAllowance(&_Token.TransactOpts, from, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_Token *TokenSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_Token *TokenTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, to, value)
}

// RenounceAllowance is a paid mutator transaction binding the contract method 0x7e21c2aa.
//
// Solidity: function renounceAllowance(to address, value uint256) returns()
func (_Token *TokenTransactor) RenounceAllowance(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceAllowance", to, value)
}

// RenounceAllowance is a paid mutator transaction binding the contract method 0x7e21c2aa.
//
// Solidity: function renounceAllowance(to address, value uint256) returns()
func (_Token *TokenSession) RenounceAllowance(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.RenounceAllowance(&_Token.TransactOpts, to, value)
}

// RenounceAllowance is a paid mutator transaction binding the contract method 0x7e21c2aa.
//
// Solidity: function renounceAllowance(to address, value uint256) returns()
func (_Token *TokenTransactorSession) RenounceAllowance(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.RenounceAllowance(&_Token.TransactOpts, to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_Token *TokenTransactor) SendTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "sendTo", to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_Token *TokenSession) SendTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SendTo(&_Token.TransactOpts, to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_Token *TokenTransactorSession) SendTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SendTo(&_Token.TransactOpts, to, value)
}

// TokenImplABI is the input ABI used to generate the binding from.
const TokenImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"claimAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"renounceAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"sendTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Allowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RevokedAllowance\",\"type\":\"event\"}]"

// TokenImplBin is the compiled bytecode used for deploying new contracts.
const TokenImplBin = `0x608060405234801561001057600080fd5b50604051610b77380380610b7783398101604090815281516020808401519284015191840180519094939093019261004e916001919086019061007f565b50815161006290600290602085019061007f565b50600355505060008054600160a060020a0319163317905561011a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c057805160ff19168380011785556100ed565b828001600101855582156100ed579182015b828111156100ed5782518255916020019190600101906100d2565b506100f99291506100fd565b5090565b61011791905b808211156100f95760008155600101610103565b90565b610a4e806101296000396000f3006080604052600436106100e55763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663047fc9aa81146100ea57806306fdde031461011157806318160ddd1461019b5780631d3aa9ce146101b057806327e235e3146101d657806340c10f19146101f757806355b6ed5c1461021b5780636c6f31f21461024257806370a08231146102665780637e21c2aa146102875780638da5cb5b146102ab57806395d89b41146102dc5780639e1a00aa146102f1578063a6f9dae114610315578063d5abeb0114610336578063dd62ed3e1461034b575b600080fd5b3480156100f657600080fd5b506100ff610372565b60408051918252519081900360200190f35b34801561011d57600080fd5b50610126610378565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610160578181015183820152602001610148565b50505050905090810190601f16801561018d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101a757600080fd5b506100ff610405565b3480156101bc57600080fd5b506101d4600160a060020a036004351660243561040b565b005b3480156101e257600080fd5b506100ff600160a060020a036004351661051e565b34801561020357600080fd5b506101d4600160a060020a0360043516602435610530565b34801561022757600080fd5b506100ff600160a060020a03600435811690602435166105f8565b34801561024e57600080fd5b506101d4600160a060020a0360043516602435610615565b34801561027257600080fd5b506100ff600160a060020a0360043516610718565b34801561029357600080fd5b506101d4600160a060020a0360043516602435610733565b3480156102b757600080fd5b506102c0610847565b60408051600160a060020a039092168252519081900360200190f35b3480156102e857600080fd5b50610126610856565b3480156102fd57600080fd5b506101d4600160a060020a03600435166024356108ae565b34801561032157600080fd5b506101d4600160a060020a03600435166109a8565b34801561034257600080fd5b506100ff6109f1565b34801561035757600080fd5b506100ff600160a060020a03600435811690602435166109f7565b60045481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103fd5780601f106103d2576101008083540402835291602001916103fd565b820191906000526020600020905b8154815290600101906020018083116103e057829003601f168201915b505050505081565b60045490565b600160a060020a03821660009081526006602090815260408083203384529091529020548110156104ac576040805160e560020a62461bcd02815260206004820152602a60248201527f43616e6e6f7420636c61696d206d6f7265207468616e2077686174204920616d60448201527f207065726d697474656400000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000818152600560209081526040808320805486019055600160a060020a0386168084526006835281842085855283529281902080548690039055805185815290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35050565b60056020526000908152604090205481565b6000543390600160a060020a0316811461054957600080fd5b600354600454830111156105cd576040805160e560020a62461bcd02815260206004820152602960248201527f43616e6e6f74206d696e74206d6f7265207468616e20746865206d6178696d7560448201527f6d20616c6c6f7765640000000000000000000000000000000000000000000000606482015290519081900360840190fd5b50600160a060020a039091166000908152600560205260409020805482019055600480549091019055565b600660209081526000928352604080842090915290825290205481565b3360008181526005602052604090205482908111156106a4576040805160e560020a62461bcd02815260206004820152602160248201527f43616e6e6f742073656e64206d6f726520636f696e73207468616e2049206f7760448201527f6e00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b3360008181526005602090815260408083208054889003905560068252808320600160a060020a03891680855290835292819020805488019055805187815290519293927e1184123078cd58061e15db3b045089949811fe5ca4206cf0eab78f30f58b1b929181900390910190a350505050565b600160a060020a031660009081526005602052604090205490565b336000908152600660209081526040808320600160a060020a03861684529091529020548110156107d4576040805160e560020a62461bcd02815260206004820152603660248201527f43616e6e6f742072656e6f756e6365206d6f7265207468616e2077686174204960448201527f20616c6c6f77656420746f207468652077616c6c657400000000000000000000606482015290519081900360840190fd5b33600081815260056020908152604080832080548601905560068252808320600160a060020a0387168085529083529281902080548690039055805185815290519293927f36ae076eee364f8f0a11eb80c0e39cce7f8054d9a3467cb0e7ab71c0af046ddc929181900390910190a35050565b600054600160a060020a031681565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103fd5780601f106103d2576101008083540402835291602001916103fd565b33600081815260056020526040902054829081111561093d576040805160e560020a62461bcd02815260206004820152602160248201527f43616e6e6f742073656e64206d6f726520636f696e73207468616e2049206f7760448201527f6e00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600081815260056020908152604080832080548890039055600160a060020a03881680845292819020805488019055805187815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350505050565b6000543390600160a060020a031681146109c157600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60035481565b600160a060020a039182166000908152600660209081526040808320939094168252919091522054905600a165627a7a7230582057230e9321790a122f6915c637d2cd1b148c3f4094662264d502c7ec6db997330029`

// DeployTokenImpl deploys a new Ethereum contract, binding an instance of TokenImpl to it.
func DeployTokenImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _maxSupply *big.Int) (common.Address, *types.Transaction, *TokenImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenImplBin), backend, _name, _symbol, _maxSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenImpl{TokenImplCaller: TokenImplCaller{contract: contract}, TokenImplTransactor: TokenImplTransactor{contract: contract}, TokenImplFilterer: TokenImplFilterer{contract: contract}}, nil
}

// TokenImpl is an auto generated Go binding around an Ethereum contract.
type TokenImpl struct {
	TokenImplCaller     // Read-only binding to the contract
	TokenImplTransactor // Write-only binding to the contract
	TokenImplFilterer   // Log filterer for contract events
}

// TokenImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenImplSession struct {
	Contract     *TokenImpl        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenImplCallerSession struct {
	Contract *TokenImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenImplTransactorSession struct {
	Contract     *TokenImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenImplRaw struct {
	Contract *TokenImpl // Generic contract binding to access the raw methods on
}

// TokenImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenImplCallerRaw struct {
	Contract *TokenImplCaller // Generic read-only contract binding to access the raw methods on
}

// TokenImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenImplTransactorRaw struct {
	Contract *TokenImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenImpl creates a new instance of TokenImpl, bound to a specific deployed contract.
func NewTokenImpl(address common.Address, backend bind.ContractBackend) (*TokenImpl, error) {
	contract, err := bindTokenImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenImpl{TokenImplCaller: TokenImplCaller{contract: contract}, TokenImplTransactor: TokenImplTransactor{contract: contract}, TokenImplFilterer: TokenImplFilterer{contract: contract}}, nil
}

// NewTokenImplCaller creates a new read-only instance of TokenImpl, bound to a specific deployed contract.
func NewTokenImplCaller(address common.Address, caller bind.ContractCaller) (*TokenImplCaller, error) {
	contract, err := bindTokenImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenImplCaller{contract: contract}, nil
}

// NewTokenImplTransactor creates a new write-only instance of TokenImpl, bound to a specific deployed contract.
func NewTokenImplTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenImplTransactor, error) {
	contract, err := bindTokenImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenImplTransactor{contract: contract}, nil
}

// NewTokenImplFilterer creates a new log filterer instance of TokenImpl, bound to a specific deployed contract.
func NewTokenImplFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenImplFilterer, error) {
	contract, err := bindTokenImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenImplFilterer{contract: contract}, nil
}

// bindTokenImpl binds a generic wrapper to an already deployed contract.
func bindTokenImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenImpl *TokenImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenImpl.Contract.TokenImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenImpl *TokenImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenImpl.Contract.TokenImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenImpl *TokenImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenImpl.Contract.TokenImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenImpl *TokenImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenImpl *TokenImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenImpl *TokenImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenImpl.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(from address, to address) constant returns(uint256)
func (_TokenImpl *TokenImplCaller) Allowance(opts *bind.CallOpts, from common.Address, to common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "allowance", from, to)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(from address, to address) constant returns(uint256)
func (_TokenImpl *TokenImplSession) Allowance(from common.Address, to common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.Allowance(&_TokenImpl.CallOpts, from, to)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(from address, to address) constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) Allowance(from common.Address, to common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.Allowance(&_TokenImpl.CallOpts, from, to)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances( address,  address) constant returns(uint256)
func (_TokenImpl *TokenImplCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "allowances", arg0, arg1)
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances( address,  address) constant returns(uint256)
func (_TokenImpl *TokenImplSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.Allowances(&_TokenImpl.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances( address,  address) constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.Allowances(&_TokenImpl.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TokenImpl *TokenImplCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TokenImpl *TokenImplSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.BalanceOf(&_TokenImpl.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.BalanceOf(&_TokenImpl.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TokenImpl *TokenImplCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TokenImpl *TokenImplSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.Balances(&_TokenImpl.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TokenImpl.Contract.Balances(&_TokenImpl.CallOpts, arg0)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_TokenImpl *TokenImplCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "maxSupply")
	return *ret0, err
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_TokenImpl *TokenImplSession) MaxSupply() (*big.Int, error) {
	return _TokenImpl.Contract.MaxSupply(&_TokenImpl.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) MaxSupply() (*big.Int, error) {
	return _TokenImpl.Contract.MaxSupply(&_TokenImpl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenImpl *TokenImplCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenImpl *TokenImplSession) Name() (string, error) {
	return _TokenImpl.Contract.Name(&_TokenImpl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenImpl *TokenImplCallerSession) Name() (string, error) {
	return _TokenImpl.Contract.Name(&_TokenImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenImpl *TokenImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenImpl *TokenImplSession) Owner() (common.Address, error) {
	return _TokenImpl.Contract.Owner(&_TokenImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenImpl *TokenImplCallerSession) Owner() (common.Address, error) {
	return _TokenImpl.Contract.Owner(&_TokenImpl.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_TokenImpl *TokenImplCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "supply")
	return *ret0, err
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_TokenImpl *TokenImplSession) Supply() (*big.Int, error) {
	return _TokenImpl.Contract.Supply(&_TokenImpl.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) Supply() (*big.Int, error) {
	return _TokenImpl.Contract.Supply(&_TokenImpl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenImpl *TokenImplCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenImpl *TokenImplSession) Symbol() (string, error) {
	return _TokenImpl.Contract.Symbol(&_TokenImpl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenImpl *TokenImplCallerSession) Symbol() (string, error) {
	return _TokenImpl.Contract.Symbol(&_TokenImpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenImpl *TokenImplCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenImpl.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenImpl *TokenImplSession) TotalSupply() (*big.Int, error) {
	return _TokenImpl.Contract.TotalSupply(&_TokenImpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenImpl *TokenImplCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenImpl.Contract.TotalSupply(&_TokenImpl.CallOpts)
}

// Allow is a paid mutator transaction binding the contract method 0x6c6f31f2.
//
// Solidity: function allow(dest address, value uint256) returns()
func (_TokenImpl *TokenImplTransactor) Allow(opts *bind.TransactOpts, dest common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.contract.Transact(opts, "allow", dest, value)
}

// Allow is a paid mutator transaction binding the contract method 0x6c6f31f2.
//
// Solidity: function allow(dest address, value uint256) returns()
func (_TokenImpl *TokenImplSession) Allow(dest common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.Allow(&_TokenImpl.TransactOpts, dest, value)
}

// Allow is a paid mutator transaction binding the contract method 0x6c6f31f2.
//
// Solidity: function allow(dest address, value uint256) returns()
func (_TokenImpl *TokenImplTransactorSession) Allow(dest common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.Allow(&_TokenImpl.TransactOpts, dest, value)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_TokenImpl *TokenImplTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenImpl.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_TokenImpl *TokenImplSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TokenImpl.Contract.ChangeOwner(&_TokenImpl.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_TokenImpl *TokenImplTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TokenImpl.Contract.ChangeOwner(&_TokenImpl.TransactOpts, newOwner)
}

// ClaimAllowance is a paid mutator transaction binding the contract method 0x1d3aa9ce.
//
// Solidity: function claimAllowance(from address, value uint256) returns()
func (_TokenImpl *TokenImplTransactor) ClaimAllowance(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.contract.Transact(opts, "claimAllowance", from, value)
}

// ClaimAllowance is a paid mutator transaction binding the contract method 0x1d3aa9ce.
//
// Solidity: function claimAllowance(from address, value uint256) returns()
func (_TokenImpl *TokenImplSession) ClaimAllowance(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.ClaimAllowance(&_TokenImpl.TransactOpts, from, value)
}

// ClaimAllowance is a paid mutator transaction binding the contract method 0x1d3aa9ce.
//
// Solidity: function claimAllowance(from address, value uint256) returns()
func (_TokenImpl *TokenImplTransactorSession) ClaimAllowance(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.ClaimAllowance(&_TokenImpl.TransactOpts, from, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_TokenImpl *TokenImplTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_TokenImpl *TokenImplSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.Mint(&_TokenImpl.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_TokenImpl *TokenImplTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.Mint(&_TokenImpl.TransactOpts, to, value)
}

// RenounceAllowance is a paid mutator transaction binding the contract method 0x7e21c2aa.
//
// Solidity: function renounceAllowance(to address, value uint256) returns()
func (_TokenImpl *TokenImplTransactor) RenounceAllowance(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.contract.Transact(opts, "renounceAllowance", to, value)
}

// RenounceAllowance is a paid mutator transaction binding the contract method 0x7e21c2aa.
//
// Solidity: function renounceAllowance(to address, value uint256) returns()
func (_TokenImpl *TokenImplSession) RenounceAllowance(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.RenounceAllowance(&_TokenImpl.TransactOpts, to, value)
}

// RenounceAllowance is a paid mutator transaction binding the contract method 0x7e21c2aa.
//
// Solidity: function renounceAllowance(to address, value uint256) returns()
func (_TokenImpl *TokenImplTransactorSession) RenounceAllowance(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.RenounceAllowance(&_TokenImpl.TransactOpts, to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_TokenImpl *TokenImplTransactor) SendTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.contract.Transact(opts, "sendTo", to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_TokenImpl *TokenImplSession) SendTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.SendTo(&_TokenImpl.TransactOpts, to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_TokenImpl *TokenImplTransactorSession) SendTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenImpl.Contract.SendTo(&_TokenImpl.TransactOpts, to, value)
}

// TokenImplAllowedIterator is returned from FilterAllowed and is used to iterate over the raw logs and unpacked data for Allowed events raised by the TokenImpl contract.
type TokenImplAllowedIterator struct {
	Event *TokenImplAllowed // Event containing the contract specifics and raw log

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
func (it *TokenImplAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenImplAllowed)
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
		it.Event = new(TokenImplAllowed)
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
func (it *TokenImplAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenImplAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenImplAllowed represents a Allowed event raised by the TokenImpl contract.
type TokenImplAllowed struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAllowed is a free log retrieval operation binding the contract event 0x001184123078cd58061e15db3b045089949811fe5ca4206cf0eab78f30f58b1b.
//
// Solidity: e Allowed(from indexed address, to indexed address, value uint256)
func (_TokenImpl *TokenImplFilterer) FilterAllowed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenImplAllowedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenImpl.contract.FilterLogs(opts, "Allowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenImplAllowedIterator{contract: _TokenImpl.contract, event: "Allowed", logs: logs, sub: sub}, nil
}

// WatchAllowed is a free log subscription operation binding the contract event 0x001184123078cd58061e15db3b045089949811fe5ca4206cf0eab78f30f58b1b.
//
// Solidity: e Allowed(from indexed address, to indexed address, value uint256)
func (_TokenImpl *TokenImplFilterer) WatchAllowed(opts *bind.WatchOpts, sink chan<- *TokenImplAllowed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenImpl.contract.WatchLogs(opts, "Allowed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenImplAllowed)
				if err := _TokenImpl.contract.UnpackLog(event, "Allowed", log); err != nil {
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

// TokenImplRevokedAllowanceIterator is returned from FilterRevokedAllowance and is used to iterate over the raw logs and unpacked data for RevokedAllowance events raised by the TokenImpl contract.
type TokenImplRevokedAllowanceIterator struct {
	Event *TokenImplRevokedAllowance // Event containing the contract specifics and raw log

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
func (it *TokenImplRevokedAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenImplRevokedAllowance)
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
		it.Event = new(TokenImplRevokedAllowance)
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
func (it *TokenImplRevokedAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenImplRevokedAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenImplRevokedAllowance represents a RevokedAllowance event raised by the TokenImpl contract.
type TokenImplRevokedAllowance struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevokedAllowance is a free log retrieval operation binding the contract event 0x36ae076eee364f8f0a11eb80c0e39cce7f8054d9a3467cb0e7ab71c0af046ddc.
//
// Solidity: e RevokedAllowance(from indexed address, to indexed address, value uint256)
func (_TokenImpl *TokenImplFilterer) FilterRevokedAllowance(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenImplRevokedAllowanceIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenImpl.contract.FilterLogs(opts, "RevokedAllowance", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenImplRevokedAllowanceIterator{contract: _TokenImpl.contract, event: "RevokedAllowance", logs: logs, sub: sub}, nil
}

// WatchRevokedAllowance is a free log subscription operation binding the contract event 0x36ae076eee364f8f0a11eb80c0e39cce7f8054d9a3467cb0e7ab71c0af046ddc.
//
// Solidity: e RevokedAllowance(from indexed address, to indexed address, value uint256)
func (_TokenImpl *TokenImplFilterer) WatchRevokedAllowance(opts *bind.WatchOpts, sink chan<- *TokenImplRevokedAllowance, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenImpl.contract.WatchLogs(opts, "RevokedAllowance", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenImplRevokedAllowance)
				if err := _TokenImpl.contract.UnpackLog(event, "RevokedAllowance", log); err != nil {
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

// TokenImplTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenImpl contract.
type TokenImplTransferIterator struct {
	Event *TokenImplTransfer // Event containing the contract specifics and raw log

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
func (it *TokenImplTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenImplTransfer)
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
		it.Event = new(TokenImplTransfer)
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
func (it *TokenImplTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenImplTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenImplTransfer represents a Transfer event raised by the TokenImpl contract.
type TokenImplTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_TokenImpl *TokenImplFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenImplTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenImpl.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenImplTransferIterator{contract: _TokenImpl.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_TokenImpl *TokenImplFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenImplTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenImpl.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenImplTransfer)
				if err := _TokenImpl.contract.UnpackLog(event, "Transfer", log); err != nil {
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
