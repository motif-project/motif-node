package comms

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/AhmadAshraf2/Judge-AVS/BitDSMServiceManagerProxy"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func getEthClient() (*ethclient.Client, error) {
	ethUrl := viper.GetString("eth_rpc_host")
	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return client, nil
}

func getCurrentBlockNumber() (uint64, error) {
	client, err := getEthClient()
	if err != nil {
		fmt.Println("Failed to get eth client: %v", err)
		return 0, err
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to get block number: %v", err)
		return 0, err
	}

	return blockNumber, nil
}

func getPrivateKeyFromKeyStore(account accounts.Account, passphrase string) (*ecdsa.PrivateKey, error) {
	// Load the keystore
	keystorePath := viper.GetString("keystore_path")
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	// Find the account in the keystore
	acc := ks.Accounts()
	for _, acc := range acc {
		if acc.Address == account.Address {
			// Unlock the account
			err := ks.Unlock(account, passphrase)
			if err != nil {
				return nil, fmt.Errorf("failed to unlock account: %v", err)
			}

			// Get the private key
			key, err := ks.Export(account, passphrase, passphrase)
			if err != nil {
				return nil, fmt.Errorf("failed to export private key: %v", err)
			}

			privateKey, err := crypto.ToECDSA(key)
			if err != nil {
				return nil, fmt.Errorf("failed to convert key to ECDSA: %v", err)
			}

			return privateKey, nil
		}
	}

	return nil, fmt.Errorf("account not found in keystore")
}

func CallConfirmBtcDeposit(depositorAddr string) (string, error) {

	ethAccountOpr := loadEthAccount()
	client, err := getEthClient()
	if err != nil {
		fmt.Println("Failed to get eth client: ", err)
		return "", err
	}

	contractAddress := common.HexToAddress(viper.GetString("service_manager_contract_address"))
	instance, err := BitDSMServiceManagerProxy.NewBitDSMServiceManagerProxy(contractAddress, client)
	if err != nil {
		fmt.Println("Failed to create contract instance: %v", err)
		return "", err
	}

	nonce, err := client.PendingNonceAt(context.Background(), ethAccountOpr.Address)
	if err != nil {
		fmt.Println("Failed to get nonce: %v", err)
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Failed to get gas price: %v", err)
		return "", err
	}

	privateKey, err := getPrivateKeyFromKeyStore(ethAccountOpr, viper.GetString("eth_keystore_password"))
	if err != nil {
		fmt.Println("Failed to get private key: ", err)
		return "", err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("Failed to get chain ID: %v", err)
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println("Failed to create transactor: %v", err)
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	currentBlock, err := getCurrentBlockNumber()
	if err != nil {
		fmt.Println("Failed to get current block number: ", err)
		return "", err
	}

	// Prepare the function parameters
	task := BitDSMServiceManagerProxy.IBitDSMServiceManagerTask{
		Name:             "Deposit Confirmed",
		TaskCreatedBlock: uint32(currentBlock),
	}
	referenceTaskIndex := uint32(0)
	signature := []byte("YOUR_SIGNATURE")

	// Call the confirmDeposit function
	tx, err := instance.ConfirmDeposit(auth, task, referenceTaskIndex, signature)
	if err != nil {
		log.Fatalf("Failed to call confirmDeposit: %v", err)
	}

	fmt.Println("Transaction submitted: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func loadEthAccount() accounts.Account {
	keystoreDir := viper.GetString("eth_keystore_dir")
	var ethAccount accounts.Account
	// Check if the keystore directory exists
	if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
		ethAccount = generateEthKeyPair()
	} else {
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
		accounts := ks.Accounts()
		ethAccount = accounts[0]
	}
	return ethAccount
}

func generateEthKeyPair() accounts.Account {
	// Generate a new random private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	password := viper.GetString("eth_keystore_password")
	ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		log.Fatalf("Failed to import private key: %v", err)
	}

	return account
}
