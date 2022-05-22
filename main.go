package main

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/block"
	"github.com/oranges0da/go-blockchain/blockchain"
)

func main() {
	blockchain := blockchain.InitBlockchain()

	newBlock := block.NewBlock(1, blockchain.LastHash, "Hello World")

	blockchain.AddBlock(*newBlock)

	newBlock = block.NewBlock(2, blockchain.LastHash, "Hello World 2")

	blockchain.AddBlock(*newBlock)

	blocks := blockchain.GetBlocks()

	for _, block := range blocks {
		fmt.Printf("Block: %x\n", *block)
	}
}
