package blockchain

import badger "github.com/dgraph-io/badger/v3"

type BlockchainIter struct {
	currentHash []byte
	db          *badger.DB
}

func (chain *Blockchain) NewIter() *BlockchainIter {
	iter := &BlockchainIter{
		currentHash: chain.LastHash,
		db:          chain.Database,
	}

	return iter
}
