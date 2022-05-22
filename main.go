package main

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/blockchain"
	"github.com/oranges0da/go-blockchain/proof"
)

func main() {
	BlockChain := blockchain.New()

	BlockChain.AddBlock("First Block") // adding blocks manually to blockchain for now
	BlockChain.AddBlock("Second Block")
	BlockChain.AddBlock("Third Block")

	for _, block := range BlockChain.GetBlocks() {
		pow := proof.New(block)
		nonce, hash := pow.Run() // running proof of work on every block
		fmt.Printf("Block: %d, Nonce: %d, Hash: %x\n", block.BlockID, nonce, hash)
	}

	fmt.Printf("Finished mining all blocks.\n")
}
