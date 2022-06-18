package utils

import (
	"encoding/hex"

	"github.com/oranges0da/goblockchain/transaction"
)

// find all unspent transactions for a certain address
func FindUnspentTxs(addr string) []transaction.Transaction {
	var unspentTxs []transaction.Transaction
	var spentTxs = make(map[string][]int)

	blocks, err := GetBlocks()
	Handle(err, "error getting blocks in FindUnspentTxs")

	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txId := hex.EncodeToString(tx.ID)

			if !tx.IsCoinbase() {
				for _, in := range tx.Inputs {
					// if tx is not coinbase, and tx belongs to the address, add it to spendTxs
					if in.InCanUnlock(addr) {
						spentTxs[txId] = append(spentTxs[txId], in.Out)
					}
				}
			}
		Outputs:
			for outIdx, out := range tx.Outputs {
				if spentTxs[txId] != nil {
					for _, spentOut := range spentTxs[txId] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				if out.OutCanUnlock(addr) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}
		}
	}
	return unspentTxs
}
