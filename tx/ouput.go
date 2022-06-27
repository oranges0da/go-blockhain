package tx

import (
	"bytes"

	"github.com/oranges0da/goblockchain/hash_utils"
)

type TxOutput struct {
	Value      int    // amt of satoshis that is being "sent"
	PubKeyHash []byte // hash of public key reciever
}

func NewTxOut(value int, addr string) TxOutput {
	pubKeyHash := hash_utils.GetPubKeyHash(addr)

	txOut := TxOutput{
		Value:      value,
		PubKeyHash: pubKeyHash,
	}

	return txOut
}

// checks if output belongs to certain address
func (out *TxOutput) OutCanUnlock(addr string) bool {
	pubKeyHash := hash_utils.GetPubKeyHash(addr)

	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

// sets public key hash in output from address provided
func (out *TxOutput) Lock(addr string) {
	pubKeyHash := hash_utils.GetPubKeyHash(addr)

	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsOutLocked(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}
