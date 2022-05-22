package utils

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/blockchain"
)

func PrintBlock(blockId int, chain *blockchain.Blockchain) {
	block := chain.GetBlock(blockId)

	fmt.Printf("Block %d: %x\n", blockId, block)
}
