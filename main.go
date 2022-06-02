package main

import (
	"fmt"

	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/cli"
	"github.com/oranges0da/goblockchain/utils"
)

func main() {
	BlockChain, err := blockchain.New("random address")
	utils.Handle(err)
	fmt.Printf("Blockchain is created: %v\n", BlockChain)

	cli := cli.New()
	cli.Run()
}
