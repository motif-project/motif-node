package BtcDepositConfirmer

import (
	"fmt"
	"math/big"
	"time"

	"github.com/AhmadAshraf2/Judge-AVS/db"
	"github.com/AhmadAshraf2/Judge-AVS/ethComms"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
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

			multisigaddresses := db.QueryMultisigAddresses(dbconn)

			for _, multiSigAddress := range multisigaddresses {
				for _, txOut := range tx.Vout {
					for _, address := range txOut.ScriptPubKey.Addresses {
						if multiSigAddress.Address == address {
							bigInt := new(big.Int)
							// Convert the string to big.Int
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
