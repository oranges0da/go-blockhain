package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/oranges0da/goblockchain/handle"
)

// to data (such as block or int) to byte, for hashing, etc
func ToByte[T any](data T) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)

	handle.Handle(err, "ToByte")

	return buff.Bytes()
}
