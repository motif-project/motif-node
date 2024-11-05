package BtcDepositConfirmer

import (
	"fmt"
	"log"
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

		if len(depositRequests) <= 0 {
			fmt.Println("no deposit requests found")
			time.Sleep(3 * time.Minute)
			continue
		}

		for _, request := range depositRequests {
			fmt.Println("found deposit requests")
			fmt.Println("request: ", request.TransactionID)
			tx, err := utils.GetRawTransaction(request.TransactionID)
			if err != nil {
				fmt.Println("Failed to get raw transaction: ", err)
				continue
			}
			fmt.Println("found raw tx")
			if tx.Confirmations <= 2 {
				continue
			}

			fmt.Println("tx confirmed : ", tx.Hex)

			t, err := utils.CreateTxFromHex(tx.Hex)
			if err != nil {
				fmt.Println("Failed to create tx from hex: ", err)
			}

			fmt.Println("tx: ", t)
			fmt.Println(t.TxOut[1])

			_, addresses, _, err := txscript.ExtractPkScriptAddrs(t.TxOut[1].PkScript, &chaincfg.SigNetParams)
			if err != nil {
				log.Fatalf("Failed to extract addresses: %v", err)
			}

			fmt.Println("============")
			fmt.Println("addresses: ", addresses)

			multisigaddresses := db.QueryMultisigAddresses(dbconn)

			for _, multiSigAddress := range multisigaddresses {
				fmt.Println("multisig found")
				fmt.Println("multisig address: ", multiSigAddress.Address)
				for _, txOut := range tx.Vout {
					fmt.Println("checking txout : ", txOut)
					fmt.Println("checking txout  hex: ", txOut.ScriptPubKey.Hex)
					for _, address := range txOut.ScriptPubKey.Addresses {
						fmt.Println("checking address: ", address)
						if multiSigAddress.Address == address {
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
