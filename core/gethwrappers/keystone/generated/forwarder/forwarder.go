// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package forwarder

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

var KeystoneForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DuplicateSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSigners\",\"type\":\"uint256\"}],\"name\":\"ExcessSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FaultToleranceMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSigners\",\"type\":\"uint256\"}],\"name\":\"InsufficientSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"donId\",\"type\":\"bytes4\"}],\"name\":\"InvalidDonId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportAlreadyProcessed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"workflowOwner\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"workflowExecutionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ReportDelivered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"workflowOwner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"workflowExecutionId\",\"type\":\"bytes32\"}],\"name\":\"getTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawReport\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"donId\",\"type\":\"bytes4\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611692806101576000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146101895780639b6f55cd146101a7578063c0965dc3146101ba578063f2fde38b146101cd57600080fd5b8063181f5a771461008257806330670580146100ca57806379ba50971461017f575b600080fd5b604080518082018252601781527f4b657973746f6e65466f7277617264657220312e302e30000000000000000000602082015290516100c19190611225565b60405180910390f35b61015a6100d83660046112ba565b6040805160609490941b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020808601919091526034850193909352605480850192909252805180850390920182526074909301835280519082012060009081526003909152205473ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c1565b6101876101e0565b005b60005473ffffffffffffffffffffffffffffffffffffffff1661015a565b6101876101b5366004611339565b6102e2565b6101876101c83660046113c9565b6107a3565b6101876101db366004611478565b610ebd565b60015473ffffffffffffffffffffffffffffffffffffffff163314610266576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60015474010000000000000000000000000000000000000000900460ff1615610337576040517f37ed32e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905560ff83166000036103b3576040517f0743bae600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601f8111156103f8576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101829052601f602482015260440161025d565b6104038360036114c9565b60ff16811161046157806104188460036114c9565b6104239060016114ec565b6040517f9dd9e6d8000000000000000000000000000000000000000000000000000000008152600481019290925260ff16602482015260440161025d565b60005b7fffffffff000000000000000000000000000000000000000000000000000000008516600090815260026020526040902060010154811015610556577fffffffff00000000000000000000000000000000000000000000000000000000851660009081526002602052604081206001018054839081106104e6576104e661150b565b60009182526020808320909101547fffffffff00000000000000000000000000000000000000000000000000000000891683526002808352604080852073ffffffffffffffffffffffffffffffffffffffff909316855291019091528120555061054f8161153a565b9050610464565b507fffffffff0000000000000000000000000000000000000000000000000000000084166000908152600260205260409020610596906001018383611169565b5060005b818110156107145760008383838181106105b6576105b661150b565b90506020020160208101906105cb9190611478565b7fffffffff000000000000000000000000000000000000000000000000000000008716600090815260026020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290920190529020549091501561065a576040517f8044bb3300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106658260016114ec565b7fffffffff000000000000000000000000000000000000000000000000000000008716600090815260026020818152604080842073ffffffffffffffffffffffffffffffffffffffff909616808552868401835290842060ff959095169094559081526001938401805494850181558252902090910180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909117905561070d8161153a565b905061059a565b5050507fffffffff0000000000000000000000000000000000000000000000000000000091909116600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff169055565b60015474010000000000000000000000000000000000000000900460ff16156107f8576040517f37ed32e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556064831015610872576040517fb55ac75400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806000806108b788888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ed192505050565b7fffffffff0000000000000000000000000000000000000000000000000000000083166000908152600260205260408120549498509296509094509250906109039060ff1660016114ec565b60ff16905085811461094b576040517f240e08a1000000000000000000000000000000000000000000000000000000008152600481018290526024810187905260440161025d565b7fffffffff00000000000000000000000000000000000000000000000000000000841660009081526002602052604081205460ff1690036109dc576040517f4d7a28770000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008516600482015260240161025d565b604080517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608d901b166020808301919091526034820185905260548083018790528351808403909101815260749092018352815191810191909120600081815260039092529190205473ffffffffffffffffffffffffffffffffffffffff1615610a94576040517f5a80160d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008a8a604051610aa6929190611572565b60405180910390209050610ab86111f1565b6000805b8a811015610d4a576000806000610b2a8f8f86818110610ade57610ade61150b565b9050602002810190610af09190611582565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f6392505050565b925092509250600060018883868660405160008152602001604052604051610b6e949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610b90573d6000803e3d6000fd5b505050602060405103519050600260008e7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060020160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205495508560ff16600003610c85576040517fbf18af4300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260240161025d565b610c906001876115e7565b955060008760ff8816601f8110610ca957610ca961150b565b602002015173ffffffffffffffffffffffffffffffffffffffff1614610cfb576040517f8044bb3300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80878760ff16601f8110610d1157610d1161150b565b73ffffffffffffffffffffffffffffffffffffffff909216602092909202015250839250610d42915082905061153a565b915050610abc565b5050505060008b905060008173ffffffffffffffffffffffffffffffffffffffff166345f5dbaa89878f8f6064908092610d8693929190611600565b6040518563ffffffff1660e01b8152600401610da5949392919061162a565b600060405180830381600087803b158015610dbf57600080fd5b505af1925050508015610dd0575060015b610ddc57506000610de0565b5060015b336003600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085858e73ffffffffffffffffffffffffffffffffffffffff167fb4383c25c1299cd52b8b0b890b27f009f787db8b41aeb4547665d08b48c3e31784604051610e7e911515815260200190565b60405180910390a45050600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690555050505050505050505050565b610ec5610ff1565b610ece81611074565b50565b6000806000806064855111610f42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c6964207265706f7274206c656e6774680000000000000000000000604482015260640161025d565b50505050602081015160408201516044830151606490930151919390929190565b60008060008351604114610fd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e6774680000000000000000604482015260640161025d565b50505060208101516040820151606090920151909260009190911a90565b60005473ffffffffffffffffffffffffffffffffffffffff163314611072576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161025d565b565b3373ffffffffffffffffffffffffffffffffffffffff8216036110f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161025d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156111e1579160200282015b828111156111e15781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190611189565b506111ed929150611210565b5090565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156111ed5760008155600101611211565b600060208083528351808285015260005b8181101561125257858101830151858201604001528201611236565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112b557600080fd5b919050565b6000806000606084860312156112cf57600080fd5b6112d884611291565b95602085013595506040909401359392505050565b60008083601f8401126112ff57600080fd5b50813567ffffffffffffffff81111561131757600080fd5b6020830191508360208260051b850101111561133257600080fd5b9250929050565b6000806000806060858703121561134f57600080fd5b84357fffffffff000000000000000000000000000000000000000000000000000000008116811461137f57600080fd5b9350602085013560ff8116811461139557600080fd5b9250604085013567ffffffffffffffff8111156113b157600080fd5b6113bd878288016112ed565b95989497509550505050565b6000806000806000606086880312156113e157600080fd5b6113ea86611291565b9450602086013567ffffffffffffffff8082111561140757600080fd5b818801915088601f83011261141b57600080fd5b81358181111561142a57600080fd5b89602082850101111561143c57600080fd5b60208301965080955050604088013591508082111561145a57600080fd5b50611467888289016112ed565b969995985093965092949392505050565b60006020828403121561148a57600080fd5b61149382611291565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff81811683821602908116908181146114e5576114e561149a565b5092915050565b60ff81811683821601908111156115055761150561149a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361156b5761156b61149a565b5060010190565b8183823760009101908152919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126115b757600080fd5b83018035915067ffffffffffffffff8211156115d257600080fd5b60200191503681900382131561133257600080fd5b60ff82811682821603908111156115055761150561149a565b6000808585111561161057600080fd5b8386111561161d57600080fd5b5050820193919092039150565b84815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101939250505056fea164736f6c6343000813000a",
}

