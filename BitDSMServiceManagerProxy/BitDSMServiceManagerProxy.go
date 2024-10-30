// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BitDSMServiceManagerProxy

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

// IBitDSMServiceManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IBitDSMServiceManagerTask struct {
	Name             string
	TaskCreatedBlock uint32
}

// IRewardsCoordinatorRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// BitDSMServiceManagerProxyMetaData contains all meta data concerning the BitDSMServiceManagerProxy contract.
var BitDSMServiceManagerProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmDeposit\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIBitDSMServiceManager.Task\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardsInitiator\",\"inputs\":[{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DepositConfirmed\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBitDSMServiceManager.Task\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBitDSMServiceManager.Task\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5060043610610111575f3560e01c80639926ee7d1161009e578063e481af9d1161006e578063e481af9d1461027f578063ead1131614610287578063f2fde38b1461029a578063fc299dee146102ad578063fce36c7d146102c0575f80fd5b80639926ee7d14610233578063a364f4da14610246578063a98fb35514610259578063c4d66de81461026c575f80fd5b80636b3aa72e116100e45780636b3aa72e146101bb578063715018a6146101e257806385edf874146101ea5780638b00ce7c146101fd5780638da5cb5b14610222575f80fd5b80632d89f6fc1461011557806333cfb7b7146101475780633bc28c8c14610167578063683048351461017c575b5f80fd5b61013461012336600461172a565b60986020525f908152604090205481565b6040519081526020015b60405180910390f35b61015a61015536600461175e565b6102d3565b60405161013e9190611779565b61017a61017536600461175e565b6102e4565b005b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161013e565b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b61017a6102f8565b61017a6101f83660046118c9565b61030b565b60975461020d9063ffffffff1681565b60405163ffffffff909116815260200161013e565b6033546001600160a01b03166101a3565b61017a610241366004611915565b6103d3565b61017a61025436600461175e565b610432565b61017a6102673660046118c9565b610483565b61017a61027a36600461175e565b610494565b61015a6105a1565b61017a6102953660046119b9565b6105b0565b61017a6102a836600461175e565b61085f565b6065546101a3906001600160a01b031681565b61017a6102ce366004611a5a565b6108d5565b60606102de826108e7565b92915050565b6102ec610c00565b6102f581610c5a565b50565b610300610c00565b6103095f610cc3565b565b6040805180820182528281524363ffffffff166020808301919091529151909161033791839101611af6565b60408051601f1981840301815282825280516020918201206097805463ffffffff9081165f908152609890945293909220555416907f58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b84790610399908490611af6565b60405180910390a26097546103b59063ffffffff166001611b41565b6097805463ffffffff191663ffffffff929092169190911790555050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104245760405162461bcd60e51b815260040161041b90611b5e565b60405180910390fd5b61042e8282610d14565b5050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461047a5760405162461bcd60e51b815260040161041b90611b5e565b6102f581610d93565b61048b610c00565b6102f581610e0d565b5f54610100900460ff16158080156104b257505f54600160ff909116105b806104cb5750303b1580156104cb57505f5460ff166001145b61052e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161041b565b5f805460ff19166001179055801561054f575f805461ff0019166101001790555b610559825f610e59565b801561042e575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b60606105ab610ed5565b905090565b60405163ec7fbb3160e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ec7fbb3190602401602060405180830381865afa158015610612573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106369190611bce565b6106825760405162461bcd60e51b815260206004820152601b60248201527f4f70657261746f72206d75737420626520726567697374657265640000000000604482015260640161041b565b63ffffffff83165f908152609860209081526040918290205491516106a991879101611c6e565b60405160208183030381529060405280519060200120146106fb5760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207461736b60a01b604482015260640161041b565b5f6107068580611c80565b604051602001610717929190611cc2565b6040516020818303038152906040528051906020012090505f610786826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018290525f90605c01604051602081830303815290604052805190602001209050919050565b90505f6107c88286868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061100992505050565b90506001600160a01b03811633146108165760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b604482015260640161041b565b8563ffffffff167f6085d9f4a68cf7f70388007e6e9d518bbe73b5d4cc6fbe5862e37f1368e4faa7883360405161084e929190611cf2565b60405180910390a250505050505050565b610867610c00565b6001600160a01b0381166108cc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161041b565b6102f581610cc3565b6108dd61102b565b61042e82826110c5565b60605f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631703a0186040518163ffffffff1660e01b81526004015f60405180830381865afa158015610945573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261096c9190810190611d56565b8051519091505f816001600160401b0381111561098b5761098b6117c5565b6040519080825280602002602001820160405280156109b4578160200160208202803683370190505b5090505f5b82811015610a125783518051829081106109d5576109d5611e4e565b60200260200101515f01518282815181106109f2576109f2611e4e565b6001600160a01b03909216602092830291909101909101526001016109b9565b50604051639004134760e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639004134790610a639089908690600401611e62565b5f60405180830381865afa158015610a7d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610aa49190810190611ebd565b90505f836001600160401b03811115610abf57610abf6117c5565b604051908082528060200260200182016040528015610ae8578160200160208202803683370190505b5090505f805b85811015610b2d575f848281518110610b0957610b09611e4e565b60200260200101511115610b255781610b2181611f48565b9250505b600101610aee565b505f816001600160401b03811115610b4757610b476117c5565b604051908082528060200260200182016040528015610b70578160200160208202803683370190505b5090505f5b86811015610bf3575f858281518110610b9057610b90611e4e565b60200260200101511115610beb57838181518110610bb057610bb0611e4e565b6020026020010151828281518110610bca57610bca611e4e565b60200260200101906001600160a01b031690816001600160a01b0316815250505b600101610b75565b5098975050505050505050565b6033546001600160a01b031633146103095760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041b565b606554604080516001600160a01b03928316815291831660208301527fe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610d629085908590600401611f60565b5f604051808303815f87803b158015610d79575f80fd5b505af1158015610d8b573d5f803e3d5ffd5b505050505050565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b5f604051808303815f87803b158015610df4575f80fd5b505af1158015610e06573d5f803e3d5ffd5b5050505050565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610ddd908490600401611faa565b5f54610100900460ff16610ec35760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840161041b565b610ecc82610cc3565b61042e81610c5a565b60605f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631703a0186040518163ffffffff1660e01b81526004015f60405180830381865afa158015610f33573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610f5a9190810190611d56565b90505f815f0151516001600160401b03811115610f7957610f796117c5565b604051908082528060200260200182016040528015610fa2578160200160208202803683370190505b5090505f5b825151811015611002578251805182908110610fc557610fc5611e4e565b60200260200101515f0151828281518110610fe257610fe2611e4e565b6001600160a01b0390921660209283029190910190910152600101610fa7565b5092915050565b5f805f61101685856113d5565b9150915061102381611440565b509392505050565b6065546001600160a01b031633146103095760405162461bcd60e51b815260206004820152605160248201527f4543445341536572766963654d616e61676572426173652e6f6e6c795265776160448201527f726473496e69746961746f723a2063616c6c6572206973206e6f7420746865206064820152703932bbb0b932399034b734ba34b0ba37b960791b608482015260a40161041b565b5f5b81811015611386578282828181106110e1576110e1611e4e565b90506020028101906110f39190611fbc565b61110490604081019060200161175e565b6001600160a01b03166323b872dd333086868681811061112657611126611e4e565b90506020028101906111389190611fbc565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303815f875af115801561118c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111b09190611bce565b505f8383838181106111c4576111c4611e4e565b90506020028101906111d69190611fbc565b6111e790604081019060200161175e565b604051636eb1769f60e11b81523060048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152919091169063dd62ed3e90604401602060405180830381865afa158015611253573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112779190611fe5565b905083838381811061128b5761128b611e4e565b905060200281019061129d9190611fbc565b6112ae90604081019060200161175e565b6001600160a01b031663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000838787878181106112f0576112f0611e4e565b90506020028101906113029190611fbc565b604001356113109190611ffc565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015611358573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061137c9190611bce565b50506001016110c7565b5060405163fce36c7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063fce36c7d90610d629085908590600401612078565b5f808251604103611409576020830151604084015160608501515f1a6113fd878285856115f5565b94509450505050611439565b825160400361143257602083015160408401516114278683836116da565b935093505050611439565b505f905060025b9250929050565b5f81600481111561145357611453612173565b0361145b5750565b600181600481111561146f5761146f612173565b036114bc5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161041b565b60028160048111156114d0576114d0612173565b0361151d5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161041b565b600381600481111561153157611531612173565b036115895760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161041b565b600481600481111561159d5761159d612173565b036102f55760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161041b565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561162a57505f905060036116d1565b8460ff16601b1415801561164257508460ff16601c14155b1561165257505f905060046116d1565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156116a3573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b0381166116cb575f600192509250506116d1565b91505f90505b94509492505050565b5f806001600160ff1b038316816116f660ff86901c601b611ffc565b9050611704878288856115f5565b935093505050935093915050565b803563ffffffff81168114611725575f80fd5b919050565b5f6020828403121561173a575f80fd5b61174382611712565b9392505050565b6001600160a01b03811681146102f5575f80fd5b5f6020828403121561176e575f80fd5b81356117438161174a565b602080825282518282018190525f9190848201906040850190845b818110156117b95783516001600160a01b031683529284019291840191600101611794565b50909695505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156117fb576117fb6117c5565b60405290565b604051602081016001600160401b03811182821017156117fb576117fb6117c5565b604080519081016001600160401b03811182821017156117fb576117fb6117c5565b604051601f8201601f191681016001600160401b038111828210171561186d5761186d6117c5565b604052919050565b5f6001600160401b0383111561188d5761188d6117c5565b6118a0601f8401601f1916602001611845565b90508281528383830111156118b3575f80fd5b828260208301375f602084830101529392505050565b5f602082840312156118d9575f80fd5b81356001600160401b038111156118ee575f80fd5b8201601f810184136118fe575f80fd5b61190d84823560208401611875565b949350505050565b5f8060408385031215611926575f80fd5b82356119318161174a565b915060208301356001600160401b038082111561194c575f80fd5b908401906060828703121561195f575f80fd5b6119676117d9565b823582811115611975575f80fd5b83019150601f82018713611987575f80fd5b61199687833560208501611875565b815260208301356020820152604083013560408201528093505050509250929050565b5f805f80606085870312156119cc575f80fd5b84356001600160401b03808211156119e2575f80fd5b90860190604082890312156119f5575f80fd5b819550611a0460208801611712565b94506040870135915080821115611a19575f80fd5b818701915087601f830112611a2c575f80fd5b813581811115611a3a575f80fd5b886020828501011115611a4b575f80fd5b95989497505060200194505050565b5f8060208385031215611a6b575f80fd5b82356001600160401b0380821115611a81575f80fd5b818501915085601f830112611a94575f80fd5b813581811115611aa2575f80fd5b8660208260051b8501011115611ab6575f80fd5b60209290920196919550909350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f825160406020840152611b116060840182611ac8565b905063ffffffff60208501511660408401528091505092915050565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff81811683821601908082111561100257611002611b2d565b6020808252604a908201527f4543445341536572766963654d616e61676572426173652e6f6e6c795374616b60408201527f6552656769737472793a2063616c6c6572206973206e6f7420746865207374616060820152696b65526567697374727960b01b608082015260a00190565b5f60208284031215611bde575f80fd5b81518015158114611743575f80fd5b5f8135601e19833603018112611c01575f80fd5b82016020810190356001600160401b03811115611c1c575f80fd5b803603821315611c2a575f80fd5b60408552806040860152808260608701375f6060828701015263ffffffff611c5460208601611712565b166020860152601f01601f19169093016060019392505050565b602081525f6117436020830184611bed565b5f808335601e19843603018112611c95575f80fd5b8301803591506001600160401b03821115611cae575f80fd5b602001915036819003821315611439575f80fd5b74021b7b73334b936903232b837b9b4ba103337b91d1605d1b8152818360158301375f9101601501908152919050565b604081525f611d046040830185611bed565b905060018060a01b03831660208301529392505050565b5f6001600160401b03821115611d3357611d336117c5565b5060051b60200190565b6bffffffffffffffffffffffff811681146102f5575f80fd5b5f6020808385031215611d67575f80fd5b82516001600160401b0380821115611d7d575f80fd5b8185019150828287031215611d90575f80fd5b611d98611801565b825182811115611da6575f80fd5b80840193505086601f840112611dba575f80fd5b82519150611dcf611dca83611d1b565b611845565b82815260069290921b83018401918481019088841115611ded575f80fd5b938501935b83851015611e41576040858a031215611e09575f80fd5b611e11611823565b8551611e1c8161174a565b815285870151611e2b81611d3d565b8188015282526040949094019390850190611df2565b8252509695505050505050565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b0383811682526040602080840182905284519184018290525f92858201929091906060860190855b81811015611eaf578551851683529483019491830191600101611e91565b509098975050505050505050565b5f6020808385031215611ece575f80fd5b82516001600160401b03811115611ee3575f80fd5b8301601f81018513611ef3575f80fd5b8051611f01611dca82611d1b565b81815260059190911b82018301908381019087831115611f1f575f80fd5b928401925b82841015611f3d57835182529284019290840190611f24565b979650505050505050565b5f60018201611f5957611f59611b2d565b5060010190565b60018060a01b0383168152604060208201525f825160606040840152611f8960a0840182611ac8565b90506020840151606084015260408401516080840152809150509392505050565b602081525f6117436020830184611ac8565b5f8235609e19833603018112611fd0575f80fd5b9190910192915050565b80356117258161174a565b5f60208284031215611ff5575f80fd5b5051919050565b808201808211156102de576102de611b2d565b8183525f60208085019450825f5b8581101561206d5781356120308161174a565b6001600160a01b031687528183013561204881611d3d565b6bffffffffffffffffffffffff1687840152604096870196919091019060010161201d565b509495945050505050565b60208082528181018390525f906040808401600586901b8501820187855b88811015611eaf57878303603f190184528135368b9003609e190181126120bb575f80fd5b8a0160a0813536839003601e190181126120d3575f80fd5b820188810190356001600160401b038111156120ed575f80fd5b8060061b36038213156120fe575f80fd5b82875261210e838801828461200f565b9250505061211d888301611fda565b6001600160a01b0316888601528187013587860152606061213f818401611712565b63ffffffff16908601526080612156838201611712565b63ffffffff16950194909452509285019290850190600101612096565b634e487b7160e01b5f52602160045260245ffdfea26469706673582212207a6c995fa8373e9bbabff00398f36ff0c9e49b41e3e2bd3ff602cb2a85fe925e64736f6c63430008190033",
}

// BitDSMServiceManagerProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use BitDSMServiceManagerProxyMetaData.ABI instead.
var BitDSMServiceManagerProxyABI = BitDSMServiceManagerProxyMetaData.ABI

// BitDSMServiceManagerProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BitDSMServiceManagerProxyMetaData.Bin instead.
var BitDSMServiceManagerProxyBin = BitDSMServiceManagerProxyMetaData.Bin

// DeployBitDSMServiceManagerProxy deploys a new Ethereum contract, binding an instance of BitDSMServiceManagerProxy to it.
func DeployBitDSMServiceManagerProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _stakeRegistry common.Address, _rewardsCoordinator common.Address, _delegationManager common.Address) (common.Address, *types.Transaction, *BitDSMServiceManagerProxy, error) {
	parsed, err := BitDSMServiceManagerProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BitDSMServiceManagerProxyBin), backend, _avsDirectory, _stakeRegistry, _rewardsCoordinator, _delegationManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BitDSMServiceManagerProxy{BitDSMServiceManagerProxyCaller: BitDSMServiceManagerProxyCaller{contract: contract}, BitDSMServiceManagerProxyTransactor: BitDSMServiceManagerProxyTransactor{contract: contract}, BitDSMServiceManagerProxyFilterer: BitDSMServiceManagerProxyFilterer{contract: contract}}, nil
}

