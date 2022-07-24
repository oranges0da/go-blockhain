package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/src/handle"
	"github.com/oranges0da/goblockchain/src/model"
)

// to data (block, tx or int) to array of bytes
func ToByte[T *model.Transaction | *model.Block | int | string](data T) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)

	handle.Handle(err, "Error converting data to []byte.")

	return buff.Bytes()
}
