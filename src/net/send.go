// main module for sending data to other peers on network

package net

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/oranges0da/goblockchain/src/chain"
	"github.com/oranges0da/goblockchain/src/model"
	"github.com/oranges0da/goblockchain/src/utils"
)

/*
	main function for sending byte data to other known peers on network using tcp
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

// send block to peer
func SendBlock(addr string, b *model.Block) {
	data := Block{nodeAddr, utils.ToByte(b)}
	payload := utils.ToByte[any](data)
	request := append(CmdToBytes("block"), payload...)

	SendData(addr, request)
}

// send transaction to peer
func SendTx(addr string, tx *model.Transaction) {
	data := Tx{nodeAddr, utils.ToByte(tx)}
	payload := utils.ToByte(data)
	request := append(CmdToBytes("tx"), payload...)

	SendData(addr, request)
}

// send inv to peer
func SendInv(address, kind string, items [][]byte) {
	inventory := Inv{nodeAddr, kind, items}
	payload := utils.ToByte(inventory)
	request := append(CmdToBytes("inv"), payload...)

	SendData(address, request)
}

// send version to peer
func SendVersion(addr string, chain *chain.Blockchain) {
	bestHeight := chain.BlockHeight
	payload := utils.ToByte(Version{version, bestHeight, nodeAddr})

	request := append(CmdToBytes("version"), payload...)

	SendData(addr, request)
}

// request for specific peer to send most recent block
func AskBlocks(addr string) {
	payload := utils.ToByte(GetBlocks{addr})
	request := append(CmdToBytes("getblocks"), payload...)

	SendData(addr, request)
}

// request for data from peer
func AskData(address, kind string, id []byte) {
	payload := utils.ToByte(GetData{nodeAddr, kind, id})
	request := append(CmdToBytes("getdata"), payload...)

	SendData(address, request)
}

func AskAllBlocks() {
	for _, node := range knownNodes {
		AskBlocks(node)
	}
}
