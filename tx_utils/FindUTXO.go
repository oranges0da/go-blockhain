package tx_utils

import (
	"github.com/oranges0da/goblockchain/tx"
)

func FindUTXO(addr string) []tx.TxOutput {
	var UTXOs []tx.TxOutput
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
