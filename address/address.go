package address

import (
	"errors"
	"fmt"

	"github.com/AhmadAshraf2/BitDSM-Node/btcComms"
	"github.com/AhmadAshraf2/BitDSM-Node/db"
	"github.com/AhmadAshraf2/BitDSM-Node/utils"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/spf13/viper"
)

func GenerateSimpleMultisigAddress(depositorPubKey string, podEthAddress string) (string, error) {
	dbconn := db.InitDB()
	wallet := viper.GetString("wallet_name")

	descriptor, err := buildSimpleMultisigDescriptor(depositorPubKey)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := btcComms.GetDescriptorInfo(descriptor, wallet)
	if err != nil {
		fmt.Println("error in getting descriptorinfo : ", err)
		return "", err
	}

	fmt.Println("Descriptor : ", resp.Descriptor)

	err = btcComms.ImportDescriptor(resp.Descriptor, wallet)
	if err != nil {
		fmt.Println("error in importing descriptor : ", err)
		return "", err
	}

	address, err := btcComms.DeriveAddress(wallet, resp.Descriptor)
	if err != nil {
		fmt.Println("error in getting address : ", err)
	}

	addressInfo, err := btcComms.GetAddressInfo(address, wallet)
	if err != nil {
		fmt.Println("Error getting address info : ", err)
		return "", err
	}
	// Decode Hex string to bytes

	err = db.InsertMultiSigAddress(dbconn, address, addressInfo.Hex, podEthAddress)
	dbconn.Close()
	if err != nil {
		fmt.Println("error in inserting multisig address : ", err)
		return "", err
	}

	dbconn.Close()
	return address, nil
}

func buildSimpleMultisigDescriptor(depositorPubKey string) (string, error) {
	required := 2
	judgePubKey := viper.GetString("btc_xpublic_key")

	descriptorScript := fmt.Sprintf("wsh(multi(%d,%s,%s))", required, judgePubKey, depositorPubKey)
	fmt.Println(descriptorScript)
	return descriptorScript, nil
}

func GenerateMultisigwithdrawTx(withdrawBTCAddress string, podEthAddr string) (string, int64, string, error) {
	dbconn := db.InitDB()
	defer dbconn.Close()
	wallet := viper.GetString("wallet_name")
	multiSigAddresses := db.QueryMultisigAddressByPodAddress(dbconn, podEthAddr)
	if len(multiSigAddresses) <= 0 {
		fmt.Println("no multisig address found")
		return "", 0, "", errors.New("no multisig found")
	}
	multiSigAddress := multiSigAddresses[0]

	utxos, err := utils.ListUnspentBtcUtxos(multiSigAddress.Address)
	if err != nil {
		fmt.Println("error in getting utxos : ", err)
		return "", 0, "", err
	}

	if len(utxos) <= 0 {
		fmt.Println("INFO : No funds in address : ", multiSigAddress.Address)
		return "", 0, "", errors.New("no BTC UTXO found")
	}

	var inputs []btcComms.TxInput
	var outputs []btcComms.TxOutput
	totalAmountTxIn := float64(0)

	for _, u := range utxos { //ideally the height should be masked with 0x0000ffff
		inputs = append(inputs, btcComms.TxInput{Txid: u.TxID, Vout: int64(u.Vout), Sequence: int64(wire.MaxTxInSequenceNum - 10)})
		totalAmountTxIn += u.Amount
	}

	withdrawAddr, err := utils.HexToBech32(withdrawBTCAddress, &chaincfg.SigNetParams)
	if err != nil {
		fmt.Println("error in converting to bech32 : ", err)
		return "", 0, "", err
	}
	fmt.Println("withdraw btc addr : ", withdrawAddr)
	fmt.Println("total amount in BTC: ", totalAmountTxIn)

	outputs = append(outputs, btcComms.TxOutput{withdrawAddr: totalAmountTxIn})

	hexTx, err := btcComms.CreateRawTx(inputs, outputs, 0, wallet)
	if err != nil {
		fmt.Println("error in creating raw tx : ", err)
		return "", 0, "", err
	}

	multisigTx, err := utils.CreateTxFromHex(hexTx)
	if err != nil {
		fmt.Println("error decoding tx : ", err)
		return "", 0, "", err
	}

	fee, err := utils.GetFeeFromBtcNode(multisigTx)
	if err != nil {
		fmt.Println("error in getting fee : ", err)
		return "", 0, "", err
	}

	totalAmountInBTC := totalAmountTxIn - utils.SatsToBtc(fee)
	fmt.Println("fee in sats : ", fee)
	fmt.Println("total amount in btc after fee : ", totalAmountInBTC)

	outputs = []btcComms.TxOutput{btcComms.TxOutput{withdrawAddr: totalAmountInBTC}}

	p, err := btcComms.CreatePsbt(inputs, outputs, 0, wallet)
	if err != nil {
		fmt.Println("error in creating psbt : ", err)
		return "", 0, "", err
	}

	fmt.Println("transaction base64 psbt: ", p)

	db.MarkMultisigProcessed(dbconn, multiSigAddress.Address)
	return p, utils.BtcToSats(totalAmountInBTC), multisigTx.TxHash().String(), nil
}

func SignMultisigPSBT(psbt string) (string, error) {
	walletName := viper.GetString("signer_wallet_name")
	_, psbt, err := btcComms.SignPsbt(psbt, walletName, false)
	if err != nil {
		fmt.Println("error in signing psbt : ", err)
		return "", err
	}

	return psbt, nil
}
