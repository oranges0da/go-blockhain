package utils

import (
	"bytes"
	"encoding/gob"
)

// to data (such as block or int) to byte, for hashing, etc
func ToByte(data any) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)

	Handle(err)

	return buff.Bytes()
}
