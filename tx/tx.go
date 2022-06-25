package tx

import (
	"bytes"

	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/hashing"
)

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

// checks that an address can unclock an input for spending
func (in *TxInput) InCanUnlock(addr string) bool {
	pubKeyHash := hashing.HashPubKey(in.PubKey)
	address := hashing.GetAddress(pubKeyHash)

	return address == addr
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
