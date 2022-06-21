package main

import (
	"fmt"
	"log"

	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/utils"
)

func main() {
	BlockChain, err := blockchain.New("example addrr")

	if err != nil {
		log.Printf("Main blockchain error: %v", err)
	}

	log.Println("Created new blockchain, exited New()")
	log.Printf("Created blockchain :%x", BlockChain)

	blocks, err := utils.GetBlocks()
	utils.Handle(err, "main")

	for _, b := range blocks {
		fmt.Printf("Block: %v\n", b)
	}
}
