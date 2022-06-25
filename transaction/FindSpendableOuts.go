package transaction

import "encoding/hex"

func FindSpendableOuts(addr string, amt int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxs := FindUnspentTxs(addr)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txId := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Outputs {
			if string(out.PubKeyHash) == addr && accumulated < amt {
				accumulated += out.Value
				unspentOuts[txId] = append(unspentOuts[txId], outIdx)

				if accumulated >= amt {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOuts
}
