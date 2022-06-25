package transaction

import (
	"encoding/hex"

	"github.com/oranges0da/goblockchain/utils"
)

// find all unspent transactions for a certain address
func FindUnspentTxs(addr string) []Transaction {
	var unspentTxs []Transaction
	var spentTxs = make(map[string][]int)

	blocks, err := utils.GetBlocks()
	utils.Handle(err, "error getting blocks in FindUnspentTxs")

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
