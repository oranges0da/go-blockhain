package block_utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/model"
)

func ToBlock(data []byte) *model.Block {
	var block model.Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	handle.Handle(err, "Error whilst trying to convert to block.")

	return &block
}
