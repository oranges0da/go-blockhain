package blockchain

type Blockchain struct {
	LastHash []byte
	blocks   []*Block
}

func InitBlockchain() *Blockchain {
	genesis := Genesis()

	blockchain := &Blockchain{
		LastHash: genesis.Hash,
		blocks:   []*Block{genesis},
	}

	return blockchain
}

func (chain *Blockchain) AddBlock(block Block) {
	chain.blocks = append(chain.blocks, &block)
}
