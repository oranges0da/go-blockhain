package tx

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/utils"
)

type Transaction struct {
	ID       []byte // hash of transaction
	Inputs   []TxInput
	Outputs  []TxOutput
	Locktime int
}

// msg is any string that reciever can put into transaction, and therefore the blockchain, forever
// e.g. "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
func NewCoinbaseTx(addr, msg string) *Transaction {
	// get actual public key from address by decoding and removing version byte and checksum
	pubKeyHash_unformatted, err := base58.Decode(addr)
	handle.Handle(err, "Error decoding address whilst trying to create new coinbase transaction.")

	pubKeyHash := pubKeyHash_unformatted[1 : len(pubKeyHash_unformatted)-4]

	in := TxInput{
		ID:     []byte{},
		Vout:   -1,
		Sig:    []byte(msg),
		PubKey: []byte{00000000000000000000000000000000},
	}
	out := TxOutput{
		Value:      50,
		PubKeyHash: pubKeyHash,
	}

	tx := &Transaction{
		ID:       nil,
		Inputs:   []TxInput{in},
		Outputs:  []TxOutput{out},
		Locktime: 0,
	}

	tx.HashTx()

	return tx
}

func (tx *Transaction) HashTx() {
	byte_tx := utils.ToByte(tx)

	hash := sha256.Sum256(byte_tx)

	tx.ID = hash[:]
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].Vout == -1
}
