package main

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		pow := blockchain.InitPow(block)
		blockchain.RunProofWork(chain)
	}

	fmt.Printf("Finished hashing all blocks.\n")
}
