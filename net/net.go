/*
	Main module for sending and requesting data (such as blocks and txs to be added to mempool) from other peers.
*/

package net

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/model"
)

const (
	protocol = "tcp"
	version  = 1
	cmdLen   = 12
)

var (
	nodeAddr        string
	minerAddr       string
	knownNodes      = []string{"localhost:3000"}
	blocksInTransit = [][]byte{}
	memPool         = make(map[string]model.Transaction)
)

// custom structs to define []byte data for each type of data (such as block or tx)
// to be used and parsed between peers

type Addr struct {
	AddrList []string
}

type Block struct {
	AddrFrom string
	Block    []byte
}

type Inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

type Tx struct {
	AddrFrom    string
	Transaction []byte
}

type Version struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

type GetBlocks struct {
	AddrFrom string
}

type GetData struct {
	AddrFrom string
	Type     string
	ID       []byte
}

/*
	SendData() main logic for sending byte data
	to another peer, byte data will be parsed to relevant data type (block, tx, etc)
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
	handle.Handle(err, "Error sending data.")
}
