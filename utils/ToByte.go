package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/transaction"
)

// to data (such as block or int) to byte, for hashing, etc
func ToByte[T int64 | *block.Block | *transaction.Transaction | []*transaction.Transaction](data T) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)

	Handle(err)

	return buff.Bytes()
}
