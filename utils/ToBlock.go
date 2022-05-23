package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/block"
)

func ToBlock(data []byte) *block.Block {
	var block block.Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}
