package block_utils

import (
	"github.com/oranges0da/goblockchain/src/db"
	"github.com/oranges0da/goblockchain/src/handle"
	"github.com/oranges0da/goblockchain/src/model"
	"github.com/xujiajun/nutsdb"
)

// get specifc block according to hash
func GetBlock(hash []byte) (model.Block, error) {
	var block model.Block

	// open db
	db, err := db.OpenDB()
	handle.Handle(err, "error opening database while trying to get specific block")

	err = db.View(func(tx *nutsdb.Tx) error {
		if e, err := tx.Get("root", hash); err != nil {
			return err
		} else {
			block = *ToBlock(e.Value)
			return nil
		}
	})

	return block, err
}
