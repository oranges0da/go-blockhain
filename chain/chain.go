package chain

import (
	"fmt"
	"runtime"

	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/db"
	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/utils"
	"github.com/xujiajun/nutsdb"
)

const (
	bucket = "root"
)

type Blockchain struct {
	LastHash    []byte
	BlockHeight int
}

// address that first transaction must take place
func New(address string) (*Blockchain, error) {
	if db.DBExists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	db, err := db.OpenDB()
	handle.Handle(err, "blockchain")
	defer db.Close()

	// create genesis block
	genesis := block.Genesis(address)

	err = db.Update(func(tx *nutsdb.Tx) error {
		// serialize genesis block
		genesis_id := utils.ToByte(genesis.BlockID)
		byte_genesis := utils.ToByte(genesis)

		//write genesis block to db
		err := tx.Put(bucket, genesis_id, byte_genesis, 0)

		return err
	})
	handle.Handle(err, "blockchain")

	// return blockchain with the genesis hash
	blockchain := &Blockchain{
		LastHash:    genesis.Hash,
		BlockHeight: 0,
	}

	return blockchain, err
}

func (chain *Blockchain) AddBlock(block *model.Block) error {
	// set lastHash to block hash and increment blockHeight
	chain.LastHash = block.Hash
	chain.BlockHeight = block.BlockID

	// serialize block
	byte_block := utils.ToByte(block)

	// open db
	db, err := db.OpenDB()
	handle.Handle(err, "blockchain")

	// add block to db, with its hash as key
	err = db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Put("root", block.Hash, byte_block, 1); err != nil {
			handle.Handle(err, "blockchain")
		}

		return nil
	})

	return err
}
