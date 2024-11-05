package main

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/AhmadAshraf2/Judge-AVS/BtcDepositConfirmer"
	"github.com/AhmadAshraf2/Judge-AVS/PodManager"
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

	request := &PodManager.PodManagerBitcoinWithdrawalPSBTRequest{
		Pod:             common.HexToAddress("0x365119a2De4e389B542AAe945f26cb21cA3f3Fbf"),
		Operator:        common.HexToAddress("0x60b3b41240Fb9d353acE1E37B0CE79054154eA40"),
		WithdrawAddress: []byte{205, 219, 47, 200, 4, 245, 199, 159, 16, 228, 183, 54, 241, 116, 5, 224, 234, 118, 227, 29},
		Raw:             types.Log{},
	}

	ethComms.HandleWithdrawalRequest(request)

	ethComms.SubscribeToWithdrawRequests()
	wg.Wait()
}
