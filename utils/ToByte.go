package utils

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/oranges0da/goblockchain/block"
)

// to convert number (such as id) or block to byte array
func ToByte[T int64 | *block.Block](data T) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, data)

	if err != nil {
		log.Fatalf("Error encoding data to byte: %s", err)
	}

	return buff.Bytes()
}
