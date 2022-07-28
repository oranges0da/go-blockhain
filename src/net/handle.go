// main module for handling requests/responses from other peers on network

package net

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func HandleAddr(request []byte) {
	var buff bytes.Buffer
	var payload Addr

	buff.Write(request[cmdLen:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)

	}

	knownNodes = append(knownNodes, payload.AddrList...)
	fmt.Printf("there are %d known nodes\n", len(knownNodes))
	AskAllBlocks()
}
