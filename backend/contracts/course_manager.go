// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"courseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"coursePrice\",\"type\":\"uint256\"}],\"name\":\"deployCourse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"courseId\",\"type\":\"uint256\"}],\"name\":\"getCourse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}],",
	Bin: "0x608060405234801561000f575f80fd5b50610b638061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063264a1d93146100385780638b11014814610068575b5f80fd5b610052600480360381019061004d919061025c565b610084565b60405161005f91906102a9565b60405180910390f35b610082600480360381019061007d91906102c2565b6100f8565b005b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b5f338284604051610108906101be565b6101149392919061030f565b604051809103905ff08015801561012d573d5f803e3d5ffd5b509050805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6107e98061034583390190565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101f8826101cf565b9050919050565b610208816101ee565b8114610212575f80fd5b50565b5f81359050610223816101ff565b92915050565b5f819050919050565b61023b81610229565b8114610245575f80fd5b50565b5f8135905061025681610232565b92915050565b5f8060408385031215610272576102716101cb565b5b5f61027f85828601610215565b925050602061029085828601610248565b9150509250929050565b6102a3816101ee565b82525050565b5f6020820190506102bc5f83018461029a565b92915050565b5f80604083850312156102d8576102d76101cb565b5b5f6102e585828601610248565b92505060206102f685828601610248565b9150509250929050565b61030981610229565b82525050565b5f6060820190506103225f83018661029a565b61032f6020830185610300565b61033c6040830184610300565b94935050505056fe608060405234801561000f575f80fd5b506040516107e93803806107e983398181016040528101906100319190610117565b825f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160028190555080600381905550505050610167565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100b38261008a565b9050919050565b6100c3816100a9565b81146100cd575f80fd5b50565b5f815190506100de816100ba565b92915050565b5f819050919050565b6100f6816100e4565b8114610100575f80fd5b50565b5f81519050610111816100ed565b92915050565b5f805f6060848603121561012e5761012d610086565b5b5f61013b868287016100d0565b935050602061014c86828701610103565b925050604061015d86828701610103565b9150509250925092565b610675806101745f395ff3fe608060405260043610610042575f3560e01c8063264c8071146102325780638da5cb5b1461026e5780639658310c14610298578063c43a4b21146102b45761022e565b3661022e5760015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156100d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c89061040d565b60405180910390fd5b600254341015610116576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010d9061049b565b60405180910390fd5b5f6002543461012591906104ef565b90505f811115610174573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610172573d5f803e3d5ffd5b505b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60025490811502906040515f60405180830381858888f193505050501580156101d8573d5f803e3d5ffd5b506001805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055005b5f80fd5b34801561023d575f80fd5b5061025860048036038101906102539190610580565b6102de565b60405161026591906105c5565b60405180910390f35b348015610279575f80fd5b50610282610330565b60405161028f91906105fe565b60405180910390f35b6102b260048036038101906102ad9190610580565b610353565b005b3480156102bf575f80fd5b506102c86103aa565b6040516102d59190610626565b60405180910390f35b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555050565b5f600254905090565b5f82825260208201905092915050565b7f416c72656164792070757263686173656420636f7572736500000000000000005f82015250565b5f6103f76018836103b3565b9150610402826103c3565b602082019050919050565b5f6020820190508181035f830152610424816103eb565b9050919050565b7f4e6f7420656e6f7567682066756e647320666f72207472616e73616374696f6e5f8201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f6104856021836103b3565b91506104908261042b565b604082019050919050565b5f6020820190508181035f8301526104b281610479565b9050919050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104f9826104b9565b9150610504836104b9565b925082820390508181111561051c5761051b6104c2565b5b92915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61054f82610526565b9050919050565b61055f81610545565b8114610569575f80fd5b50565b5f8135905061057a81610556565b92915050565b5f6020828403121561059557610594610522565b5b5f6105a28482850161056c565b91505092915050565b5f8115159050919050565b6105bf816105ab565b82525050565b5f6020820190506105d85f8301846105b6565b92915050565b5f6105e882610526565b9050919050565b6105f8816105de565b82525050565b5f6020820190506106115f8301846105ef565b92915050565b610620816104b9565b82525050565b5f6020820190506106395f830184610617565b9291505056fea2646970667358221220363e0dab936f7676974d65d77354f204ac984f02f1b4b7968926f5520169e0a764736f6c63430008140033a264697066735822122001cf85e3549ff2bb2e4946160204117aaaa0e641bcfacce259298acf5f39b9d664736f6c63430008140033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetCourse is a free data retrieval call binding the contract method 0x264a1d93.
//
// Solidity: function getCourse(address owner, uint256 courseId) view returns(address)
func (_Contracts *ContractsCaller) GetCourse(opts *bind.CallOpts, owner common.Address, courseId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCourse", owner, courseId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCourse is a free data retrieval call binding the contract method 0x264a1d93.
//
// Solidity: function getCourse(address owner, uint256 courseId) view returns(address)
func (_Contracts *ContractsSession) GetCourse(owner common.Address, courseId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetCourse(&_Contracts.CallOpts, owner, courseId)
}

// GetCourse is a free data retrieval call binding the contract method 0x264a1d93.
//
// Solidity: function getCourse(address owner, uint256 courseId) view returns(address)
func (_Contracts *ContractsCallerSession) GetCourse(owner common.Address, courseId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetCourse(&_Contracts.CallOpts, owner, courseId)
}

// DeployCourse is a paid mutator transaction binding the contract method 0x8b110148.
//
// Solidity: function deployCourse(uint256 courseId, uint256 coursePrice) returns()
func (_Contracts *ContractsTransactor) DeployCourse(opts *bind.TransactOpts, courseId *big.Int, coursePrice *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deployCourse", courseId, coursePrice)
}

// DeployCourse is a paid mutator transaction binding the contract method 0x8b110148.
//
// Solidity: function deployCourse(uint256 courseId, uint256 coursePrice) returns()
func (_Contracts *ContractsSession) DeployCourse(courseId *big.Int, coursePrice *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeployCourse(&_Contracts.TransactOpts, courseId, coursePrice)
}

// DeployCourse is a paid mutator transaction binding the contract method 0x8b110148.
//
// Solidity: function deployCourse(uint256 courseId, uint256 coursePrice) returns()
func (_Contracts *ContractsTransactorSession) DeployCourse(courseId *big.Int, coursePrice *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeployCourse(&_Contracts.TransactOpts, courseId, coursePrice)
}
