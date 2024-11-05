package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/AhmadAshraf2/Judge-AVS/btcComms"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/psbt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func InitConfigFile() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file: ", err)
	}
}

func getBitcoinRpcClient() *rpcclient.Client {
	connCfg := &rpcclient.ConnConfig{
		Host:         viper.GetString("btc_node_host"),
		User:         viper.GetString("btc_node_user"),
		Pass:         viper.GetString("btc_node_pass"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		fmt.Println("Failed to connect to the Bitcoin client : ", err)
	}

	return client
}

func LoadBtcWallet(walletName string) {
	client := getBitcoinRpcClient()
	defer client.Shutdown()
	_, err := client.LoadWallet(walletName)
	if err != nil {
		fmt.Println("Failed to load wallet : ", err)
	}
}

func CreateTxFromHex(txHex string) (*wire.MsgTx, error) {
	// Decode the transaction hex string
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex string: %v", err)
	}

	// Create a new transaction object
	tx := wire.NewMsgTx(wire.TxVersion)

	// Deserialize the transaction bytes
	err = tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize transaction: %v", err)
	}

	return tx, nil
}

func GetRawTransaction(txid string) (*btcjson.TxRawResult, error) {

	client := getBitcoinRpcClient()
	defer client.Shutdown()

	// Convert txid to chainhash.Hash
	txHash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		fmt.Println("Failed to create tx hash: ", err)
		return nil, err
	}
	tx, err := client.GetRawTransactionVerbose(txHash)
	if err != nil {
		fmt.Println("Failed to get raw transaction: ", err)
		return nil, err
	}
	return tx, nil
}

func GetFeeFromBtcNode(tx *wire.MsgTx) (int64, error) {
	walletName := viper.GetString("wallet_name")
	feeRateAdjustment := viper.GetInt64("fee_rate_adjustment")
	result, err := btcComms.GetEstimateFee(walletName)
	if err != nil {
		fmt.Println("Error getting fee rate : ", err)
		return 0, err
	}

	feeRateInBtc := result.Result.Feerate

	fmt.Printf("Estimated fee per kilobyte for a transaction to be confirmed within 2 blocks: %f BTC\n", feeRateInBtc)
	feeRate := BtcToSats(feeRateInBtc) + feeRateAdjustment
	fmt.Printf("Estimated fee per kilobyte for a transaction to be confirmed within 2 blocks: %d Sats\n", feeRate)
	baseSize := tx.SerializeSizeStripped()
	totalSize := tx.SerializeSize()
	weight := (baseSize * 3) + totalSize
	vsize := (weight + 3) / 4
	fmt.Println("tx size in bytes : ", vsize)
	fee := float64(vsize) * float64(feeRate) / float64(1024)
	fmt.Println("fee for this tx : ", fee)
	return int64(fee), nil
}

func BtcToSats(btc float64) int64 {
	return int64(btc * 1e8)
}

func SatsToBtc(sats int64) float64 {
	return float64(sats) / 100000000.0
}

func IsValidBtcPubKey(pubKeyStr string) bool {
	// Decode the hex string into bytes
	pubKeyBytes, err := hex.DecodeString(pubKeyStr)
	if err != nil {
		return false
	}

	// Parse the public key using btcec
	_, err = btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	return err == nil
}

func IsValidEthAddress(address string) bool {
	return common.IsHexAddress(address)
}

func IsValidPsbt(psbtStr string) bool {
	// Decode the base58 string into bytes
	psbtBytes := base58.Decode(psbtStr)
	if len(psbtBytes) == 0 {
		return false
	}

	// Convert psbtBytes to an io.Reader
	psbtReader := bytes.NewReader(psbtBytes)

	// Parse the PSBT bytes
	_, err := psbt.NewFromRawBytes(psbtReader, false)
	return err == nil
}

func IsValidBtcAddress(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err == nil {
		return true
	}

	_, err = btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
	if err == nil {
		return true
	}

	_, err = btcutil.DecodeAddress(address, &chaincfg.SigNetParams)
	if err == nil {
		return true
	}

	_, err = btcutil.DecodeAddress(address, &chaincfg.RegressionNetParams)
	if err == nil {
		return true
	}
	return false
}
func ListUnspentBtcUtxos(address string) ([]btcjson.ListUnspentResult, error) {
	client := getBitcoinRpcClient()
	defer client.Shutdown()

	addr, err := btcutil.DecodeAddress(address, &chaincfg.RegressionNetParams)
	if err != nil {
		log.Fatalf("Invalid Bitcoin address: %v", err)
	}

	unspent, err := client.ListUnspentMinMaxAddresses(3, 9999999, []btcutil.Address{addr})
	if err != nil {
		fmt.Println("Error listing unspent outputs: ", err)
		return nil, err
	}
	return unspent, nil
}

func Bech32ToHex(bech32Addr string) (string, error) {
	// Decode the bech32 address
	addr, err := btcutil.DecodeAddress(bech32Addr, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("failed to decode bech32 address: %v", err)
	}

	// Handle different witness address types
	var witnessProgram []byte
	switch a := addr.(type) {
	case *btcutil.AddressWitnessPubKeyHash:
		witnessProgram = a.WitnessProgram()
	case *btcutil.AddressWitnessScriptHash:
		witnessProgram = a.WitnessProgram()
	default:
		return "", fmt.Errorf("unsupported address type: %T", addr)
	}

	// Convert to hex
	hexAddr := hex.EncodeToString(witnessProgram)

	return hexAddr, nil
}

func HexToBech32(hexAddr string, network *chaincfg.Params) (string, error) {
	// Decode hex string to bytes
	decoded, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex string: %v", err)
	}

	// Create appropriate address based on the length of the decoded bytes
	var addr btcutil.Address
	switch len(decoded) {
	case 20: // P2WPKH (20 bytes)
		addr, err = btcutil.NewAddressWitnessPubKeyHash(decoded, network)
	case 32: // P2WSH (32 bytes)
		addr, err = btcutil.NewAddressWitnessScriptHash(decoded, network)
	default:
		return "", fmt.Errorf("invalid decoded length: got %d, want 20 or 32", len(decoded))
	}

	if err != nil {
		return "", fmt.Errorf("failed to create bech32 address: %v", err)
	}

	return addr.String(), nil
}