// BitDSMServiceManagerProxy is an auto generated Go binding around an Ethereum contract.
type BitDSMServiceManagerProxy struct {
	BitDSMServiceManagerProxyCaller     // Read-only binding to the contract
	BitDSMServiceManagerProxyTransactor // Write-only binding to the contract
	BitDSMServiceManagerProxyFilterer   // Log filterer for contract events
}

// BitDSMServiceManagerProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitDSMServiceManagerProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitDSMServiceManagerProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitDSMServiceManagerProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitDSMServiceManagerProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitDSMServiceManagerProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitDSMServiceManagerProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitDSMServiceManagerProxySession struct {
	Contract     *BitDSMServiceManagerProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BitDSMServiceManagerProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitDSMServiceManagerProxyCallerSession struct {
	Contract *BitDSMServiceManagerProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// BitDSMServiceManagerProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitDSMServiceManagerProxyTransactorSession struct {
	Contract     *BitDSMServiceManagerProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// BitDSMServiceManagerProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitDSMServiceManagerProxyRaw struct {
	Contract *BitDSMServiceManagerProxy // Generic contract binding to access the raw methods on
}

// BitDSMServiceManagerProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitDSMServiceManagerProxyCallerRaw struct {
	Contract *BitDSMServiceManagerProxyCaller // Generic read-only contract binding to access the raw methods on
}

// BitDSMServiceManagerProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitDSMServiceManagerProxyTransactorRaw struct {
	Contract *BitDSMServiceManagerProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitDSMServiceManagerProxy creates a new instance of BitDSMServiceManagerProxy, bound to a specific deployed contract.
func NewBitDSMServiceManagerProxy(address common.Address, backend bind.ContractBackend) (*BitDSMServiceManagerProxy, error) {
	contract, err := bindBitDSMServiceManagerProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxy{BitDSMServiceManagerProxyCaller: BitDSMServiceManagerProxyCaller{contract: contract}, BitDSMServiceManagerProxyTransactor: BitDSMServiceManagerProxyTransactor{contract: contract}, BitDSMServiceManagerProxyFilterer: BitDSMServiceManagerProxyFilterer{contract: contract}}, nil
}

// NewBitDSMServiceManagerProxyCaller creates a new read-only instance of BitDSMServiceManagerProxy, bound to a specific deployed contract.
func NewBitDSMServiceManagerProxyCaller(address common.Address, caller bind.ContractCaller) (*BitDSMServiceManagerProxyCaller, error) {
	contract, err := bindBitDSMServiceManagerProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyCaller{contract: contract}, nil
}

// NewBitDSMServiceManagerProxyTransactor creates a new write-only instance of BitDSMServiceManagerProxy, bound to a specific deployed contract.
func NewBitDSMServiceManagerProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BitDSMServiceManagerProxyTransactor, error) {
	contract, err := bindBitDSMServiceManagerProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyTransactor{contract: contract}, nil
}

