package transaction

func FindUTXO(addr string) []TxOutput {
	var UTXOs []TxOutput
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
