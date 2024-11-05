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

	// bigAmount := new(big.Int).SetInt64(1000000)
	// ethComms.CallConfirmBtcDeposit("0x6B3925c6CaC45b31192ed4fD9B05AC46E043fE58", "0x60b3b41240Fb9d353acE1E37B0CE79054154eA40", "d74b5d23c32f211a41db07e6fd8eba8cf1e9c3df9be7a9de172bbf52a7fb8a49", *bigAmount)
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
