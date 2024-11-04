package comms

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/AhmadAshraf2/Judge-AVS/BitDSMServiceManager"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func GetEthClient() (*ethclient.Client, error) {
	ethUrl := viper.GetString("eth_rpc_host")
	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return client, nil
}

func getCurrentBlockNumber() (uint64, error) {
	client, err := GetEthClient()
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

func GetPrivateKeyFromKeyStore(account accounts.Account, passphrase string) (*ecdsa.PrivateKey, error) {
	// Load the keystore
	keystorePath := viper.GetString("eth_keystore_dir")
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
			keyJson, err := ks.Export(account, passphrase, passphrase)
			if err != nil {
				return nil, fmt.Errorf("failed to export private key: %v", err)
			}

			key, err := keystore.DecryptKey(keyJson, passphrase)
			if err != nil {
				return nil, fmt.Errorf("failed to decrypt key: %v", err)
			}

			return key.PrivateKey, nil
		}
	}

	return nil, fmt.Errorf("account not found in keystore")
}

func CallConfirmBtcDeposit(podAddress string, oprAddr string, btcTxId string, amount big.Int) (string, error) {

	ethAccountOpr := LoadEthAccount()
	client, err := GetEthClient()
	if err != nil {
		fmt.Println("Failed to get eth client: ", err)
		return "", err
	}

	contractAddress := common.HexToAddress(viper.GetString("service_manager_address"))
	instance, err := BitDSMServiceManager.NewBitDSMServiceManager(contractAddress, client)
	if err != nil {
		fmt.Println("Failed to create contract instance: %v", err)
		return "", err
	}

	privateKey, err := GetPrivateKeyFromKeyStore(ethAccountOpr, viper.GetString("eth_keystore_password"))
	if err != nil {
		fmt.Println("Failed to get private key: ", err)
		return "", err
	}

	nonce, err := client.PendingNonceAt(context.Background(), ethAccountOpr.Address)
	if err != nil {
		fmt.Println("Failed to get nonce: %v", err)
		return "", err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Get the chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("failed to get chain ID: ", err)
		return "", err
	}

	// Create the auth object
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println("failed to create transactor: ", err)
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	btcTxIdBytes, err := hex.DecodeString(btcTxId)
	if err != nil {
		fmt.Println("failed to decode BTC tx ID: ", err)
		return "", err
	}

	hash := crypto.Keccak256Hash(
		common.HexToAddress(podAddress).Bytes(),
		common.HexToAddress(oprAddr).Bytes(),
		amount.Bytes(),
		btcTxIdBytes,
		[]byte{1}, // true is represented as 1 in byte form
	)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		fmt.Println("failed to sign hash: ", err)
		return "", err
	}

	// Call the confirmDeposit function
	tx, err := instance.ConfirmDeposit(auth, common.HexToAddress(podAddress), signature)
	if err != nil {
		fmt.Println("failed to call confirmDeposit:", err)
		return "", err
	}

	log.Printf("Transaction submitted: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func LoadEthAccount() accounts.Account {
	keystoreDir := viper.GetString("eth_keystore_dir")
	passphrase := viper.GetString("eth_keystore_passphrase")
	var ethAccount accounts.Account

	// Check if the keystore directory exists
	if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
		// Create a new keystore
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
		// Generate a new account with a passphrase
		ethAccount, err = ks.NewAccount(passphrase)
		if err != nil {
			log.Fatalf("Failed to create new account: %v", err)
		}
	} else {
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
		accounts := ks.Accounts()
		if len(accounts) == 0 {
			// Generate a new account with a passphrase if no accounts exist
			ethAccount, err = ks.NewAccount(passphrase)
			if err != nil {
				log.Fatalf("Failed to create new account: %v", err)
			}
		} else {
			ethAccount = accounts[0]
		}
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
