package main

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/BitDSM/BitDSM-Node/BtcDepositConfirmer"
	"github.com/BitDSM/BitDSM-Node/address"
	"github.com/BitDSM/BitDSM-Node/api"
	"github.com/BitDSM/BitDSM-Node/ethComms"
	"github.com/BitDSM/BitDSM-Node/operator"
	"github.com/BitDSM/BitDSM-Node/utils"
)

func initialize() {
	utils.InitConfigFile()
	env := viper.GetString("env")
	if env != "dev" && env != "prod" {
		fmt.Println("Invalid environment")
		panic("Invalid environment")
	}

	utils.LoadBtcWallet(viper.GetString("wallet_name"))
	ethAccount := ethComms.LoadEthAccount()
	fmt.Println("Eth account: ", ethAccount.Address.Hex())
	operator.RegisterOperator()
	// psbt, _, _ := address.GenerateMultisigwithdrawTx("tb1qw7yedz2cevhl422947hns03yrngxhu43mhf5g0", "0x32DDC1F8cb6A602B956a31c846486531AbfB98b6")
	psbt := "cHNidP8BAFICAAAAAZ7fMPyc9aKRN3aQl76FpjBcDl8TeCK9vdhLkC84wFKLAAAAAAD1////AdQGAAAAAAAAFgAUd4mWiVjLL/qpRa+vOD4kHNBr8rEAAAAAAAEAiQIAAAABHGJMV/6F6OF2vdr8VOvaNYdUjDxKtwj6U6tZc7OEJXoAAAAAAP3///8CuAsAAAAAAAAiACBnkXmevzIAqlQfQe3OHGXAFYYvs9ohwr/dMJspbgRhwGMCAAAAAAAAIlEgcG9TI/pZVVREj7lb4v9h/r0j7zHaTy6aEQswqZhB9QjvfAMAAQEruAsAAAAAAAAiACBnkXmevzIAqlQfQe3OHGXAFYYvs9ohwr/dMJspbgRhwAEFR1IhA8sjVC9pjtHmF6YjQptYXZj7keRIOZSdtBJrKg1acyCwIQI6vMM/6vT76nEQCqlIqcX+Xm4htu3+V2WjHeb8K8JCGFKuIgYCOrzDP+r0++pxEAqpSKnF/l5uIbbt/ldlox3m/CvCQhgEd4mWiSIGA8sjVC9pjtHmF6YjQptYXZj7keRIOZSdtBJrKg1acyCwBMEaNmoAAA=="
	_, psbt, _ = address.SignMultisigPSBT(psbt)
	fmt.Println(psbt)
}

func main() {
	initialize()
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Server()
	wg.Add(1)
	go BtcDepositConfirmer.CheckDeposit()
	wg.Add(1)
	go ethComms.SubscribeToDepositRequests()
	wg.Add(1)
	go BtcDepositConfirmer.CheckWithdraw()
	ethComms.SubscribeToWithdrawRequests()
	wg.Wait()
}
