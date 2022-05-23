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
	genesisText = "Hello, this is the genesis block!"
)

type Blockchain struct {
	LastHash []byte
	blocks   []*block.Block
	Database *badger.DB
}

// address that first transaction must take place
func New(address string) *Blockchain {
	if utils.DBExists() {
		log.Println("Blockchain already exists")
		runtime.Goexit() // do not create new blockchain if dbFile already exists
	}

	var lastHash []byte // hash of last block

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	defer db.Close()
	utils.Handle(err)

	db.Update(func(txn *badger.Txn) error {
		// Your code here…
		return nil
	})

	chain := &Blockchain{
		LastHash: lastHash,
		blocks:   []*block.Block{},
	}

	return chain
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
