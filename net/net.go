package net

import "github.com/oranges0da/goblockchain/model"

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
