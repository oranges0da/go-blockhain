package utils

import (
	"fmt"

	"github.com/oranges0da/go-blockchain/blockchain"
)

func PrintBlockchain(chain *blockchain.Blockchain) {
	fmt.Printf("Blockchain: %x", *chain)
}
