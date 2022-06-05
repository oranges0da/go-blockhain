package blockchain

import (
	"fmt"
	"runtime"

	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/utils"
	"github.com/xujiajun/nutsdb"
)

const (
	bucket      = "root"
	dbPath      = "/tmp/blocks"
	dbValue     = "tmp/blocks/MANIFEST"
	genesisText = "Hello, this is the genesis block!"
)

type Blockchain struct {
	LastHash []byte
	Database *nutsdb.DB
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
		byte_genesis := utils.ToByte(genesis)

		//write genesis block to db
		if err := tx.Put(bucket, genesis.Hash, byte_genesis, 0); err != nil {
			utils.Handle(err, "blockchain")
		}
	})
	utils.Handle(err, "blockchain")

	blockchain := &Blockchain{
		LastHash: genesis.Hash,
		Database: db,
	}

	return blockchain, err
}
