package utils

import "github.com/oranges0da/goblockchain/transaction"

func FindUTXO(addr string) []transaction.TxOutput {
	var UTXOs []transaction.TxOutput
	unspentTxs := FindUnspentTxs(addr)

	for _, tx := range unspentTxs {
		for _, out := range tx.Outputs {
			if out.OutCanUnlock(addr) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}