var KeystoneForwarderABI = KeystoneForwarderMetaData.ABI

var KeystoneForwarderBin = KeystoneForwarderMetaData.Bin

func DeployKeystoneForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeystoneForwarder, error) {
	parsed, err := KeystoneForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeystoneForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeystoneForwarder{address: address, abi: *parsed, KeystoneForwarderCaller: KeystoneForwarderCaller{contract: contract}, KeystoneForwarderTransactor: KeystoneForwarderTransactor{contract: contract}, KeystoneForwarderFilterer: KeystoneForwarderFilterer{contract: contract}}, nil
}

type KeystoneForwarder struct {
	address common.Address
	abi     abi.ABI
	KeystoneForwarderCaller
	KeystoneForwarderTransactor
	KeystoneForwarderFilterer
}

type KeystoneForwarderCaller struct {
	contract *bind.BoundContract
}

type KeystoneForwarderTransactor struct {
	contract *bind.BoundContract
}

type KeystoneForwarderFilterer struct {
	contract *bind.BoundContract
}

type KeystoneForwarderSession struct {
	Contract     *KeystoneForwarder
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type KeystoneForwarderCallerSession struct {
	Contract *KeystoneForwarderCaller
	CallOpts bind.CallOpts
}

type KeystoneForwarderTransactorSession struct {
	Contract     *KeystoneForwarderTransactor
	TransactOpts bind.TransactOpts
}

type KeystoneForwarderRaw struct {
	Contract *KeystoneForwarder
}

type KeystoneForwarderCallerRaw struct {
	Contract *KeystoneForwarderCaller
}

type KeystoneForwarderTransactorRaw struct {
	Contract *KeystoneForwarderTransactor
}

func NewKeystoneForwarder(address common.Address, backend bind.ContractBackend) (*KeystoneForwarder, error) {
	abi, err := abi.JSON(strings.NewReader(KeystoneForwarderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindKeystoneForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarder{address: address, abi: abi, KeystoneForwarderCaller: KeystoneForwarderCaller{contract: contract}, KeystoneForwarderTransactor: KeystoneForwarderTransactor{contract: contract}, KeystoneForwarderFilterer: KeystoneForwarderFilterer{contract: contract}}, nil
}

func NewKeystoneForwarderCaller(address common.Address, caller bind.ContractCaller) (*KeystoneForwarderCaller, error) {
	contract, err := bindKeystoneForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarderCaller{contract: contract}, nil
}

func NewKeystoneForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*KeystoneForwarderTransactor, error) {
	contract, err := bindKeystoneForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarderTransactor{contract: contract}, nil
}

func NewKeystoneForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*KeystoneForwarderFilterer, error) {
	contract, err := bindKeystoneForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarderFilterer{contract: contract}, nil
}

func bindKeystoneForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeystoneForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_KeystoneForwarder *KeystoneForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeystoneForwarder.Contract.KeystoneForwarderCaller.contract.Call(opts, result, method, params...)
}

