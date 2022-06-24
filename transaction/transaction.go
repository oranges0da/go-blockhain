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
	ID     []byte // hash of transaction that is being spent/consumed
	Vout   int    // index of output in transaction that is being spent
	Sig    string // signature of input
	PubKey []byte // pubkey of sender, used to sign and verify signature
}

type TxOutput struct {
	Value      int    // amt of satoshis that is being "sent"
	PubKeyHash []byte // hash of public key reciever
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

	tx.Inputs = append(tx.Inputs, TxInput{Sig: sig, Vout: -1})
	tx.Outputs = append(tx.Outputs, TxOutput{Value: 50, PubKey: addr})

	tx.HashTx()

	return tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Vout == -1
}

func (tx *Transaction) OutCanUnlock(addr string) bool {

}
