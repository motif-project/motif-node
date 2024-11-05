package ethComms

import (
	"bytes"
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
	"golang.org/x/crypto/sha3"
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

	fmt.Println(podAddress)
	fmt.Println(oprAddr)
	fmt.Println(btcTxId)
	fmt.Println(amount.String())

	instance, privateKey, auth, err := initializeServiceManagerContract()

	btcTxIdBytes, err := hex.DecodeString(btcTxId)
	if err != nil {
		fmt.Println("failed to decode BTC tx ID: ", err)
		return "", err
	}

	bBytes := amount.Bytes()
	paddedAmount := make([]byte, 32)
	copy(paddedAmount[32-len(bBytes):], bBytes)

	hash := crypto.Keccak256Hash(
		common.HexToAddress(podAddress).Bytes(),
		common.HexToAddress(oprAddr).Bytes(),
		paddedAmount,
		btcTxIdBytes,
		[]byte{1}, // true is represented as 1 in byte form
	)

	hash = hashWithEthereumPrefix(hash)

	signature, err := signMessage(hash, privateKey)
	if err != nil {
		fmt.Println("failed to sign message: ", err)
		return "", err
	}

	// Call the confirmDeposit function
	tx, err := instance.ConfirmDeposit(auth, common.HexToAddress(podAddress), signature)
	if err != nil {
		fmt.Println("failed to call confirmDeposit:", err)
		return "", err
	}

	fmt.Println("Transaction deposit submitted: ", tx.Hash().Hex())
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

func CallWithdrawBitcoinPSBT(podAddress string, withdrawAddress string, psbt string, amount big.Int) (string, error) {

	instance, privateKey, auth, err := initializeServiceManagerContract()

	bBytes := amount.Bytes()
	paddedAmount := make([]byte, 32)
	copy(paddedAmount[32-len(bBytes):], bBytes)

	// psbtBytes := base58.Decode(psbt)
	// if len(psbtBytes) == 0 {
	// 	fmt.Println("failed to decode PSBT")
	// 	return "", fmt.Errorf("failed to decode PSBT")
	// }

	psbtBytes := []byte(psbt)

	withdrawAddressBytes, err := hex.DecodeString(withdrawAddress)
	if err != nil {
		fmt.Println("failed to decode withdraw address")
		return "", err
	}

	hash := crypto.Keccak256Hash(
		common.HexToAddress(podAddress).Bytes(),
		paddedAmount,
		psbtBytes,
		withdrawAddressBytes,
	)

	hash = hashWithEthereumPrefix(hash)

	signature, err := signMessage(hash, privateKey)
	if err != nil {
		fmt.Println("failed to sign message: ", err)
		return "", err
	}

	// Call the confirmDeposit function
	tx, err := instance.WithdrawBitcoinPSBT(auth, common.HexToAddress(podAddress), &amount, psbtBytes, signature)
	if err != nil {
		fmt.Println("failed to call Withdraw psbt:", err)
		return "", err
	}

	fmt.Println("Transaction withdraw submitted: ", tx.Hash().Hex())
	return tx.Hash().Hex(), nil

}

func initializeServiceManagerContract() (*BitDSMServiceManager.BitDSMServiceManager, *ecdsa.PrivateKey, *bind.TransactOpts, error) {
	ethAccountOpr := LoadEthAccount()
	client, err := GetEthClient()
	if err != nil {
		fmt.Println("Failed to get eth client: ", err)
		return nil, nil, nil, err
	}

	contractAddress := common.HexToAddress(viper.GetString("service_manager_address"))
	instance, err := BitDSMServiceManager.NewBitDSMServiceManager(contractAddress, client)
	if err != nil {
		fmt.Println("Failed to create contract instance: %v", err)
		return nil, nil, nil, err
	}

	privateKey, err := GetPrivateKeyFromKeyStore(ethAccountOpr, viper.GetString("eth_keystore_password"))
	if err != nil {
		fmt.Println("Failed to get private key: ", err)
		return nil, nil, nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), ethAccountOpr.Address)
	if err != nil {
		fmt.Println("Failed to get nonce: %v", err)
		return nil, nil, nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	// Get the chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("failed to get chain ID: ", err)
		return nil, nil, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println("failed to create transactor: ", err)
		return nil, nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return instance, privateKey, auth, nil
}

func signMessage(hash common.Hash, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		fmt.Println("failed to sign hash: ", err)
		return nil, err
	}

	if signature[64] < 27 {
		signature[64] += 27
	}
	return signature, nil
}

func hashWithEthereumPrefix(hash common.Hash) common.Hash {
	prefix := []byte("\x19Ethereum Signed Message:\n32")
	fullMessage := append(prefix, hash.Bytes()...)

	hashWithPrefix := sha3.NewLegacyKeccak256()
	hashWithPrefix.Write(fullMessage)
	return common.BytesToHash(hashWithPrefix.Sum(nil))
}

func concatenateBytes(byteSlices ...[]byte) []byte {
	var buffer bytes.Buffer

	for _, b := range byteSlices {
		buffer.Write(b)
	}

	return buffer.Bytes()
}

// 6b3925c6cac45b31192ed4fd9b05ac46e043fe5860b3b41240fb9d353ace1e37b0ce79054154ea4000000000000000000000000000000000000000000000000000000000000f42406437346235643233633332663231316134316462303765366664386562613863663165396333646639626537613964653137326262663532613766623861343901
// 6b3925c6cac45b31192ed4fd9b05ac46e043fe5860b3b41240fb9d353ace1e37b0ce79054154ea4000000000000000000000000000000000000000000000000000000000000f4240d74b5d23c32f211a41db07e6fd8eba8cf1e9c3df9be7a9de172bbf52a7fb8a4901
