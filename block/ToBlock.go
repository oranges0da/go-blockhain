package block

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/model"
)

func ToBlock(data []byte) *model.Block {
	var block model.Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}
