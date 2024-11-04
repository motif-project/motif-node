package comms

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/AhmadAshraf2/Judge-AVS/PodManager"
	"github.com/AhmadAshraf2/Judge-AVS/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func SubscribeToDepositRequests() {
	// Create a new instance of the contract binding

	oprEthAccount := LoadEthAccount()
	client, err := GetEthClient()
	if err != nil {
		fmt.Println("Failed to get eth client:", err)
		panic(err)
	}

	podManagerAddr := common.HexToAddress(viper.GetString("pod_manager_address"))
	podManager, err := PodManager.NewPodManager(podManagerAddr, client)
	if err != nil {
		fmt.Println("Failed to instantiate contract:", err)
		panic(err)
	}

	// Create a channel for the events
	ch := make(chan *PodManager.PodManagerVerifyBitcoinDepositRequest)

	// Create a subscription
	sub, err := podManager.WatchVerifyBitcoinDepositRequest(
		&bind.WatchOpts{Context: context.Background()},
		ch,
		nil, // pod filter (nil means all pods)
		nil, // operator filter (nil means all operators)
	)
	if err != nil {
		fmt.Println("Failed to subscribe to events:", err)
		panic(err)
	}

	fmt.Println("Successfully subscribed to VerifyBitcoinDepositRequest events")

	// Handle events in a loop
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Subscription error:", err)

		case event := <-ch:
			if event.Operator == oprEthAccount.Address {
				handleDepositRequest(event)
			}
		}
	}
}

func handleDepositRequest(event *PodManager.PodManagerVerifyBitcoinDepositRequest) {
	dbconn := db.InitDB()
	defer dbconn.Close()

	if event.BitcoinDepositRequest.IsPending == false {
		return
	}

	transactionId := hex.EncodeToString(event.BitcoinDepositRequest.TransactionId[:])
	db.InsertDepositRequest(dbconn, event.Pod.Hex(), event.Operator.Hex(), transactionId, event.BitcoinDepositRequest.Amount)

}
