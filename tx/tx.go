package tx

import (
	"encoding/hex"
	"log"

	"github.com/oranges0da/goblockchain/block_utils"
	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/wallet"
)

func New(to, from string, amt int, locktime int) *model.Transaction {
	var inputs []model.TxInput
	var outputs []model.TxOutput

	// check if spender has enough unspent outputs to send this tx
	acc, validOuts := FindSpendableOuts(from, amt)

	if acc < amt {
		log.Panic("Insufficient funds.")
	}

	// load wallet of sender
	wallets := wallet.LoadWallets()
	w := wallets.Get(from)

	// for every valid output of sender, append prev Vout to this tx's inputs
	for txid, outs := range validOuts {
		txID, err := hex.DecodeString(txid)
		handle.Handle(err, "Error decoding txID while making new tx.")

		for _, out := range outs {
			input := model.TxInput{txID, out, nil, w.PubKey}
			inputs = append(inputs, input)
		}
	}

	// outputs include output to reciever
	outputs = append(outputs, *NewTxOut(amt, to))

	// and if sender has more than amt, the remainder of their funds will be sent back
	if acc > amt {
		outputs = append(outputs, *NewTxOut(acc-amt, from))
	}

	tx := &model.Transaction{
		ID:       nil,
		Inputs:   inputs,
		Outputs:  outputs,
		Locktime: locktime,
	}

	tx.Hash()
	tx.Sign()

	return tx
}

// msg is any string that miner can put into blockchain forever
func NewCoinbase(addr, msg string) *model.Transaction {
	// not refrencing any previous output for this txs input, so ID and PubKey will be empty, and Vout is not accesible(-1 is not an index)
	in := model.TxInput{
		ID:     []byte{},
		Vout:   -1,
		Sig:    []byte{},
		PubKey: []byte{00000000000000000000000000000000},
	}
	out := NewTxOut(50, addr)

	tx := &model.Transaction{
		ID:       nil,
		Inputs:   []model.TxInput{in},
		Outputs:  []model.TxOutput{out},
		Locktime: 0,
	}

	tx.Hash()

	return tx
}

// return array of unspent txs for a certain address
func FindUTXO(addr string) []model.TxOutput {
	var UTXOs []model.TxOutput
	unspentTxs := FindUnspentTxs(addr)

	for _, tx := range unspentTxs {
		for _, out := range tx.Outputs {
			if OutCanUnlock(out, addr) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// return amt and map with id (hash) of tx as key, and Vout as value
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

// find all unspent transactions for a certain address
func FindUnspentTxs(addr string) []model.Transaction {
	var unspentTxs []model.Transaction
	var spentTxs = make(map[string][]int)

	blocks, err := block_utils.GetBlocks()
	handle.Handle(err, "Error getting blocks in FindUnspentTxs")

	for _, block := range blocks {
		tx := block.Transaction
		txId := hex.EncodeToString(tx.ID)

	Outputs:
		for outIdx, out := range tx.Outputs {
			if spentTxs[txId] != nil {
				for _, spentOut := range spentTxs[txId] {
					if spentOut == outIdx {
						continue Outputs
					}
				}
			}
			if OutCanUnlock(out, addr) {
				unspentTxs = append(unspentTxs, *tx)
			}
		}
		if !tx.IsCoinbase() {
			for _, in := range tx.Inputs {
				if InCanUnlock(in, addr) {
					inTxId := hex.EncodeToString(in.ID)
					spentTxs[inTxId] = append(spentTxs[inTxId], in.Vout)
				}
			}
		}
	}
	return unspentTxs
}
