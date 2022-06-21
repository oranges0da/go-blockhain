package blockchain

import (
	"fmt"
	"runtime"

	block "github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/transaction"
	"github.com/oranges0da/goblockchain/utils"
	"github.com/xujiajun/nutsdb"
)

const (
	bucket      = "root"
	genesisText = "Hello, this is the genesis block!"
)

type Blockchain struct {
	LastHash    []byte
	BlockHeight int
}

// address that first transaction must take place
func New(address string) (*Blockchain, error) {
	if utils.DBExists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	db, err := utils.OpenDB()
	utils.Handle(err, "blockchain")
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
	utils.Handle(err, "blockchain")

	// return blockchain with the genesis hash
	blockchain := &Blockchain{
		LastHash:    genesis.Hash,
		BlockHeight: 0,
	}

	return blockchain, err
}

func (chain *Blockchain) AddBlock(tx *transaction.Transaction) error {
	block := block.New(chain.BlockHeight+1, chain.LastHash, tx)

	// set lastHash to block hash and increment blockHeight
	chain.LastHash = block.Hash
	chain.BlockHeight = block.BlockID

	// serialize block
	byte_block := utils.ToByte(block)

	// open db
	db, err := utils.OpenDB()
	utils.Handle(err, "blockchain")

	// add block to db, with its hash as key
	err = db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Put("root", block.Hash, byte_block, 1); err != nil {
			utils.Handle(err, "blockchain")
		}

		return nil
	})

	return err
}
