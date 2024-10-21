package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	_ "github.com/lib/pq"

	"github.com/AhmadAshraf2/Judge-AVS/api"
	"github.com/AhmadAshraf2/Judge-AVS/bridge"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/viper"
)

func initialize() (url.URL, accounts.Account) {
	utils.InitConfigFile()

	utils.LoadBtcWallet(viper.GetString("wallet_name"))

	keystoreDir := "keystore"
	var ethAccount accounts.Account
	// Check if the keystore directory exists
	if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
		ethAccount = utils.GenerateEthKeyPair()
	} else {
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
		accounts := ks.Accounts()
		ethAccount = accounts[0]
	}

	forkscannerHost := fmt.Sprintf("%v:%v", viper.Get("forkscanner_host"), viper.Get("forkscanner_ws_port"))
	forkscannerUrl := url.URL{Scheme: "ws", Host: forkscannerHost, Path: "/"}

	return forkscannerUrl, ethAccount
}

func main() {
	forkscannerUrl, _ := initialize()
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Server()
	bridge.WatchAddress(forkscannerUrl)
	wg.Wait()
}
