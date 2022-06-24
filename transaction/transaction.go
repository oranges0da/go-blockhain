package transaction

import (
	"bytes"

	"github.com/mr-tron/base58"
)

type Transaction struct {
	ID       []byte // hash of transaction
	Inputs   []TxInput
	Outputs  []TxOutput
	Locktime int
}

type TxInput struct {
	ID     []byte // hash of transaction that is being spent/consumed
	Vout   int    // index of output in the previous transaction that is being spent
	Sig    string // signature of input
	PubKey []byte // pubkey of sender, used to sign and verify signature
}

type TxOutput struct {
	Value      int    // amt of satoshis that is being "sent"
	PubKeyHash []byte // hash of public key reciever
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Vout == -1
}

func (out *TxOutput) OutCanUnlock(addr string) bool {
	pubKeyData, err := base58.Decode(addr)
	if err != nil {
		panic(err)
	}

	pubKeyHash := pubKeyData[1 : len(pubKeyData)-4] // remove version number and checksum to just get the hash

	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

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
