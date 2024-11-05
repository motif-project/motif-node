package BtcDepositConfirmer

import (
	"fmt"
	"math/big"
	"time"

	"github.com/AhmadAshraf2/Judge-AVS/db"
	"github.com/AhmadAshraf2/Judge-AVS/ethComms"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func CheckDeposit() {
	fmt.Println("starting check deposit process")
	dbconn := db.InitDB()
	defer dbconn.Close()
	for {
		depositRequests := db.QueryDepositRequests(dbconn)

		for _, request := range depositRequests {
			tx, err := utils.GetRawTransaction(request.TransactionID)
			if err != nil {
				continue
			}
			if tx.Confirmations <= 2 {
				continue
			}

			transaction, err := utils.CreateTxFromHex(tx.Hex)
			if err != nil {
				fmt.Println("Failed to create tx from hex: ", err)
				continue
			}

			multisigaddresses := db.QueryMultisigAddresses(dbconn)

			for _, multiSigAddress := range multisigaddresses {
				for _, txOut := range transaction.TxOut {
					_, addresses, _, err := txscript.ExtractPkScriptAddrs(txOut.PkScript, &chaincfg.SigNetParams)
					if err != nil {
						fmt.Println("Failed to extract addresses: ", err)
						continue
					}
					for _, address := range addresses {
						if multiSigAddress.Address == address.EncodeAddress() {
							fmt.Println("found my deposit requests")
							bigInt := new(big.Int)
							_, success := bigInt.SetString(request.Amount, 10)
							if !success {
								fmt.Print("Failed to convert string to big.Int")
								continue
							}
							_, err := ethComms.CallConfirmBtcDeposit(request.PodAddress, request.OperatorAddress, request.TransactionID, *bigInt)
							if err != nil {
								fmt.Println("Failed to call confirm btc deposit: ", err)
								continue
							}
							db.MarkDepositRequestAsConfirmed(dbconn, request.TransactionID)
						}
					}
				}
			}

		}
		time.Sleep(3 * time.Minute)

	}

}
