package tx

import (
	"bytes"

	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/handle"
)

type TxOutput struct {
	Value      int    // amt of satoshis that is being "sent"
	PubKeyHash []byte // hash of public key reciever
}

func NewTxOut(value int, addr string) TxOutput {
	pubKeyHash_unformatted, err := base58.Decode(addr)
	handle.Handle(err, "Error decoding address while creating new TxOutput.")

	pubKeyHash := pubKeyHash_unformatted[1 : len(pubKeyHash_unformatted)-4]

	txOut := TxOutput{
		Value:      value,
		PubKeyHash: pubKeyHash,
	}

	return txOut
}

// checks if output belongs to certain address
func (out *TxOutput) OutCanUnlock(addr string) bool {
	pubKeyData, err := base58.Decode(addr)
	if err != nil {
		panic(err)
	}

	// remove version number and checksum to just get the hash
	pubKeyHash := pubKeyData[1 : len(pubKeyData)-4]

	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

// sets public key hash in output from address provided
func (out *TxOutput) Lock(addr string) {
	pubKeyData, err := base58.Decode(addr)
	if err != nil {
		panic(err)
	}

	pubKeyHash := pubKeyData[1 : len(pubKeyData)-4] // remove version number and checksum to just get the hash

	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsOutLocked(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}
