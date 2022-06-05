package main

import (
	"fmt"
	"log"

	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/transaction"
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

	fmt.Printf("Block: %v", block)
	fmt.Printf("Blockchain: %v", BlockChain)
	fmt.Printf("Block Height: %v \n", BlockChain.BlockHeight)
}