func (_KeystoneForwarder *KeystoneForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.KeystoneForwarderTransactor.contract.Transfer(opts)
}

func (_KeystoneForwarder *KeystoneForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.KeystoneForwarderTransactor.contract.Transact(opts, method, params...)
}

func (_KeystoneForwarder *KeystoneForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeystoneForwarder.Contract.contract.Call(opts, result, method, params...)
}

func (_KeystoneForwarder *KeystoneForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.contract.Transfer(opts)
}

func (_KeystoneForwarder *KeystoneForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.contract.Transact(opts, method, params...)
}

func (_KeystoneForwarder *KeystoneForwarderCaller) GetTransmitter(opts *bind.CallOpts, receiver common.Address, workflowOwner [32]byte, workflowExecutionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _KeystoneForwarder.contract.Call(opts, &out, "getTransmitter", receiver, workflowOwner, workflowExecutionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_KeystoneForwarder *KeystoneForwarderSession) GetTransmitter(receiver common.Address, workflowOwner [32]byte, workflowExecutionId [32]byte) (common.Address, error) {
	return _KeystoneForwarder.Contract.GetTransmitter(&_KeystoneForwarder.CallOpts, receiver, workflowOwner, workflowExecutionId)
}

func (_KeystoneForwarder *KeystoneForwarderCallerSession) GetTransmitter(receiver common.Address, workflowOwner [32]byte, workflowExecutionId [32]byte) (common.Address, error) {
	return _KeystoneForwarder.Contract.GetTransmitter(&_KeystoneForwarder.CallOpts, receiver, workflowOwner, workflowExecutionId)
}

func (_KeystoneForwarder *KeystoneForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeystoneForwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_KeystoneForwarder *KeystoneForwarderSession) Owner() (common.Address, error) {
	return _KeystoneForwarder.Contract.Owner(&_KeystoneForwarder.CallOpts)
}

func (_KeystoneForwarder *KeystoneForwarderCallerSession) Owner() (common.Address, error) {
	return _KeystoneForwarder.Contract.Owner(&_KeystoneForwarder.CallOpts)
}

func (_KeystoneForwarder *KeystoneForwarderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeystoneForwarder.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_KeystoneForwarder *KeystoneForwarderSession) TypeAndVersion() (string, error) {
	return _KeystoneForwarder.Contract.TypeAndVersion(&_KeystoneForwarder.CallOpts)
}

func (_KeystoneForwarder *KeystoneForwarderCallerSession) TypeAndVersion() (string, error) {
	return _KeystoneForwarder.Contract.TypeAndVersion(&_KeystoneForwarder.CallOpts)
}

func (_KeystoneForwarder *KeystoneForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeystoneForwarder.contract.Transact(opts, "acceptOwnership")
}

func (_KeystoneForwarder *KeystoneForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.AcceptOwnership(&_KeystoneForwarder.TransactOpts)
}

func (_KeystoneForwarder *KeystoneForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.AcceptOwnership(&_KeystoneForwarder.TransactOpts)
}

func (_KeystoneForwarder *KeystoneForwarderTransactor) Report(opts *bind.TransactOpts, receiverAddress common.Address, rawReport []byte, signatures [][]byte) (*types.Transaction, error) {
	return _KeystoneForwarder.contract.Transact(opts, "report", receiverAddress, rawReport, signatures)
}

func (_KeystoneForwarder *KeystoneForwarderSession) Report(receiverAddress common.Address, rawReport []byte, signatures [][]byte) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.Report(&_KeystoneForwarder.TransactOpts, receiverAddress, rawReport, signatures)
}

