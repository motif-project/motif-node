package main

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/AhmadAshraf2/Judge-AVS/BtcDepositConfirmer"
	"github.com/AhmadAshraf2/Judge-AVS/api"
	"github.com/AhmadAshraf2/Judge-AVS/ethComms"
	"github.com/AhmadAshraf2/Judge-AVS/operator"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
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
	ethComms.SubscribeToWithdrawRequests()
	wg.Wait()
}
