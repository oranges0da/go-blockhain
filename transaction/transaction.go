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
		In:       []TxInput{},
		Out:      []TxOutput{},
		Locktime: 0,
	}

	tx.In = append(tx.In, TxInput{Sig: sig, Out: -1})
	tx.Out = append(tx.Out, TxOutput{Value: 50, PubKey: addr})

	tx.HashTx()

	return tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.In) == 1 && tx.In[0].Out == -1
}

func (in *TxInput) CanUnlockInput(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
