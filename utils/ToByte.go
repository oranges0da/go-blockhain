package utils

import (
	"bytes"
	"encoding/gob"
)

// to data (such as block or int) to byte, for hashing, etc
func ToByte[T any](data T) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)

	Handle(err, "ToByte")

	return buff.Bytes()
}
