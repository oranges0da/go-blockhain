package main

import (
	"fmt"
	"log"

	block "github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/transaction"
	"github.com/oranges0da/goblockchain/utils"
)

func main() {
	BlockChain, err := blockchain.New("example addrr")

	if err != nil {
		log.Printf("Main blockchain error: %v", err)
	}

	log.Println("Created new blockchain, exited New()")

	block := block.New(int(1), BlockChain.LastHash, []*transaction.Transaction{})
	log.Printf("Creating block: %v", block)

	BlockChain.AddBlock(block)

	blocks, err := utils.GetBlocks()
	utils.Handle(err, "main")

	for _, b := range blocks {
		fmt.Printf("Block: %v\n", b)
	}
}
