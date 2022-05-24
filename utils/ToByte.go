package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/transaction"
)

// to convert number (such as id) or block to byte array
func ToByte[T int64 | *block.Block | *transaction.Transaction](data T) []byte {
	var res bytes.Buffer

	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(data)

	Handle(err)

	return res.Bytes()
}
