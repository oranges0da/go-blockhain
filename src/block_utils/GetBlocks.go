package block_utils

import (
	"github.com/oranges0da/goblockchain/src/db"
	"github.com/oranges0da/goblockchain/src/handle"
	"github.com/oranges0da/goblockchain/src/model"
	"github.com/xujiajun/nutsdb"
)

func GetBlocks() ([]*model.Block, error) {
	var blocks []*model.Block

	db, err := db.OpenDB()
	handle.Handle(err, "Error opening database (block)")
	defer db.Close()

	err = db.View(func(tx *nutsdb.Tx) error {
		entries, err := tx.GetAll("root")
		handle.Handle(err, "Error getting all entries from db (block)")

		for _, entry := range entries {
			blocks = append(blocks, ToBlock(entry.Value))
		}

		return nil
	})

	return blocks, err
}
