package tx

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/oranges0da/goblockchain/utils"
)

type Transaction struct {
	ID       []byte // hash of transaction
	Inputs   []TxInput
	Outputs  []TxOutput
	Locktime int
}

func (tx *Transaction) Sign(privKey *ecdsa.PrivateKey, prevTxs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
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
