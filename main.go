package main

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/AhmadAshraf2/BitDSM-Node/BtcDepositConfirmer"
	"github.com/AhmadAshraf2/BitDSM-Node/api"
	"github.com/AhmadAshraf2/BitDSM-Node/ethComms"
	"github.com/AhmadAshraf2/BitDSM-Node/operator"
	"github.com/AhmadAshraf2/BitDSM-Node/utils"
)

func initialize() {
	utils.InitConfigFile()
	utils.LoadBtcWallet(viper.GetString("wallet_name"))
	ethAccount := ethComms.LoadEthAccount()
	fmt.Println("Eth account: ", ethAccount.Address.Hex())
	operator.RegisterOperator()
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
