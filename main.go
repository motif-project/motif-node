package main

import (
	"fmt"
	"net/url"
	"sync"

	_ "github.com/lib/pq"

	"github.com/AhmadAshraf2/Judge-AVS/api"
	"github.com/AhmadAshraf2/Judge-AVS/bridge"
	"github.com/AhmadAshraf2/Judge-AVS/utils"
	"github.com/spf13/viper"
)

func initialize() url.URL {
	utils.InitConfigFile()
	utils.LoadBtcWallet(viper.GetString("wallet_name"))

	// keystoreDir := "keystore"
	// var ethAccount accounts.Account
	// // Check if the keystore directory exists
	// if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
	// 	ethAccount = utils.GenerateEthKeyPair()
	// } else {
	// 	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	// 	accounts := ks.Accounts()
	// 	ethAccount = accounts[0]
	// }

	forkscannerHost := fmt.Sprintf("%v:%v", viper.Get("forkscanner_host"), viper.Get("forkscanner_ws_port"))
	forkscannerUrl := url.URL{Scheme: "ws", Host: forkscannerHost, Path: "/"}

	return forkscannerUrl
}

func main() {
	forkscannerUrl := initialize()
	var wg sync.WaitGroup
	wg.Add(1)
	api.Server()
	bridge.WatchAddress(forkscannerUrl)
	wg.Wait()
}
