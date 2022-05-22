package blockchain

import (
	"github.com/dgraph-io/badger"
	"github.com/oranges0da/go-blockchain/block"
	"github.com/oranges0da/go-blockchain/utils"
)

const (
	dbPath      = "./tmp/blocks"
	dbFile      = "./tmp/blocks/MANIFEST"
	genesisText = "Hello, this is the genesis block!"
)

type Blockchain struct {
	LastHash []byte // hash of last block
	Database *badger.DB
}

// address that first transaction must take place
func New(address string) *Blockchain {
	var lastHash []byte // hash of last block

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbFile

	db, err := badger.Open(badger.DefaultOptions(opts))
	defer db.Close()
	utils.Handle(err)

	if err := db.View(func(txn *badger.Txn) error{
		item, err := txn.Get([]byte("lh"))
	} { // if nothing in db, create genesis block

	}
)
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
