package address

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AhmadAshraf2/Judge-AVS/comms"
	"github.com/AhmadAshraf2/Judge-AVS/db"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
	"github.com/btcsuite/btcd/wire"
	"github.com/spf13/viper"
)

func GenerateAndRegisterNewBtcMultiSigAddress(depositorPubKey string, depositorEthAddress string) (string, error) {
	newAddress, err := GenerateSimpleMultisigAddress(depositorPubKey, depositorEthAddress)
	if err != nil {
		return "", err
	}
	registerAddressOnForkscanner(newAddress)
	return newAddress, nil
}

func GenerateSimpleMultisigAddress(depositorPubKey string, depositorEthAddress string) (string, error) {
	dbconn := db.InitDB()
	wallet := viper.GetString("wallet_name")

	descriptor, err := buildSimpleMultisigDescriptor(depositorPubKey)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := comms.GetDescriptorInfo(descriptor, wallet)
	if err != nil {
		fmt.Println("error in getting descriptorinfo : ", err)
		return "", err
	}

	fmt.Println("Descriptor : ", resp.Descriptor)

	err = comms.ImportDescriptor(resp.Descriptor, wallet)
	if err != nil {
		fmt.Println("error in importing descriptor : ", err)
		return "", err
	}

	address, err := comms.DeriveAddress(wallet, resp.Descriptor)
	if err != nil {
		fmt.Println("error in getting address : ", err)
	}

	addressInfo, err := comms.GetAddressInfo(address, wallet)
	if err != nil {
		fmt.Println("Error getting address info : ", err)
		return "", err
	}
	// Decode Hex string to bytes

	db.InsertMultiSigAddress(dbconn, address, addressInfo.Hex, depositorEthAddress)
	dbconn.Close()
	return address, nil
}

func buildSimpleMultisigDescriptor(depositorPubKey string) (string, error) {
	required := 3
	var signerKey string
	judgePubKey := viper.GetString("btc_xpublic_key")
	signerPubKey := viper.GetStringSlice("signer_xpublic_key")
	signerKey += fmt.Sprintf(",%s", signerPubKey[0])
	descriptorScript := fmt.Sprintf("wsh(multi(%d,%s%s,%s))", required, judgePubKey, signerKey, depositorPubKey)
	fmt.Println(descriptorScript)
	return descriptorScript, nil
}

func registerAddressOnForkscanner(address string) {
	dt := time.Now().UTC()
	dt = dt.AddDate(1, 0, 0)

	request_body := map[string]interface{}{
		"method":  "add_watched_addresses",
		"id":      1,
		"jsonrpc": "2.0",
		"params": map[string]interface{}{
			"add": []interface{}{
				map[string]string{
					"address":     address,
					"watch_until": dt.Format(time.RFC3339),
				},
			},
		},
	}

	data, err := json.Marshal(request_body)
	if err != nil {
		log.Fatalf("Post: %v", err)
	}
	fmt.Println(string(data))

	forkscannerHost := viper.GetString("forkscanner_host")
	forkscannerPort := viper.GetString("forkscanner_rpc_port")
	url := fmt.Sprintf("http://%s:%s", forkscannerHost, forkscannerPort)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println(result)

	fmt.Println("registered address on forkscanner : ", address)

}

func GenerateMultisigwithdrawTx(withdrawBTCAddress string, ethClientAddr string) (string, error) {
	dbconn := db.InitDB()
	wallet := viper.GetString("wallet_name")
	multiSigAddresses := db.QueryMultisigAddressByEthAddress(dbconn, ethClientAddr)
	if len(multiSigAddresses) <= 0 {
		fmt.Println("no multisig address found")
		return "", errors.New("no multisig found")
	}
	multiSigAddress := multiSigAddresses[0]

	utxos := db.QueryUtxo(dbconn, multiSigAddress.Address)

	if len(utxos) <= 0 {
		// need to decide if this needs to be enabled
		// addr := generateAndRegisterNewAddress(accountName, height+noOfMultisigs, sweepAddress.Address)
		fmt.Println("INFO : No funds in address : ", multiSigAddress.Address)
		return "", errors.New("no BTC UTXO found")
	}

	var inputs []comms.TxInput
	var outputs []comms.TxOutput
	totalAmountTxIn := uint64(0)

	for _, u := range utxos { //ideally the height should be masked with 0x0000ffff
		inputs = append(inputs, comms.TxInput{Txid: u.Txid, Vout: int64(u.Vout), Sequence: int64(wire.MaxTxInSequenceNum - 10)})
		totalAmountTxIn += u.Amount
	}

	outputs = append(outputs, comms.TxOutput{withdrawBTCAddress: float64(totalAmountTxIn)})

	hexTx, err := comms.CreateRawTx(inputs, outputs, 0, wallet)
	if err != nil {
		fmt.Println("error in creating raw tx : ", err)
		return "", err
	}

	multisigTx, err := utils.CreateTxFromHex(hexTx)
	if err != nil {
		fmt.Println("error decoding tx : ", err)
		return "", err
	}

	fee, err := utils.GetFeeFromBtcNode(multisigTx)
	if err != nil {
		fmt.Println("error in getting fee : ", err)
		return "", err
	}
	outputs = []comms.TxOutput{comms.TxOutput{withdrawBTCAddress: float64(totalAmountTxIn - uint64(fee))}}

	p, err := comms.CreatePsbt(inputs, outputs, 0, wallet)
	if err != nil {
		fmt.Println("error in creating psbt : ", err)
		return "", err
	}

	fmt.Println("transaction base64 psbt: ", p)
	dbconn.Close()
	return p, nil
}

func SignMultisigPSBT(psbt string) (string, error) {
	_, psbt, err := comms.SignPsbt(psbt, "rbf", false)
	if err != nil {
		fmt.Println("error in signing psbt : ", err)
		return "", err
	}
	_, psbt, err = comms.SignPsbt(psbt, "s1", true)
	if err != nil {
		fmt.Println("error in signing psbt : ", err)
		return "", err
	}
	return psbt, nil
}

func GenerateBTCAddress(publicKey string, ethAddr string) (string, error) {
	newAddress, err := GenerateAndRegisterNewBtcMultiSigAddress(publicKey, ethAddr)
	if err != nil {
		return "", err
	}
	return newAddress, nil
}

func SubmitSignedPsbt(psbt string) (string, error) {
	fmt.Println("inside submit signed psbt")
	p, err := SignMultisigPSBT(psbt)
	if err != nil {
		return "", err
	}
	return p, nil
}

func GetUnsignedPsbt(withdrawBtcAddress string, ethAddr string) (string, error) {
	// Here you can add your logic to get the unsigned PSBT
	// For now, it just returns an empty string
	psbt, err := GenerateMultisigwithdrawTx(withdrawBtcAddress, ethAddr)
	if err != nil {
		return "", err
	}
	return psbt, nil
}
