package blockchain

import (
	"github.com/dgraph-io/badger"
	"github.com/oranges0da/go-blockchain/block"
)

type Blockchain struct {
	LastHash []byte // hash of last block
	Database *badger.DB
}

func New() *Blockchain {
	db, err := badger.Open(badger.DefaultOptions("/tmp/blocks"))

}

func (chain *Blockchain) AddBlock(data string) {
	block := block.New(chain.blocks[len(chain.blocks)-1].BlockID, data)

	chain.blocks = append(chain.blocks, block)
	chain.LastHash = block.Hash
}

func (chain *Blockchain) GetBlocks() []*block.Block {
	return chain.blocks
}

func (chain *Blockchain) GetBlock(blockId int) *block.Block {
	var block *block.Block

	for _, v := range chain.blocks {
		if v.BlockID == blockId {
			block = v
			break
		}
	}

	return block
}