func (_KeystoneForwarder *KeystoneForwarderTransactorSession) Report(receiverAddress common.Address, rawReport []byte, signatures [][]byte) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.Report(&_KeystoneForwarder.TransactOpts, receiverAddress, rawReport, signatures)
}

func (_KeystoneForwarder *KeystoneForwarderTransactor) SetConfig(opts *bind.TransactOpts, donId [4]byte, f uint8, signers []common.Address) (*types.Transaction, error) {
	return _KeystoneForwarder.contract.Transact(opts, "setConfig", donId, f, signers)
}

func (_KeystoneForwarder *KeystoneForwarderSession) SetConfig(donId [4]byte, f uint8, signers []common.Address) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.SetConfig(&_KeystoneForwarder.TransactOpts, donId, f, signers)
}

func (_KeystoneForwarder *KeystoneForwarderTransactorSession) SetConfig(donId [4]byte, f uint8, signers []common.Address) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.SetConfig(&_KeystoneForwarder.TransactOpts, donId, f, signers)
}

func (_KeystoneForwarder *KeystoneForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _KeystoneForwarder.contract.Transact(opts, "transferOwnership", to)
}

func (_KeystoneForwarder *KeystoneForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.TransferOwnership(&_KeystoneForwarder.TransactOpts, to)
}

func (_KeystoneForwarder *KeystoneForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _KeystoneForwarder.Contract.TransferOwnership(&_KeystoneForwarder.TransactOpts, to)
}

type KeystoneForwarderOwnershipTransferRequestedIterator struct {
	Event *KeystoneForwarderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneForwarderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneForwarderOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneForwarderOwnershipTransferRequested)
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

