package block

import (
	"bytes"
	"crypto/sha256"

	"github.com/oranges0da/goblockchain/transaction"
)

type Block struct {
	PrevHash     []byte
	BlockID      int
	Nonce        int
	Transactions []*transaction.Transaction
	Hash         []byte
}

func (b *Block) GetHash(nonce []byte) []byte {
	concat_data := [][]byte{nonce, b.Transactions}

	data := bytes.Join(concat_data, []byte{})

	hash := sha256.Sum256(data)

	return hash[:]
}

func New(BlockId int, txs []*transaction.Transaction) *Block {
	block := &Block{
		PrevHash:     []byte{},
		BlockID:      BlockId,
		Transactions: txs,
	}

	hash := block.GetHash([]byte{255}) // 255 nonce for now

	block.Hash = hash

	return block
}

func Genesis(to string) *Block { // like New(), but only for genesis block of chain
	block := &Block{
		PrevHash: []byte{0},
		BlockID:  0,
		Nonce:    0,
		Transactions: []*transaction.Transaction
	}

	hash := block.GetHash([]byte{255})

	block.Hash = hash

	return block
}
