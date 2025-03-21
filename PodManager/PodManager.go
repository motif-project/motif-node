// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PodManager

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

// IBitcoinPodManagerBitcoinDepositRequest is an auto generated low-level Go binding around an user-defined struct.
type IBitcoinPodManagerBitcoinDepositRequest struct {
	TransactionId [32]byte
	Amount        *big.Int
	IsPending     bool
}

// PodManagerMetaData contains all meta data concerning the PodManager contract.
var PodManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelWithdrawalRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmBitcoinDeposit\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPod\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"btcAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"script\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatePod\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAppRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBitcoinDepositRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBitcoinPodManager.BitcoinDepositRequest\",\"components\":[{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isPending\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBitcoinWithdrawalAddress\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMotifServiceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMotifStakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPodApp\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPods\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalTVL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserPod\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPendingBitcoinDepositRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"appRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"motifStakeRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"motifServiceManager_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockPod\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignedBitcoinWithdrawTransactionPod\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signedBitcoinWithdrawTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegatePod\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockPod\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBitcoinDepositRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyPresignedBitcoinDepositRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBitcoinAsTokens\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBitcoinCompleteTxRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"preSignedWithdrawTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBitcoinPSBTRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawPresignedBitcoinRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BTCAddressVerified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"btcAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinBurned\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinDepositConfirmed\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinMinted\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinWithdrawalCompleteTxRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"preSignedBitcoinTx\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinWithdrawalPSBTRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinWithdrawnFromPod\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodCreated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodDelegated\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodUndelegated\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalTVLUpdated\",\"inputs\":[{\"name\":\"newTVL\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifyBitcoinDepositRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bitcoinDepositRequest\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBitcoinPodManager.BitcoinDepositRequest\",\"components\":[{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isPending\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifyPresignedBitcoinDepositRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bitcoinDepositRequest\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBitcoinPodManager.BitcoinDepositRequest\",\"components\":[{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isPending\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"transaction\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawPresignedBitcoinRequest\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRequestCancelled\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidBTCAddressInitialByte\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBTCAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOperatorBTCKey\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidScriptLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoWithdrawalRequestToCancel\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WithdrawalTransactionAlreadySubmitted\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WithdrawalTransactionNotSubmitted\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60806040526004361061004d575f3560e01c80633659cfe6146100645780634f1ef286146100835780635c60da1b146100965780638f283970146100c6578063f851a440146100e55761005c565b3661005c5761005a6100f9565b005b61005a6100f9565b34801561006f575f80fd5b5061005a61007e36600461068c565b610113565b61005a6100913660046106a5565b61014e565b3480156100a1575f80fd5b506100aa6101b4565b6040516001600160a01b03909116815260200160405180910390f35b3480156100d1575f80fd5b5061005a6100e036600461068c565b6101e4565b3480156100f0575f80fd5b506100aa610204565b610101610224565b61011161010c6102b9565b6102c2565b565b61011b6102e0565b6001600160a01b03163303610146576101438160405180602001604052805f8152505f610312565b50565b6101436100f9565b6101566102e0565b6001600160a01b031633036101ac576101a78383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525060019250610312915050565b505050565b6101a76100f9565b5f6101bd6102e0565b6001600160a01b031633036101d9576101d46102b9565b905090565b6101e16100f9565b90565b6101ec6102e0565b6001600160a01b03163303610146576101438161033c565b5f61020d6102e0565b6001600160a01b031633036101d9576101d46102e0565b61022c6102e0565b6001600160a01b031633036101115760405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4015b60405180910390fd5b5f6101d4610390565b365f80375f80365f845af43d5f803e8080156102dc573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b546001600160a01b0316919050565b61031b836103b7565b5f825111806103275750805b156101a75761033683836103f6565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6103656102e0565b604080516001600160a01b03928316815291841660208301520160405180910390a161014381610422565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610303565b6103c0816104cb565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b606061041b838360405180606001604052806027815260200161076d6027913961055f565b9392505050565b6001600160a01b0381166104875760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084016102b0565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6105385760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016102b0565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6104aa565b60606001600160a01b0384163b6105c75760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016102b0565b5f80856001600160a01b0316856040516105e19190610721565b5f60405180830381855af49150503d805f8114610619576040519150601f19603f3d011682016040523d82523d5f602084013e61061e565b606091505b509150915061062e828286610638565b9695505050505050565b6060831561064757508161041b565b8251156106575782518084602001fd5b8160405162461bcd60e51b81526004016102b09190610737565b80356001600160a01b0381168114610687575f80fd5b919050565b5f6020828403121561069c575f80fd5b61041b82610671565b5f805f604084860312156106b7575f80fd5b6106c084610671565b9250602084013567ffffffffffffffff808211156106dc575f80fd5b818601915086601f8301126106ef575f80fd5b8135818111156106fd575f80fd5b87602082850101111561070e575f80fd5b6020830194508093505050509250925092565b5f82518060208501845e5f920191825250919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f8301168401019150509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122076eafcc7f4d540d0096a750183a7bb5998f5856a93b8e6f7b99d17ecfc62d6da64736f6c63430008190033",
}

// PodManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PodManagerMetaData.ABI instead.
var PodManagerABI = PodManagerMetaData.ABI

// PodManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PodManagerMetaData.Bin instead.
var PodManagerBin = PodManagerMetaData.Bin

// DeployPodManager deploys a new Ethereum contract, binding an instance of PodManager to it.
func DeployPodManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PodManager, error) {
	parsed, err := PodManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PodManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PodManager{PodManagerCaller: PodManagerCaller{contract: contract}, PodManagerTransactor: PodManagerTransactor{contract: contract}, PodManagerFilterer: PodManagerFilterer{contract: contract}}, nil
}

// PodManager is an auto generated Go binding around an Ethereum contract.
type PodManager struct {
	PodManagerCaller     // Read-only binding to the contract
	PodManagerTransactor // Write-only binding to the contract
	PodManagerFilterer   // Log filterer for contract events
}

// PodManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PodManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PodManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PodManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PodManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PodManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PodManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PodManagerSession struct {
	Contract     *PodManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PodManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PodManagerCallerSession struct {
	Contract *PodManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PodManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PodManagerTransactorSession struct {
	Contract     *PodManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PodManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PodManagerRaw struct {
	Contract *PodManager // Generic contract binding to access the raw methods on
}

// PodManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PodManagerCallerRaw struct {
	Contract *PodManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PodManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PodManagerTransactorRaw struct {
	Contract *PodManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPodManager creates a new instance of PodManager, bound to a specific deployed contract.
func NewPodManager(address common.Address, backend bind.ContractBackend) (*PodManager, error) {
	contract, err := bindPodManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PodManager{PodManagerCaller: PodManagerCaller{contract: contract}, PodManagerTransactor: PodManagerTransactor{contract: contract}, PodManagerFilterer: PodManagerFilterer{contract: contract}}, nil
}

// NewPodManagerCaller creates a new read-only instance of PodManager, bound to a specific deployed contract.
func NewPodManagerCaller(address common.Address, caller bind.ContractCaller) (*PodManagerCaller, error) {
	contract, err := bindPodManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PodManagerCaller{contract: contract}, nil
}

// NewPodManagerTransactor creates a new write-only instance of PodManager, bound to a specific deployed contract.
func NewPodManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PodManagerTransactor, error) {
	contract, err := bindPodManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PodManagerTransactor{contract: contract}, nil
}

// NewPodManagerFilterer creates a new log filterer instance of PodManager, bound to a specific deployed contract.
func NewPodManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PodManagerFilterer, error) {
	contract, err := bindPodManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PodManagerFilterer{contract: contract}, nil
}

// bindPodManager binds a generic wrapper to an already deployed contract.
func bindPodManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PodManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PodManager *PodManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PodManager.Contract.PodManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PodManager *PodManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.Contract.PodManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PodManager *PodManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PodManager.Contract.PodManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PodManager *PodManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PodManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PodManager *PodManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PodManager *PodManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PodManager.Contract.contract.Transact(opts, method, params...)
}

// GetAppRegistry is a free data retrieval call binding the contract method 0xb4ff8ff1.
//
// Solidity: function getAppRegistry() view returns(address)
func (_PodManager *PodManagerCaller) GetAppRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getAppRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAppRegistry is a free data retrieval call binding the contract method 0xb4ff8ff1.
//
// Solidity: function getAppRegistry() view returns(address)
func (_PodManager *PodManagerSession) GetAppRegistry() (common.Address, error) {
	return _PodManager.Contract.GetAppRegistry(&_PodManager.CallOpts)
}

// GetAppRegistry is a free data retrieval call binding the contract method 0xb4ff8ff1.
//
// Solidity: function getAppRegistry() view returns(address)
func (_PodManager *PodManagerCallerSession) GetAppRegistry() (common.Address, error) {
	return _PodManager.Contract.GetAppRegistry(&_PodManager.CallOpts)
}

// GetBitcoinDepositRequest is a free data retrieval call binding the contract method 0x85974473.
//
// Solidity: function getBitcoinDepositRequest(address pod) view returns((bytes32,uint256,bool))
func (_PodManager *PodManagerCaller) GetBitcoinDepositRequest(opts *bind.CallOpts, pod common.Address) (IBitcoinPodManagerBitcoinDepositRequest, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getBitcoinDepositRequest", pod)

	if err != nil {
		return *new(IBitcoinPodManagerBitcoinDepositRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IBitcoinPodManagerBitcoinDepositRequest)).(*IBitcoinPodManagerBitcoinDepositRequest)

	return out0, err

}

// GetBitcoinDepositRequest is a free data retrieval call binding the contract method 0x85974473.
//
// Solidity: function getBitcoinDepositRequest(address pod) view returns((bytes32,uint256,bool))
func (_PodManager *PodManagerSession) GetBitcoinDepositRequest(pod common.Address) (IBitcoinPodManagerBitcoinDepositRequest, error) {
	return _PodManager.Contract.GetBitcoinDepositRequest(&_PodManager.CallOpts, pod)
}

