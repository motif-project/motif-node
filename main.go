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
	address.GenerateMultisigwithdrawTx("tb1qw7yedz2cevhl422947hns03yrngxhu43mhf5g0", "32DDC1F8cb6A602B956a31c846486531AbfB98b6")
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
