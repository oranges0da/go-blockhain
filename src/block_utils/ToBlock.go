package block_utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/src/handle"
	"github.com/oranges0da/goblockchain/src/model"
)

func ToBlock(data []byte) *model.Block {
	var block model.Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	handle.Handle(err, "Error whilst trying to convert to block.")

	return &block
}