// GetBitcoinDepositRequest is a free data retrieval call binding the contract method 0x85974473.
//
// Solidity: function getBitcoinDepositRequest(address pod) view returns((bytes32,uint256,bool))
func (_PodManager *PodManagerCallerSession) GetBitcoinDepositRequest(pod common.Address) (IBitcoinPodManagerBitcoinDepositRequest, error) {
	return _PodManager.Contract.GetBitcoinDepositRequest(&_PodManager.CallOpts, pod)
}

// GetBitcoinWithdrawalAddress is a free data retrieval call binding the contract method 0xc07eea48.
//
// Solidity: function getBitcoinWithdrawalAddress(address pod) view returns(string)
func (_PodManager *PodManagerCaller) GetBitcoinWithdrawalAddress(opts *bind.CallOpts, pod common.Address) (string, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getBitcoinWithdrawalAddress", pod)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBitcoinWithdrawalAddress is a free data retrieval call binding the contract method 0xc07eea48.
//
// Solidity: function getBitcoinWithdrawalAddress(address pod) view returns(string)
func (_PodManager *PodManagerSession) GetBitcoinWithdrawalAddress(pod common.Address) (string, error) {
	return _PodManager.Contract.GetBitcoinWithdrawalAddress(&_PodManager.CallOpts, pod)
}

// GetBitcoinWithdrawalAddress is a free data retrieval call binding the contract method 0xc07eea48.
//
// Solidity: function getBitcoinWithdrawalAddress(address pod) view returns(string)
func (_PodManager *PodManagerCallerSession) GetBitcoinWithdrawalAddress(pod common.Address) (string, error) {
	return _PodManager.Contract.GetBitcoinWithdrawalAddress(&_PodManager.CallOpts, pod)
}

// GetMotifServiceManager is a free data retrieval call binding the contract method 0xf829c617.
//
// Solidity: function getMotifServiceManager() view returns(address)
func (_PodManager *PodManagerCaller) GetMotifServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getMotifServiceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMotifServiceManager is a free data retrieval call binding the contract method 0xf829c617.
//
// Solidity: function getMotifServiceManager() view returns(address)
func (_PodManager *PodManagerSession) GetMotifServiceManager() (common.Address, error) {
	return _PodManager.Contract.GetMotifServiceManager(&_PodManager.CallOpts)
}

// GetMotifServiceManager is a free data retrieval call binding the contract method 0xf829c617.
//
// Solidity: function getMotifServiceManager() view returns(address)
func (_PodManager *PodManagerCallerSession) GetMotifServiceManager() (common.Address, error) {
	return _PodManager.Contract.GetMotifServiceManager(&_PodManager.CallOpts)
}

// GetMotifStakeRegistry is a free data retrieval call binding the contract method 0xaaefe4af.
//
// Solidity: function getMotifStakeRegistry() view returns(address)
func (_PodManager *PodManagerCaller) GetMotifStakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getMotifStakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMotifStakeRegistry is a free data retrieval call binding the contract method 0xaaefe4af.
//
// Solidity: function getMotifStakeRegistry() view returns(address)
func (_PodManager *PodManagerSession) GetMotifStakeRegistry() (common.Address, error) {
	return _PodManager.Contract.GetMotifStakeRegistry(&_PodManager.CallOpts)
}

// GetMotifStakeRegistry is a free data retrieval call binding the contract method 0xaaefe4af.
//
// Solidity: function getMotifStakeRegistry() view returns(address)
func (_PodManager *PodManagerCallerSession) GetMotifStakeRegistry() (common.Address, error) {
	return _PodManager.Contract.GetMotifStakeRegistry(&_PodManager.CallOpts)
}

// GetPodApp is a free data retrieval call binding the contract method 0x0f2daf30.
//
// Solidity: function getPodApp(address pod) view returns(address)
func (_PodManager *PodManagerCaller) GetPodApp(opts *bind.CallOpts, pod common.Address) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getPodApp", pod)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPodApp is a free data retrieval call binding the contract method 0x0f2daf30.
//
// Solidity: function getPodApp(address pod) view returns(address)
func (_PodManager *PodManagerSession) GetPodApp(pod common.Address) (common.Address, error) {
	return _PodManager.Contract.GetPodApp(&_PodManager.CallOpts, pod)
}

// GetPodApp is a free data retrieval call binding the contract method 0x0f2daf30.
//
// Solidity: function getPodApp(address pod) view returns(address)
func (_PodManager *PodManagerCallerSession) GetPodApp(pod common.Address) (common.Address, error) {
	return _PodManager.Contract.GetPodApp(&_PodManager.CallOpts, pod)
}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_PodManager *PodManagerCaller) GetTotalPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getTotalPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_PodManager *PodManagerSession) GetTotalPods() (*big.Int, error) {
	return _PodManager.Contract.GetTotalPods(&_PodManager.CallOpts)
}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_PodManager *PodManagerCallerSession) GetTotalPods() (*big.Int, error) {
	return _PodManager.Contract.GetTotalPods(&_PodManager.CallOpts)
}

// GetTotalTVL is a free data retrieval call binding the contract method 0x0c021ae5.
//
// Solidity: function getTotalTVL() view returns(uint256)
func (_PodManager *PodManagerCaller) GetTotalTVL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getTotalTVL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalTVL is a free data retrieval call binding the contract method 0x0c021ae5.
//
// Solidity: function getTotalTVL() view returns(uint256)
func (_PodManager *PodManagerSession) GetTotalTVL() (*big.Int, error) {
	return _PodManager.Contract.GetTotalTVL(&_PodManager.CallOpts)
}

// GetTotalTVL is a free data retrieval call binding the contract method 0x0c021ae5.
//
// Solidity: function getTotalTVL() view returns(uint256)
func (_PodManager *PodManagerCallerSession) GetTotalTVL() (*big.Int, error) {
	return _PodManager.Contract.GetTotalTVL(&_PodManager.CallOpts)
}

