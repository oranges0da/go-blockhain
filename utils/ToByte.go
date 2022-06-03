package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/transaction"
)

// to data (such as block or int) to byte, for hashing, etc
func ToByte[T *model.Block | []*transaction.Transaction | int](data T) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)

	Handle(err, "ToByte")

	return buff.Bytes()
}
