package block

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	BlockID   int
	Nonce     int
	IsGenesis bool
	Data      []byte
	Hash      []byte
}

func (b *Block) GetHash(nonce []byte) []byte {
	concat_data := [][]byte{nonce, b.Data}

	data := bytes.Join(concat_data, []byte{})

	hash := sha256.Sum256(data)

	return hash[:]
}

func CreateBlock(BlockId int, data string) *Block {
	block := &Block{
		BlockID:   BlockId,
		IsGenesis: false,
		Data:      []byte(data),
	}

	hash := block.GetHash([]byte{255})

	block.Hash = hash

	return block
}

func Genesis() *Block {
	block := &Block{
		BlockID:   0,
		IsGenesis: true,
		Data:      []byte("Genesis Block"),
	}

	hash := block.GetHash([]byte{255})

	block.Hash = hash

	return block
}
