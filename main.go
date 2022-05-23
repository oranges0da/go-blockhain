package main

import (
	"fmt"

	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/utils"
)

func main() {
	BlockChain, err := blockchain.New("askldfj")
	if err != nil {
		utils.Handle(err)
	}

	fmt.Printf("Finished mining all blocks.\n")
}
