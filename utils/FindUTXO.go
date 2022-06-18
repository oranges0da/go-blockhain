package utils

import "github.com/oranges0da/goblockchain/transaction"

func FindUTXO(addr string) []transaction.TxOutput {
	var UTXOs []transaction.TxOutput
	unspentTxs := FindUnspentTxs(addr)

	return UTXOs
}
