package btcComms

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// VerifyTransactionSignatures verifies that the signatures in a transaction are correct
// and checks if any output matches the given address.
func VerifyTransactionSignatures(txHex string, rpcClient *rpcclient.Client, outputAddr string) (bool, error) {
	// Decode the transaction hex string
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return false, fmt.Errorf("failed to decode transaction hex: %v", err)
	}

	// Deserialize the transaction
	tx := wire.NewMsgTx(wire.TxVersion)
	err = tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return false, fmt.Errorf("failed to deserialize transaction: %v", err)
	}

	// Check if any output matches the given address
	matched := false
	for _, out := range tx.TxOut {
		// Decode the PkScript to extract the address
		_, addresses, _, err := txscript.ExtractPkScriptAddrs(out.PkScript, rpcClient.ChainParams())
		if err != nil {
			return false, fmt.Errorf("failed to extract address from PkScript: %v", err)
		}

		// Check if the output address matches the given address
		for _, addr := range addresses {
			if addr.EncodeAddress() == outputAddr {
				matched = true
				break
			}
		}
	}

	if !matched {
		return false, fmt.Errorf("no output matches the given address: %s", outputAddr)
	}

	// Iterate through each input and verify the signature
	for i, txIn := range tx.TxIn {
		// Get the previous output referenced by this input
		prevOut := txIn.PreviousOutPoint

		// Fetch the previous transaction using getrawtransaction
		prevTx, err := rpcClient.GetRawTransactionVerbose(&prevOut.Hash)
		if err != nil {
			return false, fmt.Errorf("failed to fetch previous transaction for input %d: %v", i, err)
		}

		// Ensure the output index exists in the previous transaction
		if int(prevOut.Index) >= len(prevTx.Vout) {
			return false, fmt.Errorf("invalid output index %d for input %d", prevOut.Index, i)
		}

		// Get the previous output details
		prevTxOut := prevTx.Vout[prevOut.Index]

		// Decode the PkScript from the previous output
		pkScript, err := hex.DecodeString(prevTxOut.ScriptPubKey.Hex)
		if err != nil {
			return false, fmt.Errorf("failed to decode PkScript for input %d: %v", i, err)
		}

		// Handle SegWit and legacy transactions
		var value int64
		if txIn.Witness != nil && len(txIn.Witness) > 0 {
			// SegWit transaction: Use the value from the previous output
			value = int64(prevTxOut.Value * 1e8) // Convert BTC to satoshis
		} else {
			// Legacy transaction: Value is not required for signature verification
			value = 0
		}

		// Verify the signature for this input
		vm, err := txscript.NewEngine(pkScript, tx, i, txscript.StandardVerifyFlags, nil, nil, value)
		if err != nil {
			return false, fmt.Errorf("failed to create script engine for input %d: %v", i, err)
		}

		err = vm.Execute()
		if err != nil {
			return false, fmt.Errorf("signature verification failed for input %d: %v", i, err)
		}
	}

	// If all inputs are valid, the signatures are correct
	return true, nil
}