func (it *KeystoneForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *KeystoneForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_KeystoneForwarder *KeystoneForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneForwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarderOwnershipTransferRequestedIterator{contract: _KeystoneForwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_KeystoneForwarder *KeystoneForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *KeystoneForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneForwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneForwarderOwnershipTransferRequested)
				if err := _KeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_KeystoneForwarder *KeystoneForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*KeystoneForwarderOwnershipTransferRequested, error) {
	event := new(KeystoneForwarderOwnershipTransferRequested)
	if err := _KeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type KeystoneForwarderOwnershipTransferredIterator struct {
	Event *KeystoneForwarderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneForwarderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneForwarderOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneForwarderOwnershipTransferred)
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

func (it *KeystoneForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *KeystoneForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_KeystoneForwarder *KeystoneForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneForwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarderOwnershipTransferredIterator{contract: _KeystoneForwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_KeystoneForwarder *KeystoneForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeystoneForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneForwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneForwarderOwnershipTransferred)
				if err := _KeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_KeystoneForwarder *KeystoneForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*KeystoneForwarderOwnershipTransferred, error) {
	event := new(KeystoneForwarderOwnershipTransferred)
	if err := _KeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type KeystoneForwarderReportDeliveredIterator struct {
	Event *KeystoneForwarderReportDelivered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneForwarderReportDeliveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneForwarderReportDelivered)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneForwarderReportDelivered)
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

func (it *KeystoneForwarderReportDeliveredIterator) Error() error {
	return it.fail
}

func (it *KeystoneForwarderReportDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneForwarderReportDelivered struct {
	Receiver            common.Address
	WorkflowOwner       [32]byte
	WorkflowExecutionId [32]byte
	Result              bool
	Raw                 types.Log
}

func (_KeystoneForwarder *KeystoneForwarderFilterer) FilterReportDelivered(opts *bind.FilterOpts, receiver []common.Address, workflowOwner [][32]byte, workflowExecutionId [][32]byte) (*KeystoneForwarderReportDeliveredIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var workflowOwnerRule []interface{}
	for _, workflowOwnerItem := range workflowOwner {
		workflowOwnerRule = append(workflowOwnerRule, workflowOwnerItem)
	}
	var workflowExecutionIdRule []interface{}
	for _, workflowExecutionIdItem := range workflowExecutionId {
		workflowExecutionIdRule = append(workflowExecutionIdRule, workflowExecutionIdItem)
	}

	logs, sub, err := _KeystoneForwarder.contract.FilterLogs(opts, "ReportDelivered", receiverRule, workflowOwnerRule, workflowExecutionIdRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneForwarderReportDeliveredIterator{contract: _KeystoneForwarder.contract, event: "ReportDelivered", logs: logs, sub: sub}, nil
}

func (_KeystoneForwarder *KeystoneForwarderFilterer) WatchReportDelivered(opts *bind.WatchOpts, sink chan<- *KeystoneForwarderReportDelivered, receiver []common.Address, workflowOwner [][32]byte, workflowExecutionId [][32]byte) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var workflowOwnerRule []interface{}
	for _, workflowOwnerItem := range workflowOwner {
		workflowOwnerRule = append(workflowOwnerRule, workflowOwnerItem)
	}
	var workflowExecutionIdRule []interface{}
	for _, workflowExecutionIdItem := range workflowExecutionId {
		workflowExecutionIdRule = append(workflowExecutionIdRule, workflowExecutionIdItem)
	}

	logs, sub, err := _KeystoneForwarder.contract.WatchLogs(opts, "ReportDelivered", receiverRule, workflowOwnerRule, workflowExecutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneForwarderReportDelivered)
				if err := _KeystoneForwarder.contract.UnpackLog(event, "ReportDelivered", log); err != nil {
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

func (_KeystoneForwarder *KeystoneForwarderFilterer) ParseReportDelivered(log types.Log) (*KeystoneForwarderReportDelivered, error) {
	event := new(KeystoneForwarderReportDelivered)
	if err := _KeystoneForwarder.contract.UnpackLog(event, "ReportDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_KeystoneForwarder *KeystoneForwarder) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _KeystoneForwarder.abi.Events["OwnershipTransferRequested"].ID:
		return _KeystoneForwarder.ParseOwnershipTransferRequested(log)
	case _KeystoneForwarder.abi.Events["OwnershipTransferred"].ID:
		return _KeystoneForwarder.ParseOwnershipTransferred(log)
	case _KeystoneForwarder.abi.Events["ReportDelivered"].ID:
		return _KeystoneForwarder.ParseReportDelivered(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (KeystoneForwarderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (KeystoneForwarderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (KeystoneForwarderReportDelivered) Topic() common.Hash {
	return common.HexToHash("0xb4383c25c1299cd52b8b0b890b27f009f787db8b41aeb4547665d08b48c3e317")
}

func (_KeystoneForwarder *KeystoneForwarder) Address() common.Address {
	return _KeystoneForwarder.address
}

type KeystoneForwarderInterface interface {
	GetTransmitter(opts *bind.CallOpts, receiver common.Address, workflowOwner [32]byte, workflowExecutionId [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, receiverAddress common.Address, rawReport []byte, signatures [][]byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, donId [4]byte, f uint8, signers []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneForwarderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *KeystoneForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*KeystoneForwarderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneForwarderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeystoneForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*KeystoneForwarderOwnershipTransferred, error)

	FilterReportDelivered(opts *bind.FilterOpts, receiver []common.Address, workflowOwner [][32]byte, workflowExecutionId [][32]byte) (*KeystoneForwarderReportDeliveredIterator, error)

	WatchReportDelivered(opts *bind.WatchOpts, sink chan<- *KeystoneForwarderReportDelivered, receiver []common.Address, workflowOwner [][32]byte, workflowExecutionId [][32]byte) (event.Subscription, error)

	ParseReportDelivered(log types.Log) (*KeystoneForwarderReportDelivered, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
