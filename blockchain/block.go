package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	BlockID  int
	PrevHash []byte
	Data     []byte
	Hash     []byte
}

func (b *Block) GetHash() []byte {
	data := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(data)

	return hash[:]
}

func NewBlock(BlockId int, PrevHash []byte, data string) *Block {
	block := &Block{
		BlockID:  BlockId,
		PrevHash: PrevHash,
		Data:     []byte(data),
	}

	hash := block.GetHash()

	block.Hash = hash

	return block
}

func Genesis() *Block {
	block := &Block{
		BlockID:  0,
		PrevHash: []byte{},
		Data:     []byte("Genesis Block"),
	}

	return block
}
