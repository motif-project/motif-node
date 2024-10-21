package types

type WatchtowerResponse struct {
	Jsonrpc string
	Method  string
	Params  []WatchtowerNotification
}

type WatchtowerTxInput struct {
	Address string
	Amount  uint64
	Txid    string
	Vout    uint32
}

type WatchtowerNotification struct {
	Block            string
	Height           uint64
	Receiving        string
	Satoshis         uint64
	Receiving_txid   string
	Sending_txinputs []WatchtowerTxInput
	Archived         bool
	Receiving_vout   uint64
	Sending          string
	Sending_vout     int32
}

type MultiSigAddress struct {
	Address    string
	Script     string
	EthAddress string
	Signed     bool
	Archived   bool
}

type Utxo struct {
	Txid   string
	Vout   uint32
	Amount uint64
}

type BtcPubkeyArgs struct {
	BTCPubKey string
	EthAddr   string
}

type GetUnsignedPsbtArgs struct {
	EthAddr         string
	WithdrawBtcAddr string
}

type SubmitSignedPSBT struct {
	Psbt string
}
