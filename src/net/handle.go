// main module for handling requests/responses from other peers on network

package net

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/oranges0da/goblockchain/src/block_utils"
	"github.com/oranges0da/goblockchain/src/chain"
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

func HandleBlock(request []byte, chain *chain.Blockchain) {
	var buff bytes.Buffer
	var payload Block

	buff.Write(request[cmdLen:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	block := block_utils.ToBlock(blockData)

	fmt.Println("Recevied a new block!")
	chain.AddBlock(block)

	fmt.Printf("Added block %x\n", block.Hash)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		AskData(payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	}
}

func HandleInv(request []byte, chain *chain.Blockchain) {
	var buff bytes.Buffer
	var payload Inv

	buff.Write(request[cmdLen:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Recevied inventory with %d %s\n", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		AskData(payload.AddrFrom, "block", blockHash)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if bytes.Compare(b, blockHash) != 0 {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}

	if payload.Type == "tx" {
		txID := payload.Items[0]

		if memPool[hex.EncodeToString(txID)].ID == nil {
			AskData(payload.AddrFrom, "tx", txID)
		}
	}
}

func HandleGetBlocks(request []byte, chain *chain.Blockchain) {
	var buff bytes.Buffer
	var payload GetBlocks

	buff.Write(request[cmdLen:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := block_utils.GetBlockHashes()
	SendInv(payload.AddrFrom, "block", blocks)
}

func HandleAskBlocks(request []byte, chain *chain.Blockchain) {
	var buff bytes.Buffer
	var payload GetBlocks

	buff.Write(request[cmdLen:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := block_utils.GetBlockHashes()
	SendInv(payload.AddrFrom, "block", blocks)
}
