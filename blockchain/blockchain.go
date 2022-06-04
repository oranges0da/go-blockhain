package blockchain

import (
	"fmt"
	"runtime"

	"github.com/oranges0da/goblockchain/utils"
	"github.com/xujiajun/nutsdb"
)

const (
	dbPath      = "/tmp/blocks"
	dbValue     = "tmp/blocks/MANIFEST"
	genesisText = "Hello, this is the genesis block!"
)

type Blockchain struct {
	LastHash [32]byte
	Database *nutsdb.DB
}

// address that first transaction must take place
func New(address string) (*Blockchain, error) {
	var lastHash [32]byte // hash of last block

	if utils.DBExists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	db, err := utils.OpenDB()

	blockchain := &Blockchain{
		LastHash: lastHash,
		Database: db,
	}

	return blockchain, err
}
