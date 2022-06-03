package blockchain

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/utils"
)

type BlockchainIter struct {
	CurrentHash []byte
	DB          *badger.DB
}

func (chain *Blockchain) NewIter() *BlockchainIter {
	iter := &BlockchainIter{
		CurrentHash: chain.LastHash,
		DB:          chain.Database,
	}

	return iter
}

func (iter *BlockchainIter) Next() *model.Block {
	var block *model.Block

	err := iter.DB.View(func(txn *badger.Txn) error {
		var valCopy []byte

		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			log.Printf("Item error in blockchain iter: %v", err)
		}

		err = item.Value(func(val []byte) error {
			// copy value from badger db to valCopy(byte)
			valCopy = append(valCopy, val...)

			return nil
		})
		utils.Handle(err, "iter")

		return nil
	})
	utils.Handle(err, "iter")

	// set hash to prevHash for next iteration
	iter.CurrentHash = block.PrevHash

	return block
}
