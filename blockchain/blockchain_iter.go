package blockchain

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/utils"
)

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

func (iter *BlockchainIter) Next() *block.Block {
	var byte_block []byte

	err := iter.db.View(func(txn *badger.Txn) error {
		var inner_block []byte

		item, err := txn.Get(iter.currentHash)
		utils.Handle(err)

		err = item.Value(func(val []byte) error {
			inner_block := append(inner_block, val...)
			fmt.Printf("%s\n", inner_block)

			return nil
		})

		byte_block := inner_block
		fmt.Printf("%s\n", byte_block)

		return err
	})

	utils.Handle(err)

	block := utils.ToBlock(byte_block)

	return block
}
