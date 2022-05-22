package main

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/blockchain"
	"github.com/oranges0da/go-blockchain/proof"
)

func main() {
	BlockChain := blockchain.New()

	BlockChain.AddBlock("First Block")
	BlockChain.AddBlock("Second Block")
	BlockChain.AddBlock("Third Block")

	for _, block := range BlockChain.GetBlocks() {
		pow := proof.New(block)
		pow.Run()
	}

	fmt.Printf("Finished hashing all blocks.\n")
}
