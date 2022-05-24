package main

import (
	"fmt"

	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/utils"
)

func main() {
	BlockChain, err := blockchain.New("asdas")
	utils.Handle(err)

	fmt.Printf("%s\n", BlockChain)
}
