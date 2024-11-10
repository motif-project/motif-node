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
	utils.LoadBtcWallet(viper.GetString("wallet_name"))
	ethAccount := ethComms.LoadEthAccount()
	fmt.Println("Eth account: ", ethAccount.Address.Hex())
	operator.RegisterOperator()
	address.GenerateSimpleMultisigAddress("4d6b1b1c0594fbf65e6ffdd15bc01cf9984dabcc", "0x15B17cd2aF4ee8B8163EfAc54Cb740df01384215")
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
