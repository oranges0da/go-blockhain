package main

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/blockchain"
)

func main() {
	BlockChain := blockchain.New()

	BlockChain.AddBlock("First Block")
	BlockChain.AddBlock("Second Block")
	BlockChain.AddBlock("Third Block")

	for _, block := range BlockChain.GetBlocks() {
		fmt.Printf("%v\n", block)
	}

	fmt.Printf("Finished hashing all blocks.\n")
}
