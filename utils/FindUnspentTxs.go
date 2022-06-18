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

		Outputs:
			for outIdx, out := range tx.Outputs {
				if spentTxs[txId] != nil {
					if spentTxs[txID] == outIdx {
						continue Outputs
					}
				}
			}
		}
	}
}