// GetUserPod is a free data retrieval call binding the contract method 0xda600a8a.
//
// Solidity: function getUserPod(address user) view returns(address)
func (_PodManager *PodManagerCaller) GetUserPod(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "getUserPod", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserPod is a free data retrieval call binding the contract method 0xda600a8a.
//
// Solidity: function getUserPod(address user) view returns(address)
func (_PodManager *PodManagerSession) GetUserPod(user common.Address) (common.Address, error) {
	return _PodManager.Contract.GetUserPod(&_PodManager.CallOpts, user)
}

// GetUserPod is a free data retrieval call binding the contract method 0xda600a8a.
//
// Solidity: function getUserPod(address user) view returns(address)
func (_PodManager *PodManagerCallerSession) GetUserPod(user common.Address) (common.Address, error) {
	return _PodManager.Contract.GetUserPod(&_PodManager.CallOpts, user)
}

// HasPendingBitcoinDepositRequest is a free data retrieval call binding the contract method 0xac7559d5.
//
// Solidity: function hasPendingBitcoinDepositRequest(address pod) view returns(bool)
func (_PodManager *PodManagerCaller) HasPendingBitcoinDepositRequest(opts *bind.CallOpts, pod common.Address) (bool, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "hasPendingBitcoinDepositRequest", pod)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPendingBitcoinDepositRequest is a free data retrieval call binding the contract method 0xac7559d5.
//
// Solidity: function hasPendingBitcoinDepositRequest(address pod) view returns(bool)
func (_PodManager *PodManagerSession) HasPendingBitcoinDepositRequest(pod common.Address) (bool, error) {
	return _PodManager.Contract.HasPendingBitcoinDepositRequest(&_PodManager.CallOpts, pod)
}

// HasPendingBitcoinDepositRequest is a free data retrieval call binding the contract method 0xac7559d5.
//
// Solidity: function hasPendingBitcoinDepositRequest(address pod) view returns(bool)
func (_PodManager *PodManagerCallerSession) HasPendingBitcoinDepositRequest(pod common.Address) (bool, error) {
	return _PodManager.Contract.HasPendingBitcoinDepositRequest(&_PodManager.CallOpts, pod)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PodManager *PodManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PodManager *PodManagerSession) Owner() (common.Address, error) {
	return _PodManager.Contract.Owner(&_PodManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PodManager *PodManagerCallerSession) Owner() (common.Address, error) {
	return _PodManager.Contract.Owner(&_PodManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PodManager *PodManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PodManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PodManager *PodManagerSession) Paused() (bool, error) {
	return _PodManager.Contract.Paused(&_PodManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PodManager *PodManagerCallerSession) Paused() (bool, error) {
	return _PodManager.Contract.Paused(&_PodManager.CallOpts)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0x4bda23ec.
//
// Solidity: function cancelWithdrawalRequest(address pod) returns()
func (_PodManager *PodManagerTransactor) CancelWithdrawalRequest(opts *bind.TransactOpts, pod common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "cancelWithdrawalRequest", pod)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0x4bda23ec.
//
// Solidity: function cancelWithdrawalRequest(address pod) returns()
func (_PodManager *PodManagerSession) CancelWithdrawalRequest(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.CancelWithdrawalRequest(&_PodManager.TransactOpts, pod)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0x4bda23ec.
//
// Solidity: function cancelWithdrawalRequest(address pod) returns()
func (_PodManager *PodManagerTransactorSession) CancelWithdrawalRequest(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.CancelWithdrawalRequest(&_PodManager.TransactOpts, pod)
}

// ConfirmBitcoinDeposit is a paid mutator transaction binding the contract method 0x397ee69b.
//
// Solidity: function confirmBitcoinDeposit(address pod, bytes32 transactionId, uint256 amount) returns()
func (_PodManager *PodManagerTransactor) ConfirmBitcoinDeposit(opts *bind.TransactOpts, pod common.Address, transactionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "confirmBitcoinDeposit", pod, transactionId, amount)
}

// ConfirmBitcoinDeposit is a paid mutator transaction binding the contract method 0x397ee69b.
//
// Solidity: function confirmBitcoinDeposit(address pod, bytes32 transactionId, uint256 amount) returns()
func (_PodManager *PodManagerSession) ConfirmBitcoinDeposit(pod common.Address, transactionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.ConfirmBitcoinDeposit(&_PodManager.TransactOpts, pod, transactionId, amount)
}

// ConfirmBitcoinDeposit is a paid mutator transaction binding the contract method 0x397ee69b.
//
// Solidity: function confirmBitcoinDeposit(address pod, bytes32 transactionId, uint256 amount) returns()
func (_PodManager *PodManagerTransactorSession) ConfirmBitcoinDeposit(pod common.Address, transactionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.ConfirmBitcoinDeposit(&_PodManager.TransactOpts, pod, transactionId, amount)
}

// CreatePod is a paid mutator transaction binding the contract method 0x235a1c39.
//
// Solidity: function createPod(address operator, string btcAddress, bytes script) returns(address)
func (_PodManager *PodManagerTransactor) CreatePod(opts *bind.TransactOpts, operator common.Address, btcAddress string, script []byte) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "createPod", operator, btcAddress, script)
}

// CreatePod is a paid mutator transaction binding the contract method 0x235a1c39.
//
// Solidity: function createPod(address operator, string btcAddress, bytes script) returns(address)
func (_PodManager *PodManagerSession) CreatePod(operator common.Address, btcAddress string, script []byte) (*types.Transaction, error) {
	return _PodManager.Contract.CreatePod(&_PodManager.TransactOpts, operator, btcAddress, script)
}

// CreatePod is a paid mutator transaction binding the contract method 0x235a1c39.
//
// Solidity: function createPod(address operator, string btcAddress, bytes script) returns(address)
func (_PodManager *PodManagerTransactorSession) CreatePod(operator common.Address, btcAddress string, script []byte) (*types.Transaction, error) {
	return _PodManager.Contract.CreatePod(&_PodManager.TransactOpts, operator, btcAddress, script)
}

// DelegatePod is a paid mutator transaction binding the contract method 0xf99b67c9.
//
// Solidity: function delegatePod(address pod, address appContract) returns()
func (_PodManager *PodManagerTransactor) DelegatePod(opts *bind.TransactOpts, pod common.Address, appContract common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "delegatePod", pod, appContract)
}

// DelegatePod is a paid mutator transaction binding the contract method 0xf99b67c9.
//
// Solidity: function delegatePod(address pod, address appContract) returns()
func (_PodManager *PodManagerSession) DelegatePod(pod common.Address, appContract common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.DelegatePod(&_PodManager.TransactOpts, pod, appContract)
}

// DelegatePod is a paid mutator transaction binding the contract method 0xf99b67c9.
//
// Solidity: function delegatePod(address pod, address appContract) returns()
func (_PodManager *PodManagerTransactorSession) DelegatePod(pod common.Address, appContract common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.DelegatePod(&_PodManager.TransactOpts, pod, appContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address appRegistry_, address motifStakeRegistry_, address motifServiceManager_) returns()
func (_PodManager *PodManagerTransactor) Initialize(opts *bind.TransactOpts, appRegistry_ common.Address, motifStakeRegistry_ common.Address, motifServiceManager_ common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "initialize", appRegistry_, motifStakeRegistry_, motifServiceManager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address appRegistry_, address motifStakeRegistry_, address motifServiceManager_) returns()
func (_PodManager *PodManagerSession) Initialize(appRegistry_ common.Address, motifStakeRegistry_ common.Address, motifServiceManager_ common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.Initialize(&_PodManager.TransactOpts, appRegistry_, motifStakeRegistry_, motifServiceManager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address appRegistry_, address motifStakeRegistry_, address motifServiceManager_) returns()
func (_PodManager *PodManagerTransactorSession) Initialize(appRegistry_ common.Address, motifStakeRegistry_ common.Address, motifServiceManager_ common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.Initialize(&_PodManager.TransactOpts, appRegistry_, motifStakeRegistry_, motifServiceManager_)
}

// LockPod is a paid mutator transaction binding the contract method 0x221d3300.
//
// Solidity: function lockPod(address pod) returns()
func (_PodManager *PodManagerTransactor) LockPod(opts *bind.TransactOpts, pod common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "lockPod", pod)
}

// LockPod is a paid mutator transaction binding the contract method 0x221d3300.
//
// Solidity: function lockPod(address pod) returns()
func (_PodManager *PodManagerSession) LockPod(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.LockPod(&_PodManager.TransactOpts, pod)
}

// LockPod is a paid mutator transaction binding the contract method 0x221d3300.
//
// Solidity: function lockPod(address pod) returns()
func (_PodManager *PodManagerTransactorSession) LockPod(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.LockPod(&_PodManager.TransactOpts, pod)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PodManager *PodManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PodManager *PodManagerSession) Pause() (*types.Transaction, error) {
	return _PodManager.Contract.Pause(&_PodManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PodManager *PodManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _PodManager.Contract.Pause(&_PodManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PodManager *PodManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PodManager *PodManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PodManager.Contract.RenounceOwnership(&_PodManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PodManager *PodManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PodManager.Contract.RenounceOwnership(&_PodManager.TransactOpts)
}

// SetSignedBitcoinWithdrawTransactionPod is a paid mutator transaction binding the contract method 0xd5062c77.
//
// Solidity: function setSignedBitcoinWithdrawTransactionPod(address pod, bytes signedBitcoinWithdrawTransaction) returns()
func (_PodManager *PodManagerTransactor) SetSignedBitcoinWithdrawTransactionPod(opts *bind.TransactOpts, pod common.Address, signedBitcoinWithdrawTransaction []byte) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "setSignedBitcoinWithdrawTransactionPod", pod, signedBitcoinWithdrawTransaction)
}

// SetSignedBitcoinWithdrawTransactionPod is a paid mutator transaction binding the contract method 0xd5062c77.
//
// Solidity: function setSignedBitcoinWithdrawTransactionPod(address pod, bytes signedBitcoinWithdrawTransaction) returns()
func (_PodManager *PodManagerSession) SetSignedBitcoinWithdrawTransactionPod(pod common.Address, signedBitcoinWithdrawTransaction []byte) (*types.Transaction, error) {
	return _PodManager.Contract.SetSignedBitcoinWithdrawTransactionPod(&_PodManager.TransactOpts, pod, signedBitcoinWithdrawTransaction)
}

// SetSignedBitcoinWithdrawTransactionPod is a paid mutator transaction binding the contract method 0xd5062c77.
//
// Solidity: function setSignedBitcoinWithdrawTransactionPod(address pod, bytes signedBitcoinWithdrawTransaction) returns()
func (_PodManager *PodManagerTransactorSession) SetSignedBitcoinWithdrawTransactionPod(pod common.Address, signedBitcoinWithdrawTransaction []byte) (*types.Transaction, error) {
	return _PodManager.Contract.SetSignedBitcoinWithdrawTransactionPod(&_PodManager.TransactOpts, pod, signedBitcoinWithdrawTransaction)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PodManager *PodManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PodManager *PodManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.TransferOwnership(&_PodManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PodManager *PodManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.TransferOwnership(&_PodManager.TransactOpts, newOwner)
}

// UndelegatePod is a paid mutator transaction binding the contract method 0xed1dfe0b.
//
// Solidity: function undelegatePod(address pod) returns()
func (_PodManager *PodManagerTransactor) UndelegatePod(opts *bind.TransactOpts, pod common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "undelegatePod", pod)
}

// UndelegatePod is a paid mutator transaction binding the contract method 0xed1dfe0b.
//
// Solidity: function undelegatePod(address pod) returns()
func (_PodManager *PodManagerSession) UndelegatePod(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.UndelegatePod(&_PodManager.TransactOpts, pod)
}

// UndelegatePod is a paid mutator transaction binding the contract method 0xed1dfe0b.
//
// Solidity: function undelegatePod(address pod) returns()
func (_PodManager *PodManagerTransactorSession) UndelegatePod(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.UndelegatePod(&_PodManager.TransactOpts, pod)
}

// UnlockPod is a paid mutator transaction binding the contract method 0x7de2970d.
//
// Solidity: function unlockPod(address pod) returns()
func (_PodManager *PodManagerTransactor) UnlockPod(opts *bind.TransactOpts, pod common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "unlockPod", pod)
}

// UnlockPod is a paid mutator transaction binding the contract method 0x7de2970d.
//
// Solidity: function unlockPod(address pod) returns()
func (_PodManager *PodManagerSession) UnlockPod(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.UnlockPod(&_PodManager.TransactOpts, pod)
}

// UnlockPod is a paid mutator transaction binding the contract method 0x7de2970d.
//
// Solidity: function unlockPod(address pod) returns()
func (_PodManager *PodManagerTransactorSession) UnlockPod(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.UnlockPod(&_PodManager.TransactOpts, pod)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PodManager *PodManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PodManager *PodManagerSession) Unpause() (*types.Transaction, error) {
	return _PodManager.Contract.Unpause(&_PodManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PodManager *PodManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _PodManager.Contract.Unpause(&_PodManager.TransactOpts)
}

// VerifyBitcoinDepositRequest is a paid mutator transaction binding the contract method 0x5b674bc9.
//
// Solidity: function verifyBitcoinDepositRequest(address pod, bytes32 transactionId, uint256 amount) returns()
func (_PodManager *PodManagerTransactor) VerifyBitcoinDepositRequest(opts *bind.TransactOpts, pod common.Address, transactionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "verifyBitcoinDepositRequest", pod, transactionId, amount)
}

// VerifyBitcoinDepositRequest is a paid mutator transaction binding the contract method 0x5b674bc9.
//
// Solidity: function verifyBitcoinDepositRequest(address pod, bytes32 transactionId, uint256 amount) returns()
func (_PodManager *PodManagerSession) VerifyBitcoinDepositRequest(pod common.Address, transactionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.VerifyBitcoinDepositRequest(&_PodManager.TransactOpts, pod, transactionId, amount)
}

// VerifyBitcoinDepositRequest is a paid mutator transaction binding the contract method 0x5b674bc9.
//
// Solidity: function verifyBitcoinDepositRequest(address pod, bytes32 transactionId, uint256 amount) returns()
func (_PodManager *PodManagerTransactorSession) VerifyBitcoinDepositRequest(pod common.Address, transactionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.VerifyBitcoinDepositRequest(&_PodManager.TransactOpts, pod, transactionId, amount)
}

// VerifyPresignedBitcoinDepositRequest is a paid mutator transaction binding the contract method 0xba9f5e4e.
//
// Solidity: function verifyPresignedBitcoinDepositRequest(address pod, bytes32 transactionId, bytes transaction, uint256 amount) returns()
func (_PodManager *PodManagerTransactor) VerifyPresignedBitcoinDepositRequest(opts *bind.TransactOpts, pod common.Address, transactionId [32]byte, transaction []byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "verifyPresignedBitcoinDepositRequest", pod, transactionId, transaction, amount)
}

// VerifyPresignedBitcoinDepositRequest is a paid mutator transaction binding the contract method 0xba9f5e4e.
//
// Solidity: function verifyPresignedBitcoinDepositRequest(address pod, bytes32 transactionId, bytes transaction, uint256 amount) returns()
func (_PodManager *PodManagerSession) VerifyPresignedBitcoinDepositRequest(pod common.Address, transactionId [32]byte, transaction []byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.VerifyPresignedBitcoinDepositRequest(&_PodManager.TransactOpts, pod, transactionId, transaction, amount)
}

// VerifyPresignedBitcoinDepositRequest is a paid mutator transaction binding the contract method 0xba9f5e4e.
//
// Solidity: function verifyPresignedBitcoinDepositRequest(address pod, bytes32 transactionId, bytes transaction, uint256 amount) returns()
func (_PodManager *PodManagerTransactorSession) VerifyPresignedBitcoinDepositRequest(pod common.Address, transactionId [32]byte, transaction []byte, amount *big.Int) (*types.Transaction, error) {
	return _PodManager.Contract.VerifyPresignedBitcoinDepositRequest(&_PodManager.TransactOpts, pod, transactionId, transaction, amount)
}

// WithdrawBitcoinAsTokens is a paid mutator transaction binding the contract method 0xd475ab03.
//
// Solidity: function withdrawBitcoinAsTokens(address pod) returns()
func (_PodManager *PodManagerTransactor) WithdrawBitcoinAsTokens(opts *bind.TransactOpts, pod common.Address) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "withdrawBitcoinAsTokens", pod)
}

