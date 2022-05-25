package blockchain

import (
	"log"
	"runtime"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/utils"
)

const (
	dbPath      = "./tmp/blocks"
	dbValue     = "tmp/blocks/MANIFEST"
	genesisText = "Hello, this is the genesis block!"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

// address that first transaction must take place
func New(address string) (*Blockchain, error) {
	if utils.DBExists() {
		log.Println("Blockchain already exists")
		runtime.Goexit() // do not create new blockchain if dbFile already exists
	}

	var lastHash []byte // hash of last block

	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbValue

	db, err := badger.Open(opts)
	utils.Handle(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		// create genesis block and convert it to byte array
		block_genesis := block.Genesis("example address")
		genesis := utils.ToByte(block_genesis)

		err := txn.Set(block_genesis.Hash, genesis)
		utils.Handle(err)

		// set last hash to genesis hash
		err = txn.Set([]byte("lh"), block_genesis.Hash)
		utils.Handle(err)

		return err
	})

	chain := &Blockchain{
		LastHash: lastHash,
		Database: db,
	}

	return chain, err
}

func (chain *Blockchain) AddBlock(b *block.Block) error {
	block := utils.ToByte(b)

	err := chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(b.Hash, block)
		utils.Handle(err)

		err = txn.Set([]byte("lh"), b.Hash)
		return err
	})

	return err
}
