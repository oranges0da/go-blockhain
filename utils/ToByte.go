package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ToByte(num int) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Fatalf("Error encoding data to byte: %s", err)
	}

	return buff.Bytes()
}
