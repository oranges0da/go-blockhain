package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// to convert number (such as id) to byte array
func ToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Fatalf("Error encoding data to byte: %s", err)
	}

	return buff.Bytes()
}
