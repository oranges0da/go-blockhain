package main

import (
	"github.com/oranges0da/go-blockchain/block"
	"github.com/oranges0da/go-blockchain/blockchain"
	"github.com/oranges0da/go-blockchain/utils"
)

func main() {
	blockchain := blockchain.InitBlockchain()

	newBlock := block.NewBlock(1, blockchain.LastHash, "Hello World")

	blockchain.AddBlock(*newBlock)

	newBlock = block.NewBlock(2, blockchain.LastHash, "Hello World 2")

	blockchain.AddBlock(*newBlock)

	utils.PrintBlockchain(blockchain)
}
