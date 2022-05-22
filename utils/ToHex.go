package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return buff.Bytes()
}
