package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/transaction"
)

type Block struct {
	PrevHash     []byte
	BlockID      int
	Nonce        int
	Transactions []*transaction.Transaction
	Hash         []byte
}

func (b *Block) HashBlock(nonce int) error {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(nonce)

	hash = sha256.Sum256(encoded.Bytes())
	b.Hash = hash[:] // set block hash

	return err
}

func New(BlockId int, prevHash []byte, txs []*transaction.Transaction) *Block {
	block := &Block{
		PrevHash:     []byte{},
		BlockID:      BlockId,
		Transactions: txs,
	}

	block.HashBlock(255) // 255 nonce for now

	return block
}

func Genesis(to string) *Block { // like New(), but only for genesis block of chain
	coinbase := transaction.NewCoinbase(to, "example sig")

	block := &Block{
		PrevHash:     []byte{0},
		BlockID:      0,
		Nonce:        0,
		Transactions: []*transaction.Transaction{coinbase},
	}

	block.HashBlock(255)

	return block
}
