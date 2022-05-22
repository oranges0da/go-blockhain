package main

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/block"
	"github.com/oranges0da/go-blockchain/blockchain"
)

func main() {
	blockchain := blockchain.InitBlockchain()

	blockNum := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, blockId := range blockNum {
		newBlock := block.NewBlock(blockId, blockchain.LastHash, "Block #"+string(blockId))

		blockchain.AddBlock(*newBlock)
	}

	block1 := blockchain.GetBlock(0)

	fmt.Printf("Block 1: %x\n", block1)
}
