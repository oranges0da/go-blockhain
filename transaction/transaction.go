package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	ID       []byte // hash of transaction
	Inputs   []TxInput
	Outputs  []TxOutput
	Locktime int
}

type TxInput struct {
	ID  []byte // id of transaction pointing to correlated output in badger db
	Sig string
	Out int // db index location of output of transaction
}

type TxOutput struct {
	Value  int    // amt of satoshis in output (amt that is being sent)
	PubKey string // receiver's public key/address
}

func (tx *Transaction) HashTx() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)

	if err != nil {
		log.Fatalf("Error hashing transaction: %s", err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func NewCoinbase(addr string, sig string) *Transaction {
	tx := &Transaction{
		ID:       []byte{},
		Inputs:   []TxInput{},
		Outputs:  []TxOutput{},
		Locktime: 0,
	}

	tx.Inputs = append(tx.Inputs, TxInput{Sig: sig, Out: -1})
	tx.Outputs = append(tx.Outputs, TxOutput{Value: 50, PubKey: addr})

	tx.HashTx()

	return tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Out == -1
}

func (in *TxInput) CanUnlockInput(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
