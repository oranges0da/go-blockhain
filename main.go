package main

import (
	"github.com/oranges0da/go-blockchain/block"
	"github.com/oranges0da/go-blockchain/blockchain"
	"github.com/oranges0da/go-blockchain/utils"
)

func main() {
	blockchain := blockchain.InitBlockchain()

	blockNum := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, blockId := range blockNum {
		newBlock := block.NewBlock(blockId, blockchain.LastHash, "Block #"+string(blockId))

		blockchain.AddBlock(*newBlock)
	}

	utils.PrintBlock(0, blockchain)
}
