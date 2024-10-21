package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func checkBtcNodeHealth() bool {
	url := viper.GetString("btc_node_host")
	user := viper.GetString("btc_node_user")
	password := viper.GetString("btc_node_pass")

	data := &RPCRequest{
		Jsonrpc: "1.0",
		ID:      "curltest",
		Method:  "getblockchaininfo",
		Params:  []interface{}{},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return false
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return false
	}

	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	return true
}

func checkForkscannerHealth() bool {
	ipaddress := viper.GetString("forkscanner_host")
	port := viper.GetString("forkscanner_rpc_port")
	url := fmt.Sprintf("http://%v:%v", ipaddress, port)
	payload := `{"jsonrpc": "2.0", "id": 1, "method": "get_tips", "params": null}`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println(err)
		return false
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer resp.Body.Close()
	return true
}

func runHealthCheck() (bool, bool) {
	return checkBtcNodeHealth(), checkForkscannerHealth()
}
