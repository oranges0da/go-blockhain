package tx_utils

import (
	"encoding/hex"

	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/tx"
	"github.com/oranges0da/goblockchain/utils"
)

// find all unspent transactions for a certain address
func FindUnspentTxs(addr string) []tx.Transaction {
	var unspentTxs []tx.Transaction
	var spentTxs = make(map[string][]int)

	blocks, err := utils.GetBlocks()
	handle.Handle(err, "error getting blocks in FindUnspentTxs")

	for _, block := range blocks {
		tx := block.Transaction
		txId := tx.ID

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
		if !tx.IsCoinbase() {
			for _, in := range tx.Inputs {
				if in.InCanUnlock(addr) {
					inTxId := hex.EncodeToString(in.ID)
					spentTxs[inTxId] = append(spentTxs[inTxId], in.Vout)
				}
			}
		}
	}
	return unspentTxs
}
