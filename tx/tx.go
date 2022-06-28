package tx

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/oranges0da/goblockchain/block_utils"
	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/utils"
)

type Transaction struct {
	ID       []byte // hash of transaction
	Inputs   []TxInput
	Outputs  []TxOutput
	Locktime int
}

func New(to, from string, amt int) {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOuts := FindSpendableOuts(from, amt)

	if acc < amt {
		log.Panic("Insufficient funds.")
	}
}

// msg is any string that miner can put into blockchain forever
func NewCoinbase(addr, msg string) *Transaction {
	// not refrencing any previous output for this txs input, so ID and PubKey will be empty, and Vout is not accesible(-1 is not an index)
	in := TxInput{
		ID:     []byte{},
		Vout:   -1,
		Sig:    []byte{},
		PubKey: []byte{00000000000000000000000000000000},
	}
	out := NewTxOut(50, addr)

	tx := &Transaction{
		ID:       nil,
		Inputs:   []TxInput{in},
		Outputs:  []TxOutput{out},
		Locktime: 0,
	}

	tx.Hash()

	return tx
}

func (tx *Transaction) Hash() {
	byte_tx := utils.ToByte(tx)

	hash := sha256.Sum256(byte_tx)

	tx.ID = hash[:]
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Vout == -1
}

// return array of unspent txs for a certain address
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
func FindUnspentTxs(addr string) []Transaction {
	var unspentTxs []Transaction
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
