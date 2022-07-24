// main module for sending data to other peers on network

package net

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

/*
	main function for sending byte data to other known peers on network
	byte data will be parsed to relevant data type (block, tx, etc)
	will also add new discovered nodes to known nodes
*/

func SendData(addr string, data []byte) {
	conn, err := net.Dial(protocol, addr)

	if err != nil {
		fmt.Printf("%s is not available\n", addr)
		var updatedNodes []string

		for _, node := range knownNodes {
			if node != addr {
				updatedNodes = append(updatedNodes, node)
			}
		}

		knownNodes = updatedNodes

		return
	}

	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}