// WithdrawBitcoinAsTokens is a paid mutator transaction binding the contract method 0xd475ab03.
//
// Solidity: function withdrawBitcoinAsTokens(address pod) returns()
func (_PodManager *PodManagerSession) WithdrawBitcoinAsTokens(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawBitcoinAsTokens(&_PodManager.TransactOpts, pod)
}

// WithdrawBitcoinAsTokens is a paid mutator transaction binding the contract method 0xd475ab03.
//
// Solidity: function withdrawBitcoinAsTokens(address pod) returns()
func (_PodManager *PodManagerTransactorSession) WithdrawBitcoinAsTokens(pod common.Address) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawBitcoinAsTokens(&_PodManager.TransactOpts, pod)
}

// WithdrawBitcoinCompleteTxRequest is a paid mutator transaction binding the contract method 0x2b6bbf75.
//
// Solidity: function withdrawBitcoinCompleteTxRequest(address pod, bytes preSignedWithdrawTransaction, string withdrawAddress) returns()
func (_PodManager *PodManagerTransactor) WithdrawBitcoinCompleteTxRequest(opts *bind.TransactOpts, pod common.Address, preSignedWithdrawTransaction []byte, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "withdrawBitcoinCompleteTxRequest", pod, preSignedWithdrawTransaction, withdrawAddress)
}

// WithdrawBitcoinCompleteTxRequest is a paid mutator transaction binding the contract method 0x2b6bbf75.
//
// Solidity: function withdrawBitcoinCompleteTxRequest(address pod, bytes preSignedWithdrawTransaction, string withdrawAddress) returns()
func (_PodManager *PodManagerSession) WithdrawBitcoinCompleteTxRequest(pod common.Address, preSignedWithdrawTransaction []byte, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawBitcoinCompleteTxRequest(&_PodManager.TransactOpts, pod, preSignedWithdrawTransaction, withdrawAddress)
}

// WithdrawBitcoinCompleteTxRequest is a paid mutator transaction binding the contract method 0x2b6bbf75.
//
// Solidity: function withdrawBitcoinCompleteTxRequest(address pod, bytes preSignedWithdrawTransaction, string withdrawAddress) returns()
func (_PodManager *PodManagerTransactorSession) WithdrawBitcoinCompleteTxRequest(pod common.Address, preSignedWithdrawTransaction []byte, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawBitcoinCompleteTxRequest(&_PodManager.TransactOpts, pod, preSignedWithdrawTransaction, withdrawAddress)
}

// WithdrawBitcoinPSBTRequest is a paid mutator transaction binding the contract method 0x7104c3df.
//
// Solidity: function withdrawBitcoinPSBTRequest(address pod, string withdrawAddress) returns()
func (_PodManager *PodManagerTransactor) WithdrawBitcoinPSBTRequest(opts *bind.TransactOpts, pod common.Address, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "withdrawBitcoinPSBTRequest", pod, withdrawAddress)
}

// WithdrawBitcoinPSBTRequest is a paid mutator transaction binding the contract method 0x7104c3df.
//
// Solidity: function withdrawBitcoinPSBTRequest(address pod, string withdrawAddress) returns()
func (_PodManager *PodManagerSession) WithdrawBitcoinPSBTRequest(pod common.Address, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawBitcoinPSBTRequest(&_PodManager.TransactOpts, pod, withdrawAddress)
}

// WithdrawBitcoinPSBTRequest is a paid mutator transaction binding the contract method 0x7104c3df.
//
// Solidity: function withdrawBitcoinPSBTRequest(address pod, string withdrawAddress) returns()
func (_PodManager *PodManagerTransactorSession) WithdrawBitcoinPSBTRequest(pod common.Address, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawBitcoinPSBTRequest(&_PodManager.TransactOpts, pod, withdrawAddress)
}

// WithdrawPresignedBitcoinRequest is a paid mutator transaction binding the contract method 0xb33ceeff.
//
// Solidity: function withdrawPresignedBitcoinRequest(address pod, string withdrawAddress) returns()
func (_PodManager *PodManagerTransactor) WithdrawPresignedBitcoinRequest(opts *bind.TransactOpts, pod common.Address, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.contract.Transact(opts, "withdrawPresignedBitcoinRequest", pod, withdrawAddress)
}

// WithdrawPresignedBitcoinRequest is a paid mutator transaction binding the contract method 0xb33ceeff.
//
// Solidity: function withdrawPresignedBitcoinRequest(address pod, string withdrawAddress) returns()
func (_PodManager *PodManagerSession) WithdrawPresignedBitcoinRequest(pod common.Address, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawPresignedBitcoinRequest(&_PodManager.TransactOpts, pod, withdrawAddress)
}

// WithdrawPresignedBitcoinRequest is a paid mutator transaction binding the contract method 0xb33ceeff.
//
// Solidity: function withdrawPresignedBitcoinRequest(address pod, string withdrawAddress) returns()
func (_PodManager *PodManagerTransactorSession) WithdrawPresignedBitcoinRequest(pod common.Address, withdrawAddress string) (*types.Transaction, error) {
	return _PodManager.Contract.WithdrawPresignedBitcoinRequest(&_PodManager.TransactOpts, pod, withdrawAddress)
}

