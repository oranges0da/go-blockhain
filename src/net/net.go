/*
	Main network module for sending and requesting data (such as blocks and txs to be added to mempool) from other peers.
*/

package net

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/oranges0da/goblockchain/src/chain"
	"github.com/oranges0da/goblockchain/src/handle"
	"github.com/oranges0da/goblockchain/src/model"
	"github.com/oranges0da/goblockchain/src/utils"
)

const (
	protocol = "tcp"
	version  = 1
	cmdLen   = 12
)

var (
	nodeAddr        string // port of node
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

func handleConnection(conn net.Conn, chain *chain.Blockchain) {
	req, err := ioutil.ReadAll(conn)
	handle.Handle(err, "error parsing tcp request")
	defer conn.Close()

	cmd := utils.ToCmd(req[:cmdLen])
	fmt.Printf("Recieved command: %v", cmd)

	switch cmd {
	case "addr":
		HandleAddr(req)
	case "block":
		HandleBlock(req, chain)
	case "inv":
		HandleInv(req, chain)
	case "getblocks":
		HandleGetBlocks(req, chain)
	case "getdata":
		HandleGetData(req, chain)
	case "tx":
		HandleTx(req, chain)
	case "version":
		HandleVersion(req, chain)
	default:
		fmt.Println("Unknown command")
	}
}

/*
	Main function for starting tcp server.
	Vital to be able to recieve requests from peers
	and hand out according data.
*/

func StartServer(nodeID, minerAddress string, chain *chain.Blockchain) {
	nodeAddr = fmt.Sprintf("localhost:%s", nodeID)
	minerAddr = minerAddress
	ln, err := net.Listen(protocol, nodeAddr)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	if nodeAddr != knownNodes[0] {
		SendVersion(knownNodes[0], chain)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn, chain)

	}
}
