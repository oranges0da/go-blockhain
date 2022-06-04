package blockchain

import "github.com/xujiajun/nutsdb"

type BlockchainIter struct {
	CurrentHash [32]byte
	DB          *nutsdb.DB
}

func (chain *Blockchain) NewIter() *BlockchainIter {
	iter := &BlockchainIter{
		CurrentHash: chain.LastHash,
		DB:          chain.Database,
	}

	return iter
}
