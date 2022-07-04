package model

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
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
	Vout   int    // index of output in the previous transaction that is being spent (Vector Output)
	Sig    []byte // signature of input
	PubKey []byte // pubkey of sender, used to sign and verify signature
}

type TxOutput struct {
	Value      int    // amt of satoshis that is being "sent"
	PubKeyHash []byte // hash of public key reciever
}

func (tx *Transaction) Hash() {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(tx)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	byte_tx := buff.Bytes()

	hash := sha256.Sum256(byte_tx)

	tx.ID = hash[:]
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Vout == -1
}

// output tx that does not contain sig for verifying purposes "copying"
func (tx *Transaction) Copy() *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	for _, in := range tx.Inputs {
		inCopy := TxInput{in.ID, in.Vout, nil, nil}

		inputs = append(inputs, inCopy)
	}

	for _, out := range tx.Outputs {
		outCopy := TxOutput{out.Value, out.PubKeyHash}

		outputs = append(outputs, outCopy)
	}

	tx = &Transaction{
		ID:       tx.ID,
		Inputs:   inputs,
		Outputs:  outputs,
		Locktime: tx.Locktime,
	}

	return tx
}

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey) {
	if tx.IsCoinbase() {
		return
	}

	for inID, in := range tx.Inputs {
		r, s, err := ecdsa.Sign(rand.Reader, &privKey, in.ID)
		if err != nil {
			log.Panic(err)
		}

		sig := append(r.Bytes(), s.Bytes()...)

		tx.Inputs[inID].Sig = sig
	}
}
