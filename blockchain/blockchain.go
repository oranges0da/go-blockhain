package blockchain

import "github.com/oranges0da/go-blockchain/block"

type Blockchain struct {
	LastHash []byte
	blocks   []*block.Block
}

func New() *Blockchain {
	genesis := block.Genesis()

	blockchain := &Blockchain{
		LastHash: genesis.Hash,
		blocks:   []*block.Block{genesis},
	}

	return blockchain
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
