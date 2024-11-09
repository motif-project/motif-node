package operator

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/BitDSM/BitDSM-Node/AvsDirectory"
	"github.com/BitDSM/BitDSM-Node/BitdsmRegistry"
	"github.com/BitDSM/BitDSM-Node/DelegationManager"
	"github.com/BitDSM/BitDSM-Node/ethComms"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

func RegisterOperator() {
	client, err := ethComms.GetEthClient()
	if err != nil {
		fmt.Println("Failed to get eth client: ", err)
	}

	ethAccountOpr := ethComms.LoadEthAccount()
	privateKey, err := ethComms.GetPrivateKeyFromKeyStore(ethAccountOpr, viper.GetString("eth_keystore_password"))
	if err != nil {
		fmt.Println("Failed to get private key: ", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("Failed to get chain ID: ", err)
	}

	// Create an authenticated transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: ", err)
	}

	operatorDetails := DelegationManager.IDelegationManagerOperatorDetails{
		DeprecatedEarningsReceiver: ethAccountOpr.Address, // or specify an address to receive earnings
		DelegationApprover:         ethAccountOpr.Address, // or specify an approver address
		StakerOptOutWindowBlocks:   0,                     // or specify your preferred window
	}

	delegationManagerAddr := common.HexToAddress(viper.GetString("eigen_delegation_manager_address"))

	delegationManager, err := DelegationManager.NewDelegationManager(delegationManagerAddr, client)
	if err != nil {
		fmt.Println("failed to initialize delegation manager: ", err)
	}
	// Check if operator is already registered in EigenLayer
	fmt.Println("auth.From: ", auth.From)
	registeredOperator, err := delegationManager.IsOperator(&bind.CallOpts{}, auth.From)
	if err != nil {
		fmt.Println("Failed to check operator registration: ", err)
	}

	if !registeredOperator {
		fmt.Println("Registering Operator to EigenLayer")
		tx, err := delegationManager.RegisterAsOperator(
			auth,
			operatorDetails,                     // operator address
			viper.GetString("opr_metadata_uri"), // metadata URI
		)
		if err != nil {
			fmt.Println("Error in registering as operator: ", err)
			panic(err)
		}
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			fmt.Println("Failed to wait for transaction receipt: ", err)
			panic(err)
		}
		if receipt.Status == 1 {
			fmt.Println("Operator registered to Core EigenLayer contracts")
		} else {
			fmt.Println("Transaction failed: ", receipt)
			panic("Operator registration failed")
		}
	}

	// registering with AVS
	bitdsmStakeRegistryAddr := common.HexToAddress(viper.GetString("bitdsm_registry_address"))
	bitdsmRegistry, err := BitdsmRegistry.NewBitdsmRegistry(bitdsmStakeRegistryAddr, client)
	if err != nil {
		fmt.Println("failed to initialize AVS registry contract: ", err)
	}

	// Check if operator is already registered in AVS
	registered, err := bitdsmRegistry.OperatorRegistered(&bind.CallOpts{}, auth.From)
	if err != nil {
		fmt.Println("failed to check operator registration in AVS: ", err)
	}

	if registered {
		fmt.Println("Operator already registered with AVS")
		return
	}

	// Generate salt and expiry for registration
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		fmt.Println("failed to generate salt: ", err)
	}

	// Set expiry to 24 hours from now
	expiry := big.NewInt(time.Now().Add(24 * time.Hour).Unix())

	var saltArray [32]byte
	copy(saltArray[:], salt)

	avsDirectoryAddress := common.HexToAddress(viper.GetString("eigen_avs_directory_address"))
	avsDirectory, err := AvsDirectory.NewAvsDirectory(avsDirectoryAddress, client)
	if err != nil {
		fmt.Println("failed to initialize AVS directory contract: ", err)
	}

	serviceManagerAddr := common.HexToAddress(viper.GetString("service_manager_address"))
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{},
		auth.From, // operator address
		serviceManagerAddr,
		saltArray, // random salt
		expiry,    // expiry timestamp
	)
	if err != nil {
		fmt.Println("failed to calculate digest hash: ", err)
	}
	signature, err := crypto.Sign(digestHash[:], privateKey)
	if err != nil {
		fmt.Println("failed to sign digest hash: ", err)
	}

	fmt.Println("signature: ", signature)
	if signature[64] < 27 {
		signature[64] += 27
	}

	operatorSignature := BitdsmRegistry.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature, // Your signature bytes
		Salt:      saltArray,
		Expiry:    expiry,
	}

	// Register operator
	btcPubkey, _ := hex.DecodeString("02ce985f738ae8cf92969ab31a06f74f5e977e335bf226b01231c2e10b55db08dd")
	tx, err := bitdsmRegistry.RegisterOperatorWithSignature(
		auth,
		operatorSignature,
		auth.From,
		btcPubkey,
	)
	if err != nil {
		fmt.Println("failed to register operator: ", err)
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction receipt: ", err)
	}

	if receipt.Status != 1 {
		fmt.Println("registration transaction failed: ", receipt)
	}

	fmt.Println("Successfully registered operator with AVS")

}
