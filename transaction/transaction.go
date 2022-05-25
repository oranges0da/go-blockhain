package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	ID       []byte // hash of transaction
	In       []TxInput
	Out      []TxOutput
	Locktime int64
}

type TxInput struct {
	ID  []byte // id of transaction pointing to correlated output in badger db
	Sig string
	Out int // db index location of output of transaction
}

type TxOutput struct {
	Value  float64
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

func NewCoinbase(to string, sig string) *Transaction {
	tx := &Transaction{
		ID:       []byte{},
		In:       []TxInput{},
		Out:      []TxOutput{},
		Locktime: 0,
	}

	tx.HashTx()
}
