package BtcDepositConfirmer

import (
	"fmt"

	"github.com/AhmadAshraf2/Judge-AVS/comms"
	"github.com/AhmadAshraf2/Judge-AVS/db"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
)

func CheckDeposit() {
	for {
		dbconn := db.InitDB()
		defer dbconn.Close()
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
							_, err := comms.CallConfirmBtcDeposit(request.PodAddress, request.OperatorAddress, request.TransactionID, *request.Amount)
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
	}

}