// NewBitDSMServiceManagerProxyFilterer creates a new log filterer instance of BitDSMServiceManagerProxy, bound to a specific deployed contract.
func NewBitDSMServiceManagerProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BitDSMServiceManagerProxyFilterer, error) {
	contract, err := bindBitDSMServiceManagerProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyFilterer{contract: contract}, nil
}

// bindBitDSMServiceManagerProxy binds a generic wrapper to an already deployed contract.
func bindBitDSMServiceManagerProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BitDSMServiceManagerProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitDSMServiceManagerProxy.Contract.BitDSMServiceManagerProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.BitDSMServiceManagerProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.BitDSMServiceManagerProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitDSMServiceManagerProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.contract.Transact(opts, method, params...)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _BitDSMServiceManagerProxy.Contract.AllTaskHashes(&_BitDSMServiceManagerProxy.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _BitDSMServiceManagerProxy.Contract.AllTaskHashes(&_BitDSMServiceManagerProxy.CallOpts, arg0)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) AvsDirectory() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.AvsDirectory(&_BitDSMServiceManagerProxy.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) AvsDirectory() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.AvsDirectory(&_BitDSMServiceManagerProxy.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "getOperatorRestakedStrategies", _operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.GetOperatorRestakedStrategies(&_BitDSMServiceManagerProxy.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.GetOperatorRestakedStrategies(&_BitDSMServiceManagerProxy.CallOpts, _operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.GetRestakeableStrategies(&_BitDSMServiceManagerProxy.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.GetRestakeableStrategies(&_BitDSMServiceManagerProxy.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) LatestTaskNum() (uint32, error) {
	return _BitDSMServiceManagerProxy.Contract.LatestTaskNum(&_BitDSMServiceManagerProxy.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) LatestTaskNum() (uint32, error) {
	return _BitDSMServiceManagerProxy.Contract.LatestTaskNum(&_BitDSMServiceManagerProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) Owner() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.Owner(&_BitDSMServiceManagerProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) Owner() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.Owner(&_BitDSMServiceManagerProxy.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) RewardsInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "rewardsInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) RewardsInitiator() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.RewardsInitiator(&_BitDSMServiceManagerProxy.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) RewardsInitiator() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.RewardsInitiator(&_BitDSMServiceManagerProxy.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BitDSMServiceManagerProxy.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) StakeRegistry() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.StakeRegistry(&_BitDSMServiceManagerProxy.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyCallerSession) StakeRegistry() (common.Address, error) {
	return _BitDSMServiceManagerProxy.Contract.StakeRegistry(&_BitDSMServiceManagerProxy.CallOpts)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0xead11316.
//
// Solidity: function confirmDeposit((string,uint32) task, uint32 referenceTaskIndex, bytes signature) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) ConfirmDeposit(opts *bind.TransactOpts, task IBitDSMServiceManagerTask, referenceTaskIndex uint32, signature []byte) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "confirmDeposit", task, referenceTaskIndex, signature)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0xead11316.
//
// Solidity: function confirmDeposit((string,uint32) task, uint32 referenceTaskIndex, bytes signature) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) ConfirmDeposit(task IBitDSMServiceManagerTask, referenceTaskIndex uint32, signature []byte) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.ConfirmDeposit(&_BitDSMServiceManagerProxy.TransactOpts, task, referenceTaskIndex, signature)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0xead11316.
//
// Solidity: function confirmDeposit((string,uint32) task, uint32 referenceTaskIndex, bytes signature) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) ConfirmDeposit(task IBitDSMServiceManagerTask, referenceTaskIndex uint32, signature []byte) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.ConfirmDeposit(&_BitDSMServiceManagerProxy.TransactOpts, task, referenceTaskIndex, signature)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.CreateAVSRewardsSubmission(&_BitDSMServiceManagerProxy.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.CreateAVSRewardsSubmission(&_BitDSMServiceManagerProxy.TransactOpts, rewardsSubmissions)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x85edf874.
//
// Solidity: function createNewTask(string name) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) CreateNewTask(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "createNewTask", name)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x85edf874.
//
// Solidity: function createNewTask(string name) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) CreateNewTask(name string) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.CreateNewTask(&_BitDSMServiceManagerProxy.TransactOpts, name)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x85edf874.
//
// Solidity: function createNewTask(string name) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) CreateNewTask(name string) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.CreateNewTask(&_BitDSMServiceManagerProxy.TransactOpts, name)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.DeregisterOperatorFromAVS(&_BitDSMServiceManagerProxy.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.DeregisterOperatorFromAVS(&_BitDSMServiceManagerProxy.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.Initialize(&_BitDSMServiceManagerProxy.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.Initialize(&_BitDSMServiceManagerProxy.TransactOpts, _owner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.RegisterOperatorToAVS(&_BitDSMServiceManagerProxy.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.RegisterOperatorToAVS(&_BitDSMServiceManagerProxy.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.RenounceOwnership(&_BitDSMServiceManagerProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.RenounceOwnership(&_BitDSMServiceManagerProxy.TransactOpts)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "setRewardsInitiator", newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.SetRewardsInitiator(&_BitDSMServiceManagerProxy.TransactOpts, newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.SetRewardsInitiator(&_BitDSMServiceManagerProxy.TransactOpts, newRewardsInitiator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.TransferOwnership(&_BitDSMServiceManagerProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.TransferOwnership(&_BitDSMServiceManagerProxy.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxySession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.UpdateAVSMetadataURI(&_BitDSMServiceManagerProxy.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _BitDSMServiceManagerProxy.Contract.UpdateAVSMetadataURI(&_BitDSMServiceManagerProxy.TransactOpts, _metadataURI)
}

// BitDSMServiceManagerProxyDepositConfirmedIterator is returned from FilterDepositConfirmed and is used to iterate over the raw logs and unpacked data for DepositConfirmed events raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyDepositConfirmedIterator struct {
	Event *BitDSMServiceManagerProxyDepositConfirmed // Event containing the contract specifics and raw log

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
func (it *BitDSMServiceManagerProxyDepositConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitDSMServiceManagerProxyDepositConfirmed)
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
		it.Event = new(BitDSMServiceManagerProxyDepositConfirmed)
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
func (it *BitDSMServiceManagerProxyDepositConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitDSMServiceManagerProxyDepositConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitDSMServiceManagerProxyDepositConfirmed represents a DepositConfirmed event raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyDepositConfirmed struct {
	TaskId   uint32
	Task     IBitDSMServiceManagerTask
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositConfirmed is a free log retrieval operation binding the contract event 0x6085d9f4a68cf7f70388007e6e9d518bbe73b5d4cc6fbe5862e37f1368e4faa7.
//
// Solidity: event DepositConfirmed(uint32 indexed taskId, (string,uint32) task, address operator)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) FilterDepositConfirmed(opts *bind.FilterOpts, taskId []uint32) (*BitDSMServiceManagerProxyDepositConfirmedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _BitDSMServiceManagerProxy.contract.FilterLogs(opts, "DepositConfirmed", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyDepositConfirmedIterator{contract: _BitDSMServiceManagerProxy.contract, event: "DepositConfirmed", logs: logs, sub: sub}, nil
}

// WatchDepositConfirmed is a free log subscription operation binding the contract event 0x6085d9f4a68cf7f70388007e6e9d518bbe73b5d4cc6fbe5862e37f1368e4faa7.
//
// Solidity: event DepositConfirmed(uint32 indexed taskId, (string,uint32) task, address operator)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) WatchDepositConfirmed(opts *bind.WatchOpts, sink chan<- *BitDSMServiceManagerProxyDepositConfirmed, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _BitDSMServiceManagerProxy.contract.WatchLogs(opts, "DepositConfirmed", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitDSMServiceManagerProxyDepositConfirmed)
				if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "DepositConfirmed", log); err != nil {
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

// ParseDepositConfirmed is a log parse operation binding the contract event 0x6085d9f4a68cf7f70388007e6e9d518bbe73b5d4cc6fbe5862e37f1368e4faa7.
//
// Solidity: event DepositConfirmed(uint32 indexed taskId, (string,uint32) task, address operator)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) ParseDepositConfirmed(log types.Log) (*BitDSMServiceManagerProxyDepositConfirmed, error) {
	event := new(BitDSMServiceManagerProxyDepositConfirmed)
	if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "DepositConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BitDSMServiceManagerProxyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyInitializedIterator struct {
	Event *BitDSMServiceManagerProxyInitialized // Event containing the contract specifics and raw log

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
func (it *BitDSMServiceManagerProxyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitDSMServiceManagerProxyInitialized)
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
		it.Event = new(BitDSMServiceManagerProxyInitialized)
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
func (it *BitDSMServiceManagerProxyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitDSMServiceManagerProxyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitDSMServiceManagerProxyInitialized represents a Initialized event raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) FilterInitialized(opts *bind.FilterOpts) (*BitDSMServiceManagerProxyInitializedIterator, error) {

	logs, sub, err := _BitDSMServiceManagerProxy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyInitializedIterator{contract: _BitDSMServiceManagerProxy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BitDSMServiceManagerProxyInitialized) (event.Subscription, error) {

	logs, sub, err := _BitDSMServiceManagerProxy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitDSMServiceManagerProxyInitialized)
				if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) ParseInitialized(log types.Log) (*BitDSMServiceManagerProxyInitialized, error) {
	event := new(BitDSMServiceManagerProxyInitialized)
	if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BitDSMServiceManagerProxyNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyNewTaskCreatedIterator struct {
	Event *BitDSMServiceManagerProxyNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *BitDSMServiceManagerProxyNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitDSMServiceManagerProxyNewTaskCreated)
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
		it.Event = new(BitDSMServiceManagerProxyNewTaskCreated)
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
func (it *BitDSMServiceManagerProxyNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitDSMServiceManagerProxyNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitDSMServiceManagerProxyNewTaskCreated represents a NewTaskCreated event raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyNewTaskCreated struct {
	TaskId uint32
	Task   IBitDSMServiceManagerTask
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b847.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (string,uint32) task)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskId []uint32) (*BitDSMServiceManagerProxyNewTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _BitDSMServiceManagerProxy.contract.FilterLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyNewTaskCreatedIterator{contract: _BitDSMServiceManagerProxy.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b847.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (string,uint32) task)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *BitDSMServiceManagerProxyNewTaskCreated, taskId []uint32) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _BitDSMServiceManagerProxy.contract.WatchLogs(opts, "NewTaskCreated", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitDSMServiceManagerProxyNewTaskCreated)
				if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b847.
//
// Solidity: event NewTaskCreated(uint32 indexed taskId, (string,uint32) task)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) ParseNewTaskCreated(log types.Log) (*BitDSMServiceManagerProxyNewTaskCreated, error) {
	event := new(BitDSMServiceManagerProxyNewTaskCreated)
	if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BitDSMServiceManagerProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyOwnershipTransferredIterator struct {
	Event *BitDSMServiceManagerProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BitDSMServiceManagerProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitDSMServiceManagerProxyOwnershipTransferred)
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
		it.Event = new(BitDSMServiceManagerProxyOwnershipTransferred)
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
func (it *BitDSMServiceManagerProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitDSMServiceManagerProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitDSMServiceManagerProxyOwnershipTransferred represents a OwnershipTransferred event raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BitDSMServiceManagerProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BitDSMServiceManagerProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyOwnershipTransferredIterator{contract: _BitDSMServiceManagerProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BitDSMServiceManagerProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BitDSMServiceManagerProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitDSMServiceManagerProxyOwnershipTransferred)
				if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) ParseOwnershipTransferred(log types.Log) (*BitDSMServiceManagerProxyOwnershipTransferred, error) {
	event := new(BitDSMServiceManagerProxyOwnershipTransferred)
	if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator is returned from FilterRewardsInitiatorUpdated and is used to iterate over the raw logs and unpacked data for RewardsInitiatorUpdated events raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator struct {
	Event *BitDSMServiceManagerProxyRewardsInitiatorUpdated // Event containing the contract specifics and raw log

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
func (it *BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitDSMServiceManagerProxyRewardsInitiatorUpdated)
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
		it.Event = new(BitDSMServiceManagerProxyRewardsInitiatorUpdated)
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
func (it *BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitDSMServiceManagerProxyRewardsInitiatorUpdated represents a RewardsInitiatorUpdated event raised by the BitDSMServiceManagerProxy contract.
type BitDSMServiceManagerProxyRewardsInitiatorUpdated struct {
	PrevRewardsInitiator common.Address
	NewRewardsInitiator  common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsInitiatorUpdated is a free log retrieval operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator, error) {

	logs, sub, err := _BitDSMServiceManagerProxy.contract.FilterLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return &BitDSMServiceManagerProxyRewardsInitiatorUpdatedIterator{contract: _BitDSMServiceManagerProxy.contract, event: "RewardsInitiatorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsInitiatorUpdated is a free log subscription operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *BitDSMServiceManagerProxyRewardsInitiatorUpdated) (event.Subscription, error) {

	logs, sub, err := _BitDSMServiceManagerProxy.contract.WatchLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitDSMServiceManagerProxyRewardsInitiatorUpdated)
				if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
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

// ParseRewardsInitiatorUpdated is a log parse operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_BitDSMServiceManagerProxy *BitDSMServiceManagerProxyFilterer) ParseRewardsInitiatorUpdated(log types.Log) (*BitDSMServiceManagerProxyRewardsInitiatorUpdated, error) {
	event := new(BitDSMServiceManagerProxyRewardsInitiatorUpdated)
	if err := _BitDSMServiceManagerProxy.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
