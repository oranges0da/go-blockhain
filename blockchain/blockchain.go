package blockchain

import "github.com/oranges0da/go-blockchain/block"

type Blockchain struct {
	LastHash []byte
	blocks   []*block.Block
}

func InitBlockchain() *Blockchain {
	genesis := block.Genesis()

	blockchain := &Blockchain{
		LastHash: genesis.Hash,
		blocks:   []*block.Block{genesis},
	}

	return blockchain
}

func (chain *Blockchain) AddBlock(newBlock block.Block) {
	chain.blocks = append(chain.blocks, &newBlock)
	chain.LastHash = newBlock.Hash
}

func (chain *Blockchain) GetBlock(blockId int) *block.Block {
	for _, block := range chain.blocks {
		if block.BlockID == blockId {
			return block
		}
	}
}
