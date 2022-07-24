/*
	Main network module for sending and requesting data (such as blocks and txs to be added to mempool) from other peers.
*/

package net

import (
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
