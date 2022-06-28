package tx

import (
	"bytes"

	"github.com/oranges0da/goblockchain/hash_utils"
	"github.com/oranges0da/goblockchain/model"
)

type TxOutput struct {
	Value      int    // amt of satoshis that is being "sent"
	PubKeyHash []byte // hash of public key reciever
}

func NewTxOut(value int, addr string) *model.TxOutput {
	pubKeyHash := hash_utils.GetPubKeyHash(addr)

	txOut := &model.TxOutput{
		Value:      value,
		PubKeyHash: pubKeyHash,
	}

	return txOut
}

// checks if output belongs to certain address
func OutCanUnlock(out model.TxOutput, addr string) bool {
	pubKeyHash := hash_utils.GetPubKeyHash(addr)

	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

// sets public key hash in output from address provided
func Lock(out model.TxOutput, addr string) {
	pubKeyHash := hash_utils.GetPubKeyHash(addr)

	out.PubKeyHash = pubKeyHash
}

func IsOutLocked(out model.TxOutput, pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}
