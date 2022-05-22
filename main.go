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

	fmt.Printf("%+v\n", chain)
}
