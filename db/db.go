package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/AhmadAshraf2/Judge-AVS/types"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", viper.Get("DB_host"), viper.Get("DB_port"), viper.Get("DB_user"), viper.Get("DB_password"), viper.Get("DB_name"))
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println("DB error : ", err)
		panic(err)
	}
	fmt.Println("DB initialized")
	return db
}

func InsertMultiSigAddress(dbconn *sql.DB, address string, script string, ethAddr string) {
	_, err := dbconn.Exec("INSERT into multi_sig_address VALUES ($1, $2, $3, $4, $5)",
		address,
		script,
		ethAddr,
		false,
		false,
	)
	if err != nil {
		fmt.Println("An error occured while executing insert multi sig address query: ", err)
	}
}

func QueryMultisigAddressByEthAddress(dbconn *sql.DB, ethAddr string) []types.MultiSigAddress {
	// fmt.Println("getting address for height: ", height)
	var DB_reader *sql.Rows
	var err error
	DB_reader, err = dbconn.Query("select * from multi_sig_address where ethaddress = $1 and archived = false", ethAddr)

	if err != nil {
		fmt.Println("An error occured while query address by Eth Address: ", err)
	}

	defer DB_reader.Close()
	addresses := make([]types.MultiSigAddress, 0)

	for DB_reader.Next() {
		address := types.MultiSigAddress{}
		err := DB_reader.Scan(
			&address.Address,
			&address.Script,
			&address.EthAddress,
			&address.Signed,
			&address.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		addresses = append(addresses, address)
	}

	return addresses
}

func QueryMultisigAddresses(dbconn *sql.DB) []types.MultiSigAddress {
	// fmt.Println("getting address for height: ", height)
	var DB_reader *sql.Rows
	var err error
	DB_reader, err = dbconn.Query("select * from multi_sig_address where archived = false")

	if err != nil {
		fmt.Println("An error occured while query address: ", err)
	}

	defer DB_reader.Close()
	addresses := make([]types.MultiSigAddress, 0)

	for DB_reader.Next() {
		address := types.MultiSigAddress{}
		err := DB_reader.Scan(
			&address.Address,
			&address.Script,
			&address.EthAddress,
			&address.Signed,
			&address.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		addresses = append(addresses, address)
	}

	return addresses
}

func QueryUtxo(dbconn *sql.DB, address string) []types.Utxo {
	DB_reader, err := dbconn.Query("select txid, Receiving_vout, satoshis from notification where receiving = $1", address)
	if err != nil {
		fmt.Println("An error occured while query utxo: ", err)
	}

	defer DB_reader.Close()
	utxos := make([]types.Utxo, 0)

	for DB_reader.Next() {
		utxo := types.Utxo{}
		err := DB_reader.Scan(
			&utxo.Txid,
			&utxo.Vout,
			&utxo.Amount,
		)
		if err != nil {
			fmt.Println(err)
		}
		utxos = append(utxos, utxo)
	}
	return utxos
}

func InsertDepositRequest(dbconn *sql.DB, podAddr string, operatorAddr string, txid string, amount *big.Int) {
	_, err := dbconn.Exec("INSERT into deposit_requests VALUES ($1, $2, $3, $4, $5)",
		podAddr,
		operatorAddr,
		txid,
		amount,
		false,
	)
	if err != nil {
		fmt.Println("An error occured while insert deposit request query: ", err)
	}
}

func QueryDepositRequests(dbconn *sql.DB) []types.BtcDepositRequest {
	DB_reader, err := dbconn.Query("select * from deposit_requests where Archived = false")
	if err != nil {
		fmt.Println("An error occured while query deposit request: ", err)
		return nil
	}

	defer DB_reader.Close()
	depositRequests := make([]types.BtcDepositRequest, 0)

	for DB_reader.Next() {
		depositRequest := types.BtcDepositRequest{}
		err := DB_reader.Scan(
			&depositRequest.PodAddress,
			&depositRequest.OperatorAddress,
			&depositRequest.TransactionID,
			&depositRequest.Amount,
			&depositRequest.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		depositRequests = append(depositRequests, depositRequest)
	}
	return depositRequests
}

func MarkDepositRequestAsConfirmed(dbconn *sql.DB, txid string) {
	_, err := dbconn.Exec("UPDATE deposit_requests SET archived = true WHERE TransactionID = $1", txid)
	if err != nil {
		fmt.Println("An error occured while updating deposit request query: ", err)
	}
}