// PodManagerBTCAddressVerifiedIterator is returned from FilterBTCAddressVerified and is used to iterate over the raw logs and unpacked data for BTCAddressVerified events raised by the PodManager contract.
type PodManagerBTCAddressVerifiedIterator struct {
	Event *PodManagerBTCAddressVerified // Event containing the contract specifics and raw log

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
func (it *PodManagerBTCAddressVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBTCAddressVerified)
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
		it.Event = new(PodManagerBTCAddressVerified)
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
func (it *PodManagerBTCAddressVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBTCAddressVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBTCAddressVerified represents a BTCAddressVerified event raised by the PodManager contract.
type PodManagerBTCAddressVerified struct {
	Operator   common.Address
	BtcAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBTCAddressVerified is a free log retrieval operation binding the contract event 0xabc3bc0a50a83331466b147d3ce6d2a71cc52bfa15d52467b4860bb8d08c206b.
//
// Solidity: event BTCAddressVerified(address indexed operator, string btcAddress)
func (_PodManager *PodManagerFilterer) FilterBTCAddressVerified(opts *bind.FilterOpts, operator []common.Address) (*PodManagerBTCAddressVerifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BTCAddressVerified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBTCAddressVerifiedIterator{contract: _PodManager.contract, event: "BTCAddressVerified", logs: logs, sub: sub}, nil
}

// WatchBTCAddressVerified is a free log subscription operation binding the contract event 0xabc3bc0a50a83331466b147d3ce6d2a71cc52bfa15d52467b4860bb8d08c206b.
//
// Solidity: event BTCAddressVerified(address indexed operator, string btcAddress)
func (_PodManager *PodManagerFilterer) WatchBTCAddressVerified(opts *bind.WatchOpts, sink chan<- *PodManagerBTCAddressVerified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BTCAddressVerified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBTCAddressVerified)
				if err := _PodManager.contract.UnpackLog(event, "BTCAddressVerified", log); err != nil {
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

// ParseBTCAddressVerified is a log parse operation binding the contract event 0xabc3bc0a50a83331466b147d3ce6d2a71cc52bfa15d52467b4860bb8d08c206b.
//
// Solidity: event BTCAddressVerified(address indexed operator, string btcAddress)
func (_PodManager *PodManagerFilterer) ParseBTCAddressVerified(log types.Log) (*PodManagerBTCAddressVerified, error) {
	event := new(PodManagerBTCAddressVerified)
	if err := _PodManager.contract.UnpackLog(event, "BTCAddressVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBitcoinBurnedIterator is returned from FilterBitcoinBurned and is used to iterate over the raw logs and unpacked data for BitcoinBurned events raised by the PodManager contract.
type PodManagerBitcoinBurnedIterator struct {
	Event *PodManagerBitcoinBurned // Event containing the contract specifics and raw log

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
func (it *PodManagerBitcoinBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBitcoinBurned)
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
		it.Event = new(PodManagerBitcoinBurned)
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
func (it *PodManagerBitcoinBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBitcoinBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBitcoinBurned represents a BitcoinBurned event raised by the PodManager contract.
type PodManagerBitcoinBurned struct {
	Pod    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBitcoinBurned is a free log retrieval operation binding the contract event 0xab3f63d5d40e3fa7b8c0bece27a98227873fd63e10d0da6e8fcefc6e510d3f1d.
//
// Solidity: event BitcoinBurned(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) FilterBitcoinBurned(opts *bind.FilterOpts, pod []common.Address) (*PodManagerBitcoinBurnedIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BitcoinBurned", podRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBitcoinBurnedIterator{contract: _PodManager.contract, event: "BitcoinBurned", logs: logs, sub: sub}, nil
}

// WatchBitcoinBurned is a free log subscription operation binding the contract event 0xab3f63d5d40e3fa7b8c0bece27a98227873fd63e10d0da6e8fcefc6e510d3f1d.
//
// Solidity: event BitcoinBurned(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) WatchBitcoinBurned(opts *bind.WatchOpts, sink chan<- *PodManagerBitcoinBurned, pod []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BitcoinBurned", podRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBitcoinBurned)
				if err := _PodManager.contract.UnpackLog(event, "BitcoinBurned", log); err != nil {
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

// ParseBitcoinBurned is a log parse operation binding the contract event 0xab3f63d5d40e3fa7b8c0bece27a98227873fd63e10d0da6e8fcefc6e510d3f1d.
//
// Solidity: event BitcoinBurned(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) ParseBitcoinBurned(log types.Log) (*PodManagerBitcoinBurned, error) {
	event := new(PodManagerBitcoinBurned)
	if err := _PodManager.contract.UnpackLog(event, "BitcoinBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBitcoinDepositConfirmedIterator is returned from FilterBitcoinDepositConfirmed and is used to iterate over the raw logs and unpacked data for BitcoinDepositConfirmed events raised by the PodManager contract.
type PodManagerBitcoinDepositConfirmedIterator struct {
	Event *PodManagerBitcoinDepositConfirmed // Event containing the contract specifics and raw log

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
func (it *PodManagerBitcoinDepositConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBitcoinDepositConfirmed)
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
		it.Event = new(PodManagerBitcoinDepositConfirmed)
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
func (it *PodManagerBitcoinDepositConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBitcoinDepositConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBitcoinDepositConfirmed represents a BitcoinDepositConfirmed event raised by the PodManager contract.
type PodManagerBitcoinDepositConfirmed struct {
	Pod    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBitcoinDepositConfirmed is a free log retrieval operation binding the contract event 0xcd774d9acd3dba4b6906137578fb72af62a91a12e8943e076232cb691cc96cef.
//
// Solidity: event BitcoinDepositConfirmed(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) FilterBitcoinDepositConfirmed(opts *bind.FilterOpts, pod []common.Address) (*PodManagerBitcoinDepositConfirmedIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BitcoinDepositConfirmed", podRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBitcoinDepositConfirmedIterator{contract: _PodManager.contract, event: "BitcoinDepositConfirmed", logs: logs, sub: sub}, nil
}

// WatchBitcoinDepositConfirmed is a free log subscription operation binding the contract event 0xcd774d9acd3dba4b6906137578fb72af62a91a12e8943e076232cb691cc96cef.
//
// Solidity: event BitcoinDepositConfirmed(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) WatchBitcoinDepositConfirmed(opts *bind.WatchOpts, sink chan<- *PodManagerBitcoinDepositConfirmed, pod []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BitcoinDepositConfirmed", podRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBitcoinDepositConfirmed)
				if err := _PodManager.contract.UnpackLog(event, "BitcoinDepositConfirmed", log); err != nil {
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

// ParseBitcoinDepositConfirmed is a log parse operation binding the contract event 0xcd774d9acd3dba4b6906137578fb72af62a91a12e8943e076232cb691cc96cef.
//
// Solidity: event BitcoinDepositConfirmed(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) ParseBitcoinDepositConfirmed(log types.Log) (*PodManagerBitcoinDepositConfirmed, error) {
	event := new(PodManagerBitcoinDepositConfirmed)
	if err := _PodManager.contract.UnpackLog(event, "BitcoinDepositConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBitcoinMintedIterator is returned from FilterBitcoinMinted and is used to iterate over the raw logs and unpacked data for BitcoinMinted events raised by the PodManager contract.
type PodManagerBitcoinMintedIterator struct {
	Event *PodManagerBitcoinMinted // Event containing the contract specifics and raw log

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
func (it *PodManagerBitcoinMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBitcoinMinted)
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
		it.Event = new(PodManagerBitcoinMinted)
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
func (it *PodManagerBitcoinMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBitcoinMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBitcoinMinted represents a BitcoinMinted event raised by the PodManager contract.
type PodManagerBitcoinMinted struct {
	Pod    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBitcoinMinted is a free log retrieval operation binding the contract event 0x5aca106588d0e1ebb9128b3f4b196bc949a208e38d3db913363353402fe73bd3.
//
// Solidity: event BitcoinMinted(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) FilterBitcoinMinted(opts *bind.FilterOpts, pod []common.Address) (*PodManagerBitcoinMintedIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BitcoinMinted", podRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBitcoinMintedIterator{contract: _PodManager.contract, event: "BitcoinMinted", logs: logs, sub: sub}, nil
}

// WatchBitcoinMinted is a free log subscription operation binding the contract event 0x5aca106588d0e1ebb9128b3f4b196bc949a208e38d3db913363353402fe73bd3.
//
// Solidity: event BitcoinMinted(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) WatchBitcoinMinted(opts *bind.WatchOpts, sink chan<- *PodManagerBitcoinMinted, pod []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BitcoinMinted", podRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBitcoinMinted)
				if err := _PodManager.contract.UnpackLog(event, "BitcoinMinted", log); err != nil {
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

// ParseBitcoinMinted is a log parse operation binding the contract event 0x5aca106588d0e1ebb9128b3f4b196bc949a208e38d3db913363353402fe73bd3.
//
// Solidity: event BitcoinMinted(address indexed pod, uint256 amount)
func (_PodManager *PodManagerFilterer) ParseBitcoinMinted(log types.Log) (*PodManagerBitcoinMinted, error) {
	event := new(PodManagerBitcoinMinted)
	if err := _PodManager.contract.UnpackLog(event, "BitcoinMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBitcoinWithdrawalCompleteTxRequestIterator is returned from FilterBitcoinWithdrawalCompleteTxRequest and is used to iterate over the raw logs and unpacked data for BitcoinWithdrawalCompleteTxRequest events raised by the PodManager contract.
type PodManagerBitcoinWithdrawalCompleteTxRequestIterator struct {
	Event *PodManagerBitcoinWithdrawalCompleteTxRequest // Event containing the contract specifics and raw log

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
func (it *PodManagerBitcoinWithdrawalCompleteTxRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBitcoinWithdrawalCompleteTxRequest)
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
		it.Event = new(PodManagerBitcoinWithdrawalCompleteTxRequest)
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
func (it *PodManagerBitcoinWithdrawalCompleteTxRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBitcoinWithdrawalCompleteTxRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBitcoinWithdrawalCompleteTxRequest represents a BitcoinWithdrawalCompleteTxRequest event raised by the PodManager contract.
type PodManagerBitcoinWithdrawalCompleteTxRequest struct {
	Pod                common.Address
	Operator           common.Address
	PreSignedBitcoinTx []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBitcoinWithdrawalCompleteTxRequest is a free log retrieval operation binding the contract event 0xef73d097c5517b2b5d2dbffb17b0f948b5d7606055610ed1cad57aaa439e94f6.
//
// Solidity: event BitcoinWithdrawalCompleteTxRequest(address indexed pod, address indexed operator, bytes preSignedBitcoinTx)
func (_PodManager *PodManagerFilterer) FilterBitcoinWithdrawalCompleteTxRequest(opts *bind.FilterOpts, pod []common.Address, operator []common.Address) (*PodManagerBitcoinWithdrawalCompleteTxRequestIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BitcoinWithdrawalCompleteTxRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBitcoinWithdrawalCompleteTxRequestIterator{contract: _PodManager.contract, event: "BitcoinWithdrawalCompleteTxRequest", logs: logs, sub: sub}, nil
}

// WatchBitcoinWithdrawalCompleteTxRequest is a free log subscription operation binding the contract event 0xef73d097c5517b2b5d2dbffb17b0f948b5d7606055610ed1cad57aaa439e94f6.
//
// Solidity: event BitcoinWithdrawalCompleteTxRequest(address indexed pod, address indexed operator, bytes preSignedBitcoinTx)
func (_PodManager *PodManagerFilterer) WatchBitcoinWithdrawalCompleteTxRequest(opts *bind.WatchOpts, sink chan<- *PodManagerBitcoinWithdrawalCompleteTxRequest, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BitcoinWithdrawalCompleteTxRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBitcoinWithdrawalCompleteTxRequest)
				if err := _PodManager.contract.UnpackLog(event, "BitcoinWithdrawalCompleteTxRequest", log); err != nil {
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

// ParseBitcoinWithdrawalCompleteTxRequest is a log parse operation binding the contract event 0xef73d097c5517b2b5d2dbffb17b0f948b5d7606055610ed1cad57aaa439e94f6.
//
// Solidity: event BitcoinWithdrawalCompleteTxRequest(address indexed pod, address indexed operator, bytes preSignedBitcoinTx)
func (_PodManager *PodManagerFilterer) ParseBitcoinWithdrawalCompleteTxRequest(log types.Log) (*PodManagerBitcoinWithdrawalCompleteTxRequest, error) {
	event := new(PodManagerBitcoinWithdrawalCompleteTxRequest)
	if err := _PodManager.contract.UnpackLog(event, "BitcoinWithdrawalCompleteTxRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBitcoinWithdrawalPSBTRequestIterator is returned from FilterBitcoinWithdrawalPSBTRequest and is used to iterate over the raw logs and unpacked data for BitcoinWithdrawalPSBTRequest events raised by the PodManager contract.
type PodManagerBitcoinWithdrawalPSBTRequestIterator struct {
	Event *PodManagerBitcoinWithdrawalPSBTRequest // Event containing the contract specifics and raw log

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
func (it *PodManagerBitcoinWithdrawalPSBTRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBitcoinWithdrawalPSBTRequest)
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
		it.Event = new(PodManagerBitcoinWithdrawalPSBTRequest)
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
func (it *PodManagerBitcoinWithdrawalPSBTRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBitcoinWithdrawalPSBTRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBitcoinWithdrawalPSBTRequest represents a BitcoinWithdrawalPSBTRequest event raised by the PodManager contract.
type PodManagerBitcoinWithdrawalPSBTRequest struct {
	Pod             common.Address
	Operator        common.Address
	WithdrawAddress string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBitcoinWithdrawalPSBTRequest is a free log retrieval operation binding the contract event 0xe6142438106d9245d23d95a8dcc0141c9e63e96b0e2a9655b213ca6bd50af011.
//
// Solidity: event BitcoinWithdrawalPSBTRequest(address indexed pod, address indexed operator, string withdrawAddress)
func (_PodManager *PodManagerFilterer) FilterBitcoinWithdrawalPSBTRequest(opts *bind.FilterOpts, pod []common.Address, operator []common.Address) (*PodManagerBitcoinWithdrawalPSBTRequestIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BitcoinWithdrawalPSBTRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBitcoinWithdrawalPSBTRequestIterator{contract: _PodManager.contract, event: "BitcoinWithdrawalPSBTRequest", logs: logs, sub: sub}, nil
}

// WatchBitcoinWithdrawalPSBTRequest is a free log subscription operation binding the contract event 0xe6142438106d9245d23d95a8dcc0141c9e63e96b0e2a9655b213ca6bd50af011.
//
// Solidity: event BitcoinWithdrawalPSBTRequest(address indexed pod, address indexed operator, string withdrawAddress)
func (_PodManager *PodManagerFilterer) WatchBitcoinWithdrawalPSBTRequest(opts *bind.WatchOpts, sink chan<- *PodManagerBitcoinWithdrawalPSBTRequest, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BitcoinWithdrawalPSBTRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBitcoinWithdrawalPSBTRequest)
				if err := _PodManager.contract.UnpackLog(event, "BitcoinWithdrawalPSBTRequest", log); err != nil {
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

// ParseBitcoinWithdrawalPSBTRequest is a log parse operation binding the contract event 0xe6142438106d9245d23d95a8dcc0141c9e63e96b0e2a9655b213ca6bd50af011.
//
// Solidity: event BitcoinWithdrawalPSBTRequest(address indexed pod, address indexed operator, string withdrawAddress)
func (_PodManager *PodManagerFilterer) ParseBitcoinWithdrawalPSBTRequest(log types.Log) (*PodManagerBitcoinWithdrawalPSBTRequest, error) {
	event := new(PodManagerBitcoinWithdrawalPSBTRequest)
	if err := _PodManager.contract.UnpackLog(event, "BitcoinWithdrawalPSBTRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerBitcoinWithdrawnFromPodIterator is returned from FilterBitcoinWithdrawnFromPod and is used to iterate over the raw logs and unpacked data for BitcoinWithdrawnFromPod events raised by the PodManager contract.
type PodManagerBitcoinWithdrawnFromPodIterator struct {
	Event *PodManagerBitcoinWithdrawnFromPod // Event containing the contract specifics and raw log

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
func (it *PodManagerBitcoinWithdrawnFromPodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerBitcoinWithdrawnFromPod)
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
		it.Event = new(PodManagerBitcoinWithdrawnFromPod)
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
func (it *PodManagerBitcoinWithdrawnFromPodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerBitcoinWithdrawnFromPodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerBitcoinWithdrawnFromPod represents a BitcoinWithdrawnFromPod event raised by the PodManager contract.
type PodManagerBitcoinWithdrawnFromPod struct {
	Pod             common.Address
	WithdrawAddress string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBitcoinWithdrawnFromPod is a free log retrieval operation binding the contract event 0x6dd3dd303209228b16cce56d4dfeb8116c23e6b9374927841bb4557b1e686643.
//
// Solidity: event BitcoinWithdrawnFromPod(address indexed pod, string withdrawAddress)
func (_PodManager *PodManagerFilterer) FilterBitcoinWithdrawnFromPod(opts *bind.FilterOpts, pod []common.Address) (*PodManagerBitcoinWithdrawnFromPodIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "BitcoinWithdrawnFromPod", podRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerBitcoinWithdrawnFromPodIterator{contract: _PodManager.contract, event: "BitcoinWithdrawnFromPod", logs: logs, sub: sub}, nil
}

// WatchBitcoinWithdrawnFromPod is a free log subscription operation binding the contract event 0x6dd3dd303209228b16cce56d4dfeb8116c23e6b9374927841bb4557b1e686643.
//
// Solidity: event BitcoinWithdrawnFromPod(address indexed pod, string withdrawAddress)
func (_PodManager *PodManagerFilterer) WatchBitcoinWithdrawnFromPod(opts *bind.WatchOpts, sink chan<- *PodManagerBitcoinWithdrawnFromPod, pod []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "BitcoinWithdrawnFromPod", podRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerBitcoinWithdrawnFromPod)
				if err := _PodManager.contract.UnpackLog(event, "BitcoinWithdrawnFromPod", log); err != nil {
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

// ParseBitcoinWithdrawnFromPod is a log parse operation binding the contract event 0x6dd3dd303209228b16cce56d4dfeb8116c23e6b9374927841bb4557b1e686643.
//
// Solidity: event BitcoinWithdrawnFromPod(address indexed pod, string withdrawAddress)
func (_PodManager *PodManagerFilterer) ParseBitcoinWithdrawnFromPod(log types.Log) (*PodManagerBitcoinWithdrawnFromPod, error) {
	event := new(PodManagerBitcoinWithdrawnFromPod)
	if err := _PodManager.contract.UnpackLog(event, "BitcoinWithdrawnFromPod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PodManager contract.
type PodManagerInitializedIterator struct {
	Event *PodManagerInitialized // Event containing the contract specifics and raw log

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
func (it *PodManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerInitialized)
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
		it.Event = new(PodManagerInitialized)
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
func (it *PodManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerInitialized represents a Initialized event raised by the PodManager contract.
type PodManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PodManager *PodManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PodManagerInitializedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PodManagerInitializedIterator{contract: _PodManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PodManager *PodManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PodManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerInitialized)
				if err := _PodManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PodManager *PodManagerFilterer) ParseInitialized(log types.Log) (*PodManagerInitialized, error) {
	event := new(PodManagerInitialized)
	if err := _PodManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PodManager contract.
type PodManagerOwnershipTransferredIterator struct {
	Event *PodManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PodManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerOwnershipTransferred)
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
		it.Event = new(PodManagerOwnershipTransferred)
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
func (it *PodManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PodManager contract.
type PodManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PodManager *PodManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PodManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerOwnershipTransferredIterator{contract: _PodManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PodManager *PodManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PodManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerOwnershipTransferred)
				if err := _PodManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PodManager *PodManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PodManagerOwnershipTransferred, error) {
	event := new(PodManagerOwnershipTransferred)
	if err := _PodManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PodManager contract.
type PodManagerPausedIterator struct {
	Event *PodManagerPaused // Event containing the contract specifics and raw log

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
func (it *PodManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPaused)
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
		it.Event = new(PodManagerPaused)
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
func (it *PodManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPaused represents a Paused event raised by the PodManager contract.
type PodManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PodManager *PodManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*PodManagerPausedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PodManagerPausedIterator{contract: _PodManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PodManager *PodManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PodManagerPaused) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPaused)
				if err := _PodManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PodManager *PodManagerFilterer) ParsePaused(log types.Log) (*PodManagerPaused, error) {
	event := new(PodManagerPaused)
	if err := _PodManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPodCreatedIterator is returned from FilterPodCreated and is used to iterate over the raw logs and unpacked data for PodCreated events raised by the PodManager contract.
type PodManagerPodCreatedIterator struct {
	Event *PodManagerPodCreated // Event containing the contract specifics and raw log

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
func (it *PodManagerPodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPodCreated)
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
		it.Event = new(PodManagerPodCreated)
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
func (it *PodManagerPodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPodCreated represents a PodCreated event raised by the PodManager contract.
type PodManagerPodCreated struct {
	User     common.Address
	Pod      common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPodCreated is a free log retrieval operation binding the contract event 0x9bec3cb55fcd181aaece51ed577f8a95b847c8af90535a81a7fa7d1c20138ef9.
//
// Solidity: event PodCreated(address indexed user, address indexed pod, address indexed operator)
func (_PodManager *PodManagerFilterer) FilterPodCreated(opts *bind.FilterOpts, user []common.Address, pod []common.Address, operator []common.Address) (*PodManagerPodCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "PodCreated", userRule, podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerPodCreatedIterator{contract: _PodManager.contract, event: "PodCreated", logs: logs, sub: sub}, nil
}

// WatchPodCreated is a free log subscription operation binding the contract event 0x9bec3cb55fcd181aaece51ed577f8a95b847c8af90535a81a7fa7d1c20138ef9.
//
// Solidity: event PodCreated(address indexed user, address indexed pod, address indexed operator)
func (_PodManager *PodManagerFilterer) WatchPodCreated(opts *bind.WatchOpts, sink chan<- *PodManagerPodCreated, user []common.Address, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "PodCreated", userRule, podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPodCreated)
				if err := _PodManager.contract.UnpackLog(event, "PodCreated", log); err != nil {
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

// ParsePodCreated is a log parse operation binding the contract event 0x9bec3cb55fcd181aaece51ed577f8a95b847c8af90535a81a7fa7d1c20138ef9.
//
// Solidity: event PodCreated(address indexed user, address indexed pod, address indexed operator)
func (_PodManager *PodManagerFilterer) ParsePodCreated(log types.Log) (*PodManagerPodCreated, error) {
	event := new(PodManagerPodCreated)
	if err := _PodManager.contract.UnpackLog(event, "PodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPodDelegatedIterator is returned from FilterPodDelegated and is used to iterate over the raw logs and unpacked data for PodDelegated events raised by the PodManager contract.
type PodManagerPodDelegatedIterator struct {
	Event *PodManagerPodDelegated // Event containing the contract specifics and raw log

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
func (it *PodManagerPodDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPodDelegated)
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
		it.Event = new(PodManagerPodDelegated)
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
func (it *PodManagerPodDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPodDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPodDelegated represents a PodDelegated event raised by the PodManager contract.
type PodManagerPodDelegated struct {
	Pod         common.Address
	AppContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPodDelegated is a free log retrieval operation binding the contract event 0xef32074b263175118a4614cc3e2742c50b70facbf865c0c26da767dd095e167d.
//
// Solidity: event PodDelegated(address indexed pod, address indexed appContract)
func (_PodManager *PodManagerFilterer) FilterPodDelegated(opts *bind.FilterOpts, pod []common.Address, appContract []common.Address) (*PodManagerPodDelegatedIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var appContractRule []interface{}
	for _, appContractItem := range appContract {
		appContractRule = append(appContractRule, appContractItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "PodDelegated", podRule, appContractRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerPodDelegatedIterator{contract: _PodManager.contract, event: "PodDelegated", logs: logs, sub: sub}, nil
}

// WatchPodDelegated is a free log subscription operation binding the contract event 0xef32074b263175118a4614cc3e2742c50b70facbf865c0c26da767dd095e167d.
//
// Solidity: event PodDelegated(address indexed pod, address indexed appContract)
func (_PodManager *PodManagerFilterer) WatchPodDelegated(opts *bind.WatchOpts, sink chan<- *PodManagerPodDelegated, pod []common.Address, appContract []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var appContractRule []interface{}
	for _, appContractItem := range appContract {
		appContractRule = append(appContractRule, appContractItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "PodDelegated", podRule, appContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPodDelegated)
				if err := _PodManager.contract.UnpackLog(event, "PodDelegated", log); err != nil {
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

// ParsePodDelegated is a log parse operation binding the contract event 0xef32074b263175118a4614cc3e2742c50b70facbf865c0c26da767dd095e167d.
//
// Solidity: event PodDelegated(address indexed pod, address indexed appContract)
func (_PodManager *PodManagerFilterer) ParsePodDelegated(log types.Log) (*PodManagerPodDelegated, error) {
	event := new(PodManagerPodDelegated)
	if err := _PodManager.contract.UnpackLog(event, "PodDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerPodUndelegatedIterator is returned from FilterPodUndelegated and is used to iterate over the raw logs and unpacked data for PodUndelegated events raised by the PodManager contract.
type PodManagerPodUndelegatedIterator struct {
	Event *PodManagerPodUndelegated // Event containing the contract specifics and raw log

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
func (it *PodManagerPodUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerPodUndelegated)
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
		it.Event = new(PodManagerPodUndelegated)
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
func (it *PodManagerPodUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerPodUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerPodUndelegated represents a PodUndelegated event raised by the PodManager contract.
type PodManagerPodUndelegated struct {
	Pod common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPodUndelegated is a free log retrieval operation binding the contract event 0xfde114ec59a374c36954512dafbedfd68f4c4d137190476a76cddccd24697a3a.
//
// Solidity: event PodUndelegated(address indexed pod)
func (_PodManager *PodManagerFilterer) FilterPodUndelegated(opts *bind.FilterOpts, pod []common.Address) (*PodManagerPodUndelegatedIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "PodUndelegated", podRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerPodUndelegatedIterator{contract: _PodManager.contract, event: "PodUndelegated", logs: logs, sub: sub}, nil
}

// WatchPodUndelegated is a free log subscription operation binding the contract event 0xfde114ec59a374c36954512dafbedfd68f4c4d137190476a76cddccd24697a3a.
//
// Solidity: event PodUndelegated(address indexed pod)
func (_PodManager *PodManagerFilterer) WatchPodUndelegated(opts *bind.WatchOpts, sink chan<- *PodManagerPodUndelegated, pod []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "PodUndelegated", podRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerPodUndelegated)
				if err := _PodManager.contract.UnpackLog(event, "PodUndelegated", log); err != nil {
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

// ParsePodUndelegated is a log parse operation binding the contract event 0xfde114ec59a374c36954512dafbedfd68f4c4d137190476a76cddccd24697a3a.
//
// Solidity: event PodUndelegated(address indexed pod)
func (_PodManager *PodManagerFilterer) ParsePodUndelegated(log types.Log) (*PodManagerPodUndelegated, error) {
	event := new(PodManagerPodUndelegated)
	if err := _PodManager.contract.UnpackLog(event, "PodUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerTotalTVLUpdatedIterator is returned from FilterTotalTVLUpdated and is used to iterate over the raw logs and unpacked data for TotalTVLUpdated events raised by the PodManager contract.
type PodManagerTotalTVLUpdatedIterator struct {
	Event *PodManagerTotalTVLUpdated // Event containing the contract specifics and raw log

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
func (it *PodManagerTotalTVLUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerTotalTVLUpdated)
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
		it.Event = new(PodManagerTotalTVLUpdated)
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
func (it *PodManagerTotalTVLUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerTotalTVLUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerTotalTVLUpdated represents a TotalTVLUpdated event raised by the PodManager contract.
type PodManagerTotalTVLUpdated struct {
	NewTVL *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTotalTVLUpdated is a free log retrieval operation binding the contract event 0x6781bb9c83a65d946aefe506cbfb0817b2c859b3a822b74302fc199678c680de.
//
// Solidity: event TotalTVLUpdated(uint256 newTVL)
func (_PodManager *PodManagerFilterer) FilterTotalTVLUpdated(opts *bind.FilterOpts) (*PodManagerTotalTVLUpdatedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "TotalTVLUpdated")
	if err != nil {
		return nil, err
	}
	return &PodManagerTotalTVLUpdatedIterator{contract: _PodManager.contract, event: "TotalTVLUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalTVLUpdated is a free log subscription operation binding the contract event 0x6781bb9c83a65d946aefe506cbfb0817b2c859b3a822b74302fc199678c680de.
//
// Solidity: event TotalTVLUpdated(uint256 newTVL)
func (_PodManager *PodManagerFilterer) WatchTotalTVLUpdated(opts *bind.WatchOpts, sink chan<- *PodManagerTotalTVLUpdated) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "TotalTVLUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerTotalTVLUpdated)
				if err := _PodManager.contract.UnpackLog(event, "TotalTVLUpdated", log); err != nil {
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

// ParseTotalTVLUpdated is a log parse operation binding the contract event 0x6781bb9c83a65d946aefe506cbfb0817b2c859b3a822b74302fc199678c680de.
//
// Solidity: event TotalTVLUpdated(uint256 newTVL)
func (_PodManager *PodManagerFilterer) ParseTotalTVLUpdated(log types.Log) (*PodManagerTotalTVLUpdated, error) {
	event := new(PodManagerTotalTVLUpdated)
	if err := _PodManager.contract.UnpackLog(event, "TotalTVLUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PodManager contract.
type PodManagerUnpausedIterator struct {
	Event *PodManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *PodManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerUnpaused)
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
		it.Event = new(PodManagerUnpaused)
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
func (it *PodManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerUnpaused represents a Unpaused event raised by the PodManager contract.
type PodManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PodManager *PodManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PodManagerUnpausedIterator, error) {

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PodManagerUnpausedIterator{contract: _PodManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PodManager *PodManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PodManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerUnpaused)
				if err := _PodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PodManager *PodManagerFilterer) ParseUnpaused(log types.Log) (*PodManagerUnpaused, error) {
	event := new(PodManagerUnpaused)
	if err := _PodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerVerifyBitcoinDepositRequestIterator is returned from FilterVerifyBitcoinDepositRequest and is used to iterate over the raw logs and unpacked data for VerifyBitcoinDepositRequest events raised by the PodManager contract.
type PodManagerVerifyBitcoinDepositRequestIterator struct {
	Event *PodManagerVerifyBitcoinDepositRequest // Event containing the contract specifics and raw log

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
func (it *PodManagerVerifyBitcoinDepositRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerVerifyBitcoinDepositRequest)
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
		it.Event = new(PodManagerVerifyBitcoinDepositRequest)
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
func (it *PodManagerVerifyBitcoinDepositRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerVerifyBitcoinDepositRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerVerifyBitcoinDepositRequest represents a VerifyBitcoinDepositRequest event raised by the PodManager contract.
type PodManagerVerifyBitcoinDepositRequest struct {
	Pod                   common.Address
	Operator              common.Address
	BitcoinDepositRequest IBitcoinPodManagerBitcoinDepositRequest
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterVerifyBitcoinDepositRequest is a free log retrieval operation binding the contract event 0x19e46e0ef253a265a935e867bddd4047e0ea6700a3b5254cc16dea197da7595e.
//
// Solidity: event VerifyBitcoinDepositRequest(address indexed pod, address indexed operator, (bytes32,uint256,bool) bitcoinDepositRequest)
func (_PodManager *PodManagerFilterer) FilterVerifyBitcoinDepositRequest(opts *bind.FilterOpts, pod []common.Address, operator []common.Address) (*PodManagerVerifyBitcoinDepositRequestIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "VerifyBitcoinDepositRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerVerifyBitcoinDepositRequestIterator{contract: _PodManager.contract, event: "VerifyBitcoinDepositRequest", logs: logs, sub: sub}, nil
}

// WatchVerifyBitcoinDepositRequest is a free log subscription operation binding the contract event 0x19e46e0ef253a265a935e867bddd4047e0ea6700a3b5254cc16dea197da7595e.
//
// Solidity: event VerifyBitcoinDepositRequest(address indexed pod, address indexed operator, (bytes32,uint256,bool) bitcoinDepositRequest)
func (_PodManager *PodManagerFilterer) WatchVerifyBitcoinDepositRequest(opts *bind.WatchOpts, sink chan<- *PodManagerVerifyBitcoinDepositRequest, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "VerifyBitcoinDepositRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerVerifyBitcoinDepositRequest)
				if err := _PodManager.contract.UnpackLog(event, "VerifyBitcoinDepositRequest", log); err != nil {
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

// ParseVerifyBitcoinDepositRequest is a log parse operation binding the contract event 0x19e46e0ef253a265a935e867bddd4047e0ea6700a3b5254cc16dea197da7595e.
//
// Solidity: event VerifyBitcoinDepositRequest(address indexed pod, address indexed operator, (bytes32,uint256,bool) bitcoinDepositRequest)
func (_PodManager *PodManagerFilterer) ParseVerifyBitcoinDepositRequest(log types.Log) (*PodManagerVerifyBitcoinDepositRequest, error) {
	event := new(PodManagerVerifyBitcoinDepositRequest)
	if err := _PodManager.contract.UnpackLog(event, "VerifyBitcoinDepositRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerVerifyPresignedBitcoinDepositRequestIterator is returned from FilterVerifyPresignedBitcoinDepositRequest and is used to iterate over the raw logs and unpacked data for VerifyPresignedBitcoinDepositRequest events raised by the PodManager contract.
type PodManagerVerifyPresignedBitcoinDepositRequestIterator struct {
	Event *PodManagerVerifyPresignedBitcoinDepositRequest // Event containing the contract specifics and raw log

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
func (it *PodManagerVerifyPresignedBitcoinDepositRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerVerifyPresignedBitcoinDepositRequest)
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
		it.Event = new(PodManagerVerifyPresignedBitcoinDepositRequest)
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
func (it *PodManagerVerifyPresignedBitcoinDepositRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerVerifyPresignedBitcoinDepositRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerVerifyPresignedBitcoinDepositRequest represents a VerifyPresignedBitcoinDepositRequest event raised by the PodManager contract.
type PodManagerVerifyPresignedBitcoinDepositRequest struct {
	Pod                   common.Address
	Operator              common.Address
	BitcoinDepositRequest IBitcoinPodManagerBitcoinDepositRequest
	Transaction           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterVerifyPresignedBitcoinDepositRequest is a free log retrieval operation binding the contract event 0x18335a2783f18e0b12893137edcf4cc4efb74d9658939b1e2de3b2a69c1cbe8f.
//
// Solidity: event VerifyPresignedBitcoinDepositRequest(address indexed pod, address indexed operator, (bytes32,uint256,bool) bitcoinDepositRequest, bytes transaction)
func (_PodManager *PodManagerFilterer) FilterVerifyPresignedBitcoinDepositRequest(opts *bind.FilterOpts, pod []common.Address, operator []common.Address) (*PodManagerVerifyPresignedBitcoinDepositRequestIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "VerifyPresignedBitcoinDepositRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerVerifyPresignedBitcoinDepositRequestIterator{contract: _PodManager.contract, event: "VerifyPresignedBitcoinDepositRequest", logs: logs, sub: sub}, nil
}

// WatchVerifyPresignedBitcoinDepositRequest is a free log subscription operation binding the contract event 0x18335a2783f18e0b12893137edcf4cc4efb74d9658939b1e2de3b2a69c1cbe8f.
//
// Solidity: event VerifyPresignedBitcoinDepositRequest(address indexed pod, address indexed operator, (bytes32,uint256,bool) bitcoinDepositRequest, bytes transaction)
func (_PodManager *PodManagerFilterer) WatchVerifyPresignedBitcoinDepositRequest(opts *bind.WatchOpts, sink chan<- *PodManagerVerifyPresignedBitcoinDepositRequest, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "VerifyPresignedBitcoinDepositRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerVerifyPresignedBitcoinDepositRequest)
				if err := _PodManager.contract.UnpackLog(event, "VerifyPresignedBitcoinDepositRequest", log); err != nil {
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

// ParseVerifyPresignedBitcoinDepositRequest is a log parse operation binding the contract event 0x18335a2783f18e0b12893137edcf4cc4efb74d9658939b1e2de3b2a69c1cbe8f.
//
// Solidity: event VerifyPresignedBitcoinDepositRequest(address indexed pod, address indexed operator, (bytes32,uint256,bool) bitcoinDepositRequest, bytes transaction)
func (_PodManager *PodManagerFilterer) ParseVerifyPresignedBitcoinDepositRequest(log types.Log) (*PodManagerVerifyPresignedBitcoinDepositRequest, error) {
	event := new(PodManagerVerifyPresignedBitcoinDepositRequest)
	if err := _PodManager.contract.UnpackLog(event, "VerifyPresignedBitcoinDepositRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerWithdrawPresignedBitcoinRequestIterator is returned from FilterWithdrawPresignedBitcoinRequest and is used to iterate over the raw logs and unpacked data for WithdrawPresignedBitcoinRequest events raised by the PodManager contract.
type PodManagerWithdrawPresignedBitcoinRequestIterator struct {
	Event *PodManagerWithdrawPresignedBitcoinRequest // Event containing the contract specifics and raw log

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
func (it *PodManagerWithdrawPresignedBitcoinRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerWithdrawPresignedBitcoinRequest)
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
		it.Event = new(PodManagerWithdrawPresignedBitcoinRequest)
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
func (it *PodManagerWithdrawPresignedBitcoinRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerWithdrawPresignedBitcoinRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerWithdrawPresignedBitcoinRequest represents a WithdrawPresignedBitcoinRequest event raised by the PodManager contract.
type PodManagerWithdrawPresignedBitcoinRequest struct {
	Pod             common.Address
	Operator        common.Address
	WithdrawAddress string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawPresignedBitcoinRequest is a free log retrieval operation binding the contract event 0x319bd90f9360738fd1105f860deb133609e477d029e7be6d8b8508fee5207cae.
//
// Solidity: event WithdrawPresignedBitcoinRequest(address indexed pod, address indexed operator, string withdrawAddress)
func (_PodManager *PodManagerFilterer) FilterWithdrawPresignedBitcoinRequest(opts *bind.FilterOpts, pod []common.Address, operator []common.Address) (*PodManagerWithdrawPresignedBitcoinRequestIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "WithdrawPresignedBitcoinRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerWithdrawPresignedBitcoinRequestIterator{contract: _PodManager.contract, event: "WithdrawPresignedBitcoinRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawPresignedBitcoinRequest is a free log subscription operation binding the contract event 0x319bd90f9360738fd1105f860deb133609e477d029e7be6d8b8508fee5207cae.
//
// Solidity: event WithdrawPresignedBitcoinRequest(address indexed pod, address indexed operator, string withdrawAddress)
func (_PodManager *PodManagerFilterer) WatchWithdrawPresignedBitcoinRequest(opts *bind.WatchOpts, sink chan<- *PodManagerWithdrawPresignedBitcoinRequest, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "WithdrawPresignedBitcoinRequest", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerWithdrawPresignedBitcoinRequest)
				if err := _PodManager.contract.UnpackLog(event, "WithdrawPresignedBitcoinRequest", log); err != nil {
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

// ParseWithdrawPresignedBitcoinRequest is a log parse operation binding the contract event 0x319bd90f9360738fd1105f860deb133609e477d029e7be6d8b8508fee5207cae.
//
// Solidity: event WithdrawPresignedBitcoinRequest(address indexed pod, address indexed operator, string withdrawAddress)
func (_PodManager *PodManagerFilterer) ParseWithdrawPresignedBitcoinRequest(log types.Log) (*PodManagerWithdrawPresignedBitcoinRequest, error) {
	event := new(PodManagerWithdrawPresignedBitcoinRequest)
	if err := _PodManager.contract.UnpackLog(event, "WithdrawPresignedBitcoinRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PodManagerWithdrawalRequestCancelledIterator is returned from FilterWithdrawalRequestCancelled and is used to iterate over the raw logs and unpacked data for WithdrawalRequestCancelled events raised by the PodManager contract.
type PodManagerWithdrawalRequestCancelledIterator struct {
	Event *PodManagerWithdrawalRequestCancelled // Event containing the contract specifics and raw log

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
func (it *PodManagerWithdrawalRequestCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PodManagerWithdrawalRequestCancelled)
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
		it.Event = new(PodManagerWithdrawalRequestCancelled)
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
func (it *PodManagerWithdrawalRequestCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PodManagerWithdrawalRequestCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PodManagerWithdrawalRequestCancelled represents a WithdrawalRequestCancelled event raised by the PodManager contract.
type PodManagerWithdrawalRequestCancelled struct {
	Pod common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequestCancelled is a free log retrieval operation binding the contract event 0x943970facef781cb8ed6d08ce322d6e16ccaac1b591098e21464910c69fe5821.
//
// Solidity: event WithdrawalRequestCancelled(address indexed pod)
func (_PodManager *PodManagerFilterer) FilterWithdrawalRequestCancelled(opts *bind.FilterOpts, pod []common.Address) (*PodManagerWithdrawalRequestCancelledIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.FilterLogs(opts, "WithdrawalRequestCancelled", podRule)
	if err != nil {
		return nil, err
	}
	return &PodManagerWithdrawalRequestCancelledIterator{contract: _PodManager.contract, event: "WithdrawalRequestCancelled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequestCancelled is a free log subscription operation binding the contract event 0x943970facef781cb8ed6d08ce322d6e16ccaac1b591098e21464910c69fe5821.
//
// Solidity: event WithdrawalRequestCancelled(address indexed pod)
func (_PodManager *PodManagerFilterer) WatchWithdrawalRequestCancelled(opts *bind.WatchOpts, sink chan<- *PodManagerWithdrawalRequestCancelled, pod []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}

	logs, sub, err := _PodManager.contract.WatchLogs(opts, "WithdrawalRequestCancelled", podRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PodManagerWithdrawalRequestCancelled)
				if err := _PodManager.contract.UnpackLog(event, "WithdrawalRequestCancelled", log); err != nil {
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

// ParseWithdrawalRequestCancelled is a log parse operation binding the contract event 0x943970facef781cb8ed6d08ce322d6e16ccaac1b591098e21464910c69fe5821.
//
// Solidity: event WithdrawalRequestCancelled(address indexed pod)
func (_PodManager *PodManagerFilterer) ParseWithdrawalRequestCancelled(log types.Log) (*PodManagerWithdrawalRequestCancelled, error) {
	event := new(PodManagerWithdrawalRequestCancelled)
	if err := _PodManager.contract.UnpackLog(event, "WithdrawalRequestCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
